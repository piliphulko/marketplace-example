\COPY table_country_city (country, city) FROM 'C:\Users\pilip\go\src\github.com\piliphulko\marketplace-example\db\csv\table_country_city.csv' DELIMITER ',' CSV HEADER;
\COPY table_customer (login_customer, passwort_customer) FROM 'C:\Users\pilip\go\src\github.com\piliphulko\marketplace-example\db\csv\table_customer.csv' DELIMITER ',' CSV HEADER;
\COPY table_customer_info (id_customer, delivery_location_country, delivery_location_city) FROM 'C:\Users\pilip\go\src\github.com\piliphulko\marketplace-example\db\csv\table_customer_info.csv' DELIMITER ',' CSV HEADER;
\COPY table_vendor (login_vendor, passwort_vendor) FROM 'C:\Users\pilip\go\src\github.com\piliphulko\marketplace-example\db\csv\table_vendor.csv' DELIMITER ',' CSV HEADER;
\COPY table_vendor_info (id_vendor, name_vendor) FROM 'C:\Users\pilip\go\src\github.com\piliphulko\marketplace-example\db\csv\table_vendor_info.csv' DELIMITER ',' CSV HEADER;
\COPY table_goods (id_vendor, type_goods, name_goods, info_goods) FROM 'C:\Users\pilip\go\src\github.com\piliphulko\marketplace-example\db\csv\table_goods.csv' DELIMITER ',' CSV HEADER;
\COPY table_vendor_price (id_vendor, id_goods, country, price_goods) FROM 'C:\Users\pilip\go\src\github.com\piliphulko\marketplace-example\db\csv\table_vendor_price.csv' DELIMITER ',' CSV HEADER;
\COPY table_warehouse (login_warehouse, passwort_warehouse) FROM 'C:\Users\pilip\go\src\github.com\piliphulko\marketplace-example\db\csv\table_warehouse.csv' DELIMITER ',' CSV HEADER;
\COPY table_warehouse_info (id_warehouse, name_warehouse, country, city) FROM 'C:\Users\pilip\go\src\github.com\piliphulko\marketplace-example\db\csv\table_warehouse_info.csv' DELIMITER ',' CSV HEADER;
\COPY table_problem (problem_text) FROM 'C:\Users\pilip\go\src\github.com\piliphulko\marketplace-example\db\csv\table_problem.csv' DELIMITER ',' CSV HEADER;
\COPY table_consignment (id_warehouse, id_vendor, id_goods, amount_goods_available, amount_goods_defect, goods_in_stock, arrival_date_goods, id_problem) FROM 'C:\Users\pilip\go\src\github.com\piliphulko\marketplace-example\db\csv\table_consignment.csv' DELIMITER ',' CSV HEADER;
\COPY table_tax_plan (country, vat) FROM 'C:\Users\pilip\go\src\github.com\piliphulko\marketplace-example\db\csv\table_tax_plan.csv' DELIMITER ',' CSV HEADER;
\COPY table_warehouse_commission (id_warehouse, commission_percentage) FROM 'C:\Users\pilip\go\src\github.com\piliphulko\marketplace-example\db\csv\table_warehouse_commission.csv' DELIMITER ',' CSV HEADER;
\COPY table_system_commission (commission_percentage) FROM 'C:\Users\pilip\go\src\github.com\piliphulko\marketplace-example\db\csv\table_system_commission.csv' DELIMITER ',' CSV HEADER;
\COPY table_customer_wallet (id_customer, amount_money) FROM 'C:\Users\pilip\go\src\github.com\piliphulko\marketplace-example\db\csv\table_customer_wallet.csv' DELIMITER ',' CSV HEADER;
\COPY table_vendor_wallet (id_vendor) FROM 'C:\Users\pilip\go\src\github.com\piliphulko\marketplace-example\db\csv\table_vendor_wallet.csv' DELIMITER ',' CSV HEADER;
\COPY table_warehouse_wallet (id_warehouse) FROM 'C:\Users\pilip\go\src\github.com\piliphulko\marketplace-example\db\csv\table_warehouse_wallet.csv' DELIMITER ',' CSV HEADER;
\COPY table_system_wallet (id_system) FROM 'C:\Users\pilip\go\src\github.com\piliphulko\marketplace-example\db\csv\table_system_wallet.csv' DELIMITER ',' CSV HEADER;

--\i 'C:/Users/pilip/go/src/github.com/piliphulko/marketplace-example/db/psql_import_csv.sql'