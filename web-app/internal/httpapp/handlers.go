package httpapp

import (
	"encoding/json"
	"html/template"
	"net/http"

	"web-app/internal/store"
)

type Handlers struct {
	Store *store.MemoryStore
	Tmpl  *template.Template
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_ = h.Tmpl.ExecuteTemplate(w, "home.html", map[string]any{
		"Title": "Go Web App",
	})
}

func (h *Handlers) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}

func (h *Handlers) ListContacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{
		"data": h.Store.ListContacts(),
	})
}
