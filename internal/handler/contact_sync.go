package handler

import (
	"net/http"
	"week3_docker/internal/schemas"
)

func (h Handler) ContactSync(w http.ResponseWriter, r *http.Request, keys map[string]string) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var req schemas.ContactSyncRequest
	if err = decoder.Decode(&req, r.PostForm); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = h.ContactService.PrimaryContactsSync(r.Context(), req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
