package services_test

import (
	"dlls/services"
	"testing"
)

func TestActionService_Like(t *testing.T) {

	actionService := services.NewActionService()

	err := actionService.Like("user-id", "target-id")

	if err != nil {
		t.Errorf("ActionService.Like() error = %v, wantErr %v", err, false)
	}

}


