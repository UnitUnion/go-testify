package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Если count > чем всего
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=5&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, responseRecorder.Code, http.StatusOK)
	b := responseRecorder.Body.String()
	l := strings.Split(b, ",")
	assert.Len(t, l, totalCount)
}

// Если city не поддерживается
func TestMainHandlerWhenCityNotSupport(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=1&city=s1pb", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, responseRecorder.Code, http.StatusBadRequest)
	assert.Equal(t, responseRecorder.Body.String(), "wrong city value")

}

// Запрос правильный, получаем 200 и не пустой body
func TestMainHandlerWhenRequestIsOk(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=3&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, responseRecorder.Code, http.StatusBadRequest)
	assert.NotEmpty(t, responseRecorder.Body)
}
