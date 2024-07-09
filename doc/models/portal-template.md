
# Portal Template

Portal Template

## Structure

`PortalTemplate`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AccessCodeAlternateEmail` | `*string` | Optional | **Default**: `"Please provide valid alternate email"` |
| `Alignment` | `*string` | Optional | defines alignment on portal. “left” is default. |
| `AuthButtonAmazon` | `*string` | Optional | label for Amazon auth button<br>**Default**: `"Sign in with Amazon"` |
| `AuthButtonAzure` | `*string` | Optional | label for Azure auth button<br>**Default**: `"Sign in with Azure"` |
| `AuthButtonEmail` | `*string` | Optional | label for Email auth button<br>**Default**: `"Sign in with Email"` |
| `AuthButtonFacebook` | `*string` | Optional | label for Facebook auth button<br>**Default**: `"Sign in with Facebook"` |
| `AuthButtonGoogle` | `*string` | Optional | label for Google auth button<br>**Default**: `"Sign in with Google"` |
| `AuthButtonMicrosoft` | `*string` | Optional | label for Microsoft auth button<br>**Default**: `"Sign in with Microsoft"` |
| `AuthButtonPassphrase` | `*string` | Optional | label for passphrase auth button |
| `AuthButtonSms` | `*string` | Optional | label for SMS auth button<br>**Default**: `"Sign in with Text Message"` |
| `AuthButtonSponsor` | `*string` | Optional | label for Sponsor auth button<br>**Default**: `"Sign in as Guest"` |
| `AuthLabel` | `*string` | Optional | **Default**: `"“Connect to WiFi with”"` |
| `BackLink` | `*string` | Optional | label of the link to go back to /logon<br>**Default**: `"Go back and edit request form"` |
| `Color` | `*string` | Optional | **Default**: `"“#1074bc”"` |
| `ColorDark` | `*string` | Optional | **Default**: `"“#0b5183”"` |
| `ColorLight` | `*string` | Optional | **Default**: `"“#3589c6”"` |
| `Company` | `*bool` | Optional | whether company field is required<br>**Default**: `false` |
| `CompanyError` | `*string` | Optional | error message when company not provided<br>**Default**: `"Please provide company name"` |
| `CompanyLabel` | `*string` | Optional | label of company field<br>**Default**: `"Company"` |
| `CreatedTime` | `*float64` | Optional | - |
| `Email` | `*bool` | Optional | whether email field is required<br>**Default**: `false` |
| `EmailAccessDomainError` | `*string` | Optional | error message when a user has valid social login but doesn’t match specified email domains.<br>**Default**: `"Access is restricted by email domain"` |
| `EmailCancel` | `*string` | Optional | **Default**: `"Cancel"` |
| `EmailCodeError` | `*string` | Optional | **Default**: `"Please provide valid alternate email"` |
| `EmailError` | `*string` | Optional | error message when email not provided<br>**Default**: `"Please provide valid email"` |
| `EmailFieldLabel` | `*string` | Optional | **Default**: `"Enter your email address"` |
| `EmailLabel` | `*string` | Optional | label of email field<br>**Default**: `"Email"` |
| `EmailMessage` | `*string` | Optional | **Default**: `"We will email you an authentication code which you can use to connect to the WiFi network."` |
| `EmailSubmit` | `*string` | Optional | **Default**: `"“Send Access Code”"` |
| `EmailTitle` | `*string` | Optional | **Default**: `"“Sign in with Email”"` |
| `Field1` | `*bool` | Optional | whether to ask field1<br>**Default**: `false` |
| `Field1Error` | `*string` | Optional | error message when field1 not provided<br>**Default**: `"Please provide field1"` |
| `Field1Label` | `*string` | Optional | label of field1<br>**Default**: `"Custom1"` |
| `Field1Required` | `*bool` | Optional | whether field1 is required field<br>**Default**: `false` |
| `Field2` | `*bool` | Optional | whether to ask field2<br>**Default**: `false` |
| `Field2Error` | `*string` | Optional | error message when field2 not provided<br>**Default**: `"Please provide field2"` |
| `Field2Label` | `*string` | Optional | label of field2<br>**Default**: `"Custom2"` |
| `Field2Required` | `*bool` | Optional | whether field2 is required field<br>**Default**: `false` |
| `Field3` | `*bool` | Optional | whether to ask field3<br>**Default**: `false` |
| `Field3Error` | `*string` | Optional | error message when field3 not provided<br>**Default**: `"Please provide field3"` |
| `Field3Label` | `*string` | Optional | label of field3<br>**Default**: `"Custom3"` |
| `Field3Required` | `*bool` | Optional | whether field3 is required field<br>**Default**: `false` |
| `Field4` | `*bool` | Optional | whether to ask field4<br>**Default**: `false` |
| `Field4Error` | `*string` | Optional | error message when field4 not provided<br>**Default**: `"Please provide field4"` |
| `Field4Label` | `*string` | Optional | label of field4<br>**Default**: `"Custom4"` |
| `Field4Required` | `*bool` | Optional | whether field4 is required field<br>**Default**: `false` |
| `ForSite` | `*bool` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | - |
| `Message` | `*string` | Optional | **Default**: `"Please enjoy the complimentary Wifi"` |
| `ModifiedTime` | `*float64` | Optional | - |
| `Name` | `*bool` | Optional | whether name field is required<br>**Default**: `false` |
| `NameError` | `*string` | Optional | error message when name not provided<br>**Default**: `"Please provide your name"` |
| `NameLabel` | `*string` | Optional | label of name field<br>**Default**: `"Name"` |
| `Optout` | `*bool` | Optional | whether to display “Do Not Store My Personal Information” |
| `OptoutLabel` | `*string` | Optional | **Default**: `"Do Not Store My Personal Information"` |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `PageTitle` | `string` | Required | **Default**: `"Welcome"` |
| `PassphraseCancel` | `*string` | Optional | **Default**: `"Cancel"` |
| `PassphraseError` | `*string` | Optional | error message when invalid passphrase is provided<br>**Default**: `"Invalid Passphrase"` |
| `PassphraseLabel` | `*string` | Optional | **Default**: `"Passphrase"` |
| `PassphraseMessage` | `*string` | Optional | **Default**: `"Login using passphrase"` |
| `PassphraseSubmit` | `*string` | Optional | **Default**: `"Sign in"` |
| `PassphraseTitle` | `*string` | Optional | Title for passphrase details page<br>**Default**: `"Sign in with Passphrase"` |
| `PoweredBy` | `*bool` | Optional | whether to show “Powered by Mist”<br>**Default**: `true` |
| `PrivacyPolicy` | `*bool` | Optional | **Default**: `false` |
| `PrivacyPolicyAcceptLabel` | `*string` | Optional | prefix of the label of the link to go to /privacy_notice |
| `PrivacyPolicyError` | `*string` | Optional | error message when privacy notice is not accepted |
| `PrivacyPolicyLink` | `*string` | Optional | label of the link to go to /privacy_notice, |
| `PrivacyPolicyText` | `*string` | Optional | privacy notice text |
| `RequiredFieldLabel` | `*string` | Optional | label to denote required field<br>**Default**: `"required"` |
| `SignInLabel` | `*string` | Optional | label of the button to /signin<br>**Default**: `"Sign In"` |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `SmsCarrierDefault` | `*string` | Optional | **Default**: `"Please Select"` |
| `SmsCarrierError` | `*string` | Optional | **Default**: `"Please select a mobile carrier"` |
| `SmsCarrierFieldLabel` | `*string` | Optional | label for mobile carrier drop-down list<br>**Default**: `"Mobile Carrier"` |
| `SmsCodeCancel` | `*string` | Optional | Label for cancel confirmation code submission<br>**Default**: `"Cancel"` |
| `SmsCodeError` | `*string` | Optional | error message when confirmation code is invalid<br>**Default**: `"Invalid Access Code"` |
| `SmsCodeFieldLabel` | `*string` | Optional | **Default**: `"Confirmation Code"` |
| `SmsCodeMessage` | `*string` | Optional | **Default**: `"Enter the confirmation code"` |
| `SmsCodeSubmit` | `*string` | Optional | Label for confirmation code submit button<br>**Default**: `"Submit Code"` |
| `SmsCodeTitle` | `*string` | Optional | **Default**: `"Access Code"` |
| `SmsCountryFieldLabel` | `*string` | Optional | “Country Code” |
| `SmsCountryFormat` | `*string` | Optional | **Default**: `"“+1”"` |
| `SmsHaveAccessCode` | `*string` | Optional | Label for checkbox to specify that the user has access code<br>**Default**: `"I have an access code"` |
| `SmsMessageFormat` | `*string` | Optional | format of access code sms message. {{code}} and {{duration}} are place holders and should be retained as is.<br>**Default**: `"Code {{code}} expires in {{duration}} minutes."` |
| `SmsNumberCancel` | `*string` | Optional | label for canceling mobile details for SMS auth<br>**Default**: `"Cancel"` |
| `SmsNumberError` | `*string` | Optional | **Default**: `"Invalid Mobile Number"` |
| `SmsNumberFieldLabel` | `*string` | Optional | label for field to provide mobile number<br>**Default**: `"Mobile Number"` |
| `SmsNumberFormat` | `*string` | Optional | **Default**: `"“2125551212 (digits only)”"` |
| `SmsNumberMessage` | `*string` | Optional | **Default**: `"We will send an access code to your mobile number which you can use to connect to the WiFi network. Message and data rates may apply."` |
| `SmsNumberSubmit` | `*string` | Optional | label for submit button for code generation<br>**Default**: `"Sign In"` |
| `SmsNumberTitle` | `*string` | Optional | Title for phone number details<br>**Default**: `"Text Message Confirmation"` |
| `SmsUsernameFormat` | `*string` | Optional | **Default**: `"username"` |
| `SmsValidityDuration` | `*float64` | Optional | how long confirmation code should be considered valid (in minutes)<br>**Default**: `5` |
| `SponsorAutoApproved` | `*string` | Optional | **Default**: `"Your request was approved."` |
| `SponsorAutoApprovedNote` | `*string` | Optional | **Default**: `"Your notification has been sent to"` |
| `SponsorBackLink` | `*string` | Optional | **Default**: `"Go back and edit request form"` |
| `SponsorCancel` | `*string` | Optional | **Default**: `"Cancel"` |
| `SponsorEmail` | `*string` | Optional | label for Sponsor Email<br>**Default**: `"Please provide valid sponsor email"` |
| `SponsorEmailError` | `*string` | Optional | **Default**: `"Please provide valid sponsor email"` |
| `SponsorEmailTemplate` | `*string` | Optional | “html template to replace/override default sponsor email template” |
| `SponsorInfoApproved` | `*string` | Optional | **Default**: `"Your request was approved by"` |
| `SponsorInfoDenied` | `*string` | Optional | **Default**: `"Your request was denied by"` |
| `SponsorInfoPending` | `*string` | Optional | **Default**: `"Your notification has been sent to"` |
| `SponsorName` | `*string` | Optional | label for Sponsor Name<br>**Default**: `"Sponsor Name"` |
| `SponsorNameError` | `*string` | Optional | **Default**: `"Please provide sponsor’s name"` |
| `SponsorNotePending` | `*string` | Optional | **Default**: `"Please wait for them to acknowledge."` |
| `SponsorStatusApproved` | `*string` | Optional | text to display if sponsor approves request<br>**Default**: `"Your request was approved"` |
| `SponsorStatusDenied` | `*string` | Optional | text to display when sponsor denies request<br>**Default**: `"Your request was denied"` |
| `SponsorStatusPending` | `*string` | Optional | text to display if request is still pending<br>**Default**: `"Notification Sent"` |
| `SponsorSubmit` | `*string` | Optional | submit button label to notify sponsor about guest request<br>**Default**: `"Notify Sponsor"` |
| `SponsorsAutoApprovedNote` | `*string` | Optional | **Default**: `"Your notification has been sent to the sponsors.\""` |
| `SponsorsError` | `*string` | Optional | **Default**: `"Please select a sponsor"` |
| `Tos` | `*bool` | Optional | **Default**: `true` |
| `TosAcceptLabel` | `*string` | Optional | prefix of the label of the link to go to /tos<br>**Default**: `"I accept the Terms of Service"` |
| `TosError` | `*string` | Optional | error message when tos not accepted<br>**Default**: `"Please review and accept terms of service"` |
| `TosLink` | `*string` | Optional | label of the link to go to /tos<br>**Default**: `"Terms of Service"` |
| `TosText` | `*string` | Optional | text of the Terms of Service<br>**Default**: `"terms of service"` |

## Example (as JSON)

```json
{
  "accessCodeAlternateEmail": "Please provide valid alternate email",
  "authButtonAmazon": "Sign in with Amazon",
  "authButtonAzure": "Sign in with Azure",
  "authButtonEmail": "Sign in with Email",
  "authButtonFacebook": "Sign in with Facebook",
  "authButtonGoogle": "Sign in with Google",
  "authButtonMicrosoft": "Sign in with Microsoft",
  "authButtonSms": "Sign in with Text Message",
  "authButtonSponsor": "Sign in as Guest",
  "authLabel": "“Connect to WiFi with”",
  "backLink": "Go back and edit request form",
  "color": "“#1074bc”",
  "colorDark": "“#0b5183”",
  "colorLight": "“#3589c6”",
  "company": false,
  "companyError": "Please provide company name",
  "companyLabel": "Company",
  "email": false,
  "emailAccessDomainError": "Access is restricted by email domain",
  "emailCancel": "Cancel",
  "emailCodeError": "Please provide valid alternate email",
  "emailError": "Please provide valid email",
  "emailFieldLabel": "Enter your email address",
  "emailLabel": "Email",
  "emailMessage": "We will email you an authentication code which you can use to connect to the WiFi network.",
  "emailSubmit": "“Send Access Code”",
  "emailTitle": "“Sign in with Email”",
  "field1": false,
  "field1Error": "Please provide field1",
  "field1Label": "Custom1",
  "field1Required": false,
  "field2": false,
  "field2Error": "Please provide field2",
  "field2Label": "Custom2",
  "field2Required": false,
  "field3": false,
  "field3Error": "Please provide field3",
  "field3Label": "Custom3",
  "field3Required": false,
  "field4": false,
  "field4Error": "Please provide field4",
  "field4Label": "Custom4",
  "field4Required": false,
  "message": "Please enjoy the complimentary Wifi",
  "name": false,
  "nameError": "Please provide your name",
  "nameLabel": "Name",
  "optoutLabel": "Do Not Store My Personal Information",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "pageTitle": "Welcome",
  "passphraseCancel": "Cancel",
  "passphraseError": "Invalid Passphrase",
  "passphraseLabel": "Passphrase",
  "passphraseMessage": "Login using passphrase",
  "passphraseSubmit": "Sign in",
  "passphraseTitle": "Sign in with Passphrase",
  "poweredBy": true,
  "privacyPolicy": false,
  "privacyPolicyAcceptLabel": "I accept the Privacy Notice",
  "privacyPolicyError": "Please review and accept privacy notice",
  "privacyPolicyLink": "Privacy Notice",
  "privacyPolicyText": "privacy notice content",
  "requiredFieldLabel": "required",
  "signInLabel": "Sign In",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "smsCarrierDefault": "Please Select",
  "smsCarrierError": "Please select a mobile carrier",
  "smsCarrierFieldLabel": "Mobile Carrier",
  "smsCodeCancel": "Cancel",
  "smsCodeError": "Invalid Access Code",
  "smsCodeFieldLabel": "Confirmation Code",
  "smsCodeMessage": "Enter the confirmation code",
  "smsCodeSubmit": "Submit Code",
  "smsCodeTitle": "Access Code",
  "smsCountryFormat": "“+1”",
  "smsHaveAccessCode": "I have an access code",
  "smsMessageFormat": "Code {{code}} expires in {{duration}} minutes.",
  "smsNumberCancel": "Cancel",
  "smsNumberError": "Invalid Mobile Number",
  "smsNumberFieldLabel": "Mobile Number",
  "smsNumberFormat": "“2125551212 (digits only)”",
  "smsNumberMessage": "We will send an access code to your mobile number which you can use to connect to the WiFi network. Message and data rates may apply.",
  "smsNumberSubmit": "Sign In",
  "smsNumberTitle": "Text Message Confirmation",
  "smsUsernameFormat": "username",
  "smsValidityDuration": 5,
  "sponsorAutoApproved": "Your request was approved.",
  "sponsorAutoApprovedNote": "Your notification has been sent to",
  "sponsorBackLink": "Go back and edit request form",
  "sponsorCancel": "Cancel",
  "sponsorEmail": "Please provide valid sponsor email",
  "sponsorEmailError": "Please provide valid sponsor email",
  "sponsorInfoApproved": "Your request was approved by",
  "sponsorInfoDenied": "Your request was denied by",
  "sponsorInfoPending": "Your notification has been sent to",
  "sponsorName": "Sponsor Name",
  "sponsorNameError": "Please provide sponsor’s name",
  "sponsorNotePending": "Please wait for them to acknowledge.",
  "sponsorStatusApproved": "Your request was approved",
  "sponsorStatusDenied": "Your request was denied",
  "sponsorStatusPending": "Notification Sent",
  "sponsorSubmit": "Notify Sponsor",
  "sponsorsAutoApprovedNote": "Your notification has been sent to the sponsors.\"",
  "sponsorsError": "Please select a sponsor",
  "tos": true,
  "tosAcceptLabel": "I accept the Terms of Service",
  "tosError": "Please review and accept terms of service",
  "tosLink": "Terms of Service",
  "tosText": "terms of service",
  "alignment": "alignment6"
}
```

