package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// MxedgesAssign represents a MxedgesAssign struct.
type MxedgesAssign struct {
    MxedgeIds            []uuid.UUID            `json:"mxedge_ids"`
    SiteId               uuid.UUID              `json:"site_id"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MxedgesAssign,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MxedgesAssign) String() string {
    return fmt.Sprintf(
    	"MxedgesAssign[MxedgeIds=%v, SiteId=%v, AdditionalProperties=%v]",
    	m.MxedgeIds, m.SiteId, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MxedgesAssign.
// It customizes the JSON marshaling process for MxedgesAssign objects.
func (m MxedgesAssign) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "mxedge_ids", "site_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the MxedgesAssign object to a map representation for JSON marshaling.
func (m MxedgesAssign) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
    structMap["mxedge_ids"] = m.MxedgeIds
    structMap["site_id"] = m.SiteId
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxedgesAssign.
// It customizes the JSON unmarshaling process for MxedgesAssign objects.
func (m *MxedgesAssign) UnmarshalJSON(input []byte) error {
    var temp tempMxedgesAssign
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "mxedge_ids", "site_id")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.MxedgeIds = *temp.MxedgeIds
    m.SiteId = *temp.SiteId
    return nil
}

// tempMxedgesAssign is a temporary struct used for validating the fields of MxedgesAssign.
type tempMxedgesAssign  struct {
    MxedgeIds *[]uuid.UUID `json:"mxedge_ids"`
    SiteId    *uuid.UUID   `json:"site_id"`
}

func (m *tempMxedgesAssign) validate() error {
    var errs []string
    if m.MxedgeIds == nil {
        errs = append(errs, "required field `mxedge_ids` is missing for type `mxedges_assign`")
    }
    if m.SiteId == nil {
        errs = append(errs, "required field `site_id` is missing for type `mxedges_assign`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
