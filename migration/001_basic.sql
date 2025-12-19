-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create Enums
CREATE TYPE status_enum AS ENUM ('PENDING', 'ACTIVE', 'APPROVED', 'REJECTED', 'CLOSED', 'INACTIVE');
CREATE TYPE role_enum AS ENUM ('ADMIN', 'DESIGNER', 'SELLER', 'CHIEF_INSTALLER', 'INSTALLER');
CREATE TYPE material_enum AS ENUM ('MUEBLE', 'MADERA', 'METAL', 'PLASTICO', 'VIDRIO');
CREATE TYPE meter_enum AS ENUM ('LINEAL', 'CUADRADO');

-- Users Table
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    phone TEXT,
    address TEXT,
    password TEXT,
    role role_enum,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE INDEX idx_users_deleted_at ON users(deleted_at);

-- Client Table
CREATE TABLE client (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    phone TEXT,
    address TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE INDEX idx_client_deleted_at ON client(deleted_at);

-- Supplier Table
CREATE TABLE supplier (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    contact TEXT,
    phone TEXT,
    email TEXT,
    address TEXT,
    user_id UUID,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_supplier_user FOREIGN KEY(user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE INDEX idx_supplier_deleted_at ON supplier(deleted_at);

-- Material Table
CREATE TABLE material (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    type material_enum,
    unite_messure TEXT,
    unit_price FLOAT8,
    stock INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE INDEX idx_material_deleted_at ON material(deleted_at);

-- Project Table
CREATE TABLE project (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    description TEXT,
    location TEXT,
    state status_enum,
    user_id UUID,
    client_id UUID,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_project_user FOREIGN KEY(user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE SET NULL,
    CONSTRAINT fk_project_client FOREIGN KEY(client_id) REFERENCES client(id) ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE INDEX idx_project_deleted_at ON project(deleted_at);

-- Schedule Table
CREATE TABLE schedule (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title TEXT NOT NULL,
    description TEXT,
    initial_date TIMESTAMP,
    final_date TIMESTAMP,
    estimate_days SMALLINT,
    state status_enum,
    project_id UUID,
    user_id UUID,
    user_assigner_id UUID,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_schedule_project FOREIGN KEY(project_id) REFERENCES project(id) ON UPDATE CASCADE ON DELETE SET NULL,
    CONSTRAINT fk_schedule_user FOREIGN KEY(user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE SET NULL,
    CONSTRAINT fk_schedule_user_assigner FOREIGN KEY(user_assigner_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE INDEX idx_schedule_deleted_at ON schedule(deleted_at);

-- Task Table
CREATE TABLE task (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title TEXT NOT NULL,
    description TEXT,
    status status_enum DEFAULT 'PENDING',
    initial_hour TIMESTAMP,
    final_hour TIMESTAMP,
    schedule_id UUID,
    user_id UUID,
    user_assigner_id UUID,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_task_schedule FOREIGN KEY(schedule_id) REFERENCES schedule(id) ON UPDATE CASCADE ON DELETE SET NULL,
    CONSTRAINT fk_task_user FOREIGN KEY(user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE SET NULL,
    CONSTRAINT fk_task_user_assigner FOREIGN KEY(user_assigner_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE INDEX idx_task_deleted_at ON task(deleted_at);

-- Quote Table
CREATE TABLE quote (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    meter_type meter_enum,
    meter_cost FLOAT8,
    meter_quantity FLOAT8,
    furniture_number INT,
    furniture_cost FLOAT8,
    total_cost FLOAT8,
    state status_enum,
    comments TEXT,
    project_id UUID,
    user_id UUID,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_quote_project FOREIGN KEY(project_id) REFERENCES project(id) ON UPDATE CASCADE ON DELETE SET NULL,
    CONSTRAINT fk_quote_user FOREIGN KEY(user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE INDEX idx_quote_deleted_at ON quote(deleted_at);

-- Design Table
CREATE TABLE design (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    url_render TEXT,
    iluminated_plane TEXT,
    state status_enum,
    comments TEXT,
    quote_id UUID,
    user_id UUID,
    user_assigner_id UUID,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_design_quote FOREIGN KEY(quote_id) REFERENCES quote(id) ON UPDATE CASCADE ON DELETE SET NULL,
    CONSTRAINT fk_design_user FOREIGN KEY(user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE SET NULL,
    CONSTRAINT fk_design_user_assigner FOREIGN KEY(user_assigner_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE INDEX idx_design_deleted_at ON design(deleted_at);

-- ClientObservation Table
CREATE TABLE client_observation (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    commment TEXT,
    project_id UUID,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_client_observation_project FOREIGN KEY(project_id) REFERENCES project(id) ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE INDEX idx_client_observation_deleted_at ON client_observation(deleted_at);

-- ServiceEvaluation Table
CREATE TABLE service_evaluation (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    design_qualification FLOAT4,
    fabrication_quality FLOAT4,
    installation_quality FLOAT4,
    overall_satisfaction FLOAT4,
    comments TEXT,
    project_id UUID,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_service_evaluation_project FOREIGN KEY(project_id) REFERENCES project(id) ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE INDEX idx_service_evaluation_deleted_at ON service_evaluation(deleted_at);

-- MaterialProject Table (Join Table)
CREATE TABLE material_project (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    quantity INT,
    material_id UUID,
    project_id UUID,
    user_id UUID,
    CONSTRAINT fk_material_project_material FOREIGN KEY(material_id) REFERENCES material(id) ON UPDATE CASCADE ON DELETE SET NULL,
    CONSTRAINT fk_material_project_project FOREIGN KEY(project_id) REFERENCES project(id) ON UPDATE CASCADE ON DELETE SET NULL,
    CONSTRAINT fk_material_project_user FOREIGN KEY(user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE SET NULL
);

-- MaterialSupplier Table (Join Table)
CREATE TABLE material_supplier (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    quantity INT,
    unit_price FLOAT8,
    total_cost FLOAT8,
    material_id UUID,
    supplier_id UUID,
    user_id UUID,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_material_supplier_material FOREIGN KEY(material_id) REFERENCES material(id) ON UPDATE CASCADE ON DELETE SET NULL,
    CONSTRAINT fk_material_supplier_supplier FOREIGN KEY(supplier_id) REFERENCES supplier(id) ON UPDATE CASCADE ON DELETE SET NULL,
    CONSTRAINT fk_material_supplier_user FOREIGN KEY(user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE INDEX idx_material_supplier_deleted_at ON material_supplier(deleted_at);

-- +goose Down
DROP TABLE IF EXISTS material_supplier;
DROP TABLE IF EXISTS material_project;
DROP TABLE IF EXISTS service_evaluation;
DROP TABLE IF EXISTS client_observation;
DROP TABLE IF EXISTS design;
DROP TABLE IF EXISTS quote;
DROP TABLE IF EXISTS task;
DROP TABLE IF EXISTS schedule;
DROP TABLE IF EXISTS project;
DROP TABLE IF EXISTS material;
DROP TABLE IF EXISTS supplier;
DROP TABLE IF EXISTS client;
DROP TABLE IF EXISTS users;
DROP TYPE IF EXISTS meter_enum;
DROP TYPE IF EXISTS material_enum;
DROP TYPE IF EXISTS role_enum;
DROP TYPE IF EXISTS status_enum;