package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// Msp represents a Msp struct.
type Msp struct {
    AllowMist            *bool                  `json:"allow_mist,omitempty"`
    // When the object has been created, in epoch
    CreatedTime          *float64               `json:"created_time,omitempty"`
    // Unique ID of the object instance in the Mist Organization
    Id                   *uuid.UUID             `json:"id,omitempty"`
    // For advanced tier (uMSPs) only
    LogoUrl              *string                `json:"logo_url,omitempty"`
    // When the object has been modified for the last time, in epoch
    ModifiedTime         *float64               `json:"modified_time,omitempty"`
    Name                 *string                `json:"name,omitempty"`
    // enum: `advanced`, `base`
    Tier                 *MspTierEnum           `json:"tier,omitempty"`
    // For advanced tier (uMSPs) only
    Url                  *string                `json:"url,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for Msp,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m Msp) String() string {
    return fmt.Sprintf(
    	"Msp[AllowMist=%v, CreatedTime=%v, Id=%v, LogoUrl=%v, ModifiedTime=%v, Name=%v, Tier=%v, Url=%v, AdditionalProperties=%v]",
    	m.AllowMist, m.CreatedTime, m.Id, m.LogoUrl, m.ModifiedTime, m.Name, m.Tier, m.Url, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Msp.
// It customizes the JSON marshaling process for Msp objects.
func (m Msp) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "allow_mist", "created_time", "id", "logo_url", "modified_time", "name", "tier", "url"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the Msp object to a map representation for JSON marshaling.
func (m Msp) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
    if m.AllowMist != nil {
        structMap["allow_mist"] = m.AllowMist
    }
    if m.CreatedTime != nil {
        structMap["created_time"] = m.CreatedTime
    }
    if m.Id != nil {
        structMap["id"] = m.Id
    }
    if m.LogoUrl != nil {
        structMap["logo_url"] = m.LogoUrl
    }
    if m.ModifiedTime != nil {
        structMap["modified_time"] = m.ModifiedTime
    }
    if m.Name != nil {
        structMap["name"] = m.Name
    }
    if m.Tier != nil {
        structMap["tier"] = m.Tier
    }
    if m.Url != nil {
        structMap["url"] = m.Url
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Msp.
// It customizes the JSON unmarshaling process for Msp objects.
func (m *Msp) UnmarshalJSON(input []byte) error {
    var temp tempMsp
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "allow_mist", "created_time", "id", "logo_url", "modified_time", "name", "tier", "url")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.AllowMist = temp.AllowMist
    m.CreatedTime = temp.CreatedTime
    m.Id = temp.Id
    m.LogoUrl = temp.LogoUrl
    m.ModifiedTime = temp.ModifiedTime
    m.Name = temp.Name
    m.Tier = temp.Tier
    m.Url = temp.Url
    return nil
}

// tempMsp is a temporary struct used for validating the fields of Msp.
type tempMsp  struct {
    AllowMist    *bool        `json:"allow_mist,omitempty"`
    CreatedTime  *float64     `json:"created_time,omitempty"`
    Id           *uuid.UUID   `json:"id,omitempty"`
    LogoUrl      *string      `json:"logo_url,omitempty"`
    ModifiedTime *float64     `json:"modified_time,omitempty"`
    Name         *string      `json:"name,omitempty"`
    Tier         *MspTierEnum `json:"tier,omitempty"`
    Url          *string      `json:"url,omitempty"`
}
