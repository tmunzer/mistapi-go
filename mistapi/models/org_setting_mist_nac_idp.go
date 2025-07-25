// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// OrgSettingMistNacIdp represents a OrgSettingMistNacIdp struct.
type OrgSettingMistNacIdp struct {
    // When the IDP of mxedge_proxy type, exclude the following realms from proxying in addition to other valid home realms in this org
    ExcludeRealms        []string               `json:"exclude_realms,omitempty"`
    // Unique ID of the object instance in the Mist Organization
    Id                   *uuid.UUID             `json:"id,omitempty"`
    // Which realm should trigger this IDP. User Realm is extracted from:
    // * Username-AVP (`mist.com` from john@mist.com)
    // * Cert CN
    UserRealms           []string               `json:"user_realms,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for OrgSettingMistNacIdp,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OrgSettingMistNacIdp) String() string {
    return fmt.Sprintf(
    	"OrgSettingMistNacIdp[ExcludeRealms=%v, Id=%v, UserRealms=%v, AdditionalProperties=%v]",
    	o.ExcludeRealms, o.Id, o.UserRealms, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingMistNacIdp.
// It customizes the JSON marshaling process for OrgSettingMistNacIdp objects.
func (o OrgSettingMistNacIdp) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(o.AdditionalProperties,
        "exclude_realms", "id", "user_realms"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingMistNacIdp object to a map representation for JSON marshaling.
func (o OrgSettingMistNacIdp) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, o.AdditionalProperties)
    if o.ExcludeRealms != nil {
        structMap["exclude_realms"] = o.ExcludeRealms
    }
    if o.Id != nil {
        structMap["id"] = o.Id
    }
    if o.UserRealms != nil {
        structMap["user_realms"] = o.UserRealms
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingMistNacIdp.
// It customizes the JSON unmarshaling process for OrgSettingMistNacIdp objects.
func (o *OrgSettingMistNacIdp) UnmarshalJSON(input []byte) error {
    var temp tempOrgSettingMistNacIdp
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "exclude_realms", "id", "user_realms")
    if err != nil {
    	return err
    }
    o.AdditionalProperties = additionalProperties
    
    o.ExcludeRealms = temp.ExcludeRealms
    o.Id = temp.Id
    o.UserRealms = temp.UserRealms
    return nil
}

// tempOrgSettingMistNacIdp is a temporary struct used for validating the fields of OrgSettingMistNacIdp.
type tempOrgSettingMistNacIdp  struct {
    ExcludeRealms []string   `json:"exclude_realms,omitempty"`
    Id            *uuid.UUID `json:"id,omitempty"`
    UserRealms    []string   `json:"user_realms,omitempty"`
}
