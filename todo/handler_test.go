package todo_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/example/auto-demo-go-app/todo"
)

func newTestHandler(t *testing.T) (*todo.Handler, *todo.Store) {
	t.Helper()
	store := todo.NewStore()
	handler := todo.NewHandler(store, "../templates")
	return handler, store
}

func TestListHandler(t *testing.T) {
	h, store := newTestHandler(t)
	store.Add("Test task")

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	h.List(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("got status %d, want %d", w.Code, http.StatusOK)
	}
	if !strings.Contains(w.Body.String(), "Test task") {
		t.Error("response should contain the task title")
	}
}

func TestAddHandler(t *testing.T) {
	h, store := newTestHandler(t)

	form := url.Values{"title": {"New task"}}
	req := httptest.NewRequest(http.MethodPost, "/add", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	h.Add(w, req)

	if w.Code != http.StatusSeeOther {
		t.Errorf("got status %d, want %d", w.Code, http.StatusSeeOther)
	}
	if store.Count() != 1 {
		t.Errorf("got count %d, want 1", store.Count())
	}
}

func TestAddHandlerEmpty(t *testing.T) {
	h, store := newTestHandler(t)

	form := url.Values{"title": {"  "}}
	req := httptest.NewRequest(http.MethodPost, "/add", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	h.Add(w, req)

	if store.Count() != 0 {
		t.Error("empty title should not create an item")
	}
}
