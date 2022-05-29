# Crimson

Features:

1. Inventory Management System

2. User authentication

3. Web app interface

## Todo

- [ ] Add `User` service in project.
- [ ] Add `Authentication` service in project.
- [ ] Create tests.
- [ ] Add Email sender function.
- [x] Add Date Validation when sending `Add/Update` product request.

## Changes

5/23/2022

- Removed `user` service and `email` service to refactor code.

5/24/2022

- Refactored the code even more by removing the `user` model and `mailer` templates.

- Created validation function for `Product Tracking Number`.

5/25/2022

- Removed colorized outputs for logging and used `level` for debugging logs.

5/26/2022

- Added `datetime` field validation to product struct.

5/27/2022

- Implemented date validation in `Add/Update Product` query.

- Modified the http routes in `http.go` for the users using the API. User needs to specify HTTP method to be used for the endpoints to do `CRUD` functions.

- Updated migrations for `prod_svc` database.
