package models

import (
    "encoding/json"
)

// WlanPortalTemplateSettingLocale represents a WlanPortalTemplateSettingLocale struct.
type WlanPortalTemplateSettingLocale struct {
    // label for Amazon auth button
    AuthButtonAmazon         *string                `json:"authButtonAmazon,omitempty"`
    // label for Azure auth button
    AuthButtonAzure          *string                `json:"authButtonAzure,omitempty"`
    // label for Email auth button
    AuthButtonEmail          *string                `json:"authButtonEmail,omitempty"`
    // label for Facebook auth button
    AuthButtonFacebook       *string                `json:"authButtonFacebook,omitempty"`
    // label for Google auth button
    AuthButtonGoogle         *string                `json:"authButtonGoogle,omitempty"`
    // label for Microsoft auth button
    AuthButtonMicrosoft      *string                `json:"authButtonMicrosoft,omitempty"`
    // label for passphrase auth button
    AuthButtonPassphrase     *string                `json:"authButtonPassphrase,omitempty"`
    // label for SMS auth button
    AuthButtonSms            *string                `json:"authButtonSms,omitempty"`
    // label for Sponsor auth button
    AuthButtonSponsor        *string                `json:"authButtonSponsor,omitempty"`
    AuthLabel                *string                `json:"authLabel,omitempty"`
    // label of the link to go back to /logon
    BackLink                 *string                `json:"backLink,omitempty"`
    // error message when company not provided
    CompanyError             *string                `json:"companyError,omitempty"`
    // label of company field
    CompanyLabel             *string                `json:"companyLabel,omitempty"`
    // error message when a user has valid social login but doesn't match specified email domains.
    EmailAccessDomainError   *string                `json:"emailAccessDomainError,omitempty"`
    // Label for cancel confirmation code submission using email auth
    EmailCancel              *string                `json:"emailCancel,omitempty"`
    EmailCodeCancel          *string                `json:"emailCodeCancel,omitempty"`
    EmailCodeError           *string                `json:"emailCodeError,omitempty"`
    EmailCodeFieldLabel      *string                `json:"emailCodeFieldLabel,omitempty"`
    EmailCodeMessage         *string                `json:"emailCodeMessage,omitempty"`
    EmailCodeSubmit          *string                `json:"emailCodeSubmit,omitempty"`
    EmailCodeTitle           *string                `json:"emailCodeTitle,omitempty"`
    // error message when email not provided
    EmailError               *string                `json:"emailError,omitempty"`
    EmailFieldLabel          *string                `json:"emailFieldLabel,omitempty"`
    // label of email field
    EmailLabel               *string                `json:"emailLabel,omitempty"`
    EmailMessage             *string                `json:"emailMessage,omitempty"`
    // Label for confirmation code submit button using email auth
    EmailSubmit              *string                `json:"emailSubmit,omitempty"`
    // Title for the Email registration
    EmailTitle               *string                `json:"emailTitle,omitempty"`
    // error message when field1 not provided
    Field1Error              *string                `json:"field1Error,omitempty"`
    // label of field1
    Field1Label              *string                `json:"field1Label,omitempty"`
    // error message when field2 not provided
    Field2Error              *string                `json:"field2Error,omitempty"`
    // label of field2
    Field2Label              *string                `json:"field2Label,omitempty"`
    // error message when field3 not provided
    Field3Error              *string                `json:"field3Error,omitempty"`
    // label of field3
    Field3Label              *string                `json:"field3Label,omitempty"`
    // error message when field4 not provided
    Field4Error              *string                `json:"field4Error,omitempty"`
    // label of field4
    Field4Label              *string                `json:"field4Label,omitempty"`
    Message                  *string                `json:"message,omitempty"`
    // error message when name not provided
    NameError                *string                `json:"nameError,omitempty"`
    // label of name field
    NameLabel                *string                `json:"nameLabel,omitempty"`
    // label for Do Not Store My Personal Information
    OptoutLabel              *string                `json:"optoutLabel,omitempty"`
    PageTitle                *string                `json:"pageTitle,omitempty"`
    // Label for the Passphrase cancel button
    PassphraseCancel         *string                `json:"passphraseCancel,omitempty"`
    // error message when invalid passphrase is provided
    PassphraseError          *string                `json:"passphraseError,omitempty"`
    // Passphrase
    PassphraseLabel          *string                `json:"passphraseLabel,omitempty"`
    PassphraseMessage        *string                `json:"passphraseMessage,omitempty"`
    // Label for the Passphrase submit button
    PassphraseSubmit         *string                `json:"passphraseSubmit,omitempty"`
    // Title for passphrase details page
    PassphraseTitle          *string                `json:"passphraseTitle,omitempty"`
    // prefix of the label of the link to go to Privacy Policy
    PrivacyPolicyAcceptLabel *string                `json:"privacyPolicyAcceptLabel,omitempty"`
    // error message when Privacy Policy not accepted
    PrivacyPolicyError       *string                `json:"privacyPolicyError,omitempty"`
    // label of the link to go to Privacy Policy
    PrivacyPolicyLink        *string                `json:"privacyPolicyLink,omitempty"`
    // text of the Privacy Policy
    PrivacyPolicyText        *string                `json:"privacyPolicyText,omitempty"`
    // label to denote required field
    RequiredFieldLabel       *string                `json:"requiredFieldLabel,omitempty"`
    // label of the button to /signin
    SignInLabel              *string                `json:"signInLabel,omitempty"`
    SmsCarrierDefault        *string                `json:"smsCarrierDefault,omitempty"`
    SmsCarrierError          *string                `json:"smsCarrierError,omitempty"`
    // label for mobile carrier drop-down list
    SmsCarrierFieldLabel     *string                `json:"smsCarrierFieldLabel,omitempty"`
    // Label for cancel confirmation code submission
    SmsCodeCancel            *string                `json:"smsCodeCancel,omitempty"`
    // error message when confirmation code is invalid
    SmsCodeError             *string                `json:"smsCodeError,omitempty"`
    SmsCodeFieldLabel        *string                `json:"smsCodeFieldLabel,omitempty"`
    SmsCodeMessage           *string                `json:"smsCodeMessage,omitempty"`
    // Label for confirmation code submit button
    SmsCodeSubmit            *string                `json:"smsCodeSubmit,omitempty"`
    SmsCodeTitle             *string                `json:"smsCodeTitle,omitempty"`
    SmsCountryFieldLabel     *string                `json:"smsCountryFieldLabel,omitempty"`
    SmsCountryFormat         *string                `json:"smsCountryFormat,omitempty"`
    // Label for checkbox to specify that the user has access code
    SmsHaveAccessCode        *string                `json:"smsHaveAccessCode,omitempty"`
    // format of access code sms message. {{code}} and {{duration}} are place holders and should be retained as is.
    SmsMessageFormat         *string                `json:"smsMessageFormat,omitempty"`
    // label for canceling mobile details for SMS auth
    SmsNumberCancel          *string                `json:"smsNumberCancel,omitempty"`
    SmsNumberError           *string                `json:"smsNumberError,omitempty"`
    // label for field to provide mobile number
    SmsNumberFieldLabel      *string                `json:"smsNumberFieldLabel,omitempty"`
    SmsNumberFormat          *string                `json:"smsNumberFormat,omitempty"`
    SmsNumberMessage         *string                `json:"smsNumberMessage,omitempty"`
    // label for submit button for code generation
    SmsNumberSubmit          *string                `json:"smsNumberSubmit,omitempty"`
    // Title for phone number details
    SmsNumberTitle           *string                `json:"smsNumberTitle,omitempty"`
    SmsUsernameFormat        *string                `json:"smsUsernameFormat,omitempty"`
    SponsorBackLink          *string                `json:"sponsorBackLink,omitempty"`
    SponsorCancel            *string                `json:"sponsorCancel,omitempty"`
    // label for Sponsor Email
    SponsorEmail             *string                `json:"sponsorEmail,omitempty"`
    SponsorEmailError        *string                `json:"sponsorEmailError,omitempty"`
    SponsorInfoApproved      *string                `json:"sponsorInfoApproved,omitempty"`
    SponsorInfoDenied        *string                `json:"sponsorInfoDenied,omitempty"`
    SponsorInfoPending       *string                `json:"sponsorInfoPending,omitempty"`
    // label for Sponsor Name
    SponsorName              *string                `json:"sponsorName,omitempty"`
    SponsorNameError         *string                `json:"sponsorNameError,omitempty"`
    SponsorNotePending       *string                `json:"sponsorNotePending,omitempty"`
    // submit button label request Wifi Access and notify sponsor about guest request
    SponsorRequestAccess     *string                `json:"sponsorRequestAccess,omitempty"`
    // text to display if sponsor approves request
    SponsorStatusApproved    *string                `json:"sponsorStatusApproved,omitempty"`
    // text to display when sponsor denies request
    SponsorStatusDenied      *string                `json:"sponsorStatusDenied,omitempty"`
    // text to display if request is still pending
    SponsorStatusPending     *string                `json:"sponsorStatusPending,omitempty"`
    // submit button label to notify sponsor about guest request
    SponsorSubmit            *string                `json:"sponsorSubmit,omitempty"`
    SponsorsError            *string                `json:"sponsorsError,omitempty"`
    SponsorsFieldLabel       *string                `json:"sponsorsFieldLabel,omitempty"`
    // prefix of the label of the link to go to tos
    TosAcceptLabel           *string                `json:"tosAcceptLabel,omitempty"`
    // error message when tos not accepted
    TosError                 *string                `json:"tosError,omitempty"`
    // label of the link to go to tos
    TosLink                  *string                `json:"tosLink,omitempty"`
    // text of the Terms of Service
    TosText                  *string                `json:"tosText,omitempty"`
    AdditionalProperties     map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WlanPortalTemplateSettingLocale.
// It customizes the JSON marshaling process for WlanPortalTemplateSettingLocale objects.
func (w WlanPortalTemplateSettingLocale) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "authButtonAmazon", "authButtonAzure", "authButtonEmail", "authButtonFacebook", "authButtonGoogle", "authButtonMicrosoft", "authButtonPassphrase", "authButtonSms", "authButtonSponsor", "authLabel", "backLink", "companyError", "companyLabel", "emailAccessDomainError", "emailCancel", "emailCodeCancel", "emailCodeError", "emailCodeFieldLabel", "emailCodeMessage", "emailCodeSubmit", "emailCodeTitle", "emailError", "emailFieldLabel", "emailLabel", "emailMessage", "emailSubmit", "emailTitle", "field1Error", "field1Label", "field2Error", "field2Label", "field3Error", "field3Label", "field4Error", "field4Label", "message", "nameError", "nameLabel", "optoutLabel", "pageTitle", "passphraseCancel", "passphraseError", "passphraseLabel", "passphraseMessage", "passphraseSubmit", "passphraseTitle", "privacyPolicyAcceptLabel", "privacyPolicyError", "privacyPolicyLink", "privacyPolicyText", "requiredFieldLabel", "signInLabel", "smsCarrierDefault", "smsCarrierError", "smsCarrierFieldLabel", "smsCodeCancel", "smsCodeError", "smsCodeFieldLabel", "smsCodeMessage", "smsCodeSubmit", "smsCodeTitle", "smsCountryFieldLabel", "smsCountryFormat", "smsHaveAccessCode", "smsMessageFormat", "smsNumberCancel", "smsNumberError", "smsNumberFieldLabel", "smsNumberFormat", "smsNumberMessage", "smsNumberSubmit", "smsNumberTitle", "smsUsernameFormat", "sponsorBackLink", "sponsorCancel", "sponsorEmail", "sponsorEmailError", "sponsorInfoApproved", "sponsorInfoDenied", "sponsorInfoPending", "sponsorName", "sponsorNameError", "sponsorNotePending", "sponsorRequestAccess", "sponsorStatusApproved", "sponsorStatusDenied", "sponsorStatusPending", "sponsorSubmit", "sponsorsError", "sponsorsFieldLabel", "tosAcceptLabel", "tosError", "tosLink", "tosText"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WlanPortalTemplateSettingLocale object to a map representation for JSON marshaling.
func (w WlanPortalTemplateSettingLocale) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
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
    if w.CompanyError != nil {
        structMap["companyError"] = w.CompanyError
    }
    if w.CompanyLabel != nil {
        structMap["companyLabel"] = w.CompanyLabel
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
    if w.Field1Error != nil {
        structMap["field1Error"] = w.Field1Error
    }
    if w.Field1Label != nil {
        structMap["field1Label"] = w.Field1Label
    }
    if w.Field2Error != nil {
        structMap["field2Error"] = w.Field2Error
    }
    if w.Field2Label != nil {
        structMap["field2Label"] = w.Field2Label
    }
    if w.Field3Error != nil {
        structMap["field3Error"] = w.Field3Error
    }
    if w.Field3Label != nil {
        structMap["field3Label"] = w.Field3Label
    }
    if w.Field4Error != nil {
        structMap["field4Error"] = w.Field4Error
    }
    if w.Field4Label != nil {
        structMap["field4Label"] = w.Field4Label
    }
    if w.Message != nil {
        structMap["message"] = w.Message
    }
    if w.NameError != nil {
        structMap["nameError"] = w.NameError
    }
    if w.NameLabel != nil {
        structMap["nameLabel"] = w.NameLabel
    }
    if w.OptoutLabel != nil {
        structMap["optoutLabel"] = w.OptoutLabel
    }
    if w.PageTitle != nil {
        structMap["pageTitle"] = w.PageTitle
    }
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
    if w.PrivacyPolicyAcceptLabel != nil {
        structMap["privacyPolicyAcceptLabel"] = w.PrivacyPolicyAcceptLabel
    }
    if w.PrivacyPolicyError != nil {
        structMap["privacyPolicyError"] = w.PrivacyPolicyError
    }
    if w.PrivacyPolicyLink != nil {
        structMap["privacyPolicyLink"] = w.PrivacyPolicyLink
    }
    if w.PrivacyPolicyText != nil {
        structMap["privacyPolicyText"] = w.PrivacyPolicyText
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
    if w.SponsorsFieldLabel != nil {
        structMap["sponsorsFieldLabel"] = w.SponsorsFieldLabel
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

// UnmarshalJSON implements the json.Unmarshaler interface for WlanPortalTemplateSettingLocale.
// It customizes the JSON unmarshaling process for WlanPortalTemplateSettingLocale objects.
func (w *WlanPortalTemplateSettingLocale) UnmarshalJSON(input []byte) error {
    var temp tempWlanPortalTemplateSettingLocale
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "authButtonAmazon", "authButtonAzure", "authButtonEmail", "authButtonFacebook", "authButtonGoogle", "authButtonMicrosoft", "authButtonPassphrase", "authButtonSms", "authButtonSponsor", "authLabel", "backLink", "companyError", "companyLabel", "emailAccessDomainError", "emailCancel", "emailCodeCancel", "emailCodeError", "emailCodeFieldLabel", "emailCodeMessage", "emailCodeSubmit", "emailCodeTitle", "emailError", "emailFieldLabel", "emailLabel", "emailMessage", "emailSubmit", "emailTitle", "field1Error", "field1Label", "field2Error", "field2Label", "field3Error", "field3Label", "field4Error", "field4Label", "message", "nameError", "nameLabel", "optoutLabel", "pageTitle", "passphraseCancel", "passphraseError", "passphraseLabel", "passphraseMessage", "passphraseSubmit", "passphraseTitle", "privacyPolicyAcceptLabel", "privacyPolicyError", "privacyPolicyLink", "privacyPolicyText", "requiredFieldLabel", "signInLabel", "smsCarrierDefault", "smsCarrierError", "smsCarrierFieldLabel", "smsCodeCancel", "smsCodeError", "smsCodeFieldLabel", "smsCodeMessage", "smsCodeSubmit", "smsCodeTitle", "smsCountryFieldLabel", "smsCountryFormat", "smsHaveAccessCode", "smsMessageFormat", "smsNumberCancel", "smsNumberError", "smsNumberFieldLabel", "smsNumberFormat", "smsNumberMessage", "smsNumberSubmit", "smsNumberTitle", "smsUsernameFormat", "sponsorBackLink", "sponsorCancel", "sponsorEmail", "sponsorEmailError", "sponsorInfoApproved", "sponsorInfoDenied", "sponsorInfoPending", "sponsorName", "sponsorNameError", "sponsorNotePending", "sponsorRequestAccess", "sponsorStatusApproved", "sponsorStatusDenied", "sponsorStatusPending", "sponsorSubmit", "sponsorsError", "sponsorsFieldLabel", "tosAcceptLabel", "tosError", "tosLink", "tosText")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
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
    w.CompanyError = temp.CompanyError
    w.CompanyLabel = temp.CompanyLabel
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
    w.Field1Error = temp.Field1Error
    w.Field1Label = temp.Field1Label
    w.Field2Error = temp.Field2Error
    w.Field2Label = temp.Field2Label
    w.Field3Error = temp.Field3Error
    w.Field3Label = temp.Field3Label
    w.Field4Error = temp.Field4Error
    w.Field4Label = temp.Field4Label
    w.Message = temp.Message
    w.NameError = temp.NameError
    w.NameLabel = temp.NameLabel
    w.OptoutLabel = temp.OptoutLabel
    w.PageTitle = temp.PageTitle
    w.PassphraseCancel = temp.PassphraseCancel
    w.PassphraseError = temp.PassphraseError
    w.PassphraseLabel = temp.PassphraseLabel
    w.PassphraseMessage = temp.PassphraseMessage
    w.PassphraseSubmit = temp.PassphraseSubmit
    w.PassphraseTitle = temp.PassphraseTitle
    w.PrivacyPolicyAcceptLabel = temp.PrivacyPolicyAcceptLabel
    w.PrivacyPolicyError = temp.PrivacyPolicyError
    w.PrivacyPolicyLink = temp.PrivacyPolicyLink
    w.PrivacyPolicyText = temp.PrivacyPolicyText
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
    w.SponsorBackLink = temp.SponsorBackLink
    w.SponsorCancel = temp.SponsorCancel
    w.SponsorEmail = temp.SponsorEmail
    w.SponsorEmailError = temp.SponsorEmailError
    w.SponsorInfoApproved = temp.SponsorInfoApproved
    w.SponsorInfoDenied = temp.SponsorInfoDenied
    w.SponsorInfoPending = temp.SponsorInfoPending
    w.SponsorName = temp.SponsorName
    w.SponsorNameError = temp.SponsorNameError
    w.SponsorNotePending = temp.SponsorNotePending
    w.SponsorRequestAccess = temp.SponsorRequestAccess
    w.SponsorStatusApproved = temp.SponsorStatusApproved
    w.SponsorStatusDenied = temp.SponsorStatusDenied
    w.SponsorStatusPending = temp.SponsorStatusPending
    w.SponsorSubmit = temp.SponsorSubmit
    w.SponsorsError = temp.SponsorsError
    w.SponsorsFieldLabel = temp.SponsorsFieldLabel
    w.TosAcceptLabel = temp.TosAcceptLabel
    w.TosError = temp.TosError
    w.TosLink = temp.TosLink
    w.TosText = temp.TosText
    return nil
}

// tempWlanPortalTemplateSettingLocale is a temporary struct used for validating the fields of WlanPortalTemplateSettingLocale.
type tempWlanPortalTemplateSettingLocale  struct {
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
    CompanyError             *string `json:"companyError,omitempty"`
    CompanyLabel             *string `json:"companyLabel,omitempty"`
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
    Field1Error              *string `json:"field1Error,omitempty"`
    Field1Label              *string `json:"field1Label,omitempty"`
    Field2Error              *string `json:"field2Error,omitempty"`
    Field2Label              *string `json:"field2Label,omitempty"`
    Field3Error              *string `json:"field3Error,omitempty"`
    Field3Label              *string `json:"field3Label,omitempty"`
    Field4Error              *string `json:"field4Error,omitempty"`
    Field4Label              *string `json:"field4Label,omitempty"`
    Message                  *string `json:"message,omitempty"`
    NameError                *string `json:"nameError,omitempty"`
    NameLabel                *string `json:"nameLabel,omitempty"`
    OptoutLabel              *string `json:"optoutLabel,omitempty"`
    PageTitle                *string `json:"pageTitle,omitempty"`
    PassphraseCancel         *string `json:"passphraseCancel,omitempty"`
    PassphraseError          *string `json:"passphraseError,omitempty"`
    PassphraseLabel          *string `json:"passphraseLabel,omitempty"`
    PassphraseMessage        *string `json:"passphraseMessage,omitempty"`
    PassphraseSubmit         *string `json:"passphraseSubmit,omitempty"`
    PassphraseTitle          *string `json:"passphraseTitle,omitempty"`
    PrivacyPolicyAcceptLabel *string `json:"privacyPolicyAcceptLabel,omitempty"`
    PrivacyPolicyError       *string `json:"privacyPolicyError,omitempty"`
    PrivacyPolicyLink        *string `json:"privacyPolicyLink,omitempty"`
    PrivacyPolicyText        *string `json:"privacyPolicyText,omitempty"`
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
    SponsorBackLink          *string `json:"sponsorBackLink,omitempty"`
    SponsorCancel            *string `json:"sponsorCancel,omitempty"`
    SponsorEmail             *string `json:"sponsorEmail,omitempty"`
    SponsorEmailError        *string `json:"sponsorEmailError,omitempty"`
    SponsorInfoApproved      *string `json:"sponsorInfoApproved,omitempty"`
    SponsorInfoDenied        *string `json:"sponsorInfoDenied,omitempty"`
    SponsorInfoPending       *string `json:"sponsorInfoPending,omitempty"`
    SponsorName              *string `json:"sponsorName,omitempty"`
    SponsorNameError         *string `json:"sponsorNameError,omitempty"`
    SponsorNotePending       *string `json:"sponsorNotePending,omitempty"`
    SponsorRequestAccess     *string `json:"sponsorRequestAccess,omitempty"`
    SponsorStatusApproved    *string `json:"sponsorStatusApproved,omitempty"`
    SponsorStatusDenied      *string `json:"sponsorStatusDenied,omitempty"`
    SponsorStatusPending     *string `json:"sponsorStatusPending,omitempty"`
    SponsorSubmit            *string `json:"sponsorSubmit,omitempty"`
    SponsorsError            *string `json:"sponsorsError,omitempty"`
    SponsorsFieldLabel       *string `json:"sponsorsFieldLabel,omitempty"`
    TosAcceptLabel           *string `json:"tosAcceptLabel,omitempty"`
    TosError                 *string `json:"tosError,omitempty"`
    TosLink                  *string `json:"tosLink,omitempty"`
    TosText                  *string `json:"tosText,omitempty"`
}
