package main

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"log"
	"net/http"
	"os"
)

var l *ldap.Conn

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)

	fmt.Println("Server listing to :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return
	}

	l, err = ldap.DialURL("ldap://127.0.0.1:389")
	if err != nil {
		log.Fatal(err)
	}
	//defer l.Close()

	fmt.Println("Binded !")
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	err := l.Bind("cn=superadmin,ou=users,dc=yunohost,dc=org", os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	_, err = fmt.Fprintf(w, "Hello de Armel")
	if err != nil {
		return
	}
}
