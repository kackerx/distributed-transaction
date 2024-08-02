package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Order struct {
	ProductID string `json:"product_id"`
	Count     int    `json:"count"`
}

func main() {
	http.HandleFunc("/order/add", func(w http.ResponseWriter, r *http.Request) {
		var order Order
		err := json.NewDecoder(r.Body).Decode(&order)
		if err != nil {
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "get: %v", order)
	})

	if err := http.ListenAndServe(":8888", nil); err != nil {
		return
	}
}
