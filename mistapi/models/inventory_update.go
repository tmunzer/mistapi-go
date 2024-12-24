package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// InventoryUpdate represents a InventoryUpdate struct.
type InventoryUpdate struct {
    // if `op`==`assign`, a **cloud-ready** switch/gateway will be managed/configured by Mist by default, this disabled the behavior
    DisableAutoConfig    *bool                        `json:"disable_auto_config,omitempty"`
    // if `op`==`assign`, `op`==`unassign`, `op`==`upgrade_to_mist`or `op`==`downgrade_to_jsi` , list of MAC, e.g. ["5c5b350e0001"]
    Macs                 []string                     `json:"macs,omitempty"`
    // if `op`==`assign`, an **adopted** switch/gateway will not be managed/configured by Mist by default, this enables the behavior
    Managed              *bool                        `json:"managed,omitempty"`
    // if `op`==`assign`, if true, treat site assignment against an already assigned AP as error
    NoReassign           *bool                        `json:"no_reassign,omitempty"`
    // enum:
    // * `upgrade_to_mist`: Upgrade to mist-managed
    // * `downgrade_to_jsi`: Downgrade to basic monitoring. When downgrading a VC member to jsi, we will move the cloud connection of the VC to jsi-terminator and keep all VC device/inventories intact for pain-free upgrading back to mist.
    // * `assign`: Assign inventory to a site
    // * `unassign`: Unassign inventory from a site
    // * `delete`: Delete multiple inventory from org. If the device is already assigned to a site, it will be unassigned
    Op                   InventoryUpdateOperationEnum `json:"op"`
    // if `op`==`delete`, list of serial numbers, e.g. ["FXLH2015150025"]
    Serials              []string                     `json:"serials,omitempty"`
    // if `op`==`assign`, target site id
    SiteId               *uuid.UUID                   `json:"site_id,omitempty"`
    AdditionalProperties map[string]interface{}       `json:"_"`
}

// String implements the fmt.Stringer interface for InventoryUpdate,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (i InventoryUpdate) String() string {
    return fmt.Sprintf(
    	"InventoryUpdate[DisableAutoConfig=%v, Macs=%v, Managed=%v, NoReassign=%v, Op=%v, Serials=%v, SiteId=%v, AdditionalProperties=%v]",
    	i.DisableAutoConfig, i.Macs, i.Managed, i.NoReassign, i.Op, i.Serials, i.SiteId, i.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for InventoryUpdate.
// It customizes the JSON marshaling process for InventoryUpdate objects.
func (i InventoryUpdate) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(i.AdditionalProperties,
        "disable_auto_config", "macs", "managed", "no_reassign", "op", "serials", "site_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(i.toMap())
}

// toMap converts the InventoryUpdate object to a map representation for JSON marshaling.
func (i InventoryUpdate) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, i.AdditionalProperties)
    if i.DisableAutoConfig != nil {
        structMap["disable_auto_config"] = i.DisableAutoConfig
    }
    if i.Macs != nil {
        structMap["macs"] = i.Macs
    }
    if i.Managed != nil {
        structMap["managed"] = i.Managed
    }
    if i.NoReassign != nil {
        structMap["no_reassign"] = i.NoReassign
    }
    structMap["op"] = i.Op
    if i.Serials != nil {
        structMap["serials"] = i.Serials
    }
    if i.SiteId != nil {
        structMap["site_id"] = i.SiteId
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InventoryUpdate.
// It customizes the JSON unmarshaling process for InventoryUpdate objects.
func (i *InventoryUpdate) UnmarshalJSON(input []byte) error {
    var temp tempInventoryUpdate
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "disable_auto_config", "macs", "managed", "no_reassign", "op", "serials", "site_id")
    if err != nil {
    	return err
    }
    i.AdditionalProperties = additionalProperties
    
    i.DisableAutoConfig = temp.DisableAutoConfig
    i.Macs = temp.Macs
    i.Managed = temp.Managed
    i.NoReassign = temp.NoReassign
    i.Op = *temp.Op
    i.Serials = temp.Serials
    i.SiteId = temp.SiteId
    return nil
}

// tempInventoryUpdate is a temporary struct used for validating the fields of InventoryUpdate.
type tempInventoryUpdate  struct {
    DisableAutoConfig *bool                         `json:"disable_auto_config,omitempty"`
    Macs              []string                      `json:"macs,omitempty"`
    Managed           *bool                         `json:"managed,omitempty"`
    NoReassign        *bool                         `json:"no_reassign,omitempty"`
    Op                *InventoryUpdateOperationEnum `json:"op"`
    Serials           []string                      `json:"serials,omitempty"`
    SiteId            *uuid.UUID                    `json:"site_id,omitempty"`
}

func (i *tempInventoryUpdate) validate() error {
    var errs []string
    if i.Op == nil {
        errs = append(errs, "required field `op` is missing for type `inventory_update`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
