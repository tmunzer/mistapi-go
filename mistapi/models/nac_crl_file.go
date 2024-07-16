package models

import (
    "encoding/json"
)

// NacCrlFile represents a NacCrlFile struct.
type NacCrlFile struct {
    // Time at which the file was first uploaded, in epoch
    CreatedTime          *float64       `json:"created_time,omitempty"`
    // ID for file upload, used to identify file even for deletion
    Id                   *string        `json:"id,omitempty"`
    // Time at which the file was last updated, in epoch
    ModifiedTime         *float64       `json:"modified_time,omitempty"`
    // Name for the .crl file issuer if set
    Name                 *string        `json:"name,omitempty"`
    // Download URL for the .crl file
    Url                  *string        `json:"url,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for NacCrlFile.
// It customizes the JSON marshaling process for NacCrlFile objects.
func (n NacCrlFile) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(n.toMap())
}

// toMap converts the NacCrlFile object to a map representation for JSON marshaling.
func (n NacCrlFile) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, n.AdditionalProperties)
    if n.CreatedTime != nil {
        structMap["created_time"] = n.CreatedTime
    }
    if n.Id != nil {
        structMap["id"] = n.Id
    }
    if n.ModifiedTime != nil {
        structMap["modified_time"] = n.ModifiedTime
    }
    if n.Name != nil {
        structMap["name"] = n.Name
    }
    if n.Url != nil {
        structMap["url"] = n.Url
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for NacCrlFile.
// It customizes the JSON unmarshaling process for NacCrlFile objects.
func (n *NacCrlFile) UnmarshalJSON(input []byte) error {
    var temp nacCrlFile
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "created_time", "id", "modified_time", "name", "url")
    if err != nil {
    	return err
    }
    
    n.AdditionalProperties = additionalProperties
    n.CreatedTime = temp.CreatedTime
    n.Id = temp.Id
    n.ModifiedTime = temp.ModifiedTime
    n.Name = temp.Name
    n.Url = temp.Url
    return nil
}

// nacCrlFile is a temporary struct used for validating the fields of NacCrlFile.
type nacCrlFile  struct {
    CreatedTime  *float64 `json:"created_time,omitempty"`
    Id           *string  `json:"id,omitempty"`
    ModifiedTime *float64 `json:"modified_time,omitempty"`
    Name         *string  `json:"name,omitempty"`
    Url          *string  `json:"url,omitempty"`
}
