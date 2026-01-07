-- +goose Up
-- Migración segura para actualizar tipos NUMERIC a precisión apropiada para decimal.Decimal
-- De NUMERIC(15, 3) a NUMERIC(19, 4) para mejor precisión y rango

-- Material table
ALTER TABLE material
    ALTER COLUMN unit_price TYPE NUMERIC(19, 4) USING unit_price::NUMERIC(19, 4);

-- Product table
ALTER TABLE product
    ALTER COLUMN unit_price TYPE NUMERIC(19, 4) USING unit_price::NUMERIC(19, 4);

-- Quote table
ALTER TABLE quote
    ALTER COLUMN total_cost TYPE NUMERIC(19, 4) USING total_cost::NUMERIC(19, 4);

-- SubQuote table
ALTER TABLE sub_quote
    ALTER COLUMN unit_cost TYPE NUMERIC(19, 4) USING unit_cost::NUMERIC(19, 4),
    ALTER COLUMN unit_quantity TYPE NUMERIC(19, 4) USING unit_quantity::NUMERIC(19, 4),
    ALTER COLUMN total_cost TYPE NUMERIC(19, 4) USING total_cost::NUMERIC(19, 4);

-- BatchMaterialSupplier table
ALTER TABLE batch_material_supplier
    ALTER COLUMN unit_price TYPE NUMERIC(19, 4) USING unit_price::NUMERIC(19, 4),
    ALTER COLUMN total_cost TYPE NUMERIC(19, 4) USING total_cost::NUMERIC(19, 4),
    ALTER COLUMN stock TYPE NUMERIC(19, 4) USING stock::NUMERIC(19, 4);

-- BatchProductSupplier table
ALTER TABLE batch_product_supplier
    ALTER COLUMN unit_price TYPE NUMERIC(19, 4) USING unit_price::NUMERIC(19, 4),
    ALTER COLUMN total_price TYPE NUMERIC(19, 4) USING total_price::NUMERIC(19, 4);

-- BatchProductMaterial table
ALTER TABLE batch_product_material
    ALTER COLUMN unit_price TYPE NUMERIC(19, 4) USING unit_price::NUMERIC(19, 4),
    ALTER COLUMN total_cost TYPE NUMERIC(19, 4) USING total_cost::NUMERIC(19, 4),
    ALTER COLUMN stock TYPE NUMERIC(19, 4) USING stock::NUMERIC(19, 4);

-- ProductMaterial table
ALTER TABLE product_material
    ALTER COLUMN unit_price TYPE NUMERIC(19, 4) USING unit_price::NUMERIC(19, 4),
    ALTER COLUMN total_cost TYPE NUMERIC(19, 4) USING total_cost::NUMERIC(19, 4);

-- ProjectBatchMaterialSupplier table
ALTER TABLE project_batch_material_supplier
    ALTER COLUMN quantity TYPE NUMERIC(19, 4) USING quantity::NUMERIC(19, 4),
    ALTER COLUMN unit_price TYPE NUMERIC(19, 4) USING unit_price::NUMERIC(19, 4),
    ALTER COLUMN total_price TYPE NUMERIC(19, 4) USING total_price::NUMERIC(19, 4);

-- ProjectBatchProductSupplier table
ALTER TABLE project_batch_product_supplier
    ALTER COLUMN unit_price TYPE NUMERIC(19, 4) USING unit_price::NUMERIC(19, 4),
    ALTER COLUMN total_price TYPE NUMERIC(19, 4) USING total_price::NUMERIC(19, 4);

-- ProjectBatchProductMaterial table
ALTER TABLE project_batch_product_material
    ALTER COLUMN quantity TYPE NUMERIC(19, 4) USING quantity::NUMERIC(19, 4),
    ALTER COLUMN unit_price TYPE NUMERIC(19, 4) USING unit_price::NUMERIC(19, 4),
    ALTER COLUMN total_price TYPE NUMERIC(19, 4) USING total_price::NUMERIC(19, 4);

-- +goose Down
-- Revertir a la precisión anterior si es necesario

-- Material table
ALTER TABLE material
    ALTER COLUMN unit_price TYPE NUMERIC(15, 3) USING unit_price::NUMERIC(15, 3);

-- Product table
ALTER TABLE product
    ALTER COLUMN unit_price TYPE NUMERIC(15, 3) USING unit_price::NUMERIC(15, 3);

-- Quote table
ALTER TABLE quote
    ALTER COLUMN total_cost TYPE NUMERIC(15, 3) USING total_cost::NUMERIC(15, 3);

-- SubQuote table
ALTER TABLE sub_quote
    ALTER COLUMN unit_cost TYPE NUMERIC(15, 3) USING unit_cost::NUMERIC(15, 3),
    ALTER COLUMN unit_quantity TYPE NUMERIC(15, 3) USING unit_quantity::NUMERIC(15, 3),
    ALTER COLUMN total_cost TYPE NUMERIC(15, 3) USING total_cost::NUMERIC(15, 3);

-- BatchMaterialSupplier table
ALTER TABLE batch_material_supplier
    ALTER COLUMN unit_price TYPE NUMERIC(15, 3) USING unit_price::NUMERIC(15, 3),
    ALTER COLUMN total_cost TYPE NUMERIC(15, 3) USING total_cost::NUMERIC(15, 3),
    ALTER COLUMN stock TYPE NUMERIC(15, 3) USING stock::NUMERIC(15, 3);

-- BatchProductSupplier table
ALTER TABLE batch_product_supplier
    ALTER COLUMN unit_price TYPE NUMERIC(15, 3) USING unit_price::NUMERIC(15, 3),
    ALTER COLUMN total_price TYPE NUMERIC(15, 3) USING total_price::NUMERIC(15, 3);

-- BatchProductMaterial table
ALTER TABLE batch_product_material
    ALTER COLUMN unit_price TYPE NUMERIC(15, 3) USING unit_price::NUMERIC(15, 3),
    ALTER COLUMN total_cost TYPE NUMERIC(15, 3) USING total_cost::NUMERIC(15, 3),
    ALTER COLUMN stock TYPE NUMERIC(15, 3) USING stock::NUMERIC(15, 3);

-- ProductMaterial table
ALTER TABLE product_material
    ALTER COLUMN unit_price TYPE NUMERIC(15, 3) USING unit_price::NUMERIC(15, 3),
    ALTER COLUMN total_cost TYPE NUMERIC(15, 3) USING total_cost::NUMERIC(15, 3);

-- ProjectBatchMaterialSupplier table
ALTER TABLE project_batch_material_supplier
    ALTER COLUMN quantity TYPE NUMERIC(15, 3) USING quantity::NUMERIC(15, 3),
    ALTER COLUMN unit_price TYPE NUMERIC(15, 3) USING unit_price::NUMERIC(15, 3),
    ALTER COLUMN total_price TYPE NUMERIC(15, 3) USING total_price::NUMERIC(15, 3);

-- ProjectBatchProductSupplier table
ALTER TABLE project_batch_product_supplier
    ALTER COLUMN unit_price TYPE NUMERIC(15, 3) USING unit_price::NUMERIC(15, 3),
    ALTER COLUMN total_price TYPE NUMERIC(15, 3) USING total_price::NUMERIC(15, 3);

-- ProjectBatchProductMaterial table
ALTER TABLE project_batch_product_material
    ALTER COLUMN quantity TYPE NUMERIC(15, 3) USING quantity::NUMERIC(15, 3),
    ALTER COLUMN unit_price TYPE NUMERIC(15, 3) USING unit_price::NUMERIC(15, 3),
    ALTER COLUMN total_price TYPE NUMERIC(15, 3) USING total_price::NUMERIC(15, 3);
