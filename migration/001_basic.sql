-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Users Table
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    phone VARCHAR(20),
    address TEXT,
    password VARCHAR(255),
    role VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_users_deleted_at ON users(deleted_at);

-- Client Table
CREATE TABLE IF NOT EXISTS client (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    phone VARCHAR(20),
    address TEXT,
    password VARCHAR(255),
    user_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE INDEX IF NOT EXISTS idx_client_deleted_at ON client(deleted_at);
CREATE INDEX IF NOT EXISTS idx_client_user_id ON client(user_id);

-- Supplier Table
CREATE TABLE IF NOT EXISTS supplier (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    contact VARCHAR(255),
    phone VARCHAR(20),
    email VARCHAR(255) NOT NULL UNIQUE,
    address TEXT,
    user_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE INDEX IF NOT EXISTS idx_supplier_deleted_at ON supplier(deleted_at);
CREATE INDEX IF NOT EXISTS idx_supplier_user_id ON supplier(user_id);

-- Material Table
CREATE TABLE IF NOT EXISTS material (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    type VARCHAR(50),
    unit_measure VARCHAR(50),
    unit_price NUMERIC(19, 4),
    user_id UUID,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE INDEX IF NOT EXISTS idx_material_deleted_at ON material(deleted_at);
CREATE INDEX IF NOT EXISTS idx_material_user_id ON material(user_id);

-- Product Table
CREATE TABLE IF NOT EXISTS product (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    type VARCHAR(50),
    unit_price NUMERIC(19, 4),
    user_id UUID,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE INDEX IF NOT EXISTS idx_product_deleted_at ON product(deleted_at);
CREATE INDEX IF NOT EXISTS idx_product_user_id ON product(user_id);

-- Warehouse Table
CREATE TABLE IF NOT EXISTS warehouse (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    address TEXT,
    description TEXT,
    phone VARCHAR(20),
    user_id UUID,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE INDEX IF NOT EXISTS idx_warehouse_deleted_at ON warehouse(deleted_at);
CREATE INDEX IF NOT EXISTS idx_warehouse_user_id ON warehouse(user_id);

-- Project Table
CREATE TABLE IF NOT EXISTS project (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    location VARCHAR(255),
    state VARCHAR(50),
    user_id UUID NOT NULL,
    client_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE SET NULL,
    FOREIGN KEY (client_id) REFERENCES client(id) ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE INDEX IF NOT EXISTS idx_project_deleted_at ON project(deleted_at);
CREATE INDEX IF NOT EXISTS idx_project_user_id ON project(user_id);
CREATE INDEX IF NOT EXISTS idx_project_client_id ON project(client_id);

-- Quote Table
CREATE TABLE IF NOT EXISTS quote (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    total_cost NUMERIC(19, 4),
    status VARCHAR(50),
    comments TEXT,
    valid_days INTEGER,
    delivery_days INTEGER,
    project_id UUID NOT NULL,
    user_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (project_id) REFERENCES project(id) ON UPDATE CASCADE ON DELETE SET NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE INDEX IF NOT EXISTS idx_quote_deleted_at ON quote(deleted_at);
CREATE INDEX IF NOT EXISTS idx_quote_project_id ON quote(project_id);
CREATE INDEX IF NOT EXISTS idx_quote_user_id ON quote(user_id);

-- SubQuote Table
CREATE TABLE IF NOT EXISTS sub_quote (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    ambient VARCHAR(255),
    unit_cost NUMERIC(19, 4),
    unit_quantity NUMERIC(19, 4),
    unit_type VARCHAR(50),
    total_cost NUMERIC(19, 4),
    status VARCHAR(50),
    description TEXT,
    quote_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (quote_id) REFERENCES quote(id) ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE INDEX IF NOT EXISTS idx_sub_quote_deleted_at ON sub_quote(deleted_at);
CREATE INDEX IF NOT EXISTS idx_sub_quote_quote_id ON sub_quote(quote_id);

-- Design Table
CREATE TABLE IF NOT EXISTS design (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    url_render TEXT,
    iluminated_plane VARCHAR(255),
    state VARCHAR(50),
    comments TEXT,
    quote_id UUID NOT NULL,
    user_id UUID NOT NULL,
    user_assigner_id UUID,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (quote_id) REFERENCES quote(id) ON UPDATE CASCADE ON DELETE SET NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE SET NULL,
    FOREIGN KEY (user_assigner_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE INDEX IF NOT EXISTS idx_design_deleted_at ON design(deleted_at);
CREATE INDEX IF NOT EXISTS idx_design_quote_id ON design(quote_id);
CREATE INDEX IF NOT EXISTS idx_design_user_id ON design(user_id);

-- Schedule Table
CREATE TABLE IF NOT EXISTS schedule (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title VARCHAR(255) NOT NULL,
    description TEXT,
    initial_date TIMESTAMP,
    final_date TIMESTAMP,
    estimate_days SMALLINT,
    state VARCHAR(50),
    project_id UUID NOT NULL,
    user_id UUID NOT NULL,
    user_assigner_id UUID,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (project_id) REFERENCES project(id) ON UPDATE CASCADE ON DELETE SET NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE SET NULL,
    FOREIGN KEY (user_assigner_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE INDEX IF NOT EXISTS idx_schedule_deleted_at ON schedule(deleted_at);
CREATE INDEX IF NOT EXISTS idx_schedule_project_id ON schedule(project_id);
CREATE INDEX IF NOT EXISTS idx_schedule_user_id ON schedule(user_id);

-- Task Table
CREATE TABLE IF NOT EXISTS task (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title VARCHAR(255) NOT NULL,
    description TEXT,
    status VARCHAR(50),
    initial_hour TIMESTAMP,
    final_hour TIMESTAMP,
    schedule_id UUID NOT NULL,
    user_id UUID NOT NULL,
    user_assigner_id UUID,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (schedule_id) REFERENCES schedule(id) ON UPDATE CASCADE ON DELETE SET NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE SET NULL,
    FOREIGN KEY (user_assigner_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE INDEX IF NOT EXISTS idx_task_deleted_at ON task(deleted_at);
CREATE INDEX IF NOT EXISTS idx_task_schedule_id ON task(schedule_id);
CREATE INDEX IF NOT EXISTS idx_task_user_id ON task(user_id);

-- ClientObservation Table
CREATE TABLE IF NOT EXISTS client_observation (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    commment TEXT,
    project_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (project_id) REFERENCES project(id) ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE INDEX IF NOT EXISTS idx_client_observation_deleted_at ON client_observation(deleted_at);
CREATE INDEX IF NOT EXISTS idx_client_observation_project_id ON client_observation(project_id);

-- ServiceEvaluation Table
CREATE TABLE IF NOT EXISTS service_evaluation (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    design_qualification REAL,
    fabrication_quality REAL,
    installation_quality REAL,
    overall_satisfaction REAL,
    comments TEXT,
    project_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (project_id) REFERENCES project(id) ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE INDEX IF NOT EXISTS idx_service_evaluation_deleted_at ON service_evaluation(deleted_at);
CREATE INDEX IF NOT EXISTS idx_service_evaluation_project_id ON service_evaluation(project_id);

-- BatchMaterial Table
CREATE TABLE IF NOT EXISTS batch_material (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    quantity NUMERIC(19, 4),
    unit_price NUMERIC(19, 4),
    total_cost NUMERIC(19, 4),
    stock NUMERIC(19, 4) DEFAULT 0,,
    material_id UUID NOT NULL,
    warehouse_id UUID NOT NULL,
    user_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (material_id) REFERENCES material(id) ON UPDATE CASCADE ON DELETE SET NULL,
    FOREIGN KEY (warehouse_id) REFERENCES warehouse(id) ON UPDATE CASCADE ON DELETE SET NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE INDEX IF NOT EXISTS idx_batch_material_deleted_at ON batch_material(deleted_at);
CREATE INDEX IF NOT EXISTS idx_batch_material_material_id ON batch_material(material_id);
CREATE INDEX IF NOT EXISTS idx_batch_material_warehouse_id ON batch_material(warehouse_id);
CREATE INDEX IF NOT EXISTS idx_batch_material_user_id ON batch_material(user_id);

-- BatchProduct Table
CREATE TABLE IF NOT EXISTS batch_product (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    quantity NUMERIC(19, 4),
    unit_price NUMERIC(19, 4),
    total_price NUMERIC(19, 4),
    stock NUMERIC(19, 4) DEFAULT 0,
    product_id UUID NOT NULL,
    warehouse_id UUID NOT NULL,
    user_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (product_id) REFERENCES product(id) ON UPDATE CASCADE ON DELETE SET NULL,
    FOREIGN KEY (warehouse_id) REFERENCES warehouse(id) ON UPDATE CASCADE ON DELETE SET NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE INDEX IF NOT EXISTS idx_batch_product_deleted_at ON batch_product(deleted_at);
CREATE INDEX IF NOT EXISTS idx_batch_product_product_id ON batch_product(product_id);
CREATE INDEX IF NOT EXISTS idx_batch_product_warehouse_id ON batch_product(warehouse_id);
CREATE INDEX IF NOT EXISTS idx_batch_product_user_id ON batch_product(user_id);

-- BatchProductMaterial Table
CREATE TABLE IF NOT EXISTS batch_product_material (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    quantity NUMERIC(19, 4),
    unit_price NUMERIC(19, 4),
    total_cost NUMERIC(19, 4),
    stock NUMERIC(19, 4),
    product_id UUID NOT NULL,
    warehouse_id UUID NOT NULL,
    user_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (product_id) REFERENCES product(id) ON UPDATE CASCADE ON DELETE SET NULL,
    FOREIGN KEY (warehouse_id) REFERENCES warehouse(id) ON UPDATE CASCADE ON DELETE SET NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE INDEX IF NOT EXISTS idx_batch_product_material_deleted_at ON batch_product_material(deleted_at);
CREATE INDEX IF NOT EXISTS idx_batch_product_material_product_id ON batch_product_material(product_id);
CREATE INDEX IF NOT EXISTS idx_batch_product_material_warehouse_id ON batch_product_material(warehouse_id);
CREATE INDEX IF NOT EXISTS idx_batch_product_material_user_id ON batch_product_material(user_id);

-- ProductMaterial Table
CREATE TABLE IF NOT EXISTS product_material (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    quantity NUMERIC(19, 4),
    unit_price NUMERIC(19, 4),
    total_cost NUMERIC(19, 4),
    batch_product_material_id UUID NOT NULL,
    material_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (batch_product_material_id) REFERENCES batch_product_material(id) ON UPDATE CASCADE ON DELETE SET NULL,
    FOREIGN KEY (material_id) REFERENCES material(id) ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE INDEX IF NOT EXISTS idx_product_material_deleted_at ON product_material(deleted_at);
CREATE INDEX IF NOT EXISTS idx_product_material_batch_product_material_id ON product_material(batch_product_material_id);
CREATE INDEX IF NOT EXISTS idx_product_material_material_id ON product_material(material_id);

-- ProjectBatchMaterial Table
CREATE TABLE IF NOT EXISTS project_batch_material (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    quantity NUMERIC(19, 4),
    unit_price NUMERIC(19, 4),
    total_price NUMERIC(19, 4),
    project_id UUID NOT NULL,
    batch_material_id UUID NOT NULL,
    user_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (project_id) REFERENCES project(id) ON UPDATE CASCADE ON DELETE SET NULL,
    FOREIGN KEY (batch_material_id) REFERENCES batch_material(id) ON UPDATE CASCADE ON DELETE SET NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE INDEX IF NOT EXISTS idx_project_batch_material_project_id ON project_batch_material(project_id);
CREATE INDEX IF NOT EXISTS idx_project_batch_material_batch_material_id ON project_batch_material(batch_material_id);
CREATE INDEX IF NOT EXISTS idx_project_batch_material_user_id ON project_batch_material(user_id);
CREATE INDEX IF NOT EXISTS idx_project_batch_material_deleted_at ON project_batch_material(deleted_at);

-- ProjectBatchProduct Table
CREATE TABLE IF NOT EXISTS project_batch_product (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    quantity NUMERIC(19, 4),
    unit_price NUMERIC(19, 4),
    total_price NUMERIC(19, 4),
    project_id UUID NOT NULL,
    batch_product_id UUID NOT NULL,
    user_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (project_id) REFERENCES project(id) ON UPDATE CASCADE ON DELETE SET NULL,
    FOREIGN KEY (batch_product_id) REFERENCES batch_product(id) ON UPDATE CASCADE ON DELETE SET NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE INDEX IF NOT EXISTS idx_project_batch_product_project_id ON project_batch_product(project_id);
CREATE INDEX IF NOT EXISTS idx_project_batch_product_batch_product_id ON project_batch_product(batch_product_id);
CREATE INDEX IF NOT EXISTS idx_project_batch_product_user_id ON project_batch_product(user_id);
CREATE INDEX IF NOT EXISTS idx_project_batch_product_deleted_at ON project_batch_product(deleted_at);

-- ProjectBatchProductMaterial Table
CREATE TABLE IF NOT EXISTS project_batch_product_material (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    quantity NUMERIC(19, 4),
    unit_price NUMERIC(19, 4),
    total_price NUMERIC(19, 4),
    project_id UUID NOT NULL,
    batch_product_material_id UUID NOT NULL,
    user_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (project_id) REFERENCES project(id) ON UPDATE CASCADE ON DELETE SET NULL,
    FOREIGN KEY (batch_product_material_id) REFERENCES batch_product_material(id) ON UPDATE CASCADE ON DELETE SET NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE INDEX IF NOT EXISTS idx_project_batch_product_material_project_id ON project_batch_product_material(project_id);
CREATE INDEX IF NOT EXISTS idx_project_batch_product_material_batch_product_material_id ON project_batch_product_material(batch_product_material_id);
CREATE INDEX IF NOT EXISTS idx_project_batch_product_material_user_id ON project_batch_product_material(user_id);
CREATE INDEX IF NOT EXISTS idx_project_batch_product_material_deleted_at ON project_batch_product_material(deleted_at);

-- MaterialDetails Table
CREATE TABLE IF NOT EXISTS material_details (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    batch_material_id UUID NOT NULL,
    user_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (batch_material_id) REFERENCES batch_material(id) ON UPDATE CASCADE ON DELETE SET NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE INDEX IF NOT EXISTS idx_material_details_deleted_at ON material_details(deleted_at);
CREATE INDEX IF NOT EXISTS idx_material_details_batch_material_id ON material_details(batch_material_id);
CREATE INDEX IF NOT EXISTS idx_material_details_user_id ON material_details(user_id);

-- ProductDetails Table
CREATE TABLE IF NOT EXISTS product_details (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    batch_product_id UUID NOT NULL,
    user_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (batch_product_id) REFERENCES batch_product(id) ON UPDATE CASCADE ON DELETE SET NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE INDEX IF NOT EXISTS idx_product_details_deleted_at ON product_details(deleted_at);
CREATE INDEX IF NOT EXISTS idx_product_details_batch_product_id ON product_details(batch_product_id);
CREATE INDEX IF NOT EXISTS idx_product_details_user_id ON product_details(user_id);

-- Purchase Table
CREATE TABLE IF NOT EXISTS purchase (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    date TIMESTAMP,
    receipt_number VARCHAR(255),
    material_details_id UUID NOT NULL,
    product_details_id UUID NOT NULL,
    supplier_id UUID NOT NULL,
    user_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (material_details_id) REFERENCES material_details(id) ON UPDATE CASCADE ON DELETE SET NULL,
    FOREIGN KEY (product_details_id) REFERENCES product_details(id) ON UPDATE CASCADE ON DELETE SET NULL,
    FOREIGN KEY (supplier_id) REFERENCES supplier(id) ON UPDATE CASCADE ON DELETE SET NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE INDEX IF NOT EXISTS idx_purchase_deleted_at ON purchase(deleted_at);
CREATE INDEX IF NOT EXISTS idx_purchase_material_details_id ON purchase(material_details_id);
CREATE INDEX IF NOT EXISTS idx_purchase_product_details_id ON purchase(product_details_id);
CREATE INDEX IF NOT EXISTS idx_purchase_supplier_id ON purchase(supplier_id);
CREATE INDEX IF NOT EXISTS idx_purchase_user_id ON purchase(user_id);

-- +goose Down

DROP TABLE IF EXISTS purchase CASCADE;
DROP TABLE IF EXISTS product_details CASCADE;
DROP TABLE IF EXISTS material_details CASCADE;
DROP TABLE IF EXISTS project_batch_product_material CASCADE;
DROP TABLE IF EXISTS project_batch_product CASCADE;
DROP TABLE IF EXISTS project_batch_material CASCADE;
DROP TABLE IF EXISTS product_material CASCADE;
DROP TABLE IF EXISTS batch_product_material CASCADE;
DROP TABLE IF EXISTS batch_product CASCADE;
DROP TABLE IF EXISTS batch_material CASCADE;
DROP TABLE IF EXISTS service_evaluation CASCADE;
DROP TABLE IF EXISTS client_observation CASCADE;
DROP TABLE IF EXISTS task CASCADE;
DROP TABLE IF EXISTS schedule CASCADE;
DROP TABLE IF EXISTS design CASCADE;
DROP TABLE IF EXISTS sub_quote CASCADE;
DROP TABLE IF EXISTS quote CASCADE;
DROP TABLE IF EXISTS project CASCADE;
DROP TABLE IF EXISTS warehouse CASCADE;
DROP TABLE IF EXISTS product CASCADE;
DROP TABLE IF EXISTS material CASCADE;
DROP TABLE IF EXISTS supplier CASCADE;
DROP TABLE IF EXISTS client CASCADE;
DROP TABLE IF EXISTS users CASCADE;
