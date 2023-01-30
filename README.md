# Assorted Jollof

AssortedJollof is a repo of one same basic REST api server.
Rewritten in 5 languages.

The 6 languages are going to be typescript, go, python, php, java and rust.

The server will persist and read relational data to/from a mysqlite3 database.

## Design

This is just CRUD for the following data

| Model    | Attributes                                            |
| -------- | ----------------------------------------------------- |
| User     | id, name, phone                                       |
| Order    | id, packs, userId (User), location, status            |
| Payments | id, amount, orderId (Order), timestamp, userId (User) |
