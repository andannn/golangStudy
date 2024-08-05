package handler

import (
	"example.com/internal/proto/response"
	"example.com/internal/util"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUserHandler(t *testing.T) {
	h := NewUserHandler(&mockUserService{})

	t.Run("GetUser failed invalid id", func(t *testing.T) {
		rec := httptest.NewRecorder()
		c := setUpEchoForGet("/users/:id", rec)
		c.SetParamNames("id")
		c.SetParamValues("&")

		if assert.NoError(t, h.GetUser(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("GetUser failed valid id", func(t *testing.T) {
		rec := httptest.NewRecorder()
		c := setUpEchoForGet("/users/:id", rec)
		c.SetParamNames("id")
		c.SetParamValues("2")

		if assert.NoError(t, h.GetUser(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	t.Run("Register valid data", func(t *testing.T) {
		userJSON := `{"name":"Jon Snow","email":"jon@labstack.com"}`
		rec := httptest.NewRecorder()
		c := setUpEchoForPost("/register", userJSON, rec)

		if assert.NoError(t, h.Register(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
}

type mockUserService struct{}

func (m *mockUserService) GetUserById(id int) (*response.UserResponse, error) {
	return &response.UserResponse{}, nil
}

func setUpEchoForGet(path string, rec *httptest.ResponseRecorder) echo.Context {
	req := httptest.NewRequest(http.MethodGet, path, nil)
	e := echo.New()
	c := e.NewContext(req, rec)
	return c
}

func setUpEchoForPost(path string, data string, rec *httptest.ResponseRecorder) echo.Context {
	req := httptest.NewRequest(http.MethodPost, path, strings.NewReader(data))
	e := echo.New()
	e.Validator = util.NewJSONValidator()
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	c := e.NewContext(req, rec)
	return c
}
