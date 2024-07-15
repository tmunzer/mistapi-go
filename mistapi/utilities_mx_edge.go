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

// UtilitiesMxEdge represents a controller struct.
type UtilitiesMxEdge struct {
    baseController
}

// NewUtilitiesMxEdge creates a new instance of UtilitiesMxEdge.
// It takes a baseController as a parameter and returns a pointer to the UtilitiesMxEdge.
func NewUtilitiesMxEdge(baseController baseController) *UtilitiesMxEdge {
    utilitiesMxEdge := UtilitiesMxEdge{baseController: baseController}
    return &utilitiesMxEdge
}

// PreemptSitesMxTunnel takes context, siteId, mxtunnelId as parameters and
// returns an models.ApiResponse with models.ResponseMxtunnelsPreemptAps data and
// an error if there was an issue with the request or response.
// To preempt AP’s which are not connected to preferred peer to the preferred peer
func (u *UtilitiesMxEdge) PreemptSitesMxTunnel(
    ctx context.Context,
    siteId uuid.UUID,
    mxtunnelId uuid.UUID) (
    models.ApiResponse[models.ResponseMxtunnelsPreemptAps],
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/mxtunnels/%v/preempt_aps", siteId, mxtunnelId),
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
    
    var result models.ResponseMxtunnelsPreemptAps
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseMxtunnelsPreemptAps](decoder)
    return models.NewApiResponse(result, resp), err
}
