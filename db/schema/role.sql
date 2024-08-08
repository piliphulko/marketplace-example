REVOKE ALL ON SCHEMA public FROM public;

REVOKE ALL ON DATABASE test_db FROM public;

-- vault_creater_roles
CREATE ROLE vault_creater_roles WITH CREATEROLE LOGIN PASSWORD '5432';

GRANT CONNECT ON DATABASE test_db TO vault_creater_roles;

CREATE ROLE role_service_acct_auth;

-- role_service_acct_auth
GRANT CONNECT ON DATABASE test_db TO role_service_acct_auth;

GRANT USAGE ON SCHEMA public TO role_service_acct_auth;

GRANT SELECT, INSERT, UPDATE 
ON TABLE table_customer, table_customer_info, 
    table_vendor, table_vendor_info,
    table_warehouse, table_warehouse_info, table_warehouse_commission 
TO role_service_acct_auth;

GRANT INSERT ON TABLE table_customer_wallet (id_customer), 
    table_warehouse_wallet (id_warehouse), table_vendor_wallet (id_vendor)
TO role_service_acct_auth;