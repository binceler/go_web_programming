package helpers

import (
	"fmt"
	"github.com/gorilla/sessions"
	"net/http"
)

var store = sessions.NewCookieStore([]byte("123123"))

func SetAlert(w http.ResponseWriter, r *http.Request, message string) error {
	session, err := store.Get(r, "go-alert")
	if err != nil {
		fmt.Println(err)
		return err
	}

	session.AddFlash(message)

	return session.Save(r, w)

}

func GetAlert(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "go-alert")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(session.Flashes())

}
