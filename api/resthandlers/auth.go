package resthandlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/theMillenniumFalcon/microservices/api/restutil"
	"github.com/theMillenniumFalcon/microservices/pb"
	"gopkg.in/mgo.v2/bson"
)

type AuthHandlers interface {
	SignUp(w http.ResponseWriter, r *http.Request)
	SignIn(w http.ResponseWriter, r *http.Request)
	PutUser(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	GetUsers(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

type authHandlers struct {
	authSvcClient pb.AuthServiceClient
}

func NewAuthHandlers(authSvcClient pb.AuthServiceClient) AuthHandlers {
	return &authHandlers{authSvcClient: authSvcClient}
}

func (h *authHandlers) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		restutil.WriteError(w, http.StatusBadRequest, restutil.ErrEmptyBody)
		return
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		restutil.WriteError(w, http.StatusBadRequest, err)
		return
	}

	user := new(pb.User)
	err = json.Unmarshal(body, user)
	if err != nil {
		restutil.WriteError(w, http.StatusBadRequest, err)
		return
	}

	user.Created = time.Now().Unix()
	user.Updated = user.Created
	user.Id = bson.NewObjectId().Hex()
	resp, err := h.authSvcClient.SignUp(r.Context(), user)
	if err != nil {
		restutil.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	restutil.WriteAsJson(w, http.StatusCreated, resp)
}
