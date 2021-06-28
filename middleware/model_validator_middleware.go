package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/sourav977/avengers-backend/models"
)

type httpHandler http.Handler

//AvengerMW validated Avenger struct
func AvengerMW(next httpHandler) httpHandler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("in AvengerMW")
		defer log.Println("AvengerMW ended")
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
		stringonlyregex, _ := regexp.Compile(`^\[a-zA-Z\]$`)
		if !stringonlyregex.MatchString(avenger.Name) || !stringonlyregex.MatchString(avenger.Alias) || !stringonlyregex.MatchString(avenger.Weapon) {
			http.Error(w, "Error: Name, Alias, Weapon must be String", http.StatusBadRequest)
			return
		}
		fmt.Printf("request received: %+v\n", avenger)
		next.ServeHTTP(w, r)
	})
}
