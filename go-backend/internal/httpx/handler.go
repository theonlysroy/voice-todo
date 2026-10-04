package httpx

import (
	"errors"
	"log/slog"
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

func Wrap(h HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := h(w, r)
		if err == nil {
			return
		}
		var apiErr *ApiError
		if _, ok := errors.AsType[*ApiError](apiErr); !ok {
			apiErr = Internal(err)
		}
		if apiErr.Status >= 500 {
			slog.Error("request failed", "err", err, "path", r.URL.Path)
		}
		Fail(w, apiErr)
	}
}
