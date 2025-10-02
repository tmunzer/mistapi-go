// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
	"context"
	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
	"github.com/google/uuid"
	"github.com/tmunzer/mistapi-go/mistapi/errors"
	"github.com/tmunzer/mistapi-go/mistapi/models"
	"net/http"
)

// MSPs represents a controller struct.
type MSPs struct {
	baseController
}

// NewMSPs creates a new instance of MSPs.
// It takes a baseController as a parameter and returns a pointer to the MSPs.
func NewMSPs(baseController baseController) *MSPs {
	mSPs := MSPs{baseController: baseController}
	return &mSPs
}

// CreateMsp takes context, body as parameters and
// returns an models.ApiResponse with models.Msp data and
// an error if there was an issue with the request or response.
// Create MSP account
func (m *MSPs) CreateMsp(
	ctx context.Context,
	body *models.Msp) (
	models.ApiResponse[models.Msp],
	error) {
	req := m.prepareRequest(ctx, "POST", "/api/v1/msps")

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
	var result models.Msp
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Msp](decoder)
	return models.NewApiResponse(result, resp), err
}

// DeleteMsp takes context, mspId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Deleting MSP removes the MSP and OrgGroup under the MSP as well as all privileges associated with them. It does not remove any Org or Admins
func (m *MSPs) DeleteMsp(
	ctx context.Context,
	mspId uuid.UUID) (
	*http.Response,
	error) {
	req := m.prepareRequest(ctx, "DELETE", "/api/v1/msps/%v")
	req.AppendTemplateParams(mspId)
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

// GetMspDetails takes context, mspId as parameters and
// returns an models.ApiResponse with models.Msp data and
// an error if there was an issue with the request or response.
// Get MSP Detail
func (m *MSPs) GetMspDetails(
	ctx context.Context,
	mspId uuid.UUID) (
	models.ApiResponse[models.Msp],
	error) {
	req := m.prepareRequest(ctx, "GET", "/api/v1/msps/%v")
	req.AppendTemplateParams(mspId)
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

	var result models.Msp
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Msp](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpdateMsp takes context, mspId, body as parameters and
// returns an models.ApiResponse with models.Msp data and
// an error if there was an issue with the request or response.
// Update MSP
func (m *MSPs) UpdateMsp(
	ctx context.Context,
	mspId uuid.UUID,
	body *models.Msp) (
	models.ApiResponse[models.Msp],
	error) {
	req := m.prepareRequest(ctx, "PUT", "/api/v1/msps/%v")
	req.AppendTemplateParams(mspId)
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

	var result models.Msp
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Msp](decoder)
	return models.NewApiResponse(result, resp), err
}

// SearchMspOrgGroup takes context, mspId, mType, q, limit, start, end, duration, sort as parameters and
// returns an models.ApiResponse with models.ResponseSearch data and
// an error if there was an issue with the request or response.
// Search in MSP Orgs
func (m *MSPs) SearchMspOrgGroup(
	ctx context.Context,
	mspId uuid.UUID,
	mType models.MspSearchTypeEnum,
	q *string,
	limit *int,
	start *string,
	end *string,
	duration *string,
	sort *string) (
	models.ApiResponse[models.ResponseSearch],
	error) {
	req := m.prepareRequest(ctx, "GET", "/api/v1/msps/%v/search")
	req.AppendTemplateParams(mspId)
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
	req.QueryParam("type", mType)
	if q != nil {
		req.QueryParam("q", *q)
	}
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if start != nil {
		req.QueryParam("start", *start)
	}
	if end != nil {
		req.QueryParam("end", *end)
	}
	if duration != nil {
		req.QueryParam("duration", *duration)
	}
	if sort != nil {
		req.QueryParam("sort", *sort)
	}

	var result models.ResponseSearch
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseSearch](decoder)
	return models.NewApiResponse(result, resp), err
}
