table "categories" {
  schema = schema.tweetyah
  column "id" {
    null           = false
    type           = int
    auto_increment = true
  }
  column "name" {
    null = false
    type = varchar(50)
  }
  column "id_user" {
    null = false
    type = varchar(50)
  }
  column "color" {
    null = false
    type = varchar(20)
  }
  primary_key {
    columns = [column.id]
  }
}

table "library_items" {
  schema = schema.tweetyah
  column "id" {
    null           = false
    type           = int
    auto_increment = true
  }
  column "text" {
    null = false
    type = varchar(350)
  }
  column "last_used" {
    null = true
    type = timestamp
  }
  column "is_thread" {
    null = true
    type = bool
  }
  column "thread_count" {
    null = true
    type = int
  }
  column "id_user" {
    null = true
    type = varchar(50)
  }
  column "id_category" {
    null = true
    type = int
  }
  column "thread_parent" {
    null = true
    type = int
  }
  column "thread_order" {
    null = true
    type = int
  }
  primary_key {
    columns = [column.id]
  }
}

table "posts" {
  schema = schema.tweetyah
  column "id" {
    null           = false
    type           = bigint
    auto_increment = true
  }
  column "text" {
    null = false
    type = varchar(350)
  }
  column "send_at" {
    null = true
    type = timestamp
  }
  column "retweet_at" {
    null = true
    type = timestamp
  }
  column "is_thread" {
    null = true
    type = bool
  }
  column "thread_count" {
    null = true
    type = int
  }
  column "id_library_item" {
    null = true
    type = int
  }
  column "id_user" {
    null = true
    type = varchar(50)
  }
  column "id_category" {
    null = true
    type = int
  }
  column "thread_parent" {
    null = true
    type = bigint
  }
  column "thread_order" {
    null = true
    type = int
  }
  column "id_sent" {
    null = true
    type = bigint
  }
  column "status" {
    null    = false
    type    = int
    default = 0
  }
  column "error" {
    null = true
    type = json
  }
  column "service" {
    null = false
    type = int
  }
  primary_key {
    columns = [column.id]
  }
}

table "user_settings" {
  schema = schema.tweetyah
  column "id" {
    null = false
    type = varchar(50)
  }
  primary_key {
    columns = [column.id]
  }
}

table "users" {
  schema = schema.tweetyah
  column "id" {
    null = false
    type = bigint
    auto_increment = true
  }
  column "last_login" {
    null = false
    type = timestamp
  }
  primary_key {
    columns = [column.id]
  }
}

table "user_tokens" {
  schema = schema.tweetyah
  column "id" {
    null           = false
    type           = int
    auto_increment = true
  }
  column "user_id" {
    null = false
    type = int
  }
  column "access_token" {
    null = false
    type = varchar(200)
  }
  column "refresh_token" {
    null = true
    type = varchar(200)
  }
  column "access_token_expiry" {
    null = true
    type = timestamp
  }
  column "mastodon_domain" {
    null = true
    type = varchar(100)
  }
  primary_key {
    columns = [column.id]
  }
}

table "auth_providers" {
  schema = schema.tweetyah
  column "id" {
    null           = false
    type           = int
    auto_increment = true
  }
  column "user_id" {
    null = false
    type = int
  }
  column "type" {
    null = false
    type = int
  }
  column "service_id" {
    null = false
    type = varchar(50)
  }
  primary_key {
    columns = [column.id]
  }
}

schema "tweetyah" {
  charset = "utf8mb4"
  collate = "utf8mb4_0900_ai_ci"
}
