
# Wlan Portal Template Setting Locale

Localized portal template strings for a specific language

## Structure

`WlanPortalTemplateSettingLocale`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AuthButtonAmazon` | `*string` | Optional | Label for Amazon auth button |
| `AuthButtonAzure` | `*string` | Optional | Label for Azure auth button |
| `AuthButtonEmail` | `*string` | Optional | Label for Email auth button |
| `AuthButtonFacebook` | `*string` | Optional | Label for Facebook auth button |
| `AuthButtonGoogle` | `*string` | Optional | Label for Google auth button |
| `AuthButtonMicrosoft` | `*string` | Optional | Label for Microsoft auth button |
| `AuthButtonPassphrase` | `*string` | Optional | Label for passphrase auth button |
| `AuthButtonSms` | `*string` | Optional | Label for SMS auth button |
| `AuthButtonSponsor` | `*string` | Optional | Label for Sponsor auth button |
| `AuthLabel` | `*string` | Optional | Localized heading text displayed above portal authentication options |
| `BackLink` | `*string` | Optional | Label of the link to go back to /logon |
| `CompanyError` | `*string` | Optional | Error message when company not provided |
| `CompanyLabel` | `*string` | Optional | Localized label displayed for the company input field |
| `EmailAccessDomainError` | `*string` | Optional | Error message when a user has valid social login but doesn't match specified email domains. |
| `EmailCancel` | `*string` | Optional | Label for cancel confirmation code submission using email auth |
| `EmailCodeCancel` | `*string` | Optional | Localized link text for requesting help when the email access code was not received |
| `EmailCodeError` | `*string` | Optional | Localized error message shown when the alternate email address for access-code delivery is invalid |
| `EmailCodeFieldLabel` | `*string` | Optional | Localized label for the email access-code input field |
| `EmailCodeMessage` | `*string` | Optional | Localized instructional text shown before entering the email access code |
| `EmailCodeSubmit` | `*string` | Optional | Localized button label for submitting the email access code |
| `EmailCodeTitle` | `*string` | Optional | Localized title shown on the email access-code entry page |
| `EmailError` | `*string` | Optional | Error message when email not provided |
| `EmailFieldLabel` | `*string` | Optional | Localized label for the email address input field |
| `EmailLabel` | `*string` | Optional | Localized label displayed for the email input field |
| `EmailMessage` | `*string` | Optional | Localized instructional text explaining email access-code delivery |
| `EmailSubmit` | `*string` | Optional | Label for confirmation code submit button using email auth |
| `EmailTitle` | `*string` | Optional | Title for the Email registration |
| `Field1Error` | `*string` | Optional | Error message when field1 not provided |
| `Field1Label` | `*string` | Optional | Localized label for custom field 1 input |
| `Field2Error` | `*string` | Optional | Error message when field2 not provided |
| `Field2Label` | `*string` | Optional | Localized label for custom field 2 input |
| `Field3Error` | `*string` | Optional | Error message when field3 not provided |
| `Field3Label` | `*string` | Optional | Localized label for custom field 3 input |
| `Field4Error` | `*string` | Optional | Error message when field4 not provided |
| `Field4Label` | `*string` | Optional | Localized label for custom field 4 input |
| `MarketingPolicyLink` | `*string` | Optional | label of the link to go to /marketing_policy |
| `MarketingPolicyOptIn` | `*bool` | Optional | Whether marketing policy optin is enabled |
| `MarketingPolicyOptInLabel` | `*string` | Optional | label for marketing optin |
| `MarketingPolicyOptInText` | `*string` | Optional | Localized text of the marketing policy opt-in content |
| `Message` | `*string` | Optional | Localized main message displayed on the guest portal sign-in page |
| `NameError` | `*string` | Optional | Error message when name not provided |
| `NameLabel` | `*string` | Optional | Localized label displayed for the name input field |
| `OptoutLabel` | `*string` | Optional | Label for Do Not Store My Personal Information |
| `PageTitle` | `*string` | Optional | Localized browser or page title shown for the guest portal |
| `PassphraseCancel` | `*string` | Optional | Label for the Passphrase cancel button |
| `PassphraseError` | `*string` | Optional | Error message when invalid passphrase is provided |
| `PassphraseLabel` | `*string` | Optional | Localized label for the passphrase input field |
| `PassphraseMessage` | `*string` | Optional | Localized instructional text shown on the passphrase sign-in page |
| `PassphraseSubmit` | `*string` | Optional | Label for the Passphrase submit button |
| `PassphraseTitle` | `*string` | Optional | Title for passphrase details page |
| `PrivacyPolicyAcceptLabel` | `*string` | Optional | Prefix of the label of the link to go to Privacy Policy |
| `PrivacyPolicyError` | `*string` | Optional | Error message when Privacy Policy not accepted |
| `PrivacyPolicyLink` | `*string` | Optional | Label of the link to go to Privacy Policy |
| `PrivacyPolicyText` | `*string` | Optional | Text of the Privacy Policy |
| `RequiredFieldLabel` | `*string` | Optional | Localized text used to mark a form field as required |
| `SignInLabel` | `*string` | Optional | Label of the button to signin |
| `SmsCarrierDefault` | `*string` | Optional | Localized default option text shown in the SMS carrier selector |
| `SmsCarrierError` | `*string` | Optional | Localized error message shown when no mobile carrier is selected |
| `SmsCarrierFieldLabel` | `*string` | Optional | Label for mobile carrier drop-down list |
| `SmsCodeCancel` | `*string` | Optional | Label for cancel confirmation code submission |
| `SmsCodeError` | `*string` | Optional | Error message when confirmation code is invalid |
| `SmsCodeFieldLabel` | `*string` | Optional | Localized label for the SMS confirmation-code input field |
| `SmsCodeMessage` | `*string` | Optional | Localized instructional text shown before entering the SMS access code |
| `SmsCodeSubmit` | `*string` | Optional | Label for confirmation code submit button |
| `SmsCodeTitle` | `*string` | Optional | Localized title shown on the SMS access-code entry page |
| `SmsCountryFieldLabel` | `*string` | Optional | Localized label for the SMS country-code input field |
| `SmsCountryFormat` | `*string` | Optional | Localized example country code format shown for SMS authentication |
| `SmsHaveAccessCode` | `*string` | Optional | Label for checkbox to specify that the user has access code |
| `SmsMessageFormat` | `*string` | Optional | Format of access code sms message. {{code}} and {{duration}} are placeholders and should be retained as is. |
| `SmsNumberCancel` | `*string` | Optional | Label for canceling mobile details for SMS auth |
| `SmsNumberError` | `*string` | Optional | Localized error message shown when the mobile number is invalid |
| `SmsNumberFieldLabel` | `*string` | Optional | Label for field to provide mobile number |
| `SmsNumberFormat` | `*string` | Optional | Localized example mobile number format shown for SMS authentication |
| `SmsNumberMessage` | `*string` | Optional | Localized instructional text explaining SMS access-code delivery |
| `SmsNumberSubmit` | `*string` | Optional | Label for submit button for code generation |
| `SmsNumberTitle` | `*string` | Optional | Title for phone number details |
| `SmsUsernameFormat` | `*string` | Optional | Localized example username format shown for SMS authentication |
| `SponsorBackLink` | `*string` | Optional | Localized link text for returning to edit the sponsor request form |
| `SponsorCancel` | `*string` | Optional | Localized button label for canceling sponsor authentication |
| `SponsorEmail` | `*string` | Optional | Label for Sponsor Email |
| `SponsorEmailError` | `*string` | Optional | Localized error message shown when the sponsor email address is invalid |
| `SponsorInfoApproved` | `*string` | Optional | Localized status message prefix shown when a sponsor approves the request |
| `SponsorInfoDenied` | `*string` | Optional | Localized status message prefix shown when a sponsor denies the request |
| `SponsorInfoPending` | `*string` | Optional | Localized status message prefix shown after a sponsor notification is sent |
| `SponsorName` | `*string` | Optional | Label for Sponsor Name |
| `SponsorNameError` | `*string` | Optional | Localized error message shown when the sponsor name is missing |
| `SponsorNotePending` | `*string` | Optional | Localized additional status text shown while sponsor approval is pending |
| `SponsorRequestAccess` | `*string` | Optional | Submit button label request Wifi Access and notify sponsor about guest request |
| `SponsorStatusApproved` | `*string` | Optional | Text to display if sponsor approves request |
| `SponsorStatusDenied` | `*string` | Optional | Text to display when sponsor denies request |
| `SponsorStatusPending` | `*string` | Optional | Text to display if request is still pending |
| `SponsorSubmit` | `*string` | Optional | Submit button label to notify sponsor about guest request |
| `SponsorsError` | `*string` | Optional | Localized error message shown when no sponsor is selected |
| `SponsorsFieldLabel` | `*string` | Optional | Localized label for the sponsor selection field |
| `TosAcceptLabel` | `*string` | Optional | Prefix of the label of the link to go to tos |
| `TosError` | `*string` | Optional | Error message when tos not accepted |
| `TosLink` | `*string` | Optional | Label of the link to go to tos |
| `TosText` | `*string` | Optional | Text of the Terms of Service |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    wlanPortalTemplateSettingLocale := models.WlanPortalTemplateSettingLocale{
        AuthButtonAmazon:          models.ToPointer("authButtonAmazon0"),
        AuthButtonAzure:           models.ToPointer("authButtonAzure8"),
        AuthButtonEmail:           models.ToPointer("authButtonEmail6"),
        AuthButtonFacebook:        models.ToPointer("authButtonFacebook0"),
        AuthButtonGoogle:          models.ToPointer("authButtonGoogle6"),
    }

}
```

