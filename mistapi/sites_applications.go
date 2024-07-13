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

// SitesApplications represents a controller struct.
type SitesApplications struct {
	baseController
}

// NewSitesApplications creates a new instance of SitesApplications.
// It takes a baseController as a parameter and returns a pointer to the SitesApplications.
func NewSitesApplications(baseController baseController) *SitesApplications {
	sitesApplications := SitesApplications{baseController: baseController}
	return &sitesApplications
}

// ListSiteApps takes context, siteId as parameters and
// returns an models.ApiResponse with []models.SiteApp data and
// an error if there was an issue with the request or response.
// Get List of Site Applications
func (s *SitesApplications) ListSiteApps(
	ctx context.Context,
	siteId uuid.UUID) (
	models.ApiResponse[[]models.SiteApp],
	error) {
	req := s.prepareRequest(ctx, "GET", fmt.Sprintf("/api/v1/sites/%v/apps", siteId))
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

	var result []models.SiteApp
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.SiteApp](decoder)
	return models.NewApiResponse(result, resp), err
}

// CountSiteApps takes context, siteId, distinct, deviceMac, app, wired as parameters and
// returns an models.ApiResponse with models.RepsonseCount data and
// an error if there was an issue with the request or response.
// Count by Distinct Attributes of Applications
func (s *SitesApplications) CountSiteApps(
	ctx context.Context,
	siteId uuid.UUID,
	distinct *models.SiteAppsCountDistinctEnum,
	deviceMac *string,
	app *string,
	wired *string) (
	models.ApiResponse[models.RepsonseCount],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/stats/apps/count", siteId),
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
	if distinct != nil {
		req.QueryParam("distinct", *distinct)
	}
	if deviceMac != nil {
		req.QueryParam("device_mac", *deviceMac)
	}
	if app != nil {
		req.QueryParam("app", *app)
	}
	if wired != nil {
		req.QueryParam("wired", *wired)
	}

	var result models.RepsonseCount
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.RepsonseCount](decoder)
	return models.NewApiResponse(result, resp), err
}
