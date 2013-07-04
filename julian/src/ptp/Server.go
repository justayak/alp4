package ptp

import (
	"fmt"
	"net/http"
	"time"
)

// Startfunktion initalisiert alle notwendigen Werte
func Start() {
	fmt.Println("start Userbase..")
	users = make( map[ string ] *User, 0)
	
	// Behandle Webanfragen
	http.HandleFunc("/", handler)
	
	// Starte den Webserver
	http.ListenAndServe(":8080",nil)
	
	go update()
}

// U S E R B A S E
// Die Userbase verwaltet die User. Wir brauchen sie,
// um eine vernünftige Darstellung zu erzeugen (html)
var users map [ string ] *User

// Liefert den User mit dem entsprechenden Namen
// Existiert der User noch nicht, wird ein neuer erzeugt 
// und zufaellig mit den schon vorhanden verbunden
func GetUser(name string )*User {
	result, ok := users[name]
	if !ok {
		result := NewUser(name, "lol")
		// verbinde mit anderen User
		
		users[name] = result
		
		fmt.Println("Create new User <" + name + ">")
	}	
	return result
}

// ~~~~~~~~~~~~~~~~~
// S E R V E R
// ~~~~~~~~~~~~~~~~~

func handler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/favicon.ico" {
		// da immer zwei Handler gecalled werden (kp warum)
		name:=r.URL.Path[1:]
		fmt.Println(">>" + name)
		
		//user := GetUser(name)
		// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	}
}

func update() {
	for {
		
		time.Sleep(time.Second)
	}
}