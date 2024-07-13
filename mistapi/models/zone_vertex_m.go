package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ZoneVertexM represents a ZoneVertexM struct.
type ZoneVertexM struct {
    // x in pixel
    X                    float64        `json:"x"`
    // y in pixel
    Y                    float64        `json:"y"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ZoneVertexM.
// It customizes the JSON marshaling process for ZoneVertexM objects.
func (z ZoneVertexM) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(z.toMap())
}

// toMap converts the ZoneVertexM object to a map representation for JSON marshaling.
func (z ZoneVertexM) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, z.AdditionalProperties)
    structMap["x"] = z.X
    structMap["y"] = z.Y
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ZoneVertexM.
// It customizes the JSON unmarshaling process for ZoneVertexM objects.
func (z *ZoneVertexM) UnmarshalJSON(input []byte) error {
    var temp zoneVertexM
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
    
    z.AdditionalProperties = additionalProperties
    z.X = *temp.X
    z.Y = *temp.Y
    return nil
}

// zoneVertexM is a temporary struct used for validating the fields of ZoneVertexM.
type zoneVertexM  struct {
    X *float64 `json:"x"`
    Y *float64 `json:"y"`
}

func (z *zoneVertexM) validate() error {
    var errs []string
    if z.X == nil {
        errs = append(errs, "required field `x` is missing for type `Zone_Vertex_M`")
    }
    if z.Y == nil {
        errs = append(errs, "required field `y` is missing for type `Zone_Vertex_M`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
