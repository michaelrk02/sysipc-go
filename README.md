# SysIPC (Go)
Filesystem-based interprocess communication (IPC) protocol implementation in Go (with JSON serializer).

Interoperability with another programming language or platform is possible by using the same protocol as described below.

## Server implementation
1. Before running the server, make sure there are no files associated with it (e.g. requests, responses, locks)
2. Wait for a client request then acquire its lock
3. Retrieve the contents of the request
4. Delete the request file, then release the lock
5. Process the request with the handler associated with method name
6. Lock the response
7. Write the response
8. Unlock the response, then repeat the same process to handle another incoming request

NOTE: The server runs single-threaded (i.e. not creating another thread on handling each request). It is also guaranteed to be safe across different client processes (through server mutual exclusion provided by client calls)

## Client implementation
1. Lock the server first
2. Generate a random call ID
3. Lock the client request
4. Write the request
5. Unlock the client request
6. Wait for the server to process the request while ensuring if the server is still locked
7. When a response is available, acquire its lock
8. Retrieve the contents of the response
9. Authenticate the call by comparing its request and response call ID's. If same, then delete the response file
10. Release the response lock
11. Unlock the server, then repeat the same process to make another call

## Specifications
- Server address is located at `<router-path>/<server-name>`
- Request file is formatted as `<server-address>.request`
- Response file is formatted as `<server-address>.response`
- Lock file is formatted as `<object>.lock` and `object` can be one of server address, request file, or response file
- The contents of the request body: `{call_id: [random 64-bit unsigned integer], method: [string], args: [key/value pair]}`
- The contents of the response body: `{call_id: [must be the same as the request call ID], return: [any], error: [string]}`

