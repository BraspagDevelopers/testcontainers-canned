# testcontainers-canned
Canned container for tests using the library [testcontainers-go](https://github.com/testcontainers/testcontainers-go).

What does it mean?
It means that if you have tests using the library _testcontainers-go_, you can easily spin up some specific applications without worring about the _waiting_ logic or _connection string/client url_-building logic.

You can also specify your own images to be initialized using the same logic as one that is already implemented here (i.e. if you create your own SQL Server image with your tables already).

## Included Containers
### General Purpose

#### SQL Server for Linux
- **Default Image**: `mcr.microsoft.com/mssql/server`