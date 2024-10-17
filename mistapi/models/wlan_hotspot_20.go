package models

import (
	"encoding/json"
)

// WlanHotspot20 represents a WlanHotspot20 struct.
// hostspot 2.0 wlan settings
type WlanHotspot20 struct {
	DomainName []string `json:"domain_name,omitempty"`
	// whether to enable hotspot 2.0 config
	Enabled   *bool    `json:"enabled,omitempty"`
	NaiRealms []string `json:"nai_realms,omitempty"`
	// list of operators to support
	Operators []WlanHotspot20OperatorsItemEnum `json:"operators,omitempty"`
	Rcoi      []string                         `json:"rcoi,omitempty"`
	// venue name, default is site name
	VenueName            *string        `json:"venue_name,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WlanHotspot20.
// It customizes the JSON marshaling process for WlanHotspot20 objects.
func (w WlanHotspot20) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(w.toMap())
}

// toMap converts the WlanHotspot20 object to a map representation for JSON marshaling.
func (w WlanHotspot20) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, w.AdditionalProperties)
	if w.DomainName != nil {
		structMap["domain_name"] = w.DomainName
	}
	if w.Enabled != nil {
		structMap["enabled"] = w.Enabled
	}
	if w.NaiRealms != nil {
		structMap["nai_realms"] = w.NaiRealms
	}
	if w.Operators != nil {
		structMap["operators"] = w.Operators
	}
	if w.Rcoi != nil {
		structMap["rcoi"] = w.Rcoi
	}
	if w.VenueName != nil {
		structMap["venue_name"] = w.VenueName
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WlanHotspot20.
// It customizes the JSON unmarshaling process for WlanHotspot20 objects.
func (w *WlanHotspot20) UnmarshalJSON(input []byte) error {
	var temp tempWlanHotspot20
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "domain_name", "enabled", "nai_realms", "operators", "rcoi", "venue_name")
	if err != nil {
		return err
	}

	w.AdditionalProperties = additionalProperties
	w.DomainName = temp.DomainName
	w.Enabled = temp.Enabled
	w.NaiRealms = temp.NaiRealms
	w.Operators = temp.Operators
	w.Rcoi = temp.Rcoi
	w.VenueName = temp.VenueName
	return nil
}

// tempWlanHotspot20 is a temporary struct used for validating the fields of WlanHotspot20.
type tempWlanHotspot20 struct {
	DomainName []string                         `json:"domain_name,omitempty"`
	Enabled    *bool                            `json:"enabled,omitempty"`
	NaiRealms  []string                         `json:"nai_realms,omitempty"`
	Operators  []WlanHotspot20OperatorsItemEnum `json:"operators,omitempty"`
	Rcoi       []string                         `json:"rcoi,omitempty"`
	VenueName  *string                          `json:"venue_name,omitempty"`
}
