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

	var err error
	l, err = ldap.DialURL("ldap://127.0.0.1:389")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	fmt.Println(l)
	fmt.Println("Binded !")

	fmt.Println("Server listing to :8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		return
	}
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	err := l.Bind("uid=superadmin,ou=users,dc=yunohost,dc=org", os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	res, err := l.WhoAmI(nil)
	if err != nil {
		return
	}

	_, err = fmt.Fprintf(w, "Hello de %v <br> %v", res.AuthzID, r.Header)
	if err != nil {
		return
	}
}
