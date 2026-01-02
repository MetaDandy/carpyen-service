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
    unit_price NUMERIC(15, 3),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_material_deleted_at ON material(deleted_at);

-- Product Table
CREATE TABLE IF NOT EXISTS product (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    type VARCHAR(50),
    unit_price NUMERIC(15, 3),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_product_deleted_at ON product(deleted_at);

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
    total_cost NUMERIC(15, 3),
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
    unit_cost NUMERIC(15, 3),
    unit_quantity NUMERIC(15, 3),
    unit_type VARCHAR(50),
    total_cost NUMERIC(15, 3),
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

-- BatchMaterialSupplier Table
CREATE TABLE IF NOT EXISTS batch_material_supplier (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    quantity BIGINT,
    unit_price NUMERIC(15, 3),
    total_cost NUMERIC(15, 3),
    stock NUMERIC(15, 3),
    material_id UUID NOT NULL,
    supplier_id UUID NOT NULL,
    user_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (material_id) REFERENCES material(id) ON UPDATE CASCADE ON DELETE SET NULL,
    FOREIGN KEY (supplier_id) REFERENCES supplier(id) ON UPDATE CASCADE ON DELETE SET NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE INDEX IF NOT EXISTS idx_batch_material_supplier_deleted_at ON batch_material_supplier(deleted_at);
CREATE INDEX IF NOT EXISTS idx_batch_material_supplier_material_id ON batch_material_supplier(material_id);
CREATE INDEX IF NOT EXISTS idx_batch_material_supplier_supplier_id ON batch_material_supplier(supplier_id);
CREATE INDEX IF NOT EXISTS idx_batch_material_supplier_user_id ON batch_material_supplier(user_id);

-- BatchProductSupplier Table
CREATE TABLE IF NOT EXISTS batch_product_supplier (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    quantity BIGINT,
    unit_price NUMERIC(15, 3),
    total_price NUMERIC(15, 3),
    product_id UUID NOT NULL,
    supplier_id UUID NOT NULL,
    user_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (product_id) REFERENCES product(id) ON UPDATE CASCADE ON DELETE SET NULL,
    FOREIGN KEY (supplier_id) REFERENCES supplier(id) ON UPDATE CASCADE ON DELETE SET NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE INDEX IF NOT EXISTS idx_batch_product_supplier_deleted_at ON batch_product_supplier(deleted_at);
CREATE INDEX IF NOT EXISTS idx_batch_product_supplier_product_id ON batch_product_supplier(product_id);
CREATE INDEX IF NOT EXISTS idx_batch_product_supplier_supplier_id ON batch_product_supplier(supplier_id);
CREATE INDEX IF NOT EXISTS idx_batch_product_supplier_user_id ON batch_product_supplier(user_id);

-- BatchProductMaterial Table
CREATE TABLE IF NOT EXISTS batch_product_material (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    quantity BIGINT,
    unit_price NUMERIC(15, 3),
    total_cost NUMERIC(15, 3),
    stock NUMERIC(15, 3),
    product_id UUID NOT NULL,
    user_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (product_id) REFERENCES product(id) ON UPDATE CASCADE ON DELETE SET NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE INDEX IF NOT EXISTS idx_batch_product_material_deleted_at ON batch_product_material(deleted_at);
CREATE INDEX IF NOT EXISTS idx_batch_product_material_product_id ON batch_product_material(product_id);
CREATE INDEX IF NOT EXISTS idx_batch_product_material_user_id ON batch_product_material(user_id);

-- ProductMaterial Table
CREATE TABLE IF NOT EXISTS product_material (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    quantity BIGINT,
    unit_price NUMERIC(15, 3),
    total_cost NUMERIC(15, 3),
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

-- ProjectBatchMaterialSupplier Table
CREATE TABLE IF NOT EXISTS project_batch_material_supplier (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    quantity NUMERIC(15, 3),
    unit_price NUMERIC(15, 3),
    total_price NUMERIC(15, 3),
    project_id UUID NOT NULL,
    batch_material_supplier_id UUID NOT NULL,
    user_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (project_id) REFERENCES project(id) ON UPDATE CASCADE ON DELETE SET NULL,
    FOREIGN KEY (batch_material_supplier_id) REFERENCES batch_material_supplier(id) ON UPDATE CASCADE ON DELETE SET NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE INDEX IF NOT EXISTS idx_project_batch_material_supplier_project_id ON project_batch_material_supplier(project_id);
CREATE INDEX IF NOT EXISTS idx_project_batch_material_supplier_batch_material_supplier_id ON project_batch_material_supplier(batch_material_supplier_id);
CREATE INDEX IF NOT EXISTS idx_project_batch_material_supplier_user_id ON project_batch_material_supplier(user_id);
CREATE INDEX IF NOT EXISTS idx_project_batch_material_supplier_deleted_at ON project_batch_material_supplier(deleted_at);

-- ProjectBatchProductSupplier Table
CREATE TABLE IF NOT EXISTS project_batch_product_supplier (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    quantity BIGINT,
    unit_price NUMERIC(15, 3),
    total_price NUMERIC(15, 3),
    project_id UUID NOT NULL,
    batch_product_supplier_id UUID NOT NULL,
    user_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (project_id) REFERENCES project(id) ON UPDATE CASCADE ON DELETE SET NULL,
    FOREIGN KEY (batch_product_supplier_id) REFERENCES batch_product_supplier(id) ON UPDATE CASCADE ON DELETE SET NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE INDEX IF NOT EXISTS idx_project_batch_product_supplier_project_id ON project_batch_product_supplier(project_id);
CREATE INDEX IF NOT EXISTS idx_project_batch_product_supplier_batch_product_supplier_id ON project_batch_product_supplier(batch_product_supplier_id);
CREATE INDEX IF NOT EXISTS idx_project_batch_product_supplier_user_id ON project_batch_product_supplier(user_id);
CREATE INDEX IF NOT EXISTS idx_project_batch_product_supplier_deleted_at ON project_batch_product_supplier(deleted_at);

-- ProjectBatchProductMaterial Table
CREATE TABLE IF NOT EXISTS project_batch_product_material (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    quantity NUMERIC(15, 3),
    unit_price NUMERIC(15, 3),
    total_price NUMERIC(15, 3),
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

-- +goose Down

DROP TABLE IF EXISTS project_batch_product_material CASCADE;
DROP TABLE IF EXISTS project_batch_product_supplier CASCADE;
DROP TABLE IF EXISTS project_batch_material_supplier CASCADE;
DROP TABLE IF EXISTS product_material CASCADE;
DROP TABLE IF EXISTS batch_product_material CASCADE;
DROP TABLE IF EXISTS batch_product_supplier CASCADE;
DROP TABLE IF EXISTS batch_material_supplier CASCADE;
DROP TABLE IF EXISTS material_project CASCADE;
DROP TABLE IF EXISTS service_evaluation CASCADE;
DROP TABLE IF EXISTS client_observation CASCADE;
DROP TABLE IF EXISTS task CASCADE;
DROP TABLE IF EXISTS schedule CASCADE;
DROP TABLE IF EXISTS design CASCADE;
DROP TABLE IF EXISTS sub_quote CASCADE;
DROP TABLE IF EXISTS quote CASCADE;
DROP TABLE IF EXISTS project CASCADE;
DROP TABLE IF EXISTS product CASCADE;
DROP TABLE IF EXISTS material CASCADE;
DROP TABLE IF EXISTS supplier CASCADE;
DROP TABLE IF EXISTS client CASCADE;
DROP TABLE IF EXISTS users CASCADE;
