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
	req := m.prepareRequest(ctx, "GET", fmt.Sprintf("/api/v1/msps/%v/ssos", mspId))
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
	req := m.prepareRequest(ctx, "POST", fmt.Sprintf("/api/v1/msps/%v/ssos", mspId))
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

	var result models.Sso
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Sso](decoder)
	return models.NewApiResponse(result, resp), err
}

// DeleteMspSso takes context, mspId, ssoId as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Delete MSP SSO Config
func (m *MSPsSSO) DeleteMspSso(
	ctx context.Context,
	mspId uuid.UUID,
	ssoId uuid.UUID) (
	*http.Response,
	error) {
	req := m.prepareRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/api/v1/msps/%v/ssos/%v", mspId, ssoId),
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
	req := m.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/msps/%v/ssos/%v", mspId, ssoId),
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
	req := m.prepareRequest(
		ctx,
		"PUT",
		fmt.Sprintf("/api/v1/msps/%v/ssos/%v", mspId, ssoId),
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
	req := m.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/msps/%v/ssos/%v/failures", mspId, ssoId),
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

	var result models.ResponseSsoFailureSearch
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseSsoFailureSearch](decoder)
	return models.NewApiResponse(result, resp), err
}

// GetMspSsoSamlMetadata takes context, mspId, ssoId as parameters and
// returns an models.ApiResponse with models.SsoSamlMetadata data and
// an error if there was an issue with the request or response.
// Get MSP SSO SAML Metadata
func (m *MSPsSSO) GetMspSsoSamlMetadata(
	ctx context.Context,
	mspId uuid.UUID,
	ssoId uuid.UUID) (
	models.ApiResponse[models.SsoSamlMetadata],
	error) {
	req := m.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/msps/%v/ssos/%v/metadata", mspId, ssoId),
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

	var result models.SsoSamlMetadata
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.SsoSamlMetadata](decoder)
	return models.NewApiResponse(result, resp), err
}

// DownloadMspSsoSamlMetadata takes context, mspId, ssoId as parameters and
// returns an models.ApiResponse with []byte data and
// an error if there was an issue with the request or response.
// Download MSP SSO SAML Metadata
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
func (m *MSPsSSO) DownloadMspSsoSamlMetadata(
	ctx context.Context,
	mspId uuid.UUID,
	ssoId uuid.UUID) (
	models.ApiResponse[[]byte],
	error) {
	req := m.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/msps/%v/ssos/%v/metadata.xml", mspId, ssoId),
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

	stream, resp, err := req.CallAsStream()
	if err != nil {
		return models.NewApiResponse(stream, resp), err
	}
	return models.NewApiResponse(stream, resp), err
}
