package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	check "gopkg.in/check.v1"
)

func Test(t *testing.T) {
	check.TestingT(t)
}

type ListBeerHandlerSuite struct {
	handler http.Handler
}

var _ = check.Suite(&ListBeerHandlerSuite{})

func (s *ListBeerHandlerSuite) SetUpSuite(c *check.C) {
	s.handler = &ListBeerHandler{}
}

func (s *ListBeerHandlerSuite) TestOK(c *check.C) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodGet, "/beer", nil)

	s.handler.ServeHTTP(w, r)

	c.Assert(w.Code, check.Equals, http.StatusOK)
	c.Assert(w.Body.String(), check.Equals, "OK")
}
