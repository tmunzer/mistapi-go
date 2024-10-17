package mistapi

import (
	"context"
	"fmt"
	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
	"github.com/google/uuid"
	"github.com/tmunzer/mistapi-go/mistapi/errors"
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

// MSPsInventory represents a controller struct.
type MSPsInventory struct {
	baseController
}

// NewMSPsInventory creates a new instance of MSPsInventory.
// It takes a baseController as a parameter and returns a pointer to the MSPsInventory.
func NewMSPsInventory(baseController baseController) *MSPsInventory {
	mSPsInventory := MSPsInventory{baseController: baseController}
	return &mSPsInventory
}

// GetMspInventoryByMac takes context, mspId, deviceMac as parameters and
// returns an models.ApiResponse with models.ResponseMspInventoryDevice data and
// an error if there was an issue with the request or response.
// Get Inventoy By device MAC address
func (m *MSPsInventory) GetMspInventoryByMac(
	ctx context.Context,
	mspId uuid.UUID,
	deviceMac string) (
	models.ApiResponse[models.ResponseMspInventoryDevice],
	error) {
	req := m.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/msps/%v/inventory/%v", mspId, deviceMac),
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
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
	})

	var result models.ResponseMspInventoryDevice
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseMspInventoryDevice](decoder)
	return models.NewApiResponse(result, resp), err
}
