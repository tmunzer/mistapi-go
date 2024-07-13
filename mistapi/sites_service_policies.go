package mistapi

import (
	"context"
	"fmt"

	"github.com/tmunzer/mistapi-go/mistapi/errors"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
	"github.com/google/uuid"
)

// SitesServicePolicies represents a controller struct.
type SitesServicePolicies struct {
	baseController
}

// NewSitesServicePolicies creates a new instance of SitesServicePolicies.
// It takes a baseController as a parameter and returns a pointer to the SitesServicePolicies.
func NewSitesServicePolicies(baseController baseController) *SitesServicePolicies {
	sitesServicePolicies := SitesServicePolicies{baseController: baseController}
	return &sitesServicePolicies
}

// ListSiteServicePoliciesDerived takes context, siteId, resolve as parameters and
// returns an models.ApiResponse with []models.ServicePolicy data and
// an error if there was an issue with the request or response.
// Retrieves the list of Service Policies available for the Site
func (s *SitesServicePolicies) ListSiteServicePoliciesDerived(
	ctx context.Context,
	siteId uuid.UUID,
	resolve *bool) (
	models.ApiResponse[[]models.ServicePolicy],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/servicepolicies/derived", siteId),
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
	if resolve != nil {
		req.QueryParam("resolve", *resolve)
	}

	var result []models.ServicePolicy
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.ServicePolicy](decoder)
	return models.NewApiResponse(result, resp), err
}