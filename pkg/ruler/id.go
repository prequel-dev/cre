package ruler

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
	"github.com/google/uuid"
)

type IdCmd struct {
	Num int ` optional:"" arg:"" default:"1" name:"num" help:"Number to generate."`
}

func (c *IdCmd) Run() error {
	for i := 0; i < c.Num; i++ {
		id, err := randomId()
		if err != nil {
			return err
		}
		fmt.Println(id)
	}
	return nil
}

func randomId() (string, error) {

	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	uuidBytes := id[:]
	encoded := base58.Encode(uuidBytes)

	return encoded, nil
}

func hashRule(data any) (string, error) {
	// json.Marshal to produce deterministic output
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	hash := sha256.Sum256(jsonBytes)

	return base58.Encode(hash[:]), nil
}
