CREATE VIEW view_market_all_countries AS
    SELECT table_warehouse_info.country, table_vendor_info.name_vendor, table_goods.name_goods, table_goods.type_goods,
        sum(table_consignment.amount_goods_available) AS amount_goods_available, table_vendor_price.price_goods,
        trunc(avg(table_vendor_price.price_goods) OVER (PARTITION BY table_goods.name_goods), 2) AS avg_price_goods_all_countries
    FROM table_consignment
    JOIN table_vendor_price USING (id_vendor, id_goods)
    JOIN table_warehouse_info USING (id_warehouse)
    JOIN table_vendor_info USING (id_vendor)
    JOIN table_goods USING (id_goods)
    WHERE table_vendor_price.country = table_warehouse_info.country AND goods_in_stock = true
    GROUP BY table_warehouse_info.country, table_vendor.login_vendor, table_goods.name_goods, 
        table_goods.type_goods, table_vendor_price.price_goods;

CREATE VIEW view_market_all_countries_for_order AS
    SELECT table_warehouse_info.country, table_vendor.login_vendor, 
        table_warehouse_info.name_warehouse, table_goods.name_goods,
        table_consignment.amount_goods_available, table_consignment.amount_goods_blocked, 
        table_vendor_price.price_goods, table_consignment.id_consignment,
    table_vendor_price.sales_model
    FROM table_consignment
    JOIN table_vendor_price USING (id_vendor, id_goods)
    JOIN table_vendor USING (id_vendor)
    JOIN table_warehouse_info USING (id_warehouse)
    JOIN table_goods USING (id_goods)
    WHERE table_vendor_price.country = table_warehouse_info.country AND goods_in_stock = true;

CREATE VIEW view_orders_for_order AS
    SELECT table_orders.id_order, table_customer.login_customer, table_orders.id_consignment,
        table_vendor.login_vendor, table_goods.name_goods, table_warehouse_info.name_warehouse, table_orders.price_goods,
        table_orders.amount_goods, table_orders.delivery_location_country, table_orders.delivery_location_city,
        table_orders.date_order_start, table_orders.date_order_finish,
        table_orders.delivery_status_order, table_orders.cancellation_order
    FROM table_orders
    JOIN table_customer USING (id_customer)
    JOIN table_consignment USING (id_warehouse, id_vendor, id_goods)
    JOIN table_vendor USING (id_vendor)
    JOIN table_goods USING (id_goods)
    JOIN table_warehouse_info USING (id_warehouse);