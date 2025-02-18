package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// ResponseSearchItem represents a ResponseSearchItem struct.
type ResponseSearchItem struct {
    // Unique ID of the object instance in the Mist Organnization
    Id                   uuid.UUID              `json:"id"`
    Text                 string                 `json:"text"`
    Type                 string                 `json:"type"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseSearchItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseSearchItem) String() string {
    return fmt.Sprintf(
    	"ResponseSearchItem[Id=%v, Text=%v, Type=%v, AdditionalProperties=%v]",
    	r.Id, r.Text, r.Type, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseSearchItem.
// It customizes the JSON marshaling process for ResponseSearchItem objects.
func (r ResponseSearchItem) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "id", "text", "type"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseSearchItem object to a map representation for JSON marshaling.
func (r ResponseSearchItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["id"] = r.Id
    structMap["text"] = r.Text
    structMap["type"] = r.Type
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseSearchItem.
// It customizes the JSON unmarshaling process for ResponseSearchItem objects.
func (r *ResponseSearchItem) UnmarshalJSON(input []byte) error {
    var temp tempResponseSearchItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "text", "type")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.Id = *temp.Id
    r.Text = *temp.Text
    r.Type = *temp.Type
    return nil
}

// tempResponseSearchItem is a temporary struct used for validating the fields of ResponseSearchItem.
type tempResponseSearchItem  struct {
    Id   *uuid.UUID `json:"id"`
    Text *string    `json:"text"`
    Type *string    `json:"type"`
}

func (r *tempResponseSearchItem) validate() error {
    var errs []string
    if r.Id == nil {
        errs = append(errs, "required field `id` is missing for type `response_search_item`")
    }
    if r.Text == nil {
        errs = append(errs, "required field `text` is missing for type `response_search_item`")
    }
    if r.Type == nil {
        errs = append(errs, "required field `type` is missing for type `response_search_item`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
