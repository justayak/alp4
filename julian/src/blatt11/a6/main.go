/*
* (c) Programmierung in Go / Addison Wesley 2010
*     Rainer Feike / Steffen Blass
*
* Quellcode-Datei: 10_01_smallserver.go
* Beschreibung: Ein kleiner Webserver in Go
*/

package main

import (
  "fmt"    
  "ptp"
)

func main() {
	fmt.Println("start..")
	
	ptp.StartUserbase()
	ptp.GetUser("hallo")
	ptp.GetUser("hallo")
	
}