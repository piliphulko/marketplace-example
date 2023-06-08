CREATE TABLE table_country_city
(
	country enum_country,
	city varchar(128),

	CONSTRAINT primery_key_city PRIMARY KEY (city)
);

CREATE TABLE table_customer
(
	id_customer int GENERATED ALWAYS AS IDENTITY,
	login_customer varchar(32) UNIQUE NOT NULL,
	passwort_customer bytea,

	CONSTRAINT primery_key_id_customer PRIMARY KEY (id_customer)
);

CREATE TABLE table_customer_info
(
	id_customer int UNIQUE,
	delivery_location_country enum_country,
	delivery_location_city varchar(128),

	CONSTRAINT foreign_key_id_customer FOREIGN KEY (id_customer) REFERENCES table_customer(id_customer),
	CONSTRAINT foreign_key_city FOREIGN KEY (delivery_location_city) REFERENCES table_country_city(city) ON UPDATE CASCADE
);

CREATE TABLE table_vendor
(
	id_vendor int GENERATED ALWAYS AS IDENTITY,
	login_vendor varchar(32) UNIQUE NOT NULL,
	passwort_vendor bytea,

	CONSTRAINT primery_key_id_vendor PRIMARY KEY (id_vendor)
);

CREATE TABLE table_vendor_info
(
	id_vendor int UNIQUE,
	name_vendor varchar(128) UNIQUE,

	CONSTRAINT foreign_key_id_vendor FOREIGN KEY (id_vendor) REFERENCES table_vendor(id_vendor)
);

CREATE TABLE table_goods
(
	id_goods int GENERATED ALWAYS AS IDENTITY,
	id_vendor int,
	type_goods enum_type_goods,
	name_goods varchar(128),
	info_goods text,

	UNIQUE(id_vendor, id_goods),
	CONSTRAINT primery_key_id_goods PRIMARY KEY (id_goods),
	CONSTRAINT foreign_key_id_vendor FOREIGN KEY (id_vendor) REFERENCES table_vendor(id_vendor)
);

CREATE TABLE table_vendor_price
(
	id_vendor int,
	id_goods int,
	country enum_country,
	price_goods domain_money,
	sales_model enum_fifo_lifo DEFAULT 'lifo'::enum_fifo_lifo,

	UNIQUE(id_vendor, id_goods, country),
	CONSTRAINT foreign_key_id_vendor FOREIGN KEY (id_vendor) REFERENCES table_vendor(id_vendor),
	CONSTRAINT foreign_key_id_goods FOREIGN KEY (id_goods) REFERENCES table_goods(id_goods)
);

CREATE TABLE table_vendor_price_archive
(
	operation_cdu varchar(6),
	date_change TIMESTAMPTZ,
	id_vendor int,
	id_goods int,
	country enum_country,
	price_goods domain_money,
	sales_model enum_fifo_lifo DEFAULT 'lifo'::enum_fifo_lifo,

	UNIQUE(id_vendor, id_goods, country),
	CONSTRAINT foreign_key_id_vendor FOREIGN KEY (id_vendor) REFERENCES table_vendor(id_vendor),
	CONSTRAINT foreign_key_id_goods FOREIGN KEY (id_goods) REFERENCES table_goods(id_goods)
);

CREATE TABLE table_warehouse
(
	id_warehouse int GENERATED ALWAYS AS IDENTITY,
	login_warehouse varchar(32) UNIQUE NOT NULL,
	passwort_warehouse bytea,

	CONSTRAINT primery_key_id_warehouse PRIMARY KEY (id_warehouse)
);

CREATE TABLE table_warehouse_info
(
	id_warehouse int UNIQUE,
	name_warehouse varchar(128),
	info_warehouse text,
	country enum_country,
	city varchar(128),

	CONSTRAINT foreign_key_id_warehouse FOREIGN KEY (id_warehouse) REFERENCES table_warehouse(id_warehouse),
	CONSTRAINT foreign_key_table_country_city FOREIGN KEY (city) REFERENCES table_country_city(city) ON UPDATE CASCADE
);

CREATE TABLE table_problem
(
	id_problem int GENERATED ALWAYS AS IDENTITY,
	count_problem int DEFAULT 1, 
	problem_text text,

	CONSTRAINT primery_key_id_problem PRIMARY KEY (id_problem)
);

CREATE TABLE table_consignment
(
	id_consignment int GENERATED ALWAYS AS IDENTITY,
	id_warehouse int,
	id_vendor int,
	id_goods int,
	amount_goods_available domain_amount,
	amount_goods_blocked domain_amount,
	amount_goods_defect domain_amount,
	goods_in_stock bool DEFAULT false,
	arrival_date_goods TIMESTAMPTZ,
	date_sold_out TIMESTAMPTZ,
	consignment_info text,
	id_problem int,

	CONSTRAINT primery_key_id_consignment PRIMARY KEY (id_consignment),
	CONSTRAINT foreign_key_id_warehouse FOREIGN KEY (id_warehouse) REFERENCES table_warehouse (id_warehouse),
	CONSTRAINT foreign_key_id_vendor FOREIGN KEY (id_vendor) REFERENCES table_vendor(id_vendor),
	CONSTRAINT foreign_key_id_goods FOREIGN KEY (id_goods) REFERENCES table_goods(id_goods),
	CONSTRAINT foreign_key_id_problem FOREIGN KEY (id_problem) REFERENCES table_problem(id_problem)
);

CREATE TABLE table_tax_plan
(
	country enum_country,
	city varchar(128),
	vat domain_percentage,

	CONSTRAINT foreign_key_table_country_city FOREIGN KEY (city) REFERENCES table_country_city(city) ON UPDATE CASCADE
);

CREATE TABLE table_warehouse_commission
(
	id_warehouse int UNIQUE,
	commission_percentage domain_percentage,

	CONSTRAINT foreign_key_id_warehouse FOREIGN KEY (id_warehouse) REFERENCES table_warehouse(id_warehouse)
);

CREATE TABLE table_system_commission
(
	id_system int UNIQUE DEFAULT 1 CHECK (id_system = 1),
	commission_percentage domain_percentage
);

CREATE TABLE table_customer_wallet
(
	id_customer int UNIQUE NOT NULL,
	amount_money domain_money,
	blocked_money domain_money,

	CONSTRAINT foreign_key_table_customer FOREIGN KEY (id_customer) REFERENCES table_customer(id_customer)
);

CREATE TABLE table_vendor_wallet
(
	id_vendor int UNIQUE NOT NULL,
	amount_money domain_money,
	blocked_money domain_money,
	tax_money domain_money,
	blocked_tax_money domain_money,

	CONSTRAINT foreign_key_id_vendor FOREIGN KEY (id_vendor) REFERENCES table_vendor(id_vendor)
);

CREATE TABLE table_warehouse_wallet
(
	id_warehouse int UNIQUE NOT NULL,
	amount_money domain_money,
	blocked_money domain_money,

	CONSTRAINT foreign_key_id_warehouse FOREIGN KEY (id_warehouse) REFERENCES table_warehouse (id_warehouse)
);

CREATE TABLE table_system_wallet
(
	id_system int UNIQUE DEFAULT 1 CHECK (id_system = 1),
	amount_money domain_money,
	blocked_money domain_money
);

CREATE TABLE table_orders
(
	id_order int, -- seq_id_order
	id_customer int NOT NULL,
	id_consignment int NOT NULL,
	id_vendor int NOT NULL,
	id_goods int NOT NULL,
	id_warehouse int NOT NULL,
	price_goods numeric,
	amount_goods domain_amount,
	delivery_location_country varchar(128),
	delivery_location_city varchar(128),
	date_order_start TIMESTAMPTZ NOT NULL,
	date_order_finish TIMESTAMPTZ,

	delivery_status_order enum_status_order DEFAULT 'unconfirmed order'::enum_status_order,
	id_problem int,

	operation_uuid uuid NOT NULL,

	CONSTRAINT primery_key_id_order_id_consignment PRIMARY KEY (id_order, id_consignment),
	CONSTRAINT foreign_key_id_customer FOREIGN KEY (id_customer) REFERENCES table_customer(id_customer),
	CONSTRAINT foreign_key_id_consignment FOREIGN KEY (id_consignment) REFERENCES table_consignment(id_consignment),
	CONSTRAINT foreign_key_id_vendor FOREIGN KEY (id_vendor) REFERENCES table_vendor(id_vendor),
	CONSTRAINT foreign_key_id_goods FOREIGN KEY (id_goods) REFERENCES table_goods(id_goods),
	CONSTRAINT foreign_key_id_warehouse FOREIGN KEY (id_warehouse) REFERENCES table_warehouse (id_warehouse),
	CONSTRAINT foreign_key_id_problem FOREIGN KEY (id_problem) REFERENCES table_problem(id_problem)
);

CREATE TABLE table_ledger
(
	id_order int NOT NULL,
	id_consignment int NOT NULL,
	id_customer int NOT NULL,
	money_customer_debit domain_money,
	id_vendor int NOT NULL,
	money_vendor_credit domain_money,
	tax_money_vendor_credit domain_money,
	id_warehouse int NOT NULL,
	money_warehouse_credit domain_money,
	money_system_credit domain_money,

	delivery_status_order enum_status_order DEFAULT 'unconfirmed order'::enum_status_order,

	operation_uuid uuid NOT NULL,

	CONSTRAINT foreign_key_id_order FOREIGN KEY (id_order, id_consignment) REFERENCES table_orders(id_order, id_consignment),
	CONSTRAINT foreign_key_id_customer FOREIGN KEY (id_customer) REFERENCES table_customer(id_customer),
	CONSTRAINT foreign_key_id_vendor FOREIGN KEY (id_vendor) REFERENCES table_vendor(id_vendor),
	CONSTRAINT foreign_key_id_warehouse FOREIGN KEY (id_warehouse) REFERENCES table_warehouse (id_warehouse)
);