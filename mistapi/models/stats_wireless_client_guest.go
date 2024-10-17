package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// StatsWirelessClientGuest represents a StatsWirelessClientGuest struct.
// information about this portal
type StatsWirelessClientGuest struct {
	// whether this guest is authorized
	Authorized bool `json:"authorized"`
	// when the guest authorization will expire
	AuthorizedExpiringTime float64 `json:"authorized_expiring_time"`
	// when the guest is authorized
	AuthorizedTime       float64        `json:"authorized_time"`
	Company              string         `json:"company"`
	Email                string         `json:"email"`
	Field1               string         `json:"field1"`
	Name                 string         `json:"name"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for StatsWirelessClientGuest.
// It customizes the JSON marshaling process for StatsWirelessClientGuest objects.
func (s StatsWirelessClientGuest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the StatsWirelessClientGuest object to a map representation for JSON marshaling.
func (s StatsWirelessClientGuest) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, s.AdditionalProperties)
	structMap["authorized"] = s.Authorized
	structMap["authorized_expiring_time"] = s.AuthorizedExpiringTime
	structMap["authorized_time"] = s.AuthorizedTime
	structMap["company"] = s.Company
	structMap["email"] = s.Email
	structMap["field1"] = s.Field1
	structMap["name"] = s.Name
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsWirelessClientGuest.
// It customizes the JSON unmarshaling process for StatsWirelessClientGuest objects.
func (s *StatsWirelessClientGuest) UnmarshalJSON(input []byte) error {
	var temp tempStatsWirelessClientGuest
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "authorized", "authorized_expiring_time", "authorized_time", "company", "email", "field1", "name")
	if err != nil {
		return err
	}

	s.AdditionalProperties = additionalProperties
	s.Authorized = *temp.Authorized
	s.AuthorizedExpiringTime = *temp.AuthorizedExpiringTime
	s.AuthorizedTime = *temp.AuthorizedTime
	s.Company = *temp.Company
	s.Email = *temp.Email
	s.Field1 = *temp.Field1
	s.Name = *temp.Name
	return nil
}

// tempStatsWirelessClientGuest is a temporary struct used for validating the fields of StatsWirelessClientGuest.
type tempStatsWirelessClientGuest struct {
	Authorized             *bool    `json:"authorized"`
	AuthorizedExpiringTime *float64 `json:"authorized_expiring_time"`
	AuthorizedTime         *float64 `json:"authorized_time"`
	Company                *string  `json:"company"`
	Email                  *string  `json:"email"`
	Field1                 *string  `json:"field1"`
	Name                   *string  `json:"name"`
}

func (s *tempStatsWirelessClientGuest) validate() error {
	var errs []string
	if s.Authorized == nil {
		errs = append(errs, "required field `authorized` is missing for type `stats_wireless_client_guest`")
	}
	if s.AuthorizedExpiringTime == nil {
		errs = append(errs, "required field `authorized_expiring_time` is missing for type `stats_wireless_client_guest`")
	}
	if s.AuthorizedTime == nil {
		errs = append(errs, "required field `authorized_time` is missing for type `stats_wireless_client_guest`")
	}
	if s.Company == nil {
		errs = append(errs, "required field `company` is missing for type `stats_wireless_client_guest`")
	}
	if s.Email == nil {
		errs = append(errs, "required field `email` is missing for type `stats_wireless_client_guest`")
	}
	if s.Field1 == nil {
		errs = append(errs, "required field `field1` is missing for type `stats_wireless_client_guest`")
	}
	if s.Name == nil {
		errs = append(errs, "required field `name` is missing for type `stats_wireless_client_guest`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
