package main

import (
  "fmt"
  "net/http"
  "io"
  "os"
  "log"
  "strconv"
)

func hideerrors() bool {
  return os.Getenv("CFGET_HIDE_ERRORS") == "1"
}

func main() {
  if len(os.Args) < 3 {
    if !hideerrors() {
      fmt.Println("Usage: " + os.Args[0] + " [options] [tags]\nOPTIONS:\n-s			Get first snippet.\n--ctmax [count]		Get [count] snippets.\n--gos [count]		get the [count]th snippet.")
    }
    os.Exit(2)
  }
  if (os.Args[1] == "-s") {
    rsp, err := http.Get("https://raw.githubusercontent.com/aarikpokras/cfget/refs/heads/master/snippets/" + os.Args[2] + "0")
    if err != nil {
      log.Fatal(err)
    }
    defer rsp.Body.Close()
    if rsp.StatusCode == http.StatusNotFound {
      if !hideerrors() {
        fmt.Println("No snippet found for " + os.Args[2] + " :(\nMake sure you're being as specific as possible; for example, instead of using hyprland/blur, try hyprland/decoration/blur.")
      }
      os.Exit(1)
    }
    if rsp.StatusCode != http.StatusOK {
      if !hideerrors() {
        fmt.Println("HTTP error " + rsp.Status)
      }
    }
    bdy, err := io.ReadAll(rsp.Body)
    fmt.Println(string(bdy))
  } else if (os.Args[1] == "--ctmax") {
    if len(os.Args) < 4 {
      if !hideerrors() {
        fmt.Println("--ctmax requires a count and tags.")
      }
      os.Exit(4)
    }
    ctmax, err := strconv.Atoi(os.Args[2])
    if err != nil {
      if !hideerrors() {
        fmt.Printf("Error in argument 2: %v\n", err)
      }
      os.Exit(8)
    }
    for i := 0; i < ctmax; i++ {
      rsp, err := http.Get("https://raw.githubusercontent.com/aarikpokras/cfget/refs/heads/master/snippets/" + os.Args[3] + strconv.Itoa(i))
      if err != nil {
        log.Fatal(err)
      }
      defer rsp.Body.Close()
      if rsp.StatusCode == http.StatusNotFound {
        fmt.Println("--------------------\nDone with snippet list. " + strconv.Itoa(i) + " snippets found.")
        os.Exit(0)
      }
      if rsp.StatusCode != http.StatusOK {
        if !hideerrors() {
          fmt.Println("HTTP error " + rsp.Status)
        }
      }
      bdy, err := io.ReadAll(rsp.Body)
      fmt.Println(string(bdy))
      fmt.Println("********************")
    }
  } else if (os.Args[1] == "--gos") {
    if len(os.Args) < 4 {
      if !hideerrors() {
        fmt.Println("--gos requires a count and tags.")
      }
      os.Exit(4)
    }
    _, err := strconv.Atoi(os.Args[2])
    if err != nil {
      if !hideerrors() {
        fmt.Println("Argument 2 is not a number.")
      }
      os.Exit(8)
    }
    rsp, err := http.Get("https://raw.githubusercontent.com/aarikpokras/cfget/refs/heads/master/snippets/" + os.Args[3] + os.Args[2])
    if err != nil {
      log.Fatal(err)
    }
    defer rsp.Body.Close()
    if rsp.StatusCode == http.StatusNotFound {
      if !hideerrors() {
        fmt.Println("Could not find that snippet.")
      }
      os.Exit(1)
    }
    if rsp.StatusCode != http.StatusOK {
      if !hideerrors() {
        fmt.Println("HTTP error " + rsp.Status)
      }
      os.Exit(1)
    }
    bdy, err := io.ReadAll(rsp.Body)
    fmt.Println(string(bdy))
  } else {
    if !hideerrors() {
      fmt.Printf("Invalid flag! Ensure counts and flags are separated and that you're using correct flags!\n\033[0;34mFrom exception in main -- unrecognized argument\033[0m\n")
    }
    os.Exit(4)
  }
}
