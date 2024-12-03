package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// AdminInvite represents a AdminInvite struct.
type AdminInvite struct {
    // skip creating initial setup if true
    AccountOnly          *bool                  `json:"account_only,omitempty"`
    // whether to allow Mist to look at this org
    AllowMist            *bool                  `json:"allow_mist,omitempty"`
    // city of registering user
    City                 *string                `json:"city,omitempty"`
    // country/region name or ISO code of registering user
    Country              *string                `json:"country,omitempty"`
    Email                string                 `json:"email"`
    FirstName            string                 `json:"first_name"`
    // required initially
    InviteCode           *string                `json:"invite_code,omitempty"`
    LastName             string                 `json:"last_name"`
    OrgName              string                 `json:"org_name"`
    Password             string                 `json:"password"`
    // reCAPTCHA , see https://www.google.com/recaptcha/
    Recaptcha            string                 `json:"recaptcha"`
    // flavor of the captcha. enum: `google`, `hcaptcha`
    RecaptchaFlavor      *RecaptchaFlavorEnum   `json:"recaptcha_flavor,omitempty"`
    // the invite token to apply after account creation
    RefererInviteToken   *string                `json:"referer_invite_token,omitempty"`
    // the url the user should be redirected back to
    ReturnTo             *string                `json:"return_to,omitempty"`
    // state name or ISO code of registering user, optional (depends on country/region)
    State                *string                `json:"state,omitempty"`
    // street address of registering user
    StreetAddress        *string                `json:"street_address,omitempty"`
    // street address 2 of registering user
    StreetAddress2       *string                `json:"street_address 2,omitempty"`
    // zipcode of registering user
    Zipcode              *string                `json:"zipcode,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for AdminInvite.
// It customizes the JSON marshaling process for AdminInvite objects.
func (a AdminInvite) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "account_only", "allow_mist", "city", "country", "email", "first_name", "invite_code", "last_name", "org_name", "password", "recaptcha", "recaptcha_flavor", "referer_invite_token", "return_to", "state", "street_address", "street_address 2", "zipcode"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the AdminInvite object to a map representation for JSON marshaling.
func (a AdminInvite) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    if a.AccountOnly != nil {
        structMap["account_only"] = a.AccountOnly
    }
    if a.AllowMist != nil {
        structMap["allow_mist"] = a.AllowMist
    }
    if a.City != nil {
        structMap["city"] = a.City
    }
    if a.Country != nil {
        structMap["country"] = a.Country
    }
    structMap["email"] = a.Email
    structMap["first_name"] = a.FirstName
    if a.InviteCode != nil {
        structMap["invite_code"] = a.InviteCode
    }
    structMap["last_name"] = a.LastName
    structMap["org_name"] = a.OrgName
    structMap["password"] = a.Password
    structMap["recaptcha"] = a.Recaptcha
    if a.RecaptchaFlavor != nil {
        structMap["recaptcha_flavor"] = a.RecaptchaFlavor
    }
    if a.RefererInviteToken != nil {
        structMap["referer_invite_token"] = a.RefererInviteToken
    }
    if a.ReturnTo != nil {
        structMap["return_to"] = a.ReturnTo
    }
    if a.State != nil {
        structMap["state"] = a.State
    }
    if a.StreetAddress != nil {
        structMap["street_address"] = a.StreetAddress
    }
    if a.StreetAddress2 != nil {
        structMap["street_address 2"] = a.StreetAddress2
    }
    if a.Zipcode != nil {
        structMap["zipcode"] = a.Zipcode
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AdminInvite.
// It customizes the JSON unmarshaling process for AdminInvite objects.
func (a *AdminInvite) UnmarshalJSON(input []byte) error {
    var temp tempAdminInvite
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "account_only", "allow_mist", "city", "country", "email", "first_name", "invite_code", "last_name", "org_name", "password", "recaptcha", "recaptcha_flavor", "referer_invite_token", "return_to", "state", "street_address", "street_address 2", "zipcode")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.AccountOnly = temp.AccountOnly
    a.AllowMist = temp.AllowMist
    a.City = temp.City
    a.Country = temp.Country
    a.Email = *temp.Email
    a.FirstName = *temp.FirstName
    a.InviteCode = temp.InviteCode
    a.LastName = *temp.LastName
    a.OrgName = *temp.OrgName
    a.Password = *temp.Password
    a.Recaptcha = *temp.Recaptcha
    a.RecaptchaFlavor = temp.RecaptchaFlavor
    a.RefererInviteToken = temp.RefererInviteToken
    a.ReturnTo = temp.ReturnTo
    a.State = temp.State
    a.StreetAddress = temp.StreetAddress
    a.StreetAddress2 = temp.StreetAddress2
    a.Zipcode = temp.Zipcode
    return nil
}

// tempAdminInvite is a temporary struct used for validating the fields of AdminInvite.
type tempAdminInvite  struct {
    AccountOnly        *bool                `json:"account_only,omitempty"`
    AllowMist          *bool                `json:"allow_mist,omitempty"`
    City               *string              `json:"city,omitempty"`
    Country            *string              `json:"country,omitempty"`
    Email              *string              `json:"email"`
    FirstName          *string              `json:"first_name"`
    InviteCode         *string              `json:"invite_code,omitempty"`
    LastName           *string              `json:"last_name"`
    OrgName            *string              `json:"org_name"`
    Password           *string              `json:"password"`
    Recaptcha          *string              `json:"recaptcha"`
    RecaptchaFlavor    *RecaptchaFlavorEnum `json:"recaptcha_flavor,omitempty"`
    RefererInviteToken *string              `json:"referer_invite_token,omitempty"`
    ReturnTo           *string              `json:"return_to,omitempty"`
    State              *string              `json:"state,omitempty"`
    StreetAddress      *string              `json:"street_address,omitempty"`
    StreetAddress2     *string              `json:"street_address 2,omitempty"`
    Zipcode            *string              `json:"zipcode,omitempty"`
}

func (a *tempAdminInvite) validate() error {
    var errs []string
    if a.Email == nil {
        errs = append(errs, "required field `email` is missing for type `admin_invite`")
    }
    if a.FirstName == nil {
        errs = append(errs, "required field `first_name` is missing for type `admin_invite`")
    }
    if a.LastName == nil {
        errs = append(errs, "required field `last_name` is missing for type `admin_invite`")
    }
    if a.OrgName == nil {
        errs = append(errs, "required field `org_name` is missing for type `admin_invite`")
    }
    if a.Password == nil {
        errs = append(errs, "required field `password` is missing for type `admin_invite`")
    }
    if a.Recaptcha == nil {
        errs = append(errs, "required field `recaptcha` is missing for type `admin_invite`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
