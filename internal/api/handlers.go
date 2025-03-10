package api

import (
	"encoding/json"
	"monitoring-app/internal/domain"
	"monitoring-app/internal/repository"
	"net/http"
	"strconv"
)

type APIHandler struct {
	Service *domain.WebsiteService
	Repo    *repository.Repository
}

func (h *APIHandler) CheckWebsiteHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "missing url parameter", http.StatusBadRequest)
		return
	}

	result := h.Service.CheckWebsite(url)
	err := h.Repo.SaveResult(r.Context(), domain.Website{
		URL:        result.URL,
		StatusCode: result.StatusCode,
		Duration:   result.Duration,
		CheckedAt:  result.CheckedAt,
		Error:      result.Error,
	})

	if err != nil {
		http.Error(w, "failed to save result", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, "Ошибка кодирования JSON", http.StatusInternalServerError)
		return
	}
}

func (h *APIHandler) GetResultsHandler(w http.ResponseWriter, r *http.Request) {
	limitStr := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}

	results, err := h.Repo.GetLastResults(r.Context(), limit)
	if err != nil {
		http.Error(w, "failed to get results", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(results); err != nil {
		http.Error(w, "Ошибка кодирования JSON", http.StatusInternalServerError)
		return
	}
}

func (h *APIHandler) GetResultByURLHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "missing url parameter", http.StatusBadRequest)
		return
	}

	result, err := h.Repo.GetResultByURL(r.Context(), url)
	if err != nil {
		http.Error(w, "failed to get result", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, "Ошибка кодирования JSON", http.StatusInternalServerError)
		return
	}
}

func (h *APIHandler) DeleteOldResultsHandler(w http.ResponseWriter, r *http.Request) {
	err := h.Repo.DeleteOldResults(r.Context())
	if err != nil {
		http.Error(w, "failed to delete old results", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *APIHandler) CountResultsHandler(w http.ResponseWriter, r *http.Request) {
	count, err := h.Repo.CountResults(r.Context())
	if err != nil {
		http.Error(w, "failed to count results", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]int{"count": count}); err != nil {
		http.Error(w, "Ошибка кодирования JSON", http.StatusInternalServerError)
		return
	}
}
