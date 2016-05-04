package main

import "fmt"
import "os"
import "log"

func main() {
  fmt.Printf("Lightweight DHCPD Server\n")
  fmt.Printf("Network Programming 2016\n")
  fmt.Printf("Jakub Wozniak\n")

  uid := os.Getuid()

  if uid != 0 {
    log.Fatal("You have to be root to run this server")
    os.Exit(1)
  }  
}



