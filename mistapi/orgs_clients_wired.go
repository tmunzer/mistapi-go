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

// OrgsClientsWired represents a controller struct.
type OrgsClientsWired struct {
    baseController
}

// NewOrgsClientsWired creates a new instance of OrgsClientsWired.
// It takes a baseController as a parameter and returns a pointer to the OrgsClientsWired.
func NewOrgsClientsWired(baseController baseController) *OrgsClientsWired {
    orgsClientsWired := OrgsClientsWired{baseController: baseController}
    return &orgsClientsWired
}

// CountOrgWiredClients takes context, orgId, distinct, page, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.RepsonseCount data and
// an error if there was an issue with the request or response.
// Count by Distinct Attributes of Clients
// Note: For list of avaialable `type` values, please refer to [listClientEventsDefinitions]($e/Constants%20Events/listClientEventsDefinitions)
func (o *OrgsClientsWired) CountOrgWiredClients(
    ctx context.Context,
    orgId uuid.UUID,
    distinct *models.OrgWiredClientsCountDistinctEnum,
    page *int,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.RepsonseCount],
    error) {
    req := o.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/orgs/%v/wired_clients/count", orgId),
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
    if distinct != nil {
        req.QueryParam("distinct", *distinct)
    }
    if page != nil {
        req.QueryParam("page", *page)
    }
    if limit != nil {
        req.QueryParam("limit", *limit)
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
    
    var result models.RepsonseCount
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.RepsonseCount](decoder)
    return models.NewApiResponse(result, resp), err
}

// SearchOrgWiredClients takes context, orgId, siteId, deviceMac, mac, portId, vlan, ip, manufacture, text, nacruleId, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.SearchWiredClient data and
// an error if there was an issue with the request or response.
// Search for Wired Clients in org
// Note: For list of avaialable `type` values, please refer to [listClientEventsDefinitions]($e/Constants%20Events/listClientEventsDefinitions)
func (o *OrgsClientsWired) SearchOrgWiredClients(
    ctx context.Context,
    orgId uuid.UUID,
    siteId *string,
    deviceMac *string,
    mac *string,
    portId *string,
    vlan *int,
    ip *string,
    manufacture *string,
    text *string,
    nacruleId *string,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.SearchWiredClient],
    error) {
    req := o.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/orgs/%v/wired_clients/search", orgId),
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
    if siteId != nil {
        req.QueryParam("site_id", *siteId)
    }
    if deviceMac != nil {
        req.QueryParam("device_mac", *deviceMac)
    }
    if mac != nil {
        req.QueryParam("mac", *mac)
    }
    if portId != nil {
        req.QueryParam("port_id", *portId)
    }
    if vlan != nil {
        req.QueryParam("vlan", *vlan)
    }
    if ip != nil {
        req.QueryParam("ip", *ip)
    }
    if manufacture != nil {
        req.QueryParam("manufacture", *manufacture)
    }
    if text != nil {
        req.QueryParam("text", *text)
    }
    if nacruleId != nil {
        req.QueryParam("nacrule_id", *nacruleId)
    }
    if limit != nil {
        req.QueryParam("limit", *limit)
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
    
    var result models.SearchWiredClient
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SearchWiredClient](decoder)
    return models.NewApiResponse(result, resp), err
}
