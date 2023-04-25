CREATE TYPE enum_type_goods AS ENUM ('shoes', 'top secret', 'drink', 'personal protection', 'book');

CREATE TYPE enum_country AS ENUM ('POLAND', 'GERMANY', 'BELARUS', 'UKRAINE', 'USA');

CREATE TYPE enum_fifo_lifo AS ENUM ('fifo', 'lifo');

CREATE TYPE type_id_amount AS (
	id_v int,
	amount_v int
);

CREATE DOMAIN domain_amount AS int NOT NULL CHECK (value >= 0) DEFAULT 0;

CREATE DOMAIN domain_money AS numeric NOT NULL CHECK (value = round(VALUE, 2) AND value >= 0) DEFAULT 0;

CREATE DOMAIN domain_percentage AS numeric NOT NULL CHECK (value >= 0 AND value <= 1);