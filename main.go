package main

import (
	"log"
	"net/http"
	"server-application3/candidats"
	"server-application3/parsing"
	"server-application3/respond"
)

func main() {
	http.HandleFunc("/about", respond.Respond)
	http.HandleFunc("/information", respond.Respond)
	http.HandleFunc("/education", respond.Respond)
	http.HandleFunc("/contact", respond.Respond)
	http.HandleFunc("/date", respond.Respond)
	http.HandleFunc("/relax", respond.Respond)
	http.HandleFunc("/submit", respond.Respond)
	http.HandleFunc("/menu", respond.Respond)
	http.HandleFunc("/parsing", parsing.ParsingHandler)
	http.HandleFunc("/candidate", candidats.AgentHandler)
	http.HandleFunc("/", respond.Respond)
	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
	// li, err := net.Listen("tcp", ":8080")
	// if err != nil {
	// log.Fatalln("Error starting server:", err)
	// }
	// defer li.Close()
	// log.Println("Server started on :8080")

	// for {
	// conn, err := li.Accept()
	// if err != nil {
	// log.Println("Accept error:", err)
	// continue
	// }
	// go handle.Handle(conn)
	// }
}
