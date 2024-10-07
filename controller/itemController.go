package controller

import (
	"encoding/json"
	"net/http"
)

func ItemController(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("All movies Deleted.")
}
