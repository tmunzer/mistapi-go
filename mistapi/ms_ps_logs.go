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

// MSPsLogs represents a controller struct.
type MSPsLogs struct {
    baseController
}

// NewMSPsLogs creates a new instance of MSPsLogs.
// It takes a baseController as a parameter and returns a pointer to the MSPsLogs.
func NewMSPsLogs(baseController baseController) *MSPsLogs {
    mSPsLogs := MSPsLogs{baseController: baseController}
    return &mSPsLogs
}

// ListMspAuditLogs takes context, mspId, siteId, adminName, message, sort, start, end, limit, page, duration as parameters and
// returns an models.ApiResponse with models.ResponseLogSearch data and
// an error if there was an issue with the request or response.
// Get list of change logs for the current MSP
func (m *MSPsLogs) ListMspAuditLogs(
    ctx context.Context,
    mspId uuid.UUID,
    siteId *string,
    adminName *string,
    message *string,
    sort *models.ListMspLogsSortEnum,
    start *int,
    end *int,
    limit *int,
    page *int,
    duration *string) (
    models.ApiResponse[models.ResponseLogSearch],
    error) {
    req := m.prepareRequest(ctx, "GET", fmt.Sprintf("/api/v1/msps/%v/logs", mspId))
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
    if siteId != nil {
        req.QueryParam("site_id", *siteId)
    }
    if adminName != nil {
        req.QueryParam("admin_name", *adminName)
    }
    if message != nil {
        req.QueryParam("message", *message)
    }
    if sort != nil {
        req.QueryParam("sort", *sort)
    }
    if start != nil {
        req.QueryParam("start", *start)
    }
    if end != nil {
        req.QueryParam("end", *end)
    }
    if limit != nil {
        req.QueryParam("limit", *limit)
    }
    if page != nil {
        req.QueryParam("page", *page)
    }
    if duration != nil {
        req.QueryParam("duration", *duration)
    }
    
    var result models.ResponseLogSearch
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseLogSearch](decoder)
    return models.NewApiResponse(result, resp), err
}

// CountMspAuditLogs takes context, mspId, distinct as parameters and
// returns an models.ApiResponse with models.RepsonseCount data and
// an error if there was an issue with the request or response.
// Count by Distinct Attributes of Audit Logs
func (m *MSPsLogs) CountMspAuditLogs(
    ctx context.Context,
    mspId uuid.UUID,
    distinct *models.MspLogsCountDistinctEnum) (
    models.ApiResponse[models.RepsonseCount],
    error) {
    req := m.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/msps/%v/logs/count", mspId),
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
    
    var result models.RepsonseCount
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.RepsonseCount](decoder)
    return models.NewApiResponse(result, resp), err
}
