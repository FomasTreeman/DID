package main

import (
	"encoding/json"
	"fmt"
	"log"

	"did-twitter/did"
)

func main() {
	// Create new DID service
	didService := did.NewDIDService()

	// Create a new DID
	newDID, err := didService.CreateDID()
	if err != nil {
		log.Fatalf("Failed to create DID: %v", err)
	}

	// Create DID Document
	didDocument := didService.CreateDIDDocument(newDID)

	// Print DID Document as JSON
	docJSON, _ := json.MarshalIndent(didDocument, "", "  ")
	fmt.Printf("Created DID Document:\n%s\n", string(docJSON))
}
