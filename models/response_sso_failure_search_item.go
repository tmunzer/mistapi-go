package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ResponseSsoFailureSearchItem represents a ResponseSsoFailureSearchItem struct.
type ResponseSsoFailureSearchItem struct {
    Detail               string         `json:"detail"`
    SamlAssertionXml     string         `json:"saml_assertion_xml"`
    Timestamp            float64        `json:"timestamp"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseSsoFailureSearchItem.
// It customizes the JSON marshaling process for ResponseSsoFailureSearchItem objects.
func (r ResponseSsoFailureSearchItem) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseSsoFailureSearchItem object to a map representation for JSON marshaling.
func (r ResponseSsoFailureSearchItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["detail"] = r.Detail
    structMap["saml_assertion_xml"] = r.SamlAssertionXml
    structMap["timestamp"] = r.Timestamp
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseSsoFailureSearchItem.
// It customizes the JSON unmarshaling process for ResponseSsoFailureSearchItem objects.
func (r *ResponseSsoFailureSearchItem) UnmarshalJSON(input []byte) error {
    var temp responseSsoFailureSearchItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "detail", "saml_assertion_xml", "timestamp")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.Detail = *temp.Detail
    r.SamlAssertionXml = *temp.SamlAssertionXml
    r.Timestamp = *temp.Timestamp
    return nil
}

// responseSsoFailureSearchItem is a temporary struct used for validating the fields of ResponseSsoFailureSearchItem.
type responseSsoFailureSearchItem  struct {
    Detail           *string  `json:"detail"`
    SamlAssertionXml *string  `json:"saml_assertion_xml"`
    Timestamp        *float64 `json:"timestamp"`
}

func (r *responseSsoFailureSearchItem) validate() error {
    var errs []string
    if r.Detail == nil {
        errs = append(errs, "required field `detail` is missing for type `Response_Sso_Failure_Search_Item`")
    }
    if r.SamlAssertionXml == nil {
        errs = append(errs, "required field `saml_assertion_xml` is missing for type `Response_Sso_Failure_Search_Item`")
    }
    if r.Timestamp == nil {
        errs = append(errs, "required field `timestamp` is missing for type `Response_Sso_Failure_Search_Item`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}