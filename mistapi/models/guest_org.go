package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// GuestOrg represents a GuestOrg struct.
// Guest
type GuestOrg struct {
    // if `auth_method`==`email`, the email address where the authorization code has been sent to
    AccessCodeEmail        *string                `json:"access_code_email,omitempty"`
    // based on the WLAN portal configuration (field `allow_wlan_id_roam`), if the user is also authorized on other Guest WLANs of the same Org without reauthentication
    AllowWlanIdRoam        *bool                  `json:"allow_wlan_id_roam,omitempty"`
    // the MAC Address of the AP the guest was connected to during the registration process
    ApMac                  *string                `json:"ap_mac,omitempty"`
    // type of guest authorization
    AuthMethod             *string                `json:"auth_method,omitempty"`
    // whether the guest is current authorized
    Authorized             *bool                  `json:"authorized,omitempty"`
    // when the authorization would expire
    AuthorizedExpiringTime *float64               `json:"authorized_expiring_time,omitempty"`
    // when the guest was authorized
    AuthorizedTime         *float64               `json:"authorized_time,omitempty"`
    // optional, the info provided by user
    Company                *string                `json:"company,omitempty"`
    // based on the WLAN portal configuration (field `cross_site`), if the user is also authorized on other sites (same `wlan.ssid`) of the same Org without reauthentication
    CrossSite              *bool                  `json:"cross_site,omitempty"`
    // optional, the info provided by user
    Email                  *string                `json:"email,omitempty"`
    // optional, the info provided by user
    Field1                 *string                `json:"field1,omitempty"`
    Field2                 *string                `json:"field2,omitempty"`
    Field3                 *string                `json:"field3,omitempty"`
    Field4                 *string                `json:"field4,omitempty"`
    // mac
    Mac                    *string                `json:"mac,omitempty"`
    // Authorization duration, in minutes. Default is 1440 minutes (1 day), maximum is 259200 (180 days)
    Minutes                *int                   `json:"minutes,omitempty"`
    // optional, the info provided by user
    Name                   *string                `json:"name,omitempty"`
    // if the client is using a randomized MAC Address to connect the SSID
    RandomMac              *bool                  `json:"random_mac,omitempty"`
    // name of the SSID
    Ssid                   *string                `json:"ssid,omitempty"`
    // ID of the WLAN
    WlanId                 uuid.UUID              `json:"wlan_id"`
    AdditionalProperties   map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for GuestOrg.
// It customizes the JSON marshaling process for GuestOrg objects.
func (g GuestOrg) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(g.AdditionalProperties,
        "access_code_email", "allow_wlan_id_roam", "ap_mac", "auth_method", "authorized", "authorized_expiring_time", "authorized_time", "company", "cross_site", "email", "field1", "field2", "field3", "field4", "mac", "minutes", "name", "random_mac", "ssid", "wlan_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(g.toMap())
}

// toMap converts the GuestOrg object to a map representation for JSON marshaling.
func (g GuestOrg) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, g.AdditionalProperties)
    if g.AccessCodeEmail != nil {
        structMap["access_code_email"] = g.AccessCodeEmail
    }
    if g.AllowWlanIdRoam != nil {
        structMap["allow_wlan_id_roam"] = g.AllowWlanIdRoam
    }
    if g.ApMac != nil {
        structMap["ap_mac"] = g.ApMac
    }
    if g.AuthMethod != nil {
        structMap["auth_method"] = g.AuthMethod
    }
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
    if g.CrossSite != nil {
        structMap["cross_site"] = g.CrossSite
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
    if g.RandomMac != nil {
        structMap["random_mac"] = g.RandomMac
    }
    if g.Ssid != nil {
        structMap["ssid"] = g.Ssid
    }
    structMap["wlan_id"] = g.WlanId
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GuestOrg.
// It customizes the JSON unmarshaling process for GuestOrg objects.
func (g *GuestOrg) UnmarshalJSON(input []byte) error {
    var temp tempGuestOrg
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "access_code_email", "allow_wlan_id_roam", "ap_mac", "auth_method", "authorized", "authorized_expiring_time", "authorized_time", "company", "cross_site", "email", "field1", "field2", "field3", "field4", "mac", "minutes", "name", "random_mac", "ssid", "wlan_id")
    if err != nil {
    	return err
    }
    g.AdditionalProperties = additionalProperties
    
    g.AccessCodeEmail = temp.AccessCodeEmail
    g.AllowWlanIdRoam = temp.AllowWlanIdRoam
    g.ApMac = temp.ApMac
    g.AuthMethod = temp.AuthMethod
    g.Authorized = temp.Authorized
    g.AuthorizedExpiringTime = temp.AuthorizedExpiringTime
    g.AuthorizedTime = temp.AuthorizedTime
    g.Company = temp.Company
    g.CrossSite = temp.CrossSite
    g.Email = temp.Email
    g.Field1 = temp.Field1
    g.Field2 = temp.Field2
    g.Field3 = temp.Field3
    g.Field4 = temp.Field4
    g.Mac = temp.Mac
    g.Minutes = temp.Minutes
    g.Name = temp.Name
    g.RandomMac = temp.RandomMac
    g.Ssid = temp.Ssid
    g.WlanId = *temp.WlanId
    return nil
}

// tempGuestOrg is a temporary struct used for validating the fields of GuestOrg.
type tempGuestOrg  struct {
    AccessCodeEmail        *string    `json:"access_code_email,omitempty"`
    AllowWlanIdRoam        *bool      `json:"allow_wlan_id_roam,omitempty"`
    ApMac                  *string    `json:"ap_mac,omitempty"`
    AuthMethod             *string    `json:"auth_method,omitempty"`
    Authorized             *bool      `json:"authorized,omitempty"`
    AuthorizedExpiringTime *float64   `json:"authorized_expiring_time,omitempty"`
    AuthorizedTime         *float64   `json:"authorized_time,omitempty"`
    Company                *string    `json:"company,omitempty"`
    CrossSite              *bool      `json:"cross_site,omitempty"`
    Email                  *string    `json:"email,omitempty"`
    Field1                 *string    `json:"field1,omitempty"`
    Field2                 *string    `json:"field2,omitempty"`
    Field3                 *string    `json:"field3,omitempty"`
    Field4                 *string    `json:"field4,omitempty"`
    Mac                    *string    `json:"mac,omitempty"`
    Minutes                *int       `json:"minutes,omitempty"`
    Name                   *string    `json:"name,omitempty"`
    RandomMac              *bool      `json:"random_mac,omitempty"`
    Ssid                   *string    `json:"ssid,omitempty"`
    WlanId                 *uuid.UUID `json:"wlan_id"`
}

func (g *tempGuestOrg) validate() error {
    var errs []string
    if g.WlanId == nil {
        errs = append(errs, "required field `wlan_id` is missing for type `guest_org`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
