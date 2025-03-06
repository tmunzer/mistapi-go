package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// Fingerprint represents a Fingerprint struct.
type Fingerprint struct {
    Family               *string                `json:"family,omitempty"`
    Mac                  *string                `json:"mac,omitempty"`
    Mfg                  *string                `json:"mfg,omitempty"`
    Model                *string                `json:"model,omitempty"`
    OrgId                *uuid.UUID             `json:"org_id,omitempty"`
    Os                   *string                `json:"os,omitempty"`
    OsType               *string                `json:"os_type,omitempty"`
    RandomMac            *bool                  `json:"random_mac,omitempty"`
    SiteId               *uuid.UUID             `json:"site_id,omitempty"`
    // Epoch (seconds)
    Timestamp            *float64               `json:"timestamp,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for Fingerprint,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (f Fingerprint) String() string {
    return fmt.Sprintf(
    	"Fingerprint[Family=%v, Mac=%v, Mfg=%v, Model=%v, OrgId=%v, Os=%v, OsType=%v, RandomMac=%v, SiteId=%v, Timestamp=%v, AdditionalProperties=%v]",
    	f.Family, f.Mac, f.Mfg, f.Model, f.OrgId, f.Os, f.OsType, f.RandomMac, f.SiteId, f.Timestamp, f.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Fingerprint.
// It customizes the JSON marshaling process for Fingerprint objects.
func (f Fingerprint) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(f.AdditionalProperties,
        "family", "mac", "mfg", "model", "org_id", "os", "os_type", "random_mac", "site_id", "timestamp"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(f.toMap())
}

// toMap converts the Fingerprint object to a map representation for JSON marshaling.
func (f Fingerprint) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, f.AdditionalProperties)
    if f.Family != nil {
        structMap["family"] = f.Family
    }
    if f.Mac != nil {
        structMap["mac"] = f.Mac
    }
    if f.Mfg != nil {
        structMap["mfg"] = f.Mfg
    }
    if f.Model != nil {
        structMap["model"] = f.Model
    }
    if f.OrgId != nil {
        structMap["org_id"] = f.OrgId
    }
    if f.Os != nil {
        structMap["os"] = f.Os
    }
    if f.OsType != nil {
        structMap["os_type"] = f.OsType
    }
    if f.RandomMac != nil {
        structMap["random_mac"] = f.RandomMac
    }
    if f.SiteId != nil {
        structMap["site_id"] = f.SiteId
    }
    if f.Timestamp != nil {
        structMap["timestamp"] = f.Timestamp
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Fingerprint.
// It customizes the JSON unmarshaling process for Fingerprint objects.
func (f *Fingerprint) UnmarshalJSON(input []byte) error {
    var temp tempFingerprint
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "family", "mac", "mfg", "model", "org_id", "os", "os_type", "random_mac", "site_id", "timestamp")
    if err != nil {
    	return err
    }
    f.AdditionalProperties = additionalProperties
    
    f.Family = temp.Family
    f.Mac = temp.Mac
    f.Mfg = temp.Mfg
    f.Model = temp.Model
    f.OrgId = temp.OrgId
    f.Os = temp.Os
    f.OsType = temp.OsType
    f.RandomMac = temp.RandomMac
    f.SiteId = temp.SiteId
    f.Timestamp = temp.Timestamp
    return nil
}

// tempFingerprint is a temporary struct used for validating the fields of Fingerprint.
type tempFingerprint  struct {
    Family    *string    `json:"family,omitempty"`
    Mac       *string    `json:"mac,omitempty"`
    Mfg       *string    `json:"mfg,omitempty"`
    Model     *string    `json:"model,omitempty"`
    OrgId     *uuid.UUID `json:"org_id,omitempty"`
    Os        *string    `json:"os,omitempty"`
    OsType    *string    `json:"os_type,omitempty"`
    RandomMac *bool      `json:"random_mac,omitempty"`
    SiteId    *uuid.UUID `json:"site_id,omitempty"`
    Timestamp *float64   `json:"timestamp,omitempty"`
}
