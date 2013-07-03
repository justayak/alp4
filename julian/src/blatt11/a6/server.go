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

// ctor
func NewUser(name string ) *User {
	return &User{
		name, 
		make( []User, 0),
	}
}

// U S E R B A S E
var users map[ string ] *User

func StartUserbase() {
	

func GetUser(name string )*User {
	result, ok := users[name]
	if !ok {
		result := NewUser(name)
		users[name] = result
	}	
	return result
}

// ~~~~~~~~~~~~~~~~~
// S E R V E R
// ~~~~~~~~~~~~~~~~~

func handler(w http.ResponseWriter, r *http.Request){
	
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
  http.HandleFunc("/", handler)
  GetUser("hallo")
 // http.ListenAndServe(":8080",nil)  
  
}