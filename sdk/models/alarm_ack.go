package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// AlarmAck represents a AlarmAck struct.
type AlarmAck struct {
    AlarmIds             []uuid.UUID    `json:"alarm_ids"`
    // Some text note describing the intent
    Note                 *string        `json:"note,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for AlarmAck.
// It customizes the JSON marshaling process for AlarmAck objects.
func (a AlarmAck) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the AlarmAck object to a map representation for JSON marshaling.
func (a AlarmAck) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    structMap["alarm_ids"] = a.AlarmIds
    if a.Note != nil {
        structMap["note"] = a.Note
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AlarmAck.
// It customizes the JSON unmarshaling process for AlarmAck objects.
func (a *AlarmAck) UnmarshalJSON(input []byte) error {
    var temp alarmAck
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

// alarmAck is a temporary struct used for validating the fields of AlarmAck.
type alarmAck  struct {
    AlarmIds *[]uuid.UUID `json:"alarm_ids"`
    Note     *string      `json:"note,omitempty"`
}

func (a *alarmAck) validate() error {
    var errs []string
    if a.AlarmIds == nil {
        errs = append(errs, "required field `alarm_ids` is missing for type `Alarm_Ack`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
