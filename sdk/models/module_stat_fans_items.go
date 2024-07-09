package models

import (
    "encoding/json"
)

// ModuleStatFansItems represents a ModuleStatFansItems struct.
type ModuleStatFansItems struct {
    Airflow              *string        `json:"airflow,omitempty"`
    Name                 *string        `json:"name,omitempty"`
    Status               *string        `json:"status,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ModuleStatFansItems.
// It customizes the JSON marshaling process for ModuleStatFansItems objects.
func (m ModuleStatFansItems) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the ModuleStatFansItems object to a map representation for JSON marshaling.
func (m ModuleStatFansItems) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
    if m.Airflow != nil {
        structMap["airflow"] = m.Airflow
    }
    if m.Name != nil {
        structMap["name"] = m.Name
    }
    if m.Status != nil {
        structMap["status"] = m.Status
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ModuleStatFansItems.
// It customizes the JSON unmarshaling process for ModuleStatFansItems objects.
func (m *ModuleStatFansItems) UnmarshalJSON(input []byte) error {
    var temp moduleStatFansItems
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "airflow", "name", "status")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.Airflow = temp.Airflow
    m.Name = temp.Name
    m.Status = temp.Status
    return nil
}

// moduleStatFansItems is a temporary struct used for validating the fields of ModuleStatFansItems.
type moduleStatFansItems  struct {
    Airflow *string `json:"airflow,omitempty"`
    Name    *string `json:"name,omitempty"`
    Status  *string `json:"status,omitempty"`
}
