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
)

// OrgsVars represents a controller struct.
type OrgsVars struct {
	baseController
}

// NewOrgsVars creates a new instance of OrgsVars.
// It takes a baseController as a parameter and returns a pointer to the OrgsVars.
func NewOrgsVars(baseController baseController) *OrgsVars {
	orgsVars := OrgsVars{baseController: baseController}
	return &orgsVars
}

// SearchOrgVars takes context, orgId, siteId, mVar, src, limit, sort, searchAfter as parameters and
// returns an models.ApiResponse with models.ResponseSearchVar data and
// an error if there was an issue with the request or response.
// Search vars
// Example: /api/v1/orgs/{org_id}/vars/search?vars=*
func (o *OrgsVars) SearchOrgVars(
	ctx context.Context,
	orgId uuid.UUID,
	siteId *string,
	mVar *string,
	src *models.VarSourceEnum,
	limit *int,
	sort *string,
	searchAfter *string) (
	models.ApiResponse[models.ResponseSearchVar],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/vars/search")
	req.AppendTemplateParams(orgId)
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
	if siteId != nil {
		req.QueryParam("site_id", *siteId)
	}
	if mVar != nil {
		req.QueryParam("var", *mVar)
	}
	if src != nil {
		req.QueryParam("src", *src)
	}
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if sort != nil {
		req.QueryParam("sort", *sort)
	}
	if searchAfter != nil {
		req.QueryParam("search_after", *searchAfter)
	}

	var result models.ResponseSearchVar
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseSearchVar](decoder)
	return models.NewApiResponse(result, resp), err
}
