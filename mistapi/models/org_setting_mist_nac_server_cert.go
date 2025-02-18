package models

import (
    "encoding/json"
    "fmt"
)

// OrgSettingMistNacServerCert represents a OrgSettingMistNacServerCert struct.
// radius server cert to be presented in EAP TLS
type OrgSettingMistNacServerCert struct {
    Cert                 *string                `json:"cert,omitempty"`
    Key                  *string                `json:"key,omitempty"`
    // private key password (optional)
    Password             *string                `json:"password,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for OrgSettingMistNacServerCert,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OrgSettingMistNacServerCert) String() string {
    return fmt.Sprintf(
    	"OrgSettingMistNacServerCert[Cert=%v, Key=%v, Password=%v, AdditionalProperties=%v]",
    	o.Cert, o.Key, o.Password, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingMistNacServerCert.
// It customizes the JSON marshaling process for OrgSettingMistNacServerCert objects.
func (o OrgSettingMistNacServerCert) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(o.AdditionalProperties,
        "cert", "key", "password"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingMistNacServerCert object to a map representation for JSON marshaling.
func (o OrgSettingMistNacServerCert) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, o.AdditionalProperties)
    if o.Cert != nil {
        structMap["cert"] = o.Cert
    }
    if o.Key != nil {
        structMap["key"] = o.Key
    }
    if o.Password != nil {
        structMap["password"] = o.Password
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingMistNacServerCert.
// It customizes the JSON unmarshaling process for OrgSettingMistNacServerCert objects.
func (o *OrgSettingMistNacServerCert) UnmarshalJSON(input []byte) error {
    var temp tempOrgSettingMistNacServerCert
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "cert", "key", "password")
    if err != nil {
    	return err
    }
    o.AdditionalProperties = additionalProperties
    
    o.Cert = temp.Cert
    o.Key = temp.Key
    o.Password = temp.Password
    return nil
}

// tempOrgSettingMistNacServerCert is a temporary struct used for validating the fields of OrgSettingMistNacServerCert.
type tempOrgSettingMistNacServerCert  struct {
    Cert     *string `json:"cert,omitempty"`
    Key      *string `json:"key,omitempty"`
    Password *string `json:"password,omitempty"`
}
