
# Wlan Portal Template Setting Locale

## Structure

`WlanPortalTemplateSettingLocale`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AuthButtonAmazon` | `*string` | Optional | label for Amazon auth button |
| `AuthButtonAzure` | `*string` | Optional | label for Azure auth button |
| `AuthButtonEmail` | `*string` | Optional | label for Email auth button |
| `AuthButtonFacebook` | `*string` | Optional | label for Facebook auth button |
| `AuthButtonGoogle` | `*string` | Optional | label for Google auth button |
| `AuthButtonMicrosoft` | `*string` | Optional | label for Microsoft auth button |
| `AuthButtonPassphrase` | `*string` | Optional | label for passphrase auth button |
| `AuthButtonSms` | `*string` | Optional | label for SMS auth button |
| `AuthButtonSponsor` | `*string` | Optional | label for Sponsor auth button |
| `AuthLabel` | `*string` | Optional | - |
| `BackLink` | `*string` | Optional | label of the link to go back to /logon |
| `CompanyError` | `*string` | Optional | error message when company not provided |
| `CompanyLabel` | `*string` | Optional | label of company field |
| `EmailAccessDomainError` | `*string` | Optional | error message when a user has valid social login but doesn't match specified email domains. |
| `EmailCancel` | `*string` | Optional | Label for cancel confirmation code submission using email auth |
| `EmailCodeCancel` | `*string` | Optional | - |
| `EmailCodeError` | `*string` | Optional | - |
| `EmailCodeFieldLabel` | `*string` | Optional | - |
| `EmailCodeMessage` | `*string` | Optional | - |
| `EmailCodeSubmit` | `*string` | Optional | - |
| `EmailCodeTitle` | `*string` | Optional | - |
| `EmailError` | `*string` | Optional | error message when email not provided |
| `EmailFieldLabel` | `*string` | Optional | - |
| `EmailLabel` | `*string` | Optional | label of email field |
| `EmailMessage` | `*string` | Optional | - |
| `EmailSubmit` | `*string` | Optional | Label for confirmation code submit button using email auth |
| `EmailTitle` | `*string` | Optional | Title for the Email registration |
| `Field1Error` | `*string` | Optional | error message when field1 not provided |
| `Field1Label` | `*string` | Optional | label of field1 |
| `Field2Error` | `*string` | Optional | error message when field2 not provided |
| `Field2Label` | `*string` | Optional | label of field2 |
| `Field3Error` | `*string` | Optional | error message when field3 not provided |
| `Field3Label` | `*string` | Optional | label of field3 |
| `Field4Error` | `*string` | Optional | error message when field4 not provided |
| `Field4Label` | `*string` | Optional | label of field4 |
| `Message` | `*string` | Optional | - |
| `NameError` | `*string` | Optional | error message when name not provided |
| `NameLabel` | `*string` | Optional | label of name field |
| `OptoutLabel` | `*string` | Optional | label for Do Not Store My Personal Information |
| `PageTitle` | `*string` | Optional | - |
| `PassphraseCancel` | `*string` | Optional | Label for the Passphrase cancel button |
| `PassphraseError` | `*string` | Optional | error message when invalid passphrase is provided |
| `PassphraseLabel` | `*string` | Optional | Passphrase |
| `PassphraseMessage` | `*string` | Optional | - |
| `PassphraseSubmit` | `*string` | Optional | Label for the Passphrase submit button |
| `PassphraseTitle` | `*string` | Optional | Title for passphrase details page |
| `PrivacyPolicyAcceptLabel` | `*string` | Optional | prefix of the label of the link to go to Privacy Policy |
| `PrivacyPolicyError` | `*string` | Optional | error message when Privacy Policy not accepted |
| `PrivacyPolicyLink` | `*string` | Optional | label of the link to go to Privacy Policy |
| `PrivacyPolicyText` | `*string` | Optional | text of the Privacy Policy |
| `RequiredFieldLabel` | `*string` | Optional | label to denote required field |
| `SignInLabel` | `*string` | Optional | label of the button to /signin |
| `SmsCarrierDefault` | `*string` | Optional | - |
| `SmsCarrierError` | `*string` | Optional | - |
| `SmsCarrierFieldLabel` | `*string` | Optional | label for mobile carrier drop-down list |
| `SmsCodeCancel` | `*string` | Optional | Label for cancel confirmation code submission |
| `SmsCodeError` | `*string` | Optional | error message when confirmation code is invalid |
| `SmsCodeFieldLabel` | `*string` | Optional | - |
| `SmsCodeMessage` | `*string` | Optional | - |
| `SmsCodeSubmit` | `*string` | Optional | Label for confirmation code submit button |
| `SmsCodeTitle` | `*string` | Optional | - |
| `SmsCountryFieldLabel` | `*string` | Optional | - |
| `SmsCountryFormat` | `*string` | Optional | - |
| `SmsHaveAccessCode` | `*string` | Optional | Label for checkbox to specify that the user has access code |
| `SmsMessageFormat` | `*string` | Optional | format of access code sms message. {{code}} and {{duration}} are place holders and should be retained as is. |
| `SmsNumberCancel` | `*string` | Optional | label for canceling mobile details for SMS auth |
| `SmsNumberError` | `*string` | Optional | - |
| `SmsNumberFieldLabel` | `*string` | Optional | label for field to provide mobile number |
| `SmsNumberFormat` | `*string` | Optional | - |
| `SmsNumberMessage` | `*string` | Optional | - |
| `SmsNumberSubmit` | `*string` | Optional | label for submit button for code generation |
| `SmsNumberTitle` | `*string` | Optional | Title for phone number details |
| `SmsUsernameFormat` | `*string` | Optional | - |
| `SponsorBackLink` | `*string` | Optional | - |
| `SponsorCancel` | `*string` | Optional | - |
| `SponsorEmail` | `*string` | Optional | label for Sponsor Email |
| `SponsorEmailError` | `*string` | Optional | - |
| `SponsorInfoApproved` | `*string` | Optional | - |
| `SponsorInfoDenied` | `*string` | Optional | - |
| `SponsorInfoPending` | `*string` | Optional | - |
| `SponsorName` | `*string` | Optional | label for Sponsor Name |
| `SponsorNameError` | `*string` | Optional | - |
| `SponsorNotePending` | `*string` | Optional | - |
| `SponsorRequestAccess` | `*string` | Optional | submit button label request Wifi Access and notify sponsor about guest request |
| `SponsorStatusApproved` | `*string` | Optional | text to display if sponsor approves request |
| `SponsorStatusDenied` | `*string` | Optional | text to display when sponsor denies request |
| `SponsorStatusPending` | `*string` | Optional | text to display if request is still pending |
| `SponsorSubmit` | `*string` | Optional | submit button label to notify sponsor about guest request |
| `SponsorsError` | `*string` | Optional | - |
| `SponsorsFieldLabel` | `*string` | Optional | - |
| `TosAcceptLabel` | `*string` | Optional | prefix of the label of the link to go to tos |
| `TosError` | `*string` | Optional | error message when tos not accepted |
| `TosLink` | `*string` | Optional | label of the link to go to tos |
| `TosText` | `*string` | Optional | text of the Terms of Service |

## Example (as JSON)

```json
{
  "authButtonAmazon": "authButtonAmazon2",
  "authButtonAzure": "authButtonAzure0",
  "authButtonEmail": "authButtonEmail8",
  "authButtonFacebook": "authButtonFacebook2",
  "authButtonGoogle": "authButtonGoogle8"
}
```
