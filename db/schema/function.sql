CREATE OR REPLACE FUNCTION function_create_order(in_country varchar, in_login_vendor varchar, in_warehouse_name varchar,
in_name_goods varchar, in_amount_goods int, in_login_customer varchar) RETURNS varchar AS $$
DECLARE
	err varchar;
	ids_consigment_s type_id_amount ARRAY;
	ids_consigment type_id_amount ARRAY;
	i_v int; a_v int; id_order_v int;
	price_goods_v numeric; name_warehouse_v varchar;
BEGIN
	err = 'inserted data is incorrect or missing';
	--
	INSERT INTO id_vendor_v, d_goods_v
	SELECT id_vendor, id_goods FROM table_consignment
	JOIN table_vendor USING (id_vendor) JOIN table_goods USING (id_goods)
	WHERE table_vendor.login_vendor = in_login_vendor AND table_goods.name_goods = in_name_goods;
	--
  	ids_consigment_s = (SELECT id_consignment, amount_goods_available FROM view_market_all_countries_for_order
					 WHERE country = in_country AND login_vendor = in_login_vendor AND name_goods = in_name_goods
					 ORDER BY CASE sales_model WHEN 'fifo' THEN sales_model ELSE NULL END DESC, 
         			 CASE sales_model WHEN 'fifo' THEN NULL ELSE sales_model END ASC);
	IF NOT FOUND THEN
		RETURN 'No items or invalid request';
	END IF;
	FOREACH i_v, a_v IN ARRAY ids_consigment_s
	LOOP
		IF in_amount_goods > a_v THEN
			in_amount_goods = in_amount_goods - a_v;
			ids_consigment = || (i_v, a_v);
		ELSE
			ids_consigment = || (i_v, a_v);
			EXIT;
		END IF;
	END LOOP;
	IF in_name_goods != 0 THEN
		RETURN 'not enough goods';
	END IF;
	err = 'database error';
	FOREACH i_v, a_v IN ARRAY ids_consigment
	LOOP
		UPDATE view_market_all_countries_for_order
		SET amount_goods_available = amount_goods_available - a_v AND
			amount_goods_blocked = amount_goods_blocked + a_v
		WHERE country = in_country AND login_vendor = in_login_vendor AND 
		name_goods = in_name_goods AND id_consigment = i_v
		RETURNING price_goods, name_warehouse INTO price_goods_v, name_warehouse_v;
		
		IF id_order_v = 0 THEN
			INSERT INTO view_orders_for_order(login_customer, id_consignment, login_vendor, 
			name_goods, name_warehouse, price_goods, amount_goods, delivery_location_country,
			delivery_location_city, date_order_start) VALUES
			(in_login_customer, i_v, in_login_vendor, in_name_goods, name_warehouse_v, price_goods_v, a_v, '', '', now())
			RETURNING id_order INTO id_order_v;
		ELSE
			INSERT INTO view_orders_for_order(id_order, login_customer, id_consignment, login_vendor, 
			name_goods, name_warehouse, price_goods, amount_goods, delivery_location_country,
			delivery_location_city, date_order_start) VALUES
			(id_order_v, in_login_customer, i_v, in_login_vendor, in_name_goods, name_warehouse_v, price_goods_v, a_v, '', '', now());
		END IF;
	END LOOP;
	RETURN 'ok'
	EXCEPTION
  	WHEN OTHERS THEN
		RETURN err;
END;
$$ LANGUAGE PLPGSQL;

SELECT function_create_order('GERMANY', 'vatican_agents', 'lair of Teutonic explorers', 'flat earth photo', 3, 'celt_4_blood_types')

CREATE OR REPLACE FUNCTION function_create_order(in_country varchar, in_login_vendor varchar, in_warehouse_name varchar,
in_name_goods varchar, in_amount_goods int, in_login_customer varchar) RETURNS varchar AS $$
DECLARE
	--err varchar;
	ids_consigment_s type_id_amount ARRAY;
	ids_consigment type_id_amount ARRAY;
	i_v int; a_v int; id_order_v int;
	price_goods_v numeric; id_warehouse_v varchar;
	id_vendor_v int; d_goods_v int;
	id_customer_v int;
BEGIN
	--err = 'inserted data is incorrect or missing';
	--
	SELECT id_vendor, id_goods, id_warehouse 
	INTO id_vendor_v, d_goods_v, id_warehouse_v
	FROM table_consignment
	JOIN table_vendor USING (id_vendor) 
	JOIN table_goods USING (id_goods) 
	JOIN table_warehouse_info USING (id_warehouse)
	WHERE table_vendor.login_vendor = in_login_vendor AND table_goods.name_goods = in_name_goods AND
	table_warehouse_info.name_warehouse = in_warehouse_name;
	--
	id_customer_v = (SELECT id_customer FROM table_customer WHERE login_customer = in_login_customer);
  	--ids_consigment_s = 
	SELECT ARRAY (SELECT ROW (id_consignment, amount_goods_available)::type_id_amount FROM view_market_all_countries_for_order
					 WHERE country = in_country AND login_vendor = in_login_vendor AND name_goods = in_name_goods
					 ORDER BY CASE sales_model WHEN 'fifo' THEN sales_model ELSE NULL END DESC, 
         			 CASE sales_model WHEN 'fifo' THEN NULL ELSE sales_model END ASC) INTO ids_consigment_s;
	IF NOT FOUND THEN
		RETURN 'No items or invalid request';
	END IF;
	FOREACH i_v, a_v IN ARRAY ids_consigment_s
	LOOP
		IF in_amount_goods > a_v THEN
			in_amount_goods = in_amount_goods - a_v;
			ids_consigment =  array_append(ids_consigment, (i_v, a_v)::type_id_amount);
		ELSE
			ids_consigment =  array_append(ids_consigment, (i_v, a_v)::type_id_amount);
			in_amount_goods = in_amount_goods - a_v;
			EXIT;
		END IF;
	END LOOP;
	IF in_amount_goods != 0 THEN
		RETURN 'not enough goods';
	END IF;
	--err = 'database error';
	FOREACH i_v, a_v IN ARRAY ids_consigment
	LOOP
		--
		UPDATE table_consignment
		SET amount_goods_available = amount_goods_available - a_v AND
			amount_goods_blocked = amount_goods_blocked + a_v
		WHERE id_warehouse = id_warehouse_v AND id_vendor = id_vendor_v AND 
		id_goods = id_goods_v AND id_consigment = i_v
		RETURNING price_goods INTO price_goods_v;
		--
		IF id_order_v = 0 THEN
			INSERT INTO table_orders(id_customer, id_consignment, id_vendor, 
			id_goods, id_warehouse, price_goods, amount_goods, delivery_location_country,
			delivery_location_city, date_order_start) VALUES
			(id_customer_v, i_v, id_vendor_v, d_goods_v, id_warehouse_v, price_goods_v, a_v, '', '', now())
			RETURNING id_order INTO id_order_v;
		ELSE
			INSERT INTO table_orders(id_order, id_customer, id_consignment, id_vendor, 
			id_goods, id_warehouse, price_goods, amount_goods, delivery_location_country,
			delivery_location_city, date_order_start) VALUES
			(id_order_v, id_customer_v, i_v, id_vendor_v, d_goods_v, id_warehouse_v, price_goods_v, a_v, '', '', now());
		END IF;
	END LOOP;
	RETURN 'ok';
	--EXCEPTION
  	--WHEN OTHERS THEN
	--	RETURN err;
END;
$$ LANGUAGE PLPGSQL;