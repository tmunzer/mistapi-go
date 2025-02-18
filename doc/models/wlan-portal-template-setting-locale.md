
# Wlan Portal Template Setting Locale

*This model accepts additional fields of type interface{}.*

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
| `AuthLabel` | `*string` | Optional | - |
| `BackLink` | `*string` | Optional | Label of the link to go back to /logon |
| `CompanyError` | `*string` | Optional | Error message when company not provided |
| `CompanyLabel` | `*string` | Optional | Label of company field |
| `EmailAccessDomainError` | `*string` | Optional | Error message when a user has valid social login but doesn't match specified email domains. |
| `EmailCancel` | `*string` | Optional | Label for cancel confirmation code submission using email auth |
| `EmailCodeCancel` | `*string` | Optional | - |
| `EmailCodeError` | `*string` | Optional | - |
| `EmailCodeFieldLabel` | `*string` | Optional | - |
| `EmailCodeMessage` | `*string` | Optional | - |
| `EmailCodeSubmit` | `*string` | Optional | - |
| `EmailCodeTitle` | `*string` | Optional | - |
| `EmailError` | `*string` | Optional | Error message when email not provided |
| `EmailFieldLabel` | `*string` | Optional | - |
| `EmailLabel` | `*string` | Optional | Label of email field |
| `EmailMessage` | `*string` | Optional | - |
| `EmailSubmit` | `*string` | Optional | Label for confirmation code submit button using email auth |
| `EmailTitle` | `*string` | Optional | Title for the Email registration |
| `Field1Error` | `*string` | Optional | Error message when field1 not provided |
| `Field1Label` | `*string` | Optional | Label of field1 |
| `Field2Error` | `*string` | Optional | Error message when field2 not provided |
| `Field2Label` | `*string` | Optional | Label of field2 |
| `Field3Error` | `*string` | Optional | Error message when field3 not provided |
| `Field3Label` | `*string` | Optional | Label of field3 |
| `Field4Error` | `*string` | Optional | Error message when field4 not provided |
| `Field4Label` | `*string` | Optional | Label of field4 |
| `Message` | `*string` | Optional | - |
| `NameError` | `*string` | Optional | Error message when name not provided |
| `NameLabel` | `*string` | Optional | Label of name field |
| `OptoutLabel` | `*string` | Optional | Label for Do Not Store My Personal Information |
| `PageTitle` | `*string` | Optional | - |
| `PassphraseCancel` | `*string` | Optional | Label for the Passphrase cancel button |
| `PassphraseError` | `*string` | Optional | Error message when invalid passphrase is provided |
| `PassphraseLabel` | `*string` | Optional | Passphrase |
| `PassphraseMessage` | `*string` | Optional | - |
| `PassphraseSubmit` | `*string` | Optional | Label for the Passphrase submit button |
| `PassphraseTitle` | `*string` | Optional | Title for passphrase details page |
| `PrivacyPolicyAcceptLabel` | `*string` | Optional | Prefix of the label of the link to go to Privacy Policy |
| `PrivacyPolicyError` | `*string` | Optional | Error message when Privacy Policy not accepted |
| `PrivacyPolicyLink` | `*string` | Optional | Label of the link to go to Privacy Policy |
| `PrivacyPolicyText` | `*string` | Optional | Text of the Privacy Policy |
| `RequiredFieldLabel` | `*string` | Optional | Label to denote required field |
| `SignInLabel` | `*string` | Optional | Label of the button to /signin |
| `SmsCarrierDefault` | `*string` | Optional | - |
| `SmsCarrierError` | `*string` | Optional | - |
| `SmsCarrierFieldLabel` | `*string` | Optional | Label for mobile carrier drop-down list |
| `SmsCodeCancel` | `*string` | Optional | Label for cancel confirmation code submission |
| `SmsCodeError` | `*string` | Optional | Error message when confirmation code is invalid |
| `SmsCodeFieldLabel` | `*string` | Optional | - |
| `SmsCodeMessage` | `*string` | Optional | - |
| `SmsCodeSubmit` | `*string` | Optional | Label for confirmation code submit button |
| `SmsCodeTitle` | `*string` | Optional | - |
| `SmsCountryFieldLabel` | `*string` | Optional | - |
| `SmsCountryFormat` | `*string` | Optional | - |
| `SmsHaveAccessCode` | `*string` | Optional | Label for checkbox to specify that the user has access code |
| `SmsMessageFormat` | `*string` | Optional | Format of access code sms message. {{code}} and {{duration}} are place holders and should be retained as is. |
| `SmsNumberCancel` | `*string` | Optional | Label for canceling mobile details for SMS auth |
| `SmsNumberError` | `*string` | Optional | - |
| `SmsNumberFieldLabel` | `*string` | Optional | Label for field to provide mobile number |
| `SmsNumberFormat` | `*string` | Optional | - |
| `SmsNumberMessage` | `*string` | Optional | - |
| `SmsNumberSubmit` | `*string` | Optional | Label for submit button for code generation |
| `SmsNumberTitle` | `*string` | Optional | Title for phone number details |
| `SmsUsernameFormat` | `*string` | Optional | - |
| `SponsorBackLink` | `*string` | Optional | - |
| `SponsorCancel` | `*string` | Optional | - |
| `SponsorEmail` | `*string` | Optional | Label for Sponsor Email |
| `SponsorEmailError` | `*string` | Optional | - |
| `SponsorInfoApproved` | `*string` | Optional | - |
| `SponsorInfoDenied` | `*string` | Optional | - |
| `SponsorInfoPending` | `*string` | Optional | - |
| `SponsorName` | `*string` | Optional | Label for Sponsor Name |
| `SponsorNameError` | `*string` | Optional | - |
| `SponsorNotePending` | `*string` | Optional | - |
| `SponsorRequestAccess` | `*string` | Optional | Submit button label request Wifi Access and notify sponsor about guest request |
| `SponsorStatusApproved` | `*string` | Optional | Text to display if sponsor approves request |
| `SponsorStatusDenied` | `*string` | Optional | Text to display when sponsor denies request |
| `SponsorStatusPending` | `*string` | Optional | Text to display if request is still pending |
| `SponsorSubmit` | `*string` | Optional | Submit button label to notify sponsor about guest request |
| `SponsorsError` | `*string` | Optional | - |
| `SponsorsFieldLabel` | `*string` | Optional | - |
| `TosAcceptLabel` | `*string` | Optional | Prefix of the label of the link to go to tos |
| `TosError` | `*string` | Optional | Error message when tos not accepted |
| `TosLink` | `*string` | Optional | Label of the link to go to tos |
| `TosText` | `*string` | Optional | Text of the Terms of Service |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "authButtonAmazon": "authButtonAmazon2",
  "authButtonAzure": "authButtonAzure0",
  "authButtonEmail": "authButtonEmail8",
  "authButtonFacebook": "authButtonFacebook2",
  "authButtonGoogle": "authButtonGoogle8",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

