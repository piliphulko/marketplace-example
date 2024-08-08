storage "postgresql" {
  connection_url = "postgres://user_vault_storage:5432@host.docker.internal:5432/database_vault_storage"
}

listener "tcp" {
  address     = "0.0.0.0:8200"
  tls_cert_file = "/usr/local/bin/certificate.crt"
  tls_key_file  = "/usr/local/bin/private.key"
}

api_addr = "https://host.docker.internal:8200"
cluster_addr = "https://host.docker.internal:8201"
ui = true
