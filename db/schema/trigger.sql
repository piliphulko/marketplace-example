CREATE TRIGGER trigger_vendor_price_archive 
BEFORE INSERT OR UPDATE OR DELETE ON table_vendor_price_archive
FOR EACH ROW EXECUTE PROCEDURE function_trigger_archive();