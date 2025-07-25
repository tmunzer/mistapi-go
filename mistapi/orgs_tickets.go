// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/errors"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "net/http"
)

// OrgsTickets represents a controller struct.
type OrgsTickets struct {
    baseController
}

// NewOrgsTickets creates a new instance of OrgsTickets.
// It takes a baseController as a parameter and returns a pointer to the OrgsTickets.
func NewOrgsTickets(baseController baseController) *OrgsTickets {
    orgsTickets := OrgsTickets{baseController: baseController}
    return &orgsTickets
}

// ListOrgTickets takes context, orgId, start, end, duration as parameters and
// returns an models.ApiResponse with []models.Ticket data and
// an error if there was an issue with the request or response.
// Get List of Tickets of an Org
func (o *OrgsTickets) ListOrgTickets(
    ctx context.Context,
    orgId uuid.UUID,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[[]models.Ticket],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/tickets")
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
    if start != nil {
        req.QueryParam("start", *start)
    }
    if end != nil {
        req.QueryParam("end", *end)
    }
    if duration != nil {
        req.QueryParam("duration", *duration)
    }
    
    var result []models.Ticket
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.Ticket](decoder)
    return models.NewApiResponse(result, resp), err
}

// CreateOrgTicket takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.Ticket data and
// an error if there was an issue with the request or response.
// Create a support ticket
func (o *OrgsTickets) CreateOrgTicket(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.Ticket) (
    models.ApiResponse[models.Ticket],
    error) {
    req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/tickets")
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
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.Ticket
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Ticket](decoder)
    return models.NewApiResponse(result, resp), err
}

// CountOrgTickets takes context, orgId, distinct, limit as parameters and
// returns an models.ApiResponse with models.ResponseCount data and
// an error if there was an issue with the request or response.
// Count by Distinct Attributes of Org Tickets
func (o *OrgsTickets) CountOrgTickets(
    ctx context.Context,
    orgId uuid.UUID,
    distinct *models.OrgTicketsCountDistinctEnum,
    limit *int) (
    models.ApiResponse[models.ResponseCount],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/tickets/count")
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
    if distinct != nil {
        req.QueryParam("distinct", *distinct)
    }
    if limit != nil {
        req.QueryParam("limit", *limit)
    }
    
    var result models.ResponseCount
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseCount](decoder)
    return models.NewApiResponse(result, resp), err
}

// GetOrgTicket takes context, orgId, ticketId, start, end, duration as parameters and
// returns an models.ApiResponse with models.Ticket data and
// an error if there was an issue with the request or response.
// Get support ticket details
func (o *OrgsTickets) GetOrgTicket(
    ctx context.Context,
    orgId uuid.UUID,
    ticketId uuid.UUID,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.Ticket],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/tickets/%v")
    req.AppendTemplateParams(orgId, ticketId)
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
    if start != nil {
        req.QueryParam("start", *start)
    }
    if end != nil {
        req.QueryParam("end", *end)
    }
    if duration != nil {
        req.QueryParam("duration", *duration)
    }
    
    var result models.Ticket
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Ticket](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateOrgTicket takes context, orgId, ticketId, body as parameters and
// returns an models.ApiResponse with models.Ticket data and
// an error if there was an issue with the request or response.
// Update support ticket
func (o *OrgsTickets) UpdateOrgTicket(
    ctx context.Context,
    orgId uuid.UUID,
    ticketId uuid.UUID,
    body *models.Ticket) (
    models.ApiResponse[models.Ticket],
    error) {
    req := o.prepareRequest(ctx, "PUT", "/api/v1/orgs/%v/tickets/%v")
    req.AppendTemplateParams(orgId, ticketId)
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
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.Ticket
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Ticket](decoder)
    return models.NewApiResponse(result, resp), err
}

// UploadOrgTicketAttachment takes context, orgId, ticketId, file as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Get Org ticket Attachment
func (o *OrgsTickets) UploadOrgTicketAttachment(
    ctx context.Context,
    orgId uuid.UUID,
    ticketId uuid.UUID,
    file *models.FileWrapper) (
    *http.Response,
    error) {
    req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/tickets/%v/attachments")
    req.AppendTemplateParams(orgId, ticketId)
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
    formFields := []https.FormParam{}
    if file != nil {
        fileParam := https.FormParam{Key: "file", Value: *file, Headers: http.Header{}}
        formFields = append(formFields, fileParam)
    }
    req.FormData(formFields)
    
    httpCtx, err := req.Call()
    if err != nil {
        return httpCtx.Response, err
    }
    return httpCtx.Response, err
}

// GetOrgTicketAttachment takes context, orgId, ticketId, attachmentId, start, end, duration as parameters and
// returns an models.ApiResponse with models.TicketAttachment data and
// an error if there was an issue with the request or response.
// Get Org ticket Attachment
func (o *OrgsTickets) GetOrgTicketAttachment(
    ctx context.Context,
    orgId uuid.UUID,
    ticketId uuid.UUID,
    attachmentId uuid.UUID,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.TicketAttachment],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/tickets/%v/attachments/%v")
    req.AppendTemplateParams(orgId, ticketId, attachmentId)
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
    if start != nil {
        req.QueryParam("start", *start)
    }
    if end != nil {
        req.QueryParam("end", *end)
    }
    if duration != nil {
        req.QueryParam("duration", *duration)
    }
    
    var result models.TicketAttachment
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.TicketAttachment](decoder)
    return models.NewApiResponse(result, resp), err
}

// AddOrgTicketComment takes context, orgId, ticketId, comment, file as parameters and
// returns an models.ApiResponse with models.Ticket data and
// an error if there was an issue with the request or response.
// Add Comment to support ticket
func (o *OrgsTickets) AddOrgTicketComment(
    ctx context.Context,
    orgId uuid.UUID,
    ticketId uuid.UUID,
    comment *string,
    file *models.FileWrapper) (
    models.ApiResponse[models.Ticket],
    error) {
    req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/tickets/%v/comments")
    req.AppendTemplateParams(orgId, ticketId)
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
    formFields := []https.FormParam{}
    if comment != nil {
        commentParam := https.FormParam{Key: "comment", Value: *comment, Headers: http.Header{}}
        formFields = append(formFields, commentParam)
    }
    if file != nil {
        fileParam := https.FormParam{Key: "file", Value: *file, Headers: http.Header{}}
        formFields = append(formFields, fileParam)
    }
    req.FormData(formFields)
    
    var result models.Ticket
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Ticket](decoder)
    return models.NewApiResponse(result, resp), err
}
