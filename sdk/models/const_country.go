package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ConstCountry represents a ConstCountry struct.
type ConstCountry struct {
    // country code, in two-character
    Alpha2               string         `json:"alpha2"`
    Certified            bool           `json:"certified"`
    Name                 string         `json:"name"`
    // country code, ISO 3166-1 numeric
    Numeric              float64        `json:"numeric"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ConstCountry.
// It customizes the JSON marshaling process for ConstCountry objects.
func (c ConstCountry) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the ConstCountry object to a map representation for JSON marshaling.
func (c ConstCountry) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["alpha2"] = c.Alpha2
    structMap["certified"] = c.Certified
    structMap["name"] = c.Name
    structMap["numeric"] = c.Numeric
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConstCountry.
// It customizes the JSON unmarshaling process for ConstCountry objects.
func (c *ConstCountry) UnmarshalJSON(input []byte) error {
    var temp constCountry
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "alpha2", "certified", "name", "numeric")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.Alpha2 = *temp.Alpha2
    c.Certified = *temp.Certified
    c.Name = *temp.Name
    c.Numeric = *temp.Numeric
    return nil
}

// constCountry is a temporary struct used for validating the fields of ConstCountry.
type constCountry  struct {
    Alpha2    *string  `json:"alpha2"`
    Certified *bool    `json:"certified"`
    Name      *string  `json:"name"`
    Numeric   *float64 `json:"numeric"`
}

func (c *constCountry) validate() error {
    var errs []string
    if c.Alpha2 == nil {
        errs = append(errs, "required field `alpha2` is missing for type `Const_Country`")
    }
    if c.Certified == nil {
        errs = append(errs, "required field `certified` is missing for type `Const_Country`")
    }
    if c.Name == nil {
        errs = append(errs, "required field `name` is missing for type `Const_Country`")
    }
    if c.Numeric == nil {
        errs = append(errs, "required field `numeric` is missing for type `Const_Country`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
