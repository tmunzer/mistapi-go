// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// AccountSkyatpData represents a AccountSkyatpData struct.
type AccountSkyatpData struct {
	// juniper secintel_feeds are enabled depending on your license tier: infected_host, geo_ip, attacker_ip, command_and_control.
	// third party:
	// * ip-based: block_list, threatfox_ip, feodo_tracker, dshield, tor
	// * url-based: threatfox_url, urlhaus, open_phish
	// * domain-based: threatfox_domains
	Secintel             *AccountSkyatpDataSecintel `json:"secintel,omitempty"`
	SecintelAllowlistUrl *string                    `json:"secintel_allowlist_url,omitempty"`
	SecintelBlocklistUrl *string                    `json:"secintel_blocklist_url,omitempty"`
	AdditionalProperties map[string]interface{}     `json:"_"`
}

// String implements the fmt.Stringer interface for AccountSkyatpData,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AccountSkyatpData) String() string {
	return fmt.Sprintf(
		"AccountSkyatpData[Secintel=%v, SecintelAllowlistUrl=%v, SecintelBlocklistUrl=%v, AdditionalProperties=%v]",
		a.Secintel, a.SecintelAllowlistUrl, a.SecintelBlocklistUrl, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AccountSkyatpData.
// It customizes the JSON marshaling process for AccountSkyatpData objects.
func (a AccountSkyatpData) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(a.AdditionalProperties,
		"secintel", "secintel_allowlist_url", "secintel_blocklist_url"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(a.toMap())
}

// toMap converts the AccountSkyatpData object to a map representation for JSON marshaling.
func (a AccountSkyatpData) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, a.AdditionalProperties)
	if a.Secintel != nil {
		structMap["secintel"] = a.Secintel.toMap()
	}
	if a.SecintelAllowlistUrl != nil {
		structMap["secintel_allowlist_url"] = a.SecintelAllowlistUrl
	}
	if a.SecintelBlocklistUrl != nil {
		structMap["secintel_blocklist_url"] = a.SecintelBlocklistUrl
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AccountSkyatpData.
// It customizes the JSON unmarshaling process for AccountSkyatpData objects.
func (a *AccountSkyatpData) UnmarshalJSON(input []byte) error {
	var temp tempAccountSkyatpData
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "secintel", "secintel_allowlist_url", "secintel_blocklist_url")
	if err != nil {
		return err
	}
	a.AdditionalProperties = additionalProperties

	a.Secintel = temp.Secintel
	a.SecintelAllowlistUrl = temp.SecintelAllowlistUrl
	a.SecintelBlocklistUrl = temp.SecintelBlocklistUrl
	return nil
}

// tempAccountSkyatpData is a temporary struct used for validating the fields of AccountSkyatpData.
type tempAccountSkyatpData struct {
	Secintel             *AccountSkyatpDataSecintel `json:"secintel,omitempty"`
	SecintelAllowlistUrl *string                    `json:"secintel_allowlist_url,omitempty"`
	SecintelBlocklistUrl *string                    `json:"secintel_blocklist_url,omitempty"`
}
