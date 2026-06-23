// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
	"context"
	"github.com/apimatic/go-core-runtime/https"
	"github.com/google/uuid"
	"github.com/tmunzer/mistapi-go/mistapi/errors"
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

// OrgsCRL represents a controller struct.
type OrgsCRL struct {
	baseController
}

// NewOrgsCRL creates a new instance of OrgsCRL.
// It takes a baseController as a parameter and returns a pointer to the OrgsCRL.
func NewOrgsCRL(baseController baseController) *OrgsCRL {
	orgsCRL := OrgsCRL{baseController: baseController}
	return &orgsCRL
}

// GetOrgCrlFile takes context, orgId as parameters and
// returns an models.ApiResponse with []byte data and
// an error if there was an issue with the request or response.
// Download the organization certificate revocation list (CRL) file used to identify certificates that should no longer be trusted.
func (o *OrgsCRL) GetOrgCrlFile(
	ctx context.Context,
	orgId uuid.UUID) (
	models.ApiResponse[[]byte],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/crl")
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

	stream, resp, err := req.CallAsStream()
	if err != nil {
		return models.NewApiResponse(stream, resp), err
	}
	return models.NewApiResponse(stream, resp), err
}
