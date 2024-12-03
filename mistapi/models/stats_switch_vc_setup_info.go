package models

import (
    "encoding/json"
)

// StatsSwitchVcSetupInfo represents a StatsSwitchVcSetupInfo struct.
type StatsSwitchVcSetupInfo struct {
    ConfigType           *string                `json:"config_type,omitempty"`
    ErrMissingDevIdFpc   *bool                  `json:"err_missing_dev_id_fpc,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for StatsSwitchVcSetupInfo.
// It customizes the JSON marshaling process for StatsSwitchVcSetupInfo objects.
func (s StatsSwitchVcSetupInfo) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "config_type", "err_missing_dev_id_fpc"); err != nil {
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
    if s.ErrMissingDevIdFpc != nil {
        structMap["err_missing_dev_id_fpc"] = s.ErrMissingDevIdFpc
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "config_type", "err_missing_dev_id_fpc")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.ConfigType = temp.ConfigType
    s.ErrMissingDevIdFpc = temp.ErrMissingDevIdFpc
    return nil
}

// tempStatsSwitchVcSetupInfo is a temporary struct used for validating the fields of StatsSwitchVcSetupInfo.
type tempStatsSwitchVcSetupInfo  struct {
    ConfigType         *string `json:"config_type,omitempty"`
    ErrMissingDevIdFpc *bool   `json:"err_missing_dev_id_fpc,omitempty"`
}
