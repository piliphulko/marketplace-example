CREATE OR REPLACE FUNCTION function_create_order
(
	in_uuid varchar,
	in_login_customer varchar,
	in_country varchar,
	in_warehouse_name varchar,
	in_name_vendor varchar,
	in_name_goods varchar,
	in_amount_goods int
) 
RETURNS varchar AS $$
DECLARE
	ids_consigment_s type_id_amount ARRAY;
	ids_consigment type_id_amount ARRAY;
	i_v int; a_v int; 
	id_order_v int;
	price_goods_v numeric; 
	id_warehouse_v int;
	id_vendor_v int; 
	id_goods_v int;
	id_customer_v int;
	warehouse_commission_percentage domain_percentage = 0;
	tax_vendor_commission_percentage domain_percentage = 0;
	system_commission_percentage domain_percentage = 0;
	money_customer_debit_v domain_money = 0;
	money_customer_debit_net domain_money = 0;
	money_vendor_credit_v domain_money = 0;
	tax_money_vendor_credit_v domain_money = 0;
	money_warehouse_credit_v domain_money = 0;
	money_system_credit_v domain_money = 0;
	delivery_location_country_v varchar;
	delivery_location_city_v varchar;
BEGIN
	--
	SELECT table_consignment.id_vendor, table_consignment.id_goods, 
		table_consignment.id_warehouse, table_vendor_price.price_goods 
	INTO id_vendor_v, id_goods_v, id_warehouse_v, price_goods_v
	FROM table_consignment
	JOIN table_vendor_info USING (id_vendor) 
	JOIN table_goods USING (id_goods) 
	JOIN table_warehouse_info USING (id_warehouse)
	JOIN table_vendor_price USING (id_goods)
	WHERE table_vendor_info.name_vendor = in_name_vendor AND table_goods.name_goods = in_name_goods AND
		  table_warehouse_info.name_warehouse = in_warehouse_name AND 
		  table_warehouse_info.country = table_vendor_price.country;
	--
	IF NOT FOUND THEN
		RETURN 'No items or invalid request'::varchar;
	END IF;
	--
	SELECT id_customer, delivery_location_country, delivery_location_city
	INTO id_customer_v, delivery_location_country_v, delivery_location_city_v
	FROM table_customer
	JOIN table_customer_info USING (id_customer)
	WHERE login_customer = in_login_customer;
	--
	IF NOT FOUND THEN
		RETURN 'No items or invalid request'::varchar;
	END IF;
	--
	SELECT ARRAY (
		SELECT ROW (t1.id_consignment, t1.amount_goods_available)::type_id_amount 
		FROM table_consignment AS t1
		JOIN table_vendor_price AS t2 USING (id_goods)
		JOIN table_warehouse_info AS t3 USING (id_warehouse)
		WHERE t1.id_warehouse = id_warehouse_v AND t1.id_vendor = id_vendor_v AND 
			t1.id_goods = id_goods_v AND t3.country = t2.country
		ORDER BY CASE t2.sales_model WHEN 'fifo' THEN t2.sales_model ELSE NULL END DESC, 
         		 CASE t2.sales_model WHEN 'fifo' THEN NULL ELSE t2.sales_model END ASC
	) INTO ids_consigment_s;
	--
	IF NOT FOUND THEN
		RETURN 'No items or invalid request'::varchar;
	END IF;
	--
	FOREACH i_v, a_v IN ARRAY ids_consigment_s
	LOOP
		IF in_amount_goods > a_v THEN
			in_amount_goods = in_amount_goods - a_v;
			ids_consigment =  array_append(ids_consigment, (i_v, a_v)::type_id_amount);
		ELSE
			ids_consigment =  array_append(ids_consigment, (i_v, in_amount_goods)::type_id_amount);
			in_amount_goods = 0;
			EXIT;
		END IF;
	END LOOP;
	--
	IF in_amount_goods != 0 THEN
		RETURN 'not enough goods'::varchar;
	END IF;
	--
	id_order_v = (SELECT nextval('seq_id_order'));
	--
	warehouse_commission_percentage = (SELECT commission_percentage FROM table_warehouse_commission
									   WHERE id_warehouse = id_warehouse_v);
	tax_vendor_commission_percentage = (SELECT vat FROM table_tax_plan
									    WHERE country = in_country::enum_country);
	system_commission_percentage = (SELECT commission_percentage FROM table_system_commission);
	--
	FOREACH i_v, a_v IN ARRAY ids_consigment
	LOOP
		--
		UPDATE table_consignment
		SET amount_goods_available = amount_goods_available - a_v,
			amount_goods_blocked = amount_goods_blocked + a_v
		WHERE id_warehouse = id_warehouse_v AND 
			  id_vendor = id_vendor_v AND 
			  id_goods = id_goods_v AND 
			  id_consignment = i_v;
		--
		IF NOT FOUND THEN
			RETURN 'problem1'::varchar;
		END IF;
		--
		INSERT INTO table_orders(id_order, id_customer, id_consignment, id_vendor, 
		id_goods, id_warehouse, price_goods, amount_goods, delivery_location_country,
		delivery_location_city, date_order_start, operation_uuid) VALUES
		(id_order_v, id_customer_v, i_v, id_vendor_v, id_goods_v, id_warehouse_v, price_goods_v, a_v, 
		delivery_location_country_v, delivery_location_city_v, now(), in_uuid::uuid);
		--
		IF NOT FOUND THEN
			RETURN 'problem2'::varchar;
		END IF;
		--
		money_customer_debit_v = price_goods_v * a_v;
		--
		tax_money_vendor_credit_v = money_customer_debit_v * tax_vendor_commission_percentage;
		--
		money_customer_debit_net = money_customer_debit_v - tax_money_vendor_credit_v;
		--
		money_warehouse_credit_v = money_customer_debit_net * warehouse_commission_percentage;
		--
		money_system_credit_v = money_customer_debit_net * system_commission_percentage;
		--
		money_vendor_credit_v = money_customer_debit_net - (money_warehouse_credit_v + money_system_credit_v);
		--
		IF money_customer_debit_v != 
			(tax_money_vendor_credit_v + money_warehouse_credit_v + money_system_credit_v + money_vendor_credit_v) THEN
			RETURN 'problem3'::varchar;
		END IF;
		--
		INSERT INTO table_ledger(id_order, id_consignment, id_customer, money_customer_debit,
							id_vendor, money_vendor_credit, tax_money_vendor_credit, 
							id_warehouse, money_warehouse_credit, money_system_credit,
							operation_uuid)
		VALUES (id_order_v, i_v, id_customer_v, money_customer_debit_v, id_vendor_v, money_vendor_credit_v, 
			tax_money_vendor_credit_v, id_warehouse_v, money_warehouse_credit_v, money_system_credit_v, in_uuid::uuid);
		--
		IF NOT FOUND THEN
			RETURN 'problem4'::varchar;
		END IF;
		--
	END LOOP;
	RETURN 'ok'::varchar;
END;
$$ LANGUAGE PLPGSQL;