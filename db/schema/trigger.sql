CREATE TRIGGER function_trigger_archive_table_vendor_price_use
BEFORE INSERT OR UPDATE OR DELETE ON table_vendor_price
FOR EACH ROW EXECUTE PROCEDURE function_trigger_archive_table_vendor_price();

CREATE TRIGGER function_trigger_consignment_use
BEFORE UPDATE ON table_consignment
FOR EACH ROW EXECUTE PROCEDURE function_trigger_consignment();