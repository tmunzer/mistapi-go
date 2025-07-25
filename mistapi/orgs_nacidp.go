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
)

// OrgsNACIDP represents a controller struct.
type OrgsNACIDP struct {
    baseController
}

// NewOrgsNACIDP creates a new instance of OrgsNACIDP.
// It takes a baseController as a parameter and returns a pointer to the OrgsNACIDP.
func NewOrgsNACIDP(baseController baseController) *OrgsNACIDP {
    orgsNACIDP := OrgsNACIDP{baseController: baseController}
    return &orgsNACIDP
}

// ValidateOrgIdpCredential takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.WebsocketSession data and
// an error if there was an issue with the request or response.
// IDP Credential Validation. The output will be available through websocket. As there can be multiple command issued against the same device at the same time and the output all goes through the same websocket stream, `session` is introduced for demux.
// #### Subscribe to Device Command outputs
// `WS /api-ws/v1/stream`
// ``` json
// {
// "subscribe": "orgs/{org_id}/mist_nac/test_idp"
// }
// ```
// ### Response (no idp can be found)
// ``` json
// {
// "event": "data",
// "channel": "/orgs/{org_id}/mist_nac/test_idp",
// "status": 
// "data": {
// "status": "failure",
// "error": "No matching IDP found"
// }
// }
// ```
// ### Response OK
// ``` json
// {
// "event": "data",
// "channel": "/orgs/{org_id}/mist_nac/test_idp",
// "status": 
// "data": {
// "status": "success",
// "idp_id": "915793c0-1355-4e98-b1c0-23df2227b357",
// "idp_type": "ldap",
// // more attributes will be added later
// }
// }
// ```
// ### Response Invalid Credentials
// ``` json
// {
// "event": "data",
// "channel": "/orgs/{org_id}/mist_nac/test_idp",
// "status": 
// "data": {
// "status": "failure",
// "error": "Invalid Credentials",
// "idp_id": "915793c0-1355-4e98-b1c0-23df2227b357",
// "idp_type": "ldap",
// }
// }
// ```
func (o *OrgsNACIDP) ValidateOrgIdpCredential(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.UsernamePassword) (
    models.ApiResponse[models.WebsocketSession],
    error) {
    req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/mist_nac/test_idp")
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
    
    var result models.WebsocketSession
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.WebsocketSession](decoder)
    return models.NewApiResponse(result, resp), err
}
