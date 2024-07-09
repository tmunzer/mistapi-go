package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// Guest represents a Guest struct.
// Guest
type Guest struct {
    // whether the guest is current authorized
    Authorized             *bool          `json:"authorized,omitempty"`
    // when the authorization would expire
    AuthorizedExpiringTime *float64       `json:"authorized_expiring_time,omitempty"`
    // when the guest was authorized
    AuthorizedTime         *float64       `json:"authorized_time,omitempty"`
    // optional, the info provided by user
    Company                *string        `json:"company,omitempty"`
    // optional, the info provided by user
    Email                  *string        `json:"email,omitempty"`
    // optional, the info provided by user
    Field1                 *string        `json:"field1,omitempty"`
    Field2                 *string        `json:"field2,omitempty"`
    Field3                 *string        `json:"field3,omitempty"`
    Field4                 *string        `json:"field4,omitempty"`
    // mac
    Mac                    *string        `json:"mac,omitempty"`
    // minutes, the maximum is 259200 (180 days)
    Minutes                *int           `json:"minutes,omitempty"`
    // optional, the info provided by user
    Name                   *string        `json:"name,omitempty"`
    Ssid                   *string        `json:"ssid,omitempty"`
    WlanId                 *uuid.UUID     `json:"wlan_id,omitempty"`
    AdditionalProperties   map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for Guest.
// It customizes the JSON marshaling process for Guest objects.
func (g Guest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the Guest object to a map representation for JSON marshaling.
func (g Guest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, g.AdditionalProperties)
    if g.Authorized != nil {
        structMap["authorized"] = g.Authorized
    }
    if g.AuthorizedExpiringTime != nil {
        structMap["authorized_expiring_time"] = g.AuthorizedExpiringTime
    }
    if g.AuthorizedTime != nil {
        structMap["authorized_time"] = g.AuthorizedTime
    }
    if g.Company != nil {
        structMap["company"] = g.Company
    }
    if g.Email != nil {
        structMap["email"] = g.Email
    }
    if g.Field1 != nil {
        structMap["field1"] = g.Field1
    }
    if g.Field2 != nil {
        structMap["field2"] = g.Field2
    }
    if g.Field3 != nil {
        structMap["field3"] = g.Field3
    }
    if g.Field4 != nil {
        structMap["field4"] = g.Field4
    }
    if g.Mac != nil {
        structMap["mac"] = g.Mac
    }
    if g.Minutes != nil {
        structMap["minutes"] = g.Minutes
    }
    if g.Name != nil {
        structMap["name"] = g.Name
    }
    if g.Ssid != nil {
        structMap["ssid"] = g.Ssid
    }
    if g.WlanId != nil {
        structMap["wlan_id"] = g.WlanId
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Guest.
// It customizes the JSON unmarshaling process for Guest objects.
func (g *Guest) UnmarshalJSON(input []byte) error {
    var temp guest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "authorized", "authorized_expiring_time", "authorized_time", "company", "email", "field1", "field2", "field3", "field4", "mac", "minutes", "name", "ssid", "wlan_id")
    if err != nil {
    	return err
    }
    
    g.AdditionalProperties = additionalProperties
    g.Authorized = temp.Authorized
    g.AuthorizedExpiringTime = temp.AuthorizedExpiringTime
    g.AuthorizedTime = temp.AuthorizedTime
    g.Company = temp.Company
    g.Email = temp.Email
    g.Field1 = temp.Field1
    g.Field2 = temp.Field2
    g.Field3 = temp.Field3
    g.Field4 = temp.Field4
    g.Mac = temp.Mac
    g.Minutes = temp.Minutes
    g.Name = temp.Name
    g.Ssid = temp.Ssid
    g.WlanId = temp.WlanId
    return nil
}

// guest is a temporary struct used for validating the fields of Guest.
type guest  struct {
    Authorized             *bool      `json:"authorized,omitempty"`
    AuthorizedExpiringTime *float64   `json:"authorized_expiring_time,omitempty"`
    AuthorizedTime         *float64   `json:"authorized_time,omitempty"`
    Company                *string    `json:"company,omitempty"`
    Email                  *string    `json:"email,omitempty"`
    Field1                 *string    `json:"field1,omitempty"`
    Field2                 *string    `json:"field2,omitempty"`
    Field3                 *string    `json:"field3,omitempty"`
    Field4                 *string    `json:"field4,omitempty"`
    Mac                    *string    `json:"mac,omitempty"`
    Minutes                *int       `json:"minutes,omitempty"`
    Name                   *string    `json:"name,omitempty"`
    Ssid                   *string    `json:"ssid,omitempty"`
    WlanId                 *uuid.UUID `json:"wlan_id,omitempty"`
}
