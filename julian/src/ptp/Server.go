package ptp

import (
	"fmt"
	"net/http"
)

// U S E R B A S E
var users map[ string ] *User

// Die Userbase verwaltet die User. Wir brauchen sie,
// um eine vernünftige Darstellung zu erzeugen (html)
func StartUserbase() {
	fmt.Println("start Userbase..")
	users = make( map[ string ] *User, 0)
}

// Liefert den User mit dem entsprechenden Namen
// Existiert der User noch nicht, wird ein neuer erzeugt 
// und zufaellig mit den schon vorhanden verbunden
func GetUser(name string )*User {
	result, ok := users[name]
	if !ok {
		result := NewUser(name)
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
	
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

/* func main() {
  http.HandleFunc("/", handler)
  
  StartUserbase()
  GetUser("hallo")
  GetUser("hallo")
 // http.ListenAndServe(":8080",nil)  
  
} */