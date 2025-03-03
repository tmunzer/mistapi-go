package models

import (
    "encoding/json"
    "fmt"
)

// WlanDatarates represents a WlanDatarates struct.
// Data rates wlan settings
type WlanDatarates struct {
    // If `template`==`custom`. EHT MCS bitmasks for 4 streams (16-bit for each stream, MCS0 is least significant bit)
    Eht                  Optional[string]                    `json:"eht"`
    // If `template`==`custom`. HE MCS bitmasks for 4 streams (16-bit for each stream, MCS0 is least significant bit
    He                   Optional[string]                    `json:"he"`
    // If `template`==`custom`. MCS bitmasks for 4 streams (16-bit for each stream, MCS0 is least significant bit), e.g. 00ff 00f0 001f limits HT rates to MCS 0-7 for 1 stream, MCS 4-7 for 2 stream (i.e. MCS 12-15), MCS 1-5 for 3 stream (i.e. MCS 16-20)
    Ht                   Optional[string]                    `json:"ht"`
    // If `template`==`custom`. List of supported rates (IE=1) and extended supported rates (IE=50) for custom template, append ‘b’ at the end to indicate a rate being basic/mandatory. If `template`==`custom` is configured and legacy does not define at least one basic rate, it will use `no-legacy` default values
    Legacy               []WlanDataratesLegacyItemEnum       `json:"legacy,omitempty"`
    // Minimum RSSI for client to connect, 0 means not enforcing
    MinRssi              *int                                `json:"min_rssi,omitempty"`
    // Data Rates template to apply. enum:
    // * `no-legacy`: no 11b
    // * `compatible`: all, like before, default setting that Broadcom/Atheros used
    // * `legacy-only`: disable 802.11n and 802.11ac
    // * `high-density`: no 11b, no low rates
    // * `custom`: user defined
    Template             Optional[WlanDataratesTemplateEnum] `json:"template"`
    // If `template`==`custom`. MCS bitmasks for 4 streams (16-bit for each stream, MCS0 is least significant bit), e.g. 03ff 01ff 00ff limits VHT rates to MCS 0-9 for 1 stream, MCS 0-8 for 2 streams, and MCS 0-7 for 3 streams.
    Vht                  Optional[string]                    `json:"vht"`
    AdditionalProperties map[string]interface{}              `json:"_"`
}

// String implements the fmt.Stringer interface for WlanDatarates,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WlanDatarates) String() string {
    return fmt.Sprintf(
    	"WlanDatarates[Eht=%v, He=%v, Ht=%v, Legacy=%v, MinRssi=%v, Template=%v, Vht=%v, AdditionalProperties=%v]",
    	w.Eht, w.He, w.Ht, w.Legacy, w.MinRssi, w.Template, w.Vht, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WlanDatarates.
// It customizes the JSON marshaling process for WlanDatarates objects.
func (w WlanDatarates) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "eht", "he", "ht", "legacy", "min_rssi", "template", "vht"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WlanDatarates object to a map representation for JSON marshaling.
func (w WlanDatarates) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    if w.Eht.IsValueSet() {
        if w.Eht.Value() != nil {
            structMap["eht"] = w.Eht.Value()
        } else {
            structMap["eht"] = nil
        }
    }
    if w.He.IsValueSet() {
        if w.He.Value() != nil {
            structMap["he"] = w.He.Value()
        } else {
            structMap["he"] = nil
        }
    }
    if w.Ht.IsValueSet() {
        if w.Ht.Value() != nil {
            structMap["ht"] = w.Ht.Value()
        } else {
            structMap["ht"] = nil
        }
    }
    if w.Legacy != nil {
        structMap["legacy"] = w.Legacy
    }
    if w.MinRssi != nil {
        structMap["min_rssi"] = w.MinRssi
    }
    if w.Template.IsValueSet() {
        if w.Template.Value() != nil {
            structMap["template"] = w.Template.Value()
        } else {
            structMap["template"] = nil
        }
    }
    if w.Vht.IsValueSet() {
        if w.Vht.Value() != nil {
            structMap["vht"] = w.Vht.Value()
        } else {
            structMap["vht"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WlanDatarates.
// It customizes the JSON unmarshaling process for WlanDatarates objects.
func (w *WlanDatarates) UnmarshalJSON(input []byte) error {
    var temp tempWlanDatarates
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "eht", "he", "ht", "legacy", "min_rssi", "template", "vht")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.Eht = temp.Eht
    w.He = temp.He
    w.Ht = temp.Ht
    w.Legacy = temp.Legacy
    w.MinRssi = temp.MinRssi
    w.Template = temp.Template
    w.Vht = temp.Vht
    return nil
}

// tempWlanDatarates is a temporary struct used for validating the fields of WlanDatarates.
type tempWlanDatarates  struct {
    Eht      Optional[string]                    `json:"eht"`
    He       Optional[string]                    `json:"he"`
    Ht       Optional[string]                    `json:"ht"`
    Legacy   []WlanDataratesLegacyItemEnum       `json:"legacy,omitempty"`
    MinRssi  *int                                `json:"min_rssi,omitempty"`
    Template Optional[WlanDataratesTemplateEnum] `json:"template"`
    Vht      Optional[string]                    `json:"vht"`
}
