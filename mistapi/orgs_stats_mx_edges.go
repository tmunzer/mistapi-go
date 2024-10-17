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

// OrgsStatsMxEdges represents a controller struct.
type OrgsStatsMxEdges struct {
	baseController
}

// NewOrgsStatsMxEdges creates a new instance of OrgsStatsMxEdges.
// It takes a baseController as a parameter and returns a pointer to the OrgsStatsMxEdges.
func NewOrgsStatsMxEdges(baseController baseController) *OrgsStatsMxEdges {
	orgsStatsMxEdges := OrgsStatsMxEdges{baseController: baseController}
	return &orgsStatsMxEdges
}

// ListOrgMxEdgesStats takes context, orgId, forSite, start, end, duration, limit, page as parameters and
// returns an models.ApiResponse with []models.StatsMxedge data and
// an error if there was an issue with the request or response.
// Get List of Org MxEdge Stats
func (o *OrgsStatsMxEdges) ListOrgMxEdgesStats(
	ctx context.Context,
	orgId uuid.UUID,
	forSite *bool,
	start *int,
	end *int,
	duration *string,
	limit *int,
	page *int) (
	models.ApiResponse[[]models.StatsMxedge],
	error) {
	req := o.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/orgs/%v/stats/mxedges", orgId),
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
	if forSite != nil {
		req.QueryParam("for_site", *forSite)
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
	if page != nil {
		req.QueryParam("page", *page)
	}

	var result []models.StatsMxedge
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.StatsMxedge](decoder)
	return models.NewApiResponse(result, resp), err
}

// GetOrgMxEdgeStats takes context, orgId, mxedgeId, forSite as parameters and
// returns an models.ApiResponse with models.StatsMxedge data and
// an error if there was an issue with the request or response.
// Get Org MxEdge Details Stats
func (o *OrgsStatsMxEdges) GetOrgMxEdgeStats(
	ctx context.Context,
	orgId uuid.UUID,
	mxedgeId uuid.UUID,
	forSite *bool) (
	models.ApiResponse[models.StatsMxedge],
	error) {
	req := o.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/orgs/%v/stats/mxedges/%v", orgId, mxedgeId),
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
	if forSite != nil {
		req.QueryParam("for_site", *forSite)
	}

	var result models.StatsMxedge
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.StatsMxedge](decoder)
	return models.NewApiResponse(result, resp), err
}
