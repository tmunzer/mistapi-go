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

// OrgsSCEP represents a controller struct.
type OrgsSCEP struct {
	baseController
}

// NewOrgsSCEP creates a new instance of OrgsSCEP.
// It takes a baseController as a parameter and returns a pointer to the OrgsSCEP.
func NewOrgsSCEP(baseController baseController) *OrgsSCEP {
	orgsSCEP := OrgsSCEP{baseController: baseController}
	return &orgsSCEP
}

// DisableOrgMistScep takes context, orgId as parameters and
// returns an models.ApiResponse with models.OrgSettingScepResponse data and
// an error if there was an issue with the request or response.
// Disable Mist SCEP for the organization and return the updated read-only SCEP settings.
func (o *OrgsSCEP) DisableOrgMistScep(
	ctx context.Context,
	orgId uuid.UUID) (
	models.ApiResponse[models.OrgSettingScepResponse],
	error) {
	req := o.prepareRequest(ctx, "DELETE", "/api/v1/orgs/%v/setting/mist_scep")
	req.AppendTemplateParams(orgId)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("csrfToken"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429},
	})

	var result models.OrgSettingScepResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.OrgSettingScepResponse](decoder)
	return models.NewApiResponse(result, resp), err
}

// GetOrgMistScep takes context, orgId as parameters and
// returns an models.ApiResponse with models.OrgSettingScepResponse data and
// an error if there was an issue with the request or response.
// Return Mist SCEP settings for the organization, including enabled and suspended status, configured certificate providers, and generated enrollment or webhook URLs.
func (o *OrgsSCEP) GetOrgMistScep(
	ctx context.Context,
	orgId uuid.UUID) (
	models.ApiResponse[models.OrgSettingScepResponse],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/setting/mist_scep")
	req.AppendTemplateParams(orgId)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("csrfToken"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429},
	})

	var result models.OrgSettingScepResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.OrgSettingScepResponse](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpdateOrgMistScep takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.OrgSettingScepResponse data and
// an error if there was an issue with the request or response.
// Update Mist SCEP settings for the organization, including enabled state, suspension state, and certificate providers.
func (o *OrgsSCEP) UpdateOrgMistScep(
	ctx context.Context,
	orgId uuid.UUID,
	body *models.OrgSettingScep) (
	models.ApiResponse[models.OrgSettingScepResponse],
	error) {
	req := o.prepareRequest(ctx, "PUT", "/api/v1/orgs/%v/setting/mist_scep")
	req.AppendTemplateParams(orgId)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("csrfToken"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	var result models.OrgSettingScepResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.OrgSettingScepResponse](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListOrgIssuedClientCertificates takes context, orgId, commonName, certProvider, serialNumber, deviceId, expireTime, createdTime, limit, page as parameters and
// returns an models.ApiResponse with models.IssuedClientCertificatesResults data and
// an error if there was an issue with the request or response.
// List Mist SCEP client certificates issued for the organization. Results can be filtered by common name, certificate provider, serial number, device ID, or time range.
func (o *OrgsSCEP) ListOrgIssuedClientCertificates(
	ctx context.Context,
	orgId uuid.UUID,
	commonName *string,
	certProvider *string,
	serialNumber *string,
	deviceId *string,
	expireTime *int,
	createdTime *int,
	limit *int,
	page *int) (
	models.ApiResponse[models.IssuedClientCertificatesResults],
	error) {
	req := o.prepareRequest(
		ctx,
		"GET",
		"/api/v1/orgs/%v/setting/mist_scep/client_certs",
	)
	req.AppendTemplateParams(orgId)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("csrfToken"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429},
	})
	if commonName != nil {
		req.QueryParam("common_name", *commonName)
	}
	if certProvider != nil {
		req.QueryParam("cert_provider", *certProvider)
	}
	if serialNumber != nil {
		req.QueryParam("serial_number", *serialNumber)
	}
	if deviceId != nil {
		req.QueryParam("device_id", *deviceId)
	}
	if expireTime != nil {
		req.QueryParam("expire_time", *expireTime)
	}
	if createdTime != nil {
		req.QueryParam("created_time", *createdTime)
	}
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if page != nil {
		req.QueryParam("page", *page)
	}

	var result models.IssuedClientCertificatesResults
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.IssuedClientCertificatesResults](decoder)
	return models.NewApiResponse(result, resp), err
}

// RevokeOrgIssuedClientCertificates takes context, orgId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Revoke issued Mist SCEP client certificates by certificate serial number.
func (o *OrgsSCEP) RevokeOrgIssuedClientCertificates(
	ctx context.Context,
	orgId uuid.UUID,
	body *models.ClientCertSerialNumbers) (
	*http.Response,
	error) {
	req := o.prepareRequest(
		ctx,
		"POST",
		"/api/v1/orgs/%v/setting/mist_scep/client_certs/revoke",
	)
	req.AppendTemplateParams(orgId)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("csrfToken"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429},
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
