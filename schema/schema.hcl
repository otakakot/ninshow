table "accounts" {
  schema = schema.public
  column "id" {
    null    = false
    type    = uuid
    default = sql("gen_random_uuid()")
  }
  column "email" {
    null = false
    type = text
  }
  column "username" {
    null    = false
    type    = text
    default = ""
  }
  column "password" {
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
  index "accounts_username_idx" {
    columns = [column.username]
  }
}
schema "public" {
  comment = "standard public schema"
}
