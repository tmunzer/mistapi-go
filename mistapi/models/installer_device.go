package models

import (
	"encoding/json"
	"github.com/google/uuid"
)

// InstallerDevice represents a InstallerDevice struct.
type InstallerDevice struct {
	Connected            *bool            `json:"connected,omitempty"`
	DeviceprofileName    *string          `json:"deviceprofile_name,omitempty"`
	ExtIp                *string          `json:"ext_ip,omitempty"`
	Height               *float64         `json:"height,omitempty"`
	Ip                   *string          `json:"ip,omitempty"`
	LastSeen             *float64         `json:"last_seen,omitempty"`
	Mac                  *string          `json:"mac,omitempty"`
	MapId                *uuid.UUID       `json:"map_id,omitempty"`
	Model                *string          `json:"model,omitempty"`
	Name                 *string          `json:"name,omitempty"`
	Orientation          *int             `json:"orientation,omitempty"`
	Serial               *string          `json:"serial,omitempty"`
	SiteName             *string          `json:"site_name,omitempty"`
	Uptime               *int             `json:"uptime,omitempty"`
	VcMac                Optional[string] `json:"vc_mac"`
	Version              *string          `json:"version,omitempty"`
	X                    *float64         `json:"x,omitempty"`
	Y                    *float64         `json:"y,omitempty"`
	AdditionalProperties map[string]any   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for InstallerDevice.
// It customizes the JSON marshaling process for InstallerDevice objects.
func (i InstallerDevice) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(i.toMap())
}

// toMap converts the InstallerDevice object to a map representation for JSON marshaling.
func (i InstallerDevice) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, i.AdditionalProperties)
	if i.Connected != nil {
		structMap["connected"] = i.Connected
	}
	if i.DeviceprofileName != nil {
		structMap["deviceprofile_name"] = i.DeviceprofileName
	}
	if i.ExtIp != nil {
		structMap["ext_ip"] = i.ExtIp
	}
	if i.Height != nil {
		structMap["height"] = i.Height
	}
	if i.Ip != nil {
		structMap["ip"] = i.Ip
	}
	if i.LastSeen != nil {
		structMap["last_seen"] = i.LastSeen
	}
	if i.Mac != nil {
		structMap["mac"] = i.Mac
	}
	if i.MapId != nil {
		structMap["map_id"] = i.MapId
	}
	if i.Model != nil {
		structMap["model"] = i.Model
	}
	if i.Name != nil {
		structMap["name"] = i.Name
	}
	if i.Orientation != nil {
		structMap["orientation"] = i.Orientation
	}
	if i.Serial != nil {
		structMap["serial"] = i.Serial
	}
	if i.SiteName != nil {
		structMap["site_name"] = i.SiteName
	}
	if i.Uptime != nil {
		structMap["uptime"] = i.Uptime
	}
	if i.VcMac.IsValueSet() {
		if i.VcMac.Value() != nil {
			structMap["vc_mac"] = i.VcMac.Value()
		} else {
			structMap["vc_mac"] = nil
		}
	}
	if i.Version != nil {
		structMap["version"] = i.Version
	}
	if i.X != nil {
		structMap["x"] = i.X
	}
	if i.Y != nil {
		structMap["y"] = i.Y
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InstallerDevice.
// It customizes the JSON unmarshaling process for InstallerDevice objects.
func (i *InstallerDevice) UnmarshalJSON(input []byte) error {
	var temp tempInstallerDevice
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "connected", "deviceprofile_name", "ext_ip", "height", "ip", "last_seen", "mac", "map_id", "model", "name", "orientation", "serial", "site_name", "uptime", "vc_mac", "version", "x", "y")
	if err != nil {
		return err
	}

	i.AdditionalProperties = additionalProperties
	i.Connected = temp.Connected
	i.DeviceprofileName = temp.DeviceprofileName
	i.ExtIp = temp.ExtIp
	i.Height = temp.Height
	i.Ip = temp.Ip
	i.LastSeen = temp.LastSeen
	i.Mac = temp.Mac
	i.MapId = temp.MapId
	i.Model = temp.Model
	i.Name = temp.Name
	i.Orientation = temp.Orientation
	i.Serial = temp.Serial
	i.SiteName = temp.SiteName
	i.Uptime = temp.Uptime
	i.VcMac = temp.VcMac
	i.Version = temp.Version
	i.X = temp.X
	i.Y = temp.Y
	return nil
}

// tempInstallerDevice is a temporary struct used for validating the fields of InstallerDevice.
type tempInstallerDevice struct {
	Connected         *bool            `json:"connected,omitempty"`
	DeviceprofileName *string          `json:"deviceprofile_name,omitempty"`
	ExtIp             *string          `json:"ext_ip,omitempty"`
	Height            *float64         `json:"height,omitempty"`
	Ip                *string          `json:"ip,omitempty"`
	LastSeen          *float64         `json:"last_seen,omitempty"`
	Mac               *string          `json:"mac,omitempty"`
	MapId             *uuid.UUID       `json:"map_id,omitempty"`
	Model             *string          `json:"model,omitempty"`
	Name              *string          `json:"name,omitempty"`
	Orientation       *int             `json:"orientation,omitempty"`
	Serial            *string          `json:"serial,omitempty"`
	SiteName          *string          `json:"site_name,omitempty"`
	Uptime            *int             `json:"uptime,omitempty"`
	VcMac             Optional[string] `json:"vc_mac"`
	Version           *string          `json:"version,omitempty"`
	X                 *float64         `json:"x,omitempty"`
	Y                 *float64         `json:"y,omitempty"`
}
