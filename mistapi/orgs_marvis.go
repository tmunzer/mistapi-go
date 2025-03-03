package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/errors"
    "github.com/tmunzer/mistapi-go/mistapi/models"
)

// OrgsMarvis represents a controller struct.
type OrgsMarvis struct {
    baseController
}

// NewOrgsMarvis creates a new instance of OrgsMarvis.
// It takes a baseController as a parameter and returns a pointer to the OrgsMarvis.
func NewOrgsMarvis(baseController baseController) *OrgsMarvis {
    orgsMarvis := OrgsMarvis{baseController: baseController}
    return &orgsMarvis
}

// TroubleshootOrg takes context, orgId, mac, siteId, start, end, mType as parameters and
// returns an models.ApiResponse with models.ResponseTroubleshoot data and
// an error if there was an issue with the request or response.
// Troubleshoot sites, devices, clients, and wired clients for maximum of last 7 days from current time. See search APIs for device information:
// - [search Device]($e/Orgs%20Devices/searchOrgDevices)
// - [search Wireless Client]($e/Orgs%20Clients%20-%20Wireless/searchOrgWirelessClients)
// - [search Wired Client]($e/Orgs%20Clients%20-%20Wired/searchOrgWiredClients)
// - [search Wan Client]($e/Orgs%20Clients%20-%20Wan/searchOrgWanClients)
// **NOTE**: requires Marvis subscription license
func (o *OrgsMarvis) TroubleshootOrg(
    ctx context.Context,
    orgId uuid.UUID,
    mac *string,
    siteId *uuid.UUID,
    start *int,
    end *int,
    mType *models.TroubleshootTypeEnum) (
    models.ApiResponse[models.ResponseTroubleshoot],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/troubleshoot")
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
    if mac != nil {
        req.QueryParam("mac", *mac)
    }
    if siteId != nil {
        req.QueryParam("site_id", *siteId)
    }
    if start != nil {
        req.QueryParam("start", *start)
    }
    if end != nil {
        req.QueryParam("end", *end)
    }
    if mType != nil {
        req.QueryParam("type", *mType)
    }
    
    var result models.ResponseTroubleshoot
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseTroubleshoot](decoder)
    return models.NewApiResponse(result, resp), err
}
