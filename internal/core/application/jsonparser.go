package application

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/EvertonTomalok/ports-challenge/internal/core/domain"
	"github.com/EvertonTomalok/ports-challenge/internal/ports"
)

type JsonParser struct {
	service ports.Service
}

func NewJsonParser(service ports.Service) *JsonParser {
	return &JsonParser{
		service: service,
	}
}

// ParseAndUpsertFile processes the JSON file line by line, performing an upsert operation
// for each item encountered. Given the file's large size, instead of loading the entire
// document at once, it will be processed in a paginated manner, returning each JSON node
// individually as it is found until all nodes have been parsed.
func (parser *JsonParser) ParseAndUpsertFile(filePath string) error {
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
		if err := parser.service.Upsert(portData); err != nil {
			return err
		}
	}
	return nil
}
