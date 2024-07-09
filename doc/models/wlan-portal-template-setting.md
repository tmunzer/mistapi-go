
# Wlan Portal Template Setting

portal template wlan settings

## Structure

`WlanPortalTemplateSetting`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AccessCodeAlternateEmail` | `*string` | Optional | “Please provide valid alternate email” |
| `Alignment` | `*string` | Optional | defines alignment on portal. “left” is default. |
| `AuthButtonAmazon` | `*string` | Optional | label for Amazon auth button |
| `AuthButtonAzure` | `*string` | Optional | label for Azure auth button |
| `AuthButtonEmail` | `*string` | Optional | label for Email auth button |
| `AuthButtonFacebook` | `*string` | Optional | label for Facebook auth button |
| `AuthButtonGoogle` | `*string` | Optional | label for Google auth button |
| `AuthButtonMicrosoft` | `*string` | Optional | label for Microsoft auth button |
| `AuthButtonPassphrase` | `*string` | Optional | label for passphrase auth button |
| `AuthButtonSms` | `*string` | Optional | label for SMS auth button |
| `AuthButtonSponsor` | `*string` | Optional | label for Sponsor auth button |
| `AuthLabel` | `*string` | Optional | “Connect to WiFi with” |
| `BackLink` | `*string` | Optional | label of the link to go back to /logon |
| `Color` | `*string` | Optional | “#1074bc” |
| `ColorDark` | `*string` | Optional | “#0b5183” |
| `ColorLight` | `*string` | Optional | “#3589c6” |
| `Company` | `*bool` | Optional | whether company field is required<br>**Default**: `false` |
| `CompanyError` | `*string` | Optional | error message when company not provided |
| `CompanyLabel` | `*string` | Optional | label of company field |
| `Email` | `*bool` | Optional | whether email field is required<br>**Default**: `false` |
| `EmailAccessDomainError` | `*string` | Optional | error message when a user has valid social login but doesn’t match specified email domains. |
| `EmailCancel` | `*string` | Optional | Label for cancel confirmation code submission using email auth |
| `EmailCodeCancel` | `*string` | Optional | - |
| `EmailCodeError` | `*string` | Optional | “Please provide valid alternate email” |
| `EmailCodeFieldLabel` | `*string` | Optional | “Confirmation Code” |
| `EmailCodeMessage` | `*string` | Optional | “Enter the access number that was sent to your email address.” |
| `EmailCodeSubmit` | `*string` | Optional | “Sign In |
| `EmailCodeTitle` | `*string` | Optional | “Access Code” |
| `EmailError` | `*string` | Optional | error message when email not provided |
| `EmailFieldLabel` | `*string` | Optional | “Enter your email address” |
| `EmailLabel` | `*string` | Optional | label of email field |
| `EmailMessage` | `*string` | Optional | “We will email you an authentication code which you can use to connect to the WiFi network.” |
| `EmailSubmit` | `*string` | Optional | Label for confirmation code submit button using email auth |
| `EmailTitle` | `*string` | Optional | “Sign in with Email” |
| `Field1` | `*bool` | Optional | whether to ask field1 |
| `Field1Error` | `*string` | Optional | error message when field1 not provided |
| `Field1Label` | `*string` | Optional | label of field1 |
| `Field1Required` | `*bool` | Optional | whether field1 is required field |
| `Field2` | `*bool` | Optional | whether to ask field2 |
| `Field2Error` | `*string` | Optional | error message when field2 not provided |
| `Field2Label` | `*string` | Optional | label of field2 |
| `Field2Required` | `*bool` | Optional | whether field2 is required field |
| `Field3` | `*bool` | Optional | whether to ask field3 |
| `Field3Error` | `*string` | Optional | error message when field3 not provided |
| `Field3Label` | `*string` | Optional | label of field3 |
| `Field3Required` | `*bool` | Optional | whether field3 is required field |
| `Field4` | `*bool` | Optional | whether to ask field4 |
| `Field4Error` | `*string` | Optional | error message when field4 not provided |
| `Field4Label` | `*string` | Optional | label of field4 |
| `Field4Required` | `*bool` | Optional | whether field4 is required field |
| `Message` | `*string` | Optional | “Please enjoy the complimentary Wifi” |
| `Name` | `*bool` | Optional | whether name field is required<br>**Default**: `false` |
| `NameError` | `*string` | Optional | error message when name not provided |
| `NameLabel` | `*string` | Optional | label of name field |
| `Optout` | `*bool` | Optional | whether to display “Do Not Store My Personal Information”<br>**Default**: `false` |
| `OptoutDefault` | `*bool` | Optional | Default value for the \"Do not store\" checkbox"<br>**Default**: `true` |
| `OptoutLabel` | `*string` | Optional | label for “Do Not Store My Personal Information”<br>**Default**: `"Do not store"` |
| `PageTitle` | `string` | Required | “Welcome” |
| `PassphraseCancel` | `*string` | Optional | “Cancel” |
| `PassphraseError` | `*string` | Optional | error message when invalid passphrase is provided |
| `PassphraseLabel` | `*string` | Optional | Passphrase |
| `PassphraseMessage` | `*string` | Optional | “Login using passphrase” |
| `PassphraseSubmit` | `*string` | Optional | “Sign in” |
| `PassphraseTitle` | `*string` | Optional | Title for passphrase details page |
| `PoweredBy` | `*bool` | Optional | whether to show “Powered by Mist”<br>**Default**: `true` |
| `RequiredFieldLabel` | `*string` | Optional | label to denote required field |
| `SignInLabel` | `*string` | Optional | label of the button to /signin |
| `SmsCarrierDefault` | `*string` | Optional | “Please Select” |
| `SmsCarrierError` | `*string` | Optional | “Please select a mobile carrier” |
| `SmsCarrierFieldLabel` | `*string` | Optional | label for mobile carrier drop-down list |
| `SmsCodeCancel` | `*string` | Optional | Label for cancel confirmation code submission |
| `SmsCodeError` | `*string` | Optional | error message when confirmation code is invalid |
| `SmsCodeFieldLabel` | `*string` | Optional | “Confirmation Code” |
| `SmsCodeMessage` | `*string` | Optional | “Enter the confirmation code” |
| `SmsCodeSubmit` | `*string` | Optional | Label for confirmation code submit button |
| `SmsCodeTitle` | `*string` | Optional | “Access Code” |
| `SmsCountryFieldLabel` | `*string` | Optional | “Country Code” |
| `SmsCountryFormat` | `*string` | Optional | “+1” |
| `SmsHaveAccessCode` | `*string` | Optional | Label for checkbox to specify that the user has access code |
| `SmsMessageFormat` | `*string` | Optional | format of access code sms message. {{code}} and {{duration}} are place holders and should be retained as is. |
| `SmsNumberCancel` | `*string` | Optional | label for canceling mobile details for SMS auth |
| `SmsNumberError` | `*string` | Optional | “Invalid Mobile Number” |
| `SmsNumberFieldLabel` | `*string` | Optional | label for field to provide mobile number |
| `SmsNumberFormat` | `*string` | Optional | “2125551212 (digits only)” |
| `SmsNumberMessage` | `*string` | Optional | “We will send an access code to your mobile number which you can use to connect to the WiFi network. Message and data rates may apply.” |
| `SmsNumberSubmit` | `*string` | Optional | label for submit button for code generation |
| `SmsNumberTitle` | `*string` | Optional | Title for phone number details |
| `SmsUsernameFormat` | `*string` | Optional | “username” |
| `SmsValidityDuration` | `*int` | Optional | how long confirmation code should be considered valid (in minutes) |
| `SponsorBackLink` | `*string` | Optional | “Go back and edit request form” |
| `SponsorCancel` | `*string` | Optional | “Cancel” |
| `SponsorEmail` | `*string` | Optional | label for Sponsor Email |
| `SponsorEmailError` | `*string` | Optional | “Please provide valid sponsor email” |
| `SponsorEmailTemplate` | `*string` | Optional | html template to replace/override default sponsor email template<br><br>Sponsor Email Template supports following template variables:<br><br>* `approve_url`: Renders URL to approve the request; optionally &minutes=N query param can be appended to change the Authorization period of the guest, where N is a valid integer denoting number of minutes a guest remains authorized<br>* `deny_url`: Renders URL to reject the request<br>* `guest_email`: Renders Email ID of the guest<br>* `guest_name`: Renders Name of the guest<br>* `field1`: Renders value of the Custom Field 1<br>* `field2`: Renders value of the Custom Field 2<br>* `sponsor_link_validity_duration`: Renders validity time of the request (i.e. Approve/Deny URL)<br>* `auth_expire_minutes`: Renders Wlan-level configured Guest Authorization Expiration time period (in minutes), If not configured then default (1 day in minutes) |
| `SponsorInfoApproved` | `*string` | Optional | “Your request was approved by” |
| `SponsorInfoDenied` | `*string` | Optional | “Your request was denied by” |
| `SponsorInfoPending` | `*string` | Optional | “Your notification has been sent to” |
| `SponsorName` | `*string` | Optional | label for Sponsor Name |
| `SponsorNameError` | `*string` | Optional | “Please provide sponsor’s name” |
| `SponsorNotePending` | `*string` | Optional | “Please wait for them to acknowledge.” |
| `SponsorRequestAccess` | `*string` | Optional | ‘submit button label request Wifi Access and notify sponsor about guest request |
| `SponsorSelectEmail` | `*string` | Optional | “Select Sponsor” |
| `SponsorStatusApproved` | `*string` | Optional | text to display if sponsor approves request |
| `SponsorStatusDenied` | `*string` | Optional | text to display when sponsor denies request |
| `SponsorStatusPending` | `*string` | Optional | text to display if request is still pending |
| `SponsorSubmit` | `*string` | Optional | submit button label to notify sponsor about guest request |
| `SponsorsError` | `*string` | Optional | “Please select a sponsor” |
| `SponsorsInfoApproved` | `*string` | Optional | “Your request was approved” |
| `SponsorsInfoDenied` | `*string` | Optional | “Your request was denied” |
| `SponsorsInfoPending` | `*string` | Optional | “Your notification has been sent to the sponsors” |
| `Tos` | `*bool` | Optional | **Default**: `true` |
| `TosAcceptLabel` | `*string` | Optional | prefix of the label of the link to go to /tos |
| `TosError` | `*string` | Optional | error message when tos not accepted |
| `TosLink` | `*string` | Optional | label of the link to go to /tos |
| `TosText` | `*string` | Optional | text of the Terms of Service |

## Example (as JSON)

```json
{
  "company": false,
  "email": false,
  "name": false,
  "optout": false,
  "optoutDefault": true,
  "optoutLabel": "Do not store",
  "pageTitle": "pageTitle4",
  "poweredBy": true,
  "tos": true,
  "accessCodeAlternateEmail": "accessCodeAlternateEmail8",
  "alignment": "alignment4",
  "authButtonAmazon": "authButtonAmazon0",
  "authButtonAzure": "authButtonAzure8",
  "authButtonEmail": "authButtonEmail6"
}
```

