
# Wlan Portal Template Setting

Portal template wlan settings

*This model accepts additional fields of type interface{}.*

## Structure

`WlanPortalTemplateSetting`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AccessCodeAlternateEmail` | `*string` | Optional | **Default**: `"Use alternate email address"` |
| `Alignment` | [`*models.PortalTemplateAlignmentEnum`](../../doc/models/portal-template-alignment-enum.md) | Optional | defines alignment on portal. enum: `center`, `left`, `right`<br><br>**Default**: `"center"` |
| `Ar` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `AuthButtonAmazon` | `*string` | Optional | Label for Amazon auth button<br><br>**Default**: `"Sign in with Amazon"` |
| `AuthButtonAzure` | `*string` | Optional | Label for Azure auth button<br><br>**Default**: `"Sign in with Azure"` |
| `AuthButtonEmail` | `*string` | Optional | Label for Email auth button<br><br>**Default**: `"Sign in with Email"` |
| `AuthButtonFacebook` | `*string` | Optional | Label for Facebook auth button<br><br>**Default**: `"Sign in with Facebook"` |
| `AuthButtonGoogle` | `*string` | Optional | Label for Google auth button<br><br>**Default**: `"Sign in with Google"` |
| `AuthButtonMicrosoft` | `*string` | Optional | Label for Microsoft auth button<br><br>**Default**: `"Sign in with Microsoft"` |
| `AuthButtonPassphrase` | `*string` | Optional | Label for passphrase auth button<br><br>**Default**: `"Sign in with Passphrase"` |
| `AuthButtonSms` | `*string` | Optional | Label for SMS auth button<br><br>**Default**: `"Sign in with Text Message"` |
| `AuthButtonSponsor` | `*string` | Optional | Label for Sponsor auth button<br><br>**Default**: `"Sign in as Guest"` |
| `AuthLabel` | `*string` | Optional | **Default**: `"Connect to Wi-Fi with"` |
| `BackLink` | `*string` | Optional | Label of the link to go back to /logon |
| `CaES` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `Color` | `*string` | Optional | Portal main color<br><br>**Default**: `"#1074bc"` |
| `ColorDark` | `*string` | Optional | **Default**: `"#0b5183"` |
| `ColorLight` | `*string` | Optional | **Default**: `"#3589c6"` |
| `Company` | `*bool` | Optional | Whether company field is required<br><br>**Default**: `false` |
| `CompanyError` | `*string` | Optional | Error message when company not provided<br><br>**Default**: `"Please provide your company name"` |
| `CompanyLabel` | `*string` | Optional | Label of company field<br><br>**Default**: `"Company"` |
| `CsCZ` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `DaDK` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `DeDE` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `ElGR` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `Email` | `*bool` | Optional | Whether email field is required<br><br>**Default**: `false` |
| `EmailAccessDomainError` | `*string` | Optional | Error message when a user has valid social login but doesn't match specified email domains.<br><br>**Default**: `"Email Access Domain Error"` |
| `EmailCancel` | `*string` | Optional | Label for cancel confirmation code submission using email auth<br><br>**Default**: `"Cancel"` |
| `EmailCodeCancel` | `*string` | Optional | **Default**: `"I did not receive the code"` |
| `EmailCodeError` | `*string` | Optional | **Default**: `"Please provide valid alternate email"` |
| `EmailCodeFieldLabel` | `*string` | Optional | **Default**: `"Access Code"` |
| `EmailCodeMessage` | `*string` | Optional | **Default**: `"Enter the access number that was sent to your email address."` |
| `EmailCodeSubmit` | `*string` | Optional | **Default**: `"Sign In"` |
| `EmailCodeTitle` | `*string` | Optional | **Default**: `"Access Code"` |
| `EmailError` | `*string` | Optional | Error message when email not provided<br><br>**Default**: `"Please provide valid email"` |
| `EmailFieldLabel` | `*string` | Optional | **Default**: `"Enter your email address"` |
| `EmailLabel` | `*string` | Optional | Label of email field<br><br>**Default**: `"Email"` |
| `EmailMessage` | `*string` | Optional | **Default**: `"We will email you an authentication code which you can use to connect to the Wi-Fi network."` |
| `EmailSubmit` | `*string` | Optional | Label for confirmation code submit button using email auth<br><br>**Default**: `"Send Access Code"` |
| `EmailTitle` | `*string` | Optional | Title for the Email registration<br><br>**Default**: `"Sign in with Email"` |
| `EnGB` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `EnUS` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `EsES` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `FiFI` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `Field1` | `*bool` | Optional | Whether to ask field1<br><br>**Default**: `false` |
| `Field1Error` | `*string` | Optional | Error message when field1 not provided<br><br>**Default**: `"Please provide Custom Field 1"` |
| `Field1Label` | `*string` | Optional | Label of field1<br><br>**Default**: `"Custom Field 1"` |
| `Field1Required` | `*bool` | Optional | Whether field1 is required field |
| `Field2` | `*bool` | Optional | Whether to ask field2<br><br>**Default**: `false` |
| `Field2Error` | `*string` | Optional | Error message when field2 not provided<br><br>**Default**: `"Please provide Custom Field 2"` |
| `Field2Label` | `*string` | Optional | Label of field2<br><br>**Default**: `"Custom Field 2"` |
| `Field2Required` | `*bool` | Optional | Whether field2 is required field |
| `Field3` | `*bool` | Optional | Whether to ask field3<br><br>**Default**: `false` |
| `Field3Error` | `*string` | Optional | Error message when field3 not provided<br><br>**Default**: `"Please provide Custom Field 3"` |
| `Field3Label` | `*string` | Optional | Label of field3<br><br>**Default**: `"Custom Field 3"` |
| `Field3Required` | `*bool` | Optional | Whether field3 is required field |
| `Field4` | `*bool` | Optional | Whether to ask field4<br><br>**Default**: `false` |
| `Field4Error` | `*string` | Optional | Error message when field4 not provided<br><br>**Default**: `"Please provide Custom Field 4"` |
| `Field4Label` | `*string` | Optional | Label of field4<br><br>**Default**: `"Custom Field 4"` |
| `Field4Required` | `*bool` | Optional | Whether field4 is required field |
| `FrFR` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `HeIL` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `HiIN` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `HrHR` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `HuHU` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `IdID` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `ItIT` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `JaJP` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `KoKR` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `Logo` | `models.Optional[string]` | Optional | Custom logo with `data:image/png;base64,` format, default null, uses Juniper Mist Logo. File size must be less than 100kB and image dimensions must be less than 500px x 200px (width x height). |
| `LogoHeight` | `*int` | Optional | Height of the logo, in px<br><br>**Constraints**: `>= 0`, `<= 200` |
| `LogoWidth` | `*int` | Optional | Width of the logo, in px<br><br>**Constraints**: `>= 0`, `<= 500` |
| `MarketingPolicyLink` | `*string` | Optional | label of the link to go to /marketing_policy<br><br>**Default**: `"Marketing Policy"` |
| `MarketingPolicyOptIn` | `*bool` | Optional | Whether marketing policy optin is enabled<br><br>**Default**: `false` |
| `MarketingPolicyOptInLabel` | `*string` | Optional | label for marketing optin<br><br>**Default**: `"I wish to receive Marketing notifications"` |
| `MarketingPolicyOptInText` | `*string` | Optional | marketing policy text<br><br>**Default**: `"Marketing policy content"` |
| `Message` | `*string` | Optional | **Default**: `"Sign in to get online"` |
| `MsMY` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `MultiAuth` | `*bool` | Optional | **Default**: `false` |
| `Name` | `*bool` | Optional | Whether name field is required<br><br>**Default**: `false` |
| `NameError` | `*string` | Optional | Error message when name not provided<br><br>**Default**: `"Please provide your name"` |
| `NameLabel` | `*string` | Optional | Label of name field<br><br>**Default**: `"Name"` |
| `NbNO` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `NlNL` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `OptOutDefault` | `*bool` | Optional | Default value for the `Do not store` checkbox<br><br>**Default**: `true` |
| `Optout` | `*bool` | Optional | Whether to display Do Not Store My Personal Information<br><br>**Default**: `false` |
| `OptoutLabel` | `*string` | Optional | Label for Do Not Store My Personal Information<br><br>**Default**: `"Do not store"` |
| `PageTitle` | `string` | Required | **Default**: `"Welcome"` |
| `PassphraseCancel` | `*string` | Optional | Label for the Passphrase cancel button<br><br>**Default**: `"Cancel"` |
| `PassphraseError` | `*string` | Optional | Error message when invalid passphrase is provided<br><br>**Default**: `"Invalid Passphrase"` |
| `PassphraseLabel` | `*string` | Optional | Passphrase<br><br>**Default**: `"Passphrase"` |
| `PassphraseMessage` | `*string` | Optional | **Default**: `"Enter the secret passphrase to access the Wi-Fi network."` |
| `PassphraseSubmit` | `*string` | Optional | Label for the Passphrase submit button<br><br>**Default**: `"Sign in"` |
| `PassphraseTitle` | `*string` | Optional | Title for passphrase details page<br><br>**Default**: `"Sign in with Passphrase"` |
| `PlPL` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `PoweredBy` | `*bool` | Optional | Whether to show \"Powered by Mist\"<br><br>**Default**: `true` |
| `Privacy` | `*bool` | Optional | Whether to require the Privacy Term acceptance<br><br>**Default**: `false` |
| `PrivacyPolicyAcceptLabel` | `*string` | Optional | Prefix of the label of the link to go to Privacy Policy<br><br>**Default**: `"I accept the Privacy Terms"` |
| `PrivacyPolicyError` | `*string` | Optional | Error message when Privacy Policy not accepted<br><br>**Default**: `"Please review and accept the Privacy Terms"` |
| `PrivacyPolicyLink` | `*string` | Optional | Label of the link to go to Privacy Policy<br><br>**Default**: `"Privacy Terms"` |
| `PrivacyPolicyText` | `*string` | Optional | Text of the Privacy Policy<br><br>**Default**: `"<< provide your Privacy Terms here >>"` |
| `PtBR` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `PtPT` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `RequiredFieldLabel` | `*string` | Optional | Label to denote required field<br><br>**Default**: `"required"` |
| `ResponsiveLayout` | `*bool` | Optional | **Default**: `true` |
| `RoRO` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `RuRU` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `SignInLabel` | `*string` | Optional | Label of the button to signin<br><br>**Default**: `"Sign In"` |
| `SkSK` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `SmsCarrierDefault` | `*string` | Optional | **Default**: `"Please Select"` |
| `SmsCarrierError` | `*string` | Optional | **Default**: `"Please select a mobile carrier"` |
| `SmsCarrierFieldLabel` | `*string` | Optional | Label for mobile carrier drop-down list<br><br>**Default**: `"Mobile Carrier"` |
| `SmsCodeCancel` | `*string` | Optional | Label for cancel confirmation code submission<br><br>**Default**: `"I did not receive the code"` |
| `SmsCodeError` | `*string` | Optional | Error message when confirmation code is invalid<br><br>**Default**: `"Invalid Access Code"` |
| `SmsCodeFieldLabel` | `*string` | Optional | **Default**: `"Confirmation Code"` |
| `SmsCodeMessage` | `*string` | Optional | **Default**: `"Enter the access number that was sent to your mobile number."` |
| `SmsCodeSubmit` | `*string` | Optional | Label for confirmation code submit button<br><br>**Default**: `"Sign In"` |
| `SmsCodeTitle` | `*string` | Optional | **Default**: `"Access Code"` |
| `SmsCountryFieldLabel` | `*string` | Optional | **Default**: `"Country Code"` |
| `SmsCountryFormat` | `*string` | Optional | **Default**: `"+1"` |
| `SmsHaveAccessCode` | `*string` | Optional | Label for checkbox to specify that the user has access code<br><br>**Default**: `"I have an access code"` |
| `SmsIsTwilio` | `*bool` | Optional | **Default**: `false` |
| `SmsMessageFormat` | `*string` | Optional | Format of access code sms message. {{code}} and {{duration}} are placeholders and should be retained as is. |
| `SmsNumberCancel` | `*string` | Optional | Label for canceling mobile details for SMS auth<br><br>**Default**: `"Cancel"` |
| `SmsNumberError` | `*string` | Optional | **Default**: `"Invalid Mobile Number"` |
| `SmsNumberFieldLabel` | `*string` | Optional | Label for field to provide mobile number<br><br>**Default**: `"Mobile Number"` |
| `SmsNumberFormat` | `*string` | Optional | **Default**: `"2125551212 (digits only)"` |
| `SmsNumberMessage` | `*string` | Optional | **Default**: `"We will send an access code to your mobile number which you can use to connect to the Wi-Fi network. Message and data rates may apply."` |
| `SmsNumberSubmit` | `*string` | Optional | Label for submit button for code generation<br><br>**Default**: `"Send Access Code"` |
| `SmsNumberTitle` | `*string` | Optional | Title for phone number details<br><br>**Default**: `"Sign in with Text Message"` |
| `SmsUsernameFormat` | `*string` | Optional | **Default**: `"username"` |
| `SmsValidityDuration` | `*int` | Optional | How long confirmation code should be considered valid (in minutes)<br><br>**Constraints**: `>= 1`, `<= 30` |
| `SponsorBackLink` | `*string` | Optional | **Default**: `"Go back and edit request form"` |
| `SponsorCancel` | `*string` | Optional | **Default**: `"Cancel"` |
| `SponsorEmail` | `*string` | Optional | Label for Sponsor Email<br><br>**Default**: `"Sponsor Email"` |
| `SponsorEmailError` | `*string` | Optional | **Default**: `"Please provide valid sponsor email"` |
| `SponsorEmailTemplate` | `*string` | Optional | HTML template to replace/override default sponsor email template<br>Sponsor Email Template supports following template variables:<br><br>* `approve_url`: Renders URL to approve the request; optionally &minutes=N query param can be appended to change the Authorization period of the guest, where N is a valid integer denoting number of minutes a guest remains authorized<br>* `deny_url`: Renders URL to reject the request<br>* `guest_email`: Renders Email ID of the guest<br>* `guest_name`: Renders Name of the guest<br>* `field1`: Renders value of the Custom Field 1<br>* `field2`: Renders value of the Custom Field 2<br>* `sponsor_link_validity_duration`: Renders validity time of the request (i.e. Approve/Deny URL)<br>* `auth_expire_minutes`: Renders Wlan-level configured Guest Authorization Expiration time period (in minutes), If not configured then default (1 day in minutes) |
| `SponsorInfoApproved` | `*string` | Optional | **Default**: `"Your request was approved by"` |
| `SponsorInfoDenied` | `*string` | Optional | **Default**: `"Your request was denied by"` |
| `SponsorInfoPending` | `*string` | Optional | **Default**: `"Your notification has been sent to"` |
| `SponsorName` | `*string` | Optional | Label for Sponsor Name<br><br>**Default**: `"Sponsor Name"` |
| `SponsorNameError` | `*string` | Optional | **Default**: `"Please provide sponsor name"` |
| `SponsorNotePending` | `*string` | Optional | **Default**: `"Please wait for them to acknowledge."` |
| `SponsorRequestAccess` | `*string` | Optional | Submit button label request Wifi Access and notify sponsor about guest request<br><br>**Default**: `"Request Wi-Fi Access"` |
| `SponsorStatusApproved` | `*string` | Optional | Text to display if sponsor approves request<br><br>**Default**: `"Your request was approved"` |
| `SponsorStatusDenied` | `*string` | Optional | Text to display when sponsor denies request<br><br>**Default**: `"Your request was denied"` |
| `SponsorStatusPending` | `*string` | Optional | Text to display if request is still pending<br><br>**Default**: `"Notification Sent"` |
| `SponsorSubmit` | `*string` | Optional | Submit button label to notify sponsor about guest request<br><br>**Default**: `"Request Wi-Fi Access"` |
| `SponsorsError` | `*string` | Optional | **Default**: `"Please select a sponsor"` |
| `SponsorsFieldLabel` | `*string` | Optional | **Default**: `"Sponsors"` |
| `SvSE` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `ThTH` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | - |
| `Tos` | `*bool` | Optional | **Default**: `true` |
| `TosAcceptLabel` | `*string` | Optional | Prefix of the label of the link to go to tos<br><br>**Default**: `"I accept the Terms of Service"` |
| `TosError` | `*string` | Optional | Error message when tos not accepted<br><br>**Default**: `"Please review and accept the Terms of Service"` |
| `TosLink` | `*string` | Optional | Label of the link to go to tos<br><br>**Default**: `"Terms of Service"` |
| `TosText` | `*string` | Optional | Text of the Terms of Service<br><br>**Default**: `"<< provide your Terms of Service here >>"` |
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
  "authLabel": "Connect to Wi-Fi with",
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
  "emailMessage": "We will email you an authentication code which you can use to connect to the Wi-Fi network.",
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
  "marketingPolicyLink": "Marketing Policy",
  "marketingPolicyOptIn": false,
  "marketingPolicyOptInLabel": "I wish to receive Marketing notifications",
  "marketingPolicyOptInText": "Marketing policy content",
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
  "passphraseMessage": "Enter the secret passphrase to access the Wi-Fi network.",
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
  "smsNumberMessage": "We will send an access code to your mobile number which you can use to connect to the Wi-Fi network. Message and data rates may apply.",
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
  "sponsorRequestAccess": "Request Wi-Fi Access",
  "sponsorStatusApproved": "Your request was approved",
  "sponsorStatusDenied": "Your request was denied",
  "sponsorStatusPending": "Notification Sent",
  "sponsorSubmit": "Request Wi-Fi Access",
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

