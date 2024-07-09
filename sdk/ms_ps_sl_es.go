package mistapi

import (
	"context"
	"fmt"

	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
	"github.com/google/uuid"
	"github.com/tmunzer/mistapi/sdk/errors"
	"github.com/tmunzer/mistapi/sdk/models"
)

// MSPsSLEs represents a controller struct.
type MSPsSLEs struct {
	baseController
}

// NewMSPsSLEs creates a new instance of MSPsSLEs.
// It takes a baseController as a parameter and returns a pointer to the MSPsSLEs.
func NewMSPsSLEs(baseController baseController) *MSPsSLEs {
	mSPsSLEs := MSPsSLEs{baseController: baseController}
	return &mSPsSLEs
}

// GetMspSle takes context, mspId, metric, sle, duration, interval, start, end as parameters and
// returns an models.ApiResponse with models.InsightMetrics data and
// an error if there was an issue with the request or response.
// Get MSP SLEs (all/worst Orgs ...)
func (m *MSPsSLEs) GetMspSle(
	ctx context.Context,
	mspId uuid.UUID,
	metric string,
	sle *string,
	duration *string,
	interval *string,
	start *int,
	end *int) (
	models.ApiResponse[models.InsightMetrics],
	error) {
	req := m.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/msps/%v/insights/%v", mspId, metric),
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
	if sle != nil {
		req.QueryParam("sle", *sle)
	}
	if duration != nil {
		req.QueryParam("duration", *duration)
	}
	if interval != nil {
		req.QueryParam("interval", *interval)
	}
	if start != nil {
		req.QueryParam("start", *start)
	}
	if end != nil {
		req.QueryParam("end", *end)
	}

	var result models.InsightMetrics
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.InsightMetrics](decoder)
	return models.NewApiResponse(result, resp), err
}
