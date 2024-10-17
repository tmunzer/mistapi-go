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

// OrgsSLEs represents a controller struct.
type OrgsSLEs struct {
	baseController
}

// NewOrgsSLEs creates a new instance of OrgsSLEs.
// It takes a baseController as a parameter and returns a pointer to the OrgsSLEs.
func NewOrgsSLEs(baseController baseController) *OrgsSLEs {
	orgsSLEs := OrgsSLEs{baseController: baseController}
	return &orgsSLEs
}

// GetOrgSitesSle takes context, orgId, sle, start, end, duration, interval, limit, page as parameters and
// returns an models.ApiResponse with models.ResponseOrgSiteSle data and
// an error if there was an issue with the request or response.
// Get Org Sites SLE
func (o *OrgsSLEs) GetOrgSitesSle(
	ctx context.Context,
	orgId uuid.UUID,
	sle *models.OrgSiteSleTypeEnum,
	start *int,
	end *int,
	duration *string,
	interval *string,
	limit *int,
	page *int) (
	models.ApiResponse[models.ResponseOrgSiteSle],
	error) {
	req := o.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/orgs/%v/insights/sites-sle", orgId),
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
	if sle != nil {
		req.QueryParam("sle", *sle)
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
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if page != nil {
		req.QueryParam("page", *page)
	}

	var result models.ResponseOrgSiteSle
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseOrgSiteSle](decoder)
	return models.NewApiResponse(result, resp), err
}

// GetOrgSle takes context, orgId, metric, sle, duration, interval, start, end as parameters and
// returns an models.ApiResponse with models.InsightMetrics data and
// an error if there was an issue with the request or response.
// Get Org SLEs (all/worst sites, Mx Edges, ...)
func (o *OrgsSLEs) GetOrgSle(
	ctx context.Context,
	orgId uuid.UUID,
	metric string,
	sle *string,
	duration *string,
	interval *string,
	start *int,
	end *int) (
	models.ApiResponse[models.InsightMetrics],
	error) {
	req := o.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/orgs/%v/insights/%v", orgId, metric),
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
