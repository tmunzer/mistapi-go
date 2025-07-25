// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// AutoPlacement represents a AutoPlacement struct.
type AutoPlacement struct {
    // Set to `true` to perform an invalid AP check and provide an estimated run time without enqueuing the run into the auto placement service.
    Dryrun               *bool                  `json:"dryrun,omitempty"`
    // * If `force_collection`==`false`: the API attempts to start localization with existing data.
    // * If `force_collection`==`true`: maintenance the API attempts to start orchestration.
    ForceCollection      *bool                  `json:"force_collection,omitempty"`
    // List of device macs
    Macs                 []string               `json:"macs,omitempty"`
    // Set to `true` to run auto placement even if there are invalid APs in the selected APs.
    Override             *bool                  `json:"override,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for AutoPlacement,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AutoPlacement) String() string {
    return fmt.Sprintf(
    	"AutoPlacement[Dryrun=%v, ForceCollection=%v, Macs=%v, Override=%v, AdditionalProperties=%v]",
    	a.Dryrun, a.ForceCollection, a.Macs, a.Override, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AutoPlacement.
// It customizes the JSON marshaling process for AutoPlacement objects.
func (a AutoPlacement) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "dryrun", "force_collection", "macs", "override"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the AutoPlacement object to a map representation for JSON marshaling.
func (a AutoPlacement) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    if a.Dryrun != nil {
        structMap["dryrun"] = a.Dryrun
    }
    if a.ForceCollection != nil {
        structMap["force_collection"] = a.ForceCollection
    }
    if a.Macs != nil {
        structMap["macs"] = a.Macs
    }
    if a.Override != nil {
        structMap["override"] = a.Override
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AutoPlacement.
// It customizes the JSON unmarshaling process for AutoPlacement objects.
func (a *AutoPlacement) UnmarshalJSON(input []byte) error {
    var temp tempAutoPlacement
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "dryrun", "force_collection", "macs", "override")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.Dryrun = temp.Dryrun
    a.ForceCollection = temp.ForceCollection
    a.Macs = temp.Macs
    a.Override = temp.Override
    return nil
}

// tempAutoPlacement is a temporary struct used for validating the fields of AutoPlacement.
type tempAutoPlacement  struct {
    Dryrun          *bool    `json:"dryrun,omitempty"`
    ForceCollection *bool    `json:"force_collection,omitempty"`
    Macs            []string `json:"macs,omitempty"`
    Override        *bool    `json:"override,omitempty"`
}
