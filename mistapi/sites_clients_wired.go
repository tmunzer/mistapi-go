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

// SitesClientsWired represents a controller struct.
type SitesClientsWired struct {
    baseController
}

// NewSitesClientsWired creates a new instance of SitesClientsWired.
// It takes a baseController as a parameter and returns a pointer to the SitesClientsWired.
func NewSitesClientsWired(baseController baseController) *SitesClientsWired {
    sitesClientsWired := SitesClientsWired{baseController: baseController}
    return &sitesClientsWired
}

// CountSiteWiredClients takes context, siteId, distinct, mac, deviceMac, portId, vlan, start, end, duration, limit, page as parameters and
// returns an models.ApiResponse with models.RepsonseCount data and
// an error if there was an issue with the request or response.
// Count by Distinct Attributes of Clients
func (s *SitesClientsWired) CountSiteWiredClients(
    ctx context.Context,
    siteId uuid.UUID,
    distinct *models.SiteWiredClientsCountDistinctEnum,
    mac *string,
    deviceMac *string,
    portId *string,
    vlan *string,
    start *int,
    end *int,
    duration *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.RepsonseCount],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/wired_clients/count", siteId),
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
    if mac != nil {
        req.QueryParam("mac", *mac)
    }
    if deviceMac != nil {
        req.QueryParam("device_mac", *deviceMac)
    }
    if portId != nil {
        req.QueryParam("port_id", *portId)
    }
    if vlan != nil {
        req.QueryParam("vlan", *vlan)
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
    if limit != nil {
        req.QueryParam("limit", *limit)
    }
    if page != nil {
        req.QueryParam("page", *page)
    }
    
    var result models.RepsonseCount
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.RepsonseCount](decoder)
    return models.NewApiResponse(result, resp), err
}

// SearchSiteWiredClients takes context, siteId, deviceMac, mac, ip, portId, vlan, manufacture, text, nacruleId, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.SearchWiredClient data and
// an error if there was an issue with the request or response.
// Search Wired Clients
func (s *SitesClientsWired) SearchSiteWiredClients(
    ctx context.Context,
    siteId uuid.UUID,
    deviceMac *string,
    mac *string,
    ip *string,
    portId *string,
    vlan *string,
    manufacture *string,
    text *string,
    nacruleId *string,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.SearchWiredClient],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/wired_clients/search", siteId),
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
    if deviceMac != nil {
        req.QueryParam("device_mac", *deviceMac)
    }
    if mac != nil {
        req.QueryParam("mac", *mac)
    }
    if ip != nil {
        req.QueryParam("ip", *ip)
    }
    if portId != nil {
        req.QueryParam("port_id", *portId)
    }
    if vlan != nil {
        req.QueryParam("vlan", *vlan)
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
