package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// MapNodePosition represents a MapNodePosition struct.
type MapNodePosition struct {
    X                    float64        `json:"x"`
    Y                    float64        `json:"y"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MapNodePosition.
// It customizes the JSON marshaling process for MapNodePosition objects.
func (m MapNodePosition) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MapNodePosition object to a map representation for JSON marshaling.
func (m MapNodePosition) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
    structMap["x"] = m.X
    structMap["y"] = m.Y
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MapNodePosition.
// It customizes the JSON unmarshaling process for MapNodePosition objects.
func (m *MapNodePosition) UnmarshalJSON(input []byte) error {
    var temp mapNodePosition
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "x", "y")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.X = *temp.X
    m.Y = *temp.Y
    return nil
}

// mapNodePosition is a temporary struct used for validating the fields of MapNodePosition.
type mapNodePosition  struct {
    X *float64 `json:"x"`
    Y *float64 `json:"y"`
}

func (m *mapNodePosition) validate() error {
    var errs []string
    if m.X == nil {
        errs = append(errs, "required field `x` is missing for type `Map_Node_Position`")
    }
    if m.Y == nil {
        errs = append(errs, "required field `y` is missing for type `Map_Node_Position`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
