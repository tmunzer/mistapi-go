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

// SitesJSE represents a controller struct.
type SitesJSE struct {
    baseController
}

// NewSitesJSE creates a new instance of SitesJSE.
// It takes a baseController as a parameter and returns a pointer to the SitesJSE.
func NewSitesJSE(baseController baseController) *SitesJSE {
    sitesJSE := SitesJSE{baseController: baseController}
    return &sitesJSE
}

// GetSiteJseInfo takes context, siteId as parameters and
// returns an models.ApiResponse with models.AccountJseInfo data and
// an error if there was an issue with the request or response.
// Retrieves the list of JSE orgs associated with the account
func (s *SitesJSE) GetSiteJseInfo(
    ctx context.Context,
    siteId uuid.UUID) (
    models.ApiResponse[models.AccountJseInfo],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/setting/jse/info")
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
    
    var result models.AccountJseInfo
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.AccountJseInfo](decoder)
    return models.NewApiResponse(result, resp), err
}
