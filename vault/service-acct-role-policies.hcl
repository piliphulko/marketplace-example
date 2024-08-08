path "database/creds/service-acct-role" {
  capabilities = ["read"]
}

path "pki/issue/service-acct-role" {
  capabilities = ["create", "update"]
}

path "auth/token/lookup-self" {
  capabilities = ["read"]
}

path "pki/*" {
  capabilities = ["create", "read", "update", "delete", "list"]
}