package mistapi

import (
    "context"
    "fmt"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/errors"
    "github.com/tmunzer/mistapi-go/mistapi/models"
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
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
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
        "400": {Message: "OK - if any of entries are valid or there’s no errors", Unmarshaller: errors.NewResponseInventoryError},
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
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
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
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

// CountOrgInventory takes context, orgId, mType, distinct, limit, page as parameters and
// returns an models.ApiResponse with models.RepsonseCount data and
// an error if there was an issue with the request or response.
// Count in the Org Inventory
func (o *OrgsInventory) CountOrgInventory(
    ctx context.Context,
    orgId uuid.UUID,
    mType *models.DeviceTypeEnum,
    distinct *models.InventoryCountDistinctEnum,
    limit *int,
    page *int) (
    models.ApiResponse[models.RepsonseCount],
    error) {
    req := o.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/orgs/%v/inventory/count", orgId),
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
    if mType != nil {
        req.QueryParam("type", *mType)
    }
    if distinct != nil {
        req.QueryParam("distinct", *distinct)
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

// CreateOrgGatewayHaCluster takes context, orgId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Create HA Cluster from unassigned Gateways
func (o *OrgsInventory) CreateOrgGatewayHaCluster(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.HaClusterConfig) (
    *http.Response,
    error) {
    req := o.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/orgs/%v/inventory/create_ha_cluster", orgId),
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
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    httpCtx, err := req.Call()
    if err != nil {
        return httpCtx.Response, err
    }
    return httpCtx.Response, err
}

// DeleteOrgGatewayHaCluster takes context, orgId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete HA Cluster
// After HA cluster deleted, both of the nodes will be unassigned.
func (o *OrgsInventory) DeleteOrgGatewayHaCluster(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.HaClusterDelete) (
    *http.Response,
    error) {
    req := o.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/orgs/%v/inventory/delete_ha_cluster", orgId),
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
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    httpCtx, err := req.Call()
    if err != nil {
        return httpCtx.Response, err
    }
    return httpCtx.Response, err
}

// ReevaluateOrgAutoAssignment takes context, orgId as parameters and
// returns an *Response and
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
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
    })
    
    httpCtx, err := req.Call()
    if err != nil {
        return httpCtx.Response, err
    }
    return httpCtx.Response, err
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
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
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

// SearchOrgInventory takes context, orgId, mType, mac, vcMac, masterMac, siteId, serial, master, sku, version, status, text, limit, page as parameters and
// returns an models.ApiResponse with models.InventorySearch data and
// an error if there was an issue with the request or response.
// Search in the Org Inventory
func (o *OrgsInventory) SearchOrgInventory(
    ctx context.Context,
    orgId uuid.UUID,
    mType *models.DeviceTypeEnum,
    mac *string,
    vcMac *string,
    masterMac *string,
    siteId *string,
    serial *string,
    master *string,
    sku *string,
    version *string,
    status *string,
    text *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.InventorySearch],
    error) {
    req := o.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/orgs/%v/inventory/search", orgId),
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
    if mType != nil {
        req.QueryParam("type", *mType)
    }
    if mac != nil {
        req.QueryParam("mac", *mac)
    }
    if vcMac != nil {
        req.QueryParam("vc_mac", *vcMac)
    }
    if masterMac != nil {
        req.QueryParam("master_mac", *masterMac)
    }
    if siteId != nil {
        req.QueryParam("site_id", *siteId)
    }
    if serial != nil {
        req.QueryParam("serial", *serial)
    }
    if master != nil {
        req.QueryParam("master", *master)
    }
    if sku != nil {
        req.QueryParam("sku", *sku)
    }
    if version != nil {
        req.QueryParam("version", *version)
    }
    if status != nil {
        req.QueryParam("status", *status)
    }
    if text != nil {
        req.QueryParam("text", *text)
    }
    if limit != nil {
        req.QueryParam("limit", *limit)
    }
    if page != nil {
        req.QueryParam("page", *page)
    }
    
    var result models.InventorySearch
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.InventorySearch](decoder)
    return models.NewApiResponse(result, resp), err
}
