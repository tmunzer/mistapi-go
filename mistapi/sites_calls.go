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

// SitesCalls represents a controller struct.
type SitesCalls struct {
    baseController
}

// NewSitesCalls creates a new instance of SitesCalls.
// It takes a baseController as a parameter and returns a pointer to the SitesCalls.
func NewSitesCalls(baseController baseController) *SitesCalls {
    sitesCalls := SitesCalls{baseController: baseController}
    return &sitesCalls
}

// TroubleshootSiteCall takes context, siteId, clientMac, meetingId, mac, app, start, end, duration, limit, page as parameters and
// returns an models.ApiResponse with models.CallTroubleshoot data and
// an error if there was an issue with the request or response.
// Troubleshoot a call
func (s *SitesCalls) TroubleshootSiteCall(
    ctx context.Context,
    siteId uuid.UUID,
    clientMac string,
    meetingId string,
    mac *string,
    app *string,
    start *int,
    end *int,
    duration *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.CallTroubleshoot],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/stats/calls/client/%v/troubleshoot", siteId, clientMac),
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
    req.QueryParam("meeting_id", meetingId)
    if mac != nil {
        req.QueryParam("mac", *mac)
    }
    if app != nil {
        req.QueryParam("app", *app)
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
    
    var result models.CallTroubleshoot
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.CallTroubleshoot](decoder)
    return models.NewApiResponse(result, resp), err
}

// CountSiteCalls takes context, siteId, distrinct, rating, app, start, end as parameters and
// returns an models.ApiResponse with models.RepsonseCount data and
// an error if there was an issue with the request or response.
// Count by Distinct Attributes of Calls
func (s *SitesCalls) CountSiteCalls(
    ctx context.Context,
    siteId uuid.UUID,
    distrinct *models.CountSiteCallsDistrinctEnum,
    rating *int,
    app *string,
    start *string,
    end *string) (
    models.ApiResponse[models.RepsonseCount],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/stats/calls/count", siteId),
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
    if distrinct != nil {
        req.QueryParam("distrinct", *distrinct)
    }
    if rating != nil {
        req.QueryParam("rating", *rating)
    }
    if app != nil {
        req.QueryParam("app", *app)
    }
    if start != nil {
        req.QueryParam("start", *start)
    }
    if end != nil {
        req.QueryParam("end", *end)
    }
    
    var result models.RepsonseCount
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.RepsonseCount](decoder)
    return models.NewApiResponse(result, resp), err
}

// SearchSiteCalls takes context, siteId, mac, app, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.ResponseStatsCalls data and
// an error if there was an issue with the request or response.
// Search Calls
func (s *SitesCalls) SearchSiteCalls(
    ctx context.Context,
    siteId uuid.UUID,
    mac *string,
    app *string,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.ResponseStatsCalls],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/stats/calls/search", siteId),
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
    if mac != nil {
        req.QueryParam("mac", *mac)
    }
    if app != nil {
        req.QueryParam("app", *app)
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
    
    var result models.ResponseStatsCalls
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseStatsCalls](decoder)
    return models.NewApiResponse(result, resp), err
}

// ListSiteTroubleshootCalls takes context, siteId, clientMac, ap, meetingId, mac, app, start, end, duration, limit, page as parameters and
// returns an models.ApiResponse with models.CallTroubleshootSummary data and
// an error if there was an issue with the request or response.
// Summary of calls troubleshoot by site
func (s *SitesCalls) ListSiteTroubleshootCalls(
    ctx context.Context,
    siteId uuid.UUID,
    clientMac string,
    ap *string,
    meetingId *string,
    mac *string,
    app *string,
    start *int,
    end *int,
    duration *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.CallTroubleshootSummary],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/stats/calls/troubleshoot", siteId),
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
    if ap != nil {
        req.QueryParam("ap", *ap)
    }
    if meetingId != nil {
        req.QueryParam("meeting_id", *meetingId)
    }
    if mac != nil {
        req.QueryParam("mac", *mac)
    }
    if app != nil {
        req.QueryParam("app", *app)
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
    
    var result models.CallTroubleshootSummary
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.CallTroubleshootSummary](decoder)
    return models.NewApiResponse(result, resp), err
}
