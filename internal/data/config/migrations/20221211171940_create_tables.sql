-- setting timezone
SET timezone = 'Asia/Kolkata';
-- installing uuid
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- brand
CREATE TABLE IF NOT EXISTS brand (
	id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
	name VARCHAR(15) UNIQUE NOT NULL,
	description TEXT NOT NULL,
	image_url TEXT,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- category
CREATE TABLE IF NOT EXISTS category (
	id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
	name VARCHAR(30) UNIQUE NOT NULL,
	description TEXT,
	image_url TEXT,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- product
DROP TYPE IF EXISTS replacement_type;
CREATE TYPE replacement_type AS ENUM(
	'day',
	'month'
);
DROP TYPE IF EXISTS warranty_type;
CREATE TYPE warranty_type AS ENUM (
	'day',
	'month',
	'year'
);
CREATE TABLE IF NOT EXISTS product (
	id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
	name TEXT NOT NULL,
	description TEXT,
	available BOOLEAN NOT NULL DEFAULT TRUE,
	mrp INT NOT NULL,
	price INT NOT NULL CHECK (price > 0),
	on_sale BOOLEAN DEFAULT FALSE,
	discount INT NOT NULL CHECK (discount > 0 and discount < 91),
	brand uuid NOT NULL,
	category uuid NOT NULL,
	images _TEXT NOT NULL,
	replacement_period SMALLINT NOT NULL,
	replacement_type replacement_type NOT NULL DEFAULT 'day',
	warranty_period SMALLINT NOT NULL,
	warranty_type warranty_type NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (brand) REFERENCES brand(id),
	FOREIGN KEY (category) REFERENCES category(id)
);




