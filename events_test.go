package events

import (
	"net/http"
	"net/http/httptest"
	"bytes"
	"testing"
	"reflect"

	"github.com/lunny/tango"
)

type EventAction struct {
	tango.Ctx
}

func (c *EventAction) Get() {
	c.Write([]byte("get"))
}

func (c *EventAction) Before() {
	c.Write([]byte("before "))
}

func (c *EventAction) After() {
	c.Write([]byte(" after"))
}

func TestEvent(t *testing.T) {
	buff := bytes.NewBufferString("")
	recorder := httptest.NewRecorder()
	recorder.Body = buff

	tg := tango.Classic()
	tg.Use(Events())
	tg.Get("/", new(EventAction))

	req, err := http.NewRequest("GET", "http://localhost:8000/", nil)
	if err != nil {
		t.Error(err)
	}

	tg.ServeHTTP(recorder, req)
	expect(t, recorder.Code, http.StatusOK)
	refute(t, len(buff.String()), 0)
	expect(t, buff.String(), "before get after")
}

/* Test Helpers */
func expect(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}

func refute(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		t.Errorf("Did not expect %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}