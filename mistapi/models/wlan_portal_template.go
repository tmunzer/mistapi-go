package models

import (
    "encoding/json"
)

// WlanPortalTemplate represents a WlanPortalTemplate struct.
type WlanPortalTemplate struct {
    // portal template wlan settings
    PortalTemplate       *WlanPortalTemplateSetting `json:"portal_template,omitempty"`
    AdditionalProperties map[string]any             `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WlanPortalTemplate.
// It customizes the JSON marshaling process for WlanPortalTemplate objects.
func (w WlanPortalTemplate) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WlanPortalTemplate object to a map representation for JSON marshaling.
func (w WlanPortalTemplate) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
    if w.PortalTemplate != nil {
        structMap["portal_template"] = w.PortalTemplate.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WlanPortalTemplate.
// It customizes the JSON unmarshaling process for WlanPortalTemplate objects.
func (w *WlanPortalTemplate) UnmarshalJSON(input []byte) error {
    var temp wlanPortalTemplate
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "portal_template")
    if err != nil {
    	return err
    }
    
    w.AdditionalProperties = additionalProperties
    w.PortalTemplate = temp.PortalTemplate
    return nil
}

// wlanPortalTemplate is a temporary struct used for validating the fields of WlanPortalTemplate.
type wlanPortalTemplate  struct {
    PortalTemplate *WlanPortalTemplateSetting `json:"portal_template,omitempty"`
}
