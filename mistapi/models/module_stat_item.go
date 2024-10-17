package models

import (
	"encoding/json"
)

// ModuleStatItem represents a ModuleStatItem struct.
type ModuleStatItem struct {
	BackupVersion Optional[string] `json:"backup_version"`
	BiosVersion   Optional[string] `json:"bios_version"`
	CpldVersion   Optional[string] `json:"cpld_version"`
	// used to report all error states the device node is running into.
	// An error should always have `type` and `since` fields, and could have some other fields specific to that type.
	Errors            []ModuleStatItemErrorsItems      `json:"errors,omitempty"`
	Fans              []ModuleStatItemFansItems        `json:"fans,omitempty"`
	FpcIdx            *int                             `json:"fpc_idx,omitempty"`
	FpgaVersion       Optional[string]                 `json:"fpga_version"`
	LastSeen          Optional[float64]                `json:"last_seen"`
	Model             Optional[string]                 `json:"model"`
	OpticsCpldVersion Optional[string]                 `json:"optics_cpld_version"`
	PendingVersion    Optional[string]                 `json:"pending_version"`
	Pics              []ModuleStatItemPicsItem         `json:"pics,omitempty"`
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
	// master / backup / linecard
	VcRole               Optional[string] `json:"vc_role"`
	VcState              Optional[string] `json:"vc_state"`
	Version              Optional[string] `json:"version"`
	AdditionalProperties map[string]any   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ModuleStatItem.
// It customizes the JSON marshaling process for ModuleStatItem objects.
func (m ModuleStatItem) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(m.toMap())
}

// toMap converts the ModuleStatItem object to a map representation for JSON marshaling.
func (m ModuleStatItem) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, m.AdditionalProperties)
	if m.BackupVersion.IsValueSet() {
		if m.BackupVersion.Value() != nil {
			structMap["backup_version"] = m.BackupVersion.Value()
		} else {
			structMap["backup_version"] = nil
		}
	}
	if m.BiosVersion.IsValueSet() {
		if m.BiosVersion.Value() != nil {
			structMap["bios_version"] = m.BiosVersion.Value()
		} else {
			structMap["bios_version"] = nil
		}
	}
	if m.CpldVersion.IsValueSet() {
		if m.CpldVersion.Value() != nil {
			structMap["cpld_version"] = m.CpldVersion.Value()
		} else {
			structMap["cpld_version"] = nil
		}
	}
	if m.Errors != nil {
		structMap["errors"] = m.Errors
	}
	if m.Fans != nil {
		structMap["fans"] = m.Fans
	}
	if m.FpcIdx != nil {
		structMap["fpc_idx"] = m.FpcIdx
	}
	if m.FpgaVersion.IsValueSet() {
		if m.FpgaVersion.Value() != nil {
			structMap["fpga_version"] = m.FpgaVersion.Value()
		} else {
			structMap["fpga_version"] = nil
		}
	}
	if m.LastSeen.IsValueSet() {
		if m.LastSeen.Value() != nil {
			structMap["last_seen"] = m.LastSeen.Value()
		} else {
			structMap["last_seen"] = nil
		}
	}
	if m.Model.IsValueSet() {
		if m.Model.Value() != nil {
			structMap["model"] = m.Model.Value()
		} else {
			structMap["model"] = nil
		}
	}
	if m.OpticsCpldVersion.IsValueSet() {
		if m.OpticsCpldVersion.Value() != nil {
			structMap["optics_cpld_version"] = m.OpticsCpldVersion.Value()
		} else {
			structMap["optics_cpld_version"] = nil
		}
	}
	if m.PendingVersion.IsValueSet() {
		if m.PendingVersion.Value() != nil {
			structMap["pending_version"] = m.PendingVersion.Value()
		} else {
			structMap["pending_version"] = nil
		}
	}
	if m.Pics != nil {
		structMap["pics"] = m.Pics
	}
	if m.Poe != nil {
		structMap["poe"] = m.Poe.toMap()
	}
	if m.PoeVersion.IsValueSet() {
		if m.PoeVersion.Value() != nil {
			structMap["poe_version"] = m.PoeVersion.Value()
		} else {
			structMap["poe_version"] = nil
		}
	}
	if m.PowerCpldVersion.IsValueSet() {
		if m.PowerCpldVersion.Value() != nil {
			structMap["power_cpld_version"] = m.PowerCpldVersion.Value()
		} else {
			structMap["power_cpld_version"] = nil
		}
	}
	if m.Psus != nil {
		structMap["psus"] = m.Psus
	}
	if m.ReFpgaVersion.IsValueSet() {
		if m.ReFpgaVersion.Value() != nil {
			structMap["re_fpga_version"] = m.ReFpgaVersion.Value()
		} else {
			structMap["re_fpga_version"] = nil
		}
	}
	if m.RecoveryVersion.IsValueSet() {
		if m.RecoveryVersion.Value() != nil {
			structMap["recovery_version"] = m.RecoveryVersion.Value()
		} else {
			structMap["recovery_version"] = nil
		}
	}
	if m.Serial.IsValueSet() {
		if m.Serial.Value() != nil {
			structMap["serial"] = m.Serial.Value()
		} else {
			structMap["serial"] = nil
		}
	}
	if m.Status.IsValueSet() {
		if m.Status.Value() != nil {
			structMap["status"] = m.Status.Value()
		} else {
			structMap["status"] = nil
		}
	}
	if m.Temperatures != nil {
		structMap["temperatures"] = m.Temperatures
	}
	if m.TmcFpgaVersion.IsValueSet() {
		if m.TmcFpgaVersion.Value() != nil {
			structMap["tmc_fpga_version"] = m.TmcFpgaVersion.Value()
		} else {
			structMap["tmc_fpga_version"] = nil
		}
	}
	if m.UbootVersion.IsValueSet() {
		if m.UbootVersion.Value() != nil {
			structMap["uboot_version"] = m.UbootVersion.Value()
		} else {
			structMap["uboot_version"] = nil
		}
	}
	if m.Uptime.IsValueSet() {
		if m.Uptime.Value() != nil {
			structMap["uptime"] = m.Uptime.Value()
		} else {
			structMap["uptime"] = nil
		}
	}
	if m.VcLinks != nil {
		structMap["vc_links"] = m.VcLinks
	}
	if m.VcMode.IsValueSet() {
		if m.VcMode.Value() != nil {
			structMap["vc_mode"] = m.VcMode.Value()
		} else {
			structMap["vc_mode"] = nil
		}
	}
	if m.VcRole.IsValueSet() {
		if m.VcRole.Value() != nil {
			structMap["vc_role"] = m.VcRole.Value()
		} else {
			structMap["vc_role"] = nil
		}
	}
	if m.VcState.IsValueSet() {
		if m.VcState.Value() != nil {
			structMap["vc_state"] = m.VcState.Value()
		} else {
			structMap["vc_state"] = nil
		}
	}
	if m.Version.IsValueSet() {
		if m.Version.Value() != nil {
			structMap["version"] = m.Version.Value()
		} else {
			structMap["version"] = nil
		}
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ModuleStatItem.
// It customizes the JSON unmarshaling process for ModuleStatItem objects.
func (m *ModuleStatItem) UnmarshalJSON(input []byte) error {
	var temp tempModuleStatItem
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "backup_version", "bios_version", "cpld_version", "errors", "fans", "fpc_idx", "fpga_version", "last_seen", "model", "optics_cpld_version", "pending_version", "pics", "poe", "poe_version", "power_cpld_version", "psus", "re_fpga_version", "recovery_version", "serial", "status", "temperatures", "tmc_fpga_version", "uboot_version", "uptime", "vc_links", "vc_mode", "vc_role", "vc_state", "version")
	if err != nil {
		return err
	}

	m.AdditionalProperties = additionalProperties
	m.BackupVersion = temp.BackupVersion
	m.BiosVersion = temp.BiosVersion
	m.CpldVersion = temp.CpldVersion
	m.Errors = temp.Errors
	m.Fans = temp.Fans
	m.FpcIdx = temp.FpcIdx
	m.FpgaVersion = temp.FpgaVersion
	m.LastSeen = temp.LastSeen
	m.Model = temp.Model
	m.OpticsCpldVersion = temp.OpticsCpldVersion
	m.PendingVersion = temp.PendingVersion
	m.Pics = temp.Pics
	m.Poe = temp.Poe
	m.PoeVersion = temp.PoeVersion
	m.PowerCpldVersion = temp.PowerCpldVersion
	m.Psus = temp.Psus
	m.ReFpgaVersion = temp.ReFpgaVersion
	m.RecoveryVersion = temp.RecoveryVersion
	m.Serial = temp.Serial
	m.Status = temp.Status
	m.Temperatures = temp.Temperatures
	m.TmcFpgaVersion = temp.TmcFpgaVersion
	m.UbootVersion = temp.UbootVersion
	m.Uptime = temp.Uptime
	m.VcLinks = temp.VcLinks
	m.VcMode = temp.VcMode
	m.VcRole = temp.VcRole
	m.VcState = temp.VcState
	m.Version = temp.Version
	return nil
}

// tempModuleStatItem is a temporary struct used for validating the fields of ModuleStatItem.
type tempModuleStatItem struct {
	BackupVersion     Optional[string]                 `json:"backup_version"`
	BiosVersion       Optional[string]                 `json:"bios_version"`
	CpldVersion       Optional[string]                 `json:"cpld_version"`
	Errors            []ModuleStatItemErrorsItems      `json:"errors,omitempty"`
	Fans              []ModuleStatItemFansItems        `json:"fans,omitempty"`
	FpcIdx            *int                             `json:"fpc_idx,omitempty"`
	FpgaVersion       Optional[string]                 `json:"fpga_version"`
	LastSeen          Optional[float64]                `json:"last_seen"`
	Model             Optional[string]                 `json:"model"`
	OpticsCpldVersion Optional[string]                 `json:"optics_cpld_version"`
	PendingVersion    Optional[string]                 `json:"pending_version"`
	Pics              []ModuleStatItemPicsItem         `json:"pics,omitempty"`
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
