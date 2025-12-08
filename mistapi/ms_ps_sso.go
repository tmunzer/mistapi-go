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

// MSPsSSO represents a controller struct.
type MSPsSSO struct {
	baseController
}

// NewMSPsSSO creates a new instance of MSPsSSO.
// It takes a baseController as a parameter and returns a pointer to the MSPsSSO.
func NewMSPsSSO(baseController baseController) *MSPsSSO {
	mSPsSSO := MSPsSSO{baseController: baseController}
	return &mSPsSSO
}

// ListMspSsos takes context, mspId as parameters and
// returns an models.ApiResponse with []models.Sso data and
// an error if there was an issue with the request or response.
// List MSP SSO Configs
func (m *MSPsSSO) ListMspSsos(
	ctx context.Context,
	mspId uuid.UUID) (
	models.ApiResponse[[]models.Sso],
	error) {
	req := m.prepareRequest(ctx, "GET", "/api/v1/msps/%v/ssos")
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

	var result []models.Sso
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.Sso](decoder)
	return models.NewApiResponse(result, resp), err
}

// CreateMspSso takes context, mspId, body as parameters and
// returns an models.ApiResponse with models.Sso data and
// an error if there was an issue with the request or response.
// Create MSP SSO profile
func (m *MSPsSSO) CreateMspSso(
	ctx context.Context,
	mspId uuid.UUID,
	body *models.Sso) (
	models.ApiResponse[models.Sso],
	error) {
	req := m.prepareRequest(ctx, "POST", "/api/v1/msps/%v/ssos")
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

	var result models.Sso
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Sso](decoder)
	return models.NewApiResponse(result, resp), err
}

// DeleteMspSso takes context, mspId, ssoId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete MSP SSO Config
func (m *MSPsSSO) DeleteMspSso(
	ctx context.Context,
	mspId uuid.UUID,
	ssoId uuid.UUID) (
	*http.Response,
	error) {
	req := m.prepareRequest(ctx, "DELETE", "/api/v1/msps/%v/ssos/%v")
	req.AppendTemplateParams(mspId, ssoId)
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

// GetMspSso takes context, mspId, ssoId as parameters and
// returns an models.ApiResponse with models.Sso data and
// an error if there was an issue with the request or response.
// Get MSP SSO Config
func (m *MSPsSSO) GetMspSso(
	ctx context.Context,
	mspId uuid.UUID,
	ssoId uuid.UUID) (
	models.ApiResponse[models.Sso],
	error) {
	req := m.prepareRequest(ctx, "GET", "/api/v1/msps/%v/ssos/%v")
	req.AppendTemplateParams(mspId, ssoId)
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

	var result models.Sso
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Sso](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpdateMspSso takes context, mspId, ssoId, body as parameters and
// returns an models.ApiResponse with models.Sso data and
// an error if there was an issue with the request or response.
// Update MSP SSO config
func (m *MSPsSSO) UpdateMspSso(
	ctx context.Context,
	mspId uuid.UUID,
	ssoId uuid.UUID,
	body *models.Sso) (
	models.ApiResponse[models.Sso],
	error) {
	req := m.prepareRequest(ctx, "PUT", "/api/v1/msps/%v/ssos/%v")
	req.AppendTemplateParams(mspId, ssoId)
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

	var result models.Sso
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Sso](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListMspSsoLatestFailures takes context, mspId, ssoId as parameters and
// returns an models.ApiResponse with models.ResponseSsoFailureSearch data and
// an error if there was an issue with the request or response.
// Get List of MSP SSO Latest Failures
func (m *MSPsSSO) ListMspSsoLatestFailures(
	ctx context.Context,
	mspId uuid.UUID,
	ssoId uuid.UUID) (
	models.ApiResponse[models.ResponseSsoFailureSearch],
	error) {
	req := m.prepareRequest(ctx, "GET", "/api/v1/msps/%v/ssos/%v/failures")
	req.AppendTemplateParams(mspId, ssoId)
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

	var result models.ResponseSsoFailureSearch
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseSsoFailureSearch](decoder)
	return models.NewApiResponse(result, resp), err
}

// GetMspSamlMetadata takes context, mspId, ssoId as parameters and
// returns an models.ApiResponse with models.SamlMetadata data and
// an error if there was an issue with the request or response.
// Get MSP SAML Metadata
func (m *MSPsSSO) GetMspSamlMetadata(
	ctx context.Context,
	mspId uuid.UUID,
	ssoId uuid.UUID) (
	models.ApiResponse[models.SamlMetadata],
	error) {
	req := m.prepareRequest(ctx, "GET", "/api/v1/msps/%v/ssos/%v/metadata")
	req.AppendTemplateParams(mspId, ssoId)
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

	var result models.SamlMetadata
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.SamlMetadata](decoder)
	return models.NewApiResponse(result, resp), err
}

// DownloadMspSamlMetadata takes context, mspId, ssoId as parameters and
// returns an models.ApiResponse with string data and
// an error if there was an issue with the request or response.
// Download MSP SAML Metadata
// Example of metadata.xml:
// ```xml
// <?xml version="1.0" encoding="UTF-8"?><md:EntityDescriptor xmlns:md="urn:oasis:names:tc:SAML:2.0:metadata" entityID="https://api.mist.com/api/v1/saml/5hdF5g/login" validUntil="2027-10-12T21:59:01Z" xmlns:ds="http://www.w3.org/2000/09/xmldsig#">
// <md:SPSSODescriptor AuthnRequestsSigned="false" WantAssertionsSigned="true" protocolSupportEnumeration="urn:oasis:names:tc:SAML:2.0:protocol">
// <md:SingleLogoutService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST" Location="https://api.mist.com/api/v1/saml/5hdF5g/logout" />
// <md:NameIDFormat>urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified</md:NameIDFormat>
// <md:AssertionConsumerService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST" Location="https://api.mist.com/api/v1/saml/5hdF5g/login" index="0" isDefault="true"/>
// <md:AttributeConsumingService index="0">
// <md:ServiceName xml:lang="en-US">Mist</md:ServiceName>
// <md:RequestedAttribute Name="Role" NameFormat="urn:oasis:names:tc:SAML:2.0:attrname-format:basic" isRequired="true"/>
// <md:RequestedAttribute Name="FirstName" NameFormat="urn:oasis:names:tc:SAML:2.0:attrname-format:basic" isRequired="false"/>
// <md:RequestedAttribute Name="LastName" NameFormat="urn:oasis:names:tc:SAML:2.0:attrname-format:basic" isRequired="false"/>
// </md:AttributeConsumingService>
// </md:SPSSODescriptor>
// </md:EntityDescriptor>
// ```
func (m *MSPsSSO) DownloadMspSamlMetadata(
	ctx context.Context,
	mspId uuid.UUID,
	ssoId uuid.UUID) (
	models.ApiResponse[string],
	error) {
	req := m.prepareRequest(ctx, "GET", "/api/v1/msps/%v/ssos/%v/metadata.xml")
	req.AppendTemplateParams(mspId, ssoId)
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

	str, resp, err := req.CallAsText()
	var result string = str

	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	return models.NewApiResponse(result, resp), err
}
