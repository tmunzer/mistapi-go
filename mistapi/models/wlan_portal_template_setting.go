package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// WlanPortalTemplateSetting represents a WlanPortalTemplateSetting struct.
// portal template wlan settings
type WlanPortalTemplateSetting struct {
	AccessCodeAlternateEmail *string `json:"accessCodeAlternateEmail,omitempty"`
	// defines alignment on portal. enum: `center`, `left`, `right`
	Alignment *PortalTemplateAlignmentEnum     `json:"alignment,omitempty"`
	Ar        *WlanPortalTemplateSettingLocale `json:"ar,omitempty"`
	// label for Amazon auth button
	AuthButtonAmazon *string `json:"authButtonAmazon,omitempty"`
	// label for Azure auth button
	AuthButtonAzure *string `json:"authButtonAzure,omitempty"`
	// label for Email auth button
	AuthButtonEmail *string `json:"authButtonEmail,omitempty"`
	// label for Facebook auth button
	AuthButtonFacebook *string `json:"authButtonFacebook,omitempty"`
	// label for Google auth button
	AuthButtonGoogle *string `json:"authButtonGoogle,omitempty"`
	// label for Microsoft auth button
	AuthButtonMicrosoft *string `json:"authButtonMicrosoft,omitempty"`
	// label for passphrase auth button
	AuthButtonPassphrase *string `json:"authButtonPassphrase,omitempty"`
	// label for SMS auth button
	AuthButtonSms *string `json:"authButtonSms,omitempty"`
	// label for Sponsor auth button
	AuthButtonSponsor *string `json:"authButtonSponsor,omitempty"`
	AuthLabel         *string `json:"authLabel,omitempty"`
	// label of the link to go back to /logon
	BackLink *string                          `json:"backLink,omitempty"`
	CaES     *WlanPortalTemplateSettingLocale `json:"ca-ES,omitempty"`
	// Portal main color
	Color      *string `json:"color,omitempty"`
	ColorDark  *string `json:"colorDark,omitempty"`
	ColorLight *string `json:"colorLight,omitempty"`
	// whether company field is required
	Company *bool `json:"company,omitempty"`
	// error message when company not provided
	CompanyError *string `json:"companyError,omitempty"`
	// label of company field
	CompanyLabel *string                          `json:"companyLabel,omitempty"`
	CsCZ         *WlanPortalTemplateSettingLocale `json:"cs-CZ,omitempty"`
	DaDK         *WlanPortalTemplateSettingLocale `json:"da-DK,omitempty"`
	DeDE         *WlanPortalTemplateSettingLocale `json:"de-DE,omitempty"`
	ElGR         *WlanPortalTemplateSettingLocale `json:"el-GR,omitempty"`
	// whether email field is required
	Email *bool `json:"email,omitempty"`
	// error message when a user has valid social login but doesn't match specified email domains.
	EmailAccessDomainError *string `json:"emailAccessDomainError,omitempty"`
	// Label for cancel confirmation code submission using email auth
	EmailCancel         *string `json:"emailCancel,omitempty"`
	EmailCodeCancel     *string `json:"emailCodeCancel,omitempty"`
	EmailCodeError      *string `json:"emailCodeError,omitempty"`
	EmailCodeFieldLabel *string `json:"emailCodeFieldLabel,omitempty"`
	EmailCodeMessage    *string `json:"emailCodeMessage,omitempty"`
	EmailCodeSubmit     *string `json:"emailCodeSubmit,omitempty"`
	EmailCodeTitle      *string `json:"emailCodeTitle,omitempty"`
	// error message when email not provided
	EmailError      *string `json:"emailError,omitempty"`
	EmailFieldLabel *string `json:"emailFieldLabel,omitempty"`
	// label of email field
	EmailLabel   *string `json:"emailLabel,omitempty"`
	EmailMessage *string `json:"emailMessage,omitempty"`
	// Label for confirmation code submit button using email auth
	EmailSubmit *string `json:"emailSubmit,omitempty"`
	// Title for the Email registration
	EmailTitle *string                          `json:"emailTitle,omitempty"`
	EnGB       *WlanPortalTemplateSettingLocale `json:"en-GB,omitempty"`
	EnUS       *WlanPortalTemplateSettingLocale `json:"en-US,omitempty"`
	EsES       *WlanPortalTemplateSettingLocale `json:"es-ES,omitempty"`
	FiFI       *WlanPortalTemplateSettingLocale `json:"fi-FI,omitempty"`
	// whether to ask field1
	Field1 *bool `json:"field1,omitempty"`
	// error message when field1 not provided
	Field1Error *string `json:"field1Error,omitempty"`
	// label of field1
	Field1Label *string `json:"field1Label,omitempty"`
	// whether field1 is required field
	Field1Required *bool `json:"field1Required,omitempty"`
	// whether to ask field2
	Field2 *bool `json:"field2,omitempty"`
	// error message when field2 not provided
	Field2Error *string `json:"field2Error,omitempty"`
	// label of field2
	Field2Label *string `json:"field2Label,omitempty"`
	// whether field2 is required field
	Field2Required *bool `json:"field2Required,omitempty"`
	// whether to ask field3
	Field3 *bool `json:"field3,omitempty"`
	// error message when field3 not provided
	Field3Error *string `json:"field3Error,omitempty"`
	// label of field3
	Field3Label *string `json:"field3Label,omitempty"`
	// whether field3 is required field
	Field3Required *bool `json:"field3Required,omitempty"`
	// whether to ask field4
	Field4 *bool `json:"field4,omitempty"`
	// error message when field4 not provided
	Field4Error *string `json:"field4Error,omitempty"`
	// label of field4
	Field4Label *string `json:"field4Label,omitempty"`
	// whether field4 is required field
	Field4Required *bool                            `json:"field4Required,omitempty"`
	FrFR           *WlanPortalTemplateSettingLocale `json:"fr-FR,omitempty"`
	HeIL           *WlanPortalTemplateSettingLocale `json:"he-IL,omitempty"`
	HiIN           *WlanPortalTemplateSettingLocale `json:"hi-IN,omitempty"`
	HrHR           *WlanPortalTemplateSettingLocale `json:"hr-HR,omitempty"`
	HuHU           *WlanPortalTemplateSettingLocale `json:"hu-HU,omitempty"`
	IdID           *WlanPortalTemplateSettingLocale `json:"id-ID,omitempty"`
	ItIT           *WlanPortalTemplateSettingLocale `json:"it-IT,omitempty"`
	JaJP           *WlanPortalTemplateSettingLocale `json:"ja-JP,omitempty"`
	KoKR           *WlanPortalTemplateSettingLocale `json:"ko-KR,omitempty"`
	// custom logo with `data:image/png;base64,` format, default null, uses Juniper Mist Logo. File size must be less than 100kB and image dimensions must be less than 500px x 200px (width x height).
	Logo Optional[string] `json:"logo"`
	// height of the logo, in px
	LogoHeight *int `json:"logoHeight,omitempty"`
	// width of the logo, in px
	LogoWidth *int                             `json:"logoWidth,omitempty"`
	Message   *string                          `json:"message,omitempty"`
	MsMY      *WlanPortalTemplateSettingLocale `json:"ms-MY,omitempty"`
	MultiAuth *bool                            `json:"multiAuth,omitempty"`
	// whether name field is required
	Name *bool `json:"name,omitempty"`
	// error message when name not provided
	NameError *string `json:"nameError,omitempty"`
	// label of name field
	NameLabel *string                          `json:"nameLabel,omitempty"`
	NbNO      *WlanPortalTemplateSettingLocale `json:"nb-NO,omitempty"`
	NlNL      *WlanPortalTemplateSettingLocale `json:"nl-NL,omitempty"`
	// Default value for the `Do not store` checkbox
	OptOutDefault *bool `json:"optOutDefault,omitempty"`
	// whether to display Do Not Store My Personal Information
	Optout *bool `json:"optout,omitempty"`
	// label for Do Not Store My Personal Information
	OptoutLabel *string `json:"optoutLabel,omitempty"`
	PageTitle   string  `json:"pageTitle"`
	// Label for the Passphrase cancel button
	PassphraseCancel *string `json:"passphraseCancel,omitempty"`
	// error message when invalid passphrase is provided
	PassphraseError *string `json:"passphraseError,omitempty"`
	// Passphrase
	PassphraseLabel   *string `json:"passphraseLabel,omitempty"`
	PassphraseMessage *string `json:"passphraseMessage,omitempty"`
	// Label for the Passphrase submit button
	PassphraseSubmit *string `json:"passphraseSubmit,omitempty"`
	// Title for passphrase details page
	PassphraseTitle *string                          `json:"passphraseTitle,omitempty"`
	PlPL            *WlanPortalTemplateSettingLocale `json:"pl-PL,omitempty"`
	// whether to show \"Powered by Mist\"
	PoweredBy *bool `json:"poweredBy,omitempty"`
	// wheter to require the Privacy Term acceptance
	Privacy *bool `json:"privacy,omitempty"`
	// prefix of the label of the link to go to Privacy Policy
	PrivacyPolicyAcceptLabel *string `json:"privacyPolicyAcceptLabel,omitempty"`
	// error message when Privacy Policy not accepted
	PrivacyPolicyError *string `json:"privacyPolicyError,omitempty"`
	// label of the link to go to Privacy Policy
	PrivacyPolicyLink *string `json:"privacyPolicyLink,omitempty"`
	// text of the Privacy Policy
	PrivacyPolicyText *string                          `json:"privacyPolicyText,omitempty"`
	PtBR              *WlanPortalTemplateSettingLocale `json:"pt-BR,omitempty"`
	PtPT              *WlanPortalTemplateSettingLocale `json:"pt-PT,omitempty"`
	// label to denote required field
	RequiredFieldLabel *string                          `json:"requiredFieldLabel,omitempty"`
	ResponsiveLayout   *bool                            `json:"responsiveLayout,omitempty"`
	RoRO               *WlanPortalTemplateSettingLocale `json:"ro-RO,omitempty"`
	RuRU               *WlanPortalTemplateSettingLocale `json:"ru-RU,omitempty"`
	// label of the button to /signin
	SignInLabel       *string                          `json:"signInLabel,omitempty"`
	SkSK              *WlanPortalTemplateSettingLocale `json:"sk-SK,omitempty"`
	SmsCarrierDefault *string                          `json:"smsCarrierDefault,omitempty"`
	SmsCarrierError   *string                          `json:"smsCarrierError,omitempty"`
	// label for mobile carrier drop-down list
	SmsCarrierFieldLabel *string `json:"smsCarrierFieldLabel,omitempty"`
	// Label for cancel confirmation code submission
	SmsCodeCancel *string `json:"smsCodeCancel,omitempty"`
	// error message when confirmation code is invalid
	SmsCodeError      *string `json:"smsCodeError,omitempty"`
	SmsCodeFieldLabel *string `json:"smsCodeFieldLabel,omitempty"`
	SmsCodeMessage    *string `json:"smsCodeMessage,omitempty"`
	// Label for confirmation code submit button
	SmsCodeSubmit        *string `json:"smsCodeSubmit,omitempty"`
	SmsCodeTitle         *string `json:"smsCodeTitle,omitempty"`
	SmsCountryFieldLabel *string `json:"smsCountryFieldLabel,omitempty"`
	SmsCountryFormat     *string `json:"smsCountryFormat,omitempty"`
	// Label for checkbox to specify that the user has access code
	SmsHaveAccessCode *string `json:"smsHaveAccessCode,omitempty"`
	SmsIsTwilio       *bool   `json:"smsIsTwilio,omitempty"`
	// format of access code sms message. {{code}} and {{duration}} are place holders and should be retained as is.
	SmsMessageFormat *string `json:"smsMessageFormat,omitempty"`
	// label for canceling mobile details for SMS auth
	SmsNumberCancel *string `json:"smsNumberCancel,omitempty"`
	SmsNumberError  *string `json:"smsNumberError,omitempty"`
	// label for field to provide mobile number
	SmsNumberFieldLabel *string `json:"smsNumberFieldLabel,omitempty"`
	SmsNumberFormat     *string `json:"smsNumberFormat,omitempty"`
	SmsNumberMessage    *string `json:"smsNumberMessage,omitempty"`
	// label for submit button for code generation
	SmsNumberSubmit *string `json:"smsNumberSubmit,omitempty"`
	// Title for phone number details
	SmsNumberTitle    *string `json:"smsNumberTitle,omitempty"`
	SmsUsernameFormat *string `json:"smsUsernameFormat,omitempty"`
	// how long confirmation code should be considered valid (in minutes)
	SmsValidityDuration *int    `json:"smsValidityDuration,omitempty"`
	SponsorBackLink     *string `json:"sponsorBackLink,omitempty"`
	SponsorCancel       *string `json:"sponsorCancel,omitempty"`
	// label for Sponsor Email
	SponsorEmail      *string `json:"sponsorEmail,omitempty"`
	SponsorEmailError *string `json:"sponsorEmailError,omitempty"`
	// html template to replace/override default sponsor email template
	// Sponsor Email Template supports following template variables:
	// * `approve_url`: Renders URL to approve the request; optionally &minutes=N query param can be appended to change the Authorization period of the guest, where N is a valid integer denoting number of minutes a guest remains authorized
	// * `deny_url`: Renders URL to reject the request
	// * `guest_email`: Renders Email ID of the guest
	// * `guest_name`: Renders Name of the guest
	// * `field1`: Renders value of the Custom Field 1
	// * `field2`: Renders value of the Custom Field 2
	// * `sponsor_link_validity_duration`: Renders validity time of the request (i.e. Approve/Deny URL)
	// * `auth_expire_minutes`: Renders Wlan-level configured Guest Authorization Expiration time period (in minutes), If not configured then default (1 day in minutes)
	SponsorEmailTemplate *string `json:"sponsorEmailTemplate,omitempty"`
	SponsorInfoApproved  *string `json:"sponsorInfoApproved,omitempty"`
	SponsorInfoDenied    *string `json:"sponsorInfoDenied,omitempty"`
	SponsorInfoPending   *string `json:"sponsorInfoPending,omitempty"`
	// label for Sponsor Name
	SponsorName        *string `json:"sponsorName,omitempty"`
	SponsorNameError   *string `json:"sponsorNameError,omitempty"`
	SponsorNotePending *string `json:"sponsorNotePending,omitempty"`
	// submit button label request Wifi Access and notify sponsor about guest request
	SponsorRequestAccess *string `json:"sponsorRequestAccess,omitempty"`
	// text to display if sponsor approves request
	SponsorStatusApproved *string `json:"sponsorStatusApproved,omitempty"`
	// text to display when sponsor denies request
	SponsorStatusDenied *string `json:"sponsorStatusDenied,omitempty"`
	// text to display if request is still pending
	SponsorStatusPending *string `json:"sponsorStatusPending,omitempty"`
	// submit button label to notify sponsor about guest request
	SponsorSubmit      *string                          `json:"sponsorSubmit,omitempty"`
	SponsorsError      *string                          `json:"sponsorsError,omitempty"`
	SponsorsFieldLabel *string                          `json:"sponsorsFieldLabel,omitempty"`
	SvSE               *WlanPortalTemplateSettingLocale `json:"sv-SE,omitempty"`
	ThTH               *WlanPortalTemplateSettingLocale `json:"th-TH,omitempty"`
	Tos                *bool                            `json:"tos,omitempty"`
	// prefix of the label of the link to go to tos
	TosAcceptLabel *string `json:"tosAcceptLabel,omitempty"`
	// error message when tos not accepted
	TosError *string `json:"tosError,omitempty"`
	// label of the link to go to tos
	TosLink *string `json:"tosLink,omitempty"`
	// text of the Terms of Service
	TosText              *string                          `json:"tosText,omitempty"`
	TrTR                 *WlanPortalTemplateSettingLocale `json:"tr-TR,omitempty"`
	UkUA                 *WlanPortalTemplateSettingLocale `json:"uk-UA,omitempty"`
	ViVN                 *WlanPortalTemplateSettingLocale `json:"vi-VN,omitempty"`
	ZhHans               *WlanPortalTemplateSettingLocale `json:"zh-Hans,omitempty"`
	ZhHant               *WlanPortalTemplateSettingLocale `json:"zh-Hant,omitempty"`
	AdditionalProperties map[string]any                   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WlanPortalTemplateSetting.
// It customizes the JSON marshaling process for WlanPortalTemplateSetting objects.
func (w WlanPortalTemplateSetting) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(w.toMap())
}

// toMap converts the WlanPortalTemplateSetting object to a map representation for JSON marshaling.
func (w WlanPortalTemplateSetting) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, w.AdditionalProperties)
	if w.AccessCodeAlternateEmail != nil {
		structMap["accessCodeAlternateEmail"] = w.AccessCodeAlternateEmail
	}
	if w.Alignment != nil {
		structMap["alignment"] = w.Alignment
	}
	if w.Ar != nil {
		structMap["ar"] = w.Ar.toMap()
	}
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
	if w.CaES != nil {
		structMap["ca-ES"] = w.CaES.toMap()
	}
	if w.Color != nil {
		structMap["color"] = w.Color
	}
	if w.ColorDark != nil {
		structMap["colorDark"] = w.ColorDark
	}
	if w.ColorLight != nil {
		structMap["colorLight"] = w.ColorLight
	}
	if w.Company != nil {
		structMap["company"] = w.Company
	}
	if w.CompanyError != nil {
		structMap["companyError"] = w.CompanyError
	}
	if w.CompanyLabel != nil {
		structMap["companyLabel"] = w.CompanyLabel
	}
	if w.CsCZ != nil {
		structMap["cs-CZ"] = w.CsCZ.toMap()
	}
	if w.DaDK != nil {
		structMap["da-DK"] = w.DaDK.toMap()
	}
	if w.DeDE != nil {
		structMap["de-DE"] = w.DeDE.toMap()
	}
	if w.ElGR != nil {
		structMap["el-GR"] = w.ElGR.toMap()
	}
	if w.Email != nil {
		structMap["email"] = w.Email
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
	if w.EnGB != nil {
		structMap["en-GB"] = w.EnGB.toMap()
	}
	if w.EnUS != nil {
		structMap["en-US"] = w.EnUS.toMap()
	}
	if w.EsES != nil {
		structMap["es-ES"] = w.EsES.toMap()
	}
	if w.FiFI != nil {
		structMap["fi-FI"] = w.FiFI.toMap()
	}
	if w.Field1 != nil {
		structMap["field1"] = w.Field1
	}
	if w.Field1Error != nil {
		structMap["field1Error"] = w.Field1Error
	}
	if w.Field1Label != nil {
		structMap["field1Label"] = w.Field1Label
	}
	if w.Field1Required != nil {
		structMap["field1Required"] = w.Field1Required
	}
	if w.Field2 != nil {
		structMap["field2"] = w.Field2
	}
	if w.Field2Error != nil {
		structMap["field2Error"] = w.Field2Error
	}
	if w.Field2Label != nil {
		structMap["field2Label"] = w.Field2Label
	}
	if w.Field2Required != nil {
		structMap["field2Required"] = w.Field2Required
	}
	if w.Field3 != nil {
		structMap["field3"] = w.Field3
	}
	if w.Field3Error != nil {
		structMap["field3Error"] = w.Field3Error
	}
	if w.Field3Label != nil {
		structMap["field3Label"] = w.Field3Label
	}
	if w.Field3Required != nil {
		structMap["field3Required"] = w.Field3Required
	}
	if w.Field4 != nil {
		structMap["field4"] = w.Field4
	}
	if w.Field4Error != nil {
		structMap["field4Error"] = w.Field4Error
	}
	if w.Field4Label != nil {
		structMap["field4Label"] = w.Field4Label
	}
	if w.Field4Required != nil {
		structMap["field4Required"] = w.Field4Required
	}
	if w.FrFR != nil {
		structMap["fr-FR"] = w.FrFR.toMap()
	}
	if w.HeIL != nil {
		structMap["he-IL"] = w.HeIL.toMap()
	}
	if w.HiIN != nil {
		structMap["hi-IN"] = w.HiIN.toMap()
	}
	if w.HrHR != nil {
		structMap["hr-HR"] = w.HrHR.toMap()
	}
	if w.HuHU != nil {
		structMap["hu-HU"] = w.HuHU.toMap()
	}
	if w.IdID != nil {
		structMap["id-ID"] = w.IdID.toMap()
	}
	if w.ItIT != nil {
		structMap["it-IT"] = w.ItIT.toMap()
	}
	if w.JaJP != nil {
		structMap["ja-JP"] = w.JaJP.toMap()
	}
	if w.KoKR != nil {
		structMap["ko-KR"] = w.KoKR.toMap()
	}
	if w.Logo.IsValueSet() {
		if w.Logo.Value() != nil {
			structMap["logo"] = w.Logo.Value()
		} else {
			structMap["logo"] = nil
		}
	}
	if w.LogoHeight != nil {
		structMap["logoHeight"] = w.LogoHeight
	}
	if w.LogoWidth != nil {
		structMap["logoWidth"] = w.LogoWidth
	}
	if w.Message != nil {
		structMap["message"] = w.Message
	}
	if w.MsMY != nil {
		structMap["ms-MY"] = w.MsMY.toMap()
	}
	if w.MultiAuth != nil {
		structMap["multiAuth"] = w.MultiAuth
	}
	if w.Name != nil {
		structMap["name"] = w.Name
	}
	if w.NameError != nil {
		structMap["nameError"] = w.NameError
	}
	if w.NameLabel != nil {
		structMap["nameLabel"] = w.NameLabel
	}
	if w.NbNO != nil {
		structMap["nb-NO"] = w.NbNO.toMap()
	}
	if w.NlNL != nil {
		structMap["nl-NL"] = w.NlNL.toMap()
	}
	if w.OptOutDefault != nil {
		structMap["optOutDefault"] = w.OptOutDefault
	}
	if w.Optout != nil {
		structMap["optout"] = w.Optout
	}
	if w.OptoutLabel != nil {
		structMap["optoutLabel"] = w.OptoutLabel
	}
	structMap["pageTitle"] = w.PageTitle
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
	if w.PlPL != nil {
		structMap["pl-PL"] = w.PlPL.toMap()
	}
	if w.PoweredBy != nil {
		structMap["poweredBy"] = w.PoweredBy
	}
	if w.Privacy != nil {
		structMap["privacy"] = w.Privacy
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
	if w.PtBR != nil {
		structMap["pt-BR"] = w.PtBR.toMap()
	}
	if w.PtPT != nil {
		structMap["pt-PT"] = w.PtPT.toMap()
	}
	if w.RequiredFieldLabel != nil {
		structMap["requiredFieldLabel"] = w.RequiredFieldLabel
	}
	if w.ResponsiveLayout != nil {
		structMap["responsiveLayout"] = w.ResponsiveLayout
	}
	if w.RoRO != nil {
		structMap["ro-RO"] = w.RoRO.toMap()
	}
	if w.RuRU != nil {
		structMap["ru-RU"] = w.RuRU.toMap()
	}
	if w.SignInLabel != nil {
		structMap["signInLabel"] = w.SignInLabel
	}
	if w.SkSK != nil {
		structMap["sk-SK"] = w.SkSK.toMap()
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
	if w.SmsIsTwilio != nil {
		structMap["smsIsTwilio"] = w.SmsIsTwilio
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
	if w.SmsValidityDuration != nil {
		structMap["smsValidityDuration"] = w.SmsValidityDuration
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
	if w.SponsorEmailTemplate != nil {
		structMap["sponsorEmailTemplate"] = w.SponsorEmailTemplate
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
	if w.SvSE != nil {
		structMap["sv-SE"] = w.SvSE.toMap()
	}
	if w.ThTH != nil {
		structMap["th-TH"] = w.ThTH.toMap()
	}
	if w.Tos != nil {
		structMap["tos"] = w.Tos
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
	if w.TrTR != nil {
		structMap["tr-TR"] = w.TrTR.toMap()
	}
	if w.UkUA != nil {
		structMap["uk-UA"] = w.UkUA.toMap()
	}
	if w.ViVN != nil {
		structMap["vi-VN"] = w.ViVN.toMap()
	}
	if w.ZhHans != nil {
		structMap["zh-Hans"] = w.ZhHans.toMap()
	}
	if w.ZhHant != nil {
		structMap["zh-Hant"] = w.ZhHant.toMap()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WlanPortalTemplateSetting.
// It customizes the JSON unmarshaling process for WlanPortalTemplateSetting objects.
func (w *WlanPortalTemplateSetting) UnmarshalJSON(input []byte) error {
	var temp tempWlanPortalTemplateSetting
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "accessCodeAlternateEmail", "alignment", "ar", "authButtonAmazon", "authButtonAzure", "authButtonEmail", "authButtonFacebook", "authButtonGoogle", "authButtonMicrosoft", "authButtonPassphrase", "authButtonSms", "authButtonSponsor", "authLabel", "backLink", "ca-ES", "color", "colorDark", "colorLight", "company", "companyError", "companyLabel", "cs-CZ", "da-DK", "de-DE", "el-GR", "email", "emailAccessDomainError", "emailCancel", "emailCodeCancel", "emailCodeError", "emailCodeFieldLabel", "emailCodeMessage", "emailCodeSubmit", "emailCodeTitle", "emailError", "emailFieldLabel", "emailLabel", "emailMessage", "emailSubmit", "emailTitle", "en-GB", "en-US", "es-ES", "fi-FI", "field1", "field1Error", "field1Label", "field1Required", "field2", "field2Error", "field2Label", "field2Required", "field3", "field3Error", "field3Label", "field3Required", "field4", "field4Error", "field4Label", "field4Required", "fr-FR", "he-IL", "hi-IN", "hr-HR", "hu-HU", "id-ID", "it-IT", "ja-JP", "ko-KR", "logo", "logoHeight", "logoWidth", "message", "ms-MY", "multiAuth", "name", "nameError", "nameLabel", "nb-NO", "nl-NL", "optOutDefault", "optout", "optoutLabel", "pageTitle", "passphraseCancel", "passphraseError", "passphraseLabel", "passphraseMessage", "passphraseSubmit", "passphraseTitle", "pl-PL", "poweredBy", "privacy", "privacyPolicyAcceptLabel", "privacyPolicyError", "privacyPolicyLink", "privacyPolicyText", "pt-BR", "pt-PT", "requiredFieldLabel", "responsiveLayout", "ro-RO", "ru-RU", "signInLabel", "sk-SK", "smsCarrierDefault", "smsCarrierError", "smsCarrierFieldLabel", "smsCodeCancel", "smsCodeError", "smsCodeFieldLabel", "smsCodeMessage", "smsCodeSubmit", "smsCodeTitle", "smsCountryFieldLabel", "smsCountryFormat", "smsHaveAccessCode", "smsIsTwilio", "smsMessageFormat", "smsNumberCancel", "smsNumberError", "smsNumberFieldLabel", "smsNumberFormat", "smsNumberMessage", "smsNumberSubmit", "smsNumberTitle", "smsUsernameFormat", "smsValidityDuration", "sponsorBackLink", "sponsorCancel", "sponsorEmail", "sponsorEmailError", "sponsorEmailTemplate", "sponsorInfoApproved", "sponsorInfoDenied", "sponsorInfoPending", "sponsorName", "sponsorNameError", "sponsorNotePending", "sponsorRequestAccess", "sponsorStatusApproved", "sponsorStatusDenied", "sponsorStatusPending", "sponsorSubmit", "sponsorsError", "sponsorsFieldLabel", "sv-SE", "th-TH", "tos", "tosAcceptLabel", "tosError", "tosLink", "tosText", "tr-TR", "uk-UA", "vi-VN", "zh-Hans", "zh-Hant")
	if err != nil {
		return err
	}

	w.AdditionalProperties = additionalProperties
	w.AccessCodeAlternateEmail = temp.AccessCodeAlternateEmail
	w.Alignment = temp.Alignment
	w.Ar = temp.Ar
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
	w.CaES = temp.CaES
	w.Color = temp.Color
	w.ColorDark = temp.ColorDark
	w.ColorLight = temp.ColorLight
	w.Company = temp.Company
	w.CompanyError = temp.CompanyError
	w.CompanyLabel = temp.CompanyLabel
	w.CsCZ = temp.CsCZ
	w.DaDK = temp.DaDK
	w.DeDE = temp.DeDE
	w.ElGR = temp.ElGR
	w.Email = temp.Email
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
	w.EnGB = temp.EnGB
	w.EnUS = temp.EnUS
	w.EsES = temp.EsES
	w.FiFI = temp.FiFI
	w.Field1 = temp.Field1
	w.Field1Error = temp.Field1Error
	w.Field1Label = temp.Field1Label
	w.Field1Required = temp.Field1Required
	w.Field2 = temp.Field2
	w.Field2Error = temp.Field2Error
	w.Field2Label = temp.Field2Label
	w.Field2Required = temp.Field2Required
	w.Field3 = temp.Field3
	w.Field3Error = temp.Field3Error
	w.Field3Label = temp.Field3Label
	w.Field3Required = temp.Field3Required
	w.Field4 = temp.Field4
	w.Field4Error = temp.Field4Error
	w.Field4Label = temp.Field4Label
	w.Field4Required = temp.Field4Required
	w.FrFR = temp.FrFR
	w.HeIL = temp.HeIL
	w.HiIN = temp.HiIN
	w.HrHR = temp.HrHR
	w.HuHU = temp.HuHU
	w.IdID = temp.IdID
	w.ItIT = temp.ItIT
	w.JaJP = temp.JaJP
	w.KoKR = temp.KoKR
	w.Logo = temp.Logo
	w.LogoHeight = temp.LogoHeight
	w.LogoWidth = temp.LogoWidth
	w.Message = temp.Message
	w.MsMY = temp.MsMY
	w.MultiAuth = temp.MultiAuth
	w.Name = temp.Name
	w.NameError = temp.NameError
	w.NameLabel = temp.NameLabel
	w.NbNO = temp.NbNO
	w.NlNL = temp.NlNL
	w.OptOutDefault = temp.OptOutDefault
	w.Optout = temp.Optout
	w.OptoutLabel = temp.OptoutLabel
	w.PageTitle = *temp.PageTitle
	w.PassphraseCancel = temp.PassphraseCancel
	w.PassphraseError = temp.PassphraseError
	w.PassphraseLabel = temp.PassphraseLabel
	w.PassphraseMessage = temp.PassphraseMessage
	w.PassphraseSubmit = temp.PassphraseSubmit
	w.PassphraseTitle = temp.PassphraseTitle
	w.PlPL = temp.PlPL
	w.PoweredBy = temp.PoweredBy
	w.Privacy = temp.Privacy
	w.PrivacyPolicyAcceptLabel = temp.PrivacyPolicyAcceptLabel
	w.PrivacyPolicyError = temp.PrivacyPolicyError
	w.PrivacyPolicyLink = temp.PrivacyPolicyLink
	w.PrivacyPolicyText = temp.PrivacyPolicyText
	w.PtBR = temp.PtBR
	w.PtPT = temp.PtPT
	w.RequiredFieldLabel = temp.RequiredFieldLabel
	w.ResponsiveLayout = temp.ResponsiveLayout
	w.RoRO = temp.RoRO
	w.RuRU = temp.RuRU
	w.SignInLabel = temp.SignInLabel
	w.SkSK = temp.SkSK
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
	w.SmsIsTwilio = temp.SmsIsTwilio
	w.SmsMessageFormat = temp.SmsMessageFormat
	w.SmsNumberCancel = temp.SmsNumberCancel
	w.SmsNumberError = temp.SmsNumberError
	w.SmsNumberFieldLabel = temp.SmsNumberFieldLabel
	w.SmsNumberFormat = temp.SmsNumberFormat
	w.SmsNumberMessage = temp.SmsNumberMessage
	w.SmsNumberSubmit = temp.SmsNumberSubmit
	w.SmsNumberTitle = temp.SmsNumberTitle
	w.SmsUsernameFormat = temp.SmsUsernameFormat
	w.SmsValidityDuration = temp.SmsValidityDuration
	w.SponsorBackLink = temp.SponsorBackLink
	w.SponsorCancel = temp.SponsorCancel
	w.SponsorEmail = temp.SponsorEmail
	w.SponsorEmailError = temp.SponsorEmailError
	w.SponsorEmailTemplate = temp.SponsorEmailTemplate
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
	w.SvSE = temp.SvSE
	w.ThTH = temp.ThTH
	w.Tos = temp.Tos
	w.TosAcceptLabel = temp.TosAcceptLabel
	w.TosError = temp.TosError
	w.TosLink = temp.TosLink
	w.TosText = temp.TosText
	w.TrTR = temp.TrTR
	w.UkUA = temp.UkUA
	w.ViVN = temp.ViVN
	w.ZhHans = temp.ZhHans
	w.ZhHant = temp.ZhHant
	return nil
}

// tempWlanPortalTemplateSetting is a temporary struct used for validating the fields of WlanPortalTemplateSetting.
type tempWlanPortalTemplateSetting struct {
	AccessCodeAlternateEmail *string                          `json:"accessCodeAlternateEmail,omitempty"`
	Alignment                *PortalTemplateAlignmentEnum     `json:"alignment,omitempty"`
	Ar                       *WlanPortalTemplateSettingLocale `json:"ar,omitempty"`
	AuthButtonAmazon         *string                          `json:"authButtonAmazon,omitempty"`
	AuthButtonAzure          *string                          `json:"authButtonAzure,omitempty"`
	AuthButtonEmail          *string                          `json:"authButtonEmail,omitempty"`
	AuthButtonFacebook       *string                          `json:"authButtonFacebook,omitempty"`
	AuthButtonGoogle         *string                          `json:"authButtonGoogle,omitempty"`
	AuthButtonMicrosoft      *string                          `json:"authButtonMicrosoft,omitempty"`
	AuthButtonPassphrase     *string                          `json:"authButtonPassphrase,omitempty"`
	AuthButtonSms            *string                          `json:"authButtonSms,omitempty"`
	AuthButtonSponsor        *string                          `json:"authButtonSponsor,omitempty"`
	AuthLabel                *string                          `json:"authLabel,omitempty"`
	BackLink                 *string                          `json:"backLink,omitempty"`
	CaES                     *WlanPortalTemplateSettingLocale `json:"ca-ES,omitempty"`
	Color                    *string                          `json:"color,omitempty"`
	ColorDark                *string                          `json:"colorDark,omitempty"`
	ColorLight               *string                          `json:"colorLight,omitempty"`
	Company                  *bool                            `json:"company,omitempty"`
	CompanyError             *string                          `json:"companyError,omitempty"`
	CompanyLabel             *string                          `json:"companyLabel,omitempty"`
	CsCZ                     *WlanPortalTemplateSettingLocale `json:"cs-CZ,omitempty"`
	DaDK                     *WlanPortalTemplateSettingLocale `json:"da-DK,omitempty"`
	DeDE                     *WlanPortalTemplateSettingLocale `json:"de-DE,omitempty"`
	ElGR                     *WlanPortalTemplateSettingLocale `json:"el-GR,omitempty"`
	Email                    *bool                            `json:"email,omitempty"`
	EmailAccessDomainError   *string                          `json:"emailAccessDomainError,omitempty"`
	EmailCancel              *string                          `json:"emailCancel,omitempty"`
	EmailCodeCancel          *string                          `json:"emailCodeCancel,omitempty"`
	EmailCodeError           *string                          `json:"emailCodeError,omitempty"`
	EmailCodeFieldLabel      *string                          `json:"emailCodeFieldLabel,omitempty"`
	EmailCodeMessage         *string                          `json:"emailCodeMessage,omitempty"`
	EmailCodeSubmit          *string                          `json:"emailCodeSubmit,omitempty"`
	EmailCodeTitle           *string                          `json:"emailCodeTitle,omitempty"`
	EmailError               *string                          `json:"emailError,omitempty"`
	EmailFieldLabel          *string                          `json:"emailFieldLabel,omitempty"`
	EmailLabel               *string                          `json:"emailLabel,omitempty"`
	EmailMessage             *string                          `json:"emailMessage,omitempty"`
	EmailSubmit              *string                          `json:"emailSubmit,omitempty"`
	EmailTitle               *string                          `json:"emailTitle,omitempty"`
	EnGB                     *WlanPortalTemplateSettingLocale `json:"en-GB,omitempty"`
	EnUS                     *WlanPortalTemplateSettingLocale `json:"en-US,omitempty"`
	EsES                     *WlanPortalTemplateSettingLocale `json:"es-ES,omitempty"`
	FiFI                     *WlanPortalTemplateSettingLocale `json:"fi-FI,omitempty"`
	Field1                   *bool                            `json:"field1,omitempty"`
	Field1Error              *string                          `json:"field1Error,omitempty"`
	Field1Label              *string                          `json:"field1Label,omitempty"`
	Field1Required           *bool                            `json:"field1Required,omitempty"`
	Field2                   *bool                            `json:"field2,omitempty"`
	Field2Error              *string                          `json:"field2Error,omitempty"`
	Field2Label              *string                          `json:"field2Label,omitempty"`
	Field2Required           *bool                            `json:"field2Required,omitempty"`
	Field3                   *bool                            `json:"field3,omitempty"`
	Field3Error              *string                          `json:"field3Error,omitempty"`
	Field3Label              *string                          `json:"field3Label,omitempty"`
	Field3Required           *bool                            `json:"field3Required,omitempty"`
	Field4                   *bool                            `json:"field4,omitempty"`
	Field4Error              *string                          `json:"field4Error,omitempty"`
	Field4Label              *string                          `json:"field4Label,omitempty"`
	Field4Required           *bool                            `json:"field4Required,omitempty"`
	FrFR                     *WlanPortalTemplateSettingLocale `json:"fr-FR,omitempty"`
	HeIL                     *WlanPortalTemplateSettingLocale `json:"he-IL,omitempty"`
	HiIN                     *WlanPortalTemplateSettingLocale `json:"hi-IN,omitempty"`
	HrHR                     *WlanPortalTemplateSettingLocale `json:"hr-HR,omitempty"`
	HuHU                     *WlanPortalTemplateSettingLocale `json:"hu-HU,omitempty"`
	IdID                     *WlanPortalTemplateSettingLocale `json:"id-ID,omitempty"`
	ItIT                     *WlanPortalTemplateSettingLocale `json:"it-IT,omitempty"`
	JaJP                     *WlanPortalTemplateSettingLocale `json:"ja-JP,omitempty"`
	KoKR                     *WlanPortalTemplateSettingLocale `json:"ko-KR,omitempty"`
	Logo                     Optional[string]                 `json:"logo"`
	LogoHeight               *int                             `json:"logoHeight,omitempty"`
	LogoWidth                *int                             `json:"logoWidth,omitempty"`
	Message                  *string                          `json:"message,omitempty"`
	MsMY                     *WlanPortalTemplateSettingLocale `json:"ms-MY,omitempty"`
	MultiAuth                *bool                            `json:"multiAuth,omitempty"`
	Name                     *bool                            `json:"name,omitempty"`
	NameError                *string                          `json:"nameError,omitempty"`
	NameLabel                *string                          `json:"nameLabel,omitempty"`
	NbNO                     *WlanPortalTemplateSettingLocale `json:"nb-NO,omitempty"`
	NlNL                     *WlanPortalTemplateSettingLocale `json:"nl-NL,omitempty"`
	OptOutDefault            *bool                            `json:"optOutDefault,omitempty"`
	Optout                   *bool                            `json:"optout,omitempty"`
	OptoutLabel              *string                          `json:"optoutLabel,omitempty"`
	PageTitle                *string                          `json:"pageTitle"`
	PassphraseCancel         *string                          `json:"passphraseCancel,omitempty"`
	PassphraseError          *string                          `json:"passphraseError,omitempty"`
	PassphraseLabel          *string                          `json:"passphraseLabel,omitempty"`
	PassphraseMessage        *string                          `json:"passphraseMessage,omitempty"`
	PassphraseSubmit         *string                          `json:"passphraseSubmit,omitempty"`
	PassphraseTitle          *string                          `json:"passphraseTitle,omitempty"`
	PlPL                     *WlanPortalTemplateSettingLocale `json:"pl-PL,omitempty"`
	PoweredBy                *bool                            `json:"poweredBy,omitempty"`
	Privacy                  *bool                            `json:"privacy,omitempty"`
	PrivacyPolicyAcceptLabel *string                          `json:"privacyPolicyAcceptLabel,omitempty"`
	PrivacyPolicyError       *string                          `json:"privacyPolicyError,omitempty"`
	PrivacyPolicyLink        *string                          `json:"privacyPolicyLink,omitempty"`
	PrivacyPolicyText        *string                          `json:"privacyPolicyText,omitempty"`
	PtBR                     *WlanPortalTemplateSettingLocale `json:"pt-BR,omitempty"`
	PtPT                     *WlanPortalTemplateSettingLocale `json:"pt-PT,omitempty"`
	RequiredFieldLabel       *string                          `json:"requiredFieldLabel,omitempty"`
	ResponsiveLayout         *bool                            `json:"responsiveLayout,omitempty"`
	RoRO                     *WlanPortalTemplateSettingLocale `json:"ro-RO,omitempty"`
	RuRU                     *WlanPortalTemplateSettingLocale `json:"ru-RU,omitempty"`
	SignInLabel              *string                          `json:"signInLabel,omitempty"`
	SkSK                     *WlanPortalTemplateSettingLocale `json:"sk-SK,omitempty"`
	SmsCarrierDefault        *string                          `json:"smsCarrierDefault,omitempty"`
	SmsCarrierError          *string                          `json:"smsCarrierError,omitempty"`
	SmsCarrierFieldLabel     *string                          `json:"smsCarrierFieldLabel,omitempty"`
	SmsCodeCancel            *string                          `json:"smsCodeCancel,omitempty"`
	SmsCodeError             *string                          `json:"smsCodeError,omitempty"`
	SmsCodeFieldLabel        *string                          `json:"smsCodeFieldLabel,omitempty"`
	SmsCodeMessage           *string                          `json:"smsCodeMessage,omitempty"`
	SmsCodeSubmit            *string                          `json:"smsCodeSubmit,omitempty"`
	SmsCodeTitle             *string                          `json:"smsCodeTitle,omitempty"`
	SmsCountryFieldLabel     *string                          `json:"smsCountryFieldLabel,omitempty"`
	SmsCountryFormat         *string                          `json:"smsCountryFormat,omitempty"`
	SmsHaveAccessCode        *string                          `json:"smsHaveAccessCode,omitempty"`
	SmsIsTwilio              *bool                            `json:"smsIsTwilio,omitempty"`
	SmsMessageFormat         *string                          `json:"smsMessageFormat,omitempty"`
	SmsNumberCancel          *string                          `json:"smsNumberCancel,omitempty"`
	SmsNumberError           *string                          `json:"smsNumberError,omitempty"`
	SmsNumberFieldLabel      *string                          `json:"smsNumberFieldLabel,omitempty"`
	SmsNumberFormat          *string                          `json:"smsNumberFormat,omitempty"`
	SmsNumberMessage         *string                          `json:"smsNumberMessage,omitempty"`
	SmsNumberSubmit          *string                          `json:"smsNumberSubmit,omitempty"`
	SmsNumberTitle           *string                          `json:"smsNumberTitle,omitempty"`
	SmsUsernameFormat        *string                          `json:"smsUsernameFormat,omitempty"`
	SmsValidityDuration      *int                             `json:"smsValidityDuration,omitempty"`
	SponsorBackLink          *string                          `json:"sponsorBackLink,omitempty"`
	SponsorCancel            *string                          `json:"sponsorCancel,omitempty"`
	SponsorEmail             *string                          `json:"sponsorEmail,omitempty"`
	SponsorEmailError        *string                          `json:"sponsorEmailError,omitempty"`
	SponsorEmailTemplate     *string                          `json:"sponsorEmailTemplate,omitempty"`
	SponsorInfoApproved      *string                          `json:"sponsorInfoApproved,omitempty"`
	SponsorInfoDenied        *string                          `json:"sponsorInfoDenied,omitempty"`
	SponsorInfoPending       *string                          `json:"sponsorInfoPending,omitempty"`
	SponsorName              *string                          `json:"sponsorName,omitempty"`
	SponsorNameError         *string                          `json:"sponsorNameError,omitempty"`
	SponsorNotePending       *string                          `json:"sponsorNotePending,omitempty"`
	SponsorRequestAccess     *string                          `json:"sponsorRequestAccess,omitempty"`
	SponsorStatusApproved    *string                          `json:"sponsorStatusApproved,omitempty"`
	SponsorStatusDenied      *string                          `json:"sponsorStatusDenied,omitempty"`
	SponsorStatusPending     *string                          `json:"sponsorStatusPending,omitempty"`
	SponsorSubmit            *string                          `json:"sponsorSubmit,omitempty"`
	SponsorsError            *string                          `json:"sponsorsError,omitempty"`
	SponsorsFieldLabel       *string                          `json:"sponsorsFieldLabel,omitempty"`
	SvSE                     *WlanPortalTemplateSettingLocale `json:"sv-SE,omitempty"`
	ThTH                     *WlanPortalTemplateSettingLocale `json:"th-TH,omitempty"`
	Tos                      *bool                            `json:"tos,omitempty"`
	TosAcceptLabel           *string                          `json:"tosAcceptLabel,omitempty"`
	TosError                 *string                          `json:"tosError,omitempty"`
	TosLink                  *string                          `json:"tosLink,omitempty"`
	TosText                  *string                          `json:"tosText,omitempty"`
	TrTR                     *WlanPortalTemplateSettingLocale `json:"tr-TR,omitempty"`
	UkUA                     *WlanPortalTemplateSettingLocale `json:"uk-UA,omitempty"`
	ViVN                     *WlanPortalTemplateSettingLocale `json:"vi-VN,omitempty"`
	ZhHans                   *WlanPortalTemplateSettingLocale `json:"zh-Hans,omitempty"`
	ZhHant                   *WlanPortalTemplateSettingLocale `json:"zh-Hant,omitempty"`
}

func (w *tempWlanPortalTemplateSetting) validate() error {
	var errs []string
	if w.PageTitle == nil {
		errs = append(errs, "required field `pageTitle` is missing for type `wlan_portal_template_setting`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
