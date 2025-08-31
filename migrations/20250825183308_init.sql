-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "orders"
(
	"uid"                varchar PRIMARY KEY,
	"track_number"       varchar UNIQUE,
	"entry"              varchar,
	"locale"             varchar,
	"internal_signature" varchar,
	"customer_id"        varchar,
	"delivery_service"   varchar,
	"shardkey"           integer,
	"sm_id"              integer,
	"date_created"       timestamp,
	"oof_shard"          varchar
);

CREATE TABLE IF NOT EXISTS "deliveries"
(
	"order_uid" varchar PRIMARY KEY,
	"name"      varchar,
	"phone"     varchar,
	"zip"       varchar,
	"city"      varchar,
	"address"   varchar,
	"region"    varchar,
	"email"     varchar
);

CREATE TABLE IF NOT EXISTS "payments"
(
	"transaction"   varchar PRIMARY KEY,
	"request_id"    varchar,
	"currency"      varchar,
	"provider"      varchar,
	"amount"        integer,
	"payment_dt"    timestamp,
	"bank"          varchar,
	"delivery_cost" integer,
	"goods_total"   integer,
	"custom_fee"    integer
);

CREATE TABLE IF NOT EXISTS "items"
(
	"chrt_id"      integer,
	"track_number" varchar,
	"price"        integer,
	"rid"          varchar,
	"name"         varchar,
	"sale"         integer,
	"size"         varchar,
	"total_price"  integer,
	"nm_id"        integer,
	"brand"        varchar,
	"status"       integer,
	PRIMARY KEY (track_number, chrt_id)
);

ALTER TABLE IF EXISTS "deliveries"
	ADD FOREIGN KEY ("order_uid") REFERENCES "orders" ("uid");

ALTER TABLE IF EXISTS "payments"
	ADD FOREIGN KEY ("transaction") REFERENCES "orders" ("uid");

ALTER TABLE IF EXISTS "items"
	ADD FOREIGN KEY ("track_number") REFERENCES "orders" ("track_number");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS deliveries;
DROP TABLE IF EXISTS items;
DROP TABLE IF EXISTS payments;
DROP TABLE IF EXISTS orders;
-- +goose StatementEnd