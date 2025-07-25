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

// SitesAPTemplates represents a controller struct.
type SitesAPTemplates struct {
    baseController
}

// NewSitesAPTemplates creates a new instance of SitesAPTemplates.
// It takes a baseController as a parameter and returns a pointer to the SitesAPTemplates.
func NewSitesAPTemplates(baseController baseController) *SitesAPTemplates {
    sitesAPTemplates := SitesAPTemplates{baseController: baseController}
    return &sitesAPTemplates
}

// ListSiteApTemplatesDerived takes context, siteId, resolve as parameters and
// returns an models.ApiResponse with []models.ApTemplate data and
// an error if there was an issue with the request or response.
// Get the list of derived AP Templates for a site
func (s *SitesAPTemplates) ListSiteApTemplatesDerived(
    ctx context.Context,
    siteId uuid.UUID,
    resolve *bool) (
    models.ApiResponse[[]models.ApTemplate],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/aptemplates/derived")
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
    if resolve != nil {
        req.QueryParam("resolve", *resolve)
    }
    
    var result []models.ApTemplate
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.ApTemplate](decoder)
    return models.NewApiResponse(result, resp), err
}
