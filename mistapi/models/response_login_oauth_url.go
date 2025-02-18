package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ResponseLoginOauthUrl represents a ResponseLoginOauthUrl struct.
type ResponseLoginOauthUrl struct {
    AuthorizationUrl     string                 `json:"authorization_url"`
    ClientId             string                 `json:"client_id"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseLoginOauthUrl,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseLoginOauthUrl) String() string {
    return fmt.Sprintf(
    	"ResponseLoginOauthUrl[AuthorizationUrl=%v, ClientId=%v, AdditionalProperties=%v]",
    	r.AuthorizationUrl, r.ClientId, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseLoginOauthUrl.
// It customizes the JSON marshaling process for ResponseLoginOauthUrl objects.
func (r ResponseLoginOauthUrl) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "authorization_url", "client_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseLoginOauthUrl object to a map representation for JSON marshaling.
func (r ResponseLoginOauthUrl) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["authorization_url"] = r.AuthorizationUrl
    structMap["client_id"] = r.ClientId
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseLoginOauthUrl.
// It customizes the JSON unmarshaling process for ResponseLoginOauthUrl objects.
func (r *ResponseLoginOauthUrl) UnmarshalJSON(input []byte) error {
    var temp tempResponseLoginOauthUrl
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "authorization_url", "client_id")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.AuthorizationUrl = *temp.AuthorizationUrl
    r.ClientId = *temp.ClientId
    return nil
}

// tempResponseLoginOauthUrl is a temporary struct used for validating the fields of ResponseLoginOauthUrl.
type tempResponseLoginOauthUrl  struct {
    AuthorizationUrl *string `json:"authorization_url"`
    ClientId         *string `json:"client_id"`
}

func (r *tempResponseLoginOauthUrl) validate() error {
    var errs []string
    if r.AuthorizationUrl == nil {
        errs = append(errs, "required field `authorization_url` is missing for type `response_login_oauth_url`")
    }
    if r.ClientId == nil {
        errs = append(errs, "required field `client_id` is missing for type `response_login_oauth_url`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
