package models

import (
	"encoding/json"
)

// UserMacsSearch represents a UserMacsSearch struct.
type UserMacsSearch struct {
	Limit                *int           `json:"limit,omitempty"`
	Page                 *int           `json:"page,omitempty"`
	Results              []UserMac      `json:"results,omitempty"`
	Total                *int           `json:"total,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UserMacsSearch.
// It customizes the JSON marshaling process for UserMacsSearch objects.
func (u UserMacsSearch) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UserMacsSearch object to a map representation for JSON marshaling.
func (u UserMacsSearch) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, u.AdditionalProperties)
	if u.Limit != nil {
		structMap["limit"] = u.Limit
	}
	if u.Page != nil {
		structMap["page"] = u.Page
	}
	if u.Results != nil {
		structMap["results"] = u.Results
	}
	if u.Total != nil {
		structMap["total"] = u.Total
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UserMacsSearch.
// It customizes the JSON unmarshaling process for UserMacsSearch objects.
func (u *UserMacsSearch) UnmarshalJSON(input []byte) error {
	var temp tempUserMacsSearch
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "limit", "page", "results", "total")
	if err != nil {
		return err
	}

	u.AdditionalProperties = additionalProperties
	u.Limit = temp.Limit
	u.Page = temp.Page
	u.Results = temp.Results
	u.Total = temp.Total
	return nil
}

// tempUserMacsSearch is a temporary struct used for validating the fields of UserMacsSearch.
type tempUserMacsSearch struct {
	Limit   *int      `json:"limit,omitempty"`
	Page    *int      `json:"page,omitempty"`
	Results []UserMac `json:"results,omitempty"`
	Total   *int      `json:"total,omitempty"`
}
