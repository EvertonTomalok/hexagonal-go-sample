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
// Since the file receive is too large, instead of reading the whole document,
// we will "paginate" the file and return a json node every time it's found, one by one
// by all the nodes are parsed.
func (actor *JsonActor) HandleUpsertStream(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("couldn't open file: %v\n", err)
	}
	defer func() {
		if file != nil {
			_ = file.Close()
		}
	}()

	decoder := json.NewDecoder(file)

	// Move cursor to the opening brace `{`
	if _, err = decoder.Token(); err != nil {
		return fmt.Errorf("reading JSON opening brace raised an error: %v\n", err)
	}

	// this code below will read node by node inside the json file.
	for decoder.More() {
		keyToken, err := decoder.Token()
		if err != nil {
			return fmt.Errorf("reading port key raised an error: %v\n", err)
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
