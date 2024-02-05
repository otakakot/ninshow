table "accounts" {
  schema = schema.public
  column "id" {
    null    = false
    type    = uuid
    default = sql("gen_random_uuid()")
  }
  column "name" {
    null    = false
    type    = text
    default = ""
  }
  column "email" {
    null = false
    type = text
  }
  column "created_at" {
    null    = false
    type    = timestamptz(3)
    default = sql("CURRENT_TIMESTAMP")
  }
  column "updated_at" {
    null    = false
    type    = timestamptz(3)
    default = sql("CURRENT_TIMESTAMP")
  }
  column "deleted" {
    null    = false
    type    = boolean
    default = false
  }
  primary_key {
    columns = [column.id]
  }
  index "accounts_email_key" {
    unique  = true
    columns = [column.email]
  }
  index "accounts_email_not_deleted_index" {
    unique  = true
    columns = [column.email]
    where   = "(NOT deleted)"
  }
}
table "jwt_sign_keys" {
  schema = schema.public
  column "id" {
    null    = false
    type    = uuid
    default = sql("gen_random_uuid()")
  }
  column "key" {
    null = false
    type = bytea
  }
  column "created_at" {
    null    = false
    type    = timestamptz(3)
    default = sql("CURRENT_TIMESTAMP")
  }
  column "updated_at" {
    null    = false
    type    = timestamptz(3)
    default = sql("CURRENT_TIMESTAMP")
  }
  primary_key {
    columns = [column.id]
  }
}
table "oidc_clients" {
  schema = schema.public
  column "id" {
    null    = false
    type    = uuid
    default = sql("gen_random_uuid()")
  }
  column "name" {
    null    = false
    type    = text
    default = ""
  }
  column "redirect_uri" {
    null    = false
    type    = text
    default = ""
  }
  column "created_at" {
    null    = false
    type    = timestamptz(3)
    default = sql("CURRENT_TIMESTAMP")
  }
  column "updated_at" {
    null    = false
    type    = timestamptz(3)
    default = sql("CURRENT_TIMESTAMP")
  }
  column "deleted" {
    null    = false
    type    = boolean
    default = false
  }
  primary_key {
    columns = [column.id]
  }
}
table "oidc_secrets" {
  schema = schema.public
  column "id" {
    null    = false
    type    = uuid
    default = sql("gen_random_uuid()")
  }
  column "created_at" {
    null    = false
    type    = timestamptz(3)
    default = sql("CURRENT_TIMESTAMP")
  }
  column "updated_at" {
    null    = false
    type    = timestamptz(3)
    default = sql("CURRENT_TIMESTAMP")
  }
  column "value" {
    null = false
    type = bytea
  }
  column "client_id" {
    null = false
    type = uuid
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "oidc_secrets_client_id_fkey" {
    columns     = [column.client_id]
    ref_columns = [table.oidc_clients.column.id]
    on_update   = CASCADE
    on_delete   = RESTRICT
  }
  index "oidc_secrets_client_id_idx" {
    columns = [column.client_id]
  }
  index "oidc_secrets_client_id_key" {
    unique  = true
    columns = [column.client_id]
  }
}
table "password_authns" {
  schema = schema.public
  column "id" {
    null    = false
    type    = uuid
    default = sql("gen_random_uuid()")
  }
  column "account_id" {
    null = false
    type = uuid
  }
  column "created_at" {
    null    = false
    type    = timestamptz(3)
    default = sql("CURRENT_TIMESTAMP")
  }
  column "updated_at" {
    null    = false
    type    = timestamptz(3)
    default = sql("CURRENT_TIMESTAMP")
  }
  column "value" {
    null = false
    type = bytea
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "password_authns_account_id_fkey" {
    columns     = [column.account_id]
    ref_columns = [table.accounts.column.id]
    on_update   = CASCADE
    on_delete   = RESTRICT
  }
  index "password_authns_account_id_idx" {
    columns = [column.account_id]
  }
  index "password_authns_account_id_key" {
    unique  = true
    columns = [column.account_id]
  }
}
schema "public" {
  comment = "standard public schema"
}
