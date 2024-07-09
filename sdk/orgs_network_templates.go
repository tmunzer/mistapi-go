package mistapigo

import (
	"context"
	"fmt"
	"net/http"

	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
	"github.com/google/uuid"
	"github.com/tmunzer/mistapigo/sdk/errors"
	"github.com/tmunzer/mistapigo/sdk/models"
)

// OrgsNetworkTemplates represents a controller struct.
type OrgsNetworkTemplates struct {
	baseController
}

// NewOrgsNetworkTemplates creates a new instance of OrgsNetworkTemplates.
// It takes a baseController as a parameter and returns a pointer to the OrgsNetworkTemplates.
func NewOrgsNetworkTemplates(baseController baseController) *OrgsNetworkTemplates {
	orgsNetworkTemplates := OrgsNetworkTemplates{baseController: baseController}
	return &orgsNetworkTemplates
}

// ListOrgNetworkTemplates takes context, orgId, page, limit as parameters and
// returns an models.ApiResponse with []models.NetworkTemplate data and
// an error if there was an issue with the request or response.
// Get List of Org Network Templates
func (o *OrgsNetworkTemplates) ListOrgNetworkTemplates(
	ctx context.Context,
	orgId uuid.UUID,
	page *int,
	limit *int) (
	models.ApiResponse[[]models.NetworkTemplate],
	error) {
	req := o.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/orgs/%v/networktemplates", orgId),
	)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
	})
	if page != nil {
		req.QueryParam("page", *page)
	}
	if limit != nil {
		req.QueryParam("limit", *limit)
	}

	var result []models.NetworkTemplate
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.NetworkTemplate](decoder)
	return models.NewApiResponse(result, resp), err
}

// CreateOrgNetworkTemplate takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.NetworkTemplate data and
// an error if there was an issue with the request or response.
// Update Org Network Templates
func (o *OrgsNetworkTemplates) CreateOrgNetworkTemplate(
	ctx context.Context,
	orgId uuid.UUID,
	body *models.NetworkTemplate) (
	models.ApiResponse[models.NetworkTemplate],
	error) {
	req := o.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/api/v1/orgs/%v/networktemplates", orgId),
	)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
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

	var result models.NetworkTemplate
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.NetworkTemplate](decoder)
	return models.NewApiResponse(result, resp), err
}

// DeleteOrgNetworkTemplate takes context, orgId, networktemplateId as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Delete Org Network Template
func (o *OrgsNetworkTemplates) DeleteOrgNetworkTemplate(
	ctx context.Context,
	orgId uuid.UUID,
	networktemplateId uuid.UUID) (
	*http.Response,
	error) {
	req := o.prepareRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/api/v1/orgs/%v/networktemplates/%v", orgId, networktemplateId),
	)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
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

// GetOrgNetworkTemplate takes context, orgId, networktemplateId as parameters and
// returns an models.ApiResponse with models.NetworkTemplate data and
// an error if there was an issue with the request or response.
// Get Org Network Templates Details
func (o *OrgsNetworkTemplates) GetOrgNetworkTemplate(
	ctx context.Context,
	orgId uuid.UUID,
	networktemplateId uuid.UUID) (
	models.ApiResponse[models.NetworkTemplate],
	error) {
	req := o.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/orgs/%v/networktemplates/%v", orgId, networktemplateId),
	)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
	})

	var result models.NetworkTemplate
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.NetworkTemplate](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpdateOrgNetworkTemplates takes context, orgId, networktemplateId, body as parameters and
// returns an models.ApiResponse with models.NetworkTemplate data and
// an error if there was an issue with the request or response.
// Update Org Network Template
func (o *OrgsNetworkTemplates) UpdateOrgNetworkTemplates(
	ctx context.Context,
	orgId uuid.UUID,
	networktemplateId uuid.UUID,
	body *models.NetworkTemplate) (
	models.ApiResponse[models.NetworkTemplate],
	error) {
	req := o.prepareRequest(
		ctx,
		"PUT",
		fmt.Sprintf("/api/v1/orgs/%v/networktemplates/%v", orgId, networktemplateId),
	)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
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

	var result models.NetworkTemplate
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.NetworkTemplate](decoder)
	return models.NewApiResponse(result, resp), err
}
