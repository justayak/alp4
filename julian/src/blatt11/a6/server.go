/*
* (c) Programmierung in Go / Addison Wesley 2010
*     Rainer Feike / Steffen Blass
*
* Quellcode-Datei: 10_01_smallserver.go
* Beschreibung: Ein kleiner Webserver in Go
*/

package main

import (
  "fmt"        // enthält Formatierungsfunktionen
  "http"       // enthält den Servercode und das http Protokoll
  "os"         // zur Interaktion mit dem Betriebssystem
  "io/ioutil"  // Dateien lesen
  "time"       // Zur Logzeitausgabe
)

// ein globaler logchannel
var logchannel chan string

func loggerThread(c chan string) {
  // endlos die strings aus dem channel holen und 
  // auf der Konsole ausgeben
  for {
    s := <-c
    fmt.Println(s)
  }
}
  
func log(l string) {
  // die funktion stellt die logzeie in die queue
  // erst mit der Zeit prefixen
  t := time.LocalTime()
  ts := fmt.Sprintf("%02d.%02d.%04d %02d:%02d:%02d: ",
            t.Day, t.Month, t.Year, t.Hour, t.Minute, t.Second)
  logchannel <- ts+l
}

// Hier werden die Requests an den Server verarbeitet
func requesthandler(c *http.Conn, r *http.Request) {
  // Pfad aus der Url holen
  pfad := r.URL.Path[0:]
  // gemäß der üblichen Konvention mappen wir / auf index.html
  if pfad == "/" { pfad = "/index.html" }
  log("verarbeite request an ["+pfad+"]")
  
  // versuche die Datei zu laden
  content, err := ioutil.ReadFile(pfad[1:])
  if err != nil {
    // Fehler ausgeben und loggen
    e := fmt.Sprintf("404: page not found at %s : Error (%s)", 
             pfad, err.String())
    log(fmt.Sprintf("GET [%s] code 404 error %s", 
            pfad, err.String()))
    http.Error(c, e, http.StatusNotFound)
  } else {
    // Seite ausgeben und Erfolg loggen
    fmt.Fprint(c, string(content))
    log(fmt.Sprintf("GET [%s] code 200 size %d Bytes", 
            pfad, len(content)))
  }
}

func main() {
  // der Port auf dem unser Server laufen soll
  port := "8080"
  // das Basisverzeichnis (rootdir) unseres Servers
  dir := "."
  
  // Kommandozeile prüfen und Werte übernehmen
  if len(os.Args) > 1 {
    port = os.Args[1]
  }
  if len(os.Args) > 2 {
    dir = os.Args[2]
  }
  
  // logkanal gepuffert eröffnen
  logchannel = make(chan string, 2000)
  
  // logging nebenläufig starten
  go loggerThread(logchannel)
  
  // und los gehts
  log("Der Go Smallserver startet auf Port "+port+" im Verzeichnis "+dir);

  // Server initialisieren
  os.Chdir(dir)
  http.HandleFunc("/", requesthandler)
  http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}