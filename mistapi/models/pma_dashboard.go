package models

import (
	"encoding/json"
)

// PmaDashboard represents a PmaDashboard struct.
type PmaDashboard struct {
	// description of the dashboard
	Description *string `json:"description,omitempty"`
	// group label name
	Label *string `json:"label,omitempty"`
	// name of the dashboard
	Name *string `json:"name,omitempty"`
	// url to access dashboard. Url will redirect the user to the dashboard
	Url                  *string        `json:"url,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for PmaDashboard.
// It customizes the JSON marshaling process for PmaDashboard objects.
func (p PmaDashboard) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(p.toMap())
}

// toMap converts the PmaDashboard object to a map representation for JSON marshaling.
func (p PmaDashboard) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, p.AdditionalProperties)
	if p.Description != nil {
		structMap["description"] = p.Description
	}
	if p.Label != nil {
		structMap["label"] = p.Label
	}
	if p.Name != nil {
		structMap["name"] = p.Name
	}
	if p.Url != nil {
		structMap["url"] = p.Url
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PmaDashboard.
// It customizes the JSON unmarshaling process for PmaDashboard objects.
func (p *PmaDashboard) UnmarshalJSON(input []byte) error {
	var temp tempPmaDashboard
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "description", "label", "name", "url")
	if err != nil {
		return err
	}

	p.AdditionalProperties = additionalProperties
	p.Description = temp.Description
	p.Label = temp.Label
	p.Name = temp.Name
	p.Url = temp.Url
	return nil
}

// tempPmaDashboard is a temporary struct used for validating the fields of PmaDashboard.
type tempPmaDashboard struct {
	Description *string `json:"description,omitempty"`
	Label       *string `json:"label,omitempty"`
	Name        *string `json:"name,omitempty"`
	Url         *string `json:"url,omitempty"`
}
