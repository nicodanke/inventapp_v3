-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2024-12-27T19:50:49.704Z

CREATE TABLE "account" (
  "id" int8 NOT NULL DEFAULT (nextval('account_id_seq'::regclass)),
  "code" varchar NOT NULL,
  "company_name" varchar NOT NULL,
  "phone" varchar,
  "email" varchar NOT NULL,
  "web_url" varchar,
  "active" bool NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE TABLE "account_address" (
  "account_id" int8 NOT NULL DEFAULT (nextval('account_address_account_id_seq'::regclass)),
  "country" varchar NOT NULL,
  "state" varchar,
  "sub_state" varchar,
  "street" varchar,
  "number" varchar,
  "unit" varchar,
  "postal_code" varchar,
  "lat" float8,
  "lng" float8,
  PRIMARY KEY ("account_id")
);

CREATE TABLE "account_module" (
  "id" int8 NOT NULL DEFAULT (nextval('account_module_id_seq'::regclass)),
  "module_id" int8 NOT NULL DEFAULT (nextval('account_module_module_id_seq'::regclass)),
  "account_id" int8 NOT NULL DEFAULT (nextval('account_module_account_id_seq'::regclass)),
  "started_at" timestamptz NOT NULL DEFAULT (now()),
  "ended_at" timestamptz,
  "price" float8 NOT NULL,
  PRIMARY KEY ("id", "module_id", "account_id")
);

CREATE TABLE "account_plan" (
  "id" int8 NOT NULL DEFAULT (nextval('account_plan_id_seq'::regclass)),
  "plan_id" int8 NOT NULL DEFAULT (nextval('account_plan_plan_id_seq'::regclass)),
  "account_id" int8 NOT NULL DEFAULT (nextval('account_plan_account_id_seq'::regclass)),
  "started_at" timestamptz NOT NULL DEFAULT (now()),
  "ended_at" timestamptz,
  "price" float8 NOT NULL,
  PRIMARY KEY ("id", "plan_id", "account_id")
);

CREATE TABLE "account_promotion" (
  "id" int8 NOT NULL DEFAULT (nextval('account_promotion_id_seq'::regclass)),
  "promotion_id" int8 NOT NULL DEFAULT (nextval('account_promotion_promotion_id_seq'::regclass)),
  "account_id" int8 NOT NULL DEFAULT (nextval('account_promotion_account_id_seq'::regclass)),
  "started_at" timestamptz NOT NULL DEFAULT (now()),
  "ended_at" timestamptz,
  PRIMARY KEY ("id", "promotion_id", "account_id")
);

CREATE TABLE "feature" (
  "id" int8 NOT NULL DEFAULT (nextval('feature_id_seq'::regclass)),
  "name" varchar NOT NULL,
  "code" varchar NOT NULL,
  PRIMARY KEY ("id")
);

CREATE TABLE "module" (
  "id" int8 NOT NULL DEFAULT (nextval('module_id_seq'::regclass)),
  "name" varchar NOT NULL,
  "code" varchar NOT NULL,
  PRIMARY KEY ("id")
);

CREATE TABLE "module_country" (
  "module_id" int8 NOT NULL DEFAULT (nextval('module_country_module_id_seq'::regclass)),
  "country" varchar NOT NULL,
  "price" float8 NOT NULL,
  PRIMARY KEY ("module_id", "country")
);

CREATE TABLE "module_feature" (
  "module_id" int8 NOT NULL DEFAULT (nextval('module_feature_module_id_seq'::regclass)),
  "feature_id" int8 NOT NULL DEFAULT (nextval('module_feature_feature_id_seq'::regclass)),
  PRIMARY KEY ("module_id", "feature_id")
);

CREATE TABLE "permission" (
  "id" int8 NOT NULL DEFAULT (nextval('permission_id_seq'::regclass)),
  "name" varchar NOT NULL,
  "description" varchar,
  "parent_id" int8,
  PRIMARY KEY ("id")
);

CREATE TABLE "plan" (
  "id" int8 NOT NULL DEFAULT (nextval('plan_id_seq'::regclass)),
  "name" varchar NOT NULL,
  "description" varchar,
  PRIMARY KEY ("id")
);

CREATE TABLE "plan_country" (
  "plan_id" int8 NOT NULL DEFAULT (nextval('plan_country_plan_id_seq'::regclass)),
  "country" varchar NOT NULL,
  "price" float8 NOT NULL,
  PRIMARY KEY ("plan_id", "country")
);

CREATE TABLE "plan_module" (
  "plan_id" int8 NOT NULL DEFAULT (nextval('plan_module_plan_id_seq'::regclass)),
  "module_id" int8 NOT NULL DEFAULT (nextval('plan_module_module_id_seq'::regclass)),
  PRIMARY KEY ("plan_id", "module_id")
);

CREATE TABLE "promotion" (
  "id" int8 NOT NULL DEFAULT (nextval('promotion_id_seq'::regclass)),
  "name" varchar NOT NULL,
  "discount_percent" float8,
  "amount" float8,
  PRIMARY KEY ("id")
);

CREATE TABLE "role" (
  "id" int8 NOT NULL DEFAULT (nextval('role_id_seq'::regclass)),
  "name" varchar NOT NULL,
  "description" varchar,
  "account_id" int8 NOT NULL DEFAULT (nextval('role_account_id_seq'::regclass)),
  PRIMARY KEY ("id")
);

CREATE TABLE "role_permission" (
  "role_id" int8 NOT NULL DEFAULT (nextval('role_permission_role_id_seq'::regclass)),
  "permission_id" int8 NOT NULL DEFAULT (nextval('role_permission_permission_id_seq'::regclass)),
  PRIMARY KEY ("role_id", "permission_id")
);

CREATE TABLE "session" (
  "id" uuid NOT NULL,
  "user_id" int8 NOT NULL DEFAULT (nextval('user_id_seq'::regclass)),
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "is_blocked" bool NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "expires_at" timestamptz NOT NULL,
  PRIMARY KEY ("id")
);

CREATE TABLE "user" (
  "id" int8 NOT NULL DEFAULT (nextval('user_id_seq'::regclass)),
  "username" varchar NOT NULL,
  "password" varchar NOT NULL,
  "name" varchar NOT NULL,
  "lastname" varchar NOT NULL,
  "email" varchar NOT NULL,
  "phone" varchar,
  "active" bool NOT NULL DEFAULT true,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "password_changed_at" timestamptz NOT NULL DEFAULT (now()),
  "role_id" int8 NOT NULL DEFAULT (nextval('user_role_id_seq'::regclass)),
  "account_id" int8 NOT NULL DEFAULT (nextval('user_account_id_seq'::regclass)),
  PRIMARY KEY ("id")
);

COMMENT ON TABLE "permission" IS 'Common table shared by all accounts';

ALTER TABLE "account_address" ADD FOREIGN KEY ("account_id") REFERENCES "account" ("id");

ALTER TABLE "account_module" ADD FOREIGN KEY ("account_id") REFERENCES "account" ("id");

ALTER TABLE "account_module" ADD FOREIGN KEY ("module_id") REFERENCES "module" ("id");

ALTER TABLE "account_plan" ADD FOREIGN KEY ("plan_id") REFERENCES "plan" ("id");

ALTER TABLE "account_plan" ADD FOREIGN KEY ("account_id") REFERENCES "account" ("id");

ALTER TABLE "account_promotion" ADD FOREIGN KEY ("account_id") REFERENCES "account" ("id");

ALTER TABLE "account_promotion" ADD FOREIGN KEY ("promotion_id") REFERENCES "promotion" ("id");

ALTER TABLE "module_country" ADD FOREIGN KEY ("module_id") REFERENCES "module" ("id");

ALTER TABLE "module_feature" ADD FOREIGN KEY ("module_id") REFERENCES "module" ("id");

ALTER TABLE "module_feature" ADD FOREIGN KEY ("feature_id") REFERENCES "feature" ("id");

ALTER TABLE "permission" ADD FOREIGN KEY ("parent_id") REFERENCES "permission" ("id");

ALTER TABLE "plan_country" ADD FOREIGN KEY ("plan_id") REFERENCES "plan" ("id");

ALTER TABLE "plan_module" ADD FOREIGN KEY ("module_id") REFERENCES "module" ("id");

ALTER TABLE "plan_module" ADD FOREIGN KEY ("plan_id") REFERENCES "plan" ("id");

ALTER TABLE "role" ADD FOREIGN KEY ("account_id") REFERENCES "account" ("id");

ALTER TABLE "role_permission" ADD FOREIGN KEY ("role_id") REFERENCES "role" ("id");

ALTER TABLE "role_permission" ADD FOREIGN KEY ("permission_id") REFERENCES "permission" ("id");

ALTER TABLE "user" ADD FOREIGN KEY ("account_id") REFERENCES "account" ("id");

ALTER TABLE "session" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");
