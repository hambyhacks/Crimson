ALTER TABLE products ADD CONSTRAINT id_check CHECK (id > 0);

ALTER TABLE products ADD CONSTRAINT product_name_check CHECK (length(product_name) BETWEEN 1 AND 70);

ALTER TABLE products ADD CONSTRAINT stock_count_check CHECK (stock_count > 0);

ALTER TABLE products ADD CONSTRAINT tracking_number_check CHECK (length(tracking_number) BETWEEN 1 AND 50);

ALTER TABLE products ADD CONSTRAINT declared_price_check CHECK (declared_price > 0.0);

ALTER TABLE products ADD CONSTRAINT shipping_fee_check CHECK (shipping_fee > 0.0);