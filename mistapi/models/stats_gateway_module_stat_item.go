// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// StatsGatewayModuleStatItem represents a StatsGatewayModuleStatItem struct.
type StatsGatewayModuleStatItem struct {
    BackupVersion        Optional[string]                 `json:"backup_version"`
    BiosVersion          Optional[string]                 `json:"bios_version"`
    CpldVersion          Optional[string]                 `json:"cpld_version"`
    Fans                 []ModuleStatItemFansItems        `json:"fans,omitempty"`
    FpgaVersion          Optional[string]                 `json:"fpga_version"`
    // Last seen timestamp
    LastSeen             Optional[float64]                `json:"last_seen"`
    Locating             *bool                            `json:"locating,omitempty"`
    Mac                  *string                          `json:"mac,omitempty"`
    Model                Optional[string]                 `json:"model"`
    OpticsCpldVersion    Optional[string]                 `json:"optics_cpld_version"`
    PendingVersion       Optional[string]                 `json:"pending_version"`
    Poe                  *ModuleStatItemPoe               `json:"poe,omitempty"`
    PoeVersion           Optional[string]                 `json:"poe_version"`
    PowerCpldVersion     Optional[string]                 `json:"power_cpld_version"`
    Psus                 []ModuleStatItemPsusItem         `json:"psus,omitempty"`
    ReFpgaVersion        Optional[string]                 `json:"re_fpga_version"`
    RecoveryVersion      Optional[string]                 `json:"recovery_version"`
    Serial               Optional[string]                 `json:"serial"`
    Status               Optional[string]                 `json:"status"`
    Temperatures         []ModuleStatItemTemperaturesItem `json:"temperatures,omitempty"`
    TmcFpgaVersion       Optional[string]                 `json:"tmc_fpga_version"`
    UbootVersion         Optional[string]                 `json:"uboot_version"`
    Uptime               Optional[int]                    `json:"uptime"`
    VcLinks              []ModuleStatItemVcLinksItem      `json:"vc_links,omitempty"`
    VcMode               Optional[string]                 `json:"vc_mode"`
    // enum: `master`, `backup`, `linecard`
    VcRole               Optional[string]                 `json:"vc_role"`
    VcState              Optional[string]                 `json:"vc_state"`
    Version              Optional[string]                 `json:"version"`
    AdditionalProperties map[string]interface{}           `json:"_"`
}

// String implements the fmt.Stringer interface for StatsGatewayModuleStatItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsGatewayModuleStatItem) String() string {
    return fmt.Sprintf(
    	"StatsGatewayModuleStatItem[BackupVersion=%v, BiosVersion=%v, CpldVersion=%v, Fans=%v, FpgaVersion=%v, LastSeen=%v, Locating=%v, Mac=%v, Model=%v, OpticsCpldVersion=%v, PendingVersion=%v, Poe=%v, PoeVersion=%v, PowerCpldVersion=%v, Psus=%v, ReFpgaVersion=%v, RecoveryVersion=%v, Serial=%v, Status=%v, Temperatures=%v, TmcFpgaVersion=%v, UbootVersion=%v, Uptime=%v, VcLinks=%v, VcMode=%v, VcRole=%v, VcState=%v, Version=%v, AdditionalProperties=%v]",
    	s.BackupVersion, s.BiosVersion, s.CpldVersion, s.Fans, s.FpgaVersion, s.LastSeen, s.Locating, s.Mac, s.Model, s.OpticsCpldVersion, s.PendingVersion, s.Poe, s.PoeVersion, s.PowerCpldVersion, s.Psus, s.ReFpgaVersion, s.RecoveryVersion, s.Serial, s.Status, s.Temperatures, s.TmcFpgaVersion, s.UbootVersion, s.Uptime, s.VcLinks, s.VcMode, s.VcRole, s.VcState, s.Version, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsGatewayModuleStatItem.
// It customizes the JSON marshaling process for StatsGatewayModuleStatItem objects.
func (s StatsGatewayModuleStatItem) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "backup_version", "bios_version", "cpld_version", "fans", "fpga_version", "last_seen", "locating", "mac", "model", "optics_cpld_version", "pending_version", "poe", "poe_version", "power_cpld_version", "psus", "re_fpga_version", "recovery_version", "serial", "status", "temperatures", "tmc_fpga_version", "uboot_version", "uptime", "vc_links", "vc_mode", "vc_role", "vc_state", "version"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsGatewayModuleStatItem object to a map representation for JSON marshaling.
func (s StatsGatewayModuleStatItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.BackupVersion.IsValueSet() {
        if s.BackupVersion.Value() != nil {
            structMap["backup_version"] = s.BackupVersion.Value()
        } else {
            structMap["backup_version"] = nil
        }
    }
    if s.BiosVersion.IsValueSet() {
        if s.BiosVersion.Value() != nil {
            structMap["bios_version"] = s.BiosVersion.Value()
        } else {
            structMap["bios_version"] = nil
        }
    }
    if s.CpldVersion.IsValueSet() {
        if s.CpldVersion.Value() != nil {
            structMap["cpld_version"] = s.CpldVersion.Value()
        } else {
            structMap["cpld_version"] = nil
        }
    }
    if s.Fans != nil {
        structMap["fans"] = s.Fans
    }
    if s.FpgaVersion.IsValueSet() {
        if s.FpgaVersion.Value() != nil {
            structMap["fpga_version"] = s.FpgaVersion.Value()
        } else {
            structMap["fpga_version"] = nil
        }
    }
    if s.LastSeen.IsValueSet() {
        if s.LastSeen.Value() != nil {
            structMap["last_seen"] = s.LastSeen.Value()
        } else {
            structMap["last_seen"] = nil
        }
    }
    if s.Locating != nil {
        structMap["locating"] = s.Locating
    }
    if s.Mac != nil {
        structMap["mac"] = s.Mac
    }
    if s.Model.IsValueSet() {
        if s.Model.Value() != nil {
            structMap["model"] = s.Model.Value()
        } else {
            structMap["model"] = nil
        }
    }
    if s.OpticsCpldVersion.IsValueSet() {
        if s.OpticsCpldVersion.Value() != nil {
            structMap["optics_cpld_version"] = s.OpticsCpldVersion.Value()
        } else {
            structMap["optics_cpld_version"] = nil
        }
    }
    if s.PendingVersion.IsValueSet() {
        if s.PendingVersion.Value() != nil {
            structMap["pending_version"] = s.PendingVersion.Value()
        } else {
            structMap["pending_version"] = nil
        }
    }
    if s.Poe != nil {
        structMap["poe"] = s.Poe.toMap()
    }
    if s.PoeVersion.IsValueSet() {
        if s.PoeVersion.Value() != nil {
            structMap["poe_version"] = s.PoeVersion.Value()
        } else {
            structMap["poe_version"] = nil
        }
    }
    if s.PowerCpldVersion.IsValueSet() {
        if s.PowerCpldVersion.Value() != nil {
            structMap["power_cpld_version"] = s.PowerCpldVersion.Value()
        } else {
            structMap["power_cpld_version"] = nil
        }
    }
    if s.Psus != nil {
        structMap["psus"] = s.Psus
    }
    if s.ReFpgaVersion.IsValueSet() {
        if s.ReFpgaVersion.Value() != nil {
            structMap["re_fpga_version"] = s.ReFpgaVersion.Value()
        } else {
            structMap["re_fpga_version"] = nil
        }
    }
    if s.RecoveryVersion.IsValueSet() {
        if s.RecoveryVersion.Value() != nil {
            structMap["recovery_version"] = s.RecoveryVersion.Value()
        } else {
            structMap["recovery_version"] = nil
        }
    }
    if s.Serial.IsValueSet() {
        if s.Serial.Value() != nil {
            structMap["serial"] = s.Serial.Value()
        } else {
            structMap["serial"] = nil
        }
    }
    if s.Status.IsValueSet() {
        if s.Status.Value() != nil {
            structMap["status"] = s.Status.Value()
        } else {
            structMap["status"] = nil
        }
    }
    if s.Temperatures != nil {
        structMap["temperatures"] = s.Temperatures
    }
    if s.TmcFpgaVersion.IsValueSet() {
        if s.TmcFpgaVersion.Value() != nil {
            structMap["tmc_fpga_version"] = s.TmcFpgaVersion.Value()
        } else {
            structMap["tmc_fpga_version"] = nil
        }
    }
    if s.UbootVersion.IsValueSet() {
        if s.UbootVersion.Value() != nil {
            structMap["uboot_version"] = s.UbootVersion.Value()
        } else {
            structMap["uboot_version"] = nil
        }
    }
    if s.Uptime.IsValueSet() {
        if s.Uptime.Value() != nil {
            structMap["uptime"] = s.Uptime.Value()
        } else {
            structMap["uptime"] = nil
        }
    }
    if s.VcLinks != nil {
        structMap["vc_links"] = s.VcLinks
    }
    if s.VcMode.IsValueSet() {
        if s.VcMode.Value() != nil {
            structMap["vc_mode"] = s.VcMode.Value()
        } else {
            structMap["vc_mode"] = nil
        }
    }
    if s.VcRole.IsValueSet() {
        if s.VcRole.Value() != nil {
            structMap["vc_role"] = s.VcRole.Value()
        } else {
            structMap["vc_role"] = nil
        }
    }
    if s.VcState.IsValueSet() {
        if s.VcState.Value() != nil {
            structMap["vc_state"] = s.VcState.Value()
        } else {
            structMap["vc_state"] = nil
        }
    }
    if s.Version.IsValueSet() {
        if s.Version.Value() != nil {
            structMap["version"] = s.Version.Value()
        } else {
            structMap["version"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsGatewayModuleStatItem.
// It customizes the JSON unmarshaling process for StatsGatewayModuleStatItem objects.
func (s *StatsGatewayModuleStatItem) UnmarshalJSON(input []byte) error {
    var temp tempStatsGatewayModuleStatItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "backup_version", "bios_version", "cpld_version", "fans", "fpga_version", "last_seen", "locating", "mac", "model", "optics_cpld_version", "pending_version", "poe", "poe_version", "power_cpld_version", "psus", "re_fpga_version", "recovery_version", "serial", "status", "temperatures", "tmc_fpga_version", "uboot_version", "uptime", "vc_links", "vc_mode", "vc_role", "vc_state", "version")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.BackupVersion = temp.BackupVersion
    s.BiosVersion = temp.BiosVersion
    s.CpldVersion = temp.CpldVersion
    s.Fans = temp.Fans
    s.FpgaVersion = temp.FpgaVersion
    s.LastSeen = temp.LastSeen
    s.Locating = temp.Locating
    s.Mac = temp.Mac
    s.Model = temp.Model
    s.OpticsCpldVersion = temp.OpticsCpldVersion
    s.PendingVersion = temp.PendingVersion
    s.Poe = temp.Poe
    s.PoeVersion = temp.PoeVersion
    s.PowerCpldVersion = temp.PowerCpldVersion
    s.Psus = temp.Psus
    s.ReFpgaVersion = temp.ReFpgaVersion
    s.RecoveryVersion = temp.RecoveryVersion
    s.Serial = temp.Serial
    s.Status = temp.Status
    s.Temperatures = temp.Temperatures
    s.TmcFpgaVersion = temp.TmcFpgaVersion
    s.UbootVersion = temp.UbootVersion
    s.Uptime = temp.Uptime
    s.VcLinks = temp.VcLinks
    s.VcMode = temp.VcMode
    s.VcRole = temp.VcRole
    s.VcState = temp.VcState
    s.Version = temp.Version
    return nil
}

// tempStatsGatewayModuleStatItem is a temporary struct used for validating the fields of StatsGatewayModuleStatItem.
type tempStatsGatewayModuleStatItem  struct {
    BackupVersion     Optional[string]                 `json:"backup_version"`
    BiosVersion       Optional[string]                 `json:"bios_version"`
    CpldVersion       Optional[string]                 `json:"cpld_version"`
    Fans              []ModuleStatItemFansItems        `json:"fans,omitempty"`
    FpgaVersion       Optional[string]                 `json:"fpga_version"`
    LastSeen          Optional[float64]                `json:"last_seen"`
    Locating          *bool                            `json:"locating,omitempty"`
    Mac               *string                          `json:"mac,omitempty"`
    Model             Optional[string]                 `json:"model"`
    OpticsCpldVersion Optional[string]                 `json:"optics_cpld_version"`
    PendingVersion    Optional[string]                 `json:"pending_version"`
    Poe               *ModuleStatItemPoe               `json:"poe,omitempty"`
    PoeVersion        Optional[string]                 `json:"poe_version"`
    PowerCpldVersion  Optional[string]                 `json:"power_cpld_version"`
    Psus              []ModuleStatItemPsusItem         `json:"psus,omitempty"`
    ReFpgaVersion     Optional[string]                 `json:"re_fpga_version"`
    RecoveryVersion   Optional[string]                 `json:"recovery_version"`
    Serial            Optional[string]                 `json:"serial"`
    Status            Optional[string]                 `json:"status"`
    Temperatures      []ModuleStatItemTemperaturesItem `json:"temperatures,omitempty"`
    TmcFpgaVersion    Optional[string]                 `json:"tmc_fpga_version"`
    UbootVersion      Optional[string]                 `json:"uboot_version"`
    Uptime            Optional[int]                    `json:"uptime"`
    VcLinks           []ModuleStatItemVcLinksItem      `json:"vc_links,omitempty"`
    VcMode            Optional[string]                 `json:"vc_mode"`
    VcRole            Optional[string]                 `json:"vc_role"`
    VcState           Optional[string]                 `json:"vc_state"`
    Version           Optional[string]                 `json:"version"`
}
