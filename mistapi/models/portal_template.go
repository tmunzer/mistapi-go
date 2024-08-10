package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// PortalTemplate represents a PortalTemplate struct.
// Portal Template
type PortalTemplate struct {
    AccessCodeAlternateEmail *string        `json:"accessCodeAlternateEmail,omitempty"`
    // defines alignment on portal. “left” is default.
    Alignment                *string        `json:"alignment,omitempty"`
    // label for Amazon auth button
    AuthButtonAmazon         *string        `json:"authButtonAmazon,omitempty"`
    // label for Azure auth button
    AuthButtonAzure          *string        `json:"authButtonAzure,omitempty"`
    // label for Email auth button
    AuthButtonEmail          *string        `json:"authButtonEmail,omitempty"`
    // label for Facebook auth button
    AuthButtonFacebook       *string        `json:"authButtonFacebook,omitempty"`
    // label for Google auth button
    AuthButtonGoogle         *string        `json:"authButtonGoogle,omitempty"`
    // label for Microsoft auth button
    AuthButtonMicrosoft      *string        `json:"authButtonMicrosoft,omitempty"`
    // label for passphrase auth button
    AuthButtonPassphrase     *string        `json:"authButtonPassphrase,omitempty"`
    // label for SMS auth button
    AuthButtonSms            *string        `json:"authButtonSms,omitempty"`
    // label for Sponsor auth button
    AuthButtonSponsor        *string        `json:"authButtonSponsor,omitempty"`
    AuthLabel                *string        `json:"authLabel,omitempty"`
    // label of the link to go back to /logon
    BackLink                 *string        `json:"backLink,omitempty"`
    Color                    *string        `json:"color,omitempty"`
    ColorDark                *string        `json:"colorDark,omitempty"`
    ColorLight               *string        `json:"colorLight,omitempty"`
    // whether company field is required
    Company                  *bool          `json:"company,omitempty"`
    // error message when company not provided
    CompanyError             *string        `json:"companyError,omitempty"`
    // label of company field
    CompanyLabel             *string        `json:"companyLabel,omitempty"`
    CreatedTime              *float64       `json:"created_time,omitempty"`
    // whether email field is required
    Email                    *bool          `json:"email,omitempty"`
    // error message when a user has valid social login but doesn’ t match specified email domains.
    EmailAccessDomainError   *string        `json:"emailAccessDomainError,omitempty"`
    EmailCancel              *string        `json:"emailCancel,omitempty"`
    EmailCodeError           *string        `json:"emailCodeError,omitempty"`
    // error message when email not provided
    EmailError               *string        `json:"emailError,omitempty"`
    EmailFieldLabel          *string        `json:"emailFieldLabel,omitempty"`
    // label of email field
    EmailLabel               *string        `json:"emailLabel,omitempty"`
    EmailMessage             *string        `json:"emailMessage,omitempty"`
    EmailSubmit              *string        `json:"emailSubmit,omitempty"`
    EmailTitle               *string        `json:"emailTitle,omitempty"`
    // whether to ask field1
    Field1                   *bool          `json:"field1,omitempty"`
    // error message when field1 not provided
    Field1Error              *string        `json:"field1Error,omitempty"`
    // label of field1
    Field1Label              *string        `json:"field1Label,omitempty"`
    // whether field1 is required field
    Field1Required           *bool          `json:"field1Required,omitempty"`
    // whether to ask field2
    Field2                   *bool          `json:"field2,omitempty"`
    // error message when field2 not provided
    Field2Error              *string        `json:"field2Error,omitempty"`
    // label of field2
    Field2Label              *string        `json:"field2Label,omitempty"`
    // whether field2 is required field
    Field2Required           *bool          `json:"field2Required,omitempty"`
    // whether to ask field3
    Field3                   *bool          `json:"field3,omitempty"`
    // error message when field3 not provided
    Field3Error              *string        `json:"field3Error,omitempty"`
    // label of field3
    Field3Label              *string        `json:"field3Label,omitempty"`
    // whether field3 is required field
    Field3Required           *bool          `json:"field3Required,omitempty"`
    // whether to ask field4
    Field4                   *bool          `json:"field4,omitempty"`
    // error message when field4 not provided
    Field4Error              *string        `json:"field4Error,omitempty"`
    // label of field4
    Field4Label              *string        `json:"field4Label,omitempty"`
    // whether field4 is required field
    Field4Required           *bool          `json:"field4Required,omitempty"`
    ForSite                  *bool          `json:"for_site,omitempty"`
    Id                       *uuid.UUID     `json:"id,omitempty"`
    Message                  *string        `json:"message,omitempty"`
    ModifiedTime             *float64       `json:"modified_time,omitempty"`
    // whether name field is required
    Name                     *bool          `json:"name,omitempty"`
    // error message when name not provided
    NameError                *string        `json:"nameError,omitempty"`
    // label of name field
    NameLabel                *string        `json:"nameLabel,omitempty"`
    // whether to display “Do Not Store My Personal Information”
    Optout                   *bool          `json:"optout,omitempty"`
    OptoutLabel              *string        `json:"optoutLabel,omitempty"`
    OrgId                    *uuid.UUID     `json:"org_id,omitempty"`
    PageTitle                string         `json:"pageTitle"`
    PassphraseCancel         *string        `json:"passphraseCancel,omitempty"`
    // error message when invalid passphrase is provided
    PassphraseError          *string        `json:"passphraseError,omitempty"`
    PassphraseLabel          *string        `json:"passphraseLabel,omitempty"`
    PassphraseMessage        *string        `json:"passphraseMessage,omitempty"`
    PassphraseSubmit         *string        `json:"passphraseSubmit,omitempty"`
    // Title for passphrase details page
    PassphraseTitle          *string        `json:"passphraseTitle,omitempty"`
    // whether to show “Powered by Mist”
    PoweredBy                *bool          `json:"poweredBy,omitempty"`
    PrivacyPolicy            *bool          `json:"privacyPolicy,omitempty"`
    // prefix of the label of the link to go to /privacy_notice
    PrivacyPolicyAcceptLabel *string        `json:"privacyPolicyAcceptLabel,omitempty"`
    // error message when privacy notice is not accepted
    PrivacyPolicyError       *string        `json:"privacyPolicyError,omitempty"`
    // label of the link to go to /privacy_notice,
    PrivacyPolicyLink        *string        `json:"privacyPolicyLink,omitempty"`
    PrivacyPolicyText        *string        `json:"privacyPolicyText,omitempty"`
    // label to denote required field
    RequiredFieldLabel       *string        `json:"requiredFieldLabel,omitempty"`
    // label of the button to /signin
    SignInLabel              *string        `json:"signInLabel,omitempty"`
    SiteId                   *uuid.UUID     `json:"site_id,omitempty"`
    SmsCarrierDefault        *string        `json:"smsCarrierDefault,omitempty"`
    SmsCarrierError          *string        `json:"smsCarrierError,omitempty"`
    // label for mobile carrier drop-down list
    SmsCarrierFieldLabel     *string        `json:"smsCarrierFieldLabel,omitempty"`
    // Label for cancel confirmation code submission
    SmsCodeCancel            *string        `json:"smsCodeCancel,omitempty"`
    // error message when confirmation code is invalid
    SmsCodeError             *string        `json:"smsCodeError,omitempty"`
    SmsCodeFieldLabel        *string        `json:"smsCodeFieldLabel,omitempty"`
    SmsCodeMessage           *string        `json:"smsCodeMessage,omitempty"`
    // Label for confirmation code submit button
    SmsCodeSubmit            *string        `json:"smsCodeSubmit,omitempty"`
    SmsCodeTitle             *string        `json:"smsCodeTitle,omitempty"`
    // “Country Code”
    SmsCountryFieldLabel     *string        `json:"smsCountryFieldLabel,omitempty"`
    SmsCountryFormat         *string        `json:"smsCountryFormat,omitempty"`
    // Label for checkbox to specify that the user has access code
    SmsHaveAccessCode        *string        `json:"smsHaveAccessCode,omitempty"`
    // format of access code sms message. {{code}} and {{duration}} are place holders and should be retained as is.
    SmsMessageFormat         *string        `json:"smsMessageFormat,omitempty"`
    // label for canceling mobile details for SMS auth
    SmsNumberCancel          *string        `json:"smsNumberCancel,omitempty"`
    SmsNumberError           *string        `json:"smsNumberError,omitempty"`
    // label for field to provide mobile number
    SmsNumberFieldLabel      *string        `json:"smsNumberFieldLabel,omitempty"`
    SmsNumberFormat          *string        `json:"smsNumberFormat,omitempty"`
    SmsNumberMessage         *string        `json:"smsNumberMessage,omitempty"`
    // label for submit button for code generation
    SmsNumberSubmit          *string        `json:"smsNumberSubmit,omitempty"`
    // Title for phone number details
    SmsNumberTitle           *string        `json:"smsNumberTitle,omitempty"`
    SmsUsernameFormat        *string        `json:"smsUsernameFormat,omitempty"`
    // how long confirmation code should be considered valid (in minutes)
    SmsValidityDuration      *float64       `json:"smsValidityDuration,omitempty"`
    SponsorAutoApproved      *string        `json:"sponsorAutoApproved,omitempty"`
    SponsorAutoApprovedNote  *string        `json:"sponsorAutoApprovedNote,omitempty"`
    SponsorBackLink          *string        `json:"sponsorBackLink,omitempty"`
    SponsorCancel            *string        `json:"sponsorCancel,omitempty"`
    // label for Sponsor Email
    SponsorEmail             *string        `json:"sponsorEmail,omitempty"`
    SponsorEmailError        *string        `json:"sponsorEmailError,omitempty"`
    // “html template to replace/override default sponsor email template”
    SponsorEmailTemplate     *string        `json:"sponsorEmailTemplate,omitempty"`
    SponsorInfoApproved      *string        `json:"sponsorInfoApproved,omitempty"`
    SponsorInfoDenied        *string        `json:"sponsorInfoDenied,omitempty"`
    SponsorInfoPending       *string        `json:"sponsorInfoPending,omitempty"`
    // label for Sponsor Name
    SponsorName              *string        `json:"sponsorName,omitempty"`
    SponsorNameError         *string        `json:"sponsorNameError,omitempty"`
    SponsorNotePending       *string        `json:"sponsorNotePending,omitempty"`
    // text to display if sponsor approves request
    SponsorStatusApproved    *string        `json:"sponsorStatusApproved,omitempty"`
    // text to display when sponsor denies request
    SponsorStatusDenied      *string        `json:"sponsorStatusDenied,omitempty"`
    // text to display if request is still pending
    SponsorStatusPending     *string        `json:"sponsorStatusPending,omitempty"`
    // submit button label to notify sponsor about guest request
    SponsorSubmit            *string        `json:"sponsorSubmit,omitempty"`
    SponsorsAutoApprovedNote *string        `json:"sponsorsAutoApprovedNote,omitempty"`
    SponsorsError            *string        `json:"sponsorsError,omitempty"`
    Tos                      *bool          `json:"tos,omitempty"`
    // prefix of the label of the link to go to /tos
    TosAcceptLabel           *string        `json:"tosAcceptLabel,omitempty"`
    // error message when tos not accepted
    TosError                 *string        `json:"tosError,omitempty"`
    // label of the link to go to /tos
    TosLink                  *string        `json:"tosLink,omitempty"`
    // text of the Terms of Service
    TosText                  *string        `json:"tosText,omitempty"`
    AdditionalProperties     map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for PortalTemplate.
// It customizes the JSON marshaling process for PortalTemplate objects.
func (p PortalTemplate) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the PortalTemplate object to a map representation for JSON marshaling.
func (p PortalTemplate) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, p.AdditionalProperties)
    if p.AccessCodeAlternateEmail != nil {
        structMap["accessCodeAlternateEmail"] = p.AccessCodeAlternateEmail
    }
    if p.Alignment != nil {
        structMap["alignment"] = p.Alignment
    }
    if p.AuthButtonAmazon != nil {
        structMap["authButtonAmazon"] = p.AuthButtonAmazon
    }
    if p.AuthButtonAzure != nil {
        structMap["authButtonAzure"] = p.AuthButtonAzure
    }
    if p.AuthButtonEmail != nil {
        structMap["authButtonEmail"] = p.AuthButtonEmail
    }
    if p.AuthButtonFacebook != nil {
        structMap["authButtonFacebook"] = p.AuthButtonFacebook
    }
    if p.AuthButtonGoogle != nil {
        structMap["authButtonGoogle"] = p.AuthButtonGoogle
    }
    if p.AuthButtonMicrosoft != nil {
        structMap["authButtonMicrosoft"] = p.AuthButtonMicrosoft
    }
    if p.AuthButtonPassphrase != nil {
        structMap["authButtonPassphrase"] = p.AuthButtonPassphrase
    }
    if p.AuthButtonSms != nil {
        structMap["authButtonSms"] = p.AuthButtonSms
    }
    if p.AuthButtonSponsor != nil {
        structMap["authButtonSponsor"] = p.AuthButtonSponsor
    }
    if p.AuthLabel != nil {
        structMap["authLabel"] = p.AuthLabel
    }
    if p.BackLink != nil {
        structMap["backLink"] = p.BackLink
    }
    if p.Color != nil {
        structMap["color"] = p.Color
    }
    if p.ColorDark != nil {
        structMap["colorDark"] = p.ColorDark
    }
    if p.ColorLight != nil {
        structMap["colorLight"] = p.ColorLight
    }
    if p.Company != nil {
        structMap["company"] = p.Company
    }
    if p.CompanyError != nil {
        structMap["companyError"] = p.CompanyError
    }
    if p.CompanyLabel != nil {
        structMap["companyLabel"] = p.CompanyLabel
    }
    if p.CreatedTime != nil {
        structMap["created_time"] = p.CreatedTime
    }
    if p.Email != nil {
        structMap["email"] = p.Email
    }
    if p.EmailAccessDomainError != nil {
        structMap["emailAccessDomainError"] = p.EmailAccessDomainError
    }
    if p.EmailCancel != nil {
        structMap["emailCancel"] = p.EmailCancel
    }
    if p.EmailCodeError != nil {
        structMap["emailCodeError"] = p.EmailCodeError
    }
    if p.EmailError != nil {
        structMap["emailError"] = p.EmailError
    }
    if p.EmailFieldLabel != nil {
        structMap["emailFieldLabel"] = p.EmailFieldLabel
    }
    if p.EmailLabel != nil {
        structMap["emailLabel"] = p.EmailLabel
    }
    if p.EmailMessage != nil {
        structMap["emailMessage"] = p.EmailMessage
    }
    if p.EmailSubmit != nil {
        structMap["emailSubmit"] = p.EmailSubmit
    }
    if p.EmailTitle != nil {
        structMap["emailTitle"] = p.EmailTitle
    }
    if p.Field1 != nil {
        structMap["field1"] = p.Field1
    }
    if p.Field1Error != nil {
        structMap["field1Error"] = p.Field1Error
    }
    if p.Field1Label != nil {
        structMap["field1Label"] = p.Field1Label
    }
    if p.Field1Required != nil {
        structMap["field1Required"] = p.Field1Required
    }
    if p.Field2 != nil {
        structMap["field2"] = p.Field2
    }
    if p.Field2Error != nil {
        structMap["field2Error"] = p.Field2Error
    }
    if p.Field2Label != nil {
        structMap["field2Label"] = p.Field2Label
    }
    if p.Field2Required != nil {
        structMap["field2Required"] = p.Field2Required
    }
    if p.Field3 != nil {
        structMap["field3"] = p.Field3
    }
    if p.Field3Error != nil {
        structMap["field3Error"] = p.Field3Error
    }
    if p.Field3Label != nil {
        structMap["field3Label"] = p.Field3Label
    }
    if p.Field3Required != nil {
        structMap["field3Required"] = p.Field3Required
    }
    if p.Field4 != nil {
        structMap["field4"] = p.Field4
    }
    if p.Field4Error != nil {
        structMap["field4Error"] = p.Field4Error
    }
    if p.Field4Label != nil {
        structMap["field4Label"] = p.Field4Label
    }
    if p.Field4Required != nil {
        structMap["field4Required"] = p.Field4Required
    }
    if p.ForSite != nil {
        structMap["for_site"] = p.ForSite
    }
    if p.Id != nil {
        structMap["id"] = p.Id
    }
    if p.Message != nil {
        structMap["message"] = p.Message
    }
    if p.ModifiedTime != nil {
        structMap["modified_time"] = p.ModifiedTime
    }
    if p.Name != nil {
        structMap["name"] = p.Name
    }
    if p.NameError != nil {
        structMap["nameError"] = p.NameError
    }
    if p.NameLabel != nil {
        structMap["nameLabel"] = p.NameLabel
    }
    if p.Optout != nil {
        structMap["optout"] = p.Optout
    }
    if p.OptoutLabel != nil {
        structMap["optoutLabel"] = p.OptoutLabel
    }
    if p.OrgId != nil {
        structMap["org_id"] = p.OrgId
    }
    structMap["pageTitle"] = p.PageTitle
    if p.PassphraseCancel != nil {
        structMap["passphraseCancel"] = p.PassphraseCancel
    }
    if p.PassphraseError != nil {
        structMap["passphraseError"] = p.PassphraseError
    }
    if p.PassphraseLabel != nil {
        structMap["passphraseLabel"] = p.PassphraseLabel
    }
    if p.PassphraseMessage != nil {
        structMap["passphraseMessage"] = p.PassphraseMessage
    }
    if p.PassphraseSubmit != nil {
        structMap["passphraseSubmit"] = p.PassphraseSubmit
    }
    if p.PassphraseTitle != nil {
        structMap["passphraseTitle"] = p.PassphraseTitle
    }
    if p.PoweredBy != nil {
        structMap["poweredBy"] = p.PoweredBy
    }
    if p.PrivacyPolicy != nil {
        structMap["privacyPolicy"] = p.PrivacyPolicy
    }
    if p.PrivacyPolicyAcceptLabel != nil {
        structMap["privacyPolicyAcceptLabel"] = p.PrivacyPolicyAcceptLabel
    }
    if p.PrivacyPolicyError != nil {
        structMap["privacyPolicyError"] = p.PrivacyPolicyError
    }
    if p.PrivacyPolicyLink != nil {
        structMap["privacyPolicyLink"] = p.PrivacyPolicyLink
    }
    if p.PrivacyPolicyText != nil {
        structMap["privacyPolicyText"] = p.PrivacyPolicyText
    }
    if p.RequiredFieldLabel != nil {
        structMap["requiredFieldLabel"] = p.RequiredFieldLabel
    }
    if p.SignInLabel != nil {
        structMap["signInLabel"] = p.SignInLabel
    }
    if p.SiteId != nil {
        structMap["site_id"] = p.SiteId
    }
    if p.SmsCarrierDefault != nil {
        structMap["smsCarrierDefault"] = p.SmsCarrierDefault
    }
    if p.SmsCarrierError != nil {
        structMap["smsCarrierError"] = p.SmsCarrierError
    }
    if p.SmsCarrierFieldLabel != nil {
        structMap["smsCarrierFieldLabel"] = p.SmsCarrierFieldLabel
    }
    if p.SmsCodeCancel != nil {
        structMap["smsCodeCancel"] = p.SmsCodeCancel
    }
    if p.SmsCodeError != nil {
        structMap["smsCodeError"] = p.SmsCodeError
    }
    if p.SmsCodeFieldLabel != nil {
        structMap["smsCodeFieldLabel"] = p.SmsCodeFieldLabel
    }
    if p.SmsCodeMessage != nil {
        structMap["smsCodeMessage"] = p.SmsCodeMessage
    }
    if p.SmsCodeSubmit != nil {
        structMap["smsCodeSubmit"] = p.SmsCodeSubmit
    }
    if p.SmsCodeTitle != nil {
        structMap["smsCodeTitle"] = p.SmsCodeTitle
    }
    if p.SmsCountryFieldLabel != nil {
        structMap["smsCountryFieldLabel"] = p.SmsCountryFieldLabel
    }
    if p.SmsCountryFormat != nil {
        structMap["smsCountryFormat"] = p.SmsCountryFormat
    }
    if p.SmsHaveAccessCode != nil {
        structMap["smsHaveAccessCode"] = p.SmsHaveAccessCode
    }
    if p.SmsMessageFormat != nil {
        structMap["smsMessageFormat"] = p.SmsMessageFormat
    }
    if p.SmsNumberCancel != nil {
        structMap["smsNumberCancel"] = p.SmsNumberCancel
    }
    if p.SmsNumberError != nil {
        structMap["smsNumberError"] = p.SmsNumberError
    }
    if p.SmsNumberFieldLabel != nil {
        structMap["smsNumberFieldLabel"] = p.SmsNumberFieldLabel
    }
    if p.SmsNumberFormat != nil {
        structMap["smsNumberFormat"] = p.SmsNumberFormat
    }
    if p.SmsNumberMessage != nil {
        structMap["smsNumberMessage"] = p.SmsNumberMessage
    }
    if p.SmsNumberSubmit != nil {
        structMap["smsNumberSubmit"] = p.SmsNumberSubmit
    }
    if p.SmsNumberTitle != nil {
        structMap["smsNumberTitle"] = p.SmsNumberTitle
    }
    if p.SmsUsernameFormat != nil {
        structMap["smsUsernameFormat"] = p.SmsUsernameFormat
    }
    if p.SmsValidityDuration != nil {
        structMap["smsValidityDuration"] = p.SmsValidityDuration
    }
    if p.SponsorAutoApproved != nil {
        structMap["sponsorAutoApproved"] = p.SponsorAutoApproved
    }
    if p.SponsorAutoApprovedNote != nil {
        structMap["sponsorAutoApprovedNote"] = p.SponsorAutoApprovedNote
    }
    if p.SponsorBackLink != nil {
        structMap["sponsorBackLink"] = p.SponsorBackLink
    }
    if p.SponsorCancel != nil {
        structMap["sponsorCancel"] = p.SponsorCancel
    }
    if p.SponsorEmail != nil {
        structMap["sponsorEmail"] = p.SponsorEmail
    }
    if p.SponsorEmailError != nil {
        structMap["sponsorEmailError"] = p.SponsorEmailError
    }
    if p.SponsorEmailTemplate != nil {
        structMap["sponsorEmailTemplate"] = p.SponsorEmailTemplate
    }
    if p.SponsorInfoApproved != nil {
        structMap["sponsorInfoApproved"] = p.SponsorInfoApproved
    }
    if p.SponsorInfoDenied != nil {
        structMap["sponsorInfoDenied"] = p.SponsorInfoDenied
    }
    if p.SponsorInfoPending != nil {
        structMap["sponsorInfoPending"] = p.SponsorInfoPending
    }
    if p.SponsorName != nil {
        structMap["sponsorName"] = p.SponsorName
    }
    if p.SponsorNameError != nil {
        structMap["sponsorNameError"] = p.SponsorNameError
    }
    if p.SponsorNotePending != nil {
        structMap["sponsorNotePending"] = p.SponsorNotePending
    }
    if p.SponsorStatusApproved != nil {
        structMap["sponsorStatusApproved"] = p.SponsorStatusApproved
    }
    if p.SponsorStatusDenied != nil {
        structMap["sponsorStatusDenied"] = p.SponsorStatusDenied
    }
    if p.SponsorStatusPending != nil {
        structMap["sponsorStatusPending"] = p.SponsorStatusPending
    }
    if p.SponsorSubmit != nil {
        structMap["sponsorSubmit"] = p.SponsorSubmit
    }
    if p.SponsorsAutoApprovedNote != nil {
        structMap["sponsorsAutoApprovedNote"] = p.SponsorsAutoApprovedNote
    }
    if p.SponsorsError != nil {
        structMap["sponsorsError"] = p.SponsorsError
    }
    if p.Tos != nil {
        structMap["tos"] = p.Tos
    }
    if p.TosAcceptLabel != nil {
        structMap["tosAcceptLabel"] = p.TosAcceptLabel
    }
    if p.TosError != nil {
        structMap["tosError"] = p.TosError
    }
    if p.TosLink != nil {
        structMap["tosLink"] = p.TosLink
    }
    if p.TosText != nil {
        structMap["tosText"] = p.TosText
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PortalTemplate.
// It customizes the JSON unmarshaling process for PortalTemplate objects.
func (p *PortalTemplate) UnmarshalJSON(input []byte) error {
    var temp tempPortalTemplate
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "accessCodeAlternateEmail", "alignment", "authButtonAmazon", "authButtonAzure", "authButtonEmail", "authButtonFacebook", "authButtonGoogle", "authButtonMicrosoft", "authButtonPassphrase", "authButtonSms", "authButtonSponsor", "authLabel", "backLink", "color", "colorDark", "colorLight", "company", "companyError", "companyLabel", "created_time", "email", "emailAccessDomainError", "emailCancel", "emailCodeError", "emailError", "emailFieldLabel", "emailLabel", "emailMessage", "emailSubmit", "emailTitle", "field1", "field1Error", "field1Label", "field1Required", "field2", "field2Error", "field2Label", "field2Required", "field3", "field3Error", "field3Label", "field3Required", "field4", "field4Error", "field4Label", "field4Required", "for_site", "id", "message", "modified_time", "name", "nameError", "nameLabel", "optout", "optoutLabel", "org_id", "pageTitle", "passphraseCancel", "passphraseError", "passphraseLabel", "passphraseMessage", "passphraseSubmit", "passphraseTitle", "poweredBy", "privacyPolicy", "privacyPolicyAcceptLabel", "privacyPolicyError", "privacyPolicyLink", "privacyPolicyText", "requiredFieldLabel", "signInLabel", "site_id", "smsCarrierDefault", "smsCarrierError", "smsCarrierFieldLabel", "smsCodeCancel", "smsCodeError", "smsCodeFieldLabel", "smsCodeMessage", "smsCodeSubmit", "smsCodeTitle", "smsCountryFieldLabel", "smsCountryFormat", "smsHaveAccessCode", "smsMessageFormat", "smsNumberCancel", "smsNumberError", "smsNumberFieldLabel", "smsNumberFormat", "smsNumberMessage", "smsNumberSubmit", "smsNumberTitle", "smsUsernameFormat", "smsValidityDuration", "sponsorAutoApproved", "sponsorAutoApprovedNote", "sponsorBackLink", "sponsorCancel", "sponsorEmail", "sponsorEmailError", "sponsorEmailTemplate", "sponsorInfoApproved", "sponsorInfoDenied", "sponsorInfoPending", "sponsorName", "sponsorNameError", "sponsorNotePending", "sponsorStatusApproved", "sponsorStatusDenied", "sponsorStatusPending", "sponsorSubmit", "sponsorsAutoApprovedNote", "sponsorsError", "tos", "tosAcceptLabel", "tosError", "tosLink", "tosText")
    if err != nil {
    	return err
    }
    
    p.AdditionalProperties = additionalProperties
    p.AccessCodeAlternateEmail = temp.AccessCodeAlternateEmail
    p.Alignment = temp.Alignment
    p.AuthButtonAmazon = temp.AuthButtonAmazon
    p.AuthButtonAzure = temp.AuthButtonAzure
    p.AuthButtonEmail = temp.AuthButtonEmail
    p.AuthButtonFacebook = temp.AuthButtonFacebook
    p.AuthButtonGoogle = temp.AuthButtonGoogle
    p.AuthButtonMicrosoft = temp.AuthButtonMicrosoft
    p.AuthButtonPassphrase = temp.AuthButtonPassphrase
    p.AuthButtonSms = temp.AuthButtonSms
    p.AuthButtonSponsor = temp.AuthButtonSponsor
    p.AuthLabel = temp.AuthLabel
    p.BackLink = temp.BackLink
    p.Color = temp.Color
    p.ColorDark = temp.ColorDark
    p.ColorLight = temp.ColorLight
    p.Company = temp.Company
    p.CompanyError = temp.CompanyError
    p.CompanyLabel = temp.CompanyLabel
    p.CreatedTime = temp.CreatedTime
    p.Email = temp.Email
    p.EmailAccessDomainError = temp.EmailAccessDomainError
    p.EmailCancel = temp.EmailCancel
    p.EmailCodeError = temp.EmailCodeError
    p.EmailError = temp.EmailError
    p.EmailFieldLabel = temp.EmailFieldLabel
    p.EmailLabel = temp.EmailLabel
    p.EmailMessage = temp.EmailMessage
    p.EmailSubmit = temp.EmailSubmit
    p.EmailTitle = temp.EmailTitle
    p.Field1 = temp.Field1
    p.Field1Error = temp.Field1Error
    p.Field1Label = temp.Field1Label
    p.Field1Required = temp.Field1Required
    p.Field2 = temp.Field2
    p.Field2Error = temp.Field2Error
    p.Field2Label = temp.Field2Label
    p.Field2Required = temp.Field2Required
    p.Field3 = temp.Field3
    p.Field3Error = temp.Field3Error
    p.Field3Label = temp.Field3Label
    p.Field3Required = temp.Field3Required
    p.Field4 = temp.Field4
    p.Field4Error = temp.Field4Error
    p.Field4Label = temp.Field4Label
    p.Field4Required = temp.Field4Required
    p.ForSite = temp.ForSite
    p.Id = temp.Id
    p.Message = temp.Message
    p.ModifiedTime = temp.ModifiedTime
    p.Name = temp.Name
    p.NameError = temp.NameError
    p.NameLabel = temp.NameLabel
    p.Optout = temp.Optout
    p.OptoutLabel = temp.OptoutLabel
    p.OrgId = temp.OrgId
    p.PageTitle = *temp.PageTitle
    p.PassphraseCancel = temp.PassphraseCancel
    p.PassphraseError = temp.PassphraseError
    p.PassphraseLabel = temp.PassphraseLabel
    p.PassphraseMessage = temp.PassphraseMessage
    p.PassphraseSubmit = temp.PassphraseSubmit
    p.PassphraseTitle = temp.PassphraseTitle
    p.PoweredBy = temp.PoweredBy
    p.PrivacyPolicy = temp.PrivacyPolicy
    p.PrivacyPolicyAcceptLabel = temp.PrivacyPolicyAcceptLabel
    p.PrivacyPolicyError = temp.PrivacyPolicyError
    p.PrivacyPolicyLink = temp.PrivacyPolicyLink
    p.PrivacyPolicyText = temp.PrivacyPolicyText
    p.RequiredFieldLabel = temp.RequiredFieldLabel
    p.SignInLabel = temp.SignInLabel
    p.SiteId = temp.SiteId
    p.SmsCarrierDefault = temp.SmsCarrierDefault
    p.SmsCarrierError = temp.SmsCarrierError
    p.SmsCarrierFieldLabel = temp.SmsCarrierFieldLabel
    p.SmsCodeCancel = temp.SmsCodeCancel
    p.SmsCodeError = temp.SmsCodeError
    p.SmsCodeFieldLabel = temp.SmsCodeFieldLabel
    p.SmsCodeMessage = temp.SmsCodeMessage
    p.SmsCodeSubmit = temp.SmsCodeSubmit
    p.SmsCodeTitle = temp.SmsCodeTitle
    p.SmsCountryFieldLabel = temp.SmsCountryFieldLabel
    p.SmsCountryFormat = temp.SmsCountryFormat
    p.SmsHaveAccessCode = temp.SmsHaveAccessCode
    p.SmsMessageFormat = temp.SmsMessageFormat
    p.SmsNumberCancel = temp.SmsNumberCancel
    p.SmsNumberError = temp.SmsNumberError
    p.SmsNumberFieldLabel = temp.SmsNumberFieldLabel
    p.SmsNumberFormat = temp.SmsNumberFormat
    p.SmsNumberMessage = temp.SmsNumberMessage
    p.SmsNumberSubmit = temp.SmsNumberSubmit
    p.SmsNumberTitle = temp.SmsNumberTitle
    p.SmsUsernameFormat = temp.SmsUsernameFormat
    p.SmsValidityDuration = temp.SmsValidityDuration
    p.SponsorAutoApproved = temp.SponsorAutoApproved
    p.SponsorAutoApprovedNote = temp.SponsorAutoApprovedNote
    p.SponsorBackLink = temp.SponsorBackLink
    p.SponsorCancel = temp.SponsorCancel
    p.SponsorEmail = temp.SponsorEmail
    p.SponsorEmailError = temp.SponsorEmailError
    p.SponsorEmailTemplate = temp.SponsorEmailTemplate
    p.SponsorInfoApproved = temp.SponsorInfoApproved
    p.SponsorInfoDenied = temp.SponsorInfoDenied
    p.SponsorInfoPending = temp.SponsorInfoPending
    p.SponsorName = temp.SponsorName
    p.SponsorNameError = temp.SponsorNameError
    p.SponsorNotePending = temp.SponsorNotePending
    p.SponsorStatusApproved = temp.SponsorStatusApproved
    p.SponsorStatusDenied = temp.SponsorStatusDenied
    p.SponsorStatusPending = temp.SponsorStatusPending
    p.SponsorSubmit = temp.SponsorSubmit
    p.SponsorsAutoApprovedNote = temp.SponsorsAutoApprovedNote
    p.SponsorsError = temp.SponsorsError
    p.Tos = temp.Tos
    p.TosAcceptLabel = temp.TosAcceptLabel
    p.TosError = temp.TosError
    p.TosLink = temp.TosLink
    p.TosText = temp.TosText
    return nil
}

// tempPortalTemplate is a temporary struct used for validating the fields of PortalTemplate.
type tempPortalTemplate  struct {
    AccessCodeAlternateEmail *string    `json:"accessCodeAlternateEmail,omitempty"`
    Alignment                *string    `json:"alignment,omitempty"`
    AuthButtonAmazon         *string    `json:"authButtonAmazon,omitempty"`
    AuthButtonAzure          *string    `json:"authButtonAzure,omitempty"`
    AuthButtonEmail          *string    `json:"authButtonEmail,omitempty"`
    AuthButtonFacebook       *string    `json:"authButtonFacebook,omitempty"`
    AuthButtonGoogle         *string    `json:"authButtonGoogle,omitempty"`
    AuthButtonMicrosoft      *string    `json:"authButtonMicrosoft,omitempty"`
    AuthButtonPassphrase     *string    `json:"authButtonPassphrase,omitempty"`
    AuthButtonSms            *string    `json:"authButtonSms,omitempty"`
    AuthButtonSponsor        *string    `json:"authButtonSponsor,omitempty"`
    AuthLabel                *string    `json:"authLabel,omitempty"`
    BackLink                 *string    `json:"backLink,omitempty"`
    Color                    *string    `json:"color,omitempty"`
    ColorDark                *string    `json:"colorDark,omitempty"`
    ColorLight               *string    `json:"colorLight,omitempty"`
    Company                  *bool      `json:"company,omitempty"`
    CompanyError             *string    `json:"companyError,omitempty"`
    CompanyLabel             *string    `json:"companyLabel,omitempty"`
    CreatedTime              *float64   `json:"created_time,omitempty"`
    Email                    *bool      `json:"email,omitempty"`
    EmailAccessDomainError   *string    `json:"emailAccessDomainError,omitempty"`
    EmailCancel              *string    `json:"emailCancel,omitempty"`
    EmailCodeError           *string    `json:"emailCodeError,omitempty"`
    EmailError               *string    `json:"emailError,omitempty"`
    EmailFieldLabel          *string    `json:"emailFieldLabel,omitempty"`
    EmailLabel               *string    `json:"emailLabel,omitempty"`
    EmailMessage             *string    `json:"emailMessage,omitempty"`
    EmailSubmit              *string    `json:"emailSubmit,omitempty"`
    EmailTitle               *string    `json:"emailTitle,omitempty"`
    Field1                   *bool      `json:"field1,omitempty"`
    Field1Error              *string    `json:"field1Error,omitempty"`
    Field1Label              *string    `json:"field1Label,omitempty"`
    Field1Required           *bool      `json:"field1Required,omitempty"`
    Field2                   *bool      `json:"field2,omitempty"`
    Field2Error              *string    `json:"field2Error,omitempty"`
    Field2Label              *string    `json:"field2Label,omitempty"`
    Field2Required           *bool      `json:"field2Required,omitempty"`
    Field3                   *bool      `json:"field3,omitempty"`
    Field3Error              *string    `json:"field3Error,omitempty"`
    Field3Label              *string    `json:"field3Label,omitempty"`
    Field3Required           *bool      `json:"field3Required,omitempty"`
    Field4                   *bool      `json:"field4,omitempty"`
    Field4Error              *string    `json:"field4Error,omitempty"`
    Field4Label              *string    `json:"field4Label,omitempty"`
    Field4Required           *bool      `json:"field4Required,omitempty"`
    ForSite                  *bool      `json:"for_site,omitempty"`
    Id                       *uuid.UUID `json:"id,omitempty"`
    Message                  *string    `json:"message,omitempty"`
    ModifiedTime             *float64   `json:"modified_time,omitempty"`
    Name                     *bool      `json:"name,omitempty"`
    NameError                *string    `json:"nameError,omitempty"`
    NameLabel                *string    `json:"nameLabel,omitempty"`
    Optout                   *bool      `json:"optout,omitempty"`
    OptoutLabel              *string    `json:"optoutLabel,omitempty"`
    OrgId                    *uuid.UUID `json:"org_id,omitempty"`
    PageTitle                *string    `json:"pageTitle"`
    PassphraseCancel         *string    `json:"passphraseCancel,omitempty"`
    PassphraseError          *string    `json:"passphraseError,omitempty"`
    PassphraseLabel          *string    `json:"passphraseLabel,omitempty"`
    PassphraseMessage        *string    `json:"passphraseMessage,omitempty"`
    PassphraseSubmit         *string    `json:"passphraseSubmit,omitempty"`
    PassphraseTitle          *string    `json:"passphraseTitle,omitempty"`
    PoweredBy                *bool      `json:"poweredBy,omitempty"`
    PrivacyPolicy            *bool      `json:"privacyPolicy,omitempty"`
    PrivacyPolicyAcceptLabel *string    `json:"privacyPolicyAcceptLabel,omitempty"`
    PrivacyPolicyError       *string    `json:"privacyPolicyError,omitempty"`
    PrivacyPolicyLink        *string    `json:"privacyPolicyLink,omitempty"`
    PrivacyPolicyText        *string    `json:"privacyPolicyText,omitempty"`
    RequiredFieldLabel       *string    `json:"requiredFieldLabel,omitempty"`
    SignInLabel              *string    `json:"signInLabel,omitempty"`
    SiteId                   *uuid.UUID `json:"site_id,omitempty"`
    SmsCarrierDefault        *string    `json:"smsCarrierDefault,omitempty"`
    SmsCarrierError          *string    `json:"smsCarrierError,omitempty"`
    SmsCarrierFieldLabel     *string    `json:"smsCarrierFieldLabel,omitempty"`
    SmsCodeCancel            *string    `json:"smsCodeCancel,omitempty"`
    SmsCodeError             *string    `json:"smsCodeError,omitempty"`
    SmsCodeFieldLabel        *string    `json:"smsCodeFieldLabel,omitempty"`
    SmsCodeMessage           *string    `json:"smsCodeMessage,omitempty"`
    SmsCodeSubmit            *string    `json:"smsCodeSubmit,omitempty"`
    SmsCodeTitle             *string    `json:"smsCodeTitle,omitempty"`
    SmsCountryFieldLabel     *string    `json:"smsCountryFieldLabel,omitempty"`
    SmsCountryFormat         *string    `json:"smsCountryFormat,omitempty"`
    SmsHaveAccessCode        *string    `json:"smsHaveAccessCode,omitempty"`
    SmsMessageFormat         *string    `json:"smsMessageFormat,omitempty"`
    SmsNumberCancel          *string    `json:"smsNumberCancel,omitempty"`
    SmsNumberError           *string    `json:"smsNumberError,omitempty"`
    SmsNumberFieldLabel      *string    `json:"smsNumberFieldLabel,omitempty"`
    SmsNumberFormat          *string    `json:"smsNumberFormat,omitempty"`
    SmsNumberMessage         *string    `json:"smsNumberMessage,omitempty"`
    SmsNumberSubmit          *string    `json:"smsNumberSubmit,omitempty"`
    SmsNumberTitle           *string    `json:"smsNumberTitle,omitempty"`
    SmsUsernameFormat        *string    `json:"smsUsernameFormat,omitempty"`
    SmsValidityDuration      *float64   `json:"smsValidityDuration,omitempty"`
    SponsorAutoApproved      *string    `json:"sponsorAutoApproved,omitempty"`
    SponsorAutoApprovedNote  *string    `json:"sponsorAutoApprovedNote,omitempty"`
    SponsorBackLink          *string    `json:"sponsorBackLink,omitempty"`
    SponsorCancel            *string    `json:"sponsorCancel,omitempty"`
    SponsorEmail             *string    `json:"sponsorEmail,omitempty"`
    SponsorEmailError        *string    `json:"sponsorEmailError,omitempty"`
    SponsorEmailTemplate     *string    `json:"sponsorEmailTemplate,omitempty"`
    SponsorInfoApproved      *string    `json:"sponsorInfoApproved,omitempty"`
    SponsorInfoDenied        *string    `json:"sponsorInfoDenied,omitempty"`
    SponsorInfoPending       *string    `json:"sponsorInfoPending,omitempty"`
    SponsorName              *string    `json:"sponsorName,omitempty"`
    SponsorNameError         *string    `json:"sponsorNameError,omitempty"`
    SponsorNotePending       *string    `json:"sponsorNotePending,omitempty"`
    SponsorStatusApproved    *string    `json:"sponsorStatusApproved,omitempty"`
    SponsorStatusDenied      *string    `json:"sponsorStatusDenied,omitempty"`
    SponsorStatusPending     *string    `json:"sponsorStatusPending,omitempty"`
    SponsorSubmit            *string    `json:"sponsorSubmit,omitempty"`
    SponsorsAutoApprovedNote *string    `json:"sponsorsAutoApprovedNote,omitempty"`
    SponsorsError            *string    `json:"sponsorsError,omitempty"`
    Tos                      *bool      `json:"tos,omitempty"`
    TosAcceptLabel           *string    `json:"tosAcceptLabel,omitempty"`
    TosError                 *string    `json:"tosError,omitempty"`
    TosLink                  *string    `json:"tosLink,omitempty"`
    TosText                  *string    `json:"tosText,omitempty"`
}

func (p *tempPortalTemplate) validate() error {
    var errs []string
    if p.PageTitle == nil {
        errs = append(errs, "required field `pageTitle` is missing for type `portal_template`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
