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

// OrgsInventory represents a controller struct.
type OrgsInventory struct {
    baseController
}

// NewOrgsInventory creates a new instance of OrgsInventory.
// It takes a baseController as a parameter and returns a pointer to the OrgsInventory.
func NewOrgsInventory(baseController baseController) *OrgsInventory {
    orgsInventory := OrgsInventory{baseController: baseController}
    return &orgsInventory
}

// GetOrgInventory takes context, orgId, serial, model, mType, mac, siteId, vcMac, vc, unassigned, limit, page as parameters and
// returns an models.ApiResponse with []models.Inventory data and
// an error if there was an issue with the request or response.
// Get Org Inventory
// ### VC (Virtual-Chassis) Management
// Ideally VC should be managed as a single device - where - one managed entity where config / monitoring is anchored against (with a stable identify MAC) - all members appears in the inventory
// In our implementation, we strive to achieve that without manual user configurations by 
// 1. during claim or adoption a VC, we require FPC0 to exist and will use its MAC as identify for the entire chassis
// 2. other VC members will be automatically populated when they’re all present
// The perceivable result is 
// 1. from `/sites/:site_id/stats/devices/:fpc0_mac` API, you’d see the VC where module_stat contains the VC members 
// 2. from `/orgs/:org_id/inventory?vc=true` API, you’d see all VC members with vc_mac pointing to the FPC0
func (o *OrgsInventory) GetOrgInventory(
    ctx context.Context,
    orgId uuid.UUID,
    serial *string,
    model *string,
    mType *models.DeviceTypeEnum,
    mac *string,
    siteId *string,
    vcMac *string,
    vc *bool,
    unassigned *bool,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.Inventory],
    error) {
    req := o.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/orgs/%v/inventory", orgId),
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
    if serial != nil {
        req.QueryParam("serial", *serial)
    }
    if model != nil {
        req.QueryParam("model", *model)
    }
    if mType != nil {
        req.QueryParam("type", *mType)
    }
    if mac != nil {
        req.QueryParam("mac", *mac)
    }
    if siteId != nil {
        req.QueryParam("site_id", *siteId)
    }
    if vcMac != nil {
        req.QueryParam("vc_mac", *vcMac)
    }
    if vc != nil {
        req.QueryParam("vc", *vc)
    }
    if unassigned != nil {
        req.QueryParam("unassigned", *unassigned)
    }
    if limit != nil {
        req.QueryParam("limit", *limit)
    }
    if page != nil {
        req.QueryParam("page", *page)
    }
    
    var result []models.Inventory
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.Inventory](decoder)
    return models.NewApiResponse(result, resp), err
}

// AddOrgInventory takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.ResponseInventory data and
// an error if there was an issue with the request or response.
// Add Device to Org Inventory with the device claim codes
func (o *OrgsInventory) AddOrgInventory(
    ctx context.Context,
    orgId uuid.UUID,
    body []string) (
    models.ApiResponse[models.ResponseInventory],
    error) {
    req := o.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/orgs/%v/inventory", orgId),
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
        "400": {Message: "if none of the entries are valid"},
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.ResponseInventory
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseInventory](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateOrgInventoryAssignment takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.ResponseOrgInventoryChange data and
// an error if there was an issue with the request or response.
// Update Org Inventory
func (o *OrgsInventory) UpdateOrgInventoryAssignment(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.InventoryUpdate) (
    models.ApiResponse[models.ResponseOrgInventoryChange],
    error) {
    req := o.prepareRequest(
      ctx,
      "PUT",
      fmt.Sprintf("/api/v1/orgs/%v/inventory", orgId),
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
    
    var result models.ResponseOrgInventoryChange
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseOrgInventoryChange](decoder)
    return models.NewApiResponse(result, resp), err
}

// ReevaluateOrgAutoAssignment takes context, orgId as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Reevaluate Auto Assignment
func (o *OrgsInventory) ReevaluateOrgAutoAssignment(
    ctx context.Context,
    orgId uuid.UUID) (
    *http.Response,
    error) {
    req := o.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/orgs/%v/inventory/reevaluate_auto_assignment", orgId),
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

// ReplaceOrgDevices takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.ResponseOrgInventoryChange data and
// an error if there was an issue with the request or response.
// It’s a common request we get from the customers. When a AP HW has problem and need a replacement, they would want to copy the existing attributes (Device Config) of this old AP to the new one. It can be done by providing the MAC of a device that’s currently in the inventory but not assigned. The Device replaced will become unassigned.
// This API also supports replacement of Mist Edges. This API copies device agnostic attributes from old Mist edge to new one.
// Mist manufactured Mist Edges will be reset to factory settings but will still be in Inventory.Brownfield or VM’s will be
// deleted from Inventory
// **Note:** For Gateway devices only like-for-like replacements (can only replace a SRX320 with a SRX320 and not some otehr model) are allowed.
func (o *OrgsInventory) ReplaceOrgDevices(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.ReplaceDevice) (
    models.ApiResponse[models.ResponseOrgInventoryChange],
    error) {
    req := o.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/orgs/%v/inventory/replace", orgId),
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
    
    var result models.ResponseOrgInventoryChange
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseOrgInventoryChange](decoder)
    return models.NewApiResponse(result, resp), err
}
