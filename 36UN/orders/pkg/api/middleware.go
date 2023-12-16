package api

import (
    "net/http"
)

// headersMiddleware устанавливает заголовки ответа сервера.
func (api *API) headersMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        next.ServeHTTP(w, r)
    })
}