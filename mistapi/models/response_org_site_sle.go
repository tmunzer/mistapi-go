package models

import (
    "encoding/json"
)

// ResponseOrgSiteSle represents a ResponseOrgSiteSle struct.
type ResponseOrgSiteSle struct {
    End                  *float64                `json:"end,omitempty"`
    Interval             *int                    `json:"interval,omitempty"`
    Limit                *int                    `json:"limit,omitempty"`
    Page                 *int                    `json:"page,omitempty"`
    Results              []OrgSiteSleWifiResult1 `json:"results,omitempty"`
    Start                *float64                `json:"start,omitempty"`
    Total                *int                    `json:"total,omitempty"`
    AdditionalProperties map[string]any          `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseOrgSiteSle.
// It customizes the JSON marshaling process for ResponseOrgSiteSle objects.
func (r ResponseOrgSiteSle) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseOrgSiteSle object to a map representation for JSON marshaling.
func (r ResponseOrgSiteSle) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    if r.End != nil {
        structMap["end"] = r.End
    }
    if r.Interval != nil {
        structMap["interval"] = r.Interval
    }
    if r.Limit != nil {
        structMap["limit"] = r.Limit
    }
    if r.Page != nil {
        structMap["page"] = r.Page
    }
    if r.Results != nil {
        structMap["results"] = r.Results
    }
    if r.Start != nil {
        structMap["start"] = r.Start
    }
    if r.Total != nil {
        structMap["total"] = r.Total
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseOrgSiteSle.
// It customizes the JSON unmarshaling process for ResponseOrgSiteSle objects.
func (r *ResponseOrgSiteSle) UnmarshalJSON(input []byte) error {
    var temp tempResponseOrgSiteSle
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "end", "interval", "limit", "page", "results", "start", "total")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.End = temp.End
    r.Interval = temp.Interval
    r.Limit = temp.Limit
    r.Page = temp.Page
    r.Results = temp.Results
    r.Start = temp.Start
    r.Total = temp.Total
    return nil
}

// tempResponseOrgSiteSle is a temporary struct used for validating the fields of ResponseOrgSiteSle.
type tempResponseOrgSiteSle  struct {
    End      *float64                `json:"end,omitempty"`
    Interval *int                    `json:"interval,omitempty"`
    Limit    *int                    `json:"limit,omitempty"`
    Page     *int                    `json:"page,omitempty"`
    Results  []OrgSiteSleWifiResult1 `json:"results,omitempty"`
    Start    *float64                `json:"start,omitempty"`
    Total    *int                    `json:"total,omitempty"`
}
