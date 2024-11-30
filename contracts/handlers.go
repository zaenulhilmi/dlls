package contracts

import (
	"net/http"
)

type AuthHandler interface {
	SignUp(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

type SubscriptionHandler interface {
	Subscribe(w http.ResponseWriter, r *http.Request)
}

type ActionHandler interface {
	Like(w http.ResponseWriter, r *http.Request)
	Pass(w http.ResponseWriter, r *http.Request)
	PreviewNextProfile(w http.ResponseWriter, r *http.Request)
}
