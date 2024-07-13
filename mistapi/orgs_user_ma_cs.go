package mistapi

import (
	"context"
	"fmt"
	"net/http"

	"github.com/tmunzer/mistapi-go/mistapi/errors"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
	"github.com/google/uuid"
)

// OrgsUserMACs represents a controller struct.
type OrgsUserMACs struct {
	baseController
}

// NewOrgsUserMACs creates a new instance of OrgsUserMACs.
// It takes a baseController as a parameter and returns a pointer to the OrgsUserMACs.
func NewOrgsUserMACs(baseController baseController) *OrgsUserMACs {
	orgsUserMACs := OrgsUserMACs{baseController: baseController}
	return &orgsUserMACs
}

// ListOrgUserMacs takes context, orgId, blacklisted, forGuestWifi, crossSite, siteId, page, limit as parameters and
// returns an models.ApiResponse with []models.UserMac data and
// an error if there was an issue with the request or response.
// List Org User MACs
func (o *OrgsUserMACs) ListOrgUserMacs(
	ctx context.Context,
	orgId uuid.UUID,
	blacklisted *bool,
	forGuestWifi *bool,
	crossSite *bool,
	siteId *string,
	page *int,
	limit *int) (
	models.ApiResponse[[]models.UserMac],
	error) {
	req := o.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/orgs/%v/usermacs", orgId),
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
	if blacklisted != nil {
		req.QueryParam("blacklisted", *blacklisted)
	}
	if forGuestWifi != nil {
		req.QueryParam("for_guest_wifi", *forGuestWifi)
	}
	if crossSite != nil {
		req.QueryParam("cross_site", *crossSite)
	}
	if siteId != nil {
		req.QueryParam("site_id", *siteId)
	}
	if page != nil {
		req.QueryParam("page", *page)
	}
	if limit != nil {
		req.QueryParam("limit", *limit)
	}

	var result []models.UserMac
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.UserMac](decoder)
	return models.NewApiResponse(result, resp), err
}

// CreateOrgUserMacs takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.UserMacImport data and
// an error if there was an issue with the request or response.
// Create Org User MACs
// ### Usermacs import CSV file format
// mac,labels,vlan,notes
// 921b638445cd,”bldg1,flor1”,vlan-100
// 721b638445ef,”bldg2,flor2”,vlan-101,Canon Printers
// 721b638445ee,”bldg3,flor3”,vlan-102
// 921b638445ce,”bldg4,flor4”,vlan-103
// 921b638445cf,”bldg5,flor5”,vlan-104
func (o *OrgsUserMACs) CreateOrgUserMacs(
	ctx context.Context,
	orgId uuid.UUID,
	body *models.UserMac) (
	models.ApiResponse[models.UserMacImport],
	error) {
	req := o.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/api/v1/orgs/%v/usermacs", orgId),
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

	var result models.UserMacImport
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.UserMacImport](decoder)
	return models.NewApiResponse(result, resp), err
}

// ImportOrgUserMacs takes context, orgId, body as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Import Org User MACs
func (o *OrgsUserMACs) ImportOrgUserMacs(
	ctx context.Context,
	orgId uuid.UUID,
	body []models.UserMac) (
	*http.Response,
	error) {
	req := o.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/api/v1/orgs/%v/usermacs/import", orgId),
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

// SearchOrgUserMacs takes context, orgId, mac, labels, page, limit as parameters and
// returns an models.ApiResponse with models.UserMacsSearch data and
// an error if there was an issue with the request or response.
// Search Org User MACs
func (o *OrgsUserMACs) SearchOrgUserMacs(
	ctx context.Context,
	orgId uuid.UUID,
	mac *string,
	labels []string,
	page *int,
	limit *int) (
	models.ApiResponse[models.UserMacsSearch],
	error) {
	req := o.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/orgs/%v/usermacs/search", orgId),
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
	if mac != nil {
		req.QueryParam("mac", *mac)
	}
	if labels != nil {
		req.QueryParam("labels", labels)
	}
	if page != nil {
		req.QueryParam("page", *page)
	}
	if limit != nil {
		req.QueryParam("limit", *limit)
	}

	var result models.UserMacsSearch
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.UserMacsSearch](decoder)
	return models.NewApiResponse(result, resp), err
}

// DeleteOrgUserMac takes context, orgId, usermacId as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Delete Org User MAC
func (o *OrgsUserMACs) DeleteOrgUserMac(
	ctx context.Context,
	orgId uuid.UUID,
	usermacId uuid.UUID) (
	*http.Response,
	error) {
	req := o.prepareRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/api/v1/orgs/%v/usermacs/search/%v", orgId, usermacId),
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

// GetOrgUserMac takes context, orgId, usermacId as parameters and
// returns an models.ApiResponse with models.UserMac data and
// an error if there was an issue with the request or response.
// Get Org User MAC
func (o *OrgsUserMACs) GetOrgUserMac(
	ctx context.Context,
	orgId uuid.UUID,
	usermacId uuid.UUID) (
	models.ApiResponse[models.UserMac],
	error) {
	req := o.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/orgs/%v/usermacs/search/%v", orgId, usermacId),
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

	var result models.UserMac
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.UserMac](decoder)
	return models.NewApiResponse(result, resp), err
}
