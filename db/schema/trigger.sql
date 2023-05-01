CREATE TRIGGER function_trigger_archive_table_vendor_price 
BEFORE INSERT OR UPDATE OR DELETE ON table_vendor_price
FOR EACH ROW EXECUTE PROCEDURE function_trigger_archive();

CREATE TRIGGER function_trigger_consignment_check_goods_in_stock
BEFORE UPDATE ON table_consignment
FOR EACH ROW EXECUTE PROCEDURE function_trigger_consignment();