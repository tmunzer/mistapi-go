package models

import (
    "encoding/json"
)

// ModuleStat represents a ModuleStat struct.
type ModuleStat struct {
    BackupVersion        *string                       `json:"backup_version,omitempty"`
    BiosVersion          *string                       `json:"bios_version,omitempty"`
    // used to report all error states the device node is running into.
    // An error should always have `type` and `since` fields, and could have some other fields specific to that type.
    Errors               []ModuleStatErrorsItems       `json:"errors,omitempty"`
    Fans                 []ModuleStatFansItems         `json:"fans,omitempty"`
    Model                *string                       `json:"model,omitempty"`
    Poe                  *ModuleStatPoe                `json:"poe,omitempty"`
    Psus                 []ModuleStatPsusItems         `json:"psus,omitempty"`
    Serial               *string                       `json:"serial,omitempty"`
    Temperatures         []ModuleStatTemperaturesItems `json:"temperatures,omitempty"`
    VcLinks              []ModuleStatVcLinksItems      `json:"vc_links,omitempty"`
    VcMode               *string                       `json:"vc_mode,omitempty"`
    // master / backup / linecard
    VcRole               *string                       `json:"vc_role,omitempty"`
    VcState              *string                       `json:"vc_state,omitempty"`
    AdditionalProperties map[string]any                `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ModuleStat.
// It customizes the JSON marshaling process for ModuleStat objects.
func (m ModuleStat) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the ModuleStat object to a map representation for JSON marshaling.
func (m ModuleStat) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
    if m.BackupVersion != nil {
        structMap["backup_version"] = m.BackupVersion
    }
    if m.BiosVersion != nil {
        structMap["bios_version"] = m.BiosVersion
    }
    if m.Errors != nil {
        structMap["errors"] = m.Errors
    }
    if m.Fans != nil {
        structMap["fans"] = m.Fans
    }
    if m.Model != nil {
        structMap["model"] = m.Model
    }
    if m.Poe != nil {
        structMap["poe"] = m.Poe.toMap()
    }
    if m.Psus != nil {
        structMap["psus"] = m.Psus
    }
    if m.Serial != nil {
        structMap["serial"] = m.Serial
    }
    if m.Temperatures != nil {
        structMap["temperatures"] = m.Temperatures
    }
    if m.VcLinks != nil {
        structMap["vc_links"] = m.VcLinks
    }
    if m.VcMode != nil {
        structMap["vc_mode"] = m.VcMode
    }
    if m.VcRole != nil {
        structMap["vc_role"] = m.VcRole
    }
    if m.VcState != nil {
        structMap["vc_state"] = m.VcState
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ModuleStat.
// It customizes the JSON unmarshaling process for ModuleStat objects.
func (m *ModuleStat) UnmarshalJSON(input []byte) error {
    var temp moduleStat
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "backup_version", "bios_version", "errors", "fans", "model", "poe", "psus", "serial", "temperatures", "vc_links", "vc_mode", "vc_role", "vc_state")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.BackupVersion = temp.BackupVersion
    m.BiosVersion = temp.BiosVersion
    m.Errors = temp.Errors
    m.Fans = temp.Fans
    m.Model = temp.Model
    m.Poe = temp.Poe
    m.Psus = temp.Psus
    m.Serial = temp.Serial
    m.Temperatures = temp.Temperatures
    m.VcLinks = temp.VcLinks
    m.VcMode = temp.VcMode
    m.VcRole = temp.VcRole
    m.VcState = temp.VcState
    return nil
}

// moduleStat is a temporary struct used for validating the fields of ModuleStat.
type moduleStat  struct {
    BackupVersion *string                       `json:"backup_version,omitempty"`
    BiosVersion   *string                       `json:"bios_version,omitempty"`
    Errors        []ModuleStatErrorsItems       `json:"errors,omitempty"`
    Fans          []ModuleStatFansItems         `json:"fans,omitempty"`
    Model         *string                       `json:"model,omitempty"`
    Poe           *ModuleStatPoe                `json:"poe,omitempty"`
    Psus          []ModuleStatPsusItems         `json:"psus,omitempty"`
    Serial        *string                       `json:"serial,omitempty"`
    Temperatures  []ModuleStatTemperaturesItems `json:"temperatures,omitempty"`
    VcLinks       []ModuleStatVcLinksItems      `json:"vc_links,omitempty"`
    VcMode        *string                       `json:"vc_mode,omitempty"`
    VcRole        *string                       `json:"vc_role,omitempty"`
    VcState       *string                       `json:"vc_state,omitempty"`
}
