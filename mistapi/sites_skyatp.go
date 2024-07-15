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

// SitesSkyatp represents a controller struct.
type SitesSkyatp struct {
    baseController
}

// NewSitesSkyatp creates a new instance of SitesSkyatp.
// It takes a baseController as a parameter and returns a pointer to the SitesSkyatp.
func NewSitesSkyatp(baseController baseController) *SitesSkyatp {
    sitesSkyatp := SitesSkyatp{baseController: baseController}
    return &sitesSkyatp
}

// CountSiteSkyatpEvents takes context, siteId, distinct, mType, mac, deviceMac, threatLevel, ipAddress, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.RepsonseCount data and
// an error if there was an issue with the request or response.
// Count by Distinct Attributes of Skyatp Events (WIP)
func (s *SitesSkyatp) CountSiteSkyatpEvents(
    ctx context.Context,
    siteId uuid.UUID,
    distinct *models.SiteSkyAtpEventsCountDistinctEnum,
    mType *string,
    mac *string,
    deviceMac *string,
    threatLevel *int,
    ipAddress *string,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.RepsonseCount],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/skyatp/events/count", siteId),
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
    if distinct != nil {
        req.QueryParam("distinct", *distinct)
    }
    if mType != nil {
        req.QueryParam("type", *mType)
    }
    if mac != nil {
        req.QueryParam("mac", *mac)
    }
    if deviceMac != nil {
        req.QueryParam("device_mac", *deviceMac)
    }
    if threatLevel != nil {
        req.QueryParam("threat_level", *threatLevel)
    }
    if ipAddress != nil {
        req.QueryParam("ip_address", *ipAddress)
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

// SearchSiteSkyatpEvents takes context, siteId, mType, mac, deviceMac, threatLevel, ipAddress, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.ResponseEventsSkyAtpSearch data and
// an error if there was an issue with the request or response.
// Search Skyatp Events (WIP)
func (s *SitesSkyatp) SearchSiteSkyatpEvents(
    ctx context.Context,
    siteId uuid.UUID,
    mType *string,
    mac *string,
    deviceMac *string,
    threatLevel *int,
    ipAddress *string,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.ResponseEventsSkyAtpSearch],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/skyatp/events/search", siteId),
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
    if mType != nil {
        req.QueryParam("type", *mType)
    }
    if mac != nil {
        req.QueryParam("mac", *mac)
    }
    if deviceMac != nil {
        req.QueryParam("device_mac", *deviceMac)
    }
    if threatLevel != nil {
        req.QueryParam("threat_level", *threatLevel)
    }
    if ipAddress != nil {
        req.QueryParam("ip_address", *ipAddress)
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
    
    var result models.ResponseEventsSkyAtpSearch
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseEventsSkyAtpSearch](decoder)
    return models.NewApiResponse(result, resp), err
}
