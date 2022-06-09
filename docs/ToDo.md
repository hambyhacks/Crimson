# Todo

- [ ] Add `User` service in project.
- [ ] Add `Authentication` service in project.
- [ ] Create tests.
- [ ] Add Email sender function.
- [x] <strike>Add Date Validation when sending `Add/Update` product request. </strike>
- [ ] Add filtering, sorting and pagination to the API.
- [x] <strike>Error check on duplicate entries in database. </strike> (*used `ON CONFLICT DO NOTHING` clause on `insert` query.*)
- [ ] Implement full-text search for fetching records in database.
- [ ] Add `product name search` instead of `product id search`.
- [ ] Correctly implement HTTP status codes in HTTP response.

## Ideas

1. Make URL request for products should not be easily guessed. (*e.g products/1*)

    - add UUID column.

    - product should be searchable using UUID not the primary key id.
