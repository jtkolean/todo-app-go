package task

import (
	"database/sql"
	"encoding/json"
	"jtkolean/http/httputil"
	"jtkolean/task/model"
	"jtkolean/task/store"
	"net/http"
	"strings"
)

type router struct {
	store store.TaskStore
}

func New(db *sql.DB) *router {
	return &router{
		store: *store.NewTaskStore(db),
	}
}

func (this *router) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		this.get(w, r)
	case http.MethodPost:
		this.create(w, r)
	case http.MethodPut:
		this.update(w, r)
	case http.MethodDelete:
		this.delete(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (this *router) get(w http.ResponseWriter, r *http.Request) {
	if httputil.HandleNotAcceptable(w, r, "application/json") {
		return
	}

	t, err := this.store.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	j, err := json.Marshal(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(j))
}

func (this *router) create(w http.ResponseWriter, r *http.Request) {
	if httputil.HandleUnsupportedContentType(w, r, "application/json") {
		return
	}

	var t model.Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err := this.store.Create(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	j, err := json.Marshal(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(j))
}

func (this *router) update(w http.ResponseWriter, r *http.Request) {
	if httputil.HandleUnsupportedContentType(w, r, "application/json") {
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/task/")

	var t model.Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := this.store.Update(id, t); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (this *router) delete(w http.ResponseWriter, r *http.Request) {

	id := strings.TrimPrefix(r.URL.Path, "/task/")

	if err := this.store.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
