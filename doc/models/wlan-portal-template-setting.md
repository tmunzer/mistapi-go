
# Wlan Portal Template Setting

Portal template settings for the WLAN guest portal

## Structure

`WlanPortalTemplateSetting`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AccessCodeAlternateEmail` | `*string` | Optional | Link text for using an alternate email address during access-code login<br><br>**Default**: `"Use alternate email address"` |
| `Alignment` | [`*models.PortalTemplateAlignmentEnum`](../../doc/models/portal-template-alignment-enum.md) | Optional | defines alignment on portal. enum: `center`, `left`, `right`<br><br>**Default**: `"center"` |
| `Ar` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | Localized portal template strings for a specific language |
| `AuthButtonAmazon` | `*string` | Optional | Label for Amazon auth button<br><br>**Default**: `"Sign in with Amazon"` |
| `AuthButtonAzure` | `*string` | Optional | Label for Azure auth button<br><br>**Default**: `"Sign in with Azure"` |
| `AuthButtonEmail` | `*string` | Optional | Label for Email auth button<br><br>**Default**: `"Sign in with Email"` |
| `AuthButtonFacebook` | `*string` | Optional | Label for Facebook auth button<br><br>**Default**: `"Sign in with Facebook"` |
| `AuthButtonGoogle` | `*string` | Optional | Label for Google auth button<br><br>**Default**: `"Sign in with Google"` |
| `AuthButtonMicrosoft` | `*string` | Optional | Label for Microsoft auth button<br><br>**Default**: `"Sign in with Microsoft"` |
| `AuthButtonPassphrase` | `*string` | Optional | Label for passphrase auth button<br><br>**Default**: `"Sign in with Passphrase"` |
| `AuthButtonSms` | `*string` | Optional | Label for SMS auth button<br><br>**Default**: `"Sign in with Text Message"` |
| `AuthButtonSponsor` | `*string` | Optional | Label for Sponsor auth button<br><br>**Default**: `"Sign in as Guest"` |
| `AuthLabel` | `*string` | Optional | Heading text displayed above portal authentication options<br><br>**Default**: `"Connect to Wi-Fi with"` |
| `BackLink` | `*string` | Optional | Label of the link to go back to /logon |
| `CaES` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | Localized portal template strings for a specific language |
| `Color` | `*string` | Optional | Primary color used by the portal template<br><br>**Default**: `"#1074bc"` |
| `ColorDark` | `*string` | Optional | Darker accent color used by the portal template<br><br>**Default**: `"#0b5183"` |
| `ColorLight` | `*string` | Optional | Lighter accent color used by the portal template<br><br>**Default**: `"#3589c6"` |
| `Company` | `*bool` | Optional | Whether company field is required<br><br>**Default**: `false` |
| `CompanyError` | `*string` | Optional | Error message when company not provided<br><br>**Default**: `"Please provide your company name"` |
| `CompanyLabel` | `*string` | Optional | Label displayed for the company input field<br><br>**Default**: `"Company"` |
| `CsCZ` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | Localized portal template strings for a specific language |
| `DaDK` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | Localized portal template strings for a specific language |
| `DeDE` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | Localized portal template strings for a specific language |
| `ElGR` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | Localized portal template strings for a specific language |
| `Email` | `*bool` | Optional | Whether email field is required<br><br>**Default**: `false` |
| `EmailAccessDomainError` | `*string` | Optional | Error message when a user has valid social login but doesn't match specified email domains.<br><br>**Default**: `"Email Access Domain Error"` |
| `EmailCancel` | `*string` | Optional | Label for cancel confirmation code submission using email auth<br><br>**Default**: `"Cancel"` |
| `EmailCodeCancel` | `*string` | Optional | Link text for requesting help when the email access code was not received<br><br>**Default**: `"I did not receive the code"` |
| `EmailCodeError` | `*string` | Optional | Error message shown when the alternate email address for access-code delivery is invalid<br><br>**Default**: `"Please provide valid alternate email"` |
| `EmailCodeFieldLabel` | `*string` | Optional | Label for the email access-code input field<br><br>**Default**: `"Access Code"` |
| `EmailCodeMessage` | `*string` | Optional | Instructional text shown before entering the email access code<br><br>**Default**: `"Enter the access number that was sent to your email address."` |
| `EmailCodeSubmit` | `*string` | Optional | Button label for submitting the email access code<br><br>**Default**: `"Sign In"` |
| `EmailCodeTitle` | `*string` | Optional | Title shown on the email access-code entry page<br><br>**Default**: `"Access Code"` |
| `EmailError` | `*string` | Optional | Error message when email not provided<br><br>**Default**: `"Please provide valid email"` |
| `EmailFieldLabel` | `*string` | Optional | Label for the email address input field<br><br>**Default**: `"Enter your email address"` |
| `EmailLabel` | `*string` | Optional | Label displayed for the email input field<br><br>**Default**: `"Email"` |
| `EmailMessage` | `*string` | Optional | Instructional text explaining email access-code delivery<br><br>**Default**: `"We will email you an authentication code which you can use to connect to the Wi-Fi network."` |
| `EmailSubmit` | `*string` | Optional | Label for confirmation code submit button using email auth<br><br>**Default**: `"Send Access Code"` |
| `EmailTitle` | `*string` | Optional | Title for the Email registration<br><br>**Default**: `"Sign in with Email"` |
| `EnGB` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | Localized portal template strings for a specific language |
| `EnUS` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | Localized portal template strings for a specific language |
| `EsES` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | Localized portal template strings for a specific language |
| `FiFI` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | Localized portal template strings for a specific language |
| `Field1` | `*bool` | Optional | Whether to ask field1<br><br>**Default**: `false` |
| `Field1Error` | `*string` | Optional | Error message when field1 not provided<br><br>**Default**: `"Please provide Custom Field 1"` |
| `Field1Label` | `*string` | Optional | Label for custom field 1 input<br><br>**Default**: `"Custom Field 1"` |
| `Field1Required` | `*bool` | Optional | Whether custom field 1 must be provided when the field is shown |
| `Field2` | `*bool` | Optional | Whether to ask field2<br><br>**Default**: `false` |
| `Field2Error` | `*string` | Optional | Error message when field2 not provided<br><br>**Default**: `"Please provide Custom Field 2"` |
| `Field2Label` | `*string` | Optional | Label for custom field 2 input<br><br>**Default**: `"Custom Field 2"` |
| `Field2Required` | `*bool` | Optional | Whether custom field 2 must be provided when the field is shown |
| `Field3` | `*bool` | Optional | Whether to ask field3<br><br>**Default**: `false` |
| `Field3Error` | `*string` | Optional | Error message when field3 not provided<br><br>**Default**: `"Please provide Custom Field 3"` |
| `Field3Label` | `*string` | Optional | Label for custom field 3 input<br><br>**Default**: `"Custom Field 3"` |
| `Field3Required` | `*bool` | Optional | Whether custom field 3 must be provided when the field is shown |
| `Field4` | `*bool` | Optional | Whether to ask field4<br><br>**Default**: `false` |
| `Field4Error` | `*string` | Optional | Error message when field4 not provided<br><br>**Default**: `"Please provide Custom Field 4"` |
| `Field4Label` | `*string` | Optional | Label for custom field 4 input<br><br>**Default**: `"Custom Field 4"` |
| `Field4Required` | `*bool` | Optional | Whether custom field 4 must be provided when the field is shown |
| `FrFR` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | Localized portal template strings for a specific language |
| `HeIL` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | Localized portal template strings for a specific language |
| `HiIN` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | Localized portal template strings for a specific language |
| `HrHR` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | Localized portal template strings for a specific language |
| `HuHU` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | Localized portal template strings for a specific language |
| `IdID` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | Localized portal template strings for a specific language |
| `ItIT` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | Localized portal template strings for a specific language |
| `JaJP` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | Localized portal template strings for a specific language |
| `KoKR` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | Localized portal template strings for a specific language |
| `Logo` | `models.Optional[string]` | Optional | Custom logo with `data:image/png;base64,` format, default null, uses Juniper Mist Logo. File size must be less than 100kB and image dimensions must be less than 500px x 200px (width x height). |
| `LogoHeight` | `*int` | Optional | Height of the logo, in px<br><br>**Constraints**: `>= 0`, `<= 200` |
| `LogoWidth` | `*int` | Optional | Width of the logo, in px<br><br>**Constraints**: `>= 0`, `<= 500` |
| `MarketingPolicyLink` | `*string` | Optional | label of the link to go to /marketing_policy<br><br>**Default**: `"Marketing Policy"` |
| `MarketingPolicyOptIn` | `*bool` | Optional | Whether marketing policy optin is enabled<br><br>**Default**: `false` |
| `MarketingPolicyOptInLabel` | `*string` | Optional | label for marketing optin<br><br>**Default**: `"I wish to receive Marketing notifications"` |
| `MarketingPolicyOptInText` | `*string` | Optional | Text of the marketing policy opt-in content<br><br>**Default**: `"Marketing policy content"` |
| `Message` | `*string` | Optional | Main message displayed on the guest portal sign-in page<br><br>**Default**: `"Sign in to get online"` |
| `MsMY` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | Localized portal template strings for a specific language |
| `MultiAuth` | `*bool` | Optional | Whether the portal presents multiple authentication methods<br><br>**Default**: `false` |
| `Name` | `*bool` | Optional | Whether name field is required<br><br>**Default**: `false` |
| `NameError` | `*string` | Optional | Error message when name not provided<br><br>**Default**: `"Please provide your name"` |
| `NameLabel` | `*string` | Optional | Label displayed for the name input field<br><br>**Default**: `"Name"` |
| `NbNO` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | Localized portal template strings for a specific language |
| `NlNL` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | Localized portal template strings for a specific language |
| `OptOutDefault` | `*bool` | Optional | Default value for the `Do not store` checkbox<br><br>**Default**: `true` |
| `Optout` | `*bool` | Optional | Whether to display Do Not Store My Personal Information<br><br>**Default**: `false` |
| `OptoutLabel` | `*string` | Optional | Label for Do Not Store My Personal Information<br><br>**Default**: `"Do not store"` |
| `PageTitle` | `string` | Required | Browser or page title shown for the guest portal<br><br>**Default**: `"Welcome"` |
| `PassphraseCancel` | `*string` | Optional | Label for the Passphrase cancel button<br><br>**Default**: `"Cancel"` |
| `PassphraseError` | `*string` | Optional | Error message when invalid passphrase is provided<br><br>**Default**: `"Invalid Passphrase"` |
| `PassphraseLabel` | `*string` | Optional | Label for the passphrase input field<br><br>**Default**: `"Passphrase"` |
| `PassphraseMessage` | `*string` | Optional | Instructional text shown on the passphrase sign-in page<br><br>**Default**: `"Enter the secret passphrase to access the Wi-Fi network."` |
| `PassphraseSubmit` | `*string` | Optional | Label for the Passphrase submit button<br><br>**Default**: `"Sign in"` |
| `PassphraseTitle` | `*string` | Optional | Title for passphrase details page<br><br>**Default**: `"Sign in with Passphrase"` |
| `PlPL` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | Localized portal template strings for a specific language |
| `PoweredBy` | `*bool` | Optional | Whether to show \"Powered by Mist\"<br><br>**Default**: `true` |
| `Privacy` | `*bool` | Optional | Whether to require the Privacy Term acceptance<br><br>**Default**: `false` |
| `PrivacyPolicyAcceptLabel` | `*string` | Optional | Prefix of the label of the link to go to Privacy Policy<br><br>**Default**: `"I accept the Privacy Terms"` |
| `PrivacyPolicyError` | `*string` | Optional | Error message when Privacy Policy not accepted<br><br>**Default**: `"Please review and accept the Privacy Terms"` |
| `PrivacyPolicyLink` | `*string` | Optional | Label of the link to go to Privacy Policy<br><br>**Default**: `"Privacy Terms"` |
| `PrivacyPolicyText` | `*string` | Optional | Text of the Privacy Policy<br><br>**Default**: `"<< provide your Privacy Terms here >>"` |
| `PtBR` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | Localized portal template strings for a specific language |
| `PtPT` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | Localized portal template strings for a specific language |
| `RequiredFieldLabel` | `*string` | Optional | Text used to mark a form field as required<br><br>**Default**: `"required"` |
| `ResponsiveLayout` | `*bool` | Optional | Whether the portal template uses a responsive layout<br><br>**Default**: `true` |
| `RoRO` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | Localized portal template strings for a specific language |
| `RuRU` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | Localized portal template strings for a specific language |
| `SignInLabel` | `*string` | Optional | Label of the button to signin<br><br>**Default**: `"Sign In"` |
| `SkSK` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | Localized portal template strings for a specific language |
| `SmsCarrierDefault` | `*string` | Optional | Default option text shown in the SMS carrier selector<br><br>**Default**: `"Please Select"` |
| `SmsCarrierError` | `*string` | Optional | Error message shown when no mobile carrier is selected<br><br>**Default**: `"Please select a mobile carrier"` |
| `SmsCarrierFieldLabel` | `*string` | Optional | Label for mobile carrier drop-down list<br><br>**Default**: `"Mobile Carrier"` |
| `SmsCodeCancel` | `*string` | Optional | Label for cancel confirmation code submission<br><br>**Default**: `"I did not receive the code"` |
| `SmsCodeError` | `*string` | Optional | Error message when confirmation code is invalid<br><br>**Default**: `"Invalid Access Code"` |
| `SmsCodeFieldLabel` | `*string` | Optional | Label for the SMS confirmation-code input field<br><br>**Default**: `"Confirmation Code"` |
| `SmsCodeMessage` | `*string` | Optional | Instructional text shown before entering the SMS access code<br><br>**Default**: `"Enter the access number that was sent to your mobile number."` |
| `SmsCodeSubmit` | `*string` | Optional | Label for confirmation code submit button<br><br>**Default**: `"Sign In"` |
| `SmsCodeTitle` | `*string` | Optional | Title shown on the SMS access-code entry page<br><br>**Default**: `"Access Code"` |
| `SmsCountryFieldLabel` | `*string` | Optional | Label for the SMS country-code input field<br><br>**Default**: `"Country Code"` |
| `SmsCountryFormat` | `*string` | Optional | Example country code format shown for SMS authentication<br><br>**Default**: `"+1"` |
| `SmsHaveAccessCode` | `*string` | Optional | Label for checkbox to specify that the user has access code<br><br>**Default**: `"I have an access code"` |
| `SmsIsTwilio` | `*bool` | Optional | Whether the SMS portal flow uses Twilio-specific behavior<br><br>**Default**: `false` |
| `SmsMessageFormat` | `*string` | Optional | Format of access code sms message. {{code}} and {{duration}} are placeholders and should be retained as is. |
| `SmsNumberCancel` | `*string` | Optional | Label for canceling mobile details for SMS auth<br><br>**Default**: `"Cancel"` |
| `SmsNumberError` | `*string` | Optional | Error message shown when the mobile number is invalid<br><br>**Default**: `"Invalid Mobile Number"` |
| `SmsNumberFieldLabel` | `*string` | Optional | Label for field to provide mobile number<br><br>**Default**: `"Mobile Number"` |
| `SmsNumberFormat` | `*string` | Optional | Example mobile number format shown for SMS authentication<br><br>**Default**: `"2125551212 (digits only)"` |
| `SmsNumberMessage` | `*string` | Optional | Instructional text explaining SMS access-code delivery<br><br>**Default**: `"We will send an access code to your mobile number which you can use to connect to the Wi-Fi network. Message and data rates may apply."` |
| `SmsNumberSubmit` | `*string` | Optional | Label for submit button for code generation<br><br>**Default**: `"Send Access Code"` |
| `SmsNumberTitle` | `*string` | Optional | Title for phone number details<br><br>**Default**: `"Sign in with Text Message"` |
| `SmsUsernameFormat` | `*string` | Optional | Example username format shown for SMS authentication<br><br>**Default**: `"username"` |
| `SmsValidityDuration` | `*int` | Optional | How long confirmation code should be considered valid (in minutes)<br><br>**Constraints**: `>= 1`, `<= 30` |
| `SponsorBackLink` | `*string` | Optional | Link text for returning to edit the sponsor request form<br><br>**Default**: `"Go back and edit request form"` |
| `SponsorCancel` | `*string` | Optional | Button label for canceling sponsor authentication<br><br>**Default**: `"Cancel"` |
| `SponsorEmail` | `*string` | Optional | Label for Sponsor Email<br><br>**Default**: `"Sponsor Email"` |
| `SponsorEmailError` | `*string` | Optional | Error message shown when the sponsor email address is invalid<br><br>**Default**: `"Please provide valid sponsor email"` |
| `SponsorEmailTemplate` | `*string` | Optional | HTML template to replace/override default sponsor email template<br>Sponsor Email Template supports following template variables:<br><br>* `approve_url`: Renders URL to approve the request; optionally &minutes=N query param can be appended to change the Authorization period of the guest, where N is a valid integer denoting number of minutes a guest remains authorized<br>* `deny_url`: Renders URL to reject the request<br>* `guest_email`: Renders Email ID of the guest<br>* `guest_name`: Renders Name of the guest<br>* `field1`: Renders value of the Custom Field 1<br>* `field2`: Renders value of the Custom Field 2<br>* `sponsor_link_validity_duration`: Renders validity time of the request (i.e. Approve/Deny URL)<br>* `auth_expire_minutes`: Renders Wlan-level configured Guest Authorization Expiration time period (in minutes), If not configured then default (1 day in minutes) |
| `SponsorInfoApproved` | `*string` | Optional | Status message prefix shown when a sponsor approves the request<br><br>**Default**: `"Your request was approved by"` |
| `SponsorInfoDenied` | `*string` | Optional | Status message prefix shown when a sponsor denies the request<br><br>**Default**: `"Your request was denied by"` |
| `SponsorInfoPending` | `*string` | Optional | Status message prefix shown after a sponsor notification is sent<br><br>**Default**: `"Your notification has been sent to"` |
| `SponsorName` | `*string` | Optional | Label for Sponsor Name<br><br>**Default**: `"Sponsor Name"` |
| `SponsorNameError` | `*string` | Optional | Error message shown when the sponsor name is missing<br><br>**Default**: `"Please provide sponsor name"` |
| `SponsorNotePending` | `*string` | Optional | Additional status text shown while sponsor approval is pending<br><br>**Default**: `"Please wait for them to acknowledge."` |
| `SponsorRequestAccess` | `*string` | Optional | Submit button label request Wifi Access and notify sponsor about guest request<br><br>**Default**: `"Request Wi-Fi Access"` |
| `SponsorStatusApproved` | `*string` | Optional | Text to display if sponsor approves request<br><br>**Default**: `"Your request was approved"` |
| `SponsorStatusDenied` | `*string` | Optional | Text to display when sponsor denies request<br><br>**Default**: `"Your request was denied"` |
| `SponsorStatusPending` | `*string` | Optional | Text to display if request is still pending<br><br>**Default**: `"Notification Sent"` |
| `SponsorSubmit` | `*string` | Optional | Submit button label to notify sponsor about guest request<br><br>**Default**: `"Request Wi-Fi Access"` |
| `SponsorsError` | `*string` | Optional | Error message shown when no sponsor is selected<br><br>**Default**: `"Please select a sponsor"` |
| `SponsorsFieldLabel` | `*string` | Optional | Label for the sponsor selection field<br><br>**Default**: `"Sponsors"` |
| `SvSE` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | Localized portal template strings for a specific language |
| `ThTH` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | Localized portal template strings for a specific language |
| `Tos` | `*bool` | Optional | Whether the portal requires Terms of Service acceptance<br><br>**Default**: `true` |
| `TosAcceptLabel` | `*string` | Optional | Prefix of the label of the link to go to tos<br><br>**Default**: `"I accept the Terms of Service"` |
| `TosError` | `*string` | Optional | Error message when tos not accepted<br><br>**Default**: `"Please review and accept the Terms of Service"` |
| `TosLink` | `*string` | Optional | Label of the link to go to tos<br><br>**Default**: `"Terms of Service"` |
| `TosText` | `*string` | Optional | Text of the Terms of Service<br><br>**Default**: `"<< provide your Terms of Service here >>"` |
| `TrTR` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | Localized portal template strings for a specific language |
| `UkUA` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | Localized portal template strings for a specific language |
| `ViVN` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | Localized portal template strings for a specific language |
| `ZhHans` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | Localized portal template strings for a specific language |
| `ZhHant` | [`*models.WlanPortalTemplateSettingLocale`](../../doc/models/wlan-portal-template-setting-locale.md) | Optional | Localized portal template strings for a specific language |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    wlanPortalTemplateSetting := models.WlanPortalTemplateSetting{
        AccessCodeAlternateEmail:  models.ToPointer("Use alternate email address"),
        Alignment:                 models.ToPointer(models.PortalTemplateAlignmentEnum_CENTER),
        Ar:                        models.ToPointer(models.WlanPortalTemplateSettingLocale{
            AuthButtonAmazon:          models.ToPointer("authButtonAmazon6"),
            AuthButtonAzure:           models.ToPointer("authButtonAzure4"),
            AuthButtonEmail:           models.ToPointer("authButtonEmail2"),
            AuthButtonFacebook:        models.ToPointer("authButtonFacebook6"),
            AuthButtonGoogle:          models.ToPointer("authButtonGoogle2"),
        }),
        AuthButtonAmazon:          models.ToPointer("Sign in with Amazon"),
        AuthButtonAzure:           models.ToPointer("Sign in with Azure"),
        AuthButtonEmail:           models.ToPointer("Sign in with Email"),
        AuthButtonFacebook:        models.ToPointer("Sign in with Facebook"),
        AuthButtonGoogle:          models.ToPointer("Sign in with Google"),
        AuthButtonMicrosoft:       models.ToPointer("Sign in with Microsoft"),
        AuthButtonPassphrase:      models.ToPointer("Sign in with Passphrase"),
        AuthButtonSms:             models.ToPointer("Sign in with Text Message"),
        AuthButtonSponsor:         models.ToPointer("Sign in as Guest"),
        AuthLabel:                 models.ToPointer("Connect to Wi-Fi with"),
        Color:                     models.ToPointer("#1074bc"),
        ColorDark:                 models.ToPointer("#0b5183"),
        ColorLight:                models.ToPointer("#3589c6"),
        Company:                   models.ToPointer(false),
        CompanyError:              models.ToPointer("Please provide your company name"),
        CompanyLabel:              models.ToPointer("Company"),
        Email:                     models.ToPointer(false),
        EmailAccessDomainError:    models.ToPointer("Email Access Domain Error"),
        EmailCancel:               models.ToPointer("Cancel"),
        EmailCodeCancel:           models.ToPointer("I did not receive the code"),
        EmailCodeError:            models.ToPointer("Please provide valid alternate email"),
        EmailCodeFieldLabel:       models.ToPointer("Access Code"),
        EmailCodeMessage:          models.ToPointer("Enter the access number that was sent to your email address."),
        EmailCodeSubmit:           models.ToPointer("Sign In"),
        EmailCodeTitle:            models.ToPointer("Access Code"),
        EmailError:                models.ToPointer("Please provide valid email"),
        EmailFieldLabel:           models.ToPointer("Enter your email address"),
        EmailLabel:                models.ToPointer("Email"),
        EmailMessage:              models.ToPointer("We will email you an authentication code which you can use to connect to the Wi-Fi network."),
        EmailSubmit:               models.ToPointer("Send Access Code"),
        EmailTitle:                models.ToPointer("Sign in with Email"),
        Field1:                    models.ToPointer(false),
        Field1Error:               models.ToPointer("Please provide Custom Field 1"),
        Field1Label:               models.ToPointer("Custom Field 1"),
        Field2:                    models.ToPointer(false),
        Field2Error:               models.ToPointer("Please provide Custom Field 2"),
        Field2Label:               models.ToPointer("Custom Field 2"),
        Field3:                    models.ToPointer(false),
        Field3Error:               models.ToPointer("Please provide Custom Field 3"),
        Field3Label:               models.ToPointer("Custom Field 3"),
        Field4:                    models.ToPointer(false),
        Field4Error:               models.ToPointer("Please provide Custom Field 4"),
        Field4Label:               models.ToPointer("Custom Field 4"),
        Logo:                      models.NewOptional(models.ToPointer("data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAZg…")),
        LogoHeight:                models.ToPointer(123),
        LogoWidth:                 models.ToPointer(408),
        MarketingPolicyLink:       models.ToPointer("Marketing Policy"),
        MarketingPolicyOptIn:      models.ToPointer(false),
        MarketingPolicyOptInLabel: models.ToPointer("I wish to receive Marketing notifications"),
        MarketingPolicyOptInText:  models.ToPointer("Marketing policy content"),
        Message:                   models.ToPointer("Sign in to get online"),
        MultiAuth:                 models.ToPointer(false),
        Name:                      models.ToPointer(false),
        NameError:                 models.ToPointer("Please provide your name"),
        NameLabel:                 models.ToPointer("Name"),
        OptOutDefault:             models.ToPointer(true),
        Optout:                    models.ToPointer(false),
        OptoutLabel:               models.ToPointer("Do not store"),
        PageTitle:                 "Welcome",
        PassphraseCancel:          models.ToPointer("Cancel"),
        PassphraseError:           models.ToPointer("Invalid Passphrase"),
        PassphraseLabel:           models.ToPointer("Passphrase"),
        PassphraseMessage:         models.ToPointer("Enter the secret passphrase to access the Wi-Fi network."),
        PassphraseSubmit:          models.ToPointer("Sign in"),
        PassphraseTitle:           models.ToPointer("Sign in with Passphrase"),
        PoweredBy:                 models.ToPointer(true),
        Privacy:                   models.ToPointer(false),
        PrivacyPolicyAcceptLabel:  models.ToPointer("I accept the Privacy Terms"),
        PrivacyPolicyError:        models.ToPointer("Please review and accept the Privacy Terms"),
        PrivacyPolicyLink:         models.ToPointer("Privacy Terms"),
        PrivacyPolicyText:         models.ToPointer("<< provide your Privacy Terms here >>"),
        RequiredFieldLabel:        models.ToPointer("required"),
        ResponsiveLayout:          models.ToPointer(true),
        SignInLabel:               models.ToPointer("Sign In"),
        SmsCarrierDefault:         models.ToPointer("Please Select"),
        SmsCarrierError:           models.ToPointer("Please select a mobile carrier"),
        SmsCarrierFieldLabel:      models.ToPointer("Mobile Carrier"),
        SmsCodeCancel:             models.ToPointer("I did not receive the code"),
        SmsCodeError:              models.ToPointer("Invalid Access Code"),
        SmsCodeFieldLabel:         models.ToPointer("Confirmation Code"),
        SmsCodeMessage:            models.ToPointer("Enter the access number that was sent to your mobile number."),
        SmsCodeSubmit:             models.ToPointer("Sign In"),
        SmsCodeTitle:              models.ToPointer("Access Code"),
        SmsCountryFieldLabel:      models.ToPointer("Country Code"),
        SmsCountryFormat:          models.ToPointer("+1"),
        SmsHaveAccessCode:         models.ToPointer("I have an access code"),
        SmsIsTwilio:               models.ToPointer(false),
        SmsNumberCancel:           models.ToPointer("Cancel"),
        SmsNumberError:            models.ToPointer("Invalid Mobile Number"),
        SmsNumberFieldLabel:       models.ToPointer("Mobile Number"),
        SmsNumberFormat:           models.ToPointer("2125551212 (digits only)"),
        SmsNumberMessage:          models.ToPointer("We will send an access code to your mobile number which you can use to connect to the Wi-Fi network. Message and data rates may apply."),
        SmsNumberSubmit:           models.ToPointer("Send Access Code"),
        SmsNumberTitle:            models.ToPointer("Sign in with Text Message"),
        SmsUsernameFormat:         models.ToPointer("username"),
        SponsorBackLink:           models.ToPointer("Go back and edit request form"),
        SponsorCancel:             models.ToPointer("Cancel"),
        SponsorEmail:              models.ToPointer("Sponsor Email"),
        SponsorEmailError:         models.ToPointer("Please provide valid sponsor email"),
        SponsorInfoApproved:       models.ToPointer("Your request was approved by"),
        SponsorInfoDenied:         models.ToPointer("Your request was denied by"),
        SponsorInfoPending:        models.ToPointer("Your notification has been sent to"),
        SponsorName:               models.ToPointer("Sponsor Name"),
        SponsorNameError:          models.ToPointer("Please provide sponsor name"),
        SponsorNotePending:        models.ToPointer("Please wait for them to acknowledge."),
        SponsorRequestAccess:      models.ToPointer("Request Wi-Fi Access"),
        SponsorStatusApproved:     models.ToPointer("Your request was approved"),
        SponsorStatusDenied:       models.ToPointer("Your request was denied"),
        SponsorStatusPending:      models.ToPointer("Notification Sent"),
        SponsorSubmit:             models.ToPointer("Request Wi-Fi Access"),
        SponsorsError:             models.ToPointer("Please select a sponsor"),
        SponsorsFieldLabel:        models.ToPointer("Sponsors"),
        Tos:                       models.ToPointer(true),
        TosAcceptLabel:            models.ToPointer("I accept the Terms of Service"),
        TosError:                  models.ToPointer("Please review and accept the Terms of Service"),
        TosLink:                   models.ToPointer("Terms of Service"),
        TosText:                   models.ToPointer("<< provide your Terms of Service here >>"),
    }

}
```

