ALTER TABLE products ADD CONSTRAINT id_check CHECK (id > 0);

ALTER TABLE products ADD CONSTRAINT product_name_check CHECK (length(product_name) BETWEEN 1 AND 70);

ALTER TABLE products ADD CONSTRAINT stock_count_check CHECK (stock_count > 0);

ALTER TABLE products ADD CONSTRAINT sku_check CHECK (length(sku) BETWEEN 1 AND 50);

ALTER TABLE products ADD CONSTRAINT price_check CHECK (price > 0.0)