// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// RogueDetails represents a RogueDetails struct.
type RogueDetails struct {
    Manufacture          string                 `json:"manufacture"`
    SeenAsClient         bool                   `json:"seen_as_client"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for RogueDetails,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r RogueDetails) String() string {
    return fmt.Sprintf(
    	"RogueDetails[Manufacture=%v, SeenAsClient=%v, AdditionalProperties=%v]",
    	r.Manufacture, r.SeenAsClient, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for RogueDetails.
// It customizes the JSON marshaling process for RogueDetails objects.
func (r RogueDetails) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "manufacture", "seen_as_client"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the RogueDetails object to a map representation for JSON marshaling.
func (r RogueDetails) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["manufacture"] = r.Manufacture
    structMap["seen_as_client"] = r.SeenAsClient
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RogueDetails.
// It customizes the JSON unmarshaling process for RogueDetails objects.
func (r *RogueDetails) UnmarshalJSON(input []byte) error {
    var temp tempRogueDetails
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "manufacture", "seen_as_client")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.Manufacture = *temp.Manufacture
    r.SeenAsClient = *temp.SeenAsClient
    return nil
}

// tempRogueDetails is a temporary struct used for validating the fields of RogueDetails.
type tempRogueDetails  struct {
    Manufacture  *string `json:"manufacture"`
    SeenAsClient *bool   `json:"seen_as_client"`
}

func (r *tempRogueDetails) validate() error {
    var errs []string
    if r.Manufacture == nil {
        errs = append(errs, "required field `manufacture` is missing for type `rogue_details`")
    }
    if r.SeenAsClient == nil {
        errs = append(errs, "required field `seen_as_client` is missing for type `rogue_details`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
