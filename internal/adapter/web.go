package adapter

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github/michaellimmm/gooddata-demo/generated/analytics/v1"
	"github/michaellimmm/gooddata-demo/internal/usecases"
	"io"
	"log/slog"
	"net/http"
	"os"

	"github.com/rs/cors"
)

type httpServer struct {
	server  *http.Server
	usecase usecases.Usecases
}

func NewHttpServer(usecase usecases.Usecases) *httpServer {
	server := &http.Server{}

	return &httpServer{
		server:  server,
		usecase: usecase,
	}
}

func (h *httpServer) Run(addr string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /v1/login", h.login)
	mux.HandleFunc("POST /v1/register", h.register)
	mux.HandleFunc("POST /v1/token", h.getToken)

	handler := cors.Default().Handler(mux)
	handler = h.loggingMiddleware(handler)

	h.server.Addr = addr
	h.server.Handler = handler

	return h.server.ListenAndServe()
}

func (h *httpServer) Stop(ctx context.Context) error {
	return h.server.Shutdown(ctx)
}

func (h *httpServer) login(w http.ResponseWriter, r *http.Request) {
	req := analytics.LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response := ErrorResponse{Errors: err.Error()}
		http.Error(w, response.ToJsonString(), http.StatusBadRequest)
		return
	}

	res, err := h.usecase.Login(r.Context(), &req)
	if err != nil {
		response := ErrorResponse{Errors: err.Error()}
		response.WriteResponse(w, http.StatusUnprocessableEntity)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (h *httpServer) register(w http.ResponseWriter, r *http.Request) {
	req := analytics.RegisterAccountRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response := ErrorResponse{Errors: err.Error()}
		response.WriteResponse(w, http.StatusBadRequest)
		return
	}

	res, err := h.usecase.Register(r.Context(), &req)
	if err != nil {
		response := ErrorResponse{Errors: err.Error()}
		response.WriteResponse(w, http.StatusUnprocessableEntity)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (h *httpServer) getToken(w http.ResponseWriter, r *http.Request) {
	req := analytics.GetTokenRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response := ErrorResponse{Errors: err.Error()}
		response.WriteResponse(w, http.StatusBadRequest)
		return
	}

	kid := os.Getenv("KID")
	privateKey := os.Getenv("PRIVATE_KEY")
	token, err := usecases.GenerateToken(privateKey, usecases.TokenKey{
		Kid: kid,
		Sub: fmt.Sprintf("u_%s", req.TenantId),
	})
	if err != nil {
		response := ErrorResponse{Errors: err.Error()}
		response.WriteResponse(w, http.StatusUnprocessableEntity)
		return
	}

	res := &analytics.GetTokenResponse{
		AccessToken: token,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (h *httpServer) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var bodyBytes []byte
		var err error

		bodyBytes, err = io.ReadAll(r.Body)
		if err != nil {
			return
		}
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		if len(bodyBytes) > 0 {
			slog.Info("request_body", slog.String("request_body", string(bodyBytes)))
		}

		next.ServeHTTP(w, r)
	})
}
