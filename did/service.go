package did

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"did-twitter/blockchain"

	"github.com/google/uuid"
)

type DIDService struct {
	blockchain *blockchain.Blockchain
}

func NewDIDService() *DIDService {
	return &DIDService{
		blockchain: blockchain.NewBlockchain(),
	}
}

func (s *DIDService) CreateDID() (*DID, error) {
	// Generate new key pair
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("failed to generate key pair: %v", err)
	}

	// Create DID identifier
	didID := fmt.Sprintf("did:custom:%s", uuid.New().String())

	// Create proof value
	proofValue := hex.EncodeToString(privateKey.PublicKey.X.Bytes())

	did := &DID{
		ID:           didID,
		PublicKey:    &privateKey.PublicKey,
		Created:      time.Now(),
		Updated:      time.Now(),
		ProofPurpose: "authentication",
		ProofValue:   proofValue,
	}

	// Store DID document in blockchain
	didDoc := s.CreateDIDDocument(did)
	didDocBytes, err := json.Marshal(didDoc)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal DID document: %v", err)
	}

	err = s.blockchain.AddBlock(didDocBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to add DID to blockchain: %v", err)
	}

	return did, nil
}

func (s *DIDService) CreateDIDDocument(did *DID) *DIDDocument {
	pubKeyHex := hex.EncodeToString(did.PublicKey.X.Bytes()) + hex.EncodeToString(did.PublicKey.Y.Bytes())

	verifyKey := VerificationKey{
		ID:           did.ID + "#keys-1",
		Type:         "EcdsaSecp256k1VerificationKey2019",
		Controller:   did.ID,
		PublicKeyHex: pubKeyHex,
	}

	return &DIDDocument{
		Context:    []string{"https://www.w3.org/ns/did/v1"},
		ID:         did.ID,
		Controller: did.ID,
		VerifyKeys: []VerificationKey{verifyKey},
		Created:    did.Created,
		Updated:    did.Updated,
	}
}
