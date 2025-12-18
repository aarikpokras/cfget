package main

import (
  "fmt"
  "net/http"
  "io"
  "os"
  "log"
  "strconv"
)

func main() {
  if len(os.Args) < 3 {
    fmt.Println("Usage: " + os.Args[0] + " -s [tags]")
    os.Exit(2)
  }
  if (os.Args[1] == "-s") {
    rsp, err := http.Get("https://raw.githubusercontent.com/aarikpokras/cfget/refs/heads/master/snippets/" + os.Args[2] + "0")
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
  } else if (os.Args[1] == "--ctmax") {
    if len(os.Args) < 4 {
      fmt.Println("--ctmax requires a count and tags.")
      os.Exit(4)
    }
    ctmax, err := strconv.Atoi(os.Args[2])
    if err != nil {
      fmt.Printf("Error in argument 2: %v\n", err)
      os.Exit(8)
    }
    for i := 0; i < ctmax; i++ {
      rsp, err := http.Get("https://raw.githubusercontent.com/aarikpokras/cfget/refs/heads/master/snippets/" + os.Args[3] + strconv.Itoa(i))
      if err != nil {
        log.Fatal(err)
      }
      defer rsp.Body.Close()
      if rsp.StatusCode == http.StatusNotFound {
        fmt.Println("-------------------\nDone with snippet list. " + strconv.Itoa(i) + " snippets found.")
        os.Exit(1)
      }
      if rsp.StatusCode != http.StatusOK {
        fmt.Println("HTTP error " + rsp.Status)
      }
      bdy, err := io.ReadAll(rsp.Body)
      fmt.Println(string(bdy))
    }
  } else {
    fmt.Printf("Invalid flag! Ensure counts and flags are separated and that you're using correct flags!\n\e[0;34mFrom exception in main -- unrecognized argument\e[0m\n")
    os.Exit(4)
  }
}
