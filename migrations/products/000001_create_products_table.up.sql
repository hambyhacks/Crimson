CREATE TABLE IF NOT EXISTS products (
    id BIGSERIAL PRIMARY KEY,
    product_name text NOT NULL,
    declared_price real NOT NULL,
    shipping_fee real NOT NULL,
    tracking_number VARCHAR(255) NOT NULL,
    seller_name VARCHAR(255) NOT NULL,
    seller_address VARCHAR(255) NOT NULL,
    date_ordered timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    date_received timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    payment_mode VARCHAR(255) NOT NULL,
    stock_count BIGINT NOT NULL
);