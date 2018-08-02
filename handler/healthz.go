package handler

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// HealthzHandler instance
type HealthzHandler struct{}

// NewHealthzHandler initializes HealtzHandler instance
func NewHealthzHandler() *HealthzHandler {
	return &HealthzHandler{}
}

// Healthz is to check service healty
func (h *HealthzHandler) Healthz(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "ok")
}
