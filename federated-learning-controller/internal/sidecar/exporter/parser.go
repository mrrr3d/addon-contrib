package exporter

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

// flexibleFloat is a custom type that can be unmarshalled from a JSON number or a string.
type flexibleFloat float64

type metricFilePayload struct {
	Metrics map[string]flexibleFloat `json:"metrics"`
	Labels  map[string]flexibleFloat `json:"labels"`
}

// UnmarshalJSON implements the json.Unmarshaler interface for the flexibleFloat type.
// It allows the value to be either a number or a string that can be parsed into a float.
func (f *flexibleFloat) UnmarshalJSON(data []byte) error {
	var num float64
	if err := json.Unmarshal(data, &num); err == nil {
		*f = flexibleFloat(num)
		return nil
	}

	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return fmt.Errorf("value must be a number or a string representing a number: %w", err)
	}

	parsedNum, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return fmt.Errorf("string value could not be parsed into a float: %w", err)
	}

	*f = flexibleFloat(parsedNum)
	return nil
}

// ParseContetnt parses the given content into a map of metrics and a map of labels.
// The content is expected to be a JSON byte array with a specific structure.
// It should have a "metrics" key and a "labels" key, both pointing to objects.
// The values in these objects should be numbers or strings that can be converted to float64.
// It returns the parsed metrics, the parsed labels, and an error if any occurs.
func ParseContetnt(content []byte) (map[string]float64, map[string]float64, error) {
	var payload metricFilePayload

	err := json.Unmarshal(content, &payload)
	if err != nil {
		log.Printf("JSON unmarshaling failed: %s", err)
		return nil, nil, err
	}

	metrics := make(map[string]float64, len(payload.Metrics))
	for key, value := range payload.Metrics {
		metrics[key] = float64(value)
	}

	labels := make(map[string]float64, len(payload.Labels))
	for key, value := range payload.Labels {
		labels[key] = float64(value)
	}

	return metrics, labels, nil
}
