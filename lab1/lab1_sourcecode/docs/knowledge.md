# Basic server client theory

- IP address = street address
- Port = apartment number
- TCP connection = telephone line between two apartments
- net.Listen() = telling the receptionist "send visitors for apartment 8080 to me"
- Accept() = answering the door when someone arrives
- Read() = listening to what they say
- Write() = replying
- Close() = hanging up

## Server (net.Listen())
A server is basically a process that waits for connections and then responds to them.

This is a server:
`ln, err := net.Listen("tcp", ":8000")`
Its saying:
"OS, notify whenever someone tries to connect to the TCP port 8000"

Now its a listening socket, no connection yet.

1. Go asks OS for socket
2. Os reserves TCP port 
3. OS monitors incoming packets destined for that port
4. OS queues new connection attempts

`conn, err := ln.Accept()`
Using the net.Listener from the net.Listen to block until somebody connects.

The OS completes the TCP handshake and returns a `net.Conn`
The `net.Conn` represents a connection, and essentially, a client.

For many concurrent connections (clients), you need to accept repeateldy:

```Go
// Infinite loop, each iteration, blocks (waits) until a new connection, Accepts it, and spawns a new goroutine to handle the connection.
for {
    conn, _ := ln.Accept()
    go handleConnection(conn)
}

```
each `conn` is a different client

## Client (net.Dial())

to connect to a server, clients use
`conn, _ := net.Dial("tcp", "127.0.0.1:8000)`
Basically means "I want a connection".


`net.Dial` is the opposite to `net.Listen`

1. Server listens on a port: `net.Listen` -> basically, a server is now created
2. A client requests a connection to the server: `net.Dial`
3. The server accepts the connection: `net.Accept`

a TCP handshake between server and client ensures both parties agree on parameters (destination, source, sequence numbers, buffers etc)

4. `net.Conn` exists on both sides (client = server connection), (server = a connected client)
-> Both can Read(), Write(), Close()

`net.Conn` is basically an established LIVE connection,
Read = Receive bytes
Write = send bytes
