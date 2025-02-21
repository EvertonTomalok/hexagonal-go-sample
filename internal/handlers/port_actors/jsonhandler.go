package port_actor

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/EvertonTomalok/ports-challenge/internal/domain"
	"github.com/EvertonTomalok/ports-challenge/internal/ports"
)

type JsonActor struct {
	service ports.Service
}

func NewJsonActor(service ports.Service) *JsonActor {
	return &JsonActor{
		service: service,
	}
}

// HandleUpsertStream process the JSON file line by line and performs
// an upsert call for each item found.
func (actor *JsonActor) HandleUpsertStream(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("Error opening file: %v\n", err)
	}
	defer func() {
		if file != nil {
			_ = file.Close()
		}
	}()

	decoder := json.NewDecoder(file)

	// Move cursor to the opening brace `{`
	if _, err = decoder.Token(); err != nil {
		return fmt.Errorf("Error reading JSON opening brace: %v\n", err)
	}

	// this code below will read item by item inside the json file.
	for decoder.More() {
		keyToken, err := decoder.Token()
		if err != nil {
			return fmt.Errorf("Error reading port key: %v\n", err)
		}
		key := keyToken.(string)

		var value domain.Port
		if err := decoder.Decode(&value); err != nil {
			return fmt.Errorf("Error decoding port value: %v\n", err)
		}
		portData := domain.PortData{
			key: value,
		}
		if err := actor.service.Upsert(portData); err != nil {
			return err
		}
	}
	return nil
}
