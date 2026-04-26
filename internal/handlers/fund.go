package handlers

import (
	"net/http"

	"github.com/N30A/fondhav/internal/repository"
)

type FundHandler struct {
	repo *repository.FundRepository
}

func NewFundHandler(repo *repository.FundRepository) *FundHandler {
	return &FundHandler{repo}
}

func (h *FundHandler) GetFunds(w http.ResponseWriter, r *http.Request) {

}

func (h *FundHandler) GetFund(w http.ResponseWriter, r *http.Request) {

}

func (h *FundHandler) GetFundHoldings(w http.ResponseWriter, r *http.Request) {

}
