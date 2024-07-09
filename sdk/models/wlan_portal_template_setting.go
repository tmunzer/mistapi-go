package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// WlanPortalTemplateSetting represents a WlanPortalTemplateSetting struct.
// portal template wlan settings
type WlanPortalTemplateSetting struct {
    // “Please provide valid alternate email”
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
    // “Connect to WiFi with”
    AuthLabel                *string        `json:"authLabel,omitempty"`
    // label of the link to go back to /logon
    BackLink                 *string        `json:"backLink,omitempty"`
    // “#1074bc”
    Color                    *string        `json:"color,omitempty"`
    // “#0b5183”
    ColorDark                *string        `json:"colorDark,omitempty"`
    // “#3589c6”
    ColorLight               *string        `json:"colorLight,omitempty"`
    // whether company field is required
    Company                  *bool          `json:"company,omitempty"`
    // error message when company not provided
    CompanyError             *string        `json:"companyError,omitempty"`
    // label of company field
    CompanyLabel             *string        `json:"companyLabel,omitempty"`
    // whether email field is required
    Email                    *bool          `json:"email,omitempty"`
    // error message when a user has valid social login but doesn’t match specified email domains.
    EmailAccessDomainError   *string        `json:"emailAccessDomainError,omitempty"`
    // Label for cancel confirmation code submission using email auth
    EmailCancel              *string        `json:"emailCancel,omitempty"`
    EmailCodeCancel          *string        `json:"emailCodeCancel,omitempty"`
    // “Please provide valid alternate email”
    EmailCodeError           *string        `json:"emailCodeError,omitempty"`
    // “Confirmation Code”
    EmailCodeFieldLabel      *string        `json:"emailCodeFieldLabel,omitempty"`
    // “Enter the access number that was sent to your email address.”
    EmailCodeMessage         *string        `json:"emailCodeMessage,omitempty"`
    // “Sign In
    EmailCodeSubmit          *string        `json:"emailCodeSubmit,omitempty"`
    // “Access Code”
    EmailCodeTitle           *string        `json:"emailCodeTitle,omitempty"`
    // error message when email not provided
    EmailError               *string        `json:"emailError,omitempty"`
    // “Enter your email address”
    EmailFieldLabel          *string        `json:"emailFieldLabel,omitempty"`
    // label of email field
    EmailLabel               *string        `json:"emailLabel,omitempty"`
    // “We will email you an authentication code which you can use to connect to the WiFi network.”
    EmailMessage             *string        `json:"emailMessage,omitempty"`
    // Label for confirmation code submit button using email auth
    EmailSubmit              *string        `json:"emailSubmit,omitempty"`
    // “Sign in with Email”
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
    // “Please enjoy the complimentary Wifi”
    Message                  *string        `json:"message,omitempty"`
    // whether name field is required
    Name                     *bool          `json:"name,omitempty"`
    // error message when name not provided
    NameError                *string        `json:"nameError,omitempty"`
    // label of name field
    NameLabel                *string        `json:"nameLabel,omitempty"`
    // whether to display “Do Not Store My Personal Information”
    Optout                   *bool          `json:"optout,omitempty"`
    // Default value for the \"Do not store\" checkbox"
    OptoutDefault            *bool          `json:"optoutDefault,omitempty"`
    // label for “Do Not Store My Personal Information”
    OptoutLabel              *string        `json:"optoutLabel,omitempty"`
    // “Welcome”
    PageTitle                string         `json:"pageTitle"`
    // “Cancel”
    PassphraseCancel         *string        `json:"passphraseCancel,omitempty"`
    // error message when invalid passphrase is provided
    PassphraseError          *string        `json:"passphraseError,omitempty"`
    // Passphrase
    PassphraseLabel          *string        `json:"passphraseLabel,omitempty"`
    // “Login using passphrase”
    PassphraseMessage        *string        `json:"passphraseMessage,omitempty"`
    // “Sign in”
    PassphraseSubmit         *string        `json:"passphraseSubmit,omitempty"`
    // Title for passphrase details page
    PassphraseTitle          *string        `json:"passphraseTitle,omitempty"`
    // whether to show “Powered by Mist”
    PoweredBy                *bool          `json:"poweredBy,omitempty"`
    // label to denote required field
    RequiredFieldLabel       *string        `json:"requiredFieldLabel,omitempty"`
    // label of the button to /signin
    SignInLabel              *string        `json:"signInLabel,omitempty"`
    // “Please Select”
    SmsCarrierDefault        *string        `json:"smsCarrierDefault,omitempty"`
    // “Please select a mobile carrier”
    SmsCarrierError          *string        `json:"smsCarrierError,omitempty"`
    // label for mobile carrier drop-down list
    SmsCarrierFieldLabel     *string        `json:"smsCarrierFieldLabel,omitempty"`
    // Label for cancel confirmation code submission
    SmsCodeCancel            *string        `json:"smsCodeCancel,omitempty"`
    // error message when confirmation code is invalid
    SmsCodeError             *string        `json:"smsCodeError,omitempty"`
    // “Confirmation Code”
    SmsCodeFieldLabel        *string        `json:"smsCodeFieldLabel,omitempty"`
    // “Enter the confirmation code”
    SmsCodeMessage           *string        `json:"smsCodeMessage,omitempty"`
    // Label for confirmation code submit button
    SmsCodeSubmit            *string        `json:"smsCodeSubmit,omitempty"`
    // “Access Code”
    SmsCodeTitle             *string        `json:"smsCodeTitle,omitempty"`
    // “Country Code”
    SmsCountryFieldLabel     *string        `json:"smsCountryFieldLabel,omitempty"`
    // “+1”
    SmsCountryFormat         *string        `json:"smsCountryFormat,omitempty"`
    // Label for checkbox to specify that the user has access code
    SmsHaveAccessCode        *string        `json:"smsHaveAccessCode,omitempty"`
    // format of access code sms message. {{code}} and {{duration}} are place holders and should be retained as is.
    SmsMessageFormat         *string        `json:"smsMessageFormat,omitempty"`
    // label for canceling mobile details for SMS auth
    SmsNumberCancel          *string        `json:"smsNumberCancel,omitempty"`
    // “Invalid Mobile Number”
    SmsNumberError           *string        `json:"smsNumberError,omitempty"`
    // label for field to provide mobile number
    SmsNumberFieldLabel      *string        `json:"smsNumberFieldLabel,omitempty"`
    // “2125551212 (digits only)”
    SmsNumberFormat          *string        `json:"smsNumberFormat,omitempty"`
    // “We will send an access code to your mobile number which you can use to connect to the WiFi network. Message and data rates may apply.”
    SmsNumberMessage         *string        `json:"smsNumberMessage,omitempty"`
    // label for submit button for code generation
    SmsNumberSubmit          *string        `json:"smsNumberSubmit,omitempty"`
    // Title for phone number details
    SmsNumberTitle           *string        `json:"smsNumberTitle,omitempty"`
    // “username”
    SmsUsernameFormat        *string        `json:"smsUsernameFormat,omitempty"`
    // how long confirmation code should be considered valid (in minutes)
    SmsValidityDuration      *int           `json:"smsValidityDuration,omitempty"`
    // “Go back and edit request form”
    SponsorBackLink          *string        `json:"sponsorBackLink,omitempty"`
    // “Cancel”
    SponsorCancel            *string        `json:"sponsorCancel,omitempty"`
    // label for Sponsor Email
    SponsorEmail             *string        `json:"sponsorEmail,omitempty"`
    // “Please provide valid sponsor email”
    SponsorEmailError        *string        `json:"sponsorEmailError,omitempty"`
    // html template to replace/override default sponsor email template
    // Sponsor Email Template supports following template variables: 
    //  * `approve_url`: Renders URL to approve the request; optionally &minutes=N query param can be appended to change the Authorization period of the guest, where N is a valid integer denoting number of minutes a guest remains authorized
    //  * `deny_url`: Renders URL to reject the request
    //  * `guest_email`: Renders Email ID of the guest
    //  * `guest_name`: Renders Name of the guest
    //  * `field1`: Renders value of the Custom Field 1
    //  * `field2`: Renders value of the Custom Field 2
    //  * `sponsor_link_validity_duration`: Renders validity time of the request (i.e. Approve/Deny URL)
    //  * `auth_expire_minutes`: Renders Wlan-level configured Guest Authorization Expiration time period (in minutes), If not configured then default (1 day in minutes)
    SponsorEmailTemplate     *string        `json:"sponsorEmailTemplate,omitempty"`
    // “Your request was approved by”
    SponsorInfoApproved      *string        `json:"sponsorInfoApproved,omitempty"`
    // “Your request was denied by”
    SponsorInfoDenied        *string        `json:"sponsorInfoDenied,omitempty"`
    // “Your notification has been sent to”
    SponsorInfoPending       *string        `json:"sponsorInfoPending,omitempty"`
    // label for Sponsor Name
    SponsorName              *string        `json:"sponsorName,omitempty"`
    // “Please provide sponsor’s name”
    SponsorNameError         *string        `json:"sponsorNameError,omitempty"`
    // “Please wait for them to acknowledge.”
    SponsorNotePending       *string        `json:"sponsorNotePending,omitempty"`
    // ‘submit button label request Wifi Access and notify sponsor about guest request
    SponsorRequestAccess     *string        `json:"sponsorRequestAccess,omitempty"`
    // “Select Sponsor”
    SponsorSelectEmail       *string        `json:"sponsorSelectEmail,omitempty"`
    // text to display if sponsor approves request
    SponsorStatusApproved    *string        `json:"sponsorStatusApproved,omitempty"`
    // text to display when sponsor denies request
    SponsorStatusDenied      *string        `json:"sponsorStatusDenied,omitempty"`
    // text to display if request is still pending
    SponsorStatusPending     *string        `json:"sponsorStatusPending,omitempty"`
    // submit button label to notify sponsor about guest request
    SponsorSubmit            *string        `json:"sponsorSubmit,omitempty"`
    // “Please select a sponsor”
    SponsorsError            *string        `json:"sponsorsError,omitempty"`
    // “Your request was approved”
    SponsorsInfoApproved     *string        `json:"sponsorsInfoApproved,omitempty"`
    // “Your request was denied”
    SponsorsInfoDenied       *string        `json:"sponsorsInfoDenied,omitempty"`
    // “Your notification has been sent to the sponsors”
    SponsorsInfoPending      *string        `json:"sponsorsInfoPending,omitempty"`
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

// MarshalJSON implements the json.Marshaler interface for WlanPortalTemplateSetting.
// It customizes the JSON marshaling process for WlanPortalTemplateSetting objects.
func (w WlanPortalTemplateSetting) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WlanPortalTemplateSetting object to a map representation for JSON marshaling.
func (w WlanPortalTemplateSetting) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
    if w.AccessCodeAlternateEmail != nil {
        structMap["accessCodeAlternateEmail"] = w.AccessCodeAlternateEmail
    }
    if w.Alignment != nil {
        structMap["alignment"] = w.Alignment
    }
    if w.AuthButtonAmazon != nil {
        structMap["authButtonAmazon"] = w.AuthButtonAmazon
    }
    if w.AuthButtonAzure != nil {
        structMap["authButtonAzure"] = w.AuthButtonAzure
    }
    if w.AuthButtonEmail != nil {
        structMap["authButtonEmail"] = w.AuthButtonEmail
    }
    if w.AuthButtonFacebook != nil {
        structMap["authButtonFacebook"] = w.AuthButtonFacebook
    }
    if w.AuthButtonGoogle != nil {
        structMap["authButtonGoogle"] = w.AuthButtonGoogle
    }
    if w.AuthButtonMicrosoft != nil {
        structMap["authButtonMicrosoft"] = w.AuthButtonMicrosoft
    }
    if w.AuthButtonPassphrase != nil {
        structMap["authButtonPassphrase"] = w.AuthButtonPassphrase
    }
    if w.AuthButtonSms != nil {
        structMap["authButtonSms"] = w.AuthButtonSms
    }
    if w.AuthButtonSponsor != nil {
        structMap["authButtonSponsor"] = w.AuthButtonSponsor
    }
    if w.AuthLabel != nil {
        structMap["authLabel"] = w.AuthLabel
    }
    if w.BackLink != nil {
        structMap["backLink"] = w.BackLink
    }
    if w.Color != nil {
        structMap["color"] = w.Color
    }
    if w.ColorDark != nil {
        structMap["colorDark"] = w.ColorDark
    }
    if w.ColorLight != nil {
        structMap["colorLight"] = w.ColorLight
    }
    if w.Company != nil {
        structMap["company"] = w.Company
    }
    if w.CompanyError != nil {
        structMap["companyError"] = w.CompanyError
    }
    if w.CompanyLabel != nil {
        structMap["companyLabel"] = w.CompanyLabel
    }
    if w.Email != nil {
        structMap["email"] = w.Email
    }
    if w.EmailAccessDomainError != nil {
        structMap["emailAccessDomainError"] = w.EmailAccessDomainError
    }
    if w.EmailCancel != nil {
        structMap["emailCancel"] = w.EmailCancel
    }
    if w.EmailCodeCancel != nil {
        structMap["emailCodeCancel"] = w.EmailCodeCancel
    }
    if w.EmailCodeError != nil {
        structMap["emailCodeError"] = w.EmailCodeError
    }
    if w.EmailCodeFieldLabel != nil {
        structMap["emailCodeFieldLabel"] = w.EmailCodeFieldLabel
    }
    if w.EmailCodeMessage != nil {
        structMap["emailCodeMessage"] = w.EmailCodeMessage
    }
    if w.EmailCodeSubmit != nil {
        structMap["emailCodeSubmit"] = w.EmailCodeSubmit
    }
    if w.EmailCodeTitle != nil {
        structMap["emailCodeTitle"] = w.EmailCodeTitle
    }
    if w.EmailError != nil {
        structMap["emailError"] = w.EmailError
    }
    if w.EmailFieldLabel != nil {
        structMap["emailFieldLabel"] = w.EmailFieldLabel
    }
    if w.EmailLabel != nil {
        structMap["emailLabel"] = w.EmailLabel
    }
    if w.EmailMessage != nil {
        structMap["emailMessage"] = w.EmailMessage
    }
    if w.EmailSubmit != nil {
        structMap["emailSubmit"] = w.EmailSubmit
    }
    if w.EmailTitle != nil {
        structMap["emailTitle"] = w.EmailTitle
    }
    if w.Field1 != nil {
        structMap["field1"] = w.Field1
    }
    if w.Field1Error != nil {
        structMap["field1Error"] = w.Field1Error
    }
    if w.Field1Label != nil {
        structMap["field1Label"] = w.Field1Label
    }
    if w.Field1Required != nil {
        structMap["field1Required"] = w.Field1Required
    }
    if w.Field2 != nil {
        structMap["field2"] = w.Field2
    }
    if w.Field2Error != nil {
        structMap["field2Error"] = w.Field2Error
    }
    if w.Field2Label != nil {
        structMap["field2Label"] = w.Field2Label
    }
    if w.Field2Required != nil {
        structMap["field2Required"] = w.Field2Required
    }
    if w.Field3 != nil {
        structMap["field3"] = w.Field3
    }
    if w.Field3Error != nil {
        structMap["field3Error"] = w.Field3Error
    }
    if w.Field3Label != nil {
        structMap["field3Label"] = w.Field3Label
    }
    if w.Field3Required != nil {
        structMap["field3Required"] = w.Field3Required
    }
    if w.Field4 != nil {
        structMap["field4"] = w.Field4
    }
    if w.Field4Error != nil {
        structMap["field4Error"] = w.Field4Error
    }
    if w.Field4Label != nil {
        structMap["field4Label"] = w.Field4Label
    }
    if w.Field4Required != nil {
        structMap["field4Required"] = w.Field4Required
    }
    if w.Message != nil {
        structMap["message"] = w.Message
    }
    if w.Name != nil {
        structMap["name"] = w.Name
    }
    if w.NameError != nil {
        structMap["nameError"] = w.NameError
    }
    if w.NameLabel != nil {
        structMap["nameLabel"] = w.NameLabel
    }
    if w.Optout != nil {
        structMap["optout"] = w.Optout
    }
    if w.OptoutDefault != nil {
        structMap["optoutDefault"] = w.OptoutDefault
    }
    if w.OptoutLabel != nil {
        structMap["optoutLabel"] = w.OptoutLabel
    }
    structMap["pageTitle"] = w.PageTitle
    if w.PassphraseCancel != nil {
        structMap["passphraseCancel"] = w.PassphraseCancel
    }
    if w.PassphraseError != nil {
        structMap["passphraseError"] = w.PassphraseError
    }
    if w.PassphraseLabel != nil {
        structMap["passphraseLabel"] = w.PassphraseLabel
    }
    if w.PassphraseMessage != nil {
        structMap["passphraseMessage"] = w.PassphraseMessage
    }
    if w.PassphraseSubmit != nil {
        structMap["passphraseSubmit"] = w.PassphraseSubmit
    }
    if w.PassphraseTitle != nil {
        structMap["passphraseTitle"] = w.PassphraseTitle
    }
    if w.PoweredBy != nil {
        structMap["poweredBy"] = w.PoweredBy
    }
    if w.RequiredFieldLabel != nil {
        structMap["requiredFieldLabel"] = w.RequiredFieldLabel
    }
    if w.SignInLabel != nil {
        structMap["signInLabel"] = w.SignInLabel
    }
    if w.SmsCarrierDefault != nil {
        structMap["smsCarrierDefault"] = w.SmsCarrierDefault
    }
    if w.SmsCarrierError != nil {
        structMap["smsCarrierError"] = w.SmsCarrierError
    }
    if w.SmsCarrierFieldLabel != nil {
        structMap["smsCarrierFieldLabel"] = w.SmsCarrierFieldLabel
    }
    if w.SmsCodeCancel != nil {
        structMap["smsCodeCancel"] = w.SmsCodeCancel
    }
    if w.SmsCodeError != nil {
        structMap["smsCodeError"] = w.SmsCodeError
    }
    if w.SmsCodeFieldLabel != nil {
        structMap["smsCodeFieldLabel"] = w.SmsCodeFieldLabel
    }
    if w.SmsCodeMessage != nil {
        structMap["smsCodeMessage"] = w.SmsCodeMessage
    }
    if w.SmsCodeSubmit != nil {
        structMap["smsCodeSubmit"] = w.SmsCodeSubmit
    }
    if w.SmsCodeTitle != nil {
        structMap["smsCodeTitle"] = w.SmsCodeTitle
    }
    if w.SmsCountryFieldLabel != nil {
        structMap["smsCountryFieldLabel"] = w.SmsCountryFieldLabel
    }
    if w.SmsCountryFormat != nil {
        structMap["smsCountryFormat"] = w.SmsCountryFormat
    }
    if w.SmsHaveAccessCode != nil {
        structMap["smsHaveAccessCode"] = w.SmsHaveAccessCode
    }
    if w.SmsMessageFormat != nil {
        structMap["smsMessageFormat"] = w.SmsMessageFormat
    }
    if w.SmsNumberCancel != nil {
        structMap["smsNumberCancel"] = w.SmsNumberCancel
    }
    if w.SmsNumberError != nil {
        structMap["smsNumberError"] = w.SmsNumberError
    }
    if w.SmsNumberFieldLabel != nil {
        structMap["smsNumberFieldLabel"] = w.SmsNumberFieldLabel
    }
    if w.SmsNumberFormat != nil {
        structMap["smsNumberFormat"] = w.SmsNumberFormat
    }
    if w.SmsNumberMessage != nil {
        structMap["smsNumberMessage"] = w.SmsNumberMessage
    }
    if w.SmsNumberSubmit != nil {
        structMap["smsNumberSubmit"] = w.SmsNumberSubmit
    }
    if w.SmsNumberTitle != nil {
        structMap["smsNumberTitle"] = w.SmsNumberTitle
    }
    if w.SmsUsernameFormat != nil {
        structMap["smsUsernameFormat"] = w.SmsUsernameFormat
    }
    if w.SmsValidityDuration != nil {
        structMap["smsValidityDuration"] = w.SmsValidityDuration
    }
    if w.SponsorBackLink != nil {
        structMap["sponsorBackLink"] = w.SponsorBackLink
    }
    if w.SponsorCancel != nil {
        structMap["sponsorCancel"] = w.SponsorCancel
    }
    if w.SponsorEmail != nil {
        structMap["sponsorEmail"] = w.SponsorEmail
    }
    if w.SponsorEmailError != nil {
        structMap["sponsorEmailError"] = w.SponsorEmailError
    }
    if w.SponsorEmailTemplate != nil {
        structMap["sponsorEmailTemplate"] = w.SponsorEmailTemplate
    }
    if w.SponsorInfoApproved != nil {
        structMap["sponsorInfoApproved"] = w.SponsorInfoApproved
    }
    if w.SponsorInfoDenied != nil {
        structMap["sponsorInfoDenied"] = w.SponsorInfoDenied
    }
    if w.SponsorInfoPending != nil {
        structMap["sponsorInfoPending"] = w.SponsorInfoPending
    }
    if w.SponsorName != nil {
        structMap["sponsorName"] = w.SponsorName
    }
    if w.SponsorNameError != nil {
        structMap["sponsorNameError"] = w.SponsorNameError
    }
    if w.SponsorNotePending != nil {
        structMap["sponsorNotePending"] = w.SponsorNotePending
    }
    if w.SponsorRequestAccess != nil {
        structMap["sponsorRequestAccess"] = w.SponsorRequestAccess
    }
    if w.SponsorSelectEmail != nil {
        structMap["sponsorSelectEmail"] = w.SponsorSelectEmail
    }
    if w.SponsorStatusApproved != nil {
        structMap["sponsorStatusApproved"] = w.SponsorStatusApproved
    }
    if w.SponsorStatusDenied != nil {
        structMap["sponsorStatusDenied"] = w.SponsorStatusDenied
    }
    if w.SponsorStatusPending != nil {
        structMap["sponsorStatusPending"] = w.SponsorStatusPending
    }
    if w.SponsorSubmit != nil {
        structMap["sponsorSubmit"] = w.SponsorSubmit
    }
    if w.SponsorsError != nil {
        structMap["sponsorsError"] = w.SponsorsError
    }
    if w.SponsorsInfoApproved != nil {
        structMap["sponsorsInfoApproved"] = w.SponsorsInfoApproved
    }
    if w.SponsorsInfoDenied != nil {
        structMap["sponsorsInfoDenied"] = w.SponsorsInfoDenied
    }
    if w.SponsorsInfoPending != nil {
        structMap["sponsorsInfoPending"] = w.SponsorsInfoPending
    }
    if w.Tos != nil {
        structMap["tos"] = w.Tos
    }
    if w.TosAcceptLabel != nil {
        structMap["tosAcceptLabel"] = w.TosAcceptLabel
    }
    if w.TosError != nil {
        structMap["tosError"] = w.TosError
    }
    if w.TosLink != nil {
        structMap["tosLink"] = w.TosLink
    }
    if w.TosText != nil {
        structMap["tosText"] = w.TosText
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WlanPortalTemplateSetting.
// It customizes the JSON unmarshaling process for WlanPortalTemplateSetting objects.
func (w *WlanPortalTemplateSetting) UnmarshalJSON(input []byte) error {
    var temp wlanPortalTemplateSetting
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "accessCodeAlternateEmail", "alignment", "authButtonAmazon", "authButtonAzure", "authButtonEmail", "authButtonFacebook", "authButtonGoogle", "authButtonMicrosoft", "authButtonPassphrase", "authButtonSms", "authButtonSponsor", "authLabel", "backLink", "color", "colorDark", "colorLight", "company", "companyError", "companyLabel", "email", "emailAccessDomainError", "emailCancel", "emailCodeCancel", "emailCodeError", "emailCodeFieldLabel", "emailCodeMessage", "emailCodeSubmit", "emailCodeTitle", "emailError", "emailFieldLabel", "emailLabel", "emailMessage", "emailSubmit", "emailTitle", "field1", "field1Error", "field1Label", "field1Required", "field2", "field2Error", "field2Label", "field2Required", "field3", "field3Error", "field3Label", "field3Required", "field4", "field4Error", "field4Label", "field4Required", "message", "name", "nameError", "nameLabel", "optout", "optoutDefault", "optoutLabel", "pageTitle", "passphraseCancel", "passphraseError", "passphraseLabel", "passphraseMessage", "passphraseSubmit", "passphraseTitle", "poweredBy", "requiredFieldLabel", "signInLabel", "smsCarrierDefault", "smsCarrierError", "smsCarrierFieldLabel", "smsCodeCancel", "smsCodeError", "smsCodeFieldLabel", "smsCodeMessage", "smsCodeSubmit", "smsCodeTitle", "smsCountryFieldLabel", "smsCountryFormat", "smsHaveAccessCode", "smsMessageFormat", "smsNumberCancel", "smsNumberError", "smsNumberFieldLabel", "smsNumberFormat", "smsNumberMessage", "smsNumberSubmit", "smsNumberTitle", "smsUsernameFormat", "smsValidityDuration", "sponsorBackLink", "sponsorCancel", "sponsorEmail", "sponsorEmailError", "sponsorEmailTemplate", "sponsorInfoApproved", "sponsorInfoDenied", "sponsorInfoPending", "sponsorName", "sponsorNameError", "sponsorNotePending", "sponsorRequestAccess", "sponsorSelectEmail", "sponsorStatusApproved", "sponsorStatusDenied", "sponsorStatusPending", "sponsorSubmit", "sponsorsError", "sponsorsInfoApproved", "sponsorsInfoDenied", "sponsorsInfoPending", "tos", "tosAcceptLabel", "tosError", "tosLink", "tosText")
    if err != nil {
    	return err
    }
    
    w.AdditionalProperties = additionalProperties
    w.AccessCodeAlternateEmail = temp.AccessCodeAlternateEmail
    w.Alignment = temp.Alignment
    w.AuthButtonAmazon = temp.AuthButtonAmazon
    w.AuthButtonAzure = temp.AuthButtonAzure
    w.AuthButtonEmail = temp.AuthButtonEmail
    w.AuthButtonFacebook = temp.AuthButtonFacebook
    w.AuthButtonGoogle = temp.AuthButtonGoogle
    w.AuthButtonMicrosoft = temp.AuthButtonMicrosoft
    w.AuthButtonPassphrase = temp.AuthButtonPassphrase
    w.AuthButtonSms = temp.AuthButtonSms
    w.AuthButtonSponsor = temp.AuthButtonSponsor
    w.AuthLabel = temp.AuthLabel
    w.BackLink = temp.BackLink
    w.Color = temp.Color
    w.ColorDark = temp.ColorDark
    w.ColorLight = temp.ColorLight
    w.Company = temp.Company
    w.CompanyError = temp.CompanyError
    w.CompanyLabel = temp.CompanyLabel
    w.Email = temp.Email
    w.EmailAccessDomainError = temp.EmailAccessDomainError
    w.EmailCancel = temp.EmailCancel
    w.EmailCodeCancel = temp.EmailCodeCancel
    w.EmailCodeError = temp.EmailCodeError
    w.EmailCodeFieldLabel = temp.EmailCodeFieldLabel
    w.EmailCodeMessage = temp.EmailCodeMessage
    w.EmailCodeSubmit = temp.EmailCodeSubmit
    w.EmailCodeTitle = temp.EmailCodeTitle
    w.EmailError = temp.EmailError
    w.EmailFieldLabel = temp.EmailFieldLabel
    w.EmailLabel = temp.EmailLabel
    w.EmailMessage = temp.EmailMessage
    w.EmailSubmit = temp.EmailSubmit
    w.EmailTitle = temp.EmailTitle
    w.Field1 = temp.Field1
    w.Field1Error = temp.Field1Error
    w.Field1Label = temp.Field1Label
    w.Field1Required = temp.Field1Required
    w.Field2 = temp.Field2
    w.Field2Error = temp.Field2Error
    w.Field2Label = temp.Field2Label
    w.Field2Required = temp.Field2Required
    w.Field3 = temp.Field3
    w.Field3Error = temp.Field3Error
    w.Field3Label = temp.Field3Label
    w.Field3Required = temp.Field3Required
    w.Field4 = temp.Field4
    w.Field4Error = temp.Field4Error
    w.Field4Label = temp.Field4Label
    w.Field4Required = temp.Field4Required
    w.Message = temp.Message
    w.Name = temp.Name
    w.NameError = temp.NameError
    w.NameLabel = temp.NameLabel
    w.Optout = temp.Optout
    w.OptoutDefault = temp.OptoutDefault
    w.OptoutLabel = temp.OptoutLabel
    w.PageTitle = *temp.PageTitle
    w.PassphraseCancel = temp.PassphraseCancel
    w.PassphraseError = temp.PassphraseError
    w.PassphraseLabel = temp.PassphraseLabel
    w.PassphraseMessage = temp.PassphraseMessage
    w.PassphraseSubmit = temp.PassphraseSubmit
    w.PassphraseTitle = temp.PassphraseTitle
    w.PoweredBy = temp.PoweredBy
    w.RequiredFieldLabel = temp.RequiredFieldLabel
    w.SignInLabel = temp.SignInLabel
    w.SmsCarrierDefault = temp.SmsCarrierDefault
    w.SmsCarrierError = temp.SmsCarrierError
    w.SmsCarrierFieldLabel = temp.SmsCarrierFieldLabel
    w.SmsCodeCancel = temp.SmsCodeCancel
    w.SmsCodeError = temp.SmsCodeError
    w.SmsCodeFieldLabel = temp.SmsCodeFieldLabel
    w.SmsCodeMessage = temp.SmsCodeMessage
    w.SmsCodeSubmit = temp.SmsCodeSubmit
    w.SmsCodeTitle = temp.SmsCodeTitle
    w.SmsCountryFieldLabel = temp.SmsCountryFieldLabel
    w.SmsCountryFormat = temp.SmsCountryFormat
    w.SmsHaveAccessCode = temp.SmsHaveAccessCode
    w.SmsMessageFormat = temp.SmsMessageFormat
    w.SmsNumberCancel = temp.SmsNumberCancel
    w.SmsNumberError = temp.SmsNumberError
    w.SmsNumberFieldLabel = temp.SmsNumberFieldLabel
    w.SmsNumberFormat = temp.SmsNumberFormat
    w.SmsNumberMessage = temp.SmsNumberMessage
    w.SmsNumberSubmit = temp.SmsNumberSubmit
    w.SmsNumberTitle = temp.SmsNumberTitle
    w.SmsUsernameFormat = temp.SmsUsernameFormat
    w.SmsValidityDuration = temp.SmsValidityDuration
    w.SponsorBackLink = temp.SponsorBackLink
    w.SponsorCancel = temp.SponsorCancel
    w.SponsorEmail = temp.SponsorEmail
    w.SponsorEmailError = temp.SponsorEmailError
    w.SponsorEmailTemplate = temp.SponsorEmailTemplate
    w.SponsorInfoApproved = temp.SponsorInfoApproved
    w.SponsorInfoDenied = temp.SponsorInfoDenied
    w.SponsorInfoPending = temp.SponsorInfoPending
    w.SponsorName = temp.SponsorName
    w.SponsorNameError = temp.SponsorNameError
    w.SponsorNotePending = temp.SponsorNotePending
    w.SponsorRequestAccess = temp.SponsorRequestAccess
    w.SponsorSelectEmail = temp.SponsorSelectEmail
    w.SponsorStatusApproved = temp.SponsorStatusApproved
    w.SponsorStatusDenied = temp.SponsorStatusDenied
    w.SponsorStatusPending = temp.SponsorStatusPending
    w.SponsorSubmit = temp.SponsorSubmit
    w.SponsorsError = temp.SponsorsError
    w.SponsorsInfoApproved = temp.SponsorsInfoApproved
    w.SponsorsInfoDenied = temp.SponsorsInfoDenied
    w.SponsorsInfoPending = temp.SponsorsInfoPending
    w.Tos = temp.Tos
    w.TosAcceptLabel = temp.TosAcceptLabel
    w.TosError = temp.TosError
    w.TosLink = temp.TosLink
    w.TosText = temp.TosText
    return nil
}

// wlanPortalTemplateSetting is a temporary struct used for validating the fields of WlanPortalTemplateSetting.
type wlanPortalTemplateSetting  struct {
    AccessCodeAlternateEmail *string `json:"accessCodeAlternateEmail,omitempty"`
    Alignment                *string `json:"alignment,omitempty"`
    AuthButtonAmazon         *string `json:"authButtonAmazon,omitempty"`
    AuthButtonAzure          *string `json:"authButtonAzure,omitempty"`
    AuthButtonEmail          *string `json:"authButtonEmail,omitempty"`
    AuthButtonFacebook       *string `json:"authButtonFacebook,omitempty"`
    AuthButtonGoogle         *string `json:"authButtonGoogle,omitempty"`
    AuthButtonMicrosoft      *string `json:"authButtonMicrosoft,omitempty"`
    AuthButtonPassphrase     *string `json:"authButtonPassphrase,omitempty"`
    AuthButtonSms            *string `json:"authButtonSms,omitempty"`
    AuthButtonSponsor        *string `json:"authButtonSponsor,omitempty"`
    AuthLabel                *string `json:"authLabel,omitempty"`
    BackLink                 *string `json:"backLink,omitempty"`
    Color                    *string `json:"color,omitempty"`
    ColorDark                *string `json:"colorDark,omitempty"`
    ColorLight               *string `json:"colorLight,omitempty"`
    Company                  *bool   `json:"company,omitempty"`
    CompanyError             *string `json:"companyError,omitempty"`
    CompanyLabel             *string `json:"companyLabel,omitempty"`
    Email                    *bool   `json:"email,omitempty"`
    EmailAccessDomainError   *string `json:"emailAccessDomainError,omitempty"`
    EmailCancel              *string `json:"emailCancel,omitempty"`
    EmailCodeCancel          *string `json:"emailCodeCancel,omitempty"`
    EmailCodeError           *string `json:"emailCodeError,omitempty"`
    EmailCodeFieldLabel      *string `json:"emailCodeFieldLabel,omitempty"`
    EmailCodeMessage         *string `json:"emailCodeMessage,omitempty"`
    EmailCodeSubmit          *string `json:"emailCodeSubmit,omitempty"`
    EmailCodeTitle           *string `json:"emailCodeTitle,omitempty"`
    EmailError               *string `json:"emailError,omitempty"`
    EmailFieldLabel          *string `json:"emailFieldLabel,omitempty"`
    EmailLabel               *string `json:"emailLabel,omitempty"`
    EmailMessage             *string `json:"emailMessage,omitempty"`
    EmailSubmit              *string `json:"emailSubmit,omitempty"`
    EmailTitle               *string `json:"emailTitle,omitempty"`
    Field1                   *bool   `json:"field1,omitempty"`
    Field1Error              *string `json:"field1Error,omitempty"`
    Field1Label              *string `json:"field1Label,omitempty"`
    Field1Required           *bool   `json:"field1Required,omitempty"`
    Field2                   *bool   `json:"field2,omitempty"`
    Field2Error              *string `json:"field2Error,omitempty"`
    Field2Label              *string `json:"field2Label,omitempty"`
    Field2Required           *bool   `json:"field2Required,omitempty"`
    Field3                   *bool   `json:"field3,omitempty"`
    Field3Error              *string `json:"field3Error,omitempty"`
    Field3Label              *string `json:"field3Label,omitempty"`
    Field3Required           *bool   `json:"field3Required,omitempty"`
    Field4                   *bool   `json:"field4,omitempty"`
    Field4Error              *string `json:"field4Error,omitempty"`
    Field4Label              *string `json:"field4Label,omitempty"`
    Field4Required           *bool   `json:"field4Required,omitempty"`
    Message                  *string `json:"message,omitempty"`
    Name                     *bool   `json:"name,omitempty"`
    NameError                *string `json:"nameError,omitempty"`
    NameLabel                *string `json:"nameLabel,omitempty"`
    Optout                   *bool   `json:"optout,omitempty"`
    OptoutDefault            *bool   `json:"optoutDefault,omitempty"`
    OptoutLabel              *string `json:"optoutLabel,omitempty"`
    PageTitle                *string `json:"pageTitle"`
    PassphraseCancel         *string `json:"passphraseCancel,omitempty"`
    PassphraseError          *string `json:"passphraseError,omitempty"`
    PassphraseLabel          *string `json:"passphraseLabel,omitempty"`
    PassphraseMessage        *string `json:"passphraseMessage,omitempty"`
    PassphraseSubmit         *string `json:"passphraseSubmit,omitempty"`
    PassphraseTitle          *string `json:"passphraseTitle,omitempty"`
    PoweredBy                *bool   `json:"poweredBy,omitempty"`
    RequiredFieldLabel       *string `json:"requiredFieldLabel,omitempty"`
    SignInLabel              *string `json:"signInLabel,omitempty"`
    SmsCarrierDefault        *string `json:"smsCarrierDefault,omitempty"`
    SmsCarrierError          *string `json:"smsCarrierError,omitempty"`
    SmsCarrierFieldLabel     *string `json:"smsCarrierFieldLabel,omitempty"`
    SmsCodeCancel            *string `json:"smsCodeCancel,omitempty"`
    SmsCodeError             *string `json:"smsCodeError,omitempty"`
    SmsCodeFieldLabel        *string `json:"smsCodeFieldLabel,omitempty"`
    SmsCodeMessage           *string `json:"smsCodeMessage,omitempty"`
    SmsCodeSubmit            *string `json:"smsCodeSubmit,omitempty"`
    SmsCodeTitle             *string `json:"smsCodeTitle,omitempty"`
    SmsCountryFieldLabel     *string `json:"smsCountryFieldLabel,omitempty"`
    SmsCountryFormat         *string `json:"smsCountryFormat,omitempty"`
    SmsHaveAccessCode        *string `json:"smsHaveAccessCode,omitempty"`
    SmsMessageFormat         *string `json:"smsMessageFormat,omitempty"`
    SmsNumberCancel          *string `json:"smsNumberCancel,omitempty"`
    SmsNumberError           *string `json:"smsNumberError,omitempty"`
    SmsNumberFieldLabel      *string `json:"smsNumberFieldLabel,omitempty"`
    SmsNumberFormat          *string `json:"smsNumberFormat,omitempty"`
    SmsNumberMessage         *string `json:"smsNumberMessage,omitempty"`
    SmsNumberSubmit          *string `json:"smsNumberSubmit,omitempty"`
    SmsNumberTitle           *string `json:"smsNumberTitle,omitempty"`
    SmsUsernameFormat        *string `json:"smsUsernameFormat,omitempty"`
    SmsValidityDuration      *int    `json:"smsValidityDuration,omitempty"`
    SponsorBackLink          *string `json:"sponsorBackLink,omitempty"`
    SponsorCancel            *string `json:"sponsorCancel,omitempty"`
    SponsorEmail             *string `json:"sponsorEmail,omitempty"`
    SponsorEmailError        *string `json:"sponsorEmailError,omitempty"`
    SponsorEmailTemplate     *string `json:"sponsorEmailTemplate,omitempty"`
    SponsorInfoApproved      *string `json:"sponsorInfoApproved,omitempty"`
    SponsorInfoDenied        *string `json:"sponsorInfoDenied,omitempty"`
    SponsorInfoPending       *string `json:"sponsorInfoPending,omitempty"`
    SponsorName              *string `json:"sponsorName,omitempty"`
    SponsorNameError         *string `json:"sponsorNameError,omitempty"`
    SponsorNotePending       *string `json:"sponsorNotePending,omitempty"`
    SponsorRequestAccess     *string `json:"sponsorRequestAccess,omitempty"`
    SponsorSelectEmail       *string `json:"sponsorSelectEmail,omitempty"`
    SponsorStatusApproved    *string `json:"sponsorStatusApproved,omitempty"`
    SponsorStatusDenied      *string `json:"sponsorStatusDenied,omitempty"`
    SponsorStatusPending     *string `json:"sponsorStatusPending,omitempty"`
    SponsorSubmit            *string `json:"sponsorSubmit,omitempty"`
    SponsorsError            *string `json:"sponsorsError,omitempty"`
    SponsorsInfoApproved     *string `json:"sponsorsInfoApproved,omitempty"`
    SponsorsInfoDenied       *string `json:"sponsorsInfoDenied,omitempty"`
    SponsorsInfoPending      *string `json:"sponsorsInfoPending,omitempty"`
    Tos                      *bool   `json:"tos,omitempty"`
    TosAcceptLabel           *string `json:"tosAcceptLabel,omitempty"`
    TosError                 *string `json:"tosError,omitempty"`
    TosLink                  *string `json:"tosLink,omitempty"`
    TosText                  *string `json:"tosText,omitempty"`
}

func (w *wlanPortalTemplateSetting) validate() error {
    var errs []string
    if w.PageTitle == nil {
        errs = append(errs, "required field `pageTitle` is missing for type `Wlan_Portal_Template_Setting`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
