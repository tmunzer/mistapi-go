package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ResponseSelfAuditLogs represents a ResponseSelfAuditLogs struct.
type ResponseSelfAuditLogs struct {
    End                  int            `json:"end"`
    Limit                int            `json:"limit"`
    Page                 int            `json:"page"`
    Results              []AuditLog     `json:"results"`
    Start                int            `json:"start"`
    Total                int            `json:"total"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseSelfAuditLogs.
// It customizes the JSON marshaling process for ResponseSelfAuditLogs objects.
func (r ResponseSelfAuditLogs) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseSelfAuditLogs object to a map representation for JSON marshaling.
func (r ResponseSelfAuditLogs) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["end"] = r.End
    structMap["limit"] = r.Limit
    structMap["page"] = r.Page
    structMap["results"] = r.Results
    structMap["start"] = r.Start
    structMap["total"] = r.Total
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseSelfAuditLogs.
// It customizes the JSON unmarshaling process for ResponseSelfAuditLogs objects.
func (r *ResponseSelfAuditLogs) UnmarshalJSON(input []byte) error {
    var temp tempResponseSelfAuditLogs
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "end", "limit", "page", "results", "start", "total")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.End = *temp.End
    r.Limit = *temp.Limit
    r.Page = *temp.Page
    r.Results = *temp.Results
    r.Start = *temp.Start
    r.Total = *temp.Total
    return nil
}

// tempResponseSelfAuditLogs is a temporary struct used for validating the fields of ResponseSelfAuditLogs.
type tempResponseSelfAuditLogs  struct {
    End     *int        `json:"end"`
    Limit   *int        `json:"limit"`
    Page    *int        `json:"page"`
    Results *[]AuditLog `json:"results"`
    Start   *int        `json:"start"`
    Total   *int        `json:"total"`
}

func (r *tempResponseSelfAuditLogs) validate() error {
    var errs []string
    if r.End == nil {
        errs = append(errs, "required field `end` is missing for type `response_self_audit_logs`")
    }
    if r.Limit == nil {
        errs = append(errs, "required field `limit` is missing for type `response_self_audit_logs`")
    }
    if r.Page == nil {
        errs = append(errs, "required field `page` is missing for type `response_self_audit_logs`")
    }
    if r.Results == nil {
        errs = append(errs, "required field `results` is missing for type `response_self_audit_logs`")
    }
    if r.Start == nil {
        errs = append(errs, "required field `start` is missing for type `response_self_audit_logs`")
    }
    if r.Total == nil {
        errs = append(errs, "required field `total` is missing for type `response_self_audit_logs`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
