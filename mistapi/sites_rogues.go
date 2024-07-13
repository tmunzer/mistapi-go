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

// SitesRogues represents a controller struct.
type SitesRogues struct {
	baseController
}

// NewSitesRogues creates a new instance of SitesRogues.
// It takes a baseController as a parameter and returns a pointer to the SitesRogues.
func NewSitesRogues(baseController baseController) *SitesRogues {
	sitesRogues := SitesRogues{baseController: baseController}
	return &sitesRogues
}

// ListSiteRogueAPs takes context, siteId, mType, limit, start, end, duration, interval as parameters and
// returns an models.ApiResponse with models.ResponseInsightRogue data and
// an error if there was an issue with the request or response.
// Get List of Site Rogue/Neighbor APs
func (s *SitesRogues) ListSiteRogueAPs(
	ctx context.Context,
	siteId uuid.UUID,
	mType *models.RogueTypeEnum,
	limit *int,
	start *int,
	end *int,
	duration *string,
	interval *string) (
	models.ApiResponse[models.ResponseInsightRogue],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/insights/rogues", siteId),
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
	if mType != nil {
		req.QueryParam("type", *mType)
	}
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if start != nil {
		req.QueryParam("start", *start)
	}
	if end != nil {
		req.QueryParam("end", *end)
	}
	if duration != nil {
		req.QueryParam("duration", *duration)
	}
	if interval != nil {
		req.QueryParam("interval", *interval)
	}

	var result models.ResponseInsightRogue
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseInsightRogue](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListSiteRogueClients takes context, siteId, limit, start, end, duration, interval as parameters and
// returns an models.ApiResponse with models.ResponseInsightRogueClient data and
// an error if there was an issue with the request or response.
// Get List of Site Rogue Clients
func (s *SitesRogues) ListSiteRogueClients(
	ctx context.Context,
	siteId uuid.UUID,
	limit *int,
	start *int,
	end *int,
	duration *string,
	interval *string) (
	models.ApiResponse[models.ResponseInsightRogueClient],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/insights/rogues/clients", siteId),
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
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if start != nil {
		req.QueryParam("start", *start)
	}
	if end != nil {
		req.QueryParam("end", *end)
	}
	if duration != nil {
		req.QueryParam("duration", *duration)
	}
	if interval != nil {
		req.QueryParam("interval", *interval)
	}

	var result models.ResponseInsightRogueClient
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseInsightRogueClient](decoder)
	return models.NewApiResponse(result, resp), err
}

// CountSiteRogueEvents takes context, siteId, distinct, mType, ssid, bssid, apMac, channel, seenOnLan, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.RepsonseCount data and
// an error if there was an issue with the request or response.
// Count Rogue Events
func (s *SitesRogues) CountSiteRogueEvents(
	ctx context.Context,
	siteId uuid.UUID,
	distinct *models.SiteRogueEventsCountDistinctEnum,
	mType *models.RogueTypeEnum,
	ssid *string,
	bssid *string,
	apMac *string,
	channel *string,
	seenOnLan *bool,
	limit *int,
	start *int,
	end *int,
	duration *string) (
	models.ApiResponse[models.RepsonseCount],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/rogues/events/count", siteId),
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
	if mType != nil {
		req.QueryParam("type", *mType)
	}
	if ssid != nil {
		req.QueryParam("ssid", *ssid)
	}
	if bssid != nil {
		req.QueryParam("bssid", *bssid)
	}
	if apMac != nil {
		req.QueryParam("ap_mac", *apMac)
	}
	if channel != nil {
		req.QueryParam("channel", *channel)
	}
	if seenOnLan != nil {
		req.QueryParam("seen_on_lan", *seenOnLan)
	}
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if start != nil {
		req.QueryParam("start", *start)
	}
	if end != nil {
		req.QueryParam("end", *end)
	}
	if duration != nil {
		req.QueryParam("duration", *duration)
	}

	var result models.RepsonseCount
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.RepsonseCount](decoder)
	return models.NewApiResponse(result, resp), err
}

// SearchSiteRogueEvents takes context, siteId, mType, ssid, bssid, apMac, channel, seenOnLan, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.ResponseEventsRogueSearch data and
// an error if there was an issue with the request or response.
// Search Rogue Events
func (s *SitesRogues) SearchSiteRogueEvents(
	ctx context.Context,
	siteId uuid.UUID,
	mType *models.RogueTypeEnum,
	ssid *string,
	bssid *string,
	apMac *string,
	channel *int,
	seenOnLan *bool,
	limit *int,
	start *int,
	end *int,
	duration *string) (
	models.ApiResponse[models.ResponseEventsRogueSearch],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/rogues/events/search", siteId),
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
	if mType != nil {
		req.QueryParam("type", *mType)
	}
	if ssid != nil {
		req.QueryParam("ssid", *ssid)
	}
	if bssid != nil {
		req.QueryParam("bssid", *bssid)
	}
	if apMac != nil {
		req.QueryParam("ap_mac", *apMac)
	}
	if channel != nil {
		req.QueryParam("channel", *channel)
	}
	if seenOnLan != nil {
		req.QueryParam("seen_on_lan", *seenOnLan)
	}
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if start != nil {
		req.QueryParam("start", *start)
	}
	if end != nil {
		req.QueryParam("end", *end)
	}
	if duration != nil {
		req.QueryParam("duration", *duration)
	}

	var result models.ResponseEventsRogueSearch
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseEventsRogueSearch](decoder)
	return models.NewApiResponse(result, resp), err
}

// GetSiteRogueAP takes context, siteId, rogueBssid as parameters and
// returns an models.ApiResponse with models.RogueDetails data and
// an error if there was an issue with the request or response.
// Get Rogue AP Details
func (s *SitesRogues) GetSiteRogueAP(
	ctx context.Context,
	siteId uuid.UUID,
	rogueBssid string) (
	models.ApiResponse[models.RogueDetails],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/rogues/%v", siteId, rogueBssid),
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

	var result models.RogueDetails
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.RogueDetails](decoder)
	return models.NewApiResponse(result, resp), err
}
