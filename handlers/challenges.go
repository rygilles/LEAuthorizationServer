package handler

import (
	"LEAuthorizationServer/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"LEAuthorizationServer/db"
)

// TokenParam is used to identify a challenge
//
// swagger:parameters getChallenge getChallengeValue
type TokenParam struct {
	// The Token of a challenge
	//
	// in: path
	// required: true
	Token string `json:"token"`
}

// ChallengeModel Request Challenge model
//
// swagger:model challenge
type ChallengeModel struct {
	// The Token of a challenge
	//
	// required: true
	Token string `json:"token"`

	// The Value of a challenge
	//
	// required: true
	Value string `json:"value"`
}

// GetChallenges is an httpHandler for route GET /people
func GetChallenges(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /challenges challenges listChallenges
	//
	// Lists all challenges.
	//
	// This will show all recorded challenges.
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Responses:
	//       200: challengesResponse
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(db.GetChallenges())
}

// GetChallenge is an httpHandler for route GET /challenge/{token}
func GetChallenge(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /challenge/{token} challenges getChallenge
	//
	// Show token matching challenge.
	//
	// This will show the record of an identified challenge.
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Params:
	//       id: TokenParam
	//
	//     Responses:
	//       200: challengeResponse
	//       404: jsonError
	params := mux.Vars(r)
	item := db.GetChallenge(params["token"])

	if item != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(item); err != nil {
			panic(err)
		}
		return
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonError{Message: "Not Found"}); err != nil {
		panic(err)
	}
}

// GetChallengeValue is an httpHandler for route GET /.well-known/acme-challenge/{token}
func GetChallengeValue(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /.well-known/acme-challenge/{token} challenges getChallengeValue
	//
	// Show token matching challenge (ACME authorization check url).
	//
	// This will show the record of an identified challenge.
	//
	//     Produces:
	//     - text/plain
	//
	//     Schemes: http, https
	//
	//     Params:
	//       id: TokenParam
	//
	//     Responses:
	//       200: challengeValueResponse
	//       404: jsonError
	params := mux.Vars(r)
	item := db.GetChallenge(params["token"])

	if item != nil {
		w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(item.Value))
		return
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonError{Message: "Not Found"}); err != nil {
		panic(err)
	}
}

// PostChallenge is an httpHandler for route POST /challenge
func PostChallenge(w http.ResponseWriter, r *http.Request) {
	// swagger:operation POST /challenge challenges createChallenge
	//
	// Insert a challenge.
	//
	// This will add a challenge record.
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: challenge
	//   in: body
	//   description: Challenge
	//   schema:
	//     "$ref": "#/definitions/challenge"
	//   required: true
	// responses:
	//   '200':
	//     description: challengeResponse
	//   '404':
	//     description: jsonError

	var challenge model.Challenge

	if err := json.NewDecoder(r.Body).Decode(&challenge); err != nil {

		payload := map[string]string{"error": "Invalid request payload"}

		response, _ := json.Marshal(payload)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)

		return
	}

	if validErrs := challenge.Validate(); len(validErrs) > 0 {
		err := map[string]interface{}{"validationError": validErrs}
		response, _ := json.Marshal(err)
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)

		return
	}

	item := db.GetChallenge(challenge.Token)

	if item != nil {
		db.UpdateChallenge(challenge)
	} else {
		db.InsertChallenge(challenge)
	}

	response, _ := json.Marshal(challenge)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}