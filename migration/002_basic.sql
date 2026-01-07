-- +goose Up

-- MATERIAL
ALTER TABLE material
ADD COLUMN user_id UUID;

ALTER TABLE material
ADD CONSTRAINT fk_material_user
FOREIGN KEY (user_id) REFERENCES users(id)
ON UPDATE CASCADE
ON DELETE SET NULL;

CREATE INDEX IF NOT EXISTS idx_material_user_id
ON material(user_id);

-- PRODUCT
ALTER TABLE product
ADD COLUMN user_id UUID;

ALTER TABLE product
ADD CONSTRAINT fk_product_user
FOREIGN KEY (user_id) REFERENCES users(id)
ON UPDATE CASCADE
ON DELETE SET NULL;

CREATE INDEX IF NOT EXISTS idx_product_user_id
ON product(user_id);

-- +goose Down

-- PRODUCT
DROP INDEX IF EXISTS idx_product_user_id;
ALTER TABLE product DROP CONSTRAINT fk_product_user;
ALTER TABLE product DROP COLUMN user_id;

-- MATERIAL
DROP INDEX IF EXISTS idx_material_user_id;
ALTER TABLE material DROP CONSTRAINT fk_material_user;
ALTER TABLE material DROP COLUMN user_id;
