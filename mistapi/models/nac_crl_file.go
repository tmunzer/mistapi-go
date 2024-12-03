package models

import (
    "encoding/json"
)

// NacCrlFile represents a NacCrlFile struct.
type NacCrlFile struct {
    // when the object has been created, in epoch
    CreatedTime          *float64               `json:"created_time,omitempty"`
    // Unique ID for the uploaded CRL file, used to reference the file
    Id                   *string                `json:"id,omitempty"`
    // when the object has been modified for the last time, in epoch
    ModifiedTime         *float64               `json:"modified_time,omitempty"`
    // Issuer name for the CRL file
    Name                 *string                `json:"name,omitempty"`
    // URL to download the uploaded CRL file
    Url                  *string                `json:"url,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for NacCrlFile.
// It customizes the JSON marshaling process for NacCrlFile objects.
func (n NacCrlFile) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(n.AdditionalProperties,
        "created_time", "id", "modified_time", "name", "url"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(n.toMap())
}

// toMap converts the NacCrlFile object to a map representation for JSON marshaling.
func (n NacCrlFile) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, n.AdditionalProperties)
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
    var temp tempNacCrlFile
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "created_time", "id", "modified_time", "name", "url")
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

// tempNacCrlFile is a temporary struct used for validating the fields of NacCrlFile.
type tempNacCrlFile  struct {
    CreatedTime  *float64 `json:"created_time,omitempty"`
    Id           *string  `json:"id,omitempty"`
    ModifiedTime *float64 `json:"modified_time,omitempty"`
    Name         *string  `json:"name,omitempty"`
    Url          *string  `json:"url,omitempty"`
}
