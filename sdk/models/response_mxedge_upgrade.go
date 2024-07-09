package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ResponseMxedgeUpgrade represents a ResponseMxedgeUpgrade struct.
type ResponseMxedgeUpgrade struct {
    Channel              string                      `json:"channel"`
    Counts               MxedgeUpgradeResponseCounts `json:"counts"`
    Id                   string                      `json:"id"`
    Status               string                      `json:"status"`
    Strategy             string                      `json:"strategy"`
    Versions             interface{}                 `json:"versions"`
    AdditionalProperties map[string]any              `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseMxedgeUpgrade.
// It customizes the JSON marshaling process for ResponseMxedgeUpgrade objects.
func (r ResponseMxedgeUpgrade) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseMxedgeUpgrade object to a map representation for JSON marshaling.
func (r ResponseMxedgeUpgrade) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["channel"] = r.Channel
    structMap["counts"] = r.Counts.toMap()
    structMap["id"] = r.Id
    structMap["status"] = r.Status
    structMap["strategy"] = r.Strategy
    structMap["versions"] = r.Versions
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseMxedgeUpgrade.
// It customizes the JSON unmarshaling process for ResponseMxedgeUpgrade objects.
func (r *ResponseMxedgeUpgrade) UnmarshalJSON(input []byte) error {
    var temp responseMxedgeUpgrade
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "channel", "counts", "id", "status", "strategy", "versions")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.Channel = *temp.Channel
    r.Counts = *temp.Counts
    r.Id = *temp.Id
    r.Status = *temp.Status
    r.Strategy = *temp.Strategy
    r.Versions = *temp.Versions
    return nil
}

// responseMxedgeUpgrade is a temporary struct used for validating the fields of ResponseMxedgeUpgrade.
type responseMxedgeUpgrade  struct {
    Channel  *string                      `json:"channel"`
    Counts   *MxedgeUpgradeResponseCounts `json:"counts"`
    Id       *string                      `json:"id"`
    Status   *string                      `json:"status"`
    Strategy *string                      `json:"strategy"`
    Versions *interface{}                 `json:"versions"`
}

func (r *responseMxedgeUpgrade) validate() error {
    var errs []string
    if r.Channel == nil {
        errs = append(errs, "required field `channel` is missing for type `Response_Mxedge_Upgrade`")
    }
    if r.Counts == nil {
        errs = append(errs, "required field `counts` is missing for type `Response_Mxedge_Upgrade`")
    }
    if r.Id == nil {
        errs = append(errs, "required field `id` is missing for type `Response_Mxedge_Upgrade`")
    }
    if r.Status == nil {
        errs = append(errs, "required field `status` is missing for type `Response_Mxedge_Upgrade`")
    }
    if r.Strategy == nil {
        errs = append(errs, "required field `strategy` is missing for type `Response_Mxedge_Upgrade`")
    }
    if r.Versions == nil {
        errs = append(errs, "required field `versions` is missing for type `Response_Mxedge_Upgrade`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
