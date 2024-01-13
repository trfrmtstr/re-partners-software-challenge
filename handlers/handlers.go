package handlers

import (
	"encoding/json"
	"net/http"
	"re-partners-take-home-assignment/packing"

	"github.com/go-playground/validator"
)

type findOptimalPackingRequestBody struct {
	PackSizes []int `json:"pack_sizes" validate:"required,min=1"`
	NumItems  int   `json:"num_items" validate:"required"`
}

// Parse and validate JSON body and respond with solution.
func FindOptimalPackingHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody findOptimalPackingRequestBody

	// NOTE: we could improve routing by using one of existing http routing libraries/frameworks
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestBody); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	var validate = validator.New()
	if err := validate.Struct(requestBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	responseBody, err := json.Marshal(packing.FindOptimalPacking(requestBody.NumItems, requestBody.PackSizes))
	if err != nil {
		http.Error(w, "Failed to marshal response body", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}
