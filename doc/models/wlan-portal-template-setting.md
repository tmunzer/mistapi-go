
# Wlan Portal Template Setting

portal template wlan settings

*This model accepts additional fields of type interface{}.*

## Structure

`WlanPortalTemplateSetting`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AccessCodeAlternateEmail` | `*string` | Optional | **Default**: `"Use alternate email address"` |
| `Alignment` | [`*models.PortalTemplateAlignmentEnum`](../../doc/models/portal-template-alignment-enum.md) | Optional | defines alignment on portal. enum: `center`, `left`, `right`<br>**Default**: `"center"` |
| `Ar` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `AuthButtonAmazon` | `*string` | Optional | label for Amazon auth button<br>**Default**: `"Sign in with Amazon"` |
| `AuthButtonAzure` | `*string` | Optional | label for Azure auth button<br>**Default**: `"Sign in with Azure"` |
| `AuthButtonEmail` | `*string` | Optional | label for Email auth button<br>**Default**: `"Sign in with Email"` |
| `AuthButtonFacebook` | `*string` | Optional | label for Facebook auth button<br>**Default**: `"Sign in with Facebook"` |
| `AuthButtonGoogle` | `*string` | Optional | label for Google auth button<br>**Default**: `"Sign in with Google"` |
| `AuthButtonMicrosoft` | `*string` | Optional | label for Microsoft auth button<br>**Default**: `"Sign in with Microsoft"` |
| `AuthButtonPassphrase` | `*string` | Optional | label for passphrase auth button<br>**Default**: `"Sign in with Passphrase"` |
| `AuthButtonSms` | `*string` | Optional | label for SMS auth button<br>**Default**: `"Sign in with Text Message"` |
| `AuthButtonSponsor` | `*string` | Optional | label for Sponsor auth button<br>**Default**: `"Sign in as Guest"` |
| `AuthLabel` | `*string` | Optional | **Default**: `"Connect to WiFi with"` |
| `BackLink` | `*string` | Optional | label of the link to go back to /logon |
| `CaES` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `Color` | `*string` | Optional | Portal main color<br>**Default**: `"#1074bc"` |
| `ColorDark` | `*string` | Optional | **Default**: `"#0b5183"` |
| `ColorLight` | `*string` | Optional | **Default**: `"#3589c6"` |
| `Company` | `*bool` | Optional | whether company field is required<br>**Default**: `false` |
| `CompanyError` | `*string` | Optional | error message when company not provided<br>**Default**: `"Please provide your company name"` |
| `CompanyLabel` | `*string` | Optional | label of company field<br>**Default**: `"Company"` |
| `CsCZ` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `DaDK` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `DeDE` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `ElGR` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `Email` | `*bool` | Optional | whether email field is required<br>**Default**: `false` |
| `EmailAccessDomainError` | `*string` | Optional | error message when a user has valid social login but doesn't match specified email domains.<br>**Default**: `"Email Access Domain Error"` |
| `EmailCancel` | `*string` | Optional | Label for cancel confirmation code submission using email auth<br>**Default**: `"Cancel"` |
| `EmailCodeCancel` | `*string` | Optional | **Default**: `"I did not receive the code"` |
| `EmailCodeError` | `*string` | Optional | **Default**: `"Please provide valid alternate email"` |
| `EmailCodeFieldLabel` | `*string` | Optional | **Default**: `"Access Code"` |
| `EmailCodeMessage` | `*string` | Optional | **Default**: `"Enter the access number that was sent to your email address."` |
| `EmailCodeSubmit` | `*string` | Optional | **Default**: `"Sign In"` |
| `EmailCodeTitle` | `*string` | Optional | **Default**: `"Access Code"` |
| `EmailError` | `*string` | Optional | error message when email not provided<br>**Default**: `"Please provide valid email"` |
| `EmailFieldLabel` | `*string` | Optional | **Default**: `"Enter your email address"` |
| `EmailLabel` | `*string` | Optional | label of email field<br>**Default**: `"Email"` |
| `EmailMessage` | `*string` | Optional | **Default**: `"We will email you an authentication code which you can use to connect to the WiFi network."` |
| `EmailSubmit` | `*string` | Optional | Label for confirmation code submit button using email auth<br>**Default**: `"Send Access Code"` |
| `EmailTitle` | `*string` | Optional | Title for the Email registration<br>**Default**: `"Sign in with Email"` |
| `EnGB` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `EnUS` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `EsES` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `FiFI` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `Field1` | `*bool` | Optional | whether to ask field1<br>**Default**: `false` |
| `Field1Error` | `*string` | Optional | error message when field1 not provided<br>**Default**: `"Please provide Custom Field 1"` |
| `Field1Label` | `*string` | Optional | label of field1<br>**Default**: `"Custom Field 1"` |
| `Field1Required` | `*bool` | Optional | whether field1 is required field |
| `Field2` | `*bool` | Optional | whether to ask field2<br>**Default**: `false` |
| `Field2Error` | `*string` | Optional | error message when field2 not provided<br>**Default**: `"Please provide Custom Field 2"` |
| `Field2Label` | `*string` | Optional | label of field2<br>**Default**: `"Custom Field 2"` |
| `Field2Required` | `*bool` | Optional | whether field2 is required field |
| `Field3` | `*bool` | Optional | whether to ask field3<br>**Default**: `false` |
| `Field3Error` | `*string` | Optional | error message when field3 not provided<br>**Default**: `"Please provide Custom Field 3"` |
| `Field3Label` | `*string` | Optional | label of field3<br>**Default**: `"Custom Field 3"` |
| `Field3Required` | `*bool` | Optional | whether field3 is required field |
| `Field4` | `*bool` | Optional | whether to ask field4<br>**Default**: `false` |
| `Field4Error` | `*string` | Optional | error message when field4 not provided<br>**Default**: `"Please provide Custom Field 4"` |
| `Field4Label` | `*string` | Optional | label of field4<br>**Default**: `"Custom Field 4"` |
| `Field4Required` | `*bool` | Optional | whether field4 is required field |
| `FrFR` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `HeIL` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `HiIN` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `HrHR` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `HuHU` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `IdID` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `ItIT` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `JaJP` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `KoKR` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `Logo` | `models.Optional[string]` | Optional | custom logo with `data:image/png;base64,` format, default null, uses Juniper Mist Logo. File size must be less than 100kB and image dimensions must be less than 500px x 200px (width x height). |
| `LogoHeight` | `*int` | Optional | height of the logo, in px<br>**Constraints**: `>= 0`, `<= 200` |
| `LogoWidth` | `*int` | Optional | width of the logo, in px<br>**Constraints**: `>= 0`, `<= 500` |
| `Message` | `*string` | Optional | **Default**: `"Sign in to get online"` |
| `MsMY` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `MultiAuth` | `*bool` | Optional | **Default**: `false` |
| `Name` | `*bool` | Optional | whether name field is required<br>**Default**: `false` |
| `NameError` | `*string` | Optional | error message when name not provided<br>**Default**: `"Please provide your name"` |
| `NameLabel` | `*string` | Optional | label of name field<br>**Default**: `"Name"` |
| `NbNO` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `NlNL` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `OptOutDefault` | `*bool` | Optional | Default value for the `Do not store` checkbox<br>**Default**: `true` |
| `Optout` | `*bool` | Optional | whether to display Do Not Store My Personal Information<br>**Default**: `false` |
| `OptoutLabel` | `*string` | Optional | label for Do Not Store My Personal Information<br>**Default**: `"Do not store"` |
| `PageTitle` | `string` | Required | **Default**: `"Welcome"` |
| `PassphraseCancel` | `*string` | Optional | Label for the Passphrase cancel button<br>**Default**: `"Cancel"` |
| `PassphraseError` | `*string` | Optional | error message when invalid passphrase is provided<br>**Default**: `"Invalid Passphrase"` |
| `PassphraseLabel` | `*string` | Optional | Passphrase<br>**Default**: `"Passphrase"` |
| `PassphraseMessage` | `*string` | Optional | **Default**: `"Enter the secret passphrase to access the WiFi network."` |
| `PassphraseSubmit` | `*string` | Optional | Label for the Passphrase submit button<br>**Default**: `"Sign in"` |
| `PassphraseTitle` | `*string` | Optional | Title for passphrase details page<br>**Default**: `"Sign in with Passphrase"` |
| `PlPL` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `PoweredBy` | `*bool` | Optional | whether to show \"Powered by Mist\"<br>**Default**: `true` |
| `Privacy` | `*bool` | Optional | wheter to require the Privacy Term acceptance<br>**Default**: `false` |
| `PrivacyPolicyAcceptLabel` | `*string` | Optional | prefix of the label of the link to go to Privacy Policy<br>**Default**: `"I accept the Privacy Terms"` |
| `PrivacyPolicyError` | `*string` | Optional | error message when Privacy Policy not accepted<br>**Default**: `"Please review and accept the Privacy Terms"` |
| `PrivacyPolicyLink` | `*string` | Optional | label of the link to go to Privacy Policy<br>**Default**: `"Privacy Terms"` |
| `PrivacyPolicyText` | `*string` | Optional | text of the Privacy Policy<br>**Default**: `"<< provide your Privacy Terms here >>"` |
| `PtBR` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `PtPT` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `RequiredFieldLabel` | `*string` | Optional | label to denote required field<br>**Default**: `"required"` |
| `ResponsiveLayout` | `*bool` | Optional | **Default**: `true` |
| `RoRO` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `RuRU` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `SignInLabel` | `*string` | Optional | label of the button to /signin<br>**Default**: `"Sign In"` |
| `SkSK` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `SmsCarrierDefault` | `*string` | Optional | **Default**: `"Please Select"` |
| `SmsCarrierError` | `*string` | Optional | **Default**: `"Please select a mobile carrier"` |
| `SmsCarrierFieldLabel` | `*string` | Optional | label for mobile carrier drop-down list<br>**Default**: `"Mobile Carrier"` |
| `SmsCodeCancel` | `*string` | Optional | Label for cancel confirmation code submission<br>**Default**: `"I did not receive the code"` |
| `SmsCodeError` | `*string` | Optional | error message when confirmation code is invalid<br>**Default**: `"Invalid Access Code"` |
| `SmsCodeFieldLabel` | `*string` | Optional | **Default**: `"Confirmation Code"` |
| `SmsCodeMessage` | `*string` | Optional | **Default**: `"Enter the access number that was sent to your mobile number."` |
| `SmsCodeSubmit` | `*string` | Optional | Label for confirmation code submit button<br>**Default**: `"Sign In"` |
| `SmsCodeTitle` | `*string` | Optional | **Default**: `"Access Code"` |
| `SmsCountryFieldLabel` | `*string` | Optional | **Default**: `"Country Code"` |
| `SmsCountryFormat` | `*string` | Optional | **Default**: `"+1"` |
| `SmsHaveAccessCode` | `*string` | Optional | Label for checkbox to specify that the user has access code<br>**Default**: `"I have an access code"` |
| `SmsIsTwilio` | `*bool` | Optional | **Default**: `false` |
| `SmsMessageFormat` | `*string` | Optional | format of access code sms message. {{code}} and {{duration}} are place holders and should be retained as is. |
| `SmsNumberCancel` | `*string` | Optional | label for canceling mobile details for SMS auth<br>**Default**: `"Cancel"` |
| `SmsNumberError` | `*string` | Optional | **Default**: `"Invalid Mobile Number"` |
| `SmsNumberFieldLabel` | `*string` | Optional | label for field to provide mobile number<br>**Default**: `"Mobile Number"` |
| `SmsNumberFormat` | `*string` | Optional | **Default**: `"2125551212 (digits only)"` |
| `SmsNumberMessage` | `*string` | Optional | **Default**: `"We will send an access code to your mobile number which you can use to connect to the WiFi network. Message and data rates may apply."` |
| `SmsNumberSubmit` | `*string` | Optional | label for submit button for code generation<br>**Default**: `"Send Access Code"` |
| `SmsNumberTitle` | `*string` | Optional | Title for phone number details<br>**Default**: `"Sign in with Text Message"` |
| `SmsUsernameFormat` | `*string` | Optional | **Default**: `"username"` |
| `SmsValidityDuration` | `*int` | Optional | how long confirmation code should be considered valid (in minutes)<br>**Constraints**: `>= 1`, `<= 30` |
| `SponsorBackLink` | `*string` | Optional | **Default**: `"Go back and edit request form"` |
| `SponsorCancel` | `*string` | Optional | **Default**: `"Cancel"` |
| `SponsorEmail` | `*string` | Optional | label for Sponsor Email<br>**Default**: `"Sponsor Email"` |
| `SponsorEmailError` | `*string` | Optional | **Default**: `"Please provide valid sponsor email"` |
| `SponsorEmailTemplate` | `*string` | Optional | html template to replace/override default sponsor email template<br>Sponsor Email Template supports following template variables:<br><br>* `approve_url`: Renders URL to approve the request; optionally &minutes=N query param can be appended to change the Authorization period of the guest, where N is a valid integer denoting number of minutes a guest remains authorized<br>* `deny_url`: Renders URL to reject the request<br>* `guest_email`: Renders Email ID of the guest<br>* `guest_name`: Renders Name of the guest<br>* `field1`: Renders value of the Custom Field 1<br>* `field2`: Renders value of the Custom Field 2<br>* `sponsor_link_validity_duration`: Renders validity time of the request (i.e. Approve/Deny URL)<br>* `auth_expire_minutes`: Renders Wlan-level configured Guest Authorization Expiration time period (in minutes), If not configured then default (1 day in minutes) |
| `SponsorInfoApproved` | `*string` | Optional | **Default**: `"Your request was approved by"` |
| `SponsorInfoDenied` | `*string` | Optional | **Default**: `"Your request was denied by"` |
| `SponsorInfoPending` | `*string` | Optional | **Default**: `"Your notification has been sent to"` |
| `SponsorName` | `*string` | Optional | label for Sponsor Name<br>**Default**: `"Sponsor Name"` |
| `SponsorNameError` | `*string` | Optional | **Default**: `"Please provide sponsor name"` |
| `SponsorNotePending` | `*string` | Optional | **Default**: `"Please wait for them to acknowledge."` |
| `SponsorRequestAccess` | `*string` | Optional | submit button label request Wifi Access and notify sponsor about guest request<br>**Default**: `"Request WiFi Access"` |
| `SponsorStatusApproved` | `*string` | Optional | text to display if sponsor approves request<br>**Default**: `"Your request was approved"` |
| `SponsorStatusDenied` | `*string` | Optional | text to display when sponsor denies request<br>**Default**: `"Your request was denied"` |
| `SponsorStatusPending` | `*string` | Optional | text to display if request is still pending<br>**Default**: `"Notification Sent"` |
| `SponsorSubmit` | `*string` | Optional | submit button label to notify sponsor about guest request<br>**Default**: `"Request WiFi Access"` |
| `SponsorsError` | `*string` | Optional | **Default**: `"Please select a sponsor"` |
| `SponsorsFieldLabel` | `*string` | Optional | **Default**: `"Sponsors"` |
| `SvSE` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `ThTH` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `Tos` | `*bool` | Optional | **Default**: `true` |
| `TosAcceptLabel` | `*string` | Optional | prefix of the label of the link to go to tos<br>**Default**: `"I accept the Terms of Service"` |
| `TosError` | `*string` | Optional | error message when tos not accepted<br>**Default**: `"Please review and accept the Terms of Service"` |
| `TosLink` | `*string` | Optional | label of the link to go to tos<br>**Default**: `"Terms of Service"` |
| `TosText` | `*string` | Optional | text of the Terms of Service<br>**Default**: `"<< provide your Terms of Service here >>"` |
| `TrTR` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `UkUA` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `ViVN` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `ZhHans` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `ZhHant` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "accessCodeAlternateEmail": "Use alternate email address",
  "alignment": "center",
  "authButtonAmazon": "Sign in with Amazon",
  "authButtonAzure": "Sign in with Azure",
  "authButtonEmail": "Sign in with Email",
  "authButtonFacebook": "Sign in with Facebook",
  "authButtonGoogle": "Sign in with Google",
  "authButtonMicrosoft": "Sign in with Microsoft",
  "authButtonPassphrase": "Sign in with Passphrase",
  "authButtonSms": "Sign in with Text Message",
  "authButtonSponsor": "Sign in as Guest",
  "authLabel": "Connect to WiFi with",
  "color": "#1074bc",
  "colorDark": "#0b5183",
  "colorLight": "#3589c6",
  "company": false,
  "companyError": "Please provide your company name",
  "companyLabel": "Company",
  "email": false,
  "emailAccessDomainError": "Email Access Domain Error",
  "emailCancel": "Cancel",
  "emailCodeCancel": "I did not receive the code",
  "emailCodeError": "Please provide valid alternate email",
  "emailCodeFieldLabel": "Access Code",
  "emailCodeMessage": "Enter the access number that was sent to your email address.",
  "emailCodeSubmit": "Sign In",
  "emailCodeTitle": "Access Code",
  "emailError": "Please provide valid email",
  "emailFieldLabel": "Enter your email address",
  "emailLabel": "Email",
  "emailMessage": "We will email you an authentication code which you can use to connect to the WiFi network.",
  "emailSubmit": "Send Access Code",
  "emailTitle": "Sign in with Email",
  "field1": false,
  "field1Error": "Please provide Custom Field 1",
  "field1Label": "Custom Field 1",
  "field2": false,
  "field2Error": "Please provide Custom Field 2",
  "field2Label": "Custom Field 2",
  "field3": false,
  "field3Error": "Please provide Custom Field 3",
  "field3Label": "Custom Field 3",
  "field4": false,
  "field4Error": "Please provide Custom Field 4",
  "field4Label": "Custom Field 4",
  "logo": "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAZg…",
  "logoHeight": 123,
  "logoWidth": 408,
  "message": "Sign in to get online",
  "multiAuth": false,
  "name": false,
  "nameError": "Please provide your name",
  "nameLabel": "Name",
  "optOutDefault": true,
  "optout": false,
  "optoutLabel": "Do not store",
  "pageTitle": "Welcome",
  "passphraseCancel": "Cancel",
  "passphraseError": "Invalid Passphrase",
  "passphraseLabel": "Passphrase",
  "passphraseMessage": "Enter the secret passphrase to access the WiFi network.",
  "passphraseSubmit": "Sign in",
  "passphraseTitle": "Sign in with Passphrase",
  "poweredBy": true,
  "privacy": false,
  "privacyPolicyAcceptLabel": "I accept the Privacy Terms",
  "privacyPolicyError": "Please review and accept the Privacy Terms",
  "privacyPolicyLink": "Privacy Terms",
  "privacyPolicyText": "<< provide your Privacy Terms here >>",
  "requiredFieldLabel": "required",
  "responsiveLayout": true,
  "signInLabel": "Sign In",
  "smsCarrierDefault": "Please Select",
  "smsCarrierError": "Please select a mobile carrier",
  "smsCarrierFieldLabel": "Mobile Carrier",
  "smsCodeCancel": "I did not receive the code",
  "smsCodeError": "Invalid Access Code",
  "smsCodeFieldLabel": "Confirmation Code",
  "smsCodeMessage": "Enter the access number that was sent to your mobile number.",
  "smsCodeSubmit": "Sign In",
  "smsCodeTitle": "Access Code",
  "smsCountryFieldLabel": "Country Code",
  "smsCountryFormat": "+1",
  "smsHaveAccessCode": "I have an access code",
  "smsIsTwilio": false,
  "smsNumberCancel": "Cancel",
  "smsNumberError": "Invalid Mobile Number",
  "smsNumberFieldLabel": "Mobile Number",
  "smsNumberFormat": "2125551212 (digits only)",
  "smsNumberMessage": "We will send an access code to your mobile number which you can use to connect to the WiFi network. Message and data rates may apply.",
  "smsNumberSubmit": "Send Access Code",
  "smsNumberTitle": "Sign in with Text Message",
  "smsUsernameFormat": "username",
  "sponsorBackLink": "Go back and edit request form",
  "sponsorCancel": "Cancel",
  "sponsorEmail": "Sponsor Email",
  "sponsorEmailError": "Please provide valid sponsor email",
  "sponsorInfoApproved": "Your request was approved by",
  "sponsorInfoDenied": "Your request was denied by",
  "sponsorInfoPending": "Your notification has been sent to",
  "sponsorName": "Sponsor Name",
  "sponsorNameError": "Please provide sponsor name",
  "sponsorNotePending": "Please wait for them to acknowledge.",
  "sponsorRequestAccess": "Request WiFi Access",
  "sponsorStatusApproved": "Your request was approved",
  "sponsorStatusDenied": "Your request was denied",
  "sponsorStatusPending": "Notification Sent",
  "sponsorSubmit": "Request WiFi Access",
  "sponsorsError": "Please select a sponsor",
  "sponsorsFieldLabel": "Sponsors",
  "tos": true,
  "tosAcceptLabel": "I accept the Terms of Service",
  "tosError": "Please review and accept the Terms of Service",
  "tosLink": "Terms of Service",
  "tosText": "<< provide your Terms of Service here >>",
  "ar": {
    "authButtonAmazon": "authButtonAmazon6",
    "authButtonAzure": "authButtonAzure4",
    "authButtonEmail": "authButtonEmail2",
    "authButtonFacebook": "authButtonFacebook6",
    "authButtonGoogle": "authButtonGoogle2",
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

