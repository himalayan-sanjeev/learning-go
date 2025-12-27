package httpapp

import "net/http"

func NewRouter(h *Handlers) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", h.Home)
	mux.HandleFunc("/health", h.Health)
	mux.HandleFunc("/api/contacts", h.ListContacts)

	return mux
}
