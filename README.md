# LOTUS

## Description

This is a golang module to implement a new RPC technology based on QUIC.

The QUIC module is imported from: [quic-go](https://github.com/quic-go/quic-go).

## Development Environment

| Content | Implememt | Version |
|-|-|-|
| OS | Linux | Ubuntu 22.04.4 LTS  |
| Language | Golang | go1.22.5 linux/amd64 |
|PKI|TLS|1.3|

## How to use

### Server

1. Import this module

    ``` go
    import "github.com/Alonza0314/lotus/server"
    ```

2. Prepare your own pem using in TLS communication(QUIC required).

    ```go
    penPath := "test.pem"
    ```

3. Init a lotus server.

    ```go
    lserver, err := server.NewLotusServer("test.pem")
    ```

4. Write your service function and decide this funtion's identifier.

    ```go
    // identifier = "add"
    func add(a, b int) int {
        return a + b
    }
    ```

5. Register this function with its identifier to lotus server.

    ```go
    err = lserver.RegisterService("add", add)
    ```

6. Call listen function to get lotus Listener and asign the listening address and port.

    ```go
    llistener, err := lserver.Listen(":4433")
    ```

7. Use a for loop to call accept function to accept the client's coinnection.

    ```go
    lconn, err := llistener.Accept(context.Background())
    ```

8. Whenever there exists a new connection, use go routine to handle it. Be caerful that you need to pass the lotus server information since the function is register in the lotus server structure.

    ```go
    for {
        lconn, err := llistener.Accept(context.Background())
        if err != nil {
            // TODO
            continue
        }
        go lconn.HandleFunc(*lserver)
    }
    ```

### Client

## Author

You can know more about the author through this link: [Alonza](https://alonza0314.github.io/)
