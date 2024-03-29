CREATE OR REPLACE FUNCTION function_trigger_archive_table_vendor_price() RETURNS trigger AS $$
BEGIN
	IF TG_OP = 'INSERT' THEN
		INSERT INTO table_vendor_price_archive
		SELECT 'INSERT', now(), NEW.*
		FROM table_vendor_price;
	ELSEIF TG_OP = 'UPDATE' THEN
		INSERT INTO table_vendor_price_archive
		SELECT 'UPDATE', now(), NEW.*
		FROM table_vendor_price;
	ELSEIF TG_OP = 'DELETE' THEN
		INSERT INTO table_vendor_price_archive
		SELECT 'DELETE', now(), OLD.*
		FROM table_vendor_price;
	END IF;
	RETURN NULL;
END;
$$ LANGUAGE PLPGSQL;

CREATE OR REPLACE FUNCTION function_trigger_consignment() RETURNS trigger AS $$
BEGIN
	IF NEW.amount_goods_available = 0 AND NEW.amount_goods_blocked = 0 THEN
		NEW.goods_in_stock = false;
		NEW.date_sold_out = now();
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE PLPGSQL;

CREATE OR REPLACE FUNCTION function_create_order
(
	in_uuid varchar,
	in_login_customer varchar,
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
	country_warehouse_v varchar;
BEGIN
	
	SELECT table_consignment.id_vendor, table_consignment.id_goods, 
		table_consignment.id_warehouse, table_vendor_price.price_goods,
		table_warehouse_info.country
	INTO id_vendor_v, id_goods_v, id_warehouse_v, price_goods_v, country_warehouse_v
	FROM table_consignment
	JOIN table_vendor_info USING (id_vendor) 
	JOIN table_goods USING (id_goods) 
	JOIN table_warehouse_info USING (id_warehouse)
	JOIN table_vendor_price USING (id_goods)
	WHERE table_vendor_info.name_vendor = in_name_vendor AND table_goods.name_goods = in_name_goods AND
		  table_warehouse_info.name_warehouse = in_warehouse_name AND 
		  table_warehouse_info.country = table_vendor_price.country;
	
	IF NOT FOUND THEN
		RETURN 'invalid request'::varchar;
	END IF;
	
	SELECT id_customer, delivery_location_country, delivery_location_city
	INTO id_customer_v, delivery_location_country_v, delivery_location_city_v
	FROM table_customer
	JOIN table_customer_info USING (id_customer)
	WHERE login_customer = in_login_customer;
	
	IF NOT FOUND THEN
		RETURN 'db_fn_err1'::varchar;
	ELSEIF delivery_location_country_v != country_warehouse_v THEN
		RETURN 'delivery country must match with warehouse side'::varchar;
	END IF;
	
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
	
	IF NOT FOUND THEN
		RETURN 'invalid request'::varchar;
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
	
	IF NOT FOUND THEN
		RETURN 'db_fn_err2'::varchar;
	ELSEIF in_amount_goods != 0 THEN
		RETURN 'not enough goods'::varchar;
	END IF;
	
	id_order_v = (SELECT nextval('seq_id_order'));
	
	warehouse_commission_percentage = (SELECT commission_percentage FROM table_warehouse_commission
									   WHERE id_warehouse = id_warehouse_v);
	tax_vendor_commission_percentage = (SELECT vat FROM table_tax_plan
									    WHERE country = country_warehouse_v::enum_country);
	system_commission_percentage = (SELECT commission_percentage FROM table_system_commission
									WHERE id_system = 1);
	
	FOREACH i_v, a_v IN ARRAY ids_consigment
	LOOP
		
		UPDATE table_consignment
		SET amount_goods_available = amount_goods_available - a_v,
			amount_goods_blocked = amount_goods_blocked + a_v
		WHERE id_warehouse = id_warehouse_v AND 
			  id_vendor = id_vendor_v AND 
			  id_goods = id_goods_v AND 
			  id_consignment = i_v;
		
		IF NOT FOUND THEN
			RETURN 'db_fn_err3'::varchar;
		END IF;
		
		INSERT INTO table_orders(id_order, id_customer, id_consignment, id_vendor, 
		id_goods, id_warehouse, price_goods, amount_goods, delivery_location_country,
		delivery_location_city, date_order_start, operation_uuid) VALUES
		(id_order_v, id_customer_v, i_v, id_vendor_v, id_goods_v, id_warehouse_v, price_goods_v, a_v, 
		delivery_location_country_v, delivery_location_city_v, now(), in_uuid::uuid);
		
		IF NOT FOUND THEN
			RETURN 'db_fn_err4'::varchar;
		END IF;
		
		money_customer_debit_v = price_goods_v * a_v;
		
		tax_money_vendor_credit_v = round(money_customer_debit_v * tax_vendor_commission_percentage, 2);
		
		money_customer_debit_net = money_customer_debit_v - tax_money_vendor_credit_v;
		
		money_warehouse_credit_v = round(money_customer_debit_net * warehouse_commission_percentage, 2);
		
		money_system_credit_v = round(money_customer_debit_net * system_commission_percentage, 2);
		
		money_vendor_credit_v = money_customer_debit_net - (money_warehouse_credit_v + money_system_credit_v);
		
		IF money_customer_debit_v != 
			(tax_money_vendor_credit_v + money_warehouse_credit_v + money_system_credit_v + money_vendor_credit_v) THEN
			RETURN 'db_fn_err5'::varchar;
		END IF;
		
		INSERT INTO table_ledger(id_order, id_consignment, id_customer, money_customer_debit,
							id_vendor, money_vendor_credit, tax_money_vendor_credit, 
							id_warehouse, money_warehouse_credit, money_system_credit,
							operation_uuid)
		VALUES (id_order_v, i_v, id_customer_v, money_customer_debit_v, id_vendor_v, money_vendor_credit_v, 
			tax_money_vendor_credit_v, id_warehouse_v, money_warehouse_credit_v, money_system_credit_v, in_uuid::uuid);
		
		IF NOT FOUND THEN
			RETURN 'db_fn_err6'::varchar;
		END IF;
		
	END LOOP;
	RETURN 'ok'::varchar;
END;
$$ LANGUAGE PLPGSQL;

CREATE OR REPLACE FUNCTION function_confirm_order
(
	in_login_customer varchar,
	in_uuid varchar
) 
RETURNS varchar AS $$
DECLARE
	order_price_v domain_money = 0;
	id_customer_v int;
	loop_id_vendor_v int;
	loop_money_vendor_v domain_money = 0;
	loop_tax_money_vendor_v domain_money = 0;
	loop_id_warehouse_v int;
	loop_money_warehouse_v domain_money = 0;
	loop_money_system_v domain_money = 0;
	array_details_ledger_v type_details_ledger ARRAY;
	total_credit_v domain_money = 0;
BEGIN
	
	SELECT sum(money_customer_debit), id_customer
	INTO order_price_v, id_customer_v
	FROM table_ledger
	WHERE operation_uuid = in_uuid::uuid
		AND delivery_status_order = 'unconfirmed order'::enum_status_order
	GROUP BY operation_uuid, id_customer;
	
	IF NOT FOUND THEN
		RETURN 'db_fn_err7'::varchar;
	ELSEIF order_price_v > (SELECT table_customer_wallet.amount_money FROM table_customer_wallet 
						 JOIN table_customer USING (id_customer) 
						 WHERE table_customer.login_customer = in_login_customer) THEN
		RETURN 'not enough money'::varchar;
	END IF;
	
	SELECT ARRAY (
		SELECT ROW (
			id_vendor, money_vendor_credit, 
			tax_money_vendor_credit,
			id_warehouse, money_warehouse_credit,
			money_system_credit)::type_details_ledger 
		FROM table_ledger
		WHERE operation_uuid = in_uuid::uuid
	) INTO array_details_ledger_v;
	
	IF NOT FOUND THEN
			RETURN 'db_fn_err8'::varchar;
	END IF;
	
	FOREACH loop_id_vendor_v, loop_money_vendor_v, loop_tax_money_vendor_v, loop_id_warehouse_v,
		loop_money_warehouse_v, loop_money_system_v IN ARRAY array_details_ledger_v
	LOOP
		
		UPDATE table_vendor_wallet
		SET blocked_money = blocked_money + loop_money_vendor_v,
			blocked_tax_money = blocked_tax_money + loop_tax_money_vendor_v
		WHERE id_vendor = loop_id_vendor_v;
		
		IF NOT FOUND THEN
			RETURN 'db_fn_err9'::varchar;
		END IF;
		
		UPDATE table_warehouse_wallet
		SET blocked_money = blocked_money + loop_money_warehouse_v
		WHERE id_warehouse = loop_id_warehouse_v;
		
		IF NOT FOUND THEN
			RETURN 'db_fn_err10'::varchar;
		END IF;
		
		UPDATE table_system_wallet
		SET blocked_money = blocked_money + loop_money_system_v;
		
		IF NOT FOUND THEN
			RETURN 'db_fn_err11'::varchar;
		END IF;
		
		total_credit_v = total_credit_v + loop_money_vendor_v + loop_tax_money_vendor_v +  loop_money_warehouse_v + loop_money_system_v;
		
	END LOOP;
	
	IF NOT FOUND THEN
		RETURN 'db_fn_err12'::varchar;
	ELSEIF total_credit_v != order_price_v THEN
		RETURN 'db_fn_err13'::varchar;
	END IF;
	
	UPDATE table_customer_wallet
	SET amount_money = amount_money - order_price_v, 
		blocked_money = blocked_money + order_price_v
	WHERE id_customer = id_customer_v;
	
	IF NOT FOUND THEN
		RETURN 'db_fn_err14'::varchar;
	END IF;
	
	UPDATE table_ledger
	SET delivery_status_order = 'confirmed order'::enum_status_order
	WHERE operation_uuid = in_uuid::uuid;
	
	IF NOT FOUND THEN
		RETURN 'db_fn_err15'::varchar;
	END IF;
	
	UPDATE table_orders
	SET delivery_status_order = 'confirmed order'::enum_status_order
	WHERE operation_uuid = in_uuid::uuid;
	
	IF NOT FOUND THEN
		RETURN 'db_fn_err16'::varchar;
	END IF;
	
	RETURN 'ok'::varchar;
	
END;
$$ LANGUAGE PLPGSQL;

CREATE OR REPLACE FUNCTION function_complete_order
(
	in_login_customer varchar,
	in_uuid varchar
)
RETURNS varchar AS $$
DECLARE
	order_price_v domain_money = 0;
	id_customer_v int;
	loop_id_vendor_v int;
	loop_money_vendor_v domain_money = 0;
	loop_tax_money_vendor_v domain_money = 0;
	loop_id_warehouse_v int;
	loop_money_warehouse_v domain_money = 0;
	loop_money_system_v domain_money = 0;
	array_details_ledger_v type_details_ledger ARRAY;
	total_credit_v domain_money = 0;
	ids_consigment_v type_id_amount ARRAY;
	loop_id_consignment_v int;
	loop_amount_goods int;
BEGIN
	
	SELECT sum(money_customer_debit), id_customer
	INTO order_price_v, id_customer_v
	FROM table_ledger
	WHERE operation_uuid = in_uuid::uuid
		AND delivery_status_order = 'confirmed order'::enum_status_order
	GROUP BY operation_uuid, id_customer;
	
	IF NOT FOUND THEN
		RETURN 'invalid request'::varchar;
	END IF;
	
	SELECT ARRAY (
		SELECT ROW (
			id_vendor, money_vendor_credit, 
			tax_money_vendor_credit,
			id_warehouse, money_warehouse_credit,
			money_system_credit)::type_details_ledger
		FROM table_ledger
		WHERE operation_uuid = in_uuid::uuid
	) INTO array_details_ledger_v;
	
	IF NOT FOUND THEN
		RETURN 'db_fn_err17'::varchar;
	END IF;
	
	FOREACH loop_id_vendor_v, loop_money_vendor_v, loop_tax_money_vendor_v, loop_id_warehouse_v,
		loop_money_warehouse_v, loop_money_system_v IN ARRAY array_details_ledger_v
	LOOP
		
		UPDATE table_vendor_wallet
		SET blocked_money = blocked_money - loop_money_vendor_v,
			blocked_tax_money = blocked_tax_money - loop_tax_money_vendor_v,
			amount_money = amount_money + loop_money_vendor_v,
			tax_money = tax_money + loop_tax_money_vendor_v
		WHERE id_vendor = loop_id_vendor_v;
		
		IF NOT FOUND THEN
			RETURN 'db_fn_err18'::varchar;
		END IF;
		
		UPDATE table_warehouse_wallet
		SET blocked_money = blocked_money - loop_money_warehouse_v,
			amount_money = amount_money + loop_money_warehouse_v
		WHERE id_warehouse = loop_id_warehouse_v;
		
		IF NOT FOUND THEN
			RETURN 'db_fn_err19'::varchar;
		END IF;
		
		UPDATE table_system_wallet
		SET blocked_money = blocked_money - loop_money_system_v,
			amount_money = amount_money + loop_money_warehouse_v;
		
		IF NOT FOUND THEN
			RETURN 'db_fn_err20'::varchar;
		END IF;
		
		total_credit_v = total_credit_v + loop_money_vendor_v + loop_tax_money_vendor_v +  loop_money_warehouse_v + loop_money_system_v;
		
	END LOOP;
	
	IF NOT FOUND THEN
		RETURN 'db_fn_err21'::varchar;
	ELSEIF total_credit_v != order_price_v THEN
		RETURN 'db_fn_err22'::varchar;
	END IF;
	
	SELECT ARRAY (
		SELECT ROW (id_consignment, amount_goods)::type_id_amount 
		FROM table_orders
		WHERE operation_uuid = in_uuid::uuid AND delivery_status_order = 'confirmed order'::enum_status_order
	) INTO ids_consigment_v;
	
	IF NOT FOUND THEN
		RETURN 'db_fn_err23'::varchar;
	END IF;
	
	FOREACH loop_id_consignment_v, loop_amount_goods IN ARRAY ids_consigment_v
	LOOP
		
		UPDATE table_consignment
		SET amount_goods_blocked = amount_goods_blocked - loop_amount_goods
		WHERE id_consignment = loop_id_consignment_v;
		
		IF NOT FOUND THEN
			RETURN 'db_fn_err24'::varchar;
		END IF;
		
	END LOOP;
	
	IF NOT FOUND THEN
		RETURN 'db_fn_err25'::varchar;
	END IF;
	
	UPDATE table_customer_wallet
	SET blocked_money = blocked_money - order_price_v 
	WHERE id_customer = id_customer_v;
	
	IF NOT FOUND THEN
		RETURN 'db_fn_err26'::varchar;
	END IF;
	
	UPDATE table_ledger
	SET delivery_status_order = 'completed order'::enum_status_order
	WHERE operation_uuid = in_uuid::uuid;
	
	IF NOT FOUND THEN
		RETURN 'db_fn_err27'::varchar;
	END IF;
	
	UPDATE table_orders
	SET delivery_status_order = 'completed order'::enum_status_order,
		date_order_finish = now()
	WHERE operation_uuid = in_uuid::uuid;
	
	IF NOT FOUND THEN
		RETURN 'db_fn_err28'::varchar;
	END IF;
	
	RETURN 'ok'::varchar;
	
END;
$$ LANGUAGE PLPGSQL;

CREATE OR REPLACE FUNCTION function_cancellation_order
(
	in_login_customer varchar,
	in_uuid varchar
)
RETURNS varchar AS $$
DECLARE
	order_price_v domain_money = 0;
	id_customer_v int;
	loop_id_vendor_v int;
	loop_money_vendor_v domain_money = 0;
	loop_tax_money_vendor_v domain_money = 0;
	loop_id_warehouse_v int;
	loop_money_warehouse_v domain_money = 0;
	loop_money_system_v domain_money = 0;
	array_details_ledger_v type_details_ledger ARRAY;
	total_credit_v domain_money = 0;
	ids_consigment_v type_id_amount ARRAY;
	loop_id_consignment_v int;
	loop_amount_goods int;
	status1_v varchar;
	status2_v varchar;
BEGIN
	
	SELECT table_orders.delivery_status_order, table_ledger.delivery_status_order
	INTO status1_v, status2_v
	FROM table_orders
	JOIN table_ledger USING (id_order)
	WHERE table_orders.operation_uuid = table_ledger.operation_uuid AND
		table_ledger.operation_uuid = in_uuid::uuid
	GROUP BY table_orders.delivery_status_order, table_ledger.delivery_status_order;
	
	IF NOT FOUND THEN
		RETURN 'invalid request'::varchar;
	END IF;
	
	SELECT sum(money_customer_debit), id_customer
	INTO order_price_v, id_customer_v
	FROM table_ledger
	WHERE operation_uuid = in_uuid::uuid
	GROUP BY operation_uuid, id_customer;
	
	IF NOT FOUND THEN
		RETURN 'db_fn_err29'::varchar;
	ELSEIF status1_v != status2_v THEN
		RETURN 'db_fn_err30'::varchar;
	ELSEIF status1_v::enum_status_order = 'confirmed order'::enum_status_order THEN
		
		SELECT ARRAY (
			SELECT ROW (
				id_vendor, money_vendor_credit, 
				tax_money_vendor_credit,
				id_warehouse, money_warehouse_credit,
				money_system_credit)::type_details_ledger
			FROM table_ledger
			WHERE operation_uuid = in_uuid::uuid
		) INTO array_details_ledger_v;
		
		IF NOT FOUND THEN
			RETURN 'db_fn_err31'::varchar;
		END IF;
		
		FOREACH loop_id_vendor_v, loop_money_vendor_v, loop_tax_money_vendor_v, loop_id_warehouse_v,
			loop_money_warehouse_v, loop_money_system_v IN ARRAY array_details_ledger_v
		LOOP
			
			UPDATE table_vendor_wallet
			SET blocked_money = blocked_money - loop_money_vendor_v,
				blocked_tax_money = blocked_tax_money - loop_tax_money_vendor_v
			WHERE id_vendor = loop_id_vendor_v;
			
			IF NOT FOUND THEN
				RETURN 'db_fn_err32'::varchar;
			END IF;
			
			UPDATE table_warehouse_wallet
			SET blocked_money = blocked_money - loop_money_warehouse_v
			WHERE id_warehouse = loop_id_warehouse_v;
			
			IF NOT FOUND THEN
				RETURN 'db_fn_err33'::varchar;
			END IF;
			
			UPDATE table_system_wallet
			SET blocked_money = blocked_money - loop_money_system_v;
			
			IF NOT FOUND THEN
				RETURN 'db_fn_err34'::varchar;
			END IF;
			
			total_credit_v = total_credit_v + loop_money_vendor_v + loop_tax_money_vendor_v +  loop_money_warehouse_v + loop_money_system_v;
			
		END LOOP;
		
		IF NOT FOUND THEN
			RETURN 'db_fn_err35'::varchar;
		ELSEIF total_credit_v != order_price_v THEN
			RETURN 'db_fn_err36'::varchar;
		END IF;
			
			UPDATE table_customer_wallet
			SET blocked_money = blocked_money - order_price_v,
				amount_money = amount_money + order_price_v
			WHERE id_customer = id_customer_v;
			
			IF NOT FOUND THEN
				RETURN 'db_fn_err37'::varchar;
			END IF;
			
	ELSEIF status1_v::enum_status_order = 'completed order'::enum_status_order THEN 
		RETURN 'order completed cannot be canceled'::varchar;
	END IF;
	
	SELECT ARRAY (
		SELECT ROW (id_consignment, amount_goods)::type_id_amount 
		FROM table_orders
		WHERE operation_uuid = in_uuid::uuid
	) INTO ids_consigment_v;
	
	IF NOT FOUND THEN
		RETURN 'db_fn_err38'::varchar;
	END IF;
	
	FOREACH loop_id_consignment_v, loop_amount_goods IN ARRAY ids_consigment_v
	LOOP
		
		UPDATE table_consignment
		SET amount_goods_blocked = amount_goods_blocked - loop_amount_goods,
			amount_goods_available = amount_goods_available + loop_amount_goods
		WHERE id_consignment = loop_id_consignment_v;
		
		IF NOT FOUND THEN
			RETURN 'db_fn_err39'::varchar;
		END IF;
		
	END LOOP;
	
	IF NOT FOUND THEN
		RETURN 'db_fn_err40'::varchar;
	END IF;
	
	DELETE FROM table_ledger
	WHERE operation_uuid = in_uuid::uuid;
	
	IF NOT FOUND THEN
		RETURN 'db_fn_err41'::varchar;
	END IF;
	
	DELETE FROM table_orders
	WHERE operation_uuid = in_uuid::uuid;
	
	IF NOT FOUND THEN
		RETURN 'db_fn_err42'::varchar;
	END IF;
	
	RETURN 'ok'::varchar;
	
END;
$$ LANGUAGE PLPGSQL;