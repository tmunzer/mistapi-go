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
)

// OrgsDevicesEdgeConnect represents a controller struct.
type OrgsDevicesEdgeConnect struct {
	baseController
}

// NewOrgsDevicesEdgeConnect creates a new instance of OrgsDevicesEdgeConnect.
// It takes a baseController as a parameter and returns a pointer to the OrgsDevicesEdgeConnect.
func NewOrgsDevicesEdgeConnect(baseController baseController) *OrgsDevicesEdgeConnect {
	orgsDevicesEdgeConnect := OrgsDevicesEdgeConnect{baseController: baseController}
	return &orgsDevicesEdgeConnect
}

// GetOrgEdgeconnectRegisterCmd takes context, orgId as parameters and
// returns an models.ApiResponse with models.EdgeconnectRegisterCmd data and
// an error if there was an issue with the request or response.
// Returns a registration code for adopting an EdgeConnect device into Mist.
func (o *OrgsDevicesEdgeConnect) GetOrgEdgeconnectRegisterCmd(
	ctx context.Context,
	orgId uuid.UUID) (
	models.ApiResponse[models.EdgeconnectRegisterCmd],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/edgeconnect/register_cmd")
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

	var result models.EdgeconnectRegisterCmd
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.EdgeconnectRegisterCmd](decoder)
	return models.NewApiResponse(result, resp), err
}
