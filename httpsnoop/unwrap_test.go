package httpsnoop

import (
	"testing"

	"github.com/hxx258456/ccgo/gmhttp/httptest"
)

func TestUnwrap(t *testing.T) {
	w := Wrap(httptest.NewRecorder(), Hooks{})
	if _, ok := Unwrap(w).(*httptest.ResponseRecorder); !ok {
		t.Error("expected ResponseRecorder")
	}
}

func TestUnwrapWithoutWrap(t *testing.T) {
	w := httptest.NewRecorder()
	if _, ok := Unwrap(w).(*httptest.ResponseRecorder); !ok {
		t.Error("expected ResponseRecorder")
	}
}

func TestUnwrapMultipleLayers(t *testing.T) {
	w := Wrap(Wrap(httptest.NewRecorder(), Hooks{}), Hooks{})
	if _, ok := Unwrap(w).(*httptest.ResponseRecorder); !ok {
		t.Error("expected ResponseRecorder")
	}
}
