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

// SitesStats represents a controller struct.
type SitesStats struct {
	baseController
}

// NewSitesStats creates a new instance of SitesStats.
// It takes a baseController as a parameter and returns a pointer to the SitesStats.
func NewSitesStats(baseController baseController) *SitesStats {
	sitesStats := SitesStats{baseController: baseController}
	return &sitesStats
}

// GetSiteStats takes context, siteId as parameters and
// returns an models.ApiResponse with models.StatsSite data and
// an error if there was an issue with the request or response.
// Get Sites Stats
func (s *SitesStats) GetSiteStats(
	ctx context.Context,
	siteId uuid.UUID) (
	models.ApiResponse[models.StatsSite],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/stats", siteId),
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

	var result models.StatsSite
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.StatsSite](decoder)
	return models.NewApiResponse(result, resp), err
}
