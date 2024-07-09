package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ResponseSelfOauthLinkSuccess represents a ResponseSelfOauthLinkSuccess struct.
type ResponseSelfOauthLinkSuccess struct {
    Action               string         `json:"action"`
    Id                   string         `json:"id"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseSelfOauthLinkSuccess.
// It customizes the JSON marshaling process for ResponseSelfOauthLinkSuccess objects.
func (r ResponseSelfOauthLinkSuccess) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseSelfOauthLinkSuccess object to a map representation for JSON marshaling.
func (r ResponseSelfOauthLinkSuccess) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["action"] = r.Action
    structMap["id"] = r.Id
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseSelfOauthLinkSuccess.
// It customizes the JSON unmarshaling process for ResponseSelfOauthLinkSuccess objects.
func (r *ResponseSelfOauthLinkSuccess) UnmarshalJSON(input []byte) error {
    var temp responseSelfOauthLinkSuccess
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "action", "id")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.Action = *temp.Action
    r.Id = *temp.Id
    return nil
}

// responseSelfOauthLinkSuccess is a temporary struct used for validating the fields of ResponseSelfOauthLinkSuccess.
type responseSelfOauthLinkSuccess  struct {
    Action *string `json:"action"`
    Id     *string `json:"id"`
}

func (r *responseSelfOauthLinkSuccess) validate() error {
    var errs []string
    if r.Action == nil {
        errs = append(errs, "required field `action` is missing for type `Response_Self_Oauth_Link_Success`")
    }
    if r.Id == nil {
        errs = append(errs, "required field `id` is missing for type `Response_Self_Oauth_Link_Success`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
