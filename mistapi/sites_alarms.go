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
	"net/http"
)

// SitesAlarms represents a controller struct.
type SitesAlarms struct {
	baseController
}

// NewSitesAlarms creates a new instance of SitesAlarms.
// It takes a baseController as a parameter and returns a pointer to the SitesAlarms.
func NewSitesAlarms(baseController baseController) *SitesAlarms {
	sitesAlarms := SitesAlarms{baseController: baseController}
	return &sitesAlarms
}

// AckSiteMultipleAlarms takes context, siteId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Ack multiple Site Alarms
func (s *SitesAlarms) AckSiteMultipleAlarms(
	ctx context.Context,
	siteId uuid.UUID,
	body *models.AlarmAck) (
	*http.Response,
	error) {
	req := s.prepareRequest(ctx, "POST", "/api/v1/sites/%v/alarms/ack")
	req.AppendTemplateParams(siteId)
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
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}

// AckSiteAllAlarms takes context, siteId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Ack all Site Alarms
// **N.B.**: Batch size for multiple alarm ack and unack has to be less or or equal to 1000.
func (s *SitesAlarms) AckSiteAllAlarms(
	ctx context.Context,
	siteId uuid.UUID,
	body *models.NoteString) (
	*http.Response,
	error) {
	req := s.prepareRequest(ctx, "POST", "/api/v1/sites/%v/alarms/ack_all")
	req.AppendTemplateParams(siteId)
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
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}

// CountSiteAlarms takes context, siteId, distinct, ackAdminName, acked, mType, severity, group, start, end, duration, limit as parameters and
// returns an models.ApiResponse with models.ResponseCount data and
// an error if there was an issue with the request or response.
// Count by Distinct Attributes of Site Alarms
func (s *SitesAlarms) CountSiteAlarms(
	ctx context.Context,
	siteId uuid.UUID,
	distinct *models.AlarmCountDistinctEnum,
	ackAdminName *string,
	acked *bool,
	mType *string,
	severity *string,
	group *string,
	start *string,
	end *string,
	duration *string,
	limit *int) (
	models.ApiResponse[models.ResponseCount],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/alarms/count")
	req.AppendTemplateParams(siteId)
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
	if distinct != nil {
		req.QueryParam("distinct", *distinct)
	}
	if ackAdminName != nil {
		req.QueryParam("ack_admin_name", *ackAdminName)
	}
	if acked != nil {
		req.QueryParam("acked", *acked)
	}
	if mType != nil {
		req.QueryParam("type", *mType)
	}
	if severity != nil {
		req.QueryParam("severity", *severity)
	}
	if group != nil {
		req.QueryParam("group", *group)
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
	if limit != nil {
		req.QueryParam("limit", *limit)
	}

	var result models.ResponseCount
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseCount](decoder)
	return models.NewApiResponse(result, resp), err
}

// SearchSiteAlarms takes context, siteId, mType, ackAdminName, acked, severity, group, limit, start, end, duration, sort, searchAfter as parameters and
// returns an models.ApiResponse with models.AlarmSearchResult data and
// an error if there was an issue with the request or response.
// Search Site Alarms
func (s *SitesAlarms) SearchSiteAlarms(
	ctx context.Context,
	siteId uuid.UUID,
	mType *string,
	ackAdminName *string,
	acked *bool,
	severity *string,
	group *string,
	limit *int,
	start *string,
	end *string,
	duration *string,
	sort *string,
	searchAfter *string) (
	models.ApiResponse[models.AlarmSearchResult],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/alarms/search")
	req.AppendTemplateParams(siteId)
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
	if mType != nil {
		req.QueryParam("type", *mType)
	}
	if ackAdminName != nil {
		req.QueryParam("ack_admin_name", *ackAdminName)
	}
	if acked != nil {
		req.QueryParam("acked", *acked)
	}
	if severity != nil {
		req.QueryParam("severity", *severity)
	}
	if group != nil {
		req.QueryParam("group", *group)
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
	if sort != nil {
		req.QueryParam("sort", *sort)
	}
	if searchAfter != nil {
		req.QueryParam("search_after", *searchAfter)
	}

	var result models.AlarmSearchResult
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.AlarmSearchResult](decoder)
	return models.NewApiResponse(result, resp), err
}

// UnackSiteMultipleAlarms takes context, siteId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Unack multiple Site Alarms
func (s *SitesAlarms) UnackSiteMultipleAlarms(
	ctx context.Context,
	siteId uuid.UUID,
	body *models.AlarmAck) (
	*http.Response,
	error) {
	req := s.prepareRequest(ctx, "POST", "/api/v1/sites/%v/alarms/unack")
	req.AppendTemplateParams(siteId)
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
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}

// UnackSiteAllAlarms takes context, siteId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Unack all Site Alarms
// **N.B.**: Batch size for multiple alarm ack and unack has to be less or or equal to 1000.
func (s *SitesAlarms) UnackSiteAllAlarms(
	ctx context.Context,
	siteId uuid.UUID,
	body *models.NoteString) (
	*http.Response,
	error) {
	req := s.prepareRequest(ctx, "POST", "/api/v1/sites/%v/alarms/unack_all")
	req.AppendTemplateParams(siteId)
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
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}

// AckSiteAlarm takes context, siteId, alarmId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Ack Site Alarm
func (s *SitesAlarms) AckSiteAlarm(
	ctx context.Context,
	siteId uuid.UUID,
	alarmId uuid.UUID,
	body *models.NoteString) (
	*http.Response,
	error) {
	req := s.prepareRequest(ctx, "POST", "/api/v1/sites/%v/alarms/%v/ack")
	req.AppendTemplateParams(siteId, alarmId)
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
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}

// UnackSiteAlarm takes context, siteId, alarmId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Unack Site Alarm
func (s *SitesAlarms) UnackSiteAlarm(
	ctx context.Context,
	siteId uuid.UUID,
	alarmId uuid.UUID,
	body *models.NoteString) (
	*http.Response,
	error) {
	req := s.prepareRequest(ctx, "POST", "/api/v1/sites/%v/alarms/%v/unack")
	req.AppendTemplateParams(siteId, alarmId)
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
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}

// UnsubscribeSiteAlarms takes context, siteId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Unsubscribe to Site Alarms
func (s *SitesAlarms) UnsubscribeSiteAlarms(
	ctx context.Context,
	siteId uuid.UUID) (
	*http.Response,
	error) {
	req := s.prepareRequest(ctx, "DELETE", "/api/v1/sites/%v/subscriptions")
	req.AppendTemplateParams(siteId)
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

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}

// SubscribeSiteAlarms takes context, siteId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Subscribe to Site Alarms
func (s *SitesAlarms) SubscribeSiteAlarms(
	ctx context.Context,
	siteId uuid.UUID) (
	*http.Response,
	error) {
	req := s.prepareRequest(ctx, "POST", "/api/v1/sites/%v/subscriptions")
	req.AppendTemplateParams(siteId)
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

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}
