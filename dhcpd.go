/**
    Simple DHCP Server
    Copyright (C) 2016  Jakub Wozniak

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <http://www.gnu.org/licenses/>.
**/

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



