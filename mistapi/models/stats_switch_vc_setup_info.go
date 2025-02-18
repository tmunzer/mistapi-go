package models

import (
    "encoding/json"
    "fmt"
)

// StatsSwitchVcSetupInfo represents a StatsSwitchVcSetupInfo struct.
type StatsSwitchVcSetupInfo struct {
    ConfigType           *string                `json:"config_type,omitempty"`
    CurrentStats         *string                `json:"current_stats,omitempty"`
    ErrMissingDevIdFpc   *bool                  `json:"err_missing_dev_id_fpc,omitempty"`
    LastUpdate           *float64               `json:"last_update,omitempty"`
    RequestTime          *float64               `json:"request_time,omitempty"`
    RequestType          *string                `json:"request_type,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsSwitchVcSetupInfo,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsSwitchVcSetupInfo) String() string {
    return fmt.Sprintf(
    	"StatsSwitchVcSetupInfo[ConfigType=%v, CurrentStats=%v, ErrMissingDevIdFpc=%v, LastUpdate=%v, RequestTime=%v, RequestType=%v, AdditionalProperties=%v]",
    	s.ConfigType, s.CurrentStats, s.ErrMissingDevIdFpc, s.LastUpdate, s.RequestTime, s.RequestType, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsSwitchVcSetupInfo.
// It customizes the JSON marshaling process for StatsSwitchVcSetupInfo objects.
func (s StatsSwitchVcSetupInfo) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "config_type", "current_stats", "err_missing_dev_id_fpc", "last_update", "request_time", "request_type"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsSwitchVcSetupInfo object to a map representation for JSON marshaling.
func (s StatsSwitchVcSetupInfo) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.ConfigType != nil {
        structMap["config_type"] = s.ConfigType
    }
    if s.CurrentStats != nil {
        structMap["current_stats"] = s.CurrentStats
    }
    if s.ErrMissingDevIdFpc != nil {
        structMap["err_missing_dev_id_fpc"] = s.ErrMissingDevIdFpc
    }
    if s.LastUpdate != nil {
        structMap["last_update"] = s.LastUpdate
    }
    if s.RequestTime != nil {
        structMap["request_time"] = s.RequestTime
    }
    if s.RequestType != nil {
        structMap["request_type"] = s.RequestType
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsSwitchVcSetupInfo.
// It customizes the JSON unmarshaling process for StatsSwitchVcSetupInfo objects.
func (s *StatsSwitchVcSetupInfo) UnmarshalJSON(input []byte) error {
    var temp tempStatsSwitchVcSetupInfo
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "config_type", "current_stats", "err_missing_dev_id_fpc", "last_update", "request_time", "request_type")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.ConfigType = temp.ConfigType
    s.CurrentStats = temp.CurrentStats
    s.ErrMissingDevIdFpc = temp.ErrMissingDevIdFpc
    s.LastUpdate = temp.LastUpdate
    s.RequestTime = temp.RequestTime
    s.RequestType = temp.RequestType
    return nil
}

// tempStatsSwitchVcSetupInfo is a temporary struct used for validating the fields of StatsSwitchVcSetupInfo.
type tempStatsSwitchVcSetupInfo  struct {
    ConfigType         *string  `json:"config_type,omitempty"`
    CurrentStats       *string  `json:"current_stats,omitempty"`
    ErrMissingDevIdFpc *bool    `json:"err_missing_dev_id_fpc,omitempty"`
    LastUpdate         *float64 `json:"last_update,omitempty"`
    RequestTime        *float64 `json:"request_time,omitempty"`
    RequestType        *string  `json:"request_type,omitempty"`
}
