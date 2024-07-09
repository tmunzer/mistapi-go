package mistapigo

import (
	"context"

	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
	"github.com/tmunzer/mistapigo/sdk/errors"
	"github.com/tmunzer/mistapigo/sdk/models"
)

// SelfAuditLogs represents a controller struct.
type SelfAuditLogs struct {
	baseController
}

// NewSelfAuditLogs creates a new instance of SelfAuditLogs.
// It takes a baseController as a parameter and returns a pointer to the SelfAuditLogs.
func NewSelfAuditLogs(baseController baseController) *SelfAuditLogs {
	selfAuditLogs := SelfAuditLogs{baseController: baseController}
	return &selfAuditLogs
}

// ListSelfAuditLogs takes context, page, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.ResponseSelfAuditLogs data and
// an error if there was an issue with the request or response.
// Get List of change logs across all Orgs for current admin
// Audit logs records all administrative activities done by current admin across all orgs
func (s *SelfAuditLogs) ListSelfAuditLogs(
	ctx context.Context,
	page *int,
	limit *int,
	start *int,
	end *int,
	duration *string) (
	models.ApiResponse[models.ResponseSelfAuditLogs],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/self/logs")
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
	if page != nil {
		req.QueryParam("page", *page)
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
	var result models.ResponseSelfAuditLogs
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseSelfAuditLogs](decoder)
	return models.NewApiResponse(result, resp), err
}
