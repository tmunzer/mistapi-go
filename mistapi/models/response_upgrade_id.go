package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// ResponseUpgradeId represents a ResponseUpgradeId struct.
type ResponseUpgradeId struct {
    UpgradeId            uuid.UUID              `json:"upgrade_id"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseUpgradeId,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseUpgradeId) String() string {
    return fmt.Sprintf(
    	"ResponseUpgradeId[UpgradeId=%v, AdditionalProperties=%v]",
    	r.UpgradeId, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseUpgradeId.
// It customizes the JSON marshaling process for ResponseUpgradeId objects.
func (r ResponseUpgradeId) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "upgrade_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseUpgradeId object to a map representation for JSON marshaling.
func (r ResponseUpgradeId) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["upgrade_id"] = r.UpgradeId
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseUpgradeId.
// It customizes the JSON unmarshaling process for ResponseUpgradeId objects.
func (r *ResponseUpgradeId) UnmarshalJSON(input []byte) error {
    var temp tempResponseUpgradeId
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "upgrade_id")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.UpgradeId = *temp.UpgradeId
    return nil
}

// tempResponseUpgradeId is a temporary struct used for validating the fields of ResponseUpgradeId.
type tempResponseUpgradeId  struct {
    UpgradeId *uuid.UUID `json:"upgrade_id"`
}

func (r *tempResponseUpgradeId) validate() error {
    var errs []string
    if r.UpgradeId == nil {
        errs = append(errs, "required field `upgrade_id` is missing for type `response_upgrade_id`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
