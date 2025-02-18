package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// TunnelConfigAutoProvision represents a TunnelConfigAutoProvision struct.
type TunnelConfigAutoProvision struct {
    Enable               *bool                                 `json:"enable,omitempty"`
    // API override for POP selection
    Latlng               *TunnelConfigAutoProvisionLatLng      `json:"latlng,omitempty"`
    Primary              *TunnelConfigAutoProvisionNode        `json:"primary,omitempty"`
    // enum: `jse-ipsec`, `zscaler-ipsec`
    Provider             TunnelConfigAutoProvisionProviderEnum `json:"provider"`
    // API override for POP selection
    Region               *string                               `json:"region,omitempty"`
    Secondary            *TunnelConfigAutoProvisionNode        `json:"secondary,omitempty"`
    AdditionalProperties map[string]interface{}                `json:"_"`
}

// String implements the fmt.Stringer interface for TunnelConfigAutoProvision,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (t TunnelConfigAutoProvision) String() string {
    return fmt.Sprintf(
    	"TunnelConfigAutoProvision[Enable=%v, Latlng=%v, Primary=%v, Provider=%v, Region=%v, Secondary=%v, AdditionalProperties=%v]",
    	t.Enable, t.Latlng, t.Primary, t.Provider, t.Region, t.Secondary, t.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for TunnelConfigAutoProvision.
// It customizes the JSON marshaling process for TunnelConfigAutoProvision objects.
func (t TunnelConfigAutoProvision) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(t.AdditionalProperties,
        "enable", "latlng", "primary", "provider", "region", "secondary"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(t.toMap())
}

// toMap converts the TunnelConfigAutoProvision object to a map representation for JSON marshaling.
func (t TunnelConfigAutoProvision) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, t.AdditionalProperties)
    if t.Enable != nil {
        structMap["enable"] = t.Enable
    }
    if t.Latlng != nil {
        structMap["latlng"] = t.Latlng.toMap()
    }
    if t.Primary != nil {
        structMap["primary"] = t.Primary.toMap()
    }
    structMap["provider"] = t.Provider
    if t.Region != nil {
        structMap["region"] = t.Region
    }
    if t.Secondary != nil {
        structMap["secondary"] = t.Secondary.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TunnelConfigAutoProvision.
// It customizes the JSON unmarshaling process for TunnelConfigAutoProvision objects.
func (t *TunnelConfigAutoProvision) UnmarshalJSON(input []byte) error {
    var temp tempTunnelConfigAutoProvision
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "enable", "latlng", "primary", "provider", "region", "secondary")
    if err != nil {
    	return err
    }
    t.AdditionalProperties = additionalProperties
    
    t.Enable = temp.Enable
    t.Latlng = temp.Latlng
    t.Primary = temp.Primary
    t.Provider = *temp.Provider
    t.Region = temp.Region
    t.Secondary = temp.Secondary
    return nil
}

// tempTunnelConfigAutoProvision is a temporary struct used for validating the fields of TunnelConfigAutoProvision.
type tempTunnelConfigAutoProvision  struct {
    Enable    *bool                                  `json:"enable,omitempty"`
    Latlng    *TunnelConfigAutoProvisionLatLng       `json:"latlng,omitempty"`
    Primary   *TunnelConfigAutoProvisionNode         `json:"primary,omitempty"`
    Provider  *TunnelConfigAutoProvisionProviderEnum `json:"provider"`
    Region    *string                                `json:"region,omitempty"`
    Secondary *TunnelConfigAutoProvisionNode         `json:"secondary,omitempty"`
}

func (t *tempTunnelConfigAutoProvision) validate() error {
    var errs []string
    if t.Provider == nil {
        errs = append(errs, "required field `provider` is missing for type `tunnel_config_auto_provision`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
