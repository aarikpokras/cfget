package main

import (
  "fmt"
  "net/http"
  "io"
  "os"
  "log"
)

func main() {
  if len(os.Args) < 2 {
    fmt.Println("Usage: " + os.Args[0] + " [tags]")
    os.Exit(2)
  }
  if (os.Args[1] == "-s") {
    rsp, err := http.Get("https://raw.githubusercontent.com/aarikpokras/cfget/refs/heads/master/snippets/" + os.Args[2])
    if err != nil {
      log.Fatal(err)
    }
    defer rsp.Body.Close()
    if rsp.StatusCode == http.StatusNotFound {
      fmt.Println("No snippet found for " + os.Args[2] + " :(")
      os.Exit(1)
    }
    if rsp.StatusCode != http.StatusOK {
      fmt.Println("HTTP error " + rsp.Status)
    }
    bdy, err := io.ReadAll(rsp.Body)
    fmt.Println(string(bdy))
  } else {
    fmt.Println("Invalid flag.")
    os.Exit(4)
  }
}
