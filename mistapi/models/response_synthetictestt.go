package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// ResponseSynthetictest represents a ResponseSynthetictest struct.
type ResponseSynthetictest struct {
    // Unique ID of the object instance in the Mist Organization
    Id                   *uuid.UUID             `json:"id,omitempty"`
    Message              *string                `json:"message,omitempty"`
    Status               *string                `json:"status,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseSynthetictest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseSynthetictest) String() string {
    return fmt.Sprintf(
    	"ResponseSynthetictest[Id=%v, Message=%v, Status=%v, AdditionalProperties=%v]",
    	r.Id, r.Message, r.Status, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseSynthetictest.
// It customizes the JSON marshaling process for ResponseSynthetictest objects.
func (r ResponseSynthetictest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "id", "message", "status"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseSynthetictest object to a map representation for JSON marshaling.
func (r ResponseSynthetictest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    if r.Id != nil {
        structMap["id"] = r.Id
    }
    if r.Message != nil {
        structMap["message"] = r.Message
    }
    if r.Status != nil {
        structMap["status"] = r.Status
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseSynthetictest.
// It customizes the JSON unmarshaling process for ResponseSynthetictest objects.
func (r *ResponseSynthetictest) UnmarshalJSON(input []byte) error {
    var temp tempResponseSynthetictest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "message", "status")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.Id = temp.Id
    r.Message = temp.Message
    r.Status = temp.Status
    return nil
}

// tempResponseSynthetictest is a temporary struct used for validating the fields of ResponseSynthetictest.
type tempResponseSynthetictest  struct {
    Id      *uuid.UUID `json:"id,omitempty"`
    Message *string    `json:"message,omitempty"`
    Status  *string    `json:"status,omitempty"`
}
