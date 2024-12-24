package models

import (
    "encoding/json"
    "fmt"
)

// ApIot represents a ApIot struct.
// IoT AP settings
type ApIot struct {
    // IoT output AP settings
    A1                   *ApIotOutput           `json:"A1,omitempty"`
    // IoT output AP settings
    A2                   *ApIotOutput           `json:"A2,omitempty"`
    // IoT output AP settings
    A3                   *ApIotOutput           `json:"A3,omitempty"`
    // IoT output AP settings
    A4                   *ApIotOutput           `json:"A4,omitempty"`
    // IoT Input AP settings
    DI1                  *ApIotInput            `json:"DI1,omitempty"`
    // IoT Input AP settings
    DI2                  *ApIotInput            `json:"DI2,omitempty"`
    // IoT output AP settings
    DO                   *ApIotOutput           `json:"DO,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ApIot,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a ApIot) String() string {
    return fmt.Sprintf(
    	"ApIot[A1=%v, A2=%v, A3=%v, A4=%v, DI1=%v, DI2=%v, DO=%v, AdditionalProperties=%v]",
    	a.A1, a.A2, a.A3, a.A4, a.DI1, a.DI2, a.DO, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ApIot.
// It customizes the JSON marshaling process for ApIot objects.
func (a ApIot) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "A1", "A2", "A3", "A4", "DI1", "DI2", "DO"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the ApIot object to a map representation for JSON marshaling.
func (a ApIot) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    if a.A1 != nil {
        structMap["A1"] = a.A1.toMap()
    }
    if a.A2 != nil {
        structMap["A2"] = a.A2.toMap()
    }
    if a.A3 != nil {
        structMap["A3"] = a.A3.toMap()
    }
    if a.A4 != nil {
        structMap["A4"] = a.A4.toMap()
    }
    if a.DI1 != nil {
        structMap["DI1"] = a.DI1.toMap()
    }
    if a.DI2 != nil {
        structMap["DI2"] = a.DI2.toMap()
    }
    if a.DO != nil {
        structMap["DO"] = a.DO.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApIot.
// It customizes the JSON unmarshaling process for ApIot objects.
func (a *ApIot) UnmarshalJSON(input []byte) error {
    var temp tempApIot
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "A1", "A2", "A3", "A4", "DI1", "DI2", "DO")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.A1 = temp.A1
    a.A2 = temp.A2
    a.A3 = temp.A3
    a.A4 = temp.A4
    a.DI1 = temp.DI1
    a.DI2 = temp.DI2
    a.DO = temp.DO
    return nil
}

// tempApIot is a temporary struct used for validating the fields of ApIot.
type tempApIot  struct {
    A1  *ApIotOutput `json:"A1,omitempty"`
    A2  *ApIotOutput `json:"A2,omitempty"`
    A3  *ApIotOutput `json:"A3,omitempty"`
    A4  *ApIotOutput `json:"A4,omitempty"`
    DI1 *ApIotInput  `json:"DI1,omitempty"`
    DI2 *ApIotInput  `json:"DI2,omitempty"`
    DO  *ApIotOutput `json:"DO,omitempty"`
}
