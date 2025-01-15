package transport

import (
	"demo/first/Yd/internal/models"
	"demo/first/Yd/internal/service"
	"encoding/json"
	"net/http"
)

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    var req models.Request
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        // Возвращаем 400, если JSON некорректен
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(models.Response{Error: "Invalid request format"})
        return
    }

    // Проверяем, что поле Expression не пустое
    if req.Expression == "" {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(models.Response{Error: "Expression is not valid"})
        return
    }

    // Вычисляем выражение
    result, err := service.CalculateExpression(req.Expression)
    if err != nil {
        if err.Error() == "Expression is not valid" {
            w.WriteHeader(http.StatusUnprocessableEntity)
            json.NewEncoder(w).Encode(models.Response{Error: "Expression is not valid"})
        } else {
            w.WriteHeader(http.StatusInternalServerError)
            json.NewEncoder(w).Encode(models.Response{Error: "Internal server error"})
        }
        return
    }

    // Возвращаем успешный ответ
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(models.Response{Result: result})
}
