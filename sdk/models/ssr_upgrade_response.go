package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// SsrUpgradeResponse represents a SsrUpgradeResponse struct.
type SsrUpgradeResponse struct {
    Channel              string                   `json:"channel"`
    Counts               SsrUpgradeResponseCounts `json:"counts"`
    DeviceType           string                   `json:"device_type"`
    Id                   string                   `json:"id"`
    Status               string                   `json:"status"`
    Strategy             string                   `json:"strategy"`
    Versions             map[string]string        `json:"versions"`
    AdditionalProperties map[string]any           `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SsrUpgradeResponse.
// It customizes the JSON marshaling process for SsrUpgradeResponse objects.
func (s SsrUpgradeResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SsrUpgradeResponse object to a map representation for JSON marshaling.
func (s SsrUpgradeResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["channel"] = s.Channel
    structMap["counts"] = s.Counts.toMap()
    structMap["device_type"] = s.DeviceType
    structMap["id"] = s.Id
    structMap["status"] = s.Status
    structMap["strategy"] = s.Strategy
    structMap["versions"] = s.Versions
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SsrUpgradeResponse.
// It customizes the JSON unmarshaling process for SsrUpgradeResponse objects.
func (s *SsrUpgradeResponse) UnmarshalJSON(input []byte) error {
    var temp ssrUpgradeResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "channel", "counts", "device_type", "id", "status", "strategy", "versions")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Channel = *temp.Channel
    s.Counts = *temp.Counts
    s.DeviceType = *temp.DeviceType
    s.Id = *temp.Id
    s.Status = *temp.Status
    s.Strategy = *temp.Strategy
    s.Versions = *temp.Versions
    return nil
}

// ssrUpgradeResponse is a temporary struct used for validating the fields of SsrUpgradeResponse.
type ssrUpgradeResponse  struct {
    Channel    *string                   `json:"channel"`
    Counts     *SsrUpgradeResponseCounts `json:"counts"`
    DeviceType *string                   `json:"device_type"`
    Id         *string                   `json:"id"`
    Status     *string                   `json:"status"`
    Strategy   *string                   `json:"strategy"`
    Versions   *map[string]string        `json:"versions"`
}

func (s *ssrUpgradeResponse) validate() error {
    var errs []string
    if s.Channel == nil {
        errs = append(errs, "required field `channel` is missing for type `Ssr_Upgrade_Response`")
    }
    if s.Counts == nil {
        errs = append(errs, "required field `counts` is missing for type `Ssr_Upgrade_Response`")
    }
    if s.DeviceType == nil {
        errs = append(errs, "required field `device_type` is missing for type `Ssr_Upgrade_Response`")
    }
    if s.Id == nil {
        errs = append(errs, "required field `id` is missing for type `Ssr_Upgrade_Response`")
    }
    if s.Status == nil {
        errs = append(errs, "required field `status` is missing for type `Ssr_Upgrade_Response`")
    }
    if s.Strategy == nil {
        errs = append(errs, "required field `strategy` is missing for type `Ssr_Upgrade_Response`")
    }
    if s.Versions == nil {
        errs = append(errs, "required field `versions` is missing for type `Ssr_Upgrade_Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
