# Assorted Jollof

AssortedJollof is a repo of one same basic REST api server.
Rewritten in 6 languages: typescript, go, python, php, java and rust.

The server will persist and read relational data to/from a mysqlite3 database.

## Design

This is just CRUD for the following data

| Model    | Attributes                                            |
| -------- | ----------------------------------------------------- |
| User     | id, name, email, password                             |
| Customer | id, name, phone, token                                |
| Order    | id, packs, userId (User), location, status            |
| Payments | id, amount, orderId (Order), timestamp, userId (User) |

## Implemented

### Typescript (Express)

- [x] models
- [x] routers and controllers
- [ ] filters
- [ ] auth

### Rust (Rocket)

- [-] models
- [-] routers and controllers
