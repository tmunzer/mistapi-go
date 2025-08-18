// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// AccountSkyatpDataSecintel represents a AccountSkyatpDataSecintel struct.
// juniper secintel_feeds are enabled depending on your license tier: infected_host, geo_ip, attacker_ip, command_and_control.
// third party:
// * ip-based: block_list, threatfox_ip, feodo_tracker, dshield, tor
// * url-based: threatfox_url, urlhaus, open_phish
// * domain-based: threatfox_domains
type AccountSkyatpDataSecintel struct {
	ThirdPartyThreatFeeds []string               `json:"third_party_threat_feeds,omitempty"`
	AdditionalProperties  map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for AccountSkyatpDataSecintel,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AccountSkyatpDataSecintel) String() string {
	return fmt.Sprintf(
		"AccountSkyatpDataSecintel[ThirdPartyThreatFeeds=%v, AdditionalProperties=%v]",
		a.ThirdPartyThreatFeeds, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AccountSkyatpDataSecintel.
// It customizes the JSON marshaling process for AccountSkyatpDataSecintel objects.
func (a AccountSkyatpDataSecintel) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(a.AdditionalProperties,
		"third_party_threat_feeds"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(a.toMap())
}

// toMap converts the AccountSkyatpDataSecintel object to a map representation for JSON marshaling.
func (a AccountSkyatpDataSecintel) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, a.AdditionalProperties)
	if a.ThirdPartyThreatFeeds != nil {
		structMap["third_party_threat_feeds"] = a.ThirdPartyThreatFeeds
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AccountSkyatpDataSecintel.
// It customizes the JSON unmarshaling process for AccountSkyatpDataSecintel objects.
func (a *AccountSkyatpDataSecintel) UnmarshalJSON(input []byte) error {
	var temp tempAccountSkyatpDataSecintel
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "third_party_threat_feeds")
	if err != nil {
		return err
	}
	a.AdditionalProperties = additionalProperties

	a.ThirdPartyThreatFeeds = temp.ThirdPartyThreatFeeds
	return nil
}

// tempAccountSkyatpDataSecintel is a temporary struct used for validating the fields of AccountSkyatpDataSecintel.
type tempAccountSkyatpDataSecintel struct {
	ThirdPartyThreatFeeds []string `json:"third_party_threat_feeds,omitempty"`
}
