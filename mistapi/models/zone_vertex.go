package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ZoneVertex represents a ZoneVertex struct.
type ZoneVertex struct {
    // x in pixel
    X                    float64                `json:"x"`
    // y in pixel
    Y                    float64                `json:"y"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ZoneVertex,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (z ZoneVertex) String() string {
    return fmt.Sprintf(
    	"ZoneVertex[X=%v, Y=%v, AdditionalProperties=%v]",
    	z.X, z.Y, z.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ZoneVertex.
// It customizes the JSON marshaling process for ZoneVertex objects.
func (z ZoneVertex) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(z.AdditionalProperties,
        "x", "y"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(z.toMap())
}

// toMap converts the ZoneVertex object to a map representation for JSON marshaling.
func (z ZoneVertex) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, z.AdditionalProperties)
    structMap["x"] = z.X
    structMap["y"] = z.Y
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ZoneVertex.
// It customizes the JSON unmarshaling process for ZoneVertex objects.
func (z *ZoneVertex) UnmarshalJSON(input []byte) error {
    var temp tempZoneVertex
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

// tempZoneVertex is a temporary struct used for validating the fields of ZoneVertex.
type tempZoneVertex  struct {
    X *float64 `json:"x"`
    Y *float64 `json:"y"`
}

func (z *tempZoneVertex) validate() error {
    var errs []string
    if z.X == nil {
        errs = append(errs, "required field `x` is missing for type `zone_vertex`")
    }
    if z.Y == nil {
        errs = append(errs, "required field `y` is missing for type `zone_vertex`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
