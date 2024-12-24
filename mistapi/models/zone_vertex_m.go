package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ZoneVertexM represents a ZoneVertexM struct.
type ZoneVertexM struct {
    // x in pixel
    X                    float64                `json:"x"`
    // y in pixel
    Y                    float64                `json:"y"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ZoneVertexM,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (z ZoneVertexM) String() string {
    return fmt.Sprintf(
    	"ZoneVertexM[X=%v, Y=%v, AdditionalProperties=%v]",
    	z.X, z.Y, z.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ZoneVertexM.
// It customizes the JSON marshaling process for ZoneVertexM objects.
func (z ZoneVertexM) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(z.AdditionalProperties,
        "x", "y"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(z.toMap())
}

// toMap converts the ZoneVertexM object to a map representation for JSON marshaling.
func (z ZoneVertexM) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, z.AdditionalProperties)
    structMap["x"] = z.X
    structMap["y"] = z.Y
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ZoneVertexM.
// It customizes the JSON unmarshaling process for ZoneVertexM objects.
func (z *ZoneVertexM) UnmarshalJSON(input []byte) error {
    var temp tempZoneVertexM
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "x", "y")
    if err != nil {
    	return err
    }
    z.AdditionalProperties = additionalProperties
    
    z.X = *temp.X
    z.Y = *temp.Y
    return nil
}

// tempZoneVertexM is a temporary struct used for validating the fields of ZoneVertexM.
type tempZoneVertexM  struct {
    X *float64 `json:"x"`
    Y *float64 `json:"y"`
}

func (z *tempZoneVertexM) validate() error {
    var errs []string
    if z.X == nil {
        errs = append(errs, "required field `x` is missing for type `zone_vertex_m`")
    }
    if z.Y == nil {
        errs = append(errs, "required field `y` is missing for type `zone_vertex_m`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
