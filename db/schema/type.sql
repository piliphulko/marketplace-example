CREATE TYPE enum_type_goods AS ENUM ('shoes', 'top secret', 'drink', 'personal protection', 'book');

CREATE TYPE enum_country AS ENUM ('POLAND', 'GERMANY', 'BELARUS', 'UKRAINE', 'USA');

CREATE TYPE enum_fifo_lifo AS ENUM ('fifo', 'lifo');

CREATE TYPE type_id_amount AS (
	id_v int,
	amount_v int
);