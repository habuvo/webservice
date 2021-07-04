package handlers

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

func WriteJSONObject(w http.ResponseWriter, o interface{}) {
	w.Header().Set("Content-Type", "application/json")
	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	err := e.Encode(o)
	if err != nil {
		WriteAPIError(w, http.StatusInternalServerError, "fail to serialize")
		zap.L().Warn("failed to serialize object", zap.Error(err))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func WriteCSVObject(w http.ResponseWriter, o [][]string) {
	w.Header().Set("Content-Type", "text/csv")
	e := csv.NewWriter(w)
	err := e.WriteAll(o)
	if err != nil {
		WriteAPIError(w, http.StatusInternalServerError, "fail to serialize")
		zap.L().Warn("failed to serialize object", zap.Error(err))
	}
	w.WriteHeader(http.StatusOK)
}

type APIError struct {
	Result string `json:"result,omitempty"`
	Error  string `json:"error"`
}

func WriteError(w http.ResponseWriter, statusCode int, o interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	err := e.Encode(o)
	if err != nil {
		zap.L().Warn("failed to serialize error", zap.Error(err))
	}
}

func WriteAPIError(w http.ResponseWriter, statusCode int, error string) {
	WriteError(w, statusCode, &APIError{Error: error})
}

func WriteAPIErrorCode(w http.ResponseWriter, statusCode int, result, error string) {
	WriteError(w, statusCode, &APIError{Result: result, Error: error})
}

func WriteAPIErrorf(w http.ResponseWriter, statusCode int, errorFmt string, args ...interface{}) {
	WriteError(w, statusCode, &APIError{Error: fmt.Sprintf(errorFmt, args...)})
}
