BEGIN;

CREATE TABLE IF NOT EXISTS public.public_data (
	id BIGSERIAL PRIMARY KEY,
	uuid uuid NOT NULL UNIQUE,

	name varchar,
	email VARCHAR(255) NOT NULL UNIQUE,
	birthdate TIMESTAMPTZ NOT NULL,

	acc_balance numeric(10, 10) NOT NULL,
	acc_balance_null numeric(10, 10), -- For "currency" precision

	user_credentials_id INT8 NOT NULL,

	hobbies VARCHAR[],

    important_notes jsonb NOT NULL,
	additional_info jsonb NULL
);

COMMIT;
