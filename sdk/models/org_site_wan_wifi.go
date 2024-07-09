package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// OrgSiteWanWifi represents a OrgSiteWanWifi struct.
type OrgSiteWanWifi struct {
    End                  float64               `json:"end"`
    Interval             int                   `json:"interval"`
    Limit                int                   `json:"limit"`
    Page                 int                   `json:"page"`
    Results              []OrgSiteSleWanResult `json:"results"`
    Start                float64               `json:"start"`
    Total                int                   `json:"total"`
    AdditionalProperties map[string]any        `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OrgSiteWanWifi.
// It customizes the JSON marshaling process for OrgSiteWanWifi objects.
func (o OrgSiteWanWifi) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(o.toMap())
}

// toMap converts the OrgSiteWanWifi object to a map representation for JSON marshaling.
func (o OrgSiteWanWifi) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, o.AdditionalProperties)
    structMap["end"] = o.End
    structMap["interval"] = o.Interval
    structMap["limit"] = o.Limit
    structMap["page"] = o.Page
    structMap["results"] = o.Results
    structMap["start"] = o.Start
    structMap["total"] = o.Total
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSiteWanWifi.
// It customizes the JSON unmarshaling process for OrgSiteWanWifi objects.
func (o *OrgSiteWanWifi) UnmarshalJSON(input []byte) error {
    var temp orgSiteWanWifi
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "end", "interval", "limit", "page", "results", "start", "total")
    if err != nil {
    	return err
    }
    
    o.AdditionalProperties = additionalProperties
    o.End = *temp.End
    o.Interval = *temp.Interval
    o.Limit = *temp.Limit
    o.Page = *temp.Page
    o.Results = *temp.Results
    o.Start = *temp.Start
    o.Total = *temp.Total
    return nil
}

// orgSiteWanWifi is a temporary struct used for validating the fields of OrgSiteWanWifi.
type orgSiteWanWifi  struct {
    End      *float64               `json:"end"`
    Interval *int                   `json:"interval"`
    Limit    *int                   `json:"limit"`
    Page     *int                   `json:"page"`
    Results  *[]OrgSiteSleWanResult `json:"results"`
    Start    *float64               `json:"start"`
    Total    *int                   `json:"total"`
}

func (o *orgSiteWanWifi) validate() error {
    var errs []string
    if o.End == nil {
        errs = append(errs, "required field `end` is missing for type `Org_Site_Wan_Wifi`")
    }
    if o.Interval == nil {
        errs = append(errs, "required field `interval` is missing for type `Org_Site_Wan_Wifi`")
    }
    if o.Limit == nil {
        errs = append(errs, "required field `limit` is missing for type `Org_Site_Wan_Wifi`")
    }
    if o.Page == nil {
        errs = append(errs, "required field `page` is missing for type `Org_Site_Wan_Wifi`")
    }
    if o.Results == nil {
        errs = append(errs, "required field `results` is missing for type `Org_Site_Wan_Wifi`")
    }
    if o.Start == nil {
        errs = append(errs, "required field `start` is missing for type `Org_Site_Wan_Wifi`")
    }
    if o.Total == nil {
        errs = append(errs, "required field `total` is missing for type `Org_Site_Wan_Wifi`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
