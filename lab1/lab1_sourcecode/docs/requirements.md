# 1 HTTP SERVER
- net/http not allowed. only net package: https://pkg.go.dev/net


[] Accept HTTP requests (GET, POST) others -> "Not Implemented (501)"
[] Return response data from locally stored files to client
[] Handle concurrent requrests (create go routine for each request)

[X] Produce one binary that takes as first argument a port number

[] When server starts, establish a socket connection to use to listen for incoming client connections
[] Each new client is accepter and a go routine is spawned to handle request 