# 📋 Guía de Implementación: shopspring/decimal

## 1. ¿Por qué necesitas Decimal?

Tu aplicación maneja valores monetarios (precios de productos, materiales, costos totales) usando `float64`. El problema:

```
float64(0.1) + float64(0.2) ≠ 0.3  // Resultado: 0.30000000000000004
```

Esto causa inconsistencias en cálculos financieros. Con `decimal.Decimal`:
- ✅ Precisión exacta sin redondeos inesperados
- ✅ Representación exacta de valores monetarios
- ✅ Sin errores acumulativos en operaciones múltiples
- ✅ Serialización correcta a JSON/XML y base de datos

---

## 2. Conceptos Fundamentales

### ¿Qué es Decimal?
Un tipo que representa números con punto fijo de precisión arbitraria. En Go, es **inmutable** (como `string` y `int`), no mutable (como `big.Int`).

### Ventajas sobre alternativas:
- **vs float64**: Precisión exacta ✅
- **vs big.Rat**: Mejor para dinero, representación más limpia ✅
- **vs big.Int**: API más simple y segura ✅

### Limitaciones:
- Máximo 2^31 dígitos después del decimal (suficiente para cualquier caso práctico)
- Ligeramente menos performante que float64 (pero es aceptable para operaciones no críticas)

---

## 3. Casos de Uso en Tu Proyecto

Los campos que **DEBES cambiar** de `float64` a `decimal.Decimal`:

| Modelo | Campo | Razón |
|--------|-------|-------|
| `Product` | `UnitPrice` | Precio unitario debe ser exacto |
| `Material` | `UnitPrice` | Precio unitario debe ser exacto |
| `Quote` | `TotalCost` | Costo total crítico (cálculos acumulados) |
| `BatchProductMaterial` | Costos/precios | Si los tiene |
| `BatchMaterialSupplier` | Costos/precios | Si los tiene |
| `Design`, `SubQuote` | Montos | Si tienen campos monetarios |

---

## 4. Anatomía de Decimal.Decimal

```go
// Decimal es una estructura inmutable con campos internos
// NO accedas directamente a los campos
decimal.Decimal{
    // internals - NO tocar
}

// Características principales:
// • Zero-value es seguro: decimal.Decimal{} = 0
// • Todas las operaciones retornan un nuevo Decimal
// • El original nunca se modifica
// • Comparable: == y != funcionan perfectamente
```

---

## 5. Formas de Crear un Decimal

```go
// Opción 1: Desde string (recomendado para valores conocidos)
// Seguro, mantiene precisión exacta
d1, err := decimal.NewFromString("19.99")

// Opción 2: Desde entero
// Útil cuando ya tienes un int
d2 := decimal.NewFromInt(42)

// Opción 3: Desde float64 (⚠️ usar con cuidado)
// Ya contiene los errores de float64, pero es lo que tienes en DB
d3 := decimal.NewFromFloat(19.99)  // Evitar si es posible

// Opción 4: De otro Decimal
d4 := d1  // Copia (es seguro porque es inmutable)

// Opción 5: Cero
d5 := decimal.Zero
```

---

## 6. Operaciones Básicas

```go
// Suma
total := price.Add(tax)

// Resta
discount := original.Sub(reduction)

// Multiplicación
subtotal := unitPrice.Mul(quantity)

// División (necesita especificar precisión)
average := total.Div(count)  // Usa divisor.Exponent() automáticamente

// Más operaciones
absolute := d.Abs()
power := d.Pow(2)  // Elevar a potencia
sqrt := d.Sqrt()
```

---

## 7. Integración con Base de Datos (GORM)

**Decimal ya tiene soporte nativo para:**

- **SQL Driver**: `database/sql` → Implementa `Scan()` y `Value()`
- **JSON**: Implementa `MarshalJSON()` y `UnmarshalJSON()`
- **XML**: Implementa `MarshalXML()` y `UnmarshalXML()`

**En tus modelos:**

```go
type Product struct {
    ID        uuid.UUID
    Name      string
    UnitPrice decimal.Decimal  // ← Cambiar de float64 a decimal.Decimal
    UserID    uuid.UUID
    // ...
}
```

GORM detectará automáticamente que es `decimal.Decimal` y lo manejará correctamente en:
- `INSERT`: Guarda con precisión
- `SELECT`: Lee del DB correctamente
- `JSON`: Serializa/deserializa automáticamente

---

## 8. Integración con JSON API

**Serialización automática:**

```go
// Cuando retornas un modelo con decimal.Decimal en JSON:
product := Product{
    UnitPrice: decimal.NewFromString("19.99"),
}

json.Marshal(product)
// Output: {"UnitPrice": "19.99"}  ✅ String exacto, no float

// Deserialización: GORM + JSON automáticamente:
var product Product
json.Unmarshal([]byte(`{"UnitPrice": "19.99"}`), &product)
// product.UnitPrice ahora es decimal exacto
```

**Nota importante:** Tu API retornará `UnitPrice` como **string en JSON**, no como número. Esto es correcto para dinero porque:
- ✅ Evita pérdida de precisión en cliente
- ✅ JavaScript/navegadores no pierden decimales
- ✅ Es standard en APIs financieras

---

## 9. DTO y Response Layer

En tus capas de respuesta, tienes dos opciones:

### Opción A: Mantener Decimal (Recomendado)
```go
// src/response/product.go
type ProductResponse struct {
    ID        uuid.UUID
    Name      string
    UnitPrice decimal.Decimal  // Se serializa como string
}
```

### Opción B: Convertir a string en DTOs (Si quieres control total)
```go
type ProductResponse struct {
    ID        uuid.UUID
    Name      string
    UnitPrice string  // "19.99" explícitamente
}

// En mapper:
response.UnitPrice = product.UnitPrice.String()
```

---

## 10. Cálculos en Services

**Patrón recomendado:**

```go
// En services/quote.go
func (s *QuoteService) CalculateTotal(items []Item) decimal.Decimal {
    total := decimal.Zero
    
    for _, item := range items {
        unitPrice := item.Price
        quantity := decimal.NewFromInt(int64(item.Quantity))
        total = total.Add(unitPrice.Mul(quantity))
    }
    
    // Aplicar impuestos si necesitas
    taxRate := decimal.NewFromString("0.08875")
    totalWithTax := total.Mul(taxRate.Add(decimal.NewFromFloat(1)))
    
    return totalWithTax.Round(2)  // 2 decimales para moneda
}
```

---

## 11. Conversiones y Comparaciones

```go
// Convertir a otros tipos
str := decimal.NewFromString("19.99").String()  // "19.99"
f64 := decimal.NewFromString("19.99").InexactFloat64()  // 19.99 (con error)
int64 := decimal.NewFromString("100").IntPart()  // 100

// Comparaciones
price1 := decimal.NewFromString("19.99")
price2 := decimal.NewFromString("19.99")

if price1.Equal(price2) {          // true
    // ...
}

if price1.GreaterThan(price2) {    // false
    // ...
}

if price1.LessThan(price2) {       // false
    // ...
}
```

---

## 12. Precisión y Redondeo

```go
// Redondear a N decimales (importante para dinero = 2)
amount := decimal.NewFromString("19.995")
rounded := amount.Round(2)  // "20.00"

// Truncar sin redondear
truncated := amount.Truncate(2)  // "19.99"

// Especificar precisión en división
result := decimal.NewFromInt(10).DivRound(3, 2)  // 3.33
```

---

## 13. Plan de Migración para Tu Proyecto

### Fase 1: Preparación
1. Añadir dependencia: `go get github.com/shopspring/decimal`
2. Identificar todos los campos monetarios

### Fase 2: Modelos
1. Cambiar `float64` → `decimal.Decimal` en modelos
2. Migrations SQL: Los tipos en DB no cambian (NUMERIC/DECIMAL ya los soporta)

### Fase 3: Handlers/DTOs
1. Actualizar structs de request/response
2. GORM manejará serialización automáticamente

### Fase 4: Services/Lógica
1. Cambiar operaciones de `float64` a `decimal.Decimal`
2. Añadir `.Round(2)` donde hagas cálculos

### Fase 5: Testing
1. Verificar casos edge (divisiones, rondeos)
2. Validar precisión en cálculos

---

## 14. Ejemplo Completo en Tu Contexto

```go
// Flujo actual (con float64)
// Cliente pide 3 productos a $19.99 c/u + 8.875% impuesto
// 19.99 * 3 = 59.97 (posible error float)
// 59.97 * 1.08875 = 65.2848... (cliente ve $65.28 o $65.29?)

// Flujo con Decimal
price := decimal.NewFromString("19.99")
quantity := decimal.NewFromInt(3)
taxRate := decimal.NewFromString("1.08875")

subtotal := price.Mul(quantity)        // 59.97
total := subtotal.Mul(taxRate)         // 65.284875
totalRounded := total.Round(2)         // 65.28
// ✅ Exacto, sin sorpresas
```

---

## 15. Errores Comunes a Evitar

❌ **NO:**
```go
price := decimal.NewFromFloat(19.99)  // Contiene error de float
price := decimal.Decimal{...}  // Acceder a internals
a := b  // Preocuparte por mutabilidad - ¡es seguro!
```

✅ **SÍ:**
```go
price, _ := decimal.NewFromString("19.99")  // Exacto
result := price.Round(2)  // Redondear siempre para dinero
value := decimal.Zero  // Inicializar a cero
```

---

## Resumen

La migración a `decimal.Decimal` es **relativamente simple**:
1. Cambiar tipo en modelos (`float64` → `decimal.Decimal`)
2. GORM, JSON y SQL lo manejan automáticamente
3. En cálculos, cambiar operaciones (suma de floats → suma de decimals)
4. Aplicar `.Round(2)` en resultados finales
5. No hay cambios en esquema de DB

---

## Referencias

- **Repositorio**: https://github.com/shopspring/decimal
- **Documentación oficial**: http://godoc.org/github.com/shopspring/decimal
- **Instalación**: `go get github.com/shopspring/decimal`
