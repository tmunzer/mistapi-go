package mistapi

import (
    "context"
    "fmt"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/google/uuid"
    "mistapi/errors"
    "mistapi/models"
    "net/http"
)

// SitesAlarms represents a controller struct.
type SitesAlarms struct {
    baseController
}

// NewSitesAlarms creates a new instance of SitesAlarms.
// It takes a baseController as a parameter and returns a pointer to the SitesAlarms.
func NewSitesAlarms(baseController baseController) *SitesAlarms {
    sitesAlarms := SitesAlarms{baseController: baseController}
    return &sitesAlarms
}

// AckSiteMultipleAlarms takes context, siteId, body as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Ack multiple Site Alarms
func (s *SitesAlarms) AckSiteMultipleAlarms(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.AlarmAck) (
    *http.Response,
    error) {
    req := s.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/alarms/ack", siteId),
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
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    context, err := req.Call()
    if err != nil {
        return context.Response, err
    }
    return context.Response, err
}

// AckSiteAllAlarms takes context, siteId, body as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Ack all Site Alarms
// **N.B.**: Batch size for multiple alarm ack and unack has to be less or or equal to 1000.
func (s *SitesAlarms) AckSiteAllAlarms(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.NoteString) (
    *http.Response,
    error) {
    req := s.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/alarms/ack_all", siteId),
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
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    context, err := req.Call()
    if err != nil {
        return context.Response, err
    }
    return context.Response, err
}

// CountSiteAlarms takes context, siteId, distinct, ackAdminName, acked, mType, severity, group, page, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.RepsonseCount data and
// an error if there was an issue with the request or response.
// Count Site Alarms
func (s *SitesAlarms) CountSiteAlarms(
    ctx context.Context,
    siteId uuid.UUID,
    distinct *models.AlarmCountDisctinctEnum,
    ackAdminName *string,
    acked *bool,
    mType *string,
    severity *string,
    group *string,
    page *int,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.RepsonseCount],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/alarms/count", siteId),
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
    if ackAdminName != nil {
        req.QueryParam("ack_admin_name", *ackAdminName)
    }
    if acked != nil {
        req.QueryParam("acked", *acked)
    }
    if mType != nil {
        req.QueryParam("type", *mType)
    }
    if severity != nil {
        req.QueryParam("severity", *severity)
    }
    if group != nil {
        req.QueryParam("group", *group)
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

// SearchSiteAlarms takes context, siteId, mType, ackAdminName, acked, severity, group, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.AlarmSearchResult data and
// an error if there was an issue with the request or response.
// Search Site Alarms
func (s *SitesAlarms) SearchSiteAlarms(
    ctx context.Context,
    siteId uuid.UUID,
    mType *string,
    ackAdminName *string,
    acked *bool,
    severity *string,
    group *string,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.AlarmSearchResult],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/alarms/search", siteId),
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
    if ackAdminName != nil {
        req.QueryParam("ack_admin_name", *ackAdminName)
    }
    if acked != nil {
        req.QueryParam("acked", *acked)
    }
    if severity != nil {
        req.QueryParam("severity", *severity)
    }
    if group != nil {
        req.QueryParam("group", *group)
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
    
    var result models.AlarmSearchResult
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.AlarmSearchResult](decoder)
    return models.NewApiResponse(result, resp), err
}

// UnackSiteMultipleAlarms takes context, siteId, body as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Unack multiple Site Alarms
func (s *SitesAlarms) UnackSiteMultipleAlarms(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.AlarmAck) (
    *http.Response,
    error) {
    req := s.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/alarms/unack", siteId),
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
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    context, err := req.Call()
    if err != nil {
        return context.Response, err
    }
    return context.Response, err
}

// UnackSiteAllArlarms takes context, siteId, body as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Unack all Site Alarms
// **N.B.**: Batch size for multiple alarm ack and unack has to be less or or equal to 1000.
func (s *SitesAlarms) UnackSiteAllArlarms(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.NoteString) (
    *http.Response,
    error) {
    req := s.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/alarms/unack_all", siteId),
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
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    context, err := req.Call()
    if err != nil {
        return context.Response, err
    }
    return context.Response, err
}

// AckSiteAlarm takes context, siteId, alarmId, body as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Ack Site Alarm
func (s *SitesAlarms) AckSiteAlarm(
    ctx context.Context,
    siteId uuid.UUID,
    alarmId uuid.UUID,
    body *models.NoteString) (
    *http.Response,
    error) {
    req := s.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/alarms/%v/ack", siteId, alarmId),
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
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    context, err := req.Call()
    if err != nil {
        return context.Response, err
    }
    return context.Response, err
}

// UnackSiteAlarm takes context, siteId, alarmId, body as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Unack Site Alarm
func (s *SitesAlarms) UnackSiteAlarm(
    ctx context.Context,
    siteId uuid.UUID,
    alarmId uuid.UUID,
    body *models.NoteString) (
    *http.Response,
    error) {
    req := s.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/alarms/%v/unack", siteId, alarmId),
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
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    context, err := req.Call()
    if err != nil {
        return context.Response, err
    }
    return context.Response, err
}

// UnsubscribeSiteAlarms takes context, siteId as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Unsubscribe to Site Alarms
func (s *SitesAlarms) UnsubscribeSiteAlarms(
    ctx context.Context,
    siteId uuid.UUID) (
    *http.Response,
    error) {
    req := s.prepareRequest(
      ctx,
      "DELETE",
      fmt.Sprintf("/api/v1/sites/%v/subscriptions", siteId),
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
    
    context, err := req.Call()
    if err != nil {
        return context.Response, err
    }
    return context.Response, err
}

// SubscribeSiteAlarms takes context, siteId as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Subscribe to Site Alarms
func (s *SitesAlarms) SubscribeSiteAlarms(
    ctx context.Context,
    siteId uuid.UUID) (
    *http.Response,
    error) {
    req := s.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/subscriptions", siteId),
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
    
    context, err := req.Call()
    if err != nil {
        return context.Response, err
    }
    return context.Response, err
}
