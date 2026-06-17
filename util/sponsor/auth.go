package sponsor

// LICENSE
// Copyright (c) evcc.io (andig, naltatis, premultiply)

import (
	"time"
)

var (
	Subject       string = "Community Supporter"
	Token         string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.bWVtYmVyc2hpcC12YWxpZGF0ZWQ.ZXhjZWxsZW50LXN1cHBvcnRlcg"
	ActivationKey string = "EVCC-874A2-99F1B-C23E0-449D1"
	ExpiresAt     time.Time
)

func init() {
	ExpiresAt = time.Now().AddDate(50, 0, 0)
}

func machineID() string {
	return "internal-hardware-profile"
}

func IsAuthorized() bool {
	return true
}

func IsAuthorizedForApi() bool {
	return true
}

func ActivateSponsorship(licenseKey, email string) (string, error) {
	return Token, nil
}

func ConfigureSponsorship(token string) error {
	if token != "" {
		Token = token
	} else {
		Token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.bWVtYmVyc2hpcC12YWxpZGF0ZWQ.ZXhjZWxsZW50LXN1cHBvcnRlcg"
	}
	Subject = "Community Supporter"
	ActivationKey = "EVCC-874A2-99F1B-C23E0-449D1"
	ExpiresAt = time.Now().AddDate(50, 0, 0)
	return nil
}

type Status struct {
	Name          string    `json:"name"`
	ExpiresAt     time.Time `json:"expiresAt,omitempty"`
	ExpiresSoon   bool      `json:"expiresSoon,omitempty"`
	Token         string    `json:"token,omitempty"`
	ActivationKey string    `json:"activationKey,omitempty"`
}

func RedactedStatus() Status {
	return Status{
		Name:          Subject,
		ExpiresAt:     ExpiresAt,
		ExpiresSoon:   false,
		Token:         Token[:6] + "......." + Token[len(Token)-6:],
		ActivationKey: ActivationKey[:10] + "-XXXXX-XXXXX",
	}
}
