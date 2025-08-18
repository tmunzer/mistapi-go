// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
	"context"
	"github.com/apimatic/go-core-runtime/https"
	"github.com/google/uuid"
	"github.com/tmunzer/mistapi-go/mistapi/errors"
	"github.com/tmunzer/mistapi-go/mistapi/models"
	"net/http"
)

// UtilitiesWiFi represents a controller struct.
type UtilitiesWiFi struct {
	baseController
}

// NewUtilitiesWiFi creates a new instance of UtilitiesWiFi.
// It takes a baseController as a parameter and returns a pointer to the UtilitiesWiFi.
func NewUtilitiesWiFi(baseController baseController) *UtilitiesWiFi {
	utilitiesWiFi := UtilitiesWiFi{baseController: baseController}
	return &utilitiesWiFi
}

// ReauthOrgDot1xWirelessClient takes context, orgId, clientMac as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Trigger a CoA (change of authorization) against a client
func (u *UtilitiesWiFi) ReauthOrgDot1xWirelessClient(
	ctx context.Context,
	orgId uuid.UUID,
	clientMac string) (
	*http.Response,
	error) {
	req := u.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/clients/%v/coa")
	req.AppendTemplateParams(orgId, clientMac)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
			NewAndAuth(
				NewAuth("basicAuth"),
				NewAuth("csrfToken"),
			),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
	})

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}

// DisconnectSiteMultipleClients takes context, siteId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// To unauthorize multiple clients
func (u *UtilitiesWiFi) DisconnectSiteMultipleClients(
	ctx context.Context,
	siteId uuid.UUID,
	body []string) (
	*http.Response,
	error) {
	req := u.prepareRequest(ctx, "POST", "/api/v1/sites/%v/clients/disconnect")
	req.AppendTemplateParams(siteId)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
			NewAndAuth(
				NewAuth("basicAuth"),
				NewAuth("csrfToken"),
			),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}

// UnauthorizeSiteMultipleClients takes context, siteId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// This unauthorize clients (if they are guest) and disconnect them. From the guest’s perspective, they will see the splash page again and go through the flow (e.g. Terms of Use) again.
func (u *UtilitiesWiFi) UnauthorizeSiteMultipleClients(
	ctx context.Context,
	siteId uuid.UUID,
	body *models.MacAddresses) (
	*http.Response,
	error) {
	req := u.prepareRequest(ctx, "POST", "/api/v1/sites/%v/clients/unauthorize")
	req.AppendTemplateParams(siteId)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
			NewAndAuth(
				NewAuth("basicAuth"),
				NewAuth("csrfToken"),
			),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}

// ReauthSiteDot1xWirelessClient takes context, siteId, clientMac as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Trigger a CoA (change of authorization) against a Wireless client
func (u *UtilitiesWiFi) ReauthSiteDot1xWirelessClient(
	ctx context.Context,
	siteId uuid.UUID,
	clientMac string) (
	*http.Response,
	error) {
	req := u.prepareRequest(ctx, "POST", "/api/v1/sites/%v/clients/%v/coa")
	req.AppendTemplateParams(siteId, clientMac)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
			NewAndAuth(
				NewAuth("basicAuth"),
				NewAuth("csrfToken"),
			),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
	})

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}

// DisconnectSiteWirelessClient takes context, siteId, clientMac as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// This disconnect a client (and it’s likely to connect back)
func (u *UtilitiesWiFi) DisconnectSiteWirelessClient(
	ctx context.Context,
	siteId uuid.UUID,
	clientMac string) (
	*http.Response,
	error) {
	req := u.prepareRequest(ctx, "POST", "/api/v1/sites/%v/clients/%v/disconnect")
	req.AppendTemplateParams(siteId, clientMac)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
			NewAndAuth(
				NewAuth("basicAuth"),
				NewAuth("csrfToken"),
			),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
	})

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}

// UnauthorizeSiteWirelessClient takes context, siteId, clientMac as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// This unauthorize a client (if it’s a guest) and disconnect it. From the guest’s perspective, s/he will see the splash page again and go through the flow (e.g. Terms of Use) again.
func (u *UtilitiesWiFi) UnauthorizeSiteWirelessClient(
	ctx context.Context,
	siteId uuid.UUID,
	clientMac string) (
	*http.Response,
	error) {
	req := u.prepareRequest(ctx, "POST", "/api/v1/sites/%v/clients/%v/unauthorize")
	req.AppendTemplateParams(siteId, clientMac)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
			NewAndAuth(
				NewAuth("basicAuth"),
				NewAuth("csrfToken"),
			),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
	})

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}

// ReprovisionSiteAllDevices takes context, siteId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// To force all Devices to reprovision itself again.
func (u *UtilitiesWiFi) ReprovisionSiteAllDevices(
	ctx context.Context,
	siteId uuid.UUID) (
	*http.Response,
	error) {
	req := u.prepareRequest(ctx, "POST", "/api/v1/sites/%v/devices/reprovision")
	req.AppendTemplateParams(siteId)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
			NewAndAuth(
				NewAuth("basicAuth"),
				NewAuth("csrfToken"),
			),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
	})

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}

// ResetSiteAllApsToUseRrm takes context, siteId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Reset all APs in the Site to use RRM
func (u *UtilitiesWiFi) ResetSiteAllApsToUseRrm(
	ctx context.Context,
	siteId uuid.UUID,
	body *models.UtilsResetRadioConfig) (
	*http.Response,
	error) {
	req := u.prepareRequest(
		ctx,
		"POST",
		"/api/v1/sites/%v/devices/reset_radio_config",
	)
	req.AppendTemplateParams(siteId)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
			NewAndAuth(
				NewAuth("basicAuth"),
				NewAuth("csrfToken"),
			),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}

// ZeroizeSiteFipsAllAps takes context, siteId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Zeroize all FIPS APs in the Site
func (u *UtilitiesWiFi) ZeroizeSiteFipsAllAps(
	ctx context.Context,
	siteId uuid.UUID,
	body *models.UtilsZeroizeFips) (
	*http.Response,
	error) {
	req := u.prepareRequest(ctx, "POST", "/api/v1/sites/%v/devices/zeroize")
	req.AppendTemplateParams(siteId)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
			NewAndAuth(
				NewAuth("basicAuth"),
				NewAuth("csrfToken"),
			),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}

// DeauthSiteWirelessClientsConnectedToARogue takes context, siteId, rogueBssid as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Send Deauth frame to clients connected to a Rogue AP
func (u *UtilitiesWiFi) DeauthSiteWirelessClientsConnectedToARogue(
	ctx context.Context,
	siteId uuid.UUID,
	rogueBssid string) (
	*http.Response,
	error) {
	req := u.prepareRequest(ctx, "POST", "/api/v1/sites/%v/rogues/%v/deauth_clients")
	req.AppendTemplateParams(siteId, rogueBssid)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
			NewAndAuth(
				NewAuth("basicAuth"),
				NewAuth("csrfToken"),
			),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
	})

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}

// OptimizeSiteRrm takes context, siteId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Optimize Site RRM
func (u *UtilitiesWiFi) OptimizeSiteRrm(
	ctx context.Context,
	siteId uuid.UUID,
	body *models.UtilsRrmOptimize) (
	*http.Response,
	error) {
	req := u.prepareRequest(ctx, "POST", "/api/v1/sites/%v/rrm/optimize")
	req.AppendTemplateParams(siteId)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
			NewAndAuth(
				NewAuth("basicAuth"),
				NewAuth("csrfToken"),
			),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}

// TestSiteWlanSmsGlobal takes context, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Allows validation of Global sms gateway credentials.
// In case of success, a text message confirming successful setup should be received. In case of error, smsglobal error message are returned.
func (u *UtilitiesWiFi) TestSiteWlanSmsGlobal(
	ctx context.Context,
	body *models.TestSmsGlobal) (
	*http.Response,
	error) {
	req := u.prepareRequest(ctx, "POST", "/api/v1/utils/test_smsglobal")

	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
			NewAndAuth(
				NewAuth("basicAuth"),
				NewAuth("csrfToken"),
			),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}
	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}

// TestSiteWlanTelstraSetup takes context, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Allows validation of Telstra sms gateway credentials.
// In case of success, a text message confirming successful setup should be received. In case of error, telstra error message are returned.
func (u *UtilitiesWiFi) TestSiteWlanTelstraSetup(
	ctx context.Context,
	body *models.TestTelstra) (
	*http.Response,
	error) {
	req := u.prepareRequest(ctx, "POST", "/api/v1/utils/test_telstra")

	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
			NewAndAuth(
				NewAuth("basicAuth"),
				NewAuth("csrfToken"),
			),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}
	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}

// TestSiteWlanTwilioSetup takes context, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Allows validation of twilio setup
// In case of success, a text message confirming successful setup should be received. In case of error, twilio error code and message are returned.
func (u *UtilitiesWiFi) TestSiteWlanTwilioSetup(
	ctx context.Context,
	body *models.TestTwilio) (
	*http.Response,
	error) {
	req := u.prepareRequest(ctx, "POST", "/api/v1/utils/test_twilio")

	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
			NewAndAuth(
				NewAuth("basicAuth"),
				NewAuth("csrfToken"),
			),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}
	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}
