# PortalTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessCodeAlternateEmail** | Pointer to **string** |  | [optional] [default to "Please provide valid alternate email"]
**Alignment** | Pointer to **string** | defines alignment on portal. “left” is default. | [optional] 
**AuthButtonAmazon** | Pointer to **string** | label for Amazon auth button | [optional] [default to "Sign in with Amazon"]
**AuthButtonAzure** | Pointer to **string** | label for Azure auth button | [optional] [default to "Sign in with Azure"]
**AuthButtonEmail** | Pointer to **string** | label for Email auth button | [optional] [default to "Sign in with Email"]
**AuthButtonFacebook** | Pointer to **string** | label for Facebook auth button | [optional] [default to "Sign in with Facebook"]
**AuthButtonGoogle** | Pointer to **string** | label for Google auth button | [optional] [default to "Sign in with Google"]
**AuthButtonMicrosoft** | Pointer to **string** | label for Microsoft auth button | [optional] [default to "Sign in with Microsoft"]
**AuthButtonPassphrase** | Pointer to **string** | label for passphrase auth button | [optional] 
**AuthButtonSms** | Pointer to **string** | label for SMS auth button | [optional] [default to "Sign in with Text Message"]
**AuthButtonSponsor** | Pointer to **string** | label for Sponsor auth button | [optional] [default to "Sign in as Guest"]
**AuthLabel** | Pointer to **string** |  | [optional] [default to "“Connect to WiFi with”"]
**BackLink** | Pointer to **string** | label of the link to go back to /logon | [optional] [default to "Go back and edit request form"]
**Color** | Pointer to **string** |  | [optional] [default to "“#1074bc”"]
**ColorDark** | Pointer to **string** |  | [optional] [default to "“#0b5183”"]
**ColorLight** | Pointer to **string** |  | [optional] [default to "“#3589c6”"]
**Company** | Pointer to **bool** | whether company field is required | [optional] [default to false]
**CompanyError** | Pointer to **string** | error message when company not provided | [optional] [default to "Please provide company name"]
**CompanyLabel** | Pointer to **string** | label of company field | [optional] [default to "Company"]
**CreatedTime** | Pointer to **float32** |  | [optional] [readonly] 
**Email** | Pointer to **bool** | whether email field is required | [optional] [default to false]
**EmailAccessDomainError** | Pointer to **string** | error message when a user has valid social login but doesn’t match specified email domains. | [optional] [default to "Access is restricted by email domain"]
**EmailCancel** | Pointer to **string** |  | [optional] [default to "Cancel"]
**EmailCodeError** | Pointer to **string** |  | [optional] [default to "Please provide valid alternate email"]
**EmailError** | Pointer to **string** | error message when email not provided | [optional] [default to "Please provide valid email"]
**EmailFieldLabel** | Pointer to **string** |  | [optional] [default to "Enter your email address"]
**EmailLabel** | Pointer to **string** | label of email field | [optional] [default to "Email"]
**EmailMessage** | Pointer to **string** |  | [optional] [default to "We will email you an authentication code which you can use to connect to the WiFi network."]
**EmailSubmit** | Pointer to **string** |  | [optional] [default to "“Send Access Code”"]
**EmailTitle** | Pointer to **string** |  | [optional] [default to "“Sign in with Email”"]
**Field1** | Pointer to **bool** | whether to ask field1 | [optional] [default to false]
**Field1Error** | Pointer to **string** | error message when field1 not provided | [optional] [default to "Please provide field1"]
**Field1Label** | Pointer to **string** | label of field1 | [optional] [default to "Custom1"]
**Field1Required** | Pointer to **bool** | whether field1 is required field | [optional] [default to false]
**Field2** | Pointer to **bool** | whether to ask field2 | [optional] [default to false]
**Field2Error** | Pointer to **string** | error message when field2 not provided | [optional] [default to "Please provide field2"]
**Field2Label** | Pointer to **string** | label of field2 | [optional] [default to "Custom2"]
**Field2Required** | Pointer to **bool** | whether field2 is required field | [optional] [default to false]
**Field3** | Pointer to **bool** | whether to ask field3 | [optional] [default to false]
**Field3Error** | Pointer to **string** | error message when field3 not provided | [optional] [default to "Please provide field3"]
**Field3Label** | Pointer to **string** | label of field3 | [optional] [default to "Custom3"]
**Field3Required** | Pointer to **bool** | whether field3 is required field | [optional] [default to false]
**Field4** | Pointer to **bool** | whether to ask field4 | [optional] [default to false]
**Field4Error** | Pointer to **string** | error message when field4 not provided | [optional] [default to "Please provide field4"]
**Field4Label** | Pointer to **string** | label of field4 | [optional] [default to "Custom4"]
**Field4Required** | Pointer to **bool** | whether field4 is required field | [optional] [default to false]
**ForSite** | Pointer to **bool** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Message** | Pointer to **string** |  | [optional] [default to "Please enjoy the complimentary Wifi"]
**ModifiedTime** | Pointer to **float32** |  | [optional] [readonly] 
**Name** | Pointer to **bool** | whether name field is required | [optional] [default to false]
**NameError** | Pointer to **string** | error message when name not provided | [optional] [default to "Please provide your name"]
**NameLabel** | Pointer to **string** | label of name field | [optional] [default to "Name"]
**Optout** | Pointer to **bool** | whether to display “Do Not Store My Personal Information” | [optional] 
**OptoutLabel** | Pointer to **string** |  | [optional] [default to "Do Not Store My Personal Information"]
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**PageTitle** | **string** |  | [default to "Welcome"]
**PassphraseCancel** | Pointer to **string** |  | [optional] [default to "Cancel"]
**PassphraseError** | Pointer to **string** | error message when invalid passphrase is provided | [optional] [default to "Invalid Passphrase"]
**PassphraseLabel** | Pointer to **string** |  | [optional] [default to "Passphrase"]
**PassphraseMessage** | Pointer to **string** |  | [optional] [default to "Login using passphrase"]
**PassphraseSubmit** | Pointer to **string** |  | [optional] [default to "Sign in"]
**PassphraseTitle** | Pointer to **string** | Title for passphrase details page | [optional] [default to "Sign in with Passphrase"]
**PoweredBy** | Pointer to **bool** | whether to show “Powered by Mist” | [optional] [default to true]
**PrivacyPolicy** | Pointer to **bool** |  | [optional] [default to false]
**PrivacyPolicyAcceptLabel** | Pointer to **string** | prefix of the label of the link to go to /privacy_notice | [optional] 
**PrivacyPolicyError** | Pointer to **string** | error message when privacy notice is not accepted | [optional] 
**PrivacyPolicyLink** | Pointer to **string** | label of the link to go to /privacy_notice, | [optional] 
**PrivacyPolicyText** | Pointer to **string** | privacy notice text | [optional] 
**RequiredFieldLabel** | Pointer to **string** | label to denote required field | [optional] [default to "required"]
**SignInLabel** | Pointer to **string** | label of the button to /signin | [optional] [default to "Sign In"]
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**SmsCarrierDefault** | Pointer to **string** |  | [optional] [default to "Please Select"]
**SmsCarrierError** | Pointer to **string** |  | [optional] [default to "Please select a mobile carrier"]
**SmsCarrierFieldLabel** | Pointer to **string** | label for mobile carrier drop-down list | [optional] [default to "Mobile Carrier"]
**SmsCodeCancel** | Pointer to **string** | Label for cancel confirmation code submission | [optional] [default to "Cancel"]
**SmsCodeError** | Pointer to **string** | error message when confirmation code is invalid | [optional] [default to "Invalid Access Code"]
**SmsCodeFieldLabel** | Pointer to **string** |  | [optional] [default to "Confirmation Code"]
**SmsCodeMessage** | Pointer to **string** |  | [optional] [default to "Enter the confirmation code"]
**SmsCodeSubmit** | Pointer to **string** | Label for confirmation code submit button | [optional] [default to "Submit Code"]
**SmsCodeTitle** | Pointer to **string** |  | [optional] [default to "Access Code"]
**SmsCountryFieldLabel** | Pointer to **string** | “Country Code” | [optional] 
**SmsCountryFormat** | Pointer to **string** |  | [optional] [default to "“+1”"]
**SmsHaveAccessCode** | Pointer to **string** | Label for checkbox to specify that the user has access code | [optional] [default to "I have an access code"]
**SmsMessageFormat** | Pointer to **string** | format of access code sms message. {{code}} and {{duration}} are place holders and should be retained as is. | [optional] [default to "Code {{code}} expires in {{duration}} minutes."]
**SmsNumberCancel** | Pointer to **string** | label for canceling mobile details for SMS auth | [optional] [default to "Cancel"]
**SmsNumberError** | Pointer to **string** |  | [optional] [default to "Invalid Mobile Number"]
**SmsNumberFieldLabel** | Pointer to **string** | label for field to provide mobile number | [optional] [default to "Mobile Number"]
**SmsNumberFormat** | Pointer to **string** |  | [optional] [default to "“2125551212 (digits only)”"]
**SmsNumberMessage** | Pointer to **string** |  | [optional] [default to "We will send an access code to your mobile number which you can use to connect to the WiFi network. Message and data rates may apply."]
**SmsNumberSubmit** | Pointer to **string** | label for submit button for code generation | [optional] [default to "Sign In"]
**SmsNumberTitle** | Pointer to **string** | Title for phone number details | [optional] [default to "Text Message Confirmation"]
**SmsUsernameFormat** | Pointer to **string** |  | [optional] [default to "username"]
**SmsValidityDuration** | Pointer to **float32** | how long confirmation code should be considered valid (in minutes) | [optional] [default to 5]
**SponsorAutoApproved** | Pointer to **string** |  | [optional] [default to "Your request was approved."]
**SponsorAutoApprovedNote** | Pointer to **string** |  | [optional] [default to "Your notification has been sent to"]
**SponsorBackLink** | Pointer to **string** |  | [optional] [default to "Go back and edit request form"]
**SponsorCancel** | Pointer to **string** |  | [optional] [default to "Cancel"]
**SponsorEmail** | Pointer to **string** | label for Sponsor Email | [optional] [default to "Please provide valid sponsor email"]
**SponsorEmailError** | Pointer to **string** |  | [optional] [default to "Please provide valid sponsor email"]
**SponsorEmailTemplate** | Pointer to **string** | “html template to replace/override default sponsor email template” | [optional] 
**SponsorInfoApproved** | Pointer to **string** |  | [optional] [default to "Your request was approved by"]
**SponsorInfoDenied** | Pointer to **string** |  | [optional] [default to "Your request was denied by"]
**SponsorInfoPending** | Pointer to **string** |  | [optional] [default to "Your notification has been sent to"]
**SponsorName** | Pointer to **string** | label for Sponsor Name | [optional] [default to "Sponsor Name"]
**SponsorNameError** | Pointer to **string** |  | [optional] [default to "Please provide sponsor’s name"]
**SponsorNotePending** | Pointer to **string** |  | [optional] [default to "Please wait for them to acknowledge."]
**SponsorStatusApproved** | Pointer to **string** | text to display if sponsor approves request | [optional] [default to "Your request was approved"]
**SponsorStatusDenied** | Pointer to **string** | text to display when sponsor denies request | [optional] [default to "Your request was denied"]
**SponsorStatusPending** | Pointer to **string** | text to display if request is still pending | [optional] [default to "Notification Sent"]
**SponsorSubmit** | Pointer to **string** | submit button label to notify sponsor about guest request | [optional] [default to "Notify Sponsor"]
**SponsorsAutoApprovedNote** | Pointer to **string** |  | [optional] [default to "Your notification has been sent to the sponsors.\""]
**SponsorsError** | Pointer to **string** |  | [optional] [default to "Please select a sponsor"]
**Tos** | Pointer to **bool** |  | [optional] [default to true]
**TosAcceptLabel** | Pointer to **string** | prefix of the label of the link to go to /tos | [optional] [default to "I accept the Terms of Service"]
**TosError** | Pointer to **string** | error message when tos not accepted | [optional] [default to "Please review and accept terms of service"]
**TosLink** | Pointer to **string** | label of the link to go to /tos | [optional] [default to "Terms of Service"]
**TosText** | Pointer to **string** | text of the Terms of Service | [optional] [default to "terms of service"]

## Methods

### NewPortalTemplate

`func NewPortalTemplate(pageTitle string, ) *PortalTemplate`

NewPortalTemplate instantiates a new PortalTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortalTemplateWithDefaults

`func NewPortalTemplateWithDefaults() *PortalTemplate`

NewPortalTemplateWithDefaults instantiates a new PortalTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessCodeAlternateEmail

`func (o *PortalTemplate) GetAccessCodeAlternateEmail() string`

GetAccessCodeAlternateEmail returns the AccessCodeAlternateEmail field if non-nil, zero value otherwise.

### GetAccessCodeAlternateEmailOk

`func (o *PortalTemplate) GetAccessCodeAlternateEmailOk() (*string, bool)`

GetAccessCodeAlternateEmailOk returns a tuple with the AccessCodeAlternateEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessCodeAlternateEmail

`func (o *PortalTemplate) SetAccessCodeAlternateEmail(v string)`

SetAccessCodeAlternateEmail sets AccessCodeAlternateEmail field to given value.

### HasAccessCodeAlternateEmail

`func (o *PortalTemplate) HasAccessCodeAlternateEmail() bool`

HasAccessCodeAlternateEmail returns a boolean if a field has been set.

### GetAlignment

`func (o *PortalTemplate) GetAlignment() string`

GetAlignment returns the Alignment field if non-nil, zero value otherwise.

### GetAlignmentOk

`func (o *PortalTemplate) GetAlignmentOk() (*string, bool)`

GetAlignmentOk returns a tuple with the Alignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlignment

`func (o *PortalTemplate) SetAlignment(v string)`

SetAlignment sets Alignment field to given value.

### HasAlignment

`func (o *PortalTemplate) HasAlignment() bool`

HasAlignment returns a boolean if a field has been set.

### GetAuthButtonAmazon

`func (o *PortalTemplate) GetAuthButtonAmazon() string`

GetAuthButtonAmazon returns the AuthButtonAmazon field if non-nil, zero value otherwise.

### GetAuthButtonAmazonOk

`func (o *PortalTemplate) GetAuthButtonAmazonOk() (*string, bool)`

GetAuthButtonAmazonOk returns a tuple with the AuthButtonAmazon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthButtonAmazon

`func (o *PortalTemplate) SetAuthButtonAmazon(v string)`

SetAuthButtonAmazon sets AuthButtonAmazon field to given value.

### HasAuthButtonAmazon

`func (o *PortalTemplate) HasAuthButtonAmazon() bool`

HasAuthButtonAmazon returns a boolean if a field has been set.

### GetAuthButtonAzure

`func (o *PortalTemplate) GetAuthButtonAzure() string`

GetAuthButtonAzure returns the AuthButtonAzure field if non-nil, zero value otherwise.

### GetAuthButtonAzureOk

`func (o *PortalTemplate) GetAuthButtonAzureOk() (*string, bool)`

GetAuthButtonAzureOk returns a tuple with the AuthButtonAzure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthButtonAzure

`func (o *PortalTemplate) SetAuthButtonAzure(v string)`

SetAuthButtonAzure sets AuthButtonAzure field to given value.

### HasAuthButtonAzure

`func (o *PortalTemplate) HasAuthButtonAzure() bool`

HasAuthButtonAzure returns a boolean if a field has been set.

### GetAuthButtonEmail

`func (o *PortalTemplate) GetAuthButtonEmail() string`

GetAuthButtonEmail returns the AuthButtonEmail field if non-nil, zero value otherwise.

### GetAuthButtonEmailOk

`func (o *PortalTemplate) GetAuthButtonEmailOk() (*string, bool)`

GetAuthButtonEmailOk returns a tuple with the AuthButtonEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthButtonEmail

`func (o *PortalTemplate) SetAuthButtonEmail(v string)`

SetAuthButtonEmail sets AuthButtonEmail field to given value.

### HasAuthButtonEmail

`func (o *PortalTemplate) HasAuthButtonEmail() bool`

HasAuthButtonEmail returns a boolean if a field has been set.

### GetAuthButtonFacebook

`func (o *PortalTemplate) GetAuthButtonFacebook() string`

GetAuthButtonFacebook returns the AuthButtonFacebook field if non-nil, zero value otherwise.

### GetAuthButtonFacebookOk

`func (o *PortalTemplate) GetAuthButtonFacebookOk() (*string, bool)`

GetAuthButtonFacebookOk returns a tuple with the AuthButtonFacebook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthButtonFacebook

`func (o *PortalTemplate) SetAuthButtonFacebook(v string)`

SetAuthButtonFacebook sets AuthButtonFacebook field to given value.

### HasAuthButtonFacebook

`func (o *PortalTemplate) HasAuthButtonFacebook() bool`

HasAuthButtonFacebook returns a boolean if a field has been set.

### GetAuthButtonGoogle

`func (o *PortalTemplate) GetAuthButtonGoogle() string`

GetAuthButtonGoogle returns the AuthButtonGoogle field if non-nil, zero value otherwise.

### GetAuthButtonGoogleOk

`func (o *PortalTemplate) GetAuthButtonGoogleOk() (*string, bool)`

GetAuthButtonGoogleOk returns a tuple with the AuthButtonGoogle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthButtonGoogle

`func (o *PortalTemplate) SetAuthButtonGoogle(v string)`

SetAuthButtonGoogle sets AuthButtonGoogle field to given value.

### HasAuthButtonGoogle

`func (o *PortalTemplate) HasAuthButtonGoogle() bool`

HasAuthButtonGoogle returns a boolean if a field has been set.

### GetAuthButtonMicrosoft

`func (o *PortalTemplate) GetAuthButtonMicrosoft() string`

GetAuthButtonMicrosoft returns the AuthButtonMicrosoft field if non-nil, zero value otherwise.

### GetAuthButtonMicrosoftOk

`func (o *PortalTemplate) GetAuthButtonMicrosoftOk() (*string, bool)`

GetAuthButtonMicrosoftOk returns a tuple with the AuthButtonMicrosoft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthButtonMicrosoft

`func (o *PortalTemplate) SetAuthButtonMicrosoft(v string)`

SetAuthButtonMicrosoft sets AuthButtonMicrosoft field to given value.

### HasAuthButtonMicrosoft

`func (o *PortalTemplate) HasAuthButtonMicrosoft() bool`

HasAuthButtonMicrosoft returns a boolean if a field has been set.

### GetAuthButtonPassphrase

`func (o *PortalTemplate) GetAuthButtonPassphrase() string`

GetAuthButtonPassphrase returns the AuthButtonPassphrase field if non-nil, zero value otherwise.

### GetAuthButtonPassphraseOk

`func (o *PortalTemplate) GetAuthButtonPassphraseOk() (*string, bool)`

GetAuthButtonPassphraseOk returns a tuple with the AuthButtonPassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthButtonPassphrase

`func (o *PortalTemplate) SetAuthButtonPassphrase(v string)`

SetAuthButtonPassphrase sets AuthButtonPassphrase field to given value.

### HasAuthButtonPassphrase

`func (o *PortalTemplate) HasAuthButtonPassphrase() bool`

HasAuthButtonPassphrase returns a boolean if a field has been set.

### GetAuthButtonSms

`func (o *PortalTemplate) GetAuthButtonSms() string`

GetAuthButtonSms returns the AuthButtonSms field if non-nil, zero value otherwise.

### GetAuthButtonSmsOk

`func (o *PortalTemplate) GetAuthButtonSmsOk() (*string, bool)`

GetAuthButtonSmsOk returns a tuple with the AuthButtonSms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthButtonSms

`func (o *PortalTemplate) SetAuthButtonSms(v string)`

SetAuthButtonSms sets AuthButtonSms field to given value.

### HasAuthButtonSms

`func (o *PortalTemplate) HasAuthButtonSms() bool`

HasAuthButtonSms returns a boolean if a field has been set.

### GetAuthButtonSponsor

`func (o *PortalTemplate) GetAuthButtonSponsor() string`

GetAuthButtonSponsor returns the AuthButtonSponsor field if non-nil, zero value otherwise.

### GetAuthButtonSponsorOk

`func (o *PortalTemplate) GetAuthButtonSponsorOk() (*string, bool)`

GetAuthButtonSponsorOk returns a tuple with the AuthButtonSponsor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthButtonSponsor

`func (o *PortalTemplate) SetAuthButtonSponsor(v string)`

SetAuthButtonSponsor sets AuthButtonSponsor field to given value.

### HasAuthButtonSponsor

`func (o *PortalTemplate) HasAuthButtonSponsor() bool`

HasAuthButtonSponsor returns a boolean if a field has been set.

### GetAuthLabel

`func (o *PortalTemplate) GetAuthLabel() string`

GetAuthLabel returns the AuthLabel field if non-nil, zero value otherwise.

### GetAuthLabelOk

`func (o *PortalTemplate) GetAuthLabelOk() (*string, bool)`

GetAuthLabelOk returns a tuple with the AuthLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthLabel

`func (o *PortalTemplate) SetAuthLabel(v string)`

SetAuthLabel sets AuthLabel field to given value.

### HasAuthLabel

`func (o *PortalTemplate) HasAuthLabel() bool`

HasAuthLabel returns a boolean if a field has been set.

### GetBackLink

`func (o *PortalTemplate) GetBackLink() string`

GetBackLink returns the BackLink field if non-nil, zero value otherwise.

### GetBackLinkOk

`func (o *PortalTemplate) GetBackLinkOk() (*string, bool)`

GetBackLinkOk returns a tuple with the BackLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackLink

`func (o *PortalTemplate) SetBackLink(v string)`

SetBackLink sets BackLink field to given value.

### HasBackLink

`func (o *PortalTemplate) HasBackLink() bool`

HasBackLink returns a boolean if a field has been set.

### GetColor

`func (o *PortalTemplate) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *PortalTemplate) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *PortalTemplate) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *PortalTemplate) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetColorDark

`func (o *PortalTemplate) GetColorDark() string`

GetColorDark returns the ColorDark field if non-nil, zero value otherwise.

### GetColorDarkOk

`func (o *PortalTemplate) GetColorDarkOk() (*string, bool)`

GetColorDarkOk returns a tuple with the ColorDark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorDark

`func (o *PortalTemplate) SetColorDark(v string)`

SetColorDark sets ColorDark field to given value.

### HasColorDark

`func (o *PortalTemplate) HasColorDark() bool`

HasColorDark returns a boolean if a field has been set.

### GetColorLight

`func (o *PortalTemplate) GetColorLight() string`

GetColorLight returns the ColorLight field if non-nil, zero value otherwise.

### GetColorLightOk

`func (o *PortalTemplate) GetColorLightOk() (*string, bool)`

GetColorLightOk returns a tuple with the ColorLight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorLight

`func (o *PortalTemplate) SetColorLight(v string)`

SetColorLight sets ColorLight field to given value.

### HasColorLight

`func (o *PortalTemplate) HasColorLight() bool`

HasColorLight returns a boolean if a field has been set.

### GetCompany

`func (o *PortalTemplate) GetCompany() bool`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *PortalTemplate) GetCompanyOk() (*bool, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *PortalTemplate) SetCompany(v bool)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *PortalTemplate) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCompanyError

`func (o *PortalTemplate) GetCompanyError() string`

GetCompanyError returns the CompanyError field if non-nil, zero value otherwise.

### GetCompanyErrorOk

`func (o *PortalTemplate) GetCompanyErrorOk() (*string, bool)`

GetCompanyErrorOk returns a tuple with the CompanyError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyError

`func (o *PortalTemplate) SetCompanyError(v string)`

SetCompanyError sets CompanyError field to given value.

### HasCompanyError

`func (o *PortalTemplate) HasCompanyError() bool`

HasCompanyError returns a boolean if a field has been set.

### GetCompanyLabel

`func (o *PortalTemplate) GetCompanyLabel() string`

GetCompanyLabel returns the CompanyLabel field if non-nil, zero value otherwise.

### GetCompanyLabelOk

`func (o *PortalTemplate) GetCompanyLabelOk() (*string, bool)`

GetCompanyLabelOk returns a tuple with the CompanyLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyLabel

`func (o *PortalTemplate) SetCompanyLabel(v string)`

SetCompanyLabel sets CompanyLabel field to given value.

### HasCompanyLabel

`func (o *PortalTemplate) HasCompanyLabel() bool`

HasCompanyLabel returns a boolean if a field has been set.

### GetCreatedTime

`func (o *PortalTemplate) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *PortalTemplate) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *PortalTemplate) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *PortalTemplate) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetEmail

`func (o *PortalTemplate) GetEmail() bool`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PortalTemplate) GetEmailOk() (*bool, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PortalTemplate) SetEmail(v bool)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PortalTemplate) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEmailAccessDomainError

`func (o *PortalTemplate) GetEmailAccessDomainError() string`

GetEmailAccessDomainError returns the EmailAccessDomainError field if non-nil, zero value otherwise.

### GetEmailAccessDomainErrorOk

`func (o *PortalTemplate) GetEmailAccessDomainErrorOk() (*string, bool)`

GetEmailAccessDomainErrorOk returns a tuple with the EmailAccessDomainError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAccessDomainError

`func (o *PortalTemplate) SetEmailAccessDomainError(v string)`

SetEmailAccessDomainError sets EmailAccessDomainError field to given value.

### HasEmailAccessDomainError

`func (o *PortalTemplate) HasEmailAccessDomainError() bool`

HasEmailAccessDomainError returns a boolean if a field has been set.

### GetEmailCancel

`func (o *PortalTemplate) GetEmailCancel() string`

GetEmailCancel returns the EmailCancel field if non-nil, zero value otherwise.

### GetEmailCancelOk

`func (o *PortalTemplate) GetEmailCancelOk() (*string, bool)`

GetEmailCancelOk returns a tuple with the EmailCancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailCancel

`func (o *PortalTemplate) SetEmailCancel(v string)`

SetEmailCancel sets EmailCancel field to given value.

### HasEmailCancel

`func (o *PortalTemplate) HasEmailCancel() bool`

HasEmailCancel returns a boolean if a field has been set.

### GetEmailCodeError

`func (o *PortalTemplate) GetEmailCodeError() string`

GetEmailCodeError returns the EmailCodeError field if non-nil, zero value otherwise.

### GetEmailCodeErrorOk

`func (o *PortalTemplate) GetEmailCodeErrorOk() (*string, bool)`

GetEmailCodeErrorOk returns a tuple with the EmailCodeError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailCodeError

`func (o *PortalTemplate) SetEmailCodeError(v string)`

SetEmailCodeError sets EmailCodeError field to given value.

### HasEmailCodeError

`func (o *PortalTemplate) HasEmailCodeError() bool`

HasEmailCodeError returns a boolean if a field has been set.

### GetEmailError

`func (o *PortalTemplate) GetEmailError() string`

GetEmailError returns the EmailError field if non-nil, zero value otherwise.

### GetEmailErrorOk

`func (o *PortalTemplate) GetEmailErrorOk() (*string, bool)`

GetEmailErrorOk returns a tuple with the EmailError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailError

`func (o *PortalTemplate) SetEmailError(v string)`

SetEmailError sets EmailError field to given value.

### HasEmailError

`func (o *PortalTemplate) HasEmailError() bool`

HasEmailError returns a boolean if a field has been set.

### GetEmailFieldLabel

`func (o *PortalTemplate) GetEmailFieldLabel() string`

GetEmailFieldLabel returns the EmailFieldLabel field if non-nil, zero value otherwise.

### GetEmailFieldLabelOk

`func (o *PortalTemplate) GetEmailFieldLabelOk() (*string, bool)`

GetEmailFieldLabelOk returns a tuple with the EmailFieldLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailFieldLabel

`func (o *PortalTemplate) SetEmailFieldLabel(v string)`

SetEmailFieldLabel sets EmailFieldLabel field to given value.

### HasEmailFieldLabel

`func (o *PortalTemplate) HasEmailFieldLabel() bool`

HasEmailFieldLabel returns a boolean if a field has been set.

### GetEmailLabel

`func (o *PortalTemplate) GetEmailLabel() string`

GetEmailLabel returns the EmailLabel field if non-nil, zero value otherwise.

### GetEmailLabelOk

`func (o *PortalTemplate) GetEmailLabelOk() (*string, bool)`

GetEmailLabelOk returns a tuple with the EmailLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailLabel

`func (o *PortalTemplate) SetEmailLabel(v string)`

SetEmailLabel sets EmailLabel field to given value.

### HasEmailLabel

`func (o *PortalTemplate) HasEmailLabel() bool`

HasEmailLabel returns a boolean if a field has been set.

### GetEmailMessage

`func (o *PortalTemplate) GetEmailMessage() string`

GetEmailMessage returns the EmailMessage field if non-nil, zero value otherwise.

### GetEmailMessageOk

`func (o *PortalTemplate) GetEmailMessageOk() (*string, bool)`

GetEmailMessageOk returns a tuple with the EmailMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailMessage

`func (o *PortalTemplate) SetEmailMessage(v string)`

SetEmailMessage sets EmailMessage field to given value.

### HasEmailMessage

`func (o *PortalTemplate) HasEmailMessage() bool`

HasEmailMessage returns a boolean if a field has been set.

### GetEmailSubmit

`func (o *PortalTemplate) GetEmailSubmit() string`

GetEmailSubmit returns the EmailSubmit field if non-nil, zero value otherwise.

### GetEmailSubmitOk

`func (o *PortalTemplate) GetEmailSubmitOk() (*string, bool)`

GetEmailSubmitOk returns a tuple with the EmailSubmit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSubmit

`func (o *PortalTemplate) SetEmailSubmit(v string)`

SetEmailSubmit sets EmailSubmit field to given value.

### HasEmailSubmit

`func (o *PortalTemplate) HasEmailSubmit() bool`

HasEmailSubmit returns a boolean if a field has been set.

### GetEmailTitle

`func (o *PortalTemplate) GetEmailTitle() string`

GetEmailTitle returns the EmailTitle field if non-nil, zero value otherwise.

### GetEmailTitleOk

`func (o *PortalTemplate) GetEmailTitleOk() (*string, bool)`

GetEmailTitleOk returns a tuple with the EmailTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTitle

`func (o *PortalTemplate) SetEmailTitle(v string)`

SetEmailTitle sets EmailTitle field to given value.

### HasEmailTitle

`func (o *PortalTemplate) HasEmailTitle() bool`

HasEmailTitle returns a boolean if a field has been set.

### GetField1

`func (o *PortalTemplate) GetField1() bool`

GetField1 returns the Field1 field if non-nil, zero value otherwise.

### GetField1Ok

`func (o *PortalTemplate) GetField1Ok() (*bool, bool)`

GetField1Ok returns a tuple with the Field1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField1

`func (o *PortalTemplate) SetField1(v bool)`

SetField1 sets Field1 field to given value.

### HasField1

`func (o *PortalTemplate) HasField1() bool`

HasField1 returns a boolean if a field has been set.

### GetField1Error

`func (o *PortalTemplate) GetField1Error() string`

GetField1Error returns the Field1Error field if non-nil, zero value otherwise.

### GetField1ErrorOk

`func (o *PortalTemplate) GetField1ErrorOk() (*string, bool)`

GetField1ErrorOk returns a tuple with the Field1Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField1Error

`func (o *PortalTemplate) SetField1Error(v string)`

SetField1Error sets Field1Error field to given value.

### HasField1Error

`func (o *PortalTemplate) HasField1Error() bool`

HasField1Error returns a boolean if a field has been set.

### GetField1Label

`func (o *PortalTemplate) GetField1Label() string`

GetField1Label returns the Field1Label field if non-nil, zero value otherwise.

### GetField1LabelOk

`func (o *PortalTemplate) GetField1LabelOk() (*string, bool)`

GetField1LabelOk returns a tuple with the Field1Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField1Label

`func (o *PortalTemplate) SetField1Label(v string)`

SetField1Label sets Field1Label field to given value.

### HasField1Label

`func (o *PortalTemplate) HasField1Label() bool`

HasField1Label returns a boolean if a field has been set.

### GetField1Required

`func (o *PortalTemplate) GetField1Required() bool`

GetField1Required returns the Field1Required field if non-nil, zero value otherwise.

### GetField1RequiredOk

`func (o *PortalTemplate) GetField1RequiredOk() (*bool, bool)`

GetField1RequiredOk returns a tuple with the Field1Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField1Required

`func (o *PortalTemplate) SetField1Required(v bool)`

SetField1Required sets Field1Required field to given value.

### HasField1Required

`func (o *PortalTemplate) HasField1Required() bool`

HasField1Required returns a boolean if a field has been set.

### GetField2

`func (o *PortalTemplate) GetField2() bool`

GetField2 returns the Field2 field if non-nil, zero value otherwise.

### GetField2Ok

`func (o *PortalTemplate) GetField2Ok() (*bool, bool)`

GetField2Ok returns a tuple with the Field2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField2

`func (o *PortalTemplate) SetField2(v bool)`

SetField2 sets Field2 field to given value.

### HasField2

`func (o *PortalTemplate) HasField2() bool`

HasField2 returns a boolean if a field has been set.

### GetField2Error

`func (o *PortalTemplate) GetField2Error() string`

GetField2Error returns the Field2Error field if non-nil, zero value otherwise.

### GetField2ErrorOk

`func (o *PortalTemplate) GetField2ErrorOk() (*string, bool)`

GetField2ErrorOk returns a tuple with the Field2Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField2Error

`func (o *PortalTemplate) SetField2Error(v string)`

SetField2Error sets Field2Error field to given value.

### HasField2Error

`func (o *PortalTemplate) HasField2Error() bool`

HasField2Error returns a boolean if a field has been set.

### GetField2Label

`func (o *PortalTemplate) GetField2Label() string`

GetField2Label returns the Field2Label field if non-nil, zero value otherwise.

### GetField2LabelOk

`func (o *PortalTemplate) GetField2LabelOk() (*string, bool)`

GetField2LabelOk returns a tuple with the Field2Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField2Label

`func (o *PortalTemplate) SetField2Label(v string)`

SetField2Label sets Field2Label field to given value.

### HasField2Label

`func (o *PortalTemplate) HasField2Label() bool`

HasField2Label returns a boolean if a field has been set.

### GetField2Required

`func (o *PortalTemplate) GetField2Required() bool`

GetField2Required returns the Field2Required field if non-nil, zero value otherwise.

### GetField2RequiredOk

`func (o *PortalTemplate) GetField2RequiredOk() (*bool, bool)`

GetField2RequiredOk returns a tuple with the Field2Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField2Required

`func (o *PortalTemplate) SetField2Required(v bool)`

SetField2Required sets Field2Required field to given value.

### HasField2Required

`func (o *PortalTemplate) HasField2Required() bool`

HasField2Required returns a boolean if a field has been set.

### GetField3

`func (o *PortalTemplate) GetField3() bool`

GetField3 returns the Field3 field if non-nil, zero value otherwise.

### GetField3Ok

`func (o *PortalTemplate) GetField3Ok() (*bool, bool)`

GetField3Ok returns a tuple with the Field3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField3

`func (o *PortalTemplate) SetField3(v bool)`

SetField3 sets Field3 field to given value.

### HasField3

`func (o *PortalTemplate) HasField3() bool`

HasField3 returns a boolean if a field has been set.

### GetField3Error

`func (o *PortalTemplate) GetField3Error() string`

GetField3Error returns the Field3Error field if non-nil, zero value otherwise.

### GetField3ErrorOk

`func (o *PortalTemplate) GetField3ErrorOk() (*string, bool)`

GetField3ErrorOk returns a tuple with the Field3Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField3Error

`func (o *PortalTemplate) SetField3Error(v string)`

SetField3Error sets Field3Error field to given value.

### HasField3Error

`func (o *PortalTemplate) HasField3Error() bool`

HasField3Error returns a boolean if a field has been set.

### GetField3Label

`func (o *PortalTemplate) GetField3Label() string`

GetField3Label returns the Field3Label field if non-nil, zero value otherwise.

### GetField3LabelOk

`func (o *PortalTemplate) GetField3LabelOk() (*string, bool)`

GetField3LabelOk returns a tuple with the Field3Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField3Label

`func (o *PortalTemplate) SetField3Label(v string)`

SetField3Label sets Field3Label field to given value.

### HasField3Label

`func (o *PortalTemplate) HasField3Label() bool`

HasField3Label returns a boolean if a field has been set.

### GetField3Required

`func (o *PortalTemplate) GetField3Required() bool`

GetField3Required returns the Field3Required field if non-nil, zero value otherwise.

### GetField3RequiredOk

`func (o *PortalTemplate) GetField3RequiredOk() (*bool, bool)`

GetField3RequiredOk returns a tuple with the Field3Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField3Required

`func (o *PortalTemplate) SetField3Required(v bool)`

SetField3Required sets Field3Required field to given value.

### HasField3Required

`func (o *PortalTemplate) HasField3Required() bool`

HasField3Required returns a boolean if a field has been set.

### GetField4

`func (o *PortalTemplate) GetField4() bool`

GetField4 returns the Field4 field if non-nil, zero value otherwise.

### GetField4Ok

`func (o *PortalTemplate) GetField4Ok() (*bool, bool)`

GetField4Ok returns a tuple with the Field4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField4

`func (o *PortalTemplate) SetField4(v bool)`

SetField4 sets Field4 field to given value.

### HasField4

`func (o *PortalTemplate) HasField4() bool`

HasField4 returns a boolean if a field has been set.

### GetField4Error

`func (o *PortalTemplate) GetField4Error() string`

GetField4Error returns the Field4Error field if non-nil, zero value otherwise.

### GetField4ErrorOk

`func (o *PortalTemplate) GetField4ErrorOk() (*string, bool)`

GetField4ErrorOk returns a tuple with the Field4Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField4Error

`func (o *PortalTemplate) SetField4Error(v string)`

SetField4Error sets Field4Error field to given value.

### HasField4Error

`func (o *PortalTemplate) HasField4Error() bool`

HasField4Error returns a boolean if a field has been set.

### GetField4Label

`func (o *PortalTemplate) GetField4Label() string`

GetField4Label returns the Field4Label field if non-nil, zero value otherwise.

### GetField4LabelOk

`func (o *PortalTemplate) GetField4LabelOk() (*string, bool)`

GetField4LabelOk returns a tuple with the Field4Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField4Label

`func (o *PortalTemplate) SetField4Label(v string)`

SetField4Label sets Field4Label field to given value.

### HasField4Label

`func (o *PortalTemplate) HasField4Label() bool`

HasField4Label returns a boolean if a field has been set.

### GetField4Required

`func (o *PortalTemplate) GetField4Required() bool`

GetField4Required returns the Field4Required field if non-nil, zero value otherwise.

### GetField4RequiredOk

`func (o *PortalTemplate) GetField4RequiredOk() (*bool, bool)`

GetField4RequiredOk returns a tuple with the Field4Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField4Required

`func (o *PortalTemplate) SetField4Required(v bool)`

SetField4Required sets Field4Required field to given value.

### HasField4Required

`func (o *PortalTemplate) HasField4Required() bool`

HasField4Required returns a boolean if a field has been set.

### GetForSite

`func (o *PortalTemplate) GetForSite() bool`

GetForSite returns the ForSite field if non-nil, zero value otherwise.

### GetForSiteOk

`func (o *PortalTemplate) GetForSiteOk() (*bool, bool)`

GetForSiteOk returns a tuple with the ForSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSite

`func (o *PortalTemplate) SetForSite(v bool)`

SetForSite sets ForSite field to given value.

### HasForSite

`func (o *PortalTemplate) HasForSite() bool`

HasForSite returns a boolean if a field has been set.

### GetId

`func (o *PortalTemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortalTemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortalTemplate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PortalTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *PortalTemplate) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PortalTemplate) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PortalTemplate) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PortalTemplate) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetModifiedTime

`func (o *PortalTemplate) GetModifiedTime() float32`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *PortalTemplate) GetModifiedTimeOk() (*float32, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *PortalTemplate) SetModifiedTime(v float32)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *PortalTemplate) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetName

`func (o *PortalTemplate) GetName() bool`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PortalTemplate) GetNameOk() (*bool, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PortalTemplate) SetName(v bool)`

SetName sets Name field to given value.

### HasName

`func (o *PortalTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNameError

`func (o *PortalTemplate) GetNameError() string`

GetNameError returns the NameError field if non-nil, zero value otherwise.

### GetNameErrorOk

`func (o *PortalTemplate) GetNameErrorOk() (*string, bool)`

GetNameErrorOk returns a tuple with the NameError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameError

`func (o *PortalTemplate) SetNameError(v string)`

SetNameError sets NameError field to given value.

### HasNameError

`func (o *PortalTemplate) HasNameError() bool`

HasNameError returns a boolean if a field has been set.

### GetNameLabel

`func (o *PortalTemplate) GetNameLabel() string`

GetNameLabel returns the NameLabel field if non-nil, zero value otherwise.

### GetNameLabelOk

`func (o *PortalTemplate) GetNameLabelOk() (*string, bool)`

GetNameLabelOk returns a tuple with the NameLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameLabel

`func (o *PortalTemplate) SetNameLabel(v string)`

SetNameLabel sets NameLabel field to given value.

### HasNameLabel

`func (o *PortalTemplate) HasNameLabel() bool`

HasNameLabel returns a boolean if a field has been set.

### GetOptout

`func (o *PortalTemplate) GetOptout() bool`

GetOptout returns the Optout field if non-nil, zero value otherwise.

### GetOptoutOk

`func (o *PortalTemplate) GetOptoutOk() (*bool, bool)`

GetOptoutOk returns a tuple with the Optout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptout

`func (o *PortalTemplate) SetOptout(v bool)`

SetOptout sets Optout field to given value.

### HasOptout

`func (o *PortalTemplate) HasOptout() bool`

HasOptout returns a boolean if a field has been set.

### GetOptoutLabel

`func (o *PortalTemplate) GetOptoutLabel() string`

GetOptoutLabel returns the OptoutLabel field if non-nil, zero value otherwise.

### GetOptoutLabelOk

`func (o *PortalTemplate) GetOptoutLabelOk() (*string, bool)`

GetOptoutLabelOk returns a tuple with the OptoutLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptoutLabel

`func (o *PortalTemplate) SetOptoutLabel(v string)`

SetOptoutLabel sets OptoutLabel field to given value.

### HasOptoutLabel

`func (o *PortalTemplate) HasOptoutLabel() bool`

HasOptoutLabel returns a boolean if a field has been set.

### GetOrgId

`func (o *PortalTemplate) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *PortalTemplate) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *PortalTemplate) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *PortalTemplate) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPageTitle

`func (o *PortalTemplate) GetPageTitle() string`

GetPageTitle returns the PageTitle field if non-nil, zero value otherwise.

### GetPageTitleOk

`func (o *PortalTemplate) GetPageTitleOk() (*string, bool)`

GetPageTitleOk returns a tuple with the PageTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageTitle

`func (o *PortalTemplate) SetPageTitle(v string)`

SetPageTitle sets PageTitle field to given value.


### GetPassphraseCancel

`func (o *PortalTemplate) GetPassphraseCancel() string`

GetPassphraseCancel returns the PassphraseCancel field if non-nil, zero value otherwise.

### GetPassphraseCancelOk

`func (o *PortalTemplate) GetPassphraseCancelOk() (*string, bool)`

GetPassphraseCancelOk returns a tuple with the PassphraseCancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphraseCancel

`func (o *PortalTemplate) SetPassphraseCancel(v string)`

SetPassphraseCancel sets PassphraseCancel field to given value.

### HasPassphraseCancel

`func (o *PortalTemplate) HasPassphraseCancel() bool`

HasPassphraseCancel returns a boolean if a field has been set.

### GetPassphraseError

`func (o *PortalTemplate) GetPassphraseError() string`

GetPassphraseError returns the PassphraseError field if non-nil, zero value otherwise.

### GetPassphraseErrorOk

`func (o *PortalTemplate) GetPassphraseErrorOk() (*string, bool)`

GetPassphraseErrorOk returns a tuple with the PassphraseError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphraseError

`func (o *PortalTemplate) SetPassphraseError(v string)`

SetPassphraseError sets PassphraseError field to given value.

### HasPassphraseError

`func (o *PortalTemplate) HasPassphraseError() bool`

HasPassphraseError returns a boolean if a field has been set.

### GetPassphraseLabel

`func (o *PortalTemplate) GetPassphraseLabel() string`

GetPassphraseLabel returns the PassphraseLabel field if non-nil, zero value otherwise.

### GetPassphraseLabelOk

`func (o *PortalTemplate) GetPassphraseLabelOk() (*string, bool)`

GetPassphraseLabelOk returns a tuple with the PassphraseLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphraseLabel

`func (o *PortalTemplate) SetPassphraseLabel(v string)`

SetPassphraseLabel sets PassphraseLabel field to given value.

### HasPassphraseLabel

`func (o *PortalTemplate) HasPassphraseLabel() bool`

HasPassphraseLabel returns a boolean if a field has been set.

### GetPassphraseMessage

`func (o *PortalTemplate) GetPassphraseMessage() string`

GetPassphraseMessage returns the PassphraseMessage field if non-nil, zero value otherwise.

### GetPassphraseMessageOk

`func (o *PortalTemplate) GetPassphraseMessageOk() (*string, bool)`

GetPassphraseMessageOk returns a tuple with the PassphraseMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphraseMessage

`func (o *PortalTemplate) SetPassphraseMessage(v string)`

SetPassphraseMessage sets PassphraseMessage field to given value.

### HasPassphraseMessage

`func (o *PortalTemplate) HasPassphraseMessage() bool`

HasPassphraseMessage returns a boolean if a field has been set.

### GetPassphraseSubmit

`func (o *PortalTemplate) GetPassphraseSubmit() string`

GetPassphraseSubmit returns the PassphraseSubmit field if non-nil, zero value otherwise.

### GetPassphraseSubmitOk

`func (o *PortalTemplate) GetPassphraseSubmitOk() (*string, bool)`

GetPassphraseSubmitOk returns a tuple with the PassphraseSubmit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphraseSubmit

`func (o *PortalTemplate) SetPassphraseSubmit(v string)`

SetPassphraseSubmit sets PassphraseSubmit field to given value.

### HasPassphraseSubmit

`func (o *PortalTemplate) HasPassphraseSubmit() bool`

HasPassphraseSubmit returns a boolean if a field has been set.

### GetPassphraseTitle

`func (o *PortalTemplate) GetPassphraseTitle() string`

GetPassphraseTitle returns the PassphraseTitle field if non-nil, zero value otherwise.

### GetPassphraseTitleOk

`func (o *PortalTemplate) GetPassphraseTitleOk() (*string, bool)`

GetPassphraseTitleOk returns a tuple with the PassphraseTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphraseTitle

`func (o *PortalTemplate) SetPassphraseTitle(v string)`

SetPassphraseTitle sets PassphraseTitle field to given value.

### HasPassphraseTitle

`func (o *PortalTemplate) HasPassphraseTitle() bool`

HasPassphraseTitle returns a boolean if a field has been set.

### GetPoweredBy

`func (o *PortalTemplate) GetPoweredBy() bool`

GetPoweredBy returns the PoweredBy field if non-nil, zero value otherwise.

### GetPoweredByOk

`func (o *PortalTemplate) GetPoweredByOk() (*bool, bool)`

GetPoweredByOk returns a tuple with the PoweredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoweredBy

`func (o *PortalTemplate) SetPoweredBy(v bool)`

SetPoweredBy sets PoweredBy field to given value.

### HasPoweredBy

`func (o *PortalTemplate) HasPoweredBy() bool`

HasPoweredBy returns a boolean if a field has been set.

### GetPrivacyPolicy

`func (o *PortalTemplate) GetPrivacyPolicy() bool`

GetPrivacyPolicy returns the PrivacyPolicy field if non-nil, zero value otherwise.

### GetPrivacyPolicyOk

`func (o *PortalTemplate) GetPrivacyPolicyOk() (*bool, bool)`

GetPrivacyPolicyOk returns a tuple with the PrivacyPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyPolicy

`func (o *PortalTemplate) SetPrivacyPolicy(v bool)`

SetPrivacyPolicy sets PrivacyPolicy field to given value.

### HasPrivacyPolicy

`func (o *PortalTemplate) HasPrivacyPolicy() bool`

HasPrivacyPolicy returns a boolean if a field has been set.

### GetPrivacyPolicyAcceptLabel

`func (o *PortalTemplate) GetPrivacyPolicyAcceptLabel() string`

GetPrivacyPolicyAcceptLabel returns the PrivacyPolicyAcceptLabel field if non-nil, zero value otherwise.

### GetPrivacyPolicyAcceptLabelOk

`func (o *PortalTemplate) GetPrivacyPolicyAcceptLabelOk() (*string, bool)`

GetPrivacyPolicyAcceptLabelOk returns a tuple with the PrivacyPolicyAcceptLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyPolicyAcceptLabel

`func (o *PortalTemplate) SetPrivacyPolicyAcceptLabel(v string)`

SetPrivacyPolicyAcceptLabel sets PrivacyPolicyAcceptLabel field to given value.

### HasPrivacyPolicyAcceptLabel

`func (o *PortalTemplate) HasPrivacyPolicyAcceptLabel() bool`

HasPrivacyPolicyAcceptLabel returns a boolean if a field has been set.

### GetPrivacyPolicyError

`func (o *PortalTemplate) GetPrivacyPolicyError() string`

GetPrivacyPolicyError returns the PrivacyPolicyError field if non-nil, zero value otherwise.

### GetPrivacyPolicyErrorOk

`func (o *PortalTemplate) GetPrivacyPolicyErrorOk() (*string, bool)`

GetPrivacyPolicyErrorOk returns a tuple with the PrivacyPolicyError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyPolicyError

`func (o *PortalTemplate) SetPrivacyPolicyError(v string)`

SetPrivacyPolicyError sets PrivacyPolicyError field to given value.

### HasPrivacyPolicyError

`func (o *PortalTemplate) HasPrivacyPolicyError() bool`

HasPrivacyPolicyError returns a boolean if a field has been set.

### GetPrivacyPolicyLink

`func (o *PortalTemplate) GetPrivacyPolicyLink() string`

GetPrivacyPolicyLink returns the PrivacyPolicyLink field if non-nil, zero value otherwise.

### GetPrivacyPolicyLinkOk

`func (o *PortalTemplate) GetPrivacyPolicyLinkOk() (*string, bool)`

GetPrivacyPolicyLinkOk returns a tuple with the PrivacyPolicyLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyPolicyLink

`func (o *PortalTemplate) SetPrivacyPolicyLink(v string)`

SetPrivacyPolicyLink sets PrivacyPolicyLink field to given value.

### HasPrivacyPolicyLink

`func (o *PortalTemplate) HasPrivacyPolicyLink() bool`

HasPrivacyPolicyLink returns a boolean if a field has been set.

### GetPrivacyPolicyText

`func (o *PortalTemplate) GetPrivacyPolicyText() string`

GetPrivacyPolicyText returns the PrivacyPolicyText field if non-nil, zero value otherwise.

### GetPrivacyPolicyTextOk

`func (o *PortalTemplate) GetPrivacyPolicyTextOk() (*string, bool)`

GetPrivacyPolicyTextOk returns a tuple with the PrivacyPolicyText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyPolicyText

`func (o *PortalTemplate) SetPrivacyPolicyText(v string)`

SetPrivacyPolicyText sets PrivacyPolicyText field to given value.

### HasPrivacyPolicyText

`func (o *PortalTemplate) HasPrivacyPolicyText() bool`

HasPrivacyPolicyText returns a boolean if a field has been set.

### GetRequiredFieldLabel

`func (o *PortalTemplate) GetRequiredFieldLabel() string`

GetRequiredFieldLabel returns the RequiredFieldLabel field if non-nil, zero value otherwise.

### GetRequiredFieldLabelOk

`func (o *PortalTemplate) GetRequiredFieldLabelOk() (*string, bool)`

GetRequiredFieldLabelOk returns a tuple with the RequiredFieldLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredFieldLabel

`func (o *PortalTemplate) SetRequiredFieldLabel(v string)`

SetRequiredFieldLabel sets RequiredFieldLabel field to given value.

### HasRequiredFieldLabel

`func (o *PortalTemplate) HasRequiredFieldLabel() bool`

HasRequiredFieldLabel returns a boolean if a field has been set.

### GetSignInLabel

`func (o *PortalTemplate) GetSignInLabel() string`

GetSignInLabel returns the SignInLabel field if non-nil, zero value otherwise.

### GetSignInLabelOk

`func (o *PortalTemplate) GetSignInLabelOk() (*string, bool)`

GetSignInLabelOk returns a tuple with the SignInLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignInLabel

`func (o *PortalTemplate) SetSignInLabel(v string)`

SetSignInLabel sets SignInLabel field to given value.

### HasSignInLabel

`func (o *PortalTemplate) HasSignInLabel() bool`

HasSignInLabel returns a boolean if a field has been set.

### GetSiteId

`func (o *PortalTemplate) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *PortalTemplate) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *PortalTemplate) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *PortalTemplate) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSmsCarrierDefault

`func (o *PortalTemplate) GetSmsCarrierDefault() string`

GetSmsCarrierDefault returns the SmsCarrierDefault field if non-nil, zero value otherwise.

### GetSmsCarrierDefaultOk

`func (o *PortalTemplate) GetSmsCarrierDefaultOk() (*string, bool)`

GetSmsCarrierDefaultOk returns a tuple with the SmsCarrierDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsCarrierDefault

`func (o *PortalTemplate) SetSmsCarrierDefault(v string)`

SetSmsCarrierDefault sets SmsCarrierDefault field to given value.

### HasSmsCarrierDefault

`func (o *PortalTemplate) HasSmsCarrierDefault() bool`

HasSmsCarrierDefault returns a boolean if a field has been set.

### GetSmsCarrierError

`func (o *PortalTemplate) GetSmsCarrierError() string`

GetSmsCarrierError returns the SmsCarrierError field if non-nil, zero value otherwise.

### GetSmsCarrierErrorOk

`func (o *PortalTemplate) GetSmsCarrierErrorOk() (*string, bool)`

GetSmsCarrierErrorOk returns a tuple with the SmsCarrierError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsCarrierError

`func (o *PortalTemplate) SetSmsCarrierError(v string)`

SetSmsCarrierError sets SmsCarrierError field to given value.

### HasSmsCarrierError

`func (o *PortalTemplate) HasSmsCarrierError() bool`

HasSmsCarrierError returns a boolean if a field has been set.

### GetSmsCarrierFieldLabel

`func (o *PortalTemplate) GetSmsCarrierFieldLabel() string`

GetSmsCarrierFieldLabel returns the SmsCarrierFieldLabel field if non-nil, zero value otherwise.

### GetSmsCarrierFieldLabelOk

`func (o *PortalTemplate) GetSmsCarrierFieldLabelOk() (*string, bool)`

GetSmsCarrierFieldLabelOk returns a tuple with the SmsCarrierFieldLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsCarrierFieldLabel

`func (o *PortalTemplate) SetSmsCarrierFieldLabel(v string)`

SetSmsCarrierFieldLabel sets SmsCarrierFieldLabel field to given value.

### HasSmsCarrierFieldLabel

`func (o *PortalTemplate) HasSmsCarrierFieldLabel() bool`

HasSmsCarrierFieldLabel returns a boolean if a field has been set.

### GetSmsCodeCancel

`func (o *PortalTemplate) GetSmsCodeCancel() string`

GetSmsCodeCancel returns the SmsCodeCancel field if non-nil, zero value otherwise.

### GetSmsCodeCancelOk

`func (o *PortalTemplate) GetSmsCodeCancelOk() (*string, bool)`

GetSmsCodeCancelOk returns a tuple with the SmsCodeCancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsCodeCancel

`func (o *PortalTemplate) SetSmsCodeCancel(v string)`

SetSmsCodeCancel sets SmsCodeCancel field to given value.

### HasSmsCodeCancel

`func (o *PortalTemplate) HasSmsCodeCancel() bool`

HasSmsCodeCancel returns a boolean if a field has been set.

### GetSmsCodeError

`func (o *PortalTemplate) GetSmsCodeError() string`

GetSmsCodeError returns the SmsCodeError field if non-nil, zero value otherwise.

### GetSmsCodeErrorOk

`func (o *PortalTemplate) GetSmsCodeErrorOk() (*string, bool)`

GetSmsCodeErrorOk returns a tuple with the SmsCodeError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsCodeError

`func (o *PortalTemplate) SetSmsCodeError(v string)`

SetSmsCodeError sets SmsCodeError field to given value.

### HasSmsCodeError

`func (o *PortalTemplate) HasSmsCodeError() bool`

HasSmsCodeError returns a boolean if a field has been set.

### GetSmsCodeFieldLabel

`func (o *PortalTemplate) GetSmsCodeFieldLabel() string`

GetSmsCodeFieldLabel returns the SmsCodeFieldLabel field if non-nil, zero value otherwise.

### GetSmsCodeFieldLabelOk

`func (o *PortalTemplate) GetSmsCodeFieldLabelOk() (*string, bool)`

GetSmsCodeFieldLabelOk returns a tuple with the SmsCodeFieldLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsCodeFieldLabel

`func (o *PortalTemplate) SetSmsCodeFieldLabel(v string)`

SetSmsCodeFieldLabel sets SmsCodeFieldLabel field to given value.

### HasSmsCodeFieldLabel

`func (o *PortalTemplate) HasSmsCodeFieldLabel() bool`

HasSmsCodeFieldLabel returns a boolean if a field has been set.

### GetSmsCodeMessage

`func (o *PortalTemplate) GetSmsCodeMessage() string`

GetSmsCodeMessage returns the SmsCodeMessage field if non-nil, zero value otherwise.

### GetSmsCodeMessageOk

`func (o *PortalTemplate) GetSmsCodeMessageOk() (*string, bool)`

GetSmsCodeMessageOk returns a tuple with the SmsCodeMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsCodeMessage

`func (o *PortalTemplate) SetSmsCodeMessage(v string)`

SetSmsCodeMessage sets SmsCodeMessage field to given value.

### HasSmsCodeMessage

`func (o *PortalTemplate) HasSmsCodeMessage() bool`

HasSmsCodeMessage returns a boolean if a field has been set.

### GetSmsCodeSubmit

`func (o *PortalTemplate) GetSmsCodeSubmit() string`

GetSmsCodeSubmit returns the SmsCodeSubmit field if non-nil, zero value otherwise.

### GetSmsCodeSubmitOk

`func (o *PortalTemplate) GetSmsCodeSubmitOk() (*string, bool)`

GetSmsCodeSubmitOk returns a tuple with the SmsCodeSubmit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsCodeSubmit

`func (o *PortalTemplate) SetSmsCodeSubmit(v string)`

SetSmsCodeSubmit sets SmsCodeSubmit field to given value.

### HasSmsCodeSubmit

`func (o *PortalTemplate) HasSmsCodeSubmit() bool`

HasSmsCodeSubmit returns a boolean if a field has been set.

### GetSmsCodeTitle

`func (o *PortalTemplate) GetSmsCodeTitle() string`

GetSmsCodeTitle returns the SmsCodeTitle field if non-nil, zero value otherwise.

### GetSmsCodeTitleOk

`func (o *PortalTemplate) GetSmsCodeTitleOk() (*string, bool)`

GetSmsCodeTitleOk returns a tuple with the SmsCodeTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsCodeTitle

`func (o *PortalTemplate) SetSmsCodeTitle(v string)`

SetSmsCodeTitle sets SmsCodeTitle field to given value.

### HasSmsCodeTitle

`func (o *PortalTemplate) HasSmsCodeTitle() bool`

HasSmsCodeTitle returns a boolean if a field has been set.

### GetSmsCountryFieldLabel

`func (o *PortalTemplate) GetSmsCountryFieldLabel() string`

GetSmsCountryFieldLabel returns the SmsCountryFieldLabel field if non-nil, zero value otherwise.

### GetSmsCountryFieldLabelOk

`func (o *PortalTemplate) GetSmsCountryFieldLabelOk() (*string, bool)`

GetSmsCountryFieldLabelOk returns a tuple with the SmsCountryFieldLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsCountryFieldLabel

`func (o *PortalTemplate) SetSmsCountryFieldLabel(v string)`

SetSmsCountryFieldLabel sets SmsCountryFieldLabel field to given value.

### HasSmsCountryFieldLabel

`func (o *PortalTemplate) HasSmsCountryFieldLabel() bool`

HasSmsCountryFieldLabel returns a boolean if a field has been set.

### GetSmsCountryFormat

`func (o *PortalTemplate) GetSmsCountryFormat() string`

GetSmsCountryFormat returns the SmsCountryFormat field if non-nil, zero value otherwise.

### GetSmsCountryFormatOk

`func (o *PortalTemplate) GetSmsCountryFormatOk() (*string, bool)`

GetSmsCountryFormatOk returns a tuple with the SmsCountryFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsCountryFormat

`func (o *PortalTemplate) SetSmsCountryFormat(v string)`

SetSmsCountryFormat sets SmsCountryFormat field to given value.

### HasSmsCountryFormat

`func (o *PortalTemplate) HasSmsCountryFormat() bool`

HasSmsCountryFormat returns a boolean if a field has been set.

### GetSmsHaveAccessCode

`func (o *PortalTemplate) GetSmsHaveAccessCode() string`

GetSmsHaveAccessCode returns the SmsHaveAccessCode field if non-nil, zero value otherwise.

### GetSmsHaveAccessCodeOk

`func (o *PortalTemplate) GetSmsHaveAccessCodeOk() (*string, bool)`

GetSmsHaveAccessCodeOk returns a tuple with the SmsHaveAccessCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsHaveAccessCode

`func (o *PortalTemplate) SetSmsHaveAccessCode(v string)`

SetSmsHaveAccessCode sets SmsHaveAccessCode field to given value.

### HasSmsHaveAccessCode

`func (o *PortalTemplate) HasSmsHaveAccessCode() bool`

HasSmsHaveAccessCode returns a boolean if a field has been set.

### GetSmsMessageFormat

`func (o *PortalTemplate) GetSmsMessageFormat() string`

GetSmsMessageFormat returns the SmsMessageFormat field if non-nil, zero value otherwise.

### GetSmsMessageFormatOk

`func (o *PortalTemplate) GetSmsMessageFormatOk() (*string, bool)`

GetSmsMessageFormatOk returns a tuple with the SmsMessageFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsMessageFormat

`func (o *PortalTemplate) SetSmsMessageFormat(v string)`

SetSmsMessageFormat sets SmsMessageFormat field to given value.

### HasSmsMessageFormat

`func (o *PortalTemplate) HasSmsMessageFormat() bool`

HasSmsMessageFormat returns a boolean if a field has been set.

### GetSmsNumberCancel

`func (o *PortalTemplate) GetSmsNumberCancel() string`

GetSmsNumberCancel returns the SmsNumberCancel field if non-nil, zero value otherwise.

### GetSmsNumberCancelOk

`func (o *PortalTemplate) GetSmsNumberCancelOk() (*string, bool)`

GetSmsNumberCancelOk returns a tuple with the SmsNumberCancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsNumberCancel

`func (o *PortalTemplate) SetSmsNumberCancel(v string)`

SetSmsNumberCancel sets SmsNumberCancel field to given value.

### HasSmsNumberCancel

`func (o *PortalTemplate) HasSmsNumberCancel() bool`

HasSmsNumberCancel returns a boolean if a field has been set.

### GetSmsNumberError

`func (o *PortalTemplate) GetSmsNumberError() string`

GetSmsNumberError returns the SmsNumberError field if non-nil, zero value otherwise.

### GetSmsNumberErrorOk

`func (o *PortalTemplate) GetSmsNumberErrorOk() (*string, bool)`

GetSmsNumberErrorOk returns a tuple with the SmsNumberError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsNumberError

`func (o *PortalTemplate) SetSmsNumberError(v string)`

SetSmsNumberError sets SmsNumberError field to given value.

### HasSmsNumberError

`func (o *PortalTemplate) HasSmsNumberError() bool`

HasSmsNumberError returns a boolean if a field has been set.

### GetSmsNumberFieldLabel

`func (o *PortalTemplate) GetSmsNumberFieldLabel() string`

GetSmsNumberFieldLabel returns the SmsNumberFieldLabel field if non-nil, zero value otherwise.

### GetSmsNumberFieldLabelOk

`func (o *PortalTemplate) GetSmsNumberFieldLabelOk() (*string, bool)`

GetSmsNumberFieldLabelOk returns a tuple with the SmsNumberFieldLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsNumberFieldLabel

`func (o *PortalTemplate) SetSmsNumberFieldLabel(v string)`

SetSmsNumberFieldLabel sets SmsNumberFieldLabel field to given value.

### HasSmsNumberFieldLabel

`func (o *PortalTemplate) HasSmsNumberFieldLabel() bool`

HasSmsNumberFieldLabel returns a boolean if a field has been set.

### GetSmsNumberFormat

`func (o *PortalTemplate) GetSmsNumberFormat() string`

GetSmsNumberFormat returns the SmsNumberFormat field if non-nil, zero value otherwise.

### GetSmsNumberFormatOk

`func (o *PortalTemplate) GetSmsNumberFormatOk() (*string, bool)`

GetSmsNumberFormatOk returns a tuple with the SmsNumberFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsNumberFormat

`func (o *PortalTemplate) SetSmsNumberFormat(v string)`

SetSmsNumberFormat sets SmsNumberFormat field to given value.

### HasSmsNumberFormat

`func (o *PortalTemplate) HasSmsNumberFormat() bool`

HasSmsNumberFormat returns a boolean if a field has been set.

### GetSmsNumberMessage

`func (o *PortalTemplate) GetSmsNumberMessage() string`

GetSmsNumberMessage returns the SmsNumberMessage field if non-nil, zero value otherwise.

### GetSmsNumberMessageOk

`func (o *PortalTemplate) GetSmsNumberMessageOk() (*string, bool)`

GetSmsNumberMessageOk returns a tuple with the SmsNumberMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsNumberMessage

`func (o *PortalTemplate) SetSmsNumberMessage(v string)`

SetSmsNumberMessage sets SmsNumberMessage field to given value.

### HasSmsNumberMessage

`func (o *PortalTemplate) HasSmsNumberMessage() bool`

HasSmsNumberMessage returns a boolean if a field has been set.

### GetSmsNumberSubmit

`func (o *PortalTemplate) GetSmsNumberSubmit() string`

GetSmsNumberSubmit returns the SmsNumberSubmit field if non-nil, zero value otherwise.

### GetSmsNumberSubmitOk

`func (o *PortalTemplate) GetSmsNumberSubmitOk() (*string, bool)`

GetSmsNumberSubmitOk returns a tuple with the SmsNumberSubmit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsNumberSubmit

`func (o *PortalTemplate) SetSmsNumberSubmit(v string)`

SetSmsNumberSubmit sets SmsNumberSubmit field to given value.

### HasSmsNumberSubmit

`func (o *PortalTemplate) HasSmsNumberSubmit() bool`

HasSmsNumberSubmit returns a boolean if a field has been set.

### GetSmsNumberTitle

`func (o *PortalTemplate) GetSmsNumberTitle() string`

GetSmsNumberTitle returns the SmsNumberTitle field if non-nil, zero value otherwise.

### GetSmsNumberTitleOk

`func (o *PortalTemplate) GetSmsNumberTitleOk() (*string, bool)`

GetSmsNumberTitleOk returns a tuple with the SmsNumberTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsNumberTitle

`func (o *PortalTemplate) SetSmsNumberTitle(v string)`

SetSmsNumberTitle sets SmsNumberTitle field to given value.

### HasSmsNumberTitle

`func (o *PortalTemplate) HasSmsNumberTitle() bool`

HasSmsNumberTitle returns a boolean if a field has been set.

### GetSmsUsernameFormat

`func (o *PortalTemplate) GetSmsUsernameFormat() string`

GetSmsUsernameFormat returns the SmsUsernameFormat field if non-nil, zero value otherwise.

### GetSmsUsernameFormatOk

`func (o *PortalTemplate) GetSmsUsernameFormatOk() (*string, bool)`

GetSmsUsernameFormatOk returns a tuple with the SmsUsernameFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsUsernameFormat

`func (o *PortalTemplate) SetSmsUsernameFormat(v string)`

SetSmsUsernameFormat sets SmsUsernameFormat field to given value.

### HasSmsUsernameFormat

`func (o *PortalTemplate) HasSmsUsernameFormat() bool`

HasSmsUsernameFormat returns a boolean if a field has been set.

### GetSmsValidityDuration

`func (o *PortalTemplate) GetSmsValidityDuration() float32`

GetSmsValidityDuration returns the SmsValidityDuration field if non-nil, zero value otherwise.

### GetSmsValidityDurationOk

`func (o *PortalTemplate) GetSmsValidityDurationOk() (*float32, bool)`

GetSmsValidityDurationOk returns a tuple with the SmsValidityDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsValidityDuration

`func (o *PortalTemplate) SetSmsValidityDuration(v float32)`

SetSmsValidityDuration sets SmsValidityDuration field to given value.

### HasSmsValidityDuration

`func (o *PortalTemplate) HasSmsValidityDuration() bool`

HasSmsValidityDuration returns a boolean if a field has been set.

### GetSponsorAutoApproved

`func (o *PortalTemplate) GetSponsorAutoApproved() string`

GetSponsorAutoApproved returns the SponsorAutoApproved field if non-nil, zero value otherwise.

### GetSponsorAutoApprovedOk

`func (o *PortalTemplate) GetSponsorAutoApprovedOk() (*string, bool)`

GetSponsorAutoApprovedOk returns a tuple with the SponsorAutoApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorAutoApproved

`func (o *PortalTemplate) SetSponsorAutoApproved(v string)`

SetSponsorAutoApproved sets SponsorAutoApproved field to given value.

### HasSponsorAutoApproved

`func (o *PortalTemplate) HasSponsorAutoApproved() bool`

HasSponsorAutoApproved returns a boolean if a field has been set.

### GetSponsorAutoApprovedNote

`func (o *PortalTemplate) GetSponsorAutoApprovedNote() string`

GetSponsorAutoApprovedNote returns the SponsorAutoApprovedNote field if non-nil, zero value otherwise.

### GetSponsorAutoApprovedNoteOk

`func (o *PortalTemplate) GetSponsorAutoApprovedNoteOk() (*string, bool)`

GetSponsorAutoApprovedNoteOk returns a tuple with the SponsorAutoApprovedNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorAutoApprovedNote

`func (o *PortalTemplate) SetSponsorAutoApprovedNote(v string)`

SetSponsorAutoApprovedNote sets SponsorAutoApprovedNote field to given value.

### HasSponsorAutoApprovedNote

`func (o *PortalTemplate) HasSponsorAutoApprovedNote() bool`

HasSponsorAutoApprovedNote returns a boolean if a field has been set.

### GetSponsorBackLink

`func (o *PortalTemplate) GetSponsorBackLink() string`

GetSponsorBackLink returns the SponsorBackLink field if non-nil, zero value otherwise.

### GetSponsorBackLinkOk

`func (o *PortalTemplate) GetSponsorBackLinkOk() (*string, bool)`

GetSponsorBackLinkOk returns a tuple with the SponsorBackLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorBackLink

`func (o *PortalTemplate) SetSponsorBackLink(v string)`

SetSponsorBackLink sets SponsorBackLink field to given value.

### HasSponsorBackLink

`func (o *PortalTemplate) HasSponsorBackLink() bool`

HasSponsorBackLink returns a boolean if a field has been set.

### GetSponsorCancel

`func (o *PortalTemplate) GetSponsorCancel() string`

GetSponsorCancel returns the SponsorCancel field if non-nil, zero value otherwise.

### GetSponsorCancelOk

`func (o *PortalTemplate) GetSponsorCancelOk() (*string, bool)`

GetSponsorCancelOk returns a tuple with the SponsorCancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorCancel

`func (o *PortalTemplate) SetSponsorCancel(v string)`

SetSponsorCancel sets SponsorCancel field to given value.

### HasSponsorCancel

`func (o *PortalTemplate) HasSponsorCancel() bool`

HasSponsorCancel returns a boolean if a field has been set.

### GetSponsorEmail

`func (o *PortalTemplate) GetSponsorEmail() string`

GetSponsorEmail returns the SponsorEmail field if non-nil, zero value otherwise.

### GetSponsorEmailOk

`func (o *PortalTemplate) GetSponsorEmailOk() (*string, bool)`

GetSponsorEmailOk returns a tuple with the SponsorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorEmail

`func (o *PortalTemplate) SetSponsorEmail(v string)`

SetSponsorEmail sets SponsorEmail field to given value.

### HasSponsorEmail

`func (o *PortalTemplate) HasSponsorEmail() bool`

HasSponsorEmail returns a boolean if a field has been set.

### GetSponsorEmailError

`func (o *PortalTemplate) GetSponsorEmailError() string`

GetSponsorEmailError returns the SponsorEmailError field if non-nil, zero value otherwise.

### GetSponsorEmailErrorOk

`func (o *PortalTemplate) GetSponsorEmailErrorOk() (*string, bool)`

GetSponsorEmailErrorOk returns a tuple with the SponsorEmailError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorEmailError

`func (o *PortalTemplate) SetSponsorEmailError(v string)`

SetSponsorEmailError sets SponsorEmailError field to given value.

### HasSponsorEmailError

`func (o *PortalTemplate) HasSponsorEmailError() bool`

HasSponsorEmailError returns a boolean if a field has been set.

### GetSponsorEmailTemplate

`func (o *PortalTemplate) GetSponsorEmailTemplate() string`

GetSponsorEmailTemplate returns the SponsorEmailTemplate field if non-nil, zero value otherwise.

### GetSponsorEmailTemplateOk

`func (o *PortalTemplate) GetSponsorEmailTemplateOk() (*string, bool)`

GetSponsorEmailTemplateOk returns a tuple with the SponsorEmailTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorEmailTemplate

`func (o *PortalTemplate) SetSponsorEmailTemplate(v string)`

SetSponsorEmailTemplate sets SponsorEmailTemplate field to given value.

### HasSponsorEmailTemplate

`func (o *PortalTemplate) HasSponsorEmailTemplate() bool`

HasSponsorEmailTemplate returns a boolean if a field has been set.

### GetSponsorInfoApproved

`func (o *PortalTemplate) GetSponsorInfoApproved() string`

GetSponsorInfoApproved returns the SponsorInfoApproved field if non-nil, zero value otherwise.

### GetSponsorInfoApprovedOk

`func (o *PortalTemplate) GetSponsorInfoApprovedOk() (*string, bool)`

GetSponsorInfoApprovedOk returns a tuple with the SponsorInfoApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorInfoApproved

`func (o *PortalTemplate) SetSponsorInfoApproved(v string)`

SetSponsorInfoApproved sets SponsorInfoApproved field to given value.

### HasSponsorInfoApproved

`func (o *PortalTemplate) HasSponsorInfoApproved() bool`

HasSponsorInfoApproved returns a boolean if a field has been set.

### GetSponsorInfoDenied

`func (o *PortalTemplate) GetSponsorInfoDenied() string`

GetSponsorInfoDenied returns the SponsorInfoDenied field if non-nil, zero value otherwise.

### GetSponsorInfoDeniedOk

`func (o *PortalTemplate) GetSponsorInfoDeniedOk() (*string, bool)`

GetSponsorInfoDeniedOk returns a tuple with the SponsorInfoDenied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorInfoDenied

`func (o *PortalTemplate) SetSponsorInfoDenied(v string)`

SetSponsorInfoDenied sets SponsorInfoDenied field to given value.

### HasSponsorInfoDenied

`func (o *PortalTemplate) HasSponsorInfoDenied() bool`

HasSponsorInfoDenied returns a boolean if a field has been set.

### GetSponsorInfoPending

`func (o *PortalTemplate) GetSponsorInfoPending() string`

GetSponsorInfoPending returns the SponsorInfoPending field if non-nil, zero value otherwise.

### GetSponsorInfoPendingOk

`func (o *PortalTemplate) GetSponsorInfoPendingOk() (*string, bool)`

GetSponsorInfoPendingOk returns a tuple with the SponsorInfoPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorInfoPending

`func (o *PortalTemplate) SetSponsorInfoPending(v string)`

SetSponsorInfoPending sets SponsorInfoPending field to given value.

### HasSponsorInfoPending

`func (o *PortalTemplate) HasSponsorInfoPending() bool`

HasSponsorInfoPending returns a boolean if a field has been set.

### GetSponsorName

`func (o *PortalTemplate) GetSponsorName() string`

GetSponsorName returns the SponsorName field if non-nil, zero value otherwise.

### GetSponsorNameOk

`func (o *PortalTemplate) GetSponsorNameOk() (*string, bool)`

GetSponsorNameOk returns a tuple with the SponsorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorName

`func (o *PortalTemplate) SetSponsorName(v string)`

SetSponsorName sets SponsorName field to given value.

### HasSponsorName

`func (o *PortalTemplate) HasSponsorName() bool`

HasSponsorName returns a boolean if a field has been set.

### GetSponsorNameError

`func (o *PortalTemplate) GetSponsorNameError() string`

GetSponsorNameError returns the SponsorNameError field if non-nil, zero value otherwise.

### GetSponsorNameErrorOk

`func (o *PortalTemplate) GetSponsorNameErrorOk() (*string, bool)`

GetSponsorNameErrorOk returns a tuple with the SponsorNameError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorNameError

`func (o *PortalTemplate) SetSponsorNameError(v string)`

SetSponsorNameError sets SponsorNameError field to given value.

### HasSponsorNameError

`func (o *PortalTemplate) HasSponsorNameError() bool`

HasSponsorNameError returns a boolean if a field has been set.

### GetSponsorNotePending

`func (o *PortalTemplate) GetSponsorNotePending() string`

GetSponsorNotePending returns the SponsorNotePending field if non-nil, zero value otherwise.

### GetSponsorNotePendingOk

`func (o *PortalTemplate) GetSponsorNotePendingOk() (*string, bool)`

GetSponsorNotePendingOk returns a tuple with the SponsorNotePending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorNotePending

`func (o *PortalTemplate) SetSponsorNotePending(v string)`

SetSponsorNotePending sets SponsorNotePending field to given value.

### HasSponsorNotePending

`func (o *PortalTemplate) HasSponsorNotePending() bool`

HasSponsorNotePending returns a boolean if a field has been set.

### GetSponsorStatusApproved

`func (o *PortalTemplate) GetSponsorStatusApproved() string`

GetSponsorStatusApproved returns the SponsorStatusApproved field if non-nil, zero value otherwise.

### GetSponsorStatusApprovedOk

`func (o *PortalTemplate) GetSponsorStatusApprovedOk() (*string, bool)`

GetSponsorStatusApprovedOk returns a tuple with the SponsorStatusApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorStatusApproved

`func (o *PortalTemplate) SetSponsorStatusApproved(v string)`

SetSponsorStatusApproved sets SponsorStatusApproved field to given value.

### HasSponsorStatusApproved

`func (o *PortalTemplate) HasSponsorStatusApproved() bool`

HasSponsorStatusApproved returns a boolean if a field has been set.

### GetSponsorStatusDenied

`func (o *PortalTemplate) GetSponsorStatusDenied() string`

GetSponsorStatusDenied returns the SponsorStatusDenied field if non-nil, zero value otherwise.

### GetSponsorStatusDeniedOk

`func (o *PortalTemplate) GetSponsorStatusDeniedOk() (*string, bool)`

GetSponsorStatusDeniedOk returns a tuple with the SponsorStatusDenied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorStatusDenied

`func (o *PortalTemplate) SetSponsorStatusDenied(v string)`

SetSponsorStatusDenied sets SponsorStatusDenied field to given value.

### HasSponsorStatusDenied

`func (o *PortalTemplate) HasSponsorStatusDenied() bool`

HasSponsorStatusDenied returns a boolean if a field has been set.

### GetSponsorStatusPending

`func (o *PortalTemplate) GetSponsorStatusPending() string`

GetSponsorStatusPending returns the SponsorStatusPending field if non-nil, zero value otherwise.

### GetSponsorStatusPendingOk

`func (o *PortalTemplate) GetSponsorStatusPendingOk() (*string, bool)`

GetSponsorStatusPendingOk returns a tuple with the SponsorStatusPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorStatusPending

`func (o *PortalTemplate) SetSponsorStatusPending(v string)`

SetSponsorStatusPending sets SponsorStatusPending field to given value.

### HasSponsorStatusPending

`func (o *PortalTemplate) HasSponsorStatusPending() bool`

HasSponsorStatusPending returns a boolean if a field has been set.

### GetSponsorSubmit

`func (o *PortalTemplate) GetSponsorSubmit() string`

GetSponsorSubmit returns the SponsorSubmit field if non-nil, zero value otherwise.

### GetSponsorSubmitOk

`func (o *PortalTemplate) GetSponsorSubmitOk() (*string, bool)`

GetSponsorSubmitOk returns a tuple with the SponsorSubmit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorSubmit

`func (o *PortalTemplate) SetSponsorSubmit(v string)`

SetSponsorSubmit sets SponsorSubmit field to given value.

### HasSponsorSubmit

`func (o *PortalTemplate) HasSponsorSubmit() bool`

HasSponsorSubmit returns a boolean if a field has been set.

### GetSponsorsAutoApprovedNote

`func (o *PortalTemplate) GetSponsorsAutoApprovedNote() string`

GetSponsorsAutoApprovedNote returns the SponsorsAutoApprovedNote field if non-nil, zero value otherwise.

### GetSponsorsAutoApprovedNoteOk

`func (o *PortalTemplate) GetSponsorsAutoApprovedNoteOk() (*string, bool)`

GetSponsorsAutoApprovedNoteOk returns a tuple with the SponsorsAutoApprovedNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorsAutoApprovedNote

`func (o *PortalTemplate) SetSponsorsAutoApprovedNote(v string)`

SetSponsorsAutoApprovedNote sets SponsorsAutoApprovedNote field to given value.

### HasSponsorsAutoApprovedNote

`func (o *PortalTemplate) HasSponsorsAutoApprovedNote() bool`

HasSponsorsAutoApprovedNote returns a boolean if a field has been set.

### GetSponsorsError

`func (o *PortalTemplate) GetSponsorsError() string`

GetSponsorsError returns the SponsorsError field if non-nil, zero value otherwise.

### GetSponsorsErrorOk

`func (o *PortalTemplate) GetSponsorsErrorOk() (*string, bool)`

GetSponsorsErrorOk returns a tuple with the SponsorsError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorsError

`func (o *PortalTemplate) SetSponsorsError(v string)`

SetSponsorsError sets SponsorsError field to given value.

### HasSponsorsError

`func (o *PortalTemplate) HasSponsorsError() bool`

HasSponsorsError returns a boolean if a field has been set.

### GetTos

`func (o *PortalTemplate) GetTos() bool`

GetTos returns the Tos field if non-nil, zero value otherwise.

### GetTosOk

`func (o *PortalTemplate) GetTosOk() (*bool, bool)`

GetTosOk returns a tuple with the Tos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTos

`func (o *PortalTemplate) SetTos(v bool)`

SetTos sets Tos field to given value.

### HasTos

`func (o *PortalTemplate) HasTos() bool`

HasTos returns a boolean if a field has been set.

### GetTosAcceptLabel

`func (o *PortalTemplate) GetTosAcceptLabel() string`

GetTosAcceptLabel returns the TosAcceptLabel field if non-nil, zero value otherwise.

### GetTosAcceptLabelOk

`func (o *PortalTemplate) GetTosAcceptLabelOk() (*string, bool)`

GetTosAcceptLabelOk returns a tuple with the TosAcceptLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTosAcceptLabel

`func (o *PortalTemplate) SetTosAcceptLabel(v string)`

SetTosAcceptLabel sets TosAcceptLabel field to given value.

### HasTosAcceptLabel

`func (o *PortalTemplate) HasTosAcceptLabel() bool`

HasTosAcceptLabel returns a boolean if a field has been set.

### GetTosError

`func (o *PortalTemplate) GetTosError() string`

GetTosError returns the TosError field if non-nil, zero value otherwise.

### GetTosErrorOk

`func (o *PortalTemplate) GetTosErrorOk() (*string, bool)`

GetTosErrorOk returns a tuple with the TosError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTosError

`func (o *PortalTemplate) SetTosError(v string)`

SetTosError sets TosError field to given value.

### HasTosError

`func (o *PortalTemplate) HasTosError() bool`

HasTosError returns a boolean if a field has been set.

### GetTosLink

`func (o *PortalTemplate) GetTosLink() string`

GetTosLink returns the TosLink field if non-nil, zero value otherwise.

### GetTosLinkOk

`func (o *PortalTemplate) GetTosLinkOk() (*string, bool)`

GetTosLinkOk returns a tuple with the TosLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTosLink

`func (o *PortalTemplate) SetTosLink(v string)`

SetTosLink sets TosLink field to given value.

### HasTosLink

`func (o *PortalTemplate) HasTosLink() bool`

HasTosLink returns a boolean if a field has been set.

### GetTosText

`func (o *PortalTemplate) GetTosText() string`

GetTosText returns the TosText field if non-nil, zero value otherwise.

### GetTosTextOk

`func (o *PortalTemplate) GetTosTextOk() (*string, bool)`

GetTosTextOk returns a tuple with the TosText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTosText

`func (o *PortalTemplate) SetTosText(v string)`

SetTosText sets TosText field to given value.

### HasTosText

`func (o *PortalTemplate) HasTosText() bool`

HasTosText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


