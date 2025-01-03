package did

import (
	"crypto/ecdsa"
	"time"
)

// DID represents a Decentralized Identity
type DID struct {
	ID           string           `json:"id"`
	PublicKey    *ecdsa.PublicKey `json:"publicKey"`
	Created      time.Time        `json:"created"`
	Updated      time.Time        `json:"updated"`
	ProofPurpose string           `json:"proofPurpose"`
	ProofValue   string           `json:"proofValue"`
}

// DIDDocument represents the public DID document
type DIDDocument struct {
	Context    []string          `json:"@context"`
	ID         string            `json:"id"`
	Controller string            `json:"controller"`
	VerifyKeys []VerificationKey `json:"verificationMethod"`
	Created    time.Time         `json:"created"`
	Updated    time.Time         `json:"updated"`
}

// VerificationKey represents a public key associated with the DID
type VerificationKey struct {
	ID           string `json:"id"`
	Type         string `json:"type"`
	Controller   string `json:"controller"`
	PublicKeyHex string `json:"publicKeyHex"`
}
