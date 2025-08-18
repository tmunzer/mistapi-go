// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// AccountPrismaConfig represents a AccountPrismaConfig struct.
// OAuth linked CrowdStrike apps account details
type AccountPrismaConfig struct {
	// Required If `enable_probe`==`true`. This field will accept an IPv4 cidr and an IP address will be picked from this range to be used as tunnel probe source ip address and as well as BGP neighbour IP address. The subnet should be big enough for num_devices * num_tunnel * 2
	AutoProbeSubnet *string `json:"auto_probe_subnet,omitempty"`
	// Customer account api client ID
	ClientId string `json:"client_id"`
	// Customer account api client Secret
	ClientSecret string `json:"client_secret"`
	// To enable/disable tunnel probe
	EnableProbe *bool `json:"enable_probe,omitempty"`
	// Prisma Tenant Service Group id
	TsgId                string                 `json:"tsg_id"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for AccountPrismaConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AccountPrismaConfig) String() string {
	return fmt.Sprintf(
		"AccountPrismaConfig[AutoProbeSubnet=%v, ClientId=%v, ClientSecret=%v, EnableProbe=%v, TsgId=%v, AdditionalProperties=%v]",
		a.AutoProbeSubnet, a.ClientId, a.ClientSecret, a.EnableProbe, a.TsgId, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AccountPrismaConfig.
// It customizes the JSON marshaling process for AccountPrismaConfig objects.
func (a AccountPrismaConfig) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(a.AdditionalProperties,
		"auto_probe_subnet", "client_id", "client_secret", "enable_probe", "tsg_id"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(a.toMap())
}

// toMap converts the AccountPrismaConfig object to a map representation for JSON marshaling.
func (a AccountPrismaConfig) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, a.AdditionalProperties)
	if a.AutoProbeSubnet != nil {
		structMap["auto_probe_subnet"] = a.AutoProbeSubnet
	}
	structMap["client_id"] = a.ClientId
	structMap["client_secret"] = a.ClientSecret
	if a.EnableProbe != nil {
		structMap["enable_probe"] = a.EnableProbe
	}
	structMap["tsg_id"] = a.TsgId
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AccountPrismaConfig.
// It customizes the JSON unmarshaling process for AccountPrismaConfig objects.
func (a *AccountPrismaConfig) UnmarshalJSON(input []byte) error {
	var temp tempAccountPrismaConfig
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "auto_probe_subnet", "client_id", "client_secret", "enable_probe", "tsg_id")
	if err != nil {
		return err
	}
	a.AdditionalProperties = additionalProperties

	a.AutoProbeSubnet = temp.AutoProbeSubnet
	a.ClientId = *temp.ClientId
	a.ClientSecret = *temp.ClientSecret
	a.EnableProbe = temp.EnableProbe
	a.TsgId = *temp.TsgId
	return nil
}

// tempAccountPrismaConfig is a temporary struct used for validating the fields of AccountPrismaConfig.
type tempAccountPrismaConfig struct {
	AutoProbeSubnet *string `json:"auto_probe_subnet,omitempty"`
	ClientId        *string `json:"client_id"`
	ClientSecret    *string `json:"client_secret"`
	EnableProbe     *bool   `json:"enable_probe,omitempty"`
	TsgId           *string `json:"tsg_id"`
}

func (a *tempAccountPrismaConfig) validate() error {
	var errs []string
	if a.ClientId == nil {
		errs = append(errs, "required field `client_id` is missing for type `account_prisma_config`")
	}
	if a.ClientSecret == nil {
		errs = append(errs, "required field `client_secret` is missing for type `account_prisma_config`")
	}
	if a.TsgId == nil {
		errs = append(errs, "required field `tsg_id` is missing for type `account_prisma_config`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
