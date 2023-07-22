BEGIN;

CREATE TABLE public.public_data (
	id BIGSERIAL PRIMARY KEY,
	uuid uuid NULL UNIQUE,

	email VARCHAR(255) NOT NULL UNIQUE,
	name varchar,
    additional_info VARCHAR[],
	birthdate TIMESTAMPTZ NOT NULL,
    acc_balance_null numeric(10, 10), -- For "currency" precision
	acc_balance numeric(10, 10) NOT NULL,
    user_credentials_id INT8 NOT NULL,
	comments_null jsonb NULL,
    comments jsonb NOT NULL
);

COMMIT;
