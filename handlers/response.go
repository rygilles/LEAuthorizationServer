package handler

import model "LEAuthorizationServer/models"

// JsonError is a generic error in JSON format
//
// swagger:response jsonError
type jsonError struct {
	// in: body
	Message string `json:"message"`
}

// ChallengeResponse contains a single challenge information
//
// swagger:response challengeResponse
type challengeResponse struct {
	// in: body
	Payload *model.Challenge `json:"challenge"`
}

// ChallengeValueResponse contains a single challenge value string information
//
// swagger:response challengeValueResponse
type challengeValueResponse struct {
	// in: body
	Payload string `json:"value"`
}

// ChallengesResponse contains all challengese from database information
//
// swagger:response challengesResponse
type challengesResponse struct {
	// in: body
	Payload *[]model.Challenge `json:"challenge"`
}
