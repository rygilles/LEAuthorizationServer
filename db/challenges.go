package db

import (
	"LEAuthorizationServer/models"
)

var challenges []model.Challenge

// Insert allows populating database
func InsertChallenge(challenge model.Challenge) {
	challenges = append(challenges, challenge)
}

// Update token matching item in database
func UpdateChallenge(challenge model.Challenge) {
	for key := range challenges {
		if challenges[key].Token == challenge.Token {
			challenges[key].Value = challenge.Value
		}
	}
}

// Get returns the whole database
func GetChallenges() []model.Challenge {
	return challenges
}

// Get the token matching model
func GetChallenge(token string) *model.Challenge {
	for _, item := range challenges {
		if item.Token == token {
			return &item
		}
	}

	return nil
}
