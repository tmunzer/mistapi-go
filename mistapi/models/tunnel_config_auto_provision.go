// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// TunnelConfigAutoProvision represents a TunnelConfigAutoProvision struct.
// Auto Provisioning configuration for the tunne. This takes precedence over the `primary` and `secondary` nodes.
type TunnelConfigAutoProvision struct {
    // Enable auto provisioning for the tunnel. If enabled, the `primary` and `secondary` nodes will be ignored.
    Enabled              *bool                                 `json:"enabled,omitempty"`
    // API override for POP selection
    Latlng               *TunnelConfigAutoProvisionLatLng      `json:"latlng,omitempty"`
    Primary              *TunnelConfigAutoProvisionNode        `json:"primary,omitempty"`
    // enum: `jse-ipsec`, `zscaler-ipsec`
    Provider             TunnelConfigAutoProvisionProviderEnum `json:"provider"`
    // API override for POP selection in the case user wants to override the auto discovery of remote network location and force the tunnel to use the specified peer location.
    Region               *string                               `json:"region,omitempty"`
    Secondary            *TunnelConfigAutoProvisionNode        `json:"secondary,omitempty"`
    // if `provider`==`prisma-ipsec`. By default, we'll use the location of the site to determine the optimal Remote Network location, optionally, service_connection can be considered, then we'll also consider this along with the site location. Define service_connection if the traffic is to be routed to a specific service connection. This field takes a service connection name that is configured in the Prisma cloud, Prisma Access Setup -> Service Connections.
    ServiceConnection    *string                               `json:"service_connection,omitempty"`
    AdditionalProperties map[string]interface{}                `json:"_"`
}

// String implements the fmt.Stringer interface for TunnelConfigAutoProvision,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (t TunnelConfigAutoProvision) String() string {
    return fmt.Sprintf(
    	"TunnelConfigAutoProvision[Enabled=%v, Latlng=%v, Primary=%v, Provider=%v, Region=%v, Secondary=%v, ServiceConnection=%v, AdditionalProperties=%v]",
    	t.Enabled, t.Latlng, t.Primary, t.Provider, t.Region, t.Secondary, t.ServiceConnection, t.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for TunnelConfigAutoProvision.
// It customizes the JSON marshaling process for TunnelConfigAutoProvision objects.
func (t TunnelConfigAutoProvision) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(t.AdditionalProperties,
        "enabled", "latlng", "primary", "provider", "region", "secondary", "service_connection"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(t.toMap())
}

// toMap converts the TunnelConfigAutoProvision object to a map representation for JSON marshaling.
func (t TunnelConfigAutoProvision) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, t.AdditionalProperties)
    if t.Enabled != nil {
        structMap["enabled"] = t.Enabled
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
    if t.ServiceConnection != nil {
        structMap["service_connection"] = t.ServiceConnection
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "enabled", "latlng", "primary", "provider", "region", "secondary", "service_connection")
    if err != nil {
    	return err
    }
    t.AdditionalProperties = additionalProperties
    
    t.Enabled = temp.Enabled
    t.Latlng = temp.Latlng
    t.Primary = temp.Primary
    t.Provider = *temp.Provider
    t.Region = temp.Region
    t.Secondary = temp.Secondary
    t.ServiceConnection = temp.ServiceConnection
    return nil
}

// tempTunnelConfigAutoProvision is a temporary struct used for validating the fields of TunnelConfigAutoProvision.
type tempTunnelConfigAutoProvision  struct {
    Enabled           *bool                                  `json:"enabled,omitempty"`
    Latlng            *TunnelConfigAutoProvisionLatLng       `json:"latlng,omitempty"`
    Primary           *TunnelConfigAutoProvisionNode         `json:"primary,omitempty"`
    Provider          *TunnelConfigAutoProvisionProviderEnum `json:"provider"`
    Region            *string                                `json:"region,omitempty"`
    Secondary         *TunnelConfigAutoProvisionNode         `json:"secondary,omitempty"`
    ServiceConnection *string                                `json:"service_connection,omitempty"`
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
