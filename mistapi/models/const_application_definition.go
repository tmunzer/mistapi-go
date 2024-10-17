package models

import (
	"encoding/json"
)

// ConstApplicationDefinition represents a ConstApplicationDefinition struct.
type ConstApplicationDefinition struct {
	AppId                *bool          `json:"app_id,omitempty"`
	AppImageUrl          *string        `json:"app_image_url,omitempty"`
	AppProbe             *bool          `json:"app_probe,omitempty"`
	Category             *string        `json:"category,omitempty"`
	Group                *string        `json:"group,omitempty"`
	Key                  *string        `json:"key,omitempty"`
	Name                 *string        `json:"name,omitempty"`
	SignatureBased       *bool          `json:"signature_based,omitempty"`
	SsrAppId             *bool          `json:"ssr_app_id,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ConstApplicationDefinition.
// It customizes the JSON marshaling process for ConstApplicationDefinition objects.
func (c ConstApplicationDefinition) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the ConstApplicationDefinition object to a map representation for JSON marshaling.
func (c ConstApplicationDefinition) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, c.AdditionalProperties)
	if c.AppId != nil {
		structMap["app_id"] = c.AppId
	}
	if c.AppImageUrl != nil {
		structMap["app_image_url"] = c.AppImageUrl
	}
	if c.AppProbe != nil {
		structMap["app_probe"] = c.AppProbe
	}
	if c.Category != nil {
		structMap["category"] = c.Category
	}
	if c.Group != nil {
		structMap["group"] = c.Group
	}
	if c.Key != nil {
		structMap["key"] = c.Key
	}
	if c.Name != nil {
		structMap["name"] = c.Name
	}
	if c.SignatureBased != nil {
		structMap["signature_based"] = c.SignatureBased
	}
	if c.SsrAppId != nil {
		structMap["ssr_app_id"] = c.SsrAppId
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConstApplicationDefinition.
// It customizes the JSON unmarshaling process for ConstApplicationDefinition objects.
func (c *ConstApplicationDefinition) UnmarshalJSON(input []byte) error {
	var temp tempConstApplicationDefinition
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "app_id", "app_image_url", "app_probe", "category", "group", "key", "name", "signature_based", "ssr_app_id")
	if err != nil {
		return err
	}

	c.AdditionalProperties = additionalProperties
	c.AppId = temp.AppId
	c.AppImageUrl = temp.AppImageUrl
	c.AppProbe = temp.AppProbe
	c.Category = temp.Category
	c.Group = temp.Group
	c.Key = temp.Key
	c.Name = temp.Name
	c.SignatureBased = temp.SignatureBased
	c.SsrAppId = temp.SsrAppId
	return nil
}

// tempConstApplicationDefinition is a temporary struct used for validating the fields of ConstApplicationDefinition.
type tempConstApplicationDefinition struct {
	AppId          *bool   `json:"app_id,omitempty"`
	AppImageUrl    *string `json:"app_image_url,omitempty"`
	AppProbe       *bool   `json:"app_probe,omitempty"`
	Category       *string `json:"category,omitempty"`
	Group          *string `json:"group,omitempty"`
	Key            *string `json:"key,omitempty"`
	Name           *string `json:"name,omitempty"`
	SignatureBased *bool   `json:"signature_based,omitempty"`
	SsrAppId       *bool   `json:"ssr_app_id,omitempty"`
}
