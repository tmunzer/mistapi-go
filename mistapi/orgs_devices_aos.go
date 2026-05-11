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

// OrgsDevicesAOS represents a controller struct.
type OrgsDevicesAOS struct {
	baseController
}

// NewOrgsDevicesAOS creates a new instance of OrgsDevicesAOS.
// It takes a baseController as a parameter and returns a pointer to the OrgsDevicesAOS.
func NewOrgsDevicesAOS(baseController baseController) *OrgsDevicesAOS {
	orgsDevicesAOS := OrgsDevicesAOS{baseController: baseController}
	return &orgsDevicesAOS
}

// GetOrgAosRegisterCmd takes context, orgId as parameters and
// returns an models.ApiResponse with models.AosRegisterCmd data and
// an error if there was an issue with the request or response.
// Generates a registration challenge token for TPM-based brownfield registration of AOS devices. The returned command string can be copied and pasted directly into an AOS device to register it with Mist.
func (o *OrgsDevicesAOS) GetOrgAosRegisterCmd(
	ctx context.Context,
	orgId uuid.UUID) (
	models.ApiResponse[models.AosRegisterCmd],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/aos/register_cmd")
	req.AppendTemplateParams(orgId)
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

	var result models.AosRegisterCmd
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.AosRegisterCmd](decoder)
	return models.NewApiResponse(result, resp), err
}
