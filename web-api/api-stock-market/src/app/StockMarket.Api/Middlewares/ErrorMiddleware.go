package Middlewares

import (
	response "api-stock-market/src/app/StockMarket.Api/Response/Errors"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			if err := recover(); err != nil {
				var msg string
				if e, ok := err.(error); ok {
					msg = e.Error()
				} else {
					msg = fmt.Sprintf("%v", err)
				}
				log.Printf("Middleware caught error: %v", err)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				err := json.NewEncoder(w).Encode(response.ErrorInternal{

					Error:   "Internal server error",
					Code:    http.StatusInternalServerError,
					Message: msg,
				})
				if err != nil {
					return
				}

			}
		}()

		next.ServeHTTP(w, r)
	})
}
