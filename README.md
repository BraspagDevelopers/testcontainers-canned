# testcontainers-canned
Canned container for tests using the library [testcontainers-go](https://github.com/testcontainers/testcontainers-go).

What does it mean?
It means that if you have tests using the library _testcontainers-go_, you can easily spin up some specific applications without worring about the _waiting_ logic or _connection string/client url_-building logic.

You can also specify your own images to be initialized using the same logic as one that is already implemented here (i.e. if you create your own SQL Server image with your tables already).

## Included Containers

### Generic API
Generic container starter for an HTTP API
- **Default Image:** None. It must be provided on creation.

### Mock-Server
- **Default Image:** `mockserver/mockserver`
- **Product Website:** https://www.mock-server.com/

### SQL Server for Linux
- **Default Image:** `mcr.microsoft.com/mssql/server`
- **Product Documentation:** https://docs.microsoft.com/en-us/sql/linux/quickstart-install-connect-docker


## More Features

### Network Creation
The package `./networks` includes a function to create a network with a random name.