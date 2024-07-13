package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// InstallerProvisionDevice represents a InstallerProvisionDevice struct.
// Provision Device
type InstallerProvisionDevice struct {
    DeviceprofileName    *string        `json:"deviceprofile_name,omitempty"`
    ForSite              *bool          `json:"for_site,omitempty"`
    Height               *float64       `json:"height,omitempty"`
    MapId                *uuid.UUID     `json:"map_id,omitempty"`
    Name                 string         `json:"name"`
    Orientation          *int           `json:"orientation,omitempty"`
    // Onlif this is to replace an existing device
    ReplacingMac         *string        `json:"replacing_mac,omitempty"`
    // optional role for switch / gateway
    Role                 *string        `json:"role,omitempty"`
    SiteId               *uuid.UUID     `json:"site_id,omitempty"`
    SiteName             *string        `json:"site_name,omitempty"`
    X                    *float64       `json:"x,omitempty"`
    Y                    *float64       `json:"y,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for InstallerProvisionDevice.
// It customizes the JSON marshaling process for InstallerProvisionDevice objects.
func (i InstallerProvisionDevice) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(i.toMap())
}

// toMap converts the InstallerProvisionDevice object to a map representation for JSON marshaling.
func (i InstallerProvisionDevice) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, i.AdditionalProperties)
    if i.DeviceprofileName != nil {
        structMap["deviceprofile_name"] = i.DeviceprofileName
    }
    if i.ForSite != nil {
        structMap["for_site"] = i.ForSite
    }
    if i.Height != nil {
        structMap["height"] = i.Height
    }
    if i.MapId != nil {
        structMap["map_id"] = i.MapId
    }
    structMap["name"] = i.Name
    if i.Orientation != nil {
        structMap["orientation"] = i.Orientation
    }
    if i.ReplacingMac != nil {
        structMap["replacing_mac"] = i.ReplacingMac
    }
    if i.Role != nil {
        structMap["role"] = i.Role
    }
    if i.SiteId != nil {
        structMap["site_id"] = i.SiteId
    }
    if i.SiteName != nil {
        structMap["site_name"] = i.SiteName
    }
    if i.X != nil {
        structMap["x"] = i.X
    }
    if i.Y != nil {
        structMap["y"] = i.Y
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InstallerProvisionDevice.
// It customizes the JSON unmarshaling process for InstallerProvisionDevice objects.
func (i *InstallerProvisionDevice) UnmarshalJSON(input []byte) error {
    var temp installerProvisionDevice
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "deviceprofile_name", "for_site", "height", "map_id", "name", "orientation", "replacing_mac", "role", "site_id", "site_name", "x", "y")
    if err != nil {
    	return err
    }
    
    i.AdditionalProperties = additionalProperties
    i.DeviceprofileName = temp.DeviceprofileName
    i.ForSite = temp.ForSite
    i.Height = temp.Height
    i.MapId = temp.MapId
    i.Name = *temp.Name
    i.Orientation = temp.Orientation
    i.ReplacingMac = temp.ReplacingMac
    i.Role = temp.Role
    i.SiteId = temp.SiteId
    i.SiteName = temp.SiteName
    i.X = temp.X
    i.Y = temp.Y
    return nil
}

// installerProvisionDevice is a temporary struct used for validating the fields of InstallerProvisionDevice.
type installerProvisionDevice  struct {
    DeviceprofileName *string    `json:"deviceprofile_name,omitempty"`
    ForSite           *bool      `json:"for_site,omitempty"`
    Height            *float64   `json:"height,omitempty"`
    MapId             *uuid.UUID `json:"map_id,omitempty"`
    Name              *string    `json:"name"`
    Orientation       *int       `json:"orientation,omitempty"`
    ReplacingMac      *string    `json:"replacing_mac,omitempty"`
    Role              *string    `json:"role,omitempty"`
    SiteId            *uuid.UUID `json:"site_id,omitempty"`
    SiteName          *string    `json:"site_name,omitempty"`
    X                 *float64   `json:"x,omitempty"`
    Y                 *float64   `json:"y,omitempty"`
}

func (i *installerProvisionDevice) validate() error {
    var errs []string
    if i.Name == nil {
        errs = append(errs, "required field `name` is missing for type `Installer_Provision_Device`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
