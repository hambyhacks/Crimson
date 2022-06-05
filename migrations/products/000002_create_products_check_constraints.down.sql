ALTER TABLE products DROP CONSTRAINT IF EXISTS product_name_check;

ALTER TABLE products DROP CONSTRAINT IF EXISTS stock_count_check;

ALTER TABLE products DROP CONSTRAINT IF EXISTS tracking_number_check;

ALTER TABLE products DROP CONSTRAINT IF EXISTS declared_price_check;

ALTER TABLE products DROP CONSTRAINT IF EXISTS shipping_fee_check;