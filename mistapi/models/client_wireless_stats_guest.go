package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ClientWirelessStatsGuest represents a ClientWirelessStatsGuest struct.
// information about this portal
type ClientWirelessStatsGuest struct {
    // whether this guest is authorized
    Authorized             bool           `json:"authorized"`
    // when the guest authorization will expire
    AuthorizedExpiringTime float64        `json:"authorized_expiring_time"`
    // when the guest is authorized
    AuthorizedTime         float64        `json:"authorized_time"`
    Company                string         `json:"company"`
    Email                  string         `json:"email"`
    Field1                 string         `json:"field1"`
    Name                   string         `json:"name"`
    AdditionalProperties   map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ClientWirelessStatsGuest.
// It customizes the JSON marshaling process for ClientWirelessStatsGuest objects.
func (c ClientWirelessStatsGuest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the ClientWirelessStatsGuest object to a map representation for JSON marshaling.
func (c ClientWirelessStatsGuest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["authorized"] = c.Authorized
    structMap["authorized_expiring_time"] = c.AuthorizedExpiringTime
    structMap["authorized_time"] = c.AuthorizedTime
    structMap["company"] = c.Company
    structMap["email"] = c.Email
    structMap["field1"] = c.Field1
    structMap["name"] = c.Name
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ClientWirelessStatsGuest.
// It customizes the JSON unmarshaling process for ClientWirelessStatsGuest objects.
func (c *ClientWirelessStatsGuest) UnmarshalJSON(input []byte) error {
    var temp tempClientWirelessStatsGuest
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
    
    c.AdditionalProperties = additionalProperties
    c.Authorized = *temp.Authorized
    c.AuthorizedExpiringTime = *temp.AuthorizedExpiringTime
    c.AuthorizedTime = *temp.AuthorizedTime
    c.Company = *temp.Company
    c.Email = *temp.Email
    c.Field1 = *temp.Field1
    c.Name = *temp.Name
    return nil
}

// tempClientWirelessStatsGuest is a temporary struct used for validating the fields of ClientWirelessStatsGuest.
type tempClientWirelessStatsGuest  struct {
    Authorized             *bool    `json:"authorized"`
    AuthorizedExpiringTime *float64 `json:"authorized_expiring_time"`
    AuthorizedTime         *float64 `json:"authorized_time"`
    Company                *string  `json:"company"`
    Email                  *string  `json:"email"`
    Field1                 *string  `json:"field1"`
    Name                   *string  `json:"name"`
}

func (c *tempClientWirelessStatsGuest) validate() error {
    var errs []string
    if c.Authorized == nil {
        errs = append(errs, "required field `authorized` is missing for type `client_wireless_stats_guest`")
    }
    if c.AuthorizedExpiringTime == nil {
        errs = append(errs, "required field `authorized_expiring_time` is missing for type `client_wireless_stats_guest`")
    }
    if c.AuthorizedTime == nil {
        errs = append(errs, "required field `authorized_time` is missing for type `client_wireless_stats_guest`")
    }
    if c.Company == nil {
        errs = append(errs, "required field `company` is missing for type `client_wireless_stats_guest`")
    }
    if c.Email == nil {
        errs = append(errs, "required field `email` is missing for type `client_wireless_stats_guest`")
    }
    if c.Field1 == nil {
        errs = append(errs, "required field `field1` is missing for type `client_wireless_stats_guest`")
    }
    if c.Name == nil {
        errs = append(errs, "required field `name` is missing for type `client_wireless_stats_guest`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
