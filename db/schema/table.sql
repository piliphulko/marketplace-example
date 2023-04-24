CREATE TABLE table_customer
(
	id_customer int GENERATED ALWAYS AS IDENTITY,
	login_customer varchar(32) UNIQUE NOT NULL,
	passwort_customer varchar(32) CHECK (char_length(passwort_customer) > 7),

	CONSTRAINT primery_key_id_customer PRIMARY KEY (id_customer)
);

CREATE TABLE table_vendor
(
	id_vendor int GENERATED ALWAYS AS IDENTITY,
	login_vendor varchar(32) UNIQUE NOT NULL,
	passwort_vendor varchar(32) CHECK (char_length(passwort_vendor) > 7),

	CONSTRAINT primery_key_id_vendor PRIMARY KEY (id_vendor)
);

CREATE TABLE table_goods
(
	id_goods int GENERATED ALWAYS AS IDENTITY,
	name_goods varchar(128),
	type_goods enum_type_goods,
	info_goods text,

	CONSTRAINT primery_key_id_goods PRIMARY KEY (id_goods)
);

CREATE TABLE table_country_city
(
	country enum_country,
	city varchar(128),
	CONSTRAINT primery_key_city PRIMARY KEY (city)
);

CREATE TABLE table_vendor_price
(
	id_vendor int,
	id_goods int,
	country enum_country,
	price_goods numeric CHECK (price_goods >= 0) NOT NULL,
	sales_model enum_fifo_lifo DEFAULT 'lifo'::enum_fifo_lifo
);

CREATE TABLE table_warehouse
(
	id_warehouse int GENERATED ALWAYS AS IDENTITY,
	login_warehouse varchar(32) UNIQUE NOT NULL,
	passwort_warehouse varchar(32) CHECK (char_length(passwort_warehouse) > 7),

	CONSTRAINT primery_key_id_warehouse PRIMARY KEY (id_warehouse)
);

CREATE TABLE table_warehouse_info
(
	id_warehouse int,
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
	amount_goods_available int DEFAULT 0,
	amount_goods_blocked int DEFAULT 0,
	amount_goods_defect int DEFAULT 0,
	goods_in_stock bool,
	arrival_date_goods TIMESTAMPTZ,
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
	vat numeric CHECK (vat >= 0) NOT NULL,

	CONSTRAINT foreign_key_table_country_city FOREIGN KEY (city) REFERENCES table_country_city(city) ON UPDATE CASCADE
);

CREATE TABLE table_warehouse_commission
(
	id_warehouse int,
	commission_percentage numeric CHECK (commission_percentage >= 0) DEFAULT 0,

	CONSTRAINT foreign_key_id_warehouse FOREIGN KEY (id_warehouse) REFERENCES table_warehouse(id_warehouse)
);

CREATE TABLE table_system_commission
(
	commission_percentage numeric CHECK (commission_percentage >= 0) DEFAULT 0 NOT NULL
);

CREATE TABLE table_customer_wallet
(
	id_customer int NOT NULL,
	amount_money numeric CHECK (amount_money >= 0) DEFAULT 0,
	blocked_money numeric CHECK (blocked_money >= 0) DEFAULT 0,

	CONSTRAINT foreign_key_table_customer FOREIGN KEY (id_customer) REFERENCES table_customer(id_customer)
);

CREATE TABLE table_vendor_wallet
(
	id_vendor int NOT NULL,
	amount_money numeric CHECK (amount_money >= 0) DEFAULT 0,
	blocked_money numeric CHECK (blocked_money >= 0) DEFAULT 0,
	tax_money numeric CHECK (tax_money >= 0) DEFAULT 0,

	CONSTRAINT foreign_key_id_vendor FOREIGN KEY (id_vendor) REFERENCES table_vendor(id_vendor)
);

CREATE TABLE table_warehouse_wallet
(
	id_warehouse int NOT NULL,
	amount_money numeric CHECK (amount_money >= 0) DEFAULT 0,
	blocked_money numeric CHECK (blocked_money >= 0) DEFAULT 0,

	CONSTRAINT foreign_key_id_warehouse FOREIGN KEY (id_warehouse) REFERENCES table_warehouse (id_warehouse)
);

CREATE TABLE table_system_wallet
(
	amount_money numeric CHECK (amount_money >= 0) DEFAULT 0,
	blocked_money numeric CHECK (blocked_money >= 0) DEFAULT 0
);

CREATE TABLE table_orders
(
	id_order int,
	id_customer int NOT NULL,
	id_consignment int NOT NULL,
	id_vendor int NOT NULL,
	id_goods int NOT NULL,
	id_warehouse int NOT NULL,
	price_goods numeric,
	amount_goods int NOT NULL,
	delivery_location_country varchar(128),
	delivery_location_city varchar(128),
	date_order_start TIMESTAMPTZ NOT NULL,
	date_order_finish TIMESTAMPTZ,

	delivery_status_order bool DEFAULT false,
	id_problem int,
	cancellation_order bool DEFAULT false,

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
	money_customer_debit numeric DEFAULT 0,
	id_vendor int NOT NULL,
	money_vendor_credit numeric DEFAULT 0,
	tax_money_vendor_credit numeric DEFAULT 0,
	id_warehouse int NOT NULL,
	money_warehouse_credit numeric DEFAULT 0,
	money_system_credit numeric DEFAULT 0,

	cancellation_pay bool,
	confirmation_order_and_pay bool,

	operation_uuid uuid NOT NULL,

	CONSTRAINT foreign_key_id_order FOREIGN KEY (id_order, id_consignment) REFERENCES table_orders(id_order, id_consignment),
	CONSTRAINT foreign_key_id_customer FOREIGN KEY (id_customer) REFERENCES table_customer(id_customer),
	CONSTRAINT foreign_key_id_vendor FOREIGN KEY (id_vendor) REFERENCES table_vendor(id_vendor),
	CONSTRAINT foreign_key_id_warehouse FOREIGN KEY (id_warehouse) REFERENCES table_warehouse (id_warehouse)
);