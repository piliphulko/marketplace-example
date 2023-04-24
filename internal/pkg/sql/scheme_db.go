package sql

const SchemeCustomer = `
CREATE TABLE table_customer	
(
	id_customer int GENERATED ALWAYS AS IDENTITY,
	login_customer varchar(32) UNIQUE NOT NULL,
	passwort_customer varchar(32) CHECK length(passwort) > 8,

	CONSTRAINT primery_key_id_customer PRIMARY KEY (id_customer)
);

CREATE TABLE table_customer_wallets
(
	id_customer int,
	money_customer numeric,
	money_blocked_customer numeric,

	CONSTRAINT foreign_key_id_customer FOREIGN KEY (id_customer) REFERENCES table_customer(id_customer)
);
`
const SchemeVendor = `
CREATE TABLE table_vendor
(
	id_vendor int GENERATED ALWAYS AS IDENTITY,
	passwort_vendor varchar(32) CHECK length(passwort) > 8,

	CONSTRAINT primery_key_id_vendor PRIMARY KEY (id_vendor)
);

CREATE TABLE table_vendor_info
(
	id_vendor int,
	name_vendor text,
	details_vendor text,
	
	CONSTRAINT foreign_key_id_vendor FOREIGN KEY (id_vendor) REFERENCES table_vendor(vendor)
);

CREATE TABLE table_vendor_wallets
(
	id_vendor int,
	money_vendor numeric,
	money_blocked_vendor numeric,

	CONSTRAINT foreign_key_id_vendor FOREIGN KEY (id_vendor) REFERENCES table_vendor(id_vendor)
);
`
const SchemeGoods = `
CREATE TABLE table_goods
(
	id_goods int GENERATED ALWAYS AS IDENTITY,
	name_goods varchar(128),
	info_goods text,

	CONSTRAINT primery_key_id_goods PRIMARY KEY (id_goods)
);

CREATE TABLE table_vendor_goods
(
	id_vendor int,
	id_goods int,
	price_goods numeric,
	amount_goods int,
	add_info_goods text,

	CONSTRAINT foreign_key_id_vendor FOREIGN KEY (id_vendor) REFERENCES table_vendor(id_vendor)
	CONSTRAINT foreign_key_id_goods FOREIGN KEY (id_goods) REFERENCES table_goods(id_goods)
);
CREATE 

CREATE TABLE table_customer_goods
(
	id_customer int,
	id_vendor int,
	id_goods int,
	price_goods numeric,
	amount_goods int,
	order_date_goods 

	CONSTRAINT foreign_key_id_customer FOREIGN KEY (id_customer) REFERENCES table_customer(id_customer)
	CONSTRAINT foreign_key_id_vendor FOREIGN KEY (id_vendor) REFERENCES table_vendor(id_vendor)
	CONSTRAINT foreign_key_id_goods FOREIGN KEY (id_goods) REFERENCES table_goods(id_goods)
);
`

/*
CREATE TABLE table_customer
(
	id_customer int GENERATED ALWAYS AS IDENTITY,
	login_customer varchar(32) UNIQUE NOT NULL,
	passwort_customer varchar(32) CHECK length(passwort) > 8,

	CONSTRAINT primery_key_id_customer PRIMARY KEY (id_customer)
);

CREATE TABLE table_customer_wallets
(
	id_customer int,
	money_customer numeric,
	money_blocked_customer numeric,

	CONSTRAINT foreign_key_id_customer FOREIGN KEY (id_customer) REFERENCES table_customer(id_customer)
);

CREATE TABLE table_vendor
(
	id_vendor int GENERATED ALWAYS AS IDENTITY,
	login_vendor varchar(32) UNIQUE NOT NULL,
	passwort_vendor varchar(32) CHECK length(passwort) > 8,

	CONSTRAINT primery_key_id_vendor PRIMARY KEY (id_vendor)
);

CREATE TABLE table_vendor_info
(
	id_vendor int,
	name_vendor text,
	details_vendor text,

	CONSTRAINT foreign_key_id_vendor FOREIGN KEY (id_vendor) REFERENCES table_vendor(vendor)
);

CREATE TABLE table_vendor_wallets
(
	id_vendor int,
	money_vendor numeric,
	money_blocked_vendor numeric,

	CONSTRAINT foreign_key_id_vendor FOREIGN KEY (id_vendor) REFERENCES table_vendor(id_vendor)
);

CREATE TABLE table_goods
(
	id_goods int GENERATED ALWAYS AS IDENTITY,
	name_goods varchar(128),
	info_goods text,

	CONSTRAINT primery_key_id_goods PRIMARY KEY (id_goods)
);

CREATE TABLE table_warehouse
(
	id_warehouse int GENERATED ALWAYS AS IDENTITY,
	name_warehouse varchar(128),
	info_warehouse text,
	location_country_warehouse varchar(128),
	location_city_warehouse varchar(128),

	CONSTRAINT primery_key_id_warehouse PRIMARY KEY (id_warehouse)
);

CREATE TABLE table_consignment
(
	id_consignment int,
	goods_available bool,
	arrival_date_goods TIMESTAMPTZ,
	consignment_info text,

	CONSTRAINT primery_key_id_consignment PRIMARY KEY (id_consignment)
);

CREATE TABLE table_warehouse_goods
(
	id_vendor_goods int,
	id_vendor int,
	id_goods int,
	amount_goods int,
	goods_available bool,
	arrival_date_goods TIMESTAMPTZ,

	CONSTRAINT primery_key_id_vendor_goods PRIMARY KEY (id_vendor_goods)
);

CREATE TABLE table_market_place
(
	id_vendor int,
	id_goods int,
	id_warehouse int,
	price_goods numeric,
	amount_goods_available int,
	add_info_goods text,

	CONSTRAINT foreign_key_id_vendor FOREIGN KEY (id_vendor) REFERENCES table_vendor(id_vendor),
	CONSTRAINT foreign_key_id_goods FOREIGN KEY (id_goods) REFERENCES table_goods(id_goods),
	CONSTRAINT foreign_key_id_warehouse FOREIGN KEY (id_warehouse) REFERENCES table_warehouse (id_warehouse)
);

CREATE TABLE table_error_order
(
	id_error int GENERATED ALWAYS AS IDENTITY,
	error_text text,

	CONSTRAINT primery_key_id_error PRIMARY KEY (id_error)
);

CREATE TABLE table_orders
(
	id_order int GENERATED ALWAYS AS IDENTITY,
	id_customer int NOT NULL,
	id_vendor int NOT NULL,
	id_goods int NOT NULL,
	id_warehouse int NOT NULL,
	price_goods numeric,
	amount_goods int NOT NULL,
	delivery_location_country varchar(128),
	delivery_location_city varchar(128),
	date_order_start TIMESTAMPTZ NOT NULL,
	date_order_finish TIMESTAMPTZ,
	delivery_status_order bool,
	error_order int,

	CONSTRAINT primery_key_id_order PRIMARY KEY (id_order)
	CONSTRAINT foreign_key_id_customer FOREIGN KEY (id_customer) REFERENCES table_customer(id_customer),
	CONSTRAINT foreign_key_id_vendor FOREIGN KEY (id_vendor) REFERENCES table_vendor(id_vendor),
	CONSTRAINT foreign_key_id_goods FOREIGN KEY (id_goods) REFERENCES table_goods(id_goods),
	CONSTRAINT foreign_key_id_warehouse FOREIGN KEY (id_warehouse) REFERENCES table_warehouse (id_warehouse)
	CONSTRAINT foreign_key_error_order FOREIGN KEY (error_order) REFERENCES table_error_order(id_error),
);

*/
