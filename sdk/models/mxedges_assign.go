package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// MxedgesAssign represents a MxedgesAssign struct.
type MxedgesAssign struct {
    MxedgeIds            []uuid.UUID    `json:"mxedge_ids"`
    SiteId               uuid.UUID      `json:"site_id"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MxedgesAssign.
// It customizes the JSON marshaling process for MxedgesAssign objects.
func (m MxedgesAssign) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MxedgesAssign object to a map representation for JSON marshaling.
func (m MxedgesAssign) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
    structMap["mxedge_ids"] = m.MxedgeIds
    structMap["site_id"] = m.SiteId
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxedgesAssign.
// It customizes the JSON unmarshaling process for MxedgesAssign objects.
func (m *MxedgesAssign) UnmarshalJSON(input []byte) error {
    var temp mxedgesAssign
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "mxedge_ids", "site_id")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.MxedgeIds = *temp.MxedgeIds
    m.SiteId = *temp.SiteId
    return nil
}

// mxedgesAssign is a temporary struct used for validating the fields of MxedgesAssign.
type mxedgesAssign  struct {
    MxedgeIds *[]uuid.UUID `json:"mxedge_ids"`
    SiteId    *uuid.UUID   `json:"site_id"`
}

func (m *mxedgesAssign) validate() error {
    var errs []string
    if m.MxedgeIds == nil {
        errs = append(errs, "required field `mxedge_ids` is missing for type `Mxedges_Assign`")
    }
    if m.SiteId == nil {
        errs = append(errs, "required field `site_id` is missing for type `Mxedges_Assign`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
