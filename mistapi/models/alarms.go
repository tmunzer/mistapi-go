package models

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"strings"
)

// Alarms represents a Alarms struct.
type Alarms struct {
	AlarmIds []uuid.UUID `json:"alarm_ids"`
	// Some text note describing the intent
	Note                 *string        `json:"note,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for Alarms.
// It customizes the JSON marshaling process for Alarms objects.
func (a Alarms) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(a.toMap())
}

// toMap converts the Alarms object to a map representation for JSON marshaling.
func (a Alarms) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, a.AdditionalProperties)
	structMap["alarm_ids"] = a.AlarmIds
	if a.Note != nil {
		structMap["note"] = a.Note
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Alarms.
// It customizes the JSON unmarshaling process for Alarms objects.
func (a *Alarms) UnmarshalJSON(input []byte) error {
	var temp tempAlarms
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "alarm_ids", "note")
	if err != nil {
		return err
	}

	a.AdditionalProperties = additionalProperties
	a.AlarmIds = *temp.AlarmIds
	a.Note = temp.Note
	return nil
}

// tempAlarms is a temporary struct used for validating the fields of Alarms.
type tempAlarms struct {
	AlarmIds *[]uuid.UUID `json:"alarm_ids"`
	Note     *string      `json:"note,omitempty"`
}

func (a *tempAlarms) validate() error {
	var errs []string
	if a.AlarmIds == nil {
		errs = append(errs, "required field `alarm_ids` is missing for type `alarms`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
