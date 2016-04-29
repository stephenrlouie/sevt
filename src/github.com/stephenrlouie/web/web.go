package main

import (
    "fmt"
    "github.com/CiscoZeus/go-zeusclient"
    "net/http"
    "os"
    "runtime"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "ZEUS. MAC2. Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
    // print env variables
    fmt.Println("OS:", runtime.GOOS)
    fmt.Println("USER:", os.Getenv("USER"))
    fmt.Println("HOME:", os.Getenv("HOME"))
    fmt.Println("HOSTNAME:", os.Getenv("HOSTNAME"))

    // send env variables to zeus
    z := &zeus.Zeus{ApiServ: "http://api.ciscozeus.io", Token: os.Getenv("ZEUS_TOKEN")}
    logs := zeus.LogList{
        Name: "syslog",
        Logs: []zeus.Log{
            zeus.Log{
                "OS":       runtime.GOOS,
                "USER":     os.Getenv("USER"),
                "HOME":     os.Getenv("HOME"),
                "HOSTNAME": os.Getenv("HOSTNAME")},
        },
    }
    suc, err := z.PostLogs(logs)
    if err != nil {
        fmt.Println("ERROR: Zeus response:", err.Error())
    } else {
        fmt.Println("Zeus response:", suc)
    }

    total, logs, err := z.GetLogs("syslog", "", "", 0, 0, 0, 0)
    if err != nil {
        fmt.Println("ERROR: Logs zeus response:", err.Error())
    } else {
        fmt.Println("Logs zeus response: %s", logs)
        fmt.Println("Logs zeus response:", total)
    }

    fmt.Printf("Starting http server...\n")
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
 }

