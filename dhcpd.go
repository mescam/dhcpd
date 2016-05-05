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

import (
    "fmt"
    "os"
    "log"
    //"golang.org/x/sys/unix"
    )


func CheckPrivileges(uid int) {
  if uid != 0 {
    log.Fatal("You have to be root to run this server")
    os.Exit(1)
  }
}


func main() {
  log.Print("Simple DHCP Server")
  log.Print("Copyright (C) 2016 Jakub Wozniak")
  log.Print("Distributed Systems, Poznan University of Technology")

  CheckPrivileges(os.Getuid())
  fmt.Println(os.Getuid())
}
