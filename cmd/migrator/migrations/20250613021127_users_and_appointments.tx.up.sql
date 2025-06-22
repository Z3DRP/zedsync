CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS roles (
	id SERIAL PRIMARY KEY,
	name varchar(255) NOT NULL UNIQUE
);

CREATE INDEX IF NOT EXISTS idx_role_name on roles (name);

CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	uid UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
	avatar VARCHAR(255) NULL,
	email VARCHAR(150) NOT NULL UNIQUE,
	phone varchar(12) NULL,
	username VARCHAR(75) NOT NULL UNIQUE,
	password VARCHAR(255) NOT NULL,
	first_name VARCHAR(100) NOT NULL,
	last_name VARCHAR(100) NOT NULL,
	role_id INTEGER NOT NULL REFERENCES roles(id),
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_usr_uid on users (uid);
CREATE INDEX IF NOT EXISTS idx_usr_usrname on users (username);
CREATE INDEX IF NOT EXISTS idx_usr_email on users (email);

CREATE TABLE IF NOT EXISTS companies (
	id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
	name VARCHAR(255),
	owner_id UUID NOT NULL references users(uid),
	hours_of_operation JSONB NOT NULL,
	business_type VARCHAR(255) NOT NULL,
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_company_owner on companies (owner_id);
CREATE INDEX IF NOT EXISTS idx_company_name on companies (name);

CREATE TABLE IF NOT EXISTS services (
	id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
	name VARCHAR(255) NOT NULL,
	description TEXT NOT NULL,
	cost INTEGER NOT NULL, -- in pennies
	average_duration NUMERIC(3, 2) null,
	image VARCHAR(255) NULL,
	company_id UUID NOT NULL references companies(id),
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_services_name on services (name);
CREATE INDEX IF NOT EXISTS idx_services_company on services (company_id);

CREATE TABLE IF NOT EXISTS appointments (
	id SERIAL PRIMARY KEY,
	scheduled_time timestamptz NOT NULL,
	service_id UUID NOT NULL REFERENCES services(id),
	no_show BOOLEAN NOT NULL DEFAULT false,
	user_id UUID NOT NULL REFERENCES users(uid),
	paid_at_booking BOOLEAN NOT NULL DEFAULT false,
	payment_received BOOLEAN NOT NULL DEFAULT false,
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_appointments_user on appointments (user_id);
CREATE INDEX IF NOT EXISTS idx_appointments_service on appointments (service_id);
CREATE INDEX IF NOT EXISTS idx_appointments_times on appointments (scheduled_time);
