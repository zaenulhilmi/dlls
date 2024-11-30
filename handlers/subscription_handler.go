package handlers

import (
	"dlls/contracts"
	"encoding/json"
	"net/http"
)

func NewSubscriptionHandler(
	subscriptionService contracts.SubscriptionService,
) contracts.SubscriptionHandler {
	return &subscriptionHandler{
		subscriptionService: subscriptionService,
	}
}

type subscriptionHandler struct {
	subscriptionService contracts.SubscriptionService
}

// Subscribe implements contracts.SubscriptionHandler.
func (s *subscriptionHandler) Subscribe(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("authenticated_user").(*contracts.User)

	err := s.subscriptionService.Subscribe(user.ID)

	if err != nil {
		writeJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Subscribed successfully",
	})
}
