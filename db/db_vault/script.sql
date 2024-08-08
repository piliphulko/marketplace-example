-- DB CRAETE
CREATE DATABASE database_vault_storage;

-- SOLVING THE PROBLEM OF INHERITANCE 'PUBLIC'
REVOKE ALL ON SCHEMA public FROM public;

REVOKE ALL ON DATABASE database_vault_storage FROM public;

-- ROLE CREATE
CREATE ROLE user_vault_storage WITH LOGIN PASSWORD '5432';

GRANT CONNECT ON DATABASE database_vault_storage TO user_vault_storage;

GRANT USAGE ON SCHEMA public TO user_vault_storage;

GRANT ALL PRIVILEGES ON DATABASE database_vault_storage TO user_vault_storage;

GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO user_vault_storage;



-- COPY FROM https://developer.hashicorp.com/vault/docs/configuration/storage/postgresql
CREATE TABLE vault_kv_store (
  parent_path TEXT COLLATE "C" NOT NULL,
  path        TEXT COLLATE "C",
  key         TEXT COLLATE "C",
  value       BYTEA,
  CONSTRAINT pkey PRIMARY KEY (path, key)
);

CREATE INDEX parent_path_idx ON vault_kv_store (parent_path);

CREATE TABLE vault_ha_locks (
  ha_key                                      TEXT COLLATE "C" NOT NULL,
  ha_identity                                 TEXT COLLATE "C" NOT NULL,
  ha_value                                    TEXT COLLATE "C",
  valid_until                                 TIMESTAMP WITH TIME ZONE NOT NULL,
  CONSTRAINT ha_key PRIMARY KEY (ha_key)
);
