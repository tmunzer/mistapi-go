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

// SitesRFTemplates represents a controller struct.
type SitesRFTemplates struct {
    baseController
}

// NewSitesRFTemplates creates a new instance of SitesRFTemplates.
// It takes a baseController as a parameter and returns a pointer to the SitesRFTemplates.
func NewSitesRFTemplates(baseController baseController) *SitesRFTemplates {
    sitesRFTemplates := SitesRFTemplates{baseController: baseController}
    return &sitesRFTemplates
}

// GetSiteRfTemplateDerived takes context, siteId, resolve as parameters and
// returns an models.ApiResponse with models.RfTemplate data and
// an error if there was an issue with the request or response.
// Get derived RF Templates for Site
func (s *SitesRFTemplates) GetSiteRfTemplateDerived(
    ctx context.Context,
    siteId uuid.UUID,
    resolve *bool) (
    models.ApiResponse[models.RfTemplate],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/rftemplates/derived", siteId),
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
    if resolve != nil {
        req.QueryParam("resolve", *resolve)
    }
    
    var result models.RfTemplate
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.RfTemplate](decoder)
    return models.NewApiResponse(result, resp), err
}
