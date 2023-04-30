CREATE TYPE enum_country AS ENUM ('POLAND', 'UKRAINE', 'BELARUS');

CREATE TYPE enum_type_goods AS ENUM ('Others', 'CPU', 'GPU', 'Cooling', 'RAM', 'Motherboards');

CREATE TYPE enum_fifo_lifo AS ENUM ('fifo', 'lifo');

CREATE DOMAIN domain_amount AS int NOT NULL CHECK (value >= 0) DEFAULT 0;

CREATE DOMAIN domain_money AS numeric NOT NULL CHECK (value = round(value, 2) AND value >= 0) DEFAULT 0;

CREATE DOMAIN domain_percentage AS numeric NOT NULL CHECK (value >= 0 AND value <= 1);

CREATE TYPE type_id_amount AS (
	id_v int,
	amount_v int
);

CREATE TYPE type_details_ledger AS (
	id1 int,
	credit1 domain_money, 
	credit1_tax domain_money,
	id2 int,
	credit2 domain_money, 
	credit_system domain_money
);