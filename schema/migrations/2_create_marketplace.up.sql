BEGIN;

CREATE TABLE IF NOT EXISTS marketplace.user(
	id BIGSERIAL PRIMARY KEY,
	name VARCHAR(100),
	hash_password VARCHAR NOT NULL,
	email VARCHAR(100) UNIQUE,
	last_updated_password timestamptz default NOW()
);

CREATE TABLE IF NOT EXISTS marketplace.user_pocket(
	user_id INTEGER NOT NULL REFERENCES marketplace.user(id),
	balance INT8 NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS marketplace.product(
	id BIGSERIAL PRIMARY KEY,
	name VARCHAR(100) not null,
	price INTEGER NOT NULL,
	stock INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS marketplace.cart(
	id BIGSERIAL PRIMARY KEY,
	user_id INTEGER NOT NULL REFERENCES marketplace.user(id),
	cart_details JSONB NOT null
);

CREATE TABLE IF NOT EXISTS marketplace.order(
	id BIGSERIAL PRIMARY KEY,
	user_id INTEGER NOT null REFERENCES tbl_user(id),
	product_checkout_json JSONB NOT NULL
);

COMMIT;
