package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// InsightRogueAp represents a InsightRogueAp struct.
type InsightRogueAp struct {
    // MAC of the device that had strongest signal strength for ssid/bssid pair
    ApMac                string                 `json:"ap_mac"`
    // Average signal strength of ap_mac for ssid/bssid pair
    AvgRssi              float64                `json:"avg_rssi"`
    // BSSID of the network detected as threat
    Bssid                string                 `json:"bssid"`
    // Channel over which ap_mac heard ssid/bssid pair
    Channel              string                 `json:"channel"`
    // X position relative to the reporting AP (`ap_mac`)
    DeltaX               *float64               `json:"delta_x,omitempty"`
    // Y position relative to the reporting AP (`ap_mac`)
    DeltaY               *float64               `json:"delta_y,omitempty"`
    // Num of aps that heard the ssid/bssid pair
    NumAps               int                    `json:"num_aps"`
    // Whether the reporting AP see a wireless client (on LAN) connecting to it
    SeenOnLan            *bool                  `json:"seen_on_lan,omitempty"`
    // SSID of the network detected as threat
    Ssid                 *string                `json:"ssid,omitempty"`
    // Represents number of times the pair was heard in the interval. Each count roughly corresponds to a minute.
    TimesHeard           *int                   `json:"times_heard,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for InsightRogueAp,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (i InsightRogueAp) String() string {
    return fmt.Sprintf(
    	"InsightRogueAp[ApMac=%v, AvgRssi=%v, Bssid=%v, Channel=%v, DeltaX=%v, DeltaY=%v, NumAps=%v, SeenOnLan=%v, Ssid=%v, TimesHeard=%v, AdditionalProperties=%v]",
    	i.ApMac, i.AvgRssi, i.Bssid, i.Channel, i.DeltaX, i.DeltaY, i.NumAps, i.SeenOnLan, i.Ssid, i.TimesHeard, i.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for InsightRogueAp.
// It customizes the JSON marshaling process for InsightRogueAp objects.
func (i InsightRogueAp) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(i.AdditionalProperties,
        "ap_mac", "avg_rssi", "bssid", "channel", "delta_x", "delta_y", "num_aps", "seen_on_lan", "ssid", "times_heard"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(i.toMap())
}

// toMap converts the InsightRogueAp object to a map representation for JSON marshaling.
func (i InsightRogueAp) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, i.AdditionalProperties)
    structMap["ap_mac"] = i.ApMac
    structMap["avg_rssi"] = i.AvgRssi
    structMap["bssid"] = i.Bssid
    structMap["channel"] = i.Channel
    if i.DeltaX != nil {
        structMap["delta_x"] = i.DeltaX
    }
    if i.DeltaY != nil {
        structMap["delta_y"] = i.DeltaY
    }
    structMap["num_aps"] = i.NumAps
    if i.SeenOnLan != nil {
        structMap["seen_on_lan"] = i.SeenOnLan
    }
    if i.Ssid != nil {
        structMap["ssid"] = i.Ssid
    }
    if i.TimesHeard != nil {
        structMap["times_heard"] = i.TimesHeard
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InsightRogueAp.
// It customizes the JSON unmarshaling process for InsightRogueAp objects.
func (i *InsightRogueAp) UnmarshalJSON(input []byte) error {
    var temp tempInsightRogueAp
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ap_mac", "avg_rssi", "bssid", "channel", "delta_x", "delta_y", "num_aps", "seen_on_lan", "ssid", "times_heard")
    if err != nil {
    	return err
    }
    i.AdditionalProperties = additionalProperties
    
    i.ApMac = *temp.ApMac
    i.AvgRssi = *temp.AvgRssi
    i.Bssid = *temp.Bssid
    i.Channel = *temp.Channel
    i.DeltaX = temp.DeltaX
    i.DeltaY = temp.DeltaY
    i.NumAps = *temp.NumAps
    i.SeenOnLan = temp.SeenOnLan
    i.Ssid = temp.Ssid
    i.TimesHeard = temp.TimesHeard
    return nil
}

// tempInsightRogueAp is a temporary struct used for validating the fields of InsightRogueAp.
type tempInsightRogueAp  struct {
    ApMac      *string  `json:"ap_mac"`
    AvgRssi    *float64 `json:"avg_rssi"`
    Bssid      *string  `json:"bssid"`
    Channel    *string  `json:"channel"`
    DeltaX     *float64 `json:"delta_x,omitempty"`
    DeltaY     *float64 `json:"delta_y,omitempty"`
    NumAps     *int     `json:"num_aps"`
    SeenOnLan  *bool    `json:"seen_on_lan,omitempty"`
    Ssid       *string  `json:"ssid,omitempty"`
    TimesHeard *int     `json:"times_heard,omitempty"`
}

func (i *tempInsightRogueAp) validate() error {
    var errs []string
    if i.ApMac == nil {
        errs = append(errs, "required field `ap_mac` is missing for type `insight_rogue_ap`")
    }
    if i.AvgRssi == nil {
        errs = append(errs, "required field `avg_rssi` is missing for type `insight_rogue_ap`")
    }
    if i.Bssid == nil {
        errs = append(errs, "required field `bssid` is missing for type `insight_rogue_ap`")
    }
    if i.Channel == nil {
        errs = append(errs, "required field `channel` is missing for type `insight_rogue_ap`")
    }
    if i.NumAps == nil {
        errs = append(errs, "required field `num_aps` is missing for type `insight_rogue_ap`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
