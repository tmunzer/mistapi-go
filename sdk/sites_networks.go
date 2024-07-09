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

// SitesNetworks represents a controller struct.
type SitesNetworks struct {
	baseController
}

// NewSitesNetworks creates a new instance of SitesNetworks.
// It takes a baseController as a parameter and returns a pointer to the SitesNetworks.
func NewSitesNetworks(baseController baseController) *SitesNetworks {
	sitesNetworks := SitesNetworks{baseController: baseController}
	return &sitesNetworks
}

// ListSiteNetworksDerived takes context, siteId, resolve as parameters and
// returns an models.ApiResponse with []models.Network data and
// an error if there was an issue with the request or response.
// Retrieves the list of Networks available for the Site
func (s *SitesNetworks) ListSiteNetworksDerived(
	ctx context.Context,
	siteId uuid.UUID,
	resolve *bool) (
	models.ApiResponse[[]models.Network],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/networks/derived", siteId),
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
	if resolve != nil {
		req.QueryParam("resolve", *resolve)
	}

	var result []models.Network
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.Network](decoder)
	return models.NewApiResponse(result, resp), err
}