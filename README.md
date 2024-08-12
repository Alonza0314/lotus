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

Use this command to get this module:

```bash
go get github.com/Alonza0314/lotus@latest
```

### Server

1. Import this module.

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

4. Write your service function and decide this funtion's identifier. Be careful that do not use "int" directly. Instead, use float64 to represent the number's type, since the json module does not distinguish "int" and "float64".

    ```go
    // identifier = "add"
    func add(a, b interface{}) float64 {
        return a.(float64) + b.(float64)
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

1. Import this module.

    ``` go
    import "github.com/Alonza0314/lotus/client"
    ```

2. Init a lotus client. The first arg is the address and the port. The other is a bool value to identity if the tls pem used by server is signed by it self or signed by public.

    ```go
    lclient, err := client.NewLotusClient(":4433", true)
    ```

3. Defer the close function.

    ```go
    defer lconn.Close()
    ```

4. Set the Service identifier and arges slice and the reply slice. Be careful that the type of args and reply is interface slice.

    ```go
    function, args, reply := "add", []interface{}{1, 2}, []interface{}{}
    ```

5. Call the service function.

    ```go
    err := lconn.Call(context.Background(), function, args, &reply)
    ```

6. Make a type assertion on the reply according to the function definition.

    ```go
    response := reply[0].(float64)
    ```

## Author

You can know more about the author through this link: [Alonza](https://alonza0314.github.io/)
