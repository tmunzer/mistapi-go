package mistapi

import (
	"context"
	"fmt"
	"net/http"

	"github.com/tmunzer/mistapi-go/mistapi/errors"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/apimatic/go-core-runtime/https"
	"github.com/google/uuid"
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
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Trigger a CoA (change of authorization) against a client
func (u *UtilitiesWiFi) ReauthOrgDot1xWirelessClient(
	ctx context.Context,
	orgId uuid.UUID,
	clientMac string) (
	*http.Response,
	error) {
	req := u.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/api/v1/orgs/%v/clients/%v/coa", orgId, clientMac),
	)
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
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
	})

	context, err := req.Call()
	if err != nil {
		return context.Response, err
	}
	return context.Response, err
}

// TestSiteWlanTelstraSetup takes context, body as parameters and
// returns an models.ApiResponse with  data and
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
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}
	context, err := req.Call()
	if err != nil {
		return context.Response, err
	}
	return context.Response, err
}

// TestSiteWlanTwilioSetup takes context, body as parameters and
// returns an models.ApiResponse with  data and
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
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}
	context, err := req.Call()
	if err != nil {
		return context.Response, err
	}
	return context.Response, err
}

// DisconnectSiteMultipleClients takes context, siteId, body as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// To unauthorize multiple clients
func (u *UtilitiesWiFi) DisconnectSiteMultipleClients(
	ctx context.Context,
	siteId uuid.UUID,
	body *models.MacAddresses) (
	*http.Response,
	error) {
	req := u.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/api/v1/sites/%v/clients/disconnect", siteId),
	)
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
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	context, err := req.Call()
	if err != nil {
		return context.Response, err
	}
	return context.Response, err
}

// UnauthorizeSiteMultipleClients takes context, siteId, body as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// This unauthorize clients (if they are guest) and disconnect them. From the guest’s perspective, they will see the splash page again and go through the flow (e.g. Terms of Use) again.
func (u *UtilitiesWiFi) UnauthorizeSiteMultipleClients(
	ctx context.Context,
	siteId uuid.UUID,
	body *models.MacAddresses) (
	*http.Response,
	error) {
	req := u.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/api/v1/sites/%v/clients/unauthorize", siteId),
	)
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
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	context, err := req.Call()
	if err != nil {
		return context.Response, err
	}
	return context.Response, err
}

// ReauthSiteDot1xWirelessClient takes context, siteId, clientMac as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Trigger a CoA (change of authorization) against a client
func (u *UtilitiesWiFi) ReauthSiteDot1xWirelessClient(
	ctx context.Context,
	siteId uuid.UUID,
	clientMac string) (
	*http.Response,
	error) {
	req := u.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/api/v1/sites/%v/clients/%v/coa", siteId, clientMac),
	)
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
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
	})

	context, err := req.Call()
	if err != nil {
		return context.Response, err
	}
	return context.Response, err
}

// DisconnectSiteWirelessClient takes context, siteId, clientMac as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// This disconnect a client (and it’s likely to connect back)
func (u *UtilitiesWiFi) DisconnectSiteWirelessClient(
	ctx context.Context,
	siteId uuid.UUID,
	clientMac string) (
	*http.Response,
	error) {
	req := u.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/api/v1/sites/%v/clients/%v/disconnect", siteId, clientMac),
	)
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
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
	})

	context, err := req.Call()
	if err != nil {
		return context.Response, err
	}
	return context.Response, err
}

// UnauthorizeSiteWirelessClient takes context, siteId, clientMac as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// This unauthorize a client (if it’s a guest) and disconnect it. From the guest’s perspective, s/he will see the splash page again and go through the flow (e.g. Terms of Use) again.
func (u *UtilitiesWiFi) UnauthorizeSiteWirelessClient(
	ctx context.Context,
	siteId uuid.UUID,
	clientMac string) (
	*http.Response,
	error) {
	req := u.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/api/v1/sites/%v/clients/%v/unauthorize", siteId, clientMac),
	)
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
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
	})

	context, err := req.Call()
	if err != nil {
		return context.Response, err
	}
	return context.Response, err
}

// ReprovisionSiteAllAps takes context, siteId as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// To force all APs to reprovision itself again.
func (u *UtilitiesWiFi) ReprovisionSiteAllAps(
	ctx context.Context,
	siteId uuid.UUID) (
	*http.Response,
	error) {
	req := u.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/api/v1/sites/%v/devices/reprovision", siteId),
	)
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
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
	})

	context, err := req.Call()
	if err != nil {
		return context.Response, err
	}
	return context.Response, err
}

// ResetSiteAllApsToUseRrm takes context, siteId, body as parameters and
// returns an models.ApiResponse with  data and
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
		fmt.Sprintf("/api/v1/sites/%v/devices/reset_radio_config", siteId),
	)
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
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	context, err := req.Call()
	if err != nil {
		return context.Response, err
	}
	return context.Response, err
}

// ZeroizeSiteFipsAllAps takes context, siteId, body as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Zeroize all FIPS APs in the Site
func (u *UtilitiesWiFi) ZeroizeSiteFipsAllAps(
	ctx context.Context,
	siteId uuid.UUID,
	body *models.UtilsZeroiseFips) (
	*http.Response,
	error) {
	req := u.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/api/v1/sites/%v/devices/zeroize", siteId),
	)
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
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	context, err := req.Call()
	if err != nil {
		return context.Response, err
	}
	return context.Response, err
}

// DeauthSiteWirelessClientsConnectedToARogue takes context, siteId, rogueBssid as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Send Deauth frame to clients connected to a Rogue AP
func (u *UtilitiesWiFi) DeauthSiteWirelessClientsConnectedToARogue(
	ctx context.Context,
	siteId uuid.UUID,
	rogueBssid string) (
	*http.Response,
	error) {
	req := u.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/api/v1/sites/%v/rogues/%v/deauth_clients", siteId, rogueBssid),
	)
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
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
	})

	context, err := req.Call()
	if err != nil {
		return context.Response, err
	}
	return context.Response, err
}

// OptimizeSiteRrm takes context, siteId, body as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Optimize Site RRM
func (u *UtilitiesWiFi) OptimizeSiteRrm(
	ctx context.Context,
	siteId uuid.UUID,
	body *models.UtilsRrmOptimize) (
	*http.Response,
	error) {
	req := u.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/api/v1/sites/%v/rrm/optimize", siteId),
	)
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
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	context, err := req.Call()
	if err != nil {
		return context.Response, err
	}
	return context.Response, err
}
