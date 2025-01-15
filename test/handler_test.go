package main

import (
	"bytes"
	"demo/first/Yd/internal/models"
	"demo/first/Yd/internal/transport"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Тест на успешное вычисление выражения
func TestCalculateHandler_Success(t *testing.T) {
	// Создаем тестовый запрос
	requestBody := models.Request{Expression: "2 + 3 * 4"}
	body, _ := json.Marshal(requestBody)
	req, err := http.NewRequest("POST", "/api/v1/calculate", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	// Создаем ResponseRecorder для записи ответа
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(transport.CalculateHandler)

	// Выполняем запрос
	handler.ServeHTTP(rr, req)

	// Проверяем статус код
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Проверяем тело ответа
	var response models.Response
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	expectedResult := 14.0
	if response.Result != expectedResult {
		t.Errorf("handler returned unexpected result: got %v want %v", response.Result, expectedResult)
	}
}

// Тест на некорректное выражение
func TestCalculateHandler_InvalidExpression(t *testing.T) {
	// Создаем тестовый запрос с некорректным выражением
	requestBody := models.Request{Expression: "2 + abc"}
	body, _ := json.Marshal(requestBody)
	req, err := http.NewRequest("POST", "/api/v1/calculate", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	// Создаем ResponseRecorder для записи ответа
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(transport.CalculateHandler)

	// Выполняем запрос
	handler.ServeHTTP(rr, req)

	// Проверяем статус код
	if status := rr.Code; status != http.StatusUnprocessableEntity {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusUnprocessableEntity)
	}

	// Проверяем тело ответа
	var response models.Response
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	expectedError := "Expression is not valid"
	if response.Error != expectedError {
		t.Errorf("handler returned unexpected error: got %v want %v", response.Error, expectedError)
	}
}

// Тест на внутреннюю ошибку сервера
func TestCalculateHandler_InternalError(t *testing.T) {
	// Создаем тестовый запрос с выражением, которое вызовет ошибку
	requestBody := models.Request{Expression: "2 / 0"}
	body, _ := json.Marshal(requestBody)
	req, err := http.NewRequest("POST", "/api/v1/calculate", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	// Создаем ResponseRecorder для записи ответа
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(transport.CalculateHandler)

	// Выполняем запрос
	handler.ServeHTTP(rr, req)

	// Проверяем статус код
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}

	// Проверяем тело ответа
	var response models.Response
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	expectedError := "Internal server error"
	if response.Error != expectedError {
		t.Errorf("handler returned unexpected error: got %v want %v", response.Error, expectedError)
	}
}

// Тест на некорректный формат запроса
func TestCalculateHandler_BadRequest(t *testing.T) {
    // Создаем тестовый запрос с некорректным телом
    body := []byte(`{"invalid": "data"}`)
    req, err := http.NewRequest("POST", "/api/v1/calculate", bytes.NewBuffer(body))
    if err != nil {
        t.Fatal(err)
    }

    // Создаем ResponseRecorder для записи ответа
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(transport.CalculateHandler)

    // Выполняем запрос
    handler.ServeHTTP(rr, req)

    // Проверяем статус код
    if status := rr.Code; status != http.StatusBadRequest {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
    }

    // Проверяем тело ответа
    var response models.Response
    err = json.Unmarshal(rr.Body.Bytes(), &response)
    if err != nil {
        t.Fatal(err)
    }

    expectedError := "Expression is not valid"
    if response.Error != expectedError {
        t.Errorf("handler returned unexpected error: got %v want %v", response.Error, expectedError)
    }
}