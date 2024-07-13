package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ResponseConfigHistorySearchItem represents a ResponseConfigHistorySearchItem struct.
type ResponseConfigHistorySearchItem struct {
    Channel24            int                                    `json:"channel_24"`
    Channel5             int                                    `json:"channel_5"`
    RadioMacs            []string                               `json:"radio_macs,omitempty"`
    Radios               []ResponseConfigHistorySearchItemRadio `json:"radios,omitempty"`
    SecpolicyViolated    bool                                   `json:"secpolicy_violated"`
    Ssids                []string                               `json:"ssids,omitempty"`
    Ssids24              []string                               `json:"ssids_24,omitempty"`
    Ssids5               []string                               `json:"ssids_5,omitempty"`
    Timestamp            float64                                `json:"timestamp"`
    Version              string                                 `json:"version"`
    Wlans                []ResponseConfigHistorySearchItemWlan  `json:"wlans,omitempty"`
    AdditionalProperties map[string]any                         `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseConfigHistorySearchItem.
// It customizes the JSON marshaling process for ResponseConfigHistorySearchItem objects.
func (r ResponseConfigHistorySearchItem) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseConfigHistorySearchItem object to a map representation for JSON marshaling.
func (r ResponseConfigHistorySearchItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["channel_24"] = r.Channel24
    structMap["channel_5"] = r.Channel5
    if r.RadioMacs != nil {
        structMap["radio_macs"] = r.RadioMacs
    }
    if r.Radios != nil {
        structMap["radios"] = r.Radios
    }
    structMap["secpolicy_violated"] = r.SecpolicyViolated
    if r.Ssids != nil {
        structMap["ssids"] = r.Ssids
    }
    if r.Ssids24 != nil {
        structMap["ssids_24"] = r.Ssids24
    }
    if r.Ssids5 != nil {
        structMap["ssids_5"] = r.Ssids5
    }
    structMap["timestamp"] = r.Timestamp
    structMap["version"] = r.Version
    if r.Wlans != nil {
        structMap["wlans"] = r.Wlans
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseConfigHistorySearchItem.
// It customizes the JSON unmarshaling process for ResponseConfigHistorySearchItem objects.
func (r *ResponseConfigHistorySearchItem) UnmarshalJSON(input []byte) error {
    var temp responseConfigHistorySearchItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "channel_24", "channel_5", "radio_macs", "radios", "secpolicy_violated", "ssids", "ssids_24", "ssids_5", "timestamp", "version", "wlans")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.Channel24 = *temp.Channel24
    r.Channel5 = *temp.Channel5
    r.RadioMacs = temp.RadioMacs
    r.Radios = temp.Radios
    r.SecpolicyViolated = *temp.SecpolicyViolated
    r.Ssids = temp.Ssids
    r.Ssids24 = temp.Ssids24
    r.Ssids5 = temp.Ssids5
    r.Timestamp = *temp.Timestamp
    r.Version = *temp.Version
    r.Wlans = temp.Wlans
    return nil
}

// responseConfigHistorySearchItem is a temporary struct used for validating the fields of ResponseConfigHistorySearchItem.
type responseConfigHistorySearchItem  struct {
    Channel24         *int                                   `json:"channel_24"`
    Channel5          *int                                   `json:"channel_5"`
    RadioMacs         []string                               `json:"radio_macs,omitempty"`
    Radios            []ResponseConfigHistorySearchItemRadio `json:"radios,omitempty"`
    SecpolicyViolated *bool                                  `json:"secpolicy_violated"`
    Ssids             []string                               `json:"ssids,omitempty"`
    Ssids24           []string                               `json:"ssids_24,omitempty"`
    Ssids5            []string                               `json:"ssids_5,omitempty"`
    Timestamp         *float64                               `json:"timestamp"`
    Version           *string                                `json:"version"`
    Wlans             []ResponseConfigHistorySearchItemWlan  `json:"wlans,omitempty"`
}

func (r *responseConfigHistorySearchItem) validate() error {
    var errs []string
    if r.Channel24 == nil {
        errs = append(errs, "required field `channel_24` is missing for type `Response_Config_History_Search_Item`")
    }
    if r.Channel5 == nil {
        errs = append(errs, "required field `channel_5` is missing for type `Response_Config_History_Search_Item`")
    }
    if r.SecpolicyViolated == nil {
        errs = append(errs, "required field `secpolicy_violated` is missing for type `Response_Config_History_Search_Item`")
    }
    if r.Timestamp == nil {
        errs = append(errs, "required field `timestamp` is missing for type `Response_Config_History_Search_Item`")
    }
    if r.Version == nil {
        errs = append(errs, "required field `version` is missing for type `Response_Config_History_Search_Item`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}