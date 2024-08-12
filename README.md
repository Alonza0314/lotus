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
    import "github.com/Alonza0314/lotus"
    ```

2. Prepare your own pem using in TLS communication.
3. Init a lotus server.
4. Write your service function and decide this funtion's identifier.
5. Register this function to lotus server.
6. Call listen function to get lotus Listener.
7. Use a for loop to call accept function to accept the client's coinnection.
8. Whenever there exist a new connection, use go routine to handle it. Be caerful that you need to pass the lotus server information since the function is register in the lotus server structure

### Client

## Author

You can know more about the author through this link: [Alonza](https://alonza0314.github.io/)
