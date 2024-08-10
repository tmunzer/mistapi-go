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

// SitesLicenses represents a controller struct.
type SitesLicenses struct {
    baseController
}

// NewSitesLicenses creates a new instance of SitesLicenses.
// It takes a baseController as a parameter and returns a pointer to the SitesLicenses.
func NewSitesLicenses(baseController baseController) *SitesLicenses {
    sitesLicenses := SitesLicenses{baseController: baseController}
    return &sitesLicenses
}

// GetSiteLicenseUsage takes context, siteId as parameters and
// returns an models.ApiResponse with models.LicenseUsageSite data and
// an error if there was an issue with the request or response.
// This shows license usage (i.e. needed) based on the features enabled for the site.
func (s *SitesLicenses) GetSiteLicenseUsage(
    ctx context.Context,
    siteId uuid.UUID) (
    models.ApiResponse[models.LicenseUsageSite],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/licenses/usages", siteId),
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
    
    var result models.LicenseUsageSite
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.LicenseUsageSite](decoder)
    return models.NewApiResponse(result, resp), err
}
