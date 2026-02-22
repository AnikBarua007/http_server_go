package main

import (
	"net/http"

	"github.com/AnikBarua007/http_server_go/internal/auth"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerDeleteChirps(w http.ResponseWriter, r *http.Request) {
	accessToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusForbidden, "invalid refresh token")
		return
	}
	userID, err := auth.ValidateJWT(accessToken, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusForbidden, "invalid token")
		return
	}
	chirpstr := r.PathValue("chirpID")
	chirpID, err := uuid.Parse(chirpstr)
	chirp, err := cfg.dbQueries.GetOneChirp(r.Context(), chirpID)

	if chirp.UserID != userID || err != nil {
		respondWithError(w, http.StatusNotFound, "chirp not found")
		return
	}
	err = cfg.dbQueries.DeleteChirp(r.Context(), chirpID)
	w.WriteHeader(http.StatusNoContent)
}
