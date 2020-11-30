package main

import (

       "io"

        "net/http"

        "github.com/bmizerany/pat"

        "log"

)

// hello world, the web server

func HelloServer(w http.ResponseWriter, req *http.Request) {

     io.WriteString(w, "Guten Tag, "+req.URL.Query().Get(":name")+"!\n")

}

// Nachricht an Server senden

func NachrichtSenden(w http.ResponseWriter, req *http.Request) {

     io.WriteString(w, "Nachricht ist: " + req.URL.Query().Get(":message")+"\n")

     io.WriteString(w, "Uhrzeit: "+req.URL.Query().Get(":uhrzeit")+"!\n")

     io.WriteString(w, "Absender: "+req.URL.Query().Get(":absender")+"!\n")

}

func main() {

     m := pat.New() // öffentlich schreibt man Gross

     m.Get("/api/v1/hello/:name", http.HandlerFunc(HelloServer))

     // zweite route

     // eine nachricht an den chat serversenden

     // 

     m.Get("/api/v2/nachricht/:message/:uhrzeit/:absender", http.HandlerFunc(NachrichtSenden))

     m.Get("/api/v2/listen/:message/:uhrzeit/:absender", http.HandlerFunc(NachrichtSenden))

     m.Get("/api/v2/loeschen/:message/:uhrzeit/:absender", http.HandlerFunc(NachrichtSenden))

     // Register this pat with the default serve mux so that other packages

     // may also be exported. (i.e. /debug/pprof/*)

     http.Handle("/api/", m)

     

     err := http.ListenAndServe(":80", nil)

     if err != nil {

         log.Fatal("ListenAndServe: ", err)

    }

}
