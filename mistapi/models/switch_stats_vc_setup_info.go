package models

import (
    "encoding/json"
)

// SwitchStatsVcSetupInfo represents a SwitchStatsVcSetupInfo struct.
type SwitchStatsVcSetupInfo struct {
    ConfigType           *string        `json:"config_type,omitempty"`
    ErrMissingDevIdFpc   *bool          `json:"err_missing_dev_id_fpc,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SwitchStatsVcSetupInfo.
// It customizes the JSON marshaling process for SwitchStatsVcSetupInfo objects.
func (s SwitchStatsVcSetupInfo) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SwitchStatsVcSetupInfo object to a map representation for JSON marshaling.
func (s SwitchStatsVcSetupInfo) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.ConfigType != nil {
        structMap["config_type"] = s.ConfigType
    }
    if s.ErrMissingDevIdFpc != nil {
        structMap["err_missing_dev_id_fpc"] = s.ErrMissingDevIdFpc
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchStatsVcSetupInfo.
// It customizes the JSON unmarshaling process for SwitchStatsVcSetupInfo objects.
func (s *SwitchStatsVcSetupInfo) UnmarshalJSON(input []byte) error {
    var temp tempSwitchStatsVcSetupInfo
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "config_type", "err_missing_dev_id_fpc")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.ConfigType = temp.ConfigType
    s.ErrMissingDevIdFpc = temp.ErrMissingDevIdFpc
    return nil
}

// tempSwitchStatsVcSetupInfo is a temporary struct used for validating the fields of SwitchStatsVcSetupInfo.
type tempSwitchStatsVcSetupInfo  struct {
    ConfigType         *string `json:"config_type,omitempty"`
    ErrMissingDevIdFpc *bool   `json:"err_missing_dev_id_fpc,omitempty"`
}
