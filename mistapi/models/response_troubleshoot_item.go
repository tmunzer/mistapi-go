package models

import (
    "encoding/json"
)

// ResponseTroubleshootItem represents a ResponseTroubleshootItem struct.
type ResponseTroubleshootItem struct {
    Category             *string                `json:"category,omitempty"`
    Reason               *string                `json:"reason,omitempty"`
    Recommendation       *string                `json:"recommendation,omitempty"`
    Text                 *string                `json:"text,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseTroubleshootItem.
// It customizes the JSON marshaling process for ResponseTroubleshootItem objects.
func (r ResponseTroubleshootItem) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "category", "reason", "recommendation", "text"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseTroubleshootItem object to a map representation for JSON marshaling.
func (r ResponseTroubleshootItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    if r.Category != nil {
        structMap["category"] = r.Category
    }
    if r.Reason != nil {
        structMap["reason"] = r.Reason
    }
    if r.Recommendation != nil {
        structMap["recommendation"] = r.Recommendation
    }
    if r.Text != nil {
        structMap["text"] = r.Text
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseTroubleshootItem.
// It customizes the JSON unmarshaling process for ResponseTroubleshootItem objects.
func (r *ResponseTroubleshootItem) UnmarshalJSON(input []byte) error {
    var temp tempResponseTroubleshootItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "category", "reason", "recommendation", "text")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.Category = temp.Category
    r.Reason = temp.Reason
    r.Recommendation = temp.Recommendation
    r.Text = temp.Text
    return nil
}

// tempResponseTroubleshootItem is a temporary struct used for validating the fields of ResponseTroubleshootItem.
type tempResponseTroubleshootItem  struct {
    Category       *string `json:"category,omitempty"`
    Reason         *string `json:"reason,omitempty"`
    Recommendation *string `json:"recommendation,omitempty"`
    Text           *string `json:"text,omitempty"`
}
