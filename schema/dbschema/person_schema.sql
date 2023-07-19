CREATE TABLE public.person (
	id BIGSERIAL PRIMARY KEY,
	uuid uuid NULL UNIQUE,
	email VARCHAR(255) NOT NULL UNIQUE,

	birthdate TIMESTAMPTZ NOT NULL,
    acc_balance NUMERIC(10, 10), -- For "currency" precision
    additional_info VARCHAR[],
    user_credentials_id INT8 NOT NULL,
	comments_null jsonb NULL,
    comments jsonb NOT NULL
)
