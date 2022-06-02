# Updates

## 5/23/2022

- Removed `user` service and `email` service to refactor code.

## 5/24/2022

- Refactored the code even more by removing the `user` model and `mailer` templates.

- Created validation function for `Product Tracking Number`.

## 5/25/2022

- Removed colorized outputs for logging and used `level` for debugging logs.

## 5/26/2022

- Added `datetime` field validation to product struct.

## 5/27/2022

- Implemented date validation in `Add/Update Product` query.

- Modified the http routes in `http.go` for the users using the API. User needs to specify HTTP method to be used for the endpoints to do `CRUD` functions.

- Updated migrations for `prod_svc` database.

## 5/30/2022

- Added sequence reset in `DeleteProduct()` function for `id` column since it is the `primary key`.

- Removed `NOT NULL` in `primary key` query on `create_products_table_up` migration file. (*outputs error in inserting product in database.*)

## 5/31/2022

- Roughly added insert check when adding records to database (*see repo.go: line 41*).

## 6/1/2022

- Added `ON CONFLICT DO NOTHING` on `Add product` query. (*better insert check*)

## 6/2/2022

- Removed redundant error logging.
