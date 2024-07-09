package models

import (
    "encoding/json"
)

// SwitchStats represents a SwitchStats struct.
// Switch statistics
type SwitchStats struct {
    ApRedundancy         *SwitchStatsApRedundancy     `json:"ap_redundancy,omitempty"`
    Clients              []SwitchStatsClient          `json:"clients,omitempty"`
    CpuStat              *CpuStat                     `json:"cpu_stat,omitempty"`
    // whether the switch supports packet capture
    HasPcap              *bool                        `json:"has_pcap,omitempty"`
    // hostname reported by the device
    Hostname             *string                      `json:"hostname,omitempty"`
    // Property key is the interface name
    IfStat               map[string]SwitchStatsIfStat `json:"if_stat,omitempty"`
    Ip                   *string                      `json:"ip,omitempty"`
    IpStat               *IpStat                      `json:"ip_stat,omitempty"`
    LastSeen             *int                         `json:"last_seen,omitempty"`
    // last trouble code of switch
    LastTrouble          *LastTrouble                 `json:"last_trouble,omitempty"`
    // memory usage stat (for virtual chassis, memory usage of master RE)
    Mac                  *MemoryStat                  `json:"mac,omitempty"`
    Model                *string                      `json:"model,omitempty"`
    ModuleStat           []ModuleStat                 `json:"module_stat,omitempty"`
    // device name if configured
    Name                 *string                      `json:"name,omitempty"`
    NumClients           *SwitchStatsNumClients       `json:"num_clients,omitempty"`
    Serial               *string                      `json:"serial,omitempty"`
    Status               *string                      `json:"status,omitempty"`
    Type                 *string                      `json:"type,omitempty"`
    Uptime               *float64                     `json:"uptime,omitempty"`
    Version              *string                      `json:"version,omitempty"`
    AdditionalProperties map[string]any               `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SwitchStats.
// It customizes the JSON marshaling process for SwitchStats objects.
func (s SwitchStats) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SwitchStats object to a map representation for JSON marshaling.
func (s SwitchStats) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.ApRedundancy != nil {
        structMap["ap_redundancy"] = s.ApRedundancy.toMap()
    }
    if s.Clients != nil {
        structMap["clients"] = s.Clients
    }
    if s.CpuStat != nil {
        structMap["cpu_stat"] = s.CpuStat.toMap()
    }
    if s.HasPcap != nil {
        structMap["has_pcap"] = s.HasPcap
    }
    if s.Hostname != nil {
        structMap["hostname"] = s.Hostname
    }
    if s.IfStat != nil {
        structMap["if_stat"] = s.IfStat
    }
    if s.Ip != nil {
        structMap["ip"] = s.Ip
    }
    if s.IpStat != nil {
        structMap["ip_stat"] = s.IpStat.toMap()
    }
    if s.LastSeen != nil {
        structMap["last_seen"] = s.LastSeen
    }
    if s.LastTrouble != nil {
        structMap["last_trouble"] = s.LastTrouble.toMap()
    }
    if s.Mac != nil {
        structMap["mac"] = s.Mac.toMap()
    }
    if s.Model != nil {
        structMap["model"] = s.Model
    }
    if s.ModuleStat != nil {
        structMap["module_stat"] = s.ModuleStat
    }
    if s.Name != nil {
        structMap["name"] = s.Name
    }
    if s.NumClients != nil {
        structMap["num_clients"] = s.NumClients.toMap()
    }
    if s.Serial != nil {
        structMap["serial"] = s.Serial
    }
    if s.Status != nil {
        structMap["status"] = s.Status
    }
    if s.Type != nil {
        structMap["type"] = s.Type
    }
    if s.Uptime != nil {
        structMap["uptime"] = s.Uptime
    }
    if s.Version != nil {
        structMap["version"] = s.Version
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchStats.
// It customizes the JSON unmarshaling process for SwitchStats objects.
func (s *SwitchStats) UnmarshalJSON(input []byte) error {
    var temp switchStats
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "ap_redundancy", "clients", "cpu_stat", "has_pcap", "hostname", "if_stat", "ip", "ip_stat", "last_seen", "last_trouble", "mac", "model", "module_stat", "name", "num_clients", "serial", "status", "type", "uptime", "version")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.ApRedundancy = temp.ApRedundancy
    s.Clients = temp.Clients
    s.CpuStat = temp.CpuStat
    s.HasPcap = temp.HasPcap
    s.Hostname = temp.Hostname
    s.IfStat = temp.IfStat
    s.Ip = temp.Ip
    s.IpStat = temp.IpStat
    s.LastSeen = temp.LastSeen
    s.LastTrouble = temp.LastTrouble
    s.Mac = temp.Mac
    s.Model = temp.Model
    s.ModuleStat = temp.ModuleStat
    s.Name = temp.Name
    s.NumClients = temp.NumClients
    s.Serial = temp.Serial
    s.Status = temp.Status
    s.Type = temp.Type
    s.Uptime = temp.Uptime
    s.Version = temp.Version
    return nil
}

// switchStats is a temporary struct used for validating the fields of SwitchStats.
type switchStats  struct {
    ApRedundancy *SwitchStatsApRedundancy     `json:"ap_redundancy,omitempty"`
    Clients      []SwitchStatsClient          `json:"clients,omitempty"`
    CpuStat      *CpuStat                     `json:"cpu_stat,omitempty"`
    HasPcap      *bool                        `json:"has_pcap,omitempty"`
    Hostname     *string                      `json:"hostname,omitempty"`
    IfStat       map[string]SwitchStatsIfStat `json:"if_stat,omitempty"`
    Ip           *string                      `json:"ip,omitempty"`
    IpStat       *IpStat                      `json:"ip_stat,omitempty"`
    LastSeen     *int                         `json:"last_seen,omitempty"`
    LastTrouble  *LastTrouble                 `json:"last_trouble,omitempty"`
    Mac          *MemoryStat                  `json:"mac,omitempty"`
    Model        *string                      `json:"model,omitempty"`
    ModuleStat   []ModuleStat                 `json:"module_stat,omitempty"`
    Name         *string                      `json:"name,omitempty"`
    NumClients   *SwitchStatsNumClients       `json:"num_clients,omitempty"`
    Serial       *string                      `json:"serial,omitempty"`
    Status       *string                      `json:"status,omitempty"`
    Type         *string                      `json:"type,omitempty"`
    Uptime       *float64                     `json:"uptime,omitempty"`
    Version      *string                      `json:"version,omitempty"`
}
