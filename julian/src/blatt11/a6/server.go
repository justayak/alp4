/*
* (c) Programmierung in Go / Addison Wesley 2010
*     Rainer Feike / Steffen Blass
*
* Quellcode-Datei: 10_01_smallserver.go
* Beschreibung: Ein kleiner Webserver in Go
*/

package main

import (
  "fmt"        // enthaelt Formatierungsfunktionen
  "net/http"       // enthaelt den Servercode und das http Protokoll
  //"os"         // zur Interaktion mit dem Betriebssystem
  //"io/ioutil"  // Dateien lesen
  //"time"       // Zur Logzeitausgabe
)

// ~~~~~~~~~~~~~~~~~
// U S E R
// ~~~~~~~~~~~~~~~~~
type User struct {
	name string 
	others []User
}

func NewUser(name string) *User {
	user:=User{name, make([]User)}
}

// ~~~~~~~~~~~~~~~~~
// S E R V E R
// ~~~~~~~~~~~~~~~~~

func handler(w http.ResponseWriter, r *http.Request){
	
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
  http.HandleFunc("/", handler)
 // http.ListenAndServe(":8080",nil)  
  
}