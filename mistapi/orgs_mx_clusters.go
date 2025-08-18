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

// OrgsMxClusters represents a controller struct.
type OrgsMxClusters struct {
	baseController
}

// NewOrgsMxClusters creates a new instance of OrgsMxClusters.
// It takes a baseController as a parameter and returns a pointer to the OrgsMxClusters.
func NewOrgsMxClusters(baseController baseController) *OrgsMxClusters {
	orgsMxClusters := OrgsMxClusters{baseController: baseController}
	return &orgsMxClusters
}

// ListOrgMxEdgeClusters takes context, orgId, limit, page as parameters and
// returns an models.ApiResponse with []models.Mxcluster data and
// an error if there was an issue with the request or response.
// Get List of Org MxEdge Clusters
func (o *OrgsMxClusters) ListOrgMxEdgeClusters(
	ctx context.Context,
	orgId uuid.UUID,
	limit *int,
	page *int) (
	models.ApiResponse[[]models.Mxcluster],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/mxclusters")
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
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if page != nil {
		req.QueryParam("page", *page)
	}

	var result []models.Mxcluster
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.Mxcluster](decoder)
	return models.NewApiResponse(result, resp), err
}

// CreateOrgMxEdgeCluster takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.Mxcluster data and
// an error if there was an issue with the request or response.
// Create MxCluster
func (o *OrgsMxClusters) CreateOrgMxEdgeCluster(
	ctx context.Context,
	orgId uuid.UUID,
	body *models.Mxcluster) (
	models.ApiResponse[models.Mxcluster],
	error) {
	req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/mxclusters")
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
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	var result models.Mxcluster
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Mxcluster](decoder)
	return models.NewApiResponse(result, resp), err
}

// DeleteOrgMxEdgeCluster takes context, orgId, mxclusterId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete Org MXEdge Cluster
func (o *OrgsMxClusters) DeleteOrgMxEdgeCluster(
	ctx context.Context,
	orgId uuid.UUID,
	mxclusterId uuid.UUID) (
	*http.Response,
	error) {
	req := o.prepareRequest(ctx, "DELETE", "/api/v1/orgs/%v/mxclusters/%v")
	req.AppendTemplateParams(orgId, mxclusterId)
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

// GetOrgMxEdgeCluster takes context, orgId, mxclusterId as parameters and
// returns an models.ApiResponse with models.Mxcluster data and
// an error if there was an issue with the request or response.
// Get Org MxEdge Cluster Details
func (o *OrgsMxClusters) GetOrgMxEdgeCluster(
	ctx context.Context,
	orgId uuid.UUID,
	mxclusterId uuid.UUID) (
	models.ApiResponse[models.Mxcluster],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/mxclusters/%v")
	req.AppendTemplateParams(orgId, mxclusterId)
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

	var result models.Mxcluster
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Mxcluster](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpdateOrgMxEdgeCluster takes context, orgId, mxclusterId, body as parameters and
// returns an models.ApiResponse with models.Mxcluster data and
// an error if there was an issue with the request or response.
// Update Org MxEdge Cluster
func (o *OrgsMxClusters) UpdateOrgMxEdgeCluster(
	ctx context.Context,
	orgId uuid.UUID,
	mxclusterId uuid.UUID,
	body *models.Mxcluster) (
	models.ApiResponse[models.Mxcluster],
	error) {
	req := o.prepareRequest(ctx, "PUT", "/api/v1/orgs/%v/mxclusters/%v")
	req.AppendTemplateParams(orgId, mxclusterId)
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

	var result models.Mxcluster
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Mxcluster](decoder)
	return models.NewApiResponse(result, resp), err
}
