package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// SkyatpListDomain represents a SkyatpListDomain struct.
type SkyatpListDomain struct {
    Comment              *string                `json:"comment,omitempty"`
    Value                string                 `json:"value"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SkyatpListDomain,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SkyatpListDomain) String() string {
    return fmt.Sprintf(
    	"SkyatpListDomain[Comment=%v, Value=%v, AdditionalProperties=%v]",
    	s.Comment, s.Value, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SkyatpListDomain.
// It customizes the JSON marshaling process for SkyatpListDomain objects.
func (s SkyatpListDomain) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "comment", "value"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SkyatpListDomain object to a map representation for JSON marshaling.
func (s SkyatpListDomain) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Comment != nil {
        structMap["comment"] = s.Comment
    }
    structMap["value"] = s.Value
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SkyatpListDomain.
// It customizes the JSON unmarshaling process for SkyatpListDomain objects.
func (s *SkyatpListDomain) UnmarshalJSON(input []byte) error {
    var temp tempSkyatpListDomain
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "comment", "value")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Comment = temp.Comment
    s.Value = *temp.Value
    return nil
}

// tempSkyatpListDomain is a temporary struct used for validating the fields of SkyatpListDomain.
type tempSkyatpListDomain  struct {
    Comment *string `json:"comment,omitempty"`
    Value   *string `json:"value"`
}

func (s *tempSkyatpListDomain) validate() error {
    var errs []string
    if s.Value == nil {
        errs = append(errs, "required field `value` is missing for type `skyatp_list_domain`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
