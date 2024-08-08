enable_db_engine:
	vault secrets enable database

connect_db:
	vault write database/config/pgconn_vault plugin_name="postgresql-database-plugin" allowed_roles="*" connection_url="postgresql://{{username}}:{{password}}@host.docker.internal:5432/test_db" username="vault_creater_roles" password="5432" password_authentication="scram-sha-256"

shema_create_role: #нужно в cmd 
	vault write database/roles/service-acct-role db_name="pgconn_vault" creation_statements='CREATE ROLE "{{name}}" WITH LOGIN PASSWORD ''{{password}}'' VALID UNTIL ''{{expiration}}'' IN ROLE role_service_acct_auth;' default_ttl="1h" max_ttl="24h"

get_login_pass:
	vault read database/creds/service-acct-role
	
revoke-acct:
	vault lease revoke database/creds/service-acct-role/ #{{username}}

start_vault_server:
	vault server -config=config.hcl

start_init_server:
	vault operator init -key-shares=1 -key-threshold=1 -pgp-keys="base64_gpgPublicKey.asc"