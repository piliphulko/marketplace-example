SELECT function_create_order('GERMANY', 'vatican_agents', 'lair of Teutonic explorers', 'flat earth photo', 3, 'celt_4_blood_types')

CREATE OR REPLACE FUNCTION function_create_order(in_country varchar, in_login_vendor varchar, in_warehouse_name varchar,
in_name_goods varchar, in_amount_goods int, in_login_customer varchar) RETURNS TABLE (err varchar, uuid_v varchar) AS $$
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
	uuid_v uuid;
BEGIN
	--
	SELECT id_vendor, id_goods, id_warehouse, price_goods 
	INTO id_vendor_v, id_goods_v, id_warehouse_v, price_goods_v
	FROM table_consignment
	JOIN table_vendor USING (id_vendor) 
	JOIN table_goods USING (id_goods) 
	JOIN table_warehouse_info USING (id_warehouse)
	JOIN table_vendor_price USING (id_vendor, id_goods)
	WHERE table_vendor.login_vendor = in_login_vendor AND table_goods.name_goods = in_name_goods AND
	table_warehouse_info.name_warehouse = in_warehouse_name AND table_warehouse_info.country = table_vendor_price.country;
	--
	id_customer_v = (SELECT id_customer FROM table_customer WHERE login_customer = in_login_customer);
	SELECT ARRAY (SELECT ROW (id_consignment, amount_goods_available)::type_id_amount FROM view_market_all_countries_for_order
					 WHERE country = in_country::enum_country AND login_vendor = in_login_vendor AND name_goods = in_name_goods
					 ORDER BY CASE sales_model WHEN 'fifo' THEN sales_model ELSE NULL END DESC, 
         			 CASE sales_model WHEN 'fifo' THEN NULL ELSE sales_model END ASC) INTO ids_consigment_s;
	IF NOT FOUND THEN
		RETURN QUERY (SELECT 'No items or invalid request'::varchar, ''::varchar);
	END IF;
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
	IF in_amount_goods != 0 THEN
		RETURN QUERY (SELECT 'not enough goods'::varchar, ''::varchar);
	END IF;
	uuid_v = uuid_generate_v4();
	id_order_v = (SELECT nextval('seq_id_order'));
	FOREACH i_v, a_v IN ARRAY ids_consigment
	LOOP
		UPDATE table_consignment
		SET amount_goods_available = amount_goods_available - a_v,
			amount_goods_blocked = amount_goods_blocked + a_v
		WHERE id_warehouse = id_warehouse_v AND id_vendor = id_vendor_v AND 
		id_goods = id_goods_v AND id_consignment = i_v;
		IF NOT FOUND THEN
			RETURN QUERY (SELECT 'problem1'::varchar, ''::varchar);
		END IF;
		--
		INSERT INTO table_orders(id_order, id_customer, id_consignment, id_vendor, 
		id_goods, id_warehouse, price_goods, amount_goods, delivery_location_country,
		delivery_location_city, date_order_start, operation_uuid) VALUES
		(id_order_v, id_customer_v, i_v, id_vendor_v, id_goods_v, id_warehouse_v, price_goods_v, a_v, '', '', now(), uuid_v);
		IF NOT FOUND THEN
			RETURN QUERY (SELECT 'problem2'::varchar, ''::varchar);
		END IF;
	END LOOP;
	RETURN QUERY (SELECT 'ok'::varchar, uuid_v::varchar);
END;
$$ LANGUAGE PLPGSQL;