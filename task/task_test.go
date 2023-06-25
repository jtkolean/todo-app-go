package task

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"jtkolean/config"
	"jtkolean/task/model"
	"net/http"
	"net/http/httptest"
	"testing"
)

var filename = "../resources/application.yaml"

func TestGet(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/task", nil)
	r.Header.Set("accept", "application/json")
	w := httptest.NewRecorder()

	c := config.NewConfig(filename)
	db := c.ConnectDB()
	New(db).Handle(w, r)

	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	var tasks = []model.Task{}
	if err1 := json.Unmarshal(data, &tasks); err1 != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	if len(tasks) < 1 {
		t.Errorf("expected todos length to be greater than or equal to 1 got %v", len(tasks))
	}
}

func TestGetForNotAcceptable(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/task", nil)
	r.Header.Set("accept", "text/csv")
	w := httptest.NewRecorder()

	c := config.NewConfig(filename)
	db := c.ConnectDB()
	New(db).Handle(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusNotAcceptable {
		t.Errorf("expected status code to be 406 but got %v", res.StatusCode)
	}
}

func TestCreate(t *testing.T) {
	r := httptest.NewRequest(http.MethodPost, "/task", bytes.NewBuffer([]byte(`{"title":"Test", "completed":0}`)))
	r.Header.Set("content-type", "application/json")
	w := httptest.NewRecorder()

	c := config.NewConfig(filename)
	db := c.ConnectDB()
	New(db).Handle(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status code to be 200 but got %v", res.StatusCode)
	}
}

func TestCreateForUnsupportedMediaType(t *testing.T) {
	r := httptest.NewRequest(http.MethodPost, "/task", nil)
	r.Header.Set("content-type", "text/csv")
	w := httptest.NewRecorder()

	c := config.NewConfig(filename)
	db := c.ConnectDB()
	New(db).Handle(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusUnsupportedMediaType {
		t.Errorf("expected status code to be 415 but got %v", res.StatusCode)
	}
}

func TestUpdate(t *testing.T) {
	r := httptest.NewRequest(http.MethodPut, "/task/3", bytes.NewBuffer([]byte(`{"title":"Test", "completed":1}`)))
	r.Header.Set("content-type", "application/json")
	w := httptest.NewRecorder()

	c := config.NewConfig(filename)
	db := c.ConnectDB()
	New(db).Handle(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status code to be 200 but got %v", res.StatusCode)
	}
}

func TestUpdateForUnsupportedMediaType(t *testing.T) {
	r := httptest.NewRequest(http.MethodPut, "/task/3", nil)
	r.Header.Set("content-type", "text/csv")
	w := httptest.NewRecorder()

	c := config.NewConfig(filename)
	db := c.ConnectDB()
	New(db).Handle(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusUnsupportedMediaType {
		t.Errorf("expected status code to be 415 but got %v", res.StatusCode)
	}
}

func TestDelete(t *testing.T) {
	r := httptest.NewRequest(http.MethodDelete, "/task/5", nil)
	w := httptest.NewRecorder()

	c := config.NewConfig(filename)
	db := c.ConnectDB()
	New(db).Handle(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status code to be 200 but got %v", res.StatusCode)
	}
}
