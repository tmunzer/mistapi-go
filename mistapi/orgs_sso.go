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

// OrgsSSO represents a controller struct.
type OrgsSSO struct {
	baseController
}

// NewOrgsSSO creates a new instance of OrgsSSO.
// It takes a baseController as a parameter and returns a pointer to the OrgsSSO.
func NewOrgsSSO(baseController baseController) *OrgsSSO {
	orgsSSO := OrgsSSO{baseController: baseController}
	return &orgsSSO
}

// ListOrgSsos takes context, orgId, limit, page as parameters and
// returns an models.ApiResponse with []models.Sso data and
// an error if there was an issue with the request or response.
// List SSO identity provider configurations defined for this organization.
func (o *OrgsSSO) ListOrgSsos(
	ctx context.Context,
	orgId uuid.UUID,
	limit *int,
	page *int) (
	models.ApiResponse[[]models.Sso],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/ssos")
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
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if page != nil {
		req.QueryParam("page", *page)
	}

	var result []models.Sso
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.Sso](decoder)
	return models.NewApiResponse(result, resp), err
}

// CreateOrgSso takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.Sso data and
// an error if there was an issue with the request or response.
// Create an organization SSO identity provider configuration, including provider settings and role-handling behavior.
func (o *OrgsSSO) CreateOrgSso(
	ctx context.Context,
	orgId uuid.UUID,
	body *models.Sso) (
	models.ApiResponse[models.Sso],
	error) {
	req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/ssos")
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

	var result models.Sso
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Sso](decoder)
	return models.NewApiResponse(result, resp), err
}

// DeleteOrgSso takes context, orgId, ssoId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete an organization SSO identity provider configuration so it can no longer be used for administrator login.
func (o *OrgsSSO) DeleteOrgSso(
	ctx context.Context,
	orgId uuid.UUID,
	ssoId uuid.UUID) (
	*http.Response,
	error) {
	req := o.prepareRequest(ctx, "DELETE", "/api/v1/orgs/%v/ssos/%v")
	req.AppendTemplateParams(orgId, ssoId)
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

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}

// GetOrgSso takes context, orgId, ssoId as parameters and
// returns an models.ApiResponse with models.Sso data and
// an error if there was an issue with the request or response.
// Return one organization SSO identity provider configuration, including provider settings and generated SSO URLs.
func (o *OrgsSSO) GetOrgSso(
	ctx context.Context,
	orgId uuid.UUID,
	ssoId uuid.UUID) (
	models.ApiResponse[models.Sso],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/ssos/%v")
	req.AppendTemplateParams(orgId, ssoId)
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

	var result models.Sso
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Sso](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpdateOrgSso takes context, orgId, ssoId, body as parameters and
// returns an models.ApiResponse with models.Sso data and
// an error if there was an issue with the request or response.
// Update an organization SSO identity provider configuration, such as IdP URLs, certificates, issuer, NameID format, and unmatched-role handling.
func (o *OrgsSSO) UpdateOrgSso(
	ctx context.Context,
	orgId uuid.UUID,
	ssoId uuid.UUID,
	body *models.Sso) (
	models.ApiResponse[models.Sso],
	error) {
	req := o.prepareRequest(ctx, "PUT", "/api/v1/orgs/%v/ssos/%v")
	req.AppendTemplateParams(orgId, ssoId)
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

	var result models.Sso
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Sso](decoder)
	return models.NewApiResponse(result, resp), err
}

// DeleteOrgSsoAdmins takes context, orgId, ssoId, body as parameters and
// returns an models.ApiResponse with models.SsoDeleteAdminsResponse data and
// an error if there was an issue with the request or response.
// Remove SSO-linked organization administrator accounts by email for this SSO profile.
func (o *OrgsSSO) DeleteOrgSsoAdmins(
	ctx context.Context,
	orgId uuid.UUID,
	ssoId uuid.UUID,
	body *models.SsoDeleteAdmins) (
	models.ApiResponse[models.SsoDeleteAdminsResponse],
	error) {
	req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/ssos/%v/delete_admins")
	req.AppendTemplateParams(orgId, ssoId)
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

	var result models.SsoDeleteAdminsResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.SsoDeleteAdminsResponse](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListOrgSsoLatestFailures takes context, orgId, ssoId, start, end, duration, limit, page as parameters and
// returns an models.ApiResponse with models.ResponseSsoFailureSearch data and
// an error if there was an issue with the request or response.
// List recent authentication failures for this organization SSO configuration, including failure details and captured SAML assertion data when available.
func (o *OrgsSSO) ListOrgSsoLatestFailures(
	ctx context.Context,
	orgId uuid.UUID,
	ssoId uuid.UUID,
	start *string,
	end *string,
	duration *string,
	limit *int,
	page *int) (
	models.ApiResponse[models.ResponseSsoFailureSearch],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/ssos/%v/failures")
	req.AppendTemplateParams(orgId, ssoId)
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
	if start != nil {
		req.QueryParam("start", *start)
	}
	if end != nil {
		req.QueryParam("end", *end)
	}
	if duration != nil {
		req.QueryParam("duration", *duration)
	}
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if page != nil {
		req.QueryParam("page", *page)
	}

	var result models.ResponseSsoFailureSearch
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseSsoFailureSearch](decoder)
	return models.NewApiResponse(result, resp), err
}

// GetOrgSamlMetadata takes context, orgId, ssoId as parameters and
// returns an models.ApiResponse with models.SamlMetadata data and
// an error if there was an issue with the request or response.
// Return generated SAML service provider metadata for this organization SSO configuration as JSON.
func (o *OrgsSSO) GetOrgSamlMetadata(
	ctx context.Context,
	orgId uuid.UUID,
	ssoId uuid.UUID) (
	models.ApiResponse[models.SamlMetadata],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/ssos/%v/metadata")
	req.AppendTemplateParams(orgId, ssoId)
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

	var result models.SamlMetadata
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.SamlMetadata](decoder)
	return models.NewApiResponse(result, resp), err
}

// DownloadOrgSamlMetadata takes context, orgId, ssoId as parameters and
// returns an models.ApiResponse with []byte data and
// an error if there was an issue with the request or response.
// Download generated SAML service provider metadata XML for this
// organization SSO configuration. Use this XML to configure the identity provider
// with Mist service-provider details such as entity ID, ACS URL, logout URL,
// NameID format, and requested attributes.
// Example of metadata.xml:
// ```xml
// <?xml version=\"1.0\" encoding=\"UTF-8\"?><md:EntityDescriptor xmlns:md=\"urn:oasis:names:tc:SAML:2.0:metadata\" entityID=\"https://api.mist.com/api/v1/saml/5hdF5g/login\"\ validUntil=\"2027-10-12T21:59:01Z\" xmlns:ds=\"http://www.w3.org/2000/09/xmldsig#\">
// <md:SPSSODescriptor AuthnRequestsSigned=\"false\" WantAssertionsSigned=\"true\" protocolSupportEnumeration=\"urn:oasis:names:tc:SAML:2.0:protocol\">
// <md:SingleLogoutService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\"\ Location=\"https://api.mist.com/api/v1/saml/5hdF5g/logout\" />
// <md:NameIDFormat>urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified</md:NameIDFormat>
// <md:AssertionConsumerService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\"\ Location=\"https://api.mist.com/api/v1/saml/5hdF5g/login\" index=\"0\" isDefault=\"true\"/>
// <md:AttributeConsumingService index=\"0\">
// <md:ServiceName xml:lang=\"en-US\">Mist</md:ServiceName>
// <md:RequestedAttribute\ Name=\"Role\" NameFormat=\"urn:oasis:names:tc:SAML:2.0:attrname-format:basic\"\ isRequired=\"true\"/>
// <md:RequestedAttribute Name=\"FirstName\"\ NameFormat=\"urn:oasis:names:tc:SAML:2.0:attrname-format:basic\" isRequired=\"false\"/>
// <md:RequestedAttribute Name=\"LastName\" NameFormat=\"urn:oasis:names:tc:SAML:2.0:attrname-format:basic\" isRequired=\"false\"/>
// </md:AttributeConsumingService>
// </md:SPSSODescriptor>
// </md:EntityDescriptor>
// ```
func (o *OrgsSSO) DownloadOrgSamlMetadata(
	ctx context.Context,
	orgId uuid.UUID,
	ssoId uuid.UUID) (
	models.ApiResponse[[]byte],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/ssos/%v/metadata.xml")
	req.AppendTemplateParams(orgId, ssoId)
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

	stream, resp, err := req.CallAsStream()
	if err != nil {
		return models.NewApiResponse(stream, resp), err
	}
	return models.NewApiResponse(stream, resp), err
}
