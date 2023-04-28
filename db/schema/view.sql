CREATE OR REPLACE VIEW view_marketplace AS SELECT
    table_warehouse_info.country,
    table_warehouse_info.name_warehouse,
    table_vendor_info.name_vendor,
    table_goods.type_googs,
    table_goods.name_goods,
    table_goods.info_goods,
    table_consignment.amount_goods_available,
    table_vendor_price.price_goods
FROM table_consignment
JOIN table_warehouse_info USING (id_warehouse)
JOIN table_vendor_info USING (id_vendor)
JOIN table_goods USING (id_goods)
JOIN table_vendor_price USING (id_goods)
WHERE table_warehouse_info.country = table_vendor_price.country AND goods_in_stock = true
ORDER BY table_warehouse_info.country DESC, table_goods.type_googs;

CREATE OR REPLACE VIEW view_orders_active AS SELECT
    table_orders.operation_uuid,
    table_orders.id_order,
    table_customer.login_customer,
    table_orders.id_consignment,
    table_vendor_info.name_vendor,
    table_goods.name_goods,
    table_warehouse_info.name_warehouse,
    table_orders.price_goods,
    table_orders.amount_goods,
    table_orders.delivery_location_country,
    table_orders.delivery_location_city,
    table_orders.date_order_start,
    table_orders.date_order_finish
FROM table_orders
JOIN table_customer USING (id_customer)
JOIN table_vendor_info USING (id_vendor)
JOIN table_goods USING (id_goods)
JOIN table_warehouse_info USING (id_warehouse)
WHERE table_orders.delivery_status_order = false AND table_orders.cancellation_order = false
ORDER BY table_orders.operation_uuid

CREATE OR REPLACE VIEW view_orders_closed AS SELECT
    table_orders.operation_uuid,
    table_orders.id_order,
    table_customer.login_customer,
    table_orders.id_consignment,
    table_vendor_info.name_vendor,
    table_goods.name_goods,
    table_warehouse_info.name_warehouse,
    table_orders.price_goods,
    table_orders.amount_goods,
    table_orders.delivery_location_country,
    table_orders.delivery_location_city,
    table_orders.date_order_start,
    table_orders.date_order_finish
FROM table_orders
JOIN table_customer USING (id_customer)
JOIN table_vendor_info USING (id_vendor)
JOIN table_goods USING (id_goods)
JOIN table_warehouse_info USING (id_warehouse)
WHERE table_orders.delivery_status_order = true AND table_orders.cancellation_order = false
ORDER BY table_orders.operation_uuid

CREATE OR REPLACE VIEW view_ledger_active AS 
SELECT 
	ledger_active.operation_uuid,
	sum(ledger_active.money_customer_debit) OVER (PARTITION BY ledger_active.operation_uuid 
		ORDER BY ledger_active.name_vendor) AS accumulative_balance_sheet,
	ledger_active.login_customer,
	ledger_active.money_customer_debit,
	ledger_active.name_vendor,
	ledger_active.money_vendor_credit,
	ledger_active.tax_money_vendor_credit,
	ledger_active.name_warehouse,
	ledger_active.money_warehouse_credit,
	ledger_active.money_system_credit
FROM ( 
	SELECT
		table_ledger.operation_uuid,
		table_customer.login_customer,
		sum(table_ledger.money_customer_debit) AS money_customer_debit,
		table_vendor_info.name_vendor,
		sum(table_ledger.money_vendor_credit) AS money_vendor_credit,
		sum(table_ledger.tax_money_vendor_credit) AS tax_money_vendor_credit,
		table_warehouse_info.name_warehouse,
		sum(table_ledger.money_warehouse_credit) AS money_warehouse_credit,
		sum(table_ledger.money_system_credit) AS money_system_credit
	FROM table_ledger
	JOIN table_customer USING (id_customer)
	JOIN table_vendor_info USING (id_vendor)
	JOIN table_warehouse_info USING (id_warehouse)
	WHERE table_ledger.delivery_status_order = false AND 
		table_ledger.cancellation_pay = false AND
		table_ledger.confirmation_order_and_pay = true
	GROUP BY table_ledger.operation_uuid, table_customer.login_customer, 
		table_vendor_info.name_vendor, table_warehouse_info.name_warehouse
	ORDER BY table_ledger.operation_uuid, table_vendor_info.name_vendor
) ledger_active;

CREATE OR REPLACE VIEW view_ledger_closed AS 
SELECT 
	ledger_active.operation_uuid,
	sum(ledger_active.money_customer_debit) OVER (PARTITION BY ledger_active.operation_uuid 
		ORDER BY ledger_active.name_vendor) AS accumulative_balance_sheet,
	ledger_active.login_customer,
	ledger_active.money_customer_debit,
	ledger_active.name_vendor,
	ledger_active.money_vendor_credit,
	ledger_active.tax_money_vendor_credit,
	ledger_active.name_warehouse,
	ledger_active.money_warehouse_credit,
	ledger_active.money_system_credit
FROM ( 
	SELECT
		table_ledger.operation_uuid,
		table_customer.login_customer,
		sum(table_ledger.money_customer_debit) AS money_customer_debit,
		table_vendor_info.name_vendor,
		sum(table_ledger.money_vendor_credit) AS money_vendor_credit,
		sum(table_ledger.tax_money_vendor_credit) AS tax_money_vendor_credit,
		table_warehouse_info.name_warehouse,
		sum(table_ledger.money_warehouse_credit) AS money_warehouse_credit,
		sum(table_ledger.money_system_credit) AS money_system_credit
	FROM table_ledger
	JOIN table_customer USING (id_customer)
	JOIN table_vendor_info USING (id_vendor)
	JOIN table_warehouse_info USING (id_warehouse)
	WHERE table_ledger.delivery_status_order = false AND 
		table_ledger.cancellation_pay = false AND
		table_ledger.confirmation_order_and_pay = true
	GROUP BY table_ledger.operation_uuid, table_customer.login_customer, 
		table_vendor_info.name_vendor, table_warehouse_info.name_warehouse
	ORDER BY table_ledger.operation_uuid, table_vendor_info.name_vendor
) ledger_active;