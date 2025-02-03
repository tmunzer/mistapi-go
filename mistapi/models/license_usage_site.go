package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// LicenseUsageSite represents a LicenseUsageSite struct.
type LicenseUsageSite struct {
    // License entitlement for the entire org
    OrgEntitled          map[string]int         `json:"org_entitled"`
    // Eligibility for the Switch SLE
    SvnaEnabled          bool                   `json:"svna_enabled"`
    TrialEnabled         bool                   `json:"trial_enabled"`
    // Subscriptions and their quantities
    Usages               map[string]int         `json:"usages"`
    // Eligibility for the AP/Client SLE
    VnaEligible          bool                   `json:"vna_eligible"`
    // If True, Conversational Assistant and Marvis Action available
    VnaUi                bool                   `json:"vna_ui"`
    // Eligibility for the WAN SLE
    WvnaEligible         bool                   `json:"wvna_eligible"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for LicenseUsageSite,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (l LicenseUsageSite) String() string {
    return fmt.Sprintf(
    	"LicenseUsageSite[OrgEntitled=%v, SvnaEnabled=%v, TrialEnabled=%v, Usages=%v, VnaEligible=%v, VnaUi=%v, WvnaEligible=%v, AdditionalProperties=%v]",
    	l.OrgEntitled, l.SvnaEnabled, l.TrialEnabled, l.Usages, l.VnaEligible, l.VnaUi, l.WvnaEligible, l.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for LicenseUsageSite.
// It customizes the JSON marshaling process for LicenseUsageSite objects.
func (l LicenseUsageSite) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(l.AdditionalProperties,
        "org_entitled", "svna_enabled", "trial_enabled", "usages", "vna_eligible", "vna_ui", "wvna_eligible"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(l.toMap())
}

// toMap converts the LicenseUsageSite object to a map representation for JSON marshaling.
func (l LicenseUsageSite) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, l.AdditionalProperties)
    structMap["org_entitled"] = l.OrgEntitled
    structMap["svna_enabled"] = l.SvnaEnabled
    structMap["trial_enabled"] = l.TrialEnabled
    structMap["usages"] = l.Usages
    structMap["vna_eligible"] = l.VnaEligible
    structMap["vna_ui"] = l.VnaUi
    structMap["wvna_eligible"] = l.WvnaEligible
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for LicenseUsageSite.
// It customizes the JSON unmarshaling process for LicenseUsageSite objects.
func (l *LicenseUsageSite) UnmarshalJSON(input []byte) error {
    var temp tempLicenseUsageSite
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "org_entitled", "svna_enabled", "trial_enabled", "usages", "vna_eligible", "vna_ui", "wvna_eligible")
    if err != nil {
    	return err
    }
    l.AdditionalProperties = additionalProperties
    
    l.OrgEntitled = *temp.OrgEntitled
    l.SvnaEnabled = *temp.SvnaEnabled
    l.TrialEnabled = *temp.TrialEnabled
    l.Usages = *temp.Usages
    l.VnaEligible = *temp.VnaEligible
    l.VnaUi = *temp.VnaUi
    l.WvnaEligible = *temp.WvnaEligible
    return nil
}

// tempLicenseUsageSite is a temporary struct used for validating the fields of LicenseUsageSite.
type tempLicenseUsageSite  struct {
    OrgEntitled  *map[string]int `json:"org_entitled"`
    SvnaEnabled  *bool           `json:"svna_enabled"`
    TrialEnabled *bool           `json:"trial_enabled"`
    Usages       *map[string]int `json:"usages"`
    VnaEligible  *bool           `json:"vna_eligible"`
    VnaUi        *bool           `json:"vna_ui"`
    WvnaEligible *bool           `json:"wvna_eligible"`
}

func (l *tempLicenseUsageSite) validate() error {
    var errs []string
    if l.OrgEntitled == nil {
        errs = append(errs, "required field `org_entitled` is missing for type `license_usage_site`")
    }
    if l.SvnaEnabled == nil {
        errs = append(errs, "required field `svna_enabled` is missing for type `license_usage_site`")
    }
    if l.TrialEnabled == nil {
        errs = append(errs, "required field `trial_enabled` is missing for type `license_usage_site`")
    }
    if l.Usages == nil {
        errs = append(errs, "required field `usages` is missing for type `license_usage_site`")
    }
    if l.VnaEligible == nil {
        errs = append(errs, "required field `vna_eligible` is missing for type `license_usage_site`")
    }
    if l.VnaUi == nil {
        errs = append(errs, "required field `vna_ui` is missing for type `license_usage_site`")
    }
    if l.WvnaEligible == nil {
        errs = append(errs, "required field `wvna_eligible` is missing for type `license_usage_site`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
