package models

import (
    "encoding/json"
    "fmt"
)

// RoutingPolicy represents a RoutingPolicy struct.
type RoutingPolicy struct {
    // zero or more criteria/filter can be specified to match the term, all criteria have to be met
    Terms                []RoutingPolicyTerm    `json:"terms,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for RoutingPolicy,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r RoutingPolicy) String() string {
    return fmt.Sprintf(
    	"RoutingPolicy[Terms=%v, AdditionalProperties=%v]",
    	r.Terms, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for RoutingPolicy.
// It customizes the JSON marshaling process for RoutingPolicy objects.
func (r RoutingPolicy) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "terms"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the RoutingPolicy object to a map representation for JSON marshaling.
func (r RoutingPolicy) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    if r.Terms != nil {
        structMap["terms"] = r.Terms
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RoutingPolicy.
// It customizes the JSON unmarshaling process for RoutingPolicy objects.
func (r *RoutingPolicy) UnmarshalJSON(input []byte) error {
    var temp tempRoutingPolicy
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "terms")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.Terms = temp.Terms
    return nil
}

// tempRoutingPolicy is a temporary struct used for validating the fields of RoutingPolicy.
type tempRoutingPolicy  struct {
    Terms []RoutingPolicyTerm `json:"terms,omitempty"`
}
