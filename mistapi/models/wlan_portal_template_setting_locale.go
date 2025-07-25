// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// WlanPortalTemplateSettingLocale represents a WlanPortalTemplateSettingLocale struct.
type WlanPortalTemplateSettingLocale struct {
    // Label for Amazon auth button
    AuthButtonAmazon          *string                `json:"authButtonAmazon,omitempty"`
    // Label for Azure auth button
    AuthButtonAzure           *string                `json:"authButtonAzure,omitempty"`
    // Label for Email auth button
    AuthButtonEmail           *string                `json:"authButtonEmail,omitempty"`
    // Label for Facebook auth button
    AuthButtonFacebook        *string                `json:"authButtonFacebook,omitempty"`
    // Label for Google auth button
    AuthButtonGoogle          *string                `json:"authButtonGoogle,omitempty"`
    // Label for Microsoft auth button
    AuthButtonMicrosoft       *string                `json:"authButtonMicrosoft,omitempty"`
    // Label for passphrase auth button
    AuthButtonPassphrase      *string                `json:"authButtonPassphrase,omitempty"`
    // Label for SMS auth button
    AuthButtonSms             *string                `json:"authButtonSms,omitempty"`
    // Label for Sponsor auth button
    AuthButtonSponsor         *string                `json:"authButtonSponsor,omitempty"`
    AuthLabel                 *string                `json:"authLabel,omitempty"`
    // Label of the link to go back to /logon
    BackLink                  *string                `json:"backLink,omitempty"`
    // Error message when company not provided
    CompanyError              *string                `json:"companyError,omitempty"`
    // Label of company field
    CompanyLabel              *string                `json:"companyLabel,omitempty"`
    // Error message when a user has valid social login but doesn't match specified email domains.
    EmailAccessDomainError    *string                `json:"emailAccessDomainError,omitempty"`
    // Label for cancel confirmation code submission using email auth
    EmailCancel               *string                `json:"emailCancel,omitempty"`
    EmailCodeCancel           *string                `json:"emailCodeCancel,omitempty"`
    EmailCodeError            *string                `json:"emailCodeError,omitempty"`
    EmailCodeFieldLabel       *string                `json:"emailCodeFieldLabel,omitempty"`
    EmailCodeMessage          *string                `json:"emailCodeMessage,omitempty"`
    EmailCodeSubmit           *string                `json:"emailCodeSubmit,omitempty"`
    EmailCodeTitle            *string                `json:"emailCodeTitle,omitempty"`
    // Error message when email not provided
    EmailError                *string                `json:"emailError,omitempty"`
    EmailFieldLabel           *string                `json:"emailFieldLabel,omitempty"`
    // Label of email field
    EmailLabel                *string                `json:"emailLabel,omitempty"`
    EmailMessage              *string                `json:"emailMessage,omitempty"`
    // Label for confirmation code submit button using email auth
    EmailSubmit               *string                `json:"emailSubmit,omitempty"`
    // Title for the Email registration
    EmailTitle                *string                `json:"emailTitle,omitempty"`
    // Error message when field1 not provided
    Field1Error               *string                `json:"field1Error,omitempty"`
    // Label of field1
    Field1Label               *string                `json:"field1Label,omitempty"`
    // Error message when field2 not provided
    Field2Error               *string                `json:"field2Error,omitempty"`
    // Label of field2
    Field2Label               *string                `json:"field2Label,omitempty"`
    // Error message when field3 not provided
    Field3Error               *string                `json:"field3Error,omitempty"`
    // Label of field3
    Field3Label               *string                `json:"field3Label,omitempty"`
    // Error message when field4 not provided
    Field4Error               *string                `json:"field4Error,omitempty"`
    // Label of field4
    Field4Label               *string                `json:"field4Label,omitempty"`
    // label of the link to go to /marketing_policy
    MarketingPolicyLink       *string                `json:"marketingPolicyLink,omitempty"`
    // Whether marketing policy optin is enabled
    MarketingPolicyOptIn      *bool                  `json:"marketingPolicyOptIn,omitempty"`
    // label for marketing optin
    MarketingPolicyOptInLabel *string                `json:"marketingPolicyOptInLabel,omitempty"`
    // marketing policy text
    MarketingPolicyOptInText  *string                `json:"marketingPolicyOptInText,omitempty"`
    Message                   *string                `json:"message,omitempty"`
    // Error message when name not provided
    NameError                 *string                `json:"nameError,omitempty"`
    // Label of name field
    NameLabel                 *string                `json:"nameLabel,omitempty"`
    // Label for Do Not Store My Personal Information
    OptoutLabel               *string                `json:"optoutLabel,omitempty"`
    PageTitle                 *string                `json:"pageTitle,omitempty"`
    // Label for the Passphrase cancel button
    PassphraseCancel          *string                `json:"passphraseCancel,omitempty"`
    // Error message when invalid passphrase is provided
    PassphraseError           *string                `json:"passphraseError,omitempty"`
    // Passphrase
    PassphraseLabel           *string                `json:"passphraseLabel,omitempty"`
    PassphraseMessage         *string                `json:"passphraseMessage,omitempty"`
    // Label for the Passphrase submit button
    PassphraseSubmit          *string                `json:"passphraseSubmit,omitempty"`
    // Title for passphrase details page
    PassphraseTitle           *string                `json:"passphraseTitle,omitempty"`
    // Prefix of the label of the link to go to Privacy Policy
    PrivacyPolicyAcceptLabel  *string                `json:"privacyPolicyAcceptLabel,omitempty"`
    // Error message when Privacy Policy not accepted
    PrivacyPolicyError        *string                `json:"privacyPolicyError,omitempty"`
    // Label of the link to go to Privacy Policy
    PrivacyPolicyLink         *string                `json:"privacyPolicyLink,omitempty"`
    // Text of the Privacy Policy
    PrivacyPolicyText         *string                `json:"privacyPolicyText,omitempty"`
    // Label to denote required field
    RequiredFieldLabel        *string                `json:"requiredFieldLabel,omitempty"`
    // Label of the button to signin
    SignInLabel               *string                `json:"signInLabel,omitempty"`
    SmsCarrierDefault         *string                `json:"smsCarrierDefault,omitempty"`
    SmsCarrierError           *string                `json:"smsCarrierError,omitempty"`
    // Label for mobile carrier drop-down list
    SmsCarrierFieldLabel      *string                `json:"smsCarrierFieldLabel,omitempty"`
    // Label for cancel confirmation code submission
    SmsCodeCancel             *string                `json:"smsCodeCancel,omitempty"`
    // Error message when confirmation code is invalid
    SmsCodeError              *string                `json:"smsCodeError,omitempty"`
    SmsCodeFieldLabel         *string                `json:"smsCodeFieldLabel,omitempty"`
    SmsCodeMessage            *string                `json:"smsCodeMessage,omitempty"`
    // Label for confirmation code submit button
    SmsCodeSubmit             *string                `json:"smsCodeSubmit,omitempty"`
    SmsCodeTitle              *string                `json:"smsCodeTitle,omitempty"`
    SmsCountryFieldLabel      *string                `json:"smsCountryFieldLabel,omitempty"`
    SmsCountryFormat          *string                `json:"smsCountryFormat,omitempty"`
    // Label for checkbox to specify that the user has access code
    SmsHaveAccessCode         *string                `json:"smsHaveAccessCode,omitempty"`
    // Format of access code sms message. {{code}} and {{duration}} are placeholders and should be retained as is.
    SmsMessageFormat          *string                `json:"smsMessageFormat,omitempty"`
    // Label for canceling mobile details for SMS auth
    SmsNumberCancel           *string                `json:"smsNumberCancel,omitempty"`
    SmsNumberError            *string                `json:"smsNumberError,omitempty"`
    // Label for field to provide mobile number
    SmsNumberFieldLabel       *string                `json:"smsNumberFieldLabel,omitempty"`
    SmsNumberFormat           *string                `json:"smsNumberFormat,omitempty"`
    SmsNumberMessage          *string                `json:"smsNumberMessage,omitempty"`
    // Label for submit button for code generation
    SmsNumberSubmit           *string                `json:"smsNumberSubmit,omitempty"`
    // Title for phone number details
    SmsNumberTitle            *string                `json:"smsNumberTitle,omitempty"`
    SmsUsernameFormat         *string                `json:"smsUsernameFormat,omitempty"`
    SponsorBackLink           *string                `json:"sponsorBackLink,omitempty"`
    SponsorCancel             *string                `json:"sponsorCancel,omitempty"`
    // Label for Sponsor Email
    SponsorEmail              *string                `json:"sponsorEmail,omitempty"`
    SponsorEmailError         *string                `json:"sponsorEmailError,omitempty"`
    SponsorInfoApproved       *string                `json:"sponsorInfoApproved,omitempty"`
    SponsorInfoDenied         *string                `json:"sponsorInfoDenied,omitempty"`
    SponsorInfoPending        *string                `json:"sponsorInfoPending,omitempty"`
    // Label for Sponsor Name
    SponsorName               *string                `json:"sponsorName,omitempty"`
    SponsorNameError          *string                `json:"sponsorNameError,omitempty"`
    SponsorNotePending        *string                `json:"sponsorNotePending,omitempty"`
    // Submit button label request Wifi Access and notify sponsor about guest request
    SponsorRequestAccess      *string                `json:"sponsorRequestAccess,omitempty"`
    // Text to display if sponsor approves request
    SponsorStatusApproved     *string                `json:"sponsorStatusApproved,omitempty"`
    // Text to display when sponsor denies request
    SponsorStatusDenied       *string                `json:"sponsorStatusDenied,omitempty"`
    // Text to display if request is still pending
    SponsorStatusPending      *string                `json:"sponsorStatusPending,omitempty"`
    // Submit button label to notify sponsor about guest request
    SponsorSubmit             *string                `json:"sponsorSubmit,omitempty"`
    SponsorsError             *string                `json:"sponsorsError,omitempty"`
    SponsorsFieldLabel        *string                `json:"sponsorsFieldLabel,omitempty"`
    // Prefix of the label of the link to go to tos
    TosAcceptLabel            *string                `json:"tosAcceptLabel,omitempty"`
    // Error message when tos not accepted
    TosError                  *string                `json:"tosError,omitempty"`
    // Label of the link to go to tos
    TosLink                   *string                `json:"tosLink,omitempty"`
    // Text of the Terms of Service
    TosText                   *string                `json:"tosText,omitempty"`
    AdditionalProperties      map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WlanPortalTemplateSettingLocale,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WlanPortalTemplateSettingLocale) String() string {
    return fmt.Sprintf(
    	"WlanPortalTemplateSettingLocale[AuthButtonAmazon=%v, AuthButtonAzure=%v, AuthButtonEmail=%v, AuthButtonFacebook=%v, AuthButtonGoogle=%v, AuthButtonMicrosoft=%v, AuthButtonPassphrase=%v, AuthButtonSms=%v, AuthButtonSponsor=%v, AuthLabel=%v, BackLink=%v, CompanyError=%v, CompanyLabel=%v, EmailAccessDomainError=%v, EmailCancel=%v, EmailCodeCancel=%v, EmailCodeError=%v, EmailCodeFieldLabel=%v, EmailCodeMessage=%v, EmailCodeSubmit=%v, EmailCodeTitle=%v, EmailError=%v, EmailFieldLabel=%v, EmailLabel=%v, EmailMessage=%v, EmailSubmit=%v, EmailTitle=%v, Field1Error=%v, Field1Label=%v, Field2Error=%v, Field2Label=%v, Field3Error=%v, Field3Label=%v, Field4Error=%v, Field4Label=%v, MarketingPolicyLink=%v, MarketingPolicyOptIn=%v, MarketingPolicyOptInLabel=%v, MarketingPolicyOptInText=%v, Message=%v, NameError=%v, NameLabel=%v, OptoutLabel=%v, PageTitle=%v, PassphraseCancel=%v, PassphraseError=%v, PassphraseLabel=%v, PassphraseMessage=%v, PassphraseSubmit=%v, PassphraseTitle=%v, PrivacyPolicyAcceptLabel=%v, PrivacyPolicyError=%v, PrivacyPolicyLink=%v, PrivacyPolicyText=%v, RequiredFieldLabel=%v, SignInLabel=%v, SmsCarrierDefault=%v, SmsCarrierError=%v, SmsCarrierFieldLabel=%v, SmsCodeCancel=%v, SmsCodeError=%v, SmsCodeFieldLabel=%v, SmsCodeMessage=%v, SmsCodeSubmit=%v, SmsCodeTitle=%v, SmsCountryFieldLabel=%v, SmsCountryFormat=%v, SmsHaveAccessCode=%v, SmsMessageFormat=%v, SmsNumberCancel=%v, SmsNumberError=%v, SmsNumberFieldLabel=%v, SmsNumberFormat=%v, SmsNumberMessage=%v, SmsNumberSubmit=%v, SmsNumberTitle=%v, SmsUsernameFormat=%v, SponsorBackLink=%v, SponsorCancel=%v, SponsorEmail=%v, SponsorEmailError=%v, SponsorInfoApproved=%v, SponsorInfoDenied=%v, SponsorInfoPending=%v, SponsorName=%v, SponsorNameError=%v, SponsorNotePending=%v, SponsorRequestAccess=%v, SponsorStatusApproved=%v, SponsorStatusDenied=%v, SponsorStatusPending=%v, SponsorSubmit=%v, SponsorsError=%v, SponsorsFieldLabel=%v, TosAcceptLabel=%v, TosError=%v, TosLink=%v, TosText=%v, AdditionalProperties=%v]",
    	w.AuthButtonAmazon, w.AuthButtonAzure, w.AuthButtonEmail, w.AuthButtonFacebook, w.AuthButtonGoogle, w.AuthButtonMicrosoft, w.AuthButtonPassphrase, w.AuthButtonSms, w.AuthButtonSponsor, w.AuthLabel, w.BackLink, w.CompanyError, w.CompanyLabel, w.EmailAccessDomainError, w.EmailCancel, w.EmailCodeCancel, w.EmailCodeError, w.EmailCodeFieldLabel, w.EmailCodeMessage, w.EmailCodeSubmit, w.EmailCodeTitle, w.EmailError, w.EmailFieldLabel, w.EmailLabel, w.EmailMessage, w.EmailSubmit, w.EmailTitle, w.Field1Error, w.Field1Label, w.Field2Error, w.Field2Label, w.Field3Error, w.Field3Label, w.Field4Error, w.Field4Label, w.MarketingPolicyLink, w.MarketingPolicyOptIn, w.MarketingPolicyOptInLabel, w.MarketingPolicyOptInText, w.Message, w.NameError, w.NameLabel, w.OptoutLabel, w.PageTitle, w.PassphraseCancel, w.PassphraseError, w.PassphraseLabel, w.PassphraseMessage, w.PassphraseSubmit, w.PassphraseTitle, w.PrivacyPolicyAcceptLabel, w.PrivacyPolicyError, w.PrivacyPolicyLink, w.PrivacyPolicyText, w.RequiredFieldLabel, w.SignInLabel, w.SmsCarrierDefault, w.SmsCarrierError, w.SmsCarrierFieldLabel, w.SmsCodeCancel, w.SmsCodeError, w.SmsCodeFieldLabel, w.SmsCodeMessage, w.SmsCodeSubmit, w.SmsCodeTitle, w.SmsCountryFieldLabel, w.SmsCountryFormat, w.SmsHaveAccessCode, w.SmsMessageFormat, w.SmsNumberCancel, w.SmsNumberError, w.SmsNumberFieldLabel, w.SmsNumberFormat, w.SmsNumberMessage, w.SmsNumberSubmit, w.SmsNumberTitle, w.SmsUsernameFormat, w.SponsorBackLink, w.SponsorCancel, w.SponsorEmail, w.SponsorEmailError, w.SponsorInfoApproved, w.SponsorInfoDenied, w.SponsorInfoPending, w.SponsorName, w.SponsorNameError, w.SponsorNotePending, w.SponsorRequestAccess, w.SponsorStatusApproved, w.SponsorStatusDenied, w.SponsorStatusPending, w.SponsorSubmit, w.SponsorsError, w.SponsorsFieldLabel, w.TosAcceptLabel, w.TosError, w.TosLink, w.TosText, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WlanPortalTemplateSettingLocale.
// It customizes the JSON marshaling process for WlanPortalTemplateSettingLocale objects.
func (w WlanPortalTemplateSettingLocale) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "authButtonAmazon", "authButtonAzure", "authButtonEmail", "authButtonFacebook", "authButtonGoogle", "authButtonMicrosoft", "authButtonPassphrase", "authButtonSms", "authButtonSponsor", "authLabel", "backLink", "companyError", "companyLabel", "emailAccessDomainError", "emailCancel", "emailCodeCancel", "emailCodeError", "emailCodeFieldLabel", "emailCodeMessage", "emailCodeSubmit", "emailCodeTitle", "emailError", "emailFieldLabel", "emailLabel", "emailMessage", "emailSubmit", "emailTitle", "field1Error", "field1Label", "field2Error", "field2Label", "field3Error", "field3Label", "field4Error", "field4Label", "marketingPolicyLink", "marketingPolicyOptIn", "marketingPolicyOptInLabel", "marketingPolicyOptInText", "message", "nameError", "nameLabel", "optoutLabel", "pageTitle", "passphraseCancel", "passphraseError", "passphraseLabel", "passphraseMessage", "passphraseSubmit", "passphraseTitle", "privacyPolicyAcceptLabel", "privacyPolicyError", "privacyPolicyLink", "privacyPolicyText", "requiredFieldLabel", "signInLabel", "smsCarrierDefault", "smsCarrierError", "smsCarrierFieldLabel", "smsCodeCancel", "smsCodeError", "smsCodeFieldLabel", "smsCodeMessage", "smsCodeSubmit", "smsCodeTitle", "smsCountryFieldLabel", "smsCountryFormat", "smsHaveAccessCode", "smsMessageFormat", "smsNumberCancel", "smsNumberError", "smsNumberFieldLabel", "smsNumberFormat", "smsNumberMessage", "smsNumberSubmit", "smsNumberTitle", "smsUsernameFormat", "sponsorBackLink", "sponsorCancel", "sponsorEmail", "sponsorEmailError", "sponsorInfoApproved", "sponsorInfoDenied", "sponsorInfoPending", "sponsorName", "sponsorNameError", "sponsorNotePending", "sponsorRequestAccess", "sponsorStatusApproved", "sponsorStatusDenied", "sponsorStatusPending", "sponsorSubmit", "sponsorsError", "sponsorsFieldLabel", "tosAcceptLabel", "tosError", "tosLink", "tosText"); err != nil {
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
    if w.MarketingPolicyLink != nil {
        structMap["marketingPolicyLink"] = w.MarketingPolicyLink
    }
    if w.MarketingPolicyOptIn != nil {
        structMap["marketingPolicyOptIn"] = w.MarketingPolicyOptIn
    }
    if w.MarketingPolicyOptInLabel != nil {
        structMap["marketingPolicyOptInLabel"] = w.MarketingPolicyOptInLabel
    }
    if w.MarketingPolicyOptInText != nil {
        structMap["marketingPolicyOptInText"] = w.MarketingPolicyOptInText
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "authButtonAmazon", "authButtonAzure", "authButtonEmail", "authButtonFacebook", "authButtonGoogle", "authButtonMicrosoft", "authButtonPassphrase", "authButtonSms", "authButtonSponsor", "authLabel", "backLink", "companyError", "companyLabel", "emailAccessDomainError", "emailCancel", "emailCodeCancel", "emailCodeError", "emailCodeFieldLabel", "emailCodeMessage", "emailCodeSubmit", "emailCodeTitle", "emailError", "emailFieldLabel", "emailLabel", "emailMessage", "emailSubmit", "emailTitle", "field1Error", "field1Label", "field2Error", "field2Label", "field3Error", "field3Label", "field4Error", "field4Label", "marketingPolicyLink", "marketingPolicyOptIn", "marketingPolicyOptInLabel", "marketingPolicyOptInText", "message", "nameError", "nameLabel", "optoutLabel", "pageTitle", "passphraseCancel", "passphraseError", "passphraseLabel", "passphraseMessage", "passphraseSubmit", "passphraseTitle", "privacyPolicyAcceptLabel", "privacyPolicyError", "privacyPolicyLink", "privacyPolicyText", "requiredFieldLabel", "signInLabel", "smsCarrierDefault", "smsCarrierError", "smsCarrierFieldLabel", "smsCodeCancel", "smsCodeError", "smsCodeFieldLabel", "smsCodeMessage", "smsCodeSubmit", "smsCodeTitle", "smsCountryFieldLabel", "smsCountryFormat", "smsHaveAccessCode", "smsMessageFormat", "smsNumberCancel", "smsNumberError", "smsNumberFieldLabel", "smsNumberFormat", "smsNumberMessage", "smsNumberSubmit", "smsNumberTitle", "smsUsernameFormat", "sponsorBackLink", "sponsorCancel", "sponsorEmail", "sponsorEmailError", "sponsorInfoApproved", "sponsorInfoDenied", "sponsorInfoPending", "sponsorName", "sponsorNameError", "sponsorNotePending", "sponsorRequestAccess", "sponsorStatusApproved", "sponsorStatusDenied", "sponsorStatusPending", "sponsorSubmit", "sponsorsError", "sponsorsFieldLabel", "tosAcceptLabel", "tosError", "tosLink", "tosText")
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
    w.MarketingPolicyLink = temp.MarketingPolicyLink
    w.MarketingPolicyOptIn = temp.MarketingPolicyOptIn
    w.MarketingPolicyOptInLabel = temp.MarketingPolicyOptInLabel
    w.MarketingPolicyOptInText = temp.MarketingPolicyOptInText
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
    AuthButtonAmazon          *string `json:"authButtonAmazon,omitempty"`
    AuthButtonAzure           *string `json:"authButtonAzure,omitempty"`
    AuthButtonEmail           *string `json:"authButtonEmail,omitempty"`
    AuthButtonFacebook        *string `json:"authButtonFacebook,omitempty"`
    AuthButtonGoogle          *string `json:"authButtonGoogle,omitempty"`
    AuthButtonMicrosoft       *string `json:"authButtonMicrosoft,omitempty"`
    AuthButtonPassphrase      *string `json:"authButtonPassphrase,omitempty"`
    AuthButtonSms             *string `json:"authButtonSms,omitempty"`
    AuthButtonSponsor         *string `json:"authButtonSponsor,omitempty"`
    AuthLabel                 *string `json:"authLabel,omitempty"`
    BackLink                  *string `json:"backLink,omitempty"`
    CompanyError              *string `json:"companyError,omitempty"`
    CompanyLabel              *string `json:"companyLabel,omitempty"`
    EmailAccessDomainError    *string `json:"emailAccessDomainError,omitempty"`
    EmailCancel               *string `json:"emailCancel,omitempty"`
    EmailCodeCancel           *string `json:"emailCodeCancel,omitempty"`
    EmailCodeError            *string `json:"emailCodeError,omitempty"`
    EmailCodeFieldLabel       *string `json:"emailCodeFieldLabel,omitempty"`
    EmailCodeMessage          *string `json:"emailCodeMessage,omitempty"`
    EmailCodeSubmit           *string `json:"emailCodeSubmit,omitempty"`
    EmailCodeTitle            *string `json:"emailCodeTitle,omitempty"`
    EmailError                *string `json:"emailError,omitempty"`
    EmailFieldLabel           *string `json:"emailFieldLabel,omitempty"`
    EmailLabel                *string `json:"emailLabel,omitempty"`
    EmailMessage              *string `json:"emailMessage,omitempty"`
    EmailSubmit               *string `json:"emailSubmit,omitempty"`
    EmailTitle                *string `json:"emailTitle,omitempty"`
    Field1Error               *string `json:"field1Error,omitempty"`
    Field1Label               *string `json:"field1Label,omitempty"`
    Field2Error               *string `json:"field2Error,omitempty"`
    Field2Label               *string `json:"field2Label,omitempty"`
    Field3Error               *string `json:"field3Error,omitempty"`
    Field3Label               *string `json:"field3Label,omitempty"`
    Field4Error               *string `json:"field4Error,omitempty"`
    Field4Label               *string `json:"field4Label,omitempty"`
    MarketingPolicyLink       *string `json:"marketingPolicyLink,omitempty"`
    MarketingPolicyOptIn      *bool   `json:"marketingPolicyOptIn,omitempty"`
    MarketingPolicyOptInLabel *string `json:"marketingPolicyOptInLabel,omitempty"`
    MarketingPolicyOptInText  *string `json:"marketingPolicyOptInText,omitempty"`
    Message                   *string `json:"message,omitempty"`
    NameError                 *string `json:"nameError,omitempty"`
    NameLabel                 *string `json:"nameLabel,omitempty"`
    OptoutLabel               *string `json:"optoutLabel,omitempty"`
    PageTitle                 *string `json:"pageTitle,omitempty"`
    PassphraseCancel          *string `json:"passphraseCancel,omitempty"`
    PassphraseError           *string `json:"passphraseError,omitempty"`
    PassphraseLabel           *string `json:"passphraseLabel,omitempty"`
    PassphraseMessage         *string `json:"passphraseMessage,omitempty"`
    PassphraseSubmit          *string `json:"passphraseSubmit,omitempty"`
    PassphraseTitle           *string `json:"passphraseTitle,omitempty"`
    PrivacyPolicyAcceptLabel  *string `json:"privacyPolicyAcceptLabel,omitempty"`
    PrivacyPolicyError        *string `json:"privacyPolicyError,omitempty"`
    PrivacyPolicyLink         *string `json:"privacyPolicyLink,omitempty"`
    PrivacyPolicyText         *string `json:"privacyPolicyText,omitempty"`
    RequiredFieldLabel        *string `json:"requiredFieldLabel,omitempty"`
    SignInLabel               *string `json:"signInLabel,omitempty"`
    SmsCarrierDefault         *string `json:"smsCarrierDefault,omitempty"`
    SmsCarrierError           *string `json:"smsCarrierError,omitempty"`
    SmsCarrierFieldLabel      *string `json:"smsCarrierFieldLabel,omitempty"`
    SmsCodeCancel             *string `json:"smsCodeCancel,omitempty"`
    SmsCodeError              *string `json:"smsCodeError,omitempty"`
    SmsCodeFieldLabel         *string `json:"smsCodeFieldLabel,omitempty"`
    SmsCodeMessage            *string `json:"smsCodeMessage,omitempty"`
    SmsCodeSubmit             *string `json:"smsCodeSubmit,omitempty"`
    SmsCodeTitle              *string `json:"smsCodeTitle,omitempty"`
    SmsCountryFieldLabel      *string `json:"smsCountryFieldLabel,omitempty"`
    SmsCountryFormat          *string `json:"smsCountryFormat,omitempty"`
    SmsHaveAccessCode         *string `json:"smsHaveAccessCode,omitempty"`
    SmsMessageFormat          *string `json:"smsMessageFormat,omitempty"`
    SmsNumberCancel           *string `json:"smsNumberCancel,omitempty"`
    SmsNumberError            *string `json:"smsNumberError,omitempty"`
    SmsNumberFieldLabel       *string `json:"smsNumberFieldLabel,omitempty"`
    SmsNumberFormat           *string `json:"smsNumberFormat,omitempty"`
    SmsNumberMessage          *string `json:"smsNumberMessage,omitempty"`
    SmsNumberSubmit           *string `json:"smsNumberSubmit,omitempty"`
    SmsNumberTitle            *string `json:"smsNumberTitle,omitempty"`
    SmsUsernameFormat         *string `json:"smsUsernameFormat,omitempty"`
    SponsorBackLink           *string `json:"sponsorBackLink,omitempty"`
    SponsorCancel             *string `json:"sponsorCancel,omitempty"`
    SponsorEmail              *string `json:"sponsorEmail,omitempty"`
    SponsorEmailError         *string `json:"sponsorEmailError,omitempty"`
    SponsorInfoApproved       *string `json:"sponsorInfoApproved,omitempty"`
    SponsorInfoDenied         *string `json:"sponsorInfoDenied,omitempty"`
    SponsorInfoPending        *string `json:"sponsorInfoPending,omitempty"`
    SponsorName               *string `json:"sponsorName,omitempty"`
    SponsorNameError          *string `json:"sponsorNameError,omitempty"`
    SponsorNotePending        *string `json:"sponsorNotePending,omitempty"`
    SponsorRequestAccess      *string `json:"sponsorRequestAccess,omitempty"`
    SponsorStatusApproved     *string `json:"sponsorStatusApproved,omitempty"`
    SponsorStatusDenied       *string `json:"sponsorStatusDenied,omitempty"`
    SponsorStatusPending      *string `json:"sponsorStatusPending,omitempty"`
    SponsorSubmit             *string `json:"sponsorSubmit,omitempty"`
    SponsorsError             *string `json:"sponsorsError,omitempty"`
    SponsorsFieldLabel        *string `json:"sponsorsFieldLabel,omitempty"`
    TosAcceptLabel            *string `json:"tosAcceptLabel,omitempty"`
    TosError                  *string `json:"tosError,omitempty"`
    TosLink                   *string `json:"tosLink,omitempty"`
    TosText                   *string `json:"tosText,omitempty"`
}
