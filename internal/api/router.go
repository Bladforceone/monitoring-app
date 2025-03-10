package api

import (
	"net/http"
)

func SetupRoutes(handler *APIHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/check", handler.CheckWebsiteHandler)
	mux.HandleFunc("/results", handler.GetResultsHandler)
	mux.HandleFunc("/result", handler.GetResultByURLHandler)
	mux.HandleFunc("/delete-old", handler.DeleteOldResultsHandler)
	mux.HandleFunc("/count", handler.CountResultsHandler)

	return mux
}
