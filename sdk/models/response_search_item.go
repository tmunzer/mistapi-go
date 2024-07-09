package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// ResponseSearchItem represents a ResponseSearchItem struct.
type ResponseSearchItem struct {
    Id                   uuid.UUID      `json:"id"`
    Text                 string         `json:"text"`
    Type                 string         `json:"type"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseSearchItem.
// It customizes the JSON marshaling process for ResponseSearchItem objects.
func (r ResponseSearchItem) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseSearchItem object to a map representation for JSON marshaling.
func (r ResponseSearchItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["id"] = r.Id
    structMap["text"] = r.Text
    structMap["type"] = r.Type
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseSearchItem.
// It customizes the JSON unmarshaling process for ResponseSearchItem objects.
func (r *ResponseSearchItem) UnmarshalJSON(input []byte) error {
    var temp responseSearchItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "id", "text", "type")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.Id = *temp.Id
    r.Text = *temp.Text
    r.Type = *temp.Type
    return nil
}

// responseSearchItem is a temporary struct used for validating the fields of ResponseSearchItem.
type responseSearchItem  struct {
    Id   *uuid.UUID `json:"id"`
    Text *string    `json:"text"`
    Type *string    `json:"type"`
}

func (r *responseSearchItem) validate() error {
    var errs []string
    if r.Id == nil {
        errs = append(errs, "required field `id` is missing for type `Response_Search_Item`")
    }
    if r.Text == nil {
        errs = append(errs, "required field `text` is missing for type `Response_Search_Item`")
    }
    if r.Type == nil {
        errs = append(errs, "required field `type` is missing for type `Response_Search_Item`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
