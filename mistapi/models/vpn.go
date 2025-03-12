package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// Vpn represents a Vpn struct.
type Vpn struct {
    // When the object has been created, in epoch
    CreatedTime          *float64               `json:"created_time,omitempty"`
    // Unique ID of the object instance in the Mist Organization
    Id                   *uuid.UUID             `json:"id,omitempty"`
    // When the object has been modified for the last time, in epoch
    ModifiedTime         *float64               `json:"modified_time,omitempty"`
    Name                 string                 `json:"name"`
    OrgId                *uuid.UUID             `json:"org_id,omitempty"`
    // Only if `type`==`hub_spoke`
    PathSelection        *VpnPathSelection      `json:"path_selection,omitempty"`
    // For `type`==`hub_spoke`, Property key is the VPN name. For `type`==`mesh`, Property key is the Interface name
    Paths                map[string]VpnPath     `json:"paths"`
    // enum: `hub_spoke`, `mesh`
    Type                 *VpnModeEnum           `json:"type,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for Vpn,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (v Vpn) String() string {
    return fmt.Sprintf(
    	"Vpn[CreatedTime=%v, Id=%v, ModifiedTime=%v, Name=%v, OrgId=%v, PathSelection=%v, Paths=%v, Type=%v, AdditionalProperties=%v]",
    	v.CreatedTime, v.Id, v.ModifiedTime, v.Name, v.OrgId, v.PathSelection, v.Paths, v.Type, v.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Vpn.
// It customizes the JSON marshaling process for Vpn objects.
func (v Vpn) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(v.AdditionalProperties,
        "created_time", "id", "modified_time", "name", "org_id", "path_selection", "paths", "type"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(v.toMap())
}

// toMap converts the Vpn object to a map representation for JSON marshaling.
func (v Vpn) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, v.AdditionalProperties)
    if v.CreatedTime != nil {
        structMap["created_time"] = v.CreatedTime
    }
    if v.Id != nil {
        structMap["id"] = v.Id
    }
    if v.ModifiedTime != nil {
        structMap["modified_time"] = v.ModifiedTime
    }
    structMap["name"] = v.Name
    if v.OrgId != nil {
        structMap["org_id"] = v.OrgId
    }
    if v.PathSelection != nil {
        structMap["path_selection"] = v.PathSelection.toMap()
    }
    structMap["paths"] = v.Paths
    if v.Type != nil {
        structMap["type"] = v.Type
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Vpn.
// It customizes the JSON unmarshaling process for Vpn objects.
func (v *Vpn) UnmarshalJSON(input []byte) error {
    var temp tempVpn
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "created_time", "id", "modified_time", "name", "org_id", "path_selection", "paths", "type")
    if err != nil {
    	return err
    }
    v.AdditionalProperties = additionalProperties
    
    v.CreatedTime = temp.CreatedTime
    v.Id = temp.Id
    v.ModifiedTime = temp.ModifiedTime
    v.Name = *temp.Name
    v.OrgId = temp.OrgId
    v.PathSelection = temp.PathSelection
    v.Paths = *temp.Paths
    v.Type = temp.Type
    return nil
}

// tempVpn is a temporary struct used for validating the fields of Vpn.
type tempVpn  struct {
    CreatedTime   *float64            `json:"created_time,omitempty"`
    Id            *uuid.UUID          `json:"id,omitempty"`
    ModifiedTime  *float64            `json:"modified_time,omitempty"`
    Name          *string             `json:"name"`
    OrgId         *uuid.UUID          `json:"org_id,omitempty"`
    PathSelection *VpnPathSelection   `json:"path_selection,omitempty"`
    Paths         *map[string]VpnPath `json:"paths"`
    Type          *VpnModeEnum        `json:"type,omitempty"`
}

func (v *tempVpn) validate() error {
    var errs []string
    if v.Name == nil {
        errs = append(errs, "required field `name` is missing for type `vpn`")
    }
    if v.Paths == nil {
        errs = append(errs, "required field `paths` is missing for type `vpn`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
