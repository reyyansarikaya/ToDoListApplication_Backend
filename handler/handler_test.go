//go:build unit_test
// +build unit_test

package handler_test

import (
	"bootcamp/handler"
	"bootcamp/mock"
	"bootcamp/model"
	"bytes"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type saveRequestBody struct {
	Todo string `json:"todo"`
}

func TestHandler_Get(t *testing.T) {
	responseWriter := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/api/v1/todo", http.NoBody)

	mockService := mock.NewMockIService(gomock.NewController(t))
	mockToDo := []model.ToDoModel{
		{ID: 0, Todo: "buy some milk"},
		{ID: 1, Todo: "go to swim"},
	}
	mockService.EXPECT().Get().Return(mockToDo).Times(1)

	h := handler.NewHandler(mockService)
	h.HandlerEndpoints(responseWriter, request)

	responseBody, _ := json.Marshal(mockToDo)

	assert.Equal(t, http.StatusOK, responseWriter.Result().StatusCode)
	assert.Equal(t, string(responseBody), responseWriter.Body.String())
}

func TestHandler_Save(t *testing.T) {
	t.Run("Succesfuly save todo", func(t *testing.T) {
		saveReqBody := &saveRequestBody{
			Todo: "buy some milk",
		}
		body, _ := json.Marshal(saveReqBody)

		responseWriter := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/api/v1/todo", bytes.NewReader(body))

		mockService := mock.NewMockIService(gomock.NewController(t))
		mockService.EXPECT().Save("buy some milk").Return(http.StatusOK).Times(1)

		h := handler.NewHandler(mockService)
		h.HandlerEndpoints(responseWriter, request)

		assert.Equal(t, http.StatusOK, responseWriter.Result().StatusCode)

	})
}
