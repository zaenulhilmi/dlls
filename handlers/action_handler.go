package handlers

import (
	"dlls/contracts"
	"encoding/json"
	"fmt"
	"net/http"
)

func NewActionHandler(
	actionService contracts.ActionService,
	userService contracts.UserService,
) contracts.ActionHandler {
	return &actionHandlerImpl{
		actionService: actionService,
		userService:   userService,
	}
}

type actionHandlerImpl struct {
	actionService contracts.ActionService
	userService   contracts.UserService
}

// Like implements contracts.ActionHandler.
func (a *actionHandlerImpl) Like(w http.ResponseWriter, r *http.Request) {
	var request actionRequest

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		writeJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := r.Context().Value("authenticated_user").(*contracts.User)

	fmt.Println("user.ID", user.ID)
	fmt.Println("request.TargetID", request.TargetID)
	err = a.actionService.Like(user.ID, request.TargetID)

	if err != nil {
		writeJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Liked successfully",
	})
}

// Pass implements contracts.ActionHandler.
func (a *actionHandlerImpl) Pass(w http.ResponseWriter, r *http.Request) {
	var request actionRequest

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		writeJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := r.Context().Value("authenticated_user").(*contracts.User)

	err = a.actionService.Pass(user.ID, request.TargetID)

	if err != nil {
		writeJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Passed successfully",
	})
}

// PreviewProfile implements contracts.ActionHandler.
func (a *actionHandlerImpl) PreviewNextProfile(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("authenticated_user").(*contracts.User)

	targetID, err := a.actionService.NextTarget(user.ID)

	if err != nil {
		writeJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	target, err := a.userService.FindByID(targetID)

	if err != nil {
		writeJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(profileResponse{
		ID:   target.ID,
		Name: target.Name,
	})

}

type actionRequest struct {
	TargetID string `json:"target_id"`
}

type profileResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
}

