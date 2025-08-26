-- +goose Up
CREATE TABLE "orders"
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
	"date_created"       time,
	"oof_shard"          integer
);

CREATE TABLE "deliveries"
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

CREATE TABLE "payments"
(
	"transaction"   varchar PRIMARY KEY,
	"request_id"    varchar,
	"currency"      varchar(3),
	"provider"      varchar,
	"amount"        integer,
	"payment_dt"    timestamp,
	"bank"          varchar,
	"delivery_cost" integer,
	"goods_total"   integer,
	"custom_fee"    integer
);

CREATE TABLE "items"
(
	"chrt_id"      integer,
	"track_number" varchar,
	"price"        integer,
	"rid"          varchar,
	"name"         varchar,
	"sale"         integer,
	"size"         integer,
	"total_price"  integer,
	"nm_id"        integer,
	"brand"        varchar,
	"status"       int,
	PRIMARY KEY (track_number, nm_id)
);

ALTER TABLE "deliveries"
	ADD FOREIGN KEY ("order_uid") REFERENCES "orders" ("uid");

ALTER TABLE "payments"
	ADD FOREIGN KEY ("transaction") REFERENCES "orders" ("uid");

ALTER TABLE "items"
	ADD FOREIGN KEY ("track_number") REFERENCES "orders" ("track_number");


-- +goose Down
DROP TABLE deliveries;
DROP TABLE items;
DROP TABLE payments;
DROP TABLE orders;
