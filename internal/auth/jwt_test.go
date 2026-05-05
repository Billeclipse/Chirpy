package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestMakeAndValidateJWT(t *testing.T) {
	userID := uuid.New()
	secret := "super-secret"

	token, err := MakeJWT(userID, secret, time.Hour)
	if err != nil {
		t.Fatalf("MakeJWT returned error: %v", err)
	}

	gotID, err := ValidateJWT(token, secret)
	if err != nil {
		t.Fatalf("ValidateJWT returned error: %v", err)
	}

	if gotID != userID {
		t.Fatalf("expected %v, got %v", userID, gotID)
	}
}

func TestExpiredJWTRejected(t *testing.T) {
	userID := uuid.New()
	secret := "super-secret"

	token, err := MakeJWT(userID, secret, -time.Hour)
	if err != nil {
		t.Fatalf("MakeJWT returned error: %v", err)
	}

	_, err = ValidateJWT(token, secret)
	if err == nil {
		t.Fatal("expected expired token to be rejected")
	}
}

func TestWrongSecretRejected(t *testing.T) {
	userID := uuid.New()

	token, err := MakeJWT(userID, "correct-secret", time.Hour)
	if err != nil {
		t.Fatalf("MakeJWT returned error: %v", err)
	}

	_, err = ValidateJWT(token, "wrong-secret")
	if err == nil {
		t.Fatal("expected token signed with wrong secret to be rejected")
	}
}
