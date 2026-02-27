package todo_test

import (
	"testing"

	"github.com/example/auto-demo-go-app/todo"
)

func TestAddAndGet(t *testing.T) {
	s := todo.NewStore()
	item := s.Add("Buy milk")

	if item.Title != "Buy milk" {
		t.Errorf("got title %q, want %q", item.Title, "Buy milk")
	}
	if item.Done {
		t.Error("new item should not be done")
	}

	got, ok := s.Get(item.ID)
	if !ok {
		t.Fatal("item not found")
	}
	if got.Title != "Buy milk" {
		t.Errorf("got title %q, want %q", got.Title, "Buy milk")
	}
}

func TestAll(t *testing.T) {
	s := todo.NewStore()
	s.Add("Task 1")
	s.Add("Task 2")
	s.Add("Task 3")

	items := s.All()
	if len(items) != 3 {
		t.Errorf("got %d items, want 3", len(items))
	}
}

func TestToggle(t *testing.T) {
	s := todo.NewStore()
	item := s.Add("Task")

	if item.Done {
		t.Fatal("should start as not done")
	}

	ok := s.Toggle(item.ID)
	if !ok {
		t.Fatal("toggle returned false")
	}

	got, _ := s.Get(item.ID)
	if !got.Done {
		t.Error("should be done after toggle")
	}

	s.Toggle(item.ID)
	got, _ = s.Get(item.ID)
	if got.Done {
		t.Error("should be undone after second toggle")
	}
}

func TestToggleNotFound(t *testing.T) {
	s := todo.NewStore()
	if s.Toggle(999) {
		t.Error("toggle should return false for missing item")
	}
}

func TestDelete(t *testing.T) {
	s := todo.NewStore()
	item := s.Add("To delete")

	ok := s.Delete(item.ID)
	if !ok {
		t.Fatal("delete returned false")
	}
	if s.Count() != 0 {
		t.Errorf("got count %d, want 0", s.Count())
	}
}

func TestDeleteNotFound(t *testing.T) {
	s := todo.NewStore()
	if s.Delete(999) {
		t.Error("delete should return false for missing item")
	}
}

func TestPending(t *testing.T) {
	s := todo.NewStore()
	s.Add("Task 1")
	item2 := s.Add("Task 2")
	s.Add("Task 3")

	if s.Pending() != 3 {
		t.Errorf("got pending %d, want 3", s.Pending())
	}

	s.Toggle(item2.ID)
	if s.Pending() != 2 {
		t.Errorf("got pending %d, want 2", s.Pending())
	}
}
