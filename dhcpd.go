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
    "os"
    "log"
    "golang.org/x/sys/unix"
    )


// it will fail if uid is not 0
func CheckPrivileges(uid int) {
  if uid != 0 {
    log.Fatal("You have to be root to run this server")
    os.Exit(1)
  }
}


// Create UDP socket, bind it to device and set some options
func BuildSocket(device string) int {
  fd, err := unix.Socket(unix.AF_INET, unix.SOCK_DGRAM, 0)
  if err != nil {
    log.Fatal("Cannot create socket: %v", err)
    os.Exit(1)
  }

  err = unix.BindToDevice(fd, device)
  if err != nil {
    log.Fatal("Cannot bind to device %v: %v", device, err)
    os.Exit(1)
  }

  err = unix.SetsockoptInt(fd, unix.SO_REUSEADDR, 1)
  if err != nil {
    log.Fatal("Cannot set REUSEADDR: %v", err)
    os.Exit(1)
  }

  err = unix.SetsockoptInt(fd, unix.SO_BROADCAST, 1)
  if err != nil {
    log.Fatal("Cannot set BROADCAST: %v", err)
    os.Exit(1)
  }


  return fd
}



func main() {
  log.Print("Simple DHCP Server")
  log.Print("Copyright (C) 2016 Jakub Wozniak")
  log.Print("Distributed Systems, Poznan University of Technology")

  CheckPrivileges(os.Getuid())
  
  // create socket
  sfd := BuildSocket("wlan0")
  log.Print(sfd)
}
