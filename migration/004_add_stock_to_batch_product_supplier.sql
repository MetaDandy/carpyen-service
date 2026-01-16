-- +goose Up
-- Agregar campo stock a la tabla batch_product_supplier

ALTER TABLE batch_product_supplier
ADD COLUMN stock NUMERIC(19, 4) DEFAULT 0;

CREATE INDEX IF NOT EXISTS idx_batch_product_supplier_stock ON batch_product_supplier(stock);

-- +goose Down
-- Revertir el cambio

DROP INDEX IF EXISTS idx_batch_product_supplier_stock;
ALTER TABLE batch_product_supplier DROP COLUMN stock;
