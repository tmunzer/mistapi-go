package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// ResponseClaimMxEdge represents a ResponseClaimMxEdge struct.
type ResponseClaimMxEdge struct {
    // Unique ID of the object instance in the Mist Organnization
    Id                   uuid.UUID              `json:"id"`
    Magic                string                 `json:"magic"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseClaimMxEdge.
// It customizes the JSON marshaling process for ResponseClaimMxEdge objects.
func (r ResponseClaimMxEdge) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "id", "magic"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseClaimMxEdge object to a map representation for JSON marshaling.
func (r ResponseClaimMxEdge) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["id"] = r.Id
    structMap["magic"] = r.Magic
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseClaimMxEdge.
// It customizes the JSON unmarshaling process for ResponseClaimMxEdge objects.
func (r *ResponseClaimMxEdge) UnmarshalJSON(input []byte) error {
    var temp tempResponseClaimMxEdge
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "magic")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.Id = *temp.Id
    r.Magic = *temp.Magic
    return nil
}

// tempResponseClaimMxEdge is a temporary struct used for validating the fields of ResponseClaimMxEdge.
type tempResponseClaimMxEdge  struct {
    Id    *uuid.UUID `json:"id"`
    Magic *string    `json:"magic"`
}

func (r *tempResponseClaimMxEdge) validate() error {
    var errs []string
    if r.Id == nil {
        errs = append(errs, "required field `id` is missing for type `response_claim_mx_edge`")
    }
    if r.Magic == nil {
        errs = append(errs, "required field `magic` is missing for type `response_claim_mx_edge`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
