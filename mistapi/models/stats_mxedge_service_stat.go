package models

import (
    "encoding/json"
    "fmt"
)

// StatsMxedgeServiceStat represents a StatsMxedgeServiceStat struct.
type StatsMxedgeServiceStat struct {
    // external IP from ep-terminator’s point of view. valid only for service having its own cloud connection
    ExtIp                *string                `json:"ext_ip,omitempty"`
    // timestamp when the last stats is seen (cloud unix time, in second). valid only for service having its own stats or whole system (last among last_seen of all services)
    LastSeen             *float64               `json:"last_seen,omitempty"`
    // package/service installation state.
    PackageState         *string                `json:"package_state,omitempty"`
    // package/service installation state.
    PackageVersion       *string                `json:"package_version,omitempty"`
    // service running state.
    RunningState         *string                `json:"running_state,omitempty"`
    // service uptime.
    Uptime               *int                   `json:"uptime,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsMxedgeServiceStat,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsMxedgeServiceStat) String() string {
    return fmt.Sprintf(
    	"StatsMxedgeServiceStat[ExtIp=%v, LastSeen=%v, PackageState=%v, PackageVersion=%v, RunningState=%v, Uptime=%v, AdditionalProperties=%v]",
    	s.ExtIp, s.LastSeen, s.PackageState, s.PackageVersion, s.RunningState, s.Uptime, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsMxedgeServiceStat.
// It customizes the JSON marshaling process for StatsMxedgeServiceStat objects.
func (s StatsMxedgeServiceStat) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "ext_ip", "last_seen", "package_state", "package_version", "running_state", "uptime"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsMxedgeServiceStat object to a map representation for JSON marshaling.
func (s StatsMxedgeServiceStat) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.ExtIp != nil {
        structMap["ext_ip"] = s.ExtIp
    }
    if s.LastSeen != nil {
        structMap["last_seen"] = s.LastSeen
    }
    if s.PackageState != nil {
        structMap["package_state"] = s.PackageState
    }
    if s.PackageVersion != nil {
        structMap["package_version"] = s.PackageVersion
    }
    if s.RunningState != nil {
        structMap["running_state"] = s.RunningState
    }
    if s.Uptime != nil {
        structMap["uptime"] = s.Uptime
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsMxedgeServiceStat.
// It customizes the JSON unmarshaling process for StatsMxedgeServiceStat objects.
func (s *StatsMxedgeServiceStat) UnmarshalJSON(input []byte) error {
    var temp tempStatsMxedgeServiceStat
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ext_ip", "last_seen", "package_state", "package_version", "running_state", "uptime")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.ExtIp = temp.ExtIp
    s.LastSeen = temp.LastSeen
    s.PackageState = temp.PackageState
    s.PackageVersion = temp.PackageVersion
    s.RunningState = temp.RunningState
    s.Uptime = temp.Uptime
    return nil
}

// tempStatsMxedgeServiceStat is a temporary struct used for validating the fields of StatsMxedgeServiceStat.
type tempStatsMxedgeServiceStat  struct {
    ExtIp          *string  `json:"ext_ip,omitempty"`
    LastSeen       *float64 `json:"last_seen,omitempty"`
    PackageState   *string  `json:"package_state,omitempty"`
    PackageVersion *string  `json:"package_version,omitempty"`
    RunningState   *string  `json:"running_state,omitempty"`
    Uptime         *int     `json:"uptime,omitempty"`
}
