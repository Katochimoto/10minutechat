package main

import (
    // "fmt"
    // "flag"
    "log"
    "io"
    "net/http"
)

func index(res http.ResponseWriter, req *http.Request)  {
    res.Header().Set("Content-Type", "text/html")
    io.WriteString(
        res,
        `<doctype html>
<html>
    <head>
        <title>Hello World</title>
    </head>
    <body>
        Hello World!
    </body>
</html>`,
    )
}

func main()  {
    // var port = flag.Int("port", 10443, "port")
    // flag.Parse()

    http.HandleFunc("/", index)

    log.Printf("About to listen on 10443. Go to https://127.0.0.1:10443/")
    log.Fatal(http.ListenAndServeTLS(":10443", "cert/10minutechat.crt", "cert/10minutechat.key", nil))
}
