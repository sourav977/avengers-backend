package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/sourav977/avengers-backend/models"
)

type httpHandler http.Handler

//AvengerValidatorMW validated Avenger struct
func AvengerValidatorMW(next httpHandler) httpHandler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("in AvengerValidatorMW")
		defer log.Println("AvengerValidatorMW ended")
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()
		//we can use 3rd party pkgs to validate struct
		//for this example lets do the manual way
		var avenger models.Avenger
		//decode req body into avenger
		dec := json.NewDecoder(r.Body).Decode(&avenger)
		if dec != nil {
			http.Error(w, "Parse Error, required a valid json-request body", http.StatusBadRequest)
			return
		}
		//empty string not allowed
		if len(strings.TrimSpace(avenger.Name)) == 0 || len(strings.TrimSpace(avenger.Alias)) == 0 || len(strings.TrimSpace(avenger.Weapon)) == 0 {
			http.Error(w, "Error: Name, Alias, Weapon can not be empty", http.StatusBadRequest)
			return
		}
		/*
			stringonlyregex := regexp.MustCompile(`^[a-zA-Z]$`)
			if !stringonlyregex.MatchString(avenger.Name) || !stringonlyregex.MatchString(avenger.Alias) || !stringonlyregex.MatchString(avenger.Weapon) {
				http.Error(w, "Error: Name, Alias, Weapon must be String", http.StatusBadRequest)
				return
			}
		*/
		log.Printf("request received: %+v\n", avenger)
		next.ServeHTTP(w, r)
	})
}
