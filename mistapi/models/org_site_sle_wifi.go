package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// OrgSiteSleWifi represents a OrgSiteSleWifi struct.
type OrgSiteSleWifi struct {
    End                  float64                `json:"end"`
    Interval             int                    `json:"interval"`
    Limit                int                    `json:"limit"`
    Page                 int                    `json:"page"`
    Results              []OrgSiteSleWifiResult `json:"results"`
    Start                float64                `json:"start"`
    Total                int                    `json:"total"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for OrgSiteSleWifi,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OrgSiteSleWifi) String() string {
    return fmt.Sprintf(
    	"OrgSiteSleWifi[End=%v, Interval=%v, Limit=%v, Page=%v, Results=%v, Start=%v, Total=%v, AdditionalProperties=%v]",
    	o.End, o.Interval, o.Limit, o.Page, o.Results, o.Start, o.Total, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OrgSiteSleWifi.
// It customizes the JSON marshaling process for OrgSiteSleWifi objects.
func (o OrgSiteSleWifi) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(o.AdditionalProperties,
        "end", "interval", "limit", "page", "results", "start", "total"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(o.toMap())
}

// toMap converts the OrgSiteSleWifi object to a map representation for JSON marshaling.
func (o OrgSiteSleWifi) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, o.AdditionalProperties)
    structMap["end"] = o.End
    structMap["interval"] = o.Interval
    structMap["limit"] = o.Limit
    structMap["page"] = o.Page
    structMap["results"] = o.Results
    structMap["start"] = o.Start
    structMap["total"] = o.Total
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSiteSleWifi.
// It customizes the JSON unmarshaling process for OrgSiteSleWifi objects.
func (o *OrgSiteSleWifi) UnmarshalJSON(input []byte) error {
    var temp tempOrgSiteSleWifi
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "end", "interval", "limit", "page", "results", "start", "total")
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

// tempOrgSiteSleWifi is a temporary struct used for validating the fields of OrgSiteSleWifi.
type tempOrgSiteSleWifi  struct {
    End      *float64                `json:"end"`
    Interval *int                    `json:"interval"`
    Limit    *int                    `json:"limit"`
    Page     *int                    `json:"page"`
    Results  *[]OrgSiteSleWifiResult `json:"results"`
    Start    *float64                `json:"start"`
    Total    *int                    `json:"total"`
}

func (o *tempOrgSiteSleWifi) validate() error {
    var errs []string
    if o.End == nil {
        errs = append(errs, "required field `end` is missing for type `org_site_sle_wifi`")
    }
    if o.Interval == nil {
        errs = append(errs, "required field `interval` is missing for type `org_site_sle_wifi`")
    }
    if o.Limit == nil {
        errs = append(errs, "required field `limit` is missing for type `org_site_sle_wifi`")
    }
    if o.Page == nil {
        errs = append(errs, "required field `page` is missing for type `org_site_sle_wifi`")
    }
    if o.Results == nil {
        errs = append(errs, "required field `results` is missing for type `org_site_sle_wifi`")
    }
    if o.Start == nil {
        errs = append(errs, "required field `start` is missing for type `org_site_sle_wifi`")
    }
    if o.Total == nil {
        errs = append(errs, "required field `total` is missing for type `org_site_sle_wifi`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
