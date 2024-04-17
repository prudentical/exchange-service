CREATE TABLE public.exchanges (
    id bigserial PRIMARY KEY,
    created_at timestamp with time zone NULL,
    updated_at timestamp with time zone NULL,
    deleted_at timestamp with time zone NULL,
    name text UNIQUE NOT NULL,
    description text NULL,
    website text NULL,
    documentation_url text NULL,
    api_url text NULL,
    status text NULL
);
CREATE INDEX idx_exchanges_del ON public.exchanges USING HASH (deleted_at);

CREATE TABLE public.currencies (
    id bigserial PRIMARY KEY,
    created_at timestamp with time zone NULL,
    updated_at timestamp with time zone NULL,
    deleted_at timestamp with time zone NULL,
    name text NULL,
    symbol text UNIQUE NOT NULL
);
CREATE INDEX idx_currencies_del ON public.currencies USING HASH (deleted_at);

CREATE TABLE public.pairs (
    id bigserial PRIMARY KEY,
    created_at timestamp with time zone NULL,
    updated_at timestamp with time zone NULL,
    deleted_at timestamp with time zone NULL,
    base_id bigint NULL,
    quote_id bigint NULL,
    exchange_id bigint NULL,
    symbol text UNIQUE NOT NULL
);
CREATE INDEX idx_pairs_del ON public.pairs USING HASH (deleted_at);