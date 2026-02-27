package todo

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	store *Store
	tmpl  *template.Template
}

func NewHandler(store *Store, templateDir string) *Handler {
	tmpl := template.Must(template.ParseGlob(templateDir + "/*.html"))
	return &Handler{store: store, tmpl: tmpl}
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Items   []*Item
		Total   int
		Pending int
	}{
		Items:   h.store.All(),
		Total:   h.store.Count(),
		Pending: h.store.Pending(),
	}
	h.tmpl.ExecuteTemplate(w, "index.html", data)
}

func (h *Handler) Add(w http.ResponseWriter, r *http.Request) {
	title := strings.TrimSpace(r.FormValue("title"))
	if title != "" {
		h.store.Add(title)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *Handler) Toggle(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err == nil {
		h.store.Toggle(id)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err == nil {
		h.store.Delete(id)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
