package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"week3_docker/internal/schemas"
)

func (h Handler) ContactActionsHook(w http.ResponseWriter, r *http.Request, keys map[string]string) {
	id, ok := keys["id"]
	if !ok {
		http.Error(w, "id not set", http.StatusBadRequest)
	}

	unsignedID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var req schemas.ContactActionsHookRequest
	req.ID = unsignedID

	if err = r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	jsonS, err := FormUrlEncodeToJSON(r.PostForm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err := json.Unmarshal([]byte(jsonS), &req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = h.ContactService.ContactActionsHook(r.Context(), req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
