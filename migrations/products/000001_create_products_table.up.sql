CREATE TABLE IF NOT EXISTS products (
    id BIGSERIAL PRIMARY KEY NOT NULL ,
    product_name text NOT NULL,
    price real NOT NULL,
    SKU VARCHAR(255) NOT NULL,
    date_ordered timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    date_received timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    stock_count BIGINT NOT NULL
);