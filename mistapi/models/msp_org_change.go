package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// MspOrgChange represents a MspOrgChange struct.
type MspOrgChange struct {
    // enum: `assign`, `unassign`
    Op                   MspOrgChangeOperationEnum `json:"op"`
    // list of org_id
    OrgIds               []string                  `json:"org_ids"`
    AdditionalProperties map[string]any            `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MspOrgChange.
// It customizes the JSON marshaling process for MspOrgChange objects.
func (m MspOrgChange) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MspOrgChange object to a map representation for JSON marshaling.
func (m MspOrgChange) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
    structMap["op"] = m.Op
    structMap["org_ids"] = m.OrgIds
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MspOrgChange.
// It customizes the JSON unmarshaling process for MspOrgChange objects.
func (m *MspOrgChange) UnmarshalJSON(input []byte) error {
    var temp tempMspOrgChange
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "op", "org_ids")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.Op = *temp.Op
    m.OrgIds = *temp.OrgIds
    return nil
}

// tempMspOrgChange is a temporary struct used for validating the fields of MspOrgChange.
type tempMspOrgChange  struct {
    Op     *MspOrgChangeOperationEnum `json:"op"`
    OrgIds *[]string                  `json:"org_ids"`
}

func (m *tempMspOrgChange) validate() error {
    var errs []string
    if m.Op == nil {
        errs = append(errs, "required field `op` is missing for type `msp_org_change`")
    }
    if m.OrgIds == nil {
        errs = append(errs, "required field `org_ids` is missing for type `msp_org_change`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
