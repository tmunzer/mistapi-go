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

// SitesDevicesWiredVirtualChassis represents a controller struct.
type SitesDevicesWiredVirtualChassis struct {
    baseController
}

// NewSitesDevicesWiredVirtualChassis creates a new instance of SitesDevicesWiredVirtualChassis.
// It takes a baseController as a parameter and returns a pointer to the SitesDevicesWiredVirtualChassis.
func NewSitesDevicesWiredVirtualChassis(baseController baseController) *SitesDevicesWiredVirtualChassis {
    sitesDevicesWiredVirtualChassis := SitesDevicesWiredVirtualChassis{baseController: baseController}
    return &sitesDevicesWiredVirtualChassis
}

// DeleteSiteVirtualChassis takes context, siteId, deviceId as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// When all the member switches of VC are removed and only member ID 0 is left, the cloud would detect this situation and automatically changes the single switch to non-VC role.
// For some unexpected cases that the VC is gone and disconncted, the API below could be used to change the state of VC’s switches to be standalone. After it is executed, all the switches will be shown as standalone switches under Inventory.
func (s *SitesDevicesWiredVirtualChassis) DeleteSiteVirtualChassis(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID) (
    *http.Response,
    error) {
    req := s.prepareRequest(
      ctx,
      "DELETE",
      fmt.Sprintf("/api/v1/sites/%v/devices/%v/vc", siteId, deviceId),
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

// GetSiteDeviceVirtualChassis takes context, siteId, deviceId as parameters and
// returns an models.ApiResponse with models.ResponseVirtualChassisConfig data and
// an error if there was an issue with the request or response.
// Get VC Status
// The API returns a combined view of the VC status which includes topology and stats_
func (s *SitesDevicesWiredVirtualChassis) GetSiteDeviceVirtualChassis(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID) (
    models.ApiResponse[models.ResponseVirtualChassisConfig],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/devices/%v/vc", siteId, deviceId),
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
    
    var result models.ResponseVirtualChassisConfig
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseVirtualChassisConfig](decoder)
    return models.NewApiResponse(result, resp), err
}

// CreateSiteVirtualChassis takes context, siteId, deviceId, body as parameters and
// returns an models.ApiResponse with models.ResponseVirtualChassisConfig data and
// an error if there was an issue with the request or response.
// For models (e.g. EX3400 and up) having dedicated VC ports, it is easier to form a VC by just connecting cables with the dedicated VC ports. Cloud will detect the new VC and update the inventory.
// In case that the user would like to choose the dedicated switch as a VC master. Or for EX2300-C-12P and EX2300-C-12T which doesn’t have dedicated VC ports, below are procedures to automate the VC creation:
// 1. Power on the switch that is choosen as the VC master first. And the powering on the other member switches.
// 2. Claim or adopt all these switches under the same organization’s Inventory
// 3. Assign these switches into the same Site
// 4. Invoke vc command on the switch choosen to be the VC master. For EX2300-C-12P, VC ports will be created automatically.
// 5. Connect the cables to the VC ports for these switches
// 6. Wait for the VC to be formed. The Org’s inventory will be updated for the new VC.
func (s *SitesDevicesWiredVirtualChassis) CreateSiteVirtualChassis(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.VirtualChassisConfig) (
    models.ApiResponse[models.ResponseVirtualChassisConfig],
    error) {
    req := s.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/devices/%v/vc", siteId, deviceId),
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
    
    var result models.ResponseVirtualChassisConfig
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseVirtualChassisConfig](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateSiteVirtualChassisMember takes context, siteId, deviceId, body as parameters and
// returns an models.ApiResponse with models.ResponseVirtualChassisConfig data and
// an error if there was an issue with the request or response.
// The VC creation and adding member switch API will update the device’s virtual chassis config which is applied after VC is formed to create JUNOS pre-provisioned virtual chassis configuration.
// ## Change to use preprovisioned VC
// To switch the VC to use preprovisioned VC, enable preprovisioned in virtual_chassis config. Both vc_role master and backup will be matched to routing-engine role in Junos preprovisioned VC config.
// In this config, fpc0 has to be the same as the mac of device_id. Use renumber if you want to replace fpc0 which involves device_id change.
// Notice: to configure preprovisioned VC, every member of the VC must be in the inventory.
// ## Add new members
// For models (e.g. EX4300 and up) having dedicated VC ports, it is easier to add new member switches into a VC by just connecting cables with the dedicated VC ports. Cloud will detect the new members and update the inventory.
// For EX2300 VC, adding new members requires to follow the procedures below:
// 1. Powering on the new member switches and ensuring cables are not connected to any VC ports.
// 2. Claim or adopt all new member switches under the VC’s organization Inventory
// 3. Assign all new member switches to the same Site as the VC
// 4. Invoke vc command to add switches to the VC.
// 5. Connect the cables to the VC ports for these switches
// 6. After a while, the Org’s Inventory shows this new switches has been added into the VC.
// ## Removing member switch
// To remove a member switch from the VC, following the procedures below:
// 1. Ensuring the VC is connected to the cloud first
// 2. Unplug the cable from the VC port of the switch
// 3. Waiting for the VC state (vc_state) of this switch is changed to not-present
// 4. Invoke update_vc with remove to remove this switch from the VC
// 5. The Org’s Inventory shows the switch is removed.
// Please notice that member ID 0 (fpc0) cannot be removed. When a VC has two switches left, unpluging the cable may result in the situation that fpc0 becomes a line card (LC). When this situation is happened, please re-plug in the cable, wait for both switches becoming present (show virtual-chassis) and then removing the cable again.
// ## Renumber a member switch
// When a member switch doesn’t work properly and needed to be replaced, the renumber API could be used. The following two types of renumber are supported:
// 1. Replace a non-fpc0 member switch
// 2. Replace fpc0. When fpc0 is relaced, PAPI device config and JUNOS config will be both updated.
// For renumber to work, the following procedures are needed: 
// 1. Ensuring the VC is connected to the cloud and the state of the member switch to be replaced must be non present. 
// 2. Adding the new member switch to the VC 
// 3. Waiting for the VC state (vc_state) of this VC to be updated to API server 
// 4. Invoke vc with renumber to replace the new member switch from fpc X to
// ## Perprovision VC members
// By specifying “preprovision” op, you can convert the current VC to pre-provisioned mode, update VC members as well as specify vc_ports when adding new members for device models without dedicated vc ports. Use renumber for fpc0 replacement which involves device_id change.
// Note: 
// 1. vc_ports is used for adding new members and not needed if * the device model has dedicated vc ports, or * no new member is added 
// 2. New VC members to be added should exist in the same Site as the VC
// Update Device’s VC config can achieve similar purpose by directly modifying current virtual_chassis config. However, it cannot fulfill requests to enabling vc_ports on new members that are yet to belong to current VC.
func (s *SitesDevicesWiredVirtualChassis) UpdateSiteVirtualChassisMember(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.VirtualChassisUpdate) (
    models.ApiResponse[models.ResponseVirtualChassisConfig],
    error) {
    req := s.prepareRequest(
      ctx,
      "PUT",
      fmt.Sprintf("/api/v1/sites/%v/devices/%v/vc", siteId, deviceId),
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
    
    var result models.ResponseVirtualChassisConfig
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseVirtualChassisConfig](decoder)
    return models.NewApiResponse(result, resp), err
}

// SetSiteVcPort takes context, siteId, deviceId, body as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Set VC port
func (s *SitesDevicesWiredVirtualChassis) SetSiteVcPort(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.VirtualChassisPort) (
    *http.Response,
    error) {
    req := s.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/devices/%v/vc/vc_port", siteId, deviceId),
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
