package loafer

import (
	"encoding/json"
	"fmt"
)

type JSONFormatter struct{}

func (j *JSONFormatter) Format(f Fields) ([]byte, error) {
	json, err := json.Marshal(f)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal fields to JSON, %v", err)
	}

	return append(json, '\n'), nil
}
