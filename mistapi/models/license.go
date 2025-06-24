package models

import (
    "encoding/json"
    "fmt"
)

// License represents a License struct.
// License
type License struct {
    Amendments           []LicenseAmendment     `json:"amendments,omitempty"`
    // Property key is license type (e.g. SUB-MAN) and Property value is the number of licenses entitled.
    Entitled             map[string]int         `json:"entitled,omitempty"`
    // Maximum number of licenses that may be required if the service is enabled on all the Organization Devices. Property key is the service name (e.g. "SUB-MAN").
    FullyLoaded          map[string]int         `json:"fully_loaded,omitempty"`
    Licenses             []LicenseSub           `json:"licenses,omitempty"`
    // Number of licenses currently consumed. Property key is license type (e.g. SUB-MAN).
    Summary              map[string]int         `json:"summary,omitempty"`
    // Number of available licenes. Property key is the service name (e.g. "SUB-MAN"). name (e.g. "SUB-MAN")
    Usages               map[string]int         `json:"usages,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for License,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (l License) String() string {
    return fmt.Sprintf(
    	"License[Amendments=%v, Entitled=%v, FullyLoaded=%v, Licenses=%v, Summary=%v, Usages=%v, AdditionalProperties=%v]",
    	l.Amendments, l.Entitled, l.FullyLoaded, l.Licenses, l.Summary, l.Usages, l.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for License.
// It customizes the JSON marshaling process for License objects.
func (l License) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(l.AdditionalProperties,
        "amendments", "entitled", "fully_loaded", "licenses", "summary", "usages"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(l.toMap())
}

// toMap converts the License object to a map representation for JSON marshaling.
func (l License) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, l.AdditionalProperties)
    if l.Amendments != nil {
        structMap["amendments"] = l.Amendments
    }
    if l.Entitled != nil {
        structMap["entitled"] = l.Entitled
    }
    if l.FullyLoaded != nil {
        structMap["fully_loaded"] = l.FullyLoaded
    }
    if l.Licenses != nil {
        structMap["licenses"] = l.Licenses
    }
    if l.Summary != nil {
        structMap["summary"] = l.Summary
    }
    if l.Usages != nil {
        structMap["usages"] = l.Usages
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for License.
// It customizes the JSON unmarshaling process for License objects.
func (l *License) UnmarshalJSON(input []byte) error {
    var temp tempLicense
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "amendments", "entitled", "fully_loaded", "licenses", "summary", "usages")
    if err != nil {
    	return err
    }
    l.AdditionalProperties = additionalProperties
    
    l.Amendments = temp.Amendments
    l.Entitled = temp.Entitled
    l.FullyLoaded = temp.FullyLoaded
    l.Licenses = temp.Licenses
    l.Summary = temp.Summary
    l.Usages = temp.Usages
    return nil
}

// tempLicense is a temporary struct used for validating the fields of License.
type tempLicense  struct {
    Amendments  []LicenseAmendment `json:"amendments,omitempty"`
    Entitled    map[string]int     `json:"entitled,omitempty"`
    FullyLoaded map[string]int     `json:"fully_loaded,omitempty"`
    Licenses    []LicenseSub       `json:"licenses,omitempty"`
    Summary     map[string]int     `json:"summary,omitempty"`
    Usages      map[string]int     `json:"usages,omitempty"`
}
