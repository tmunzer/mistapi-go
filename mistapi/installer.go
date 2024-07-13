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

// Installer represents a controller struct.
type Installer struct {
    baseController
}

// NewInstaller creates a new instance of Installer.
// It takes a baseController as a parameter and returns a pointer to the Installer.
func NewInstaller(baseController baseController) *Installer {
    installer := Installer{baseController: baseController}
    return &installer
}

// ListInstallerAlarmTemplates takes context, orgId as parameters and
// returns an models.ApiResponse with []models.InstallersItem data and
// an error if there was an issue with the request or response.
// Get List of alarm templates
func (i *Installer) ListInstallerAlarmTemplates(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[[]models.InstallersItem],
    error) {
    req := i.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/installer/orgs/%v/alarmtemplates", orgId),
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
    
    var result []models.InstallersItem
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.InstallersItem](decoder)
    return models.NewApiResponse(result, resp), err
}

// ListInstallerDeviceProfiles takes context, orgId, mType as parameters and
// returns an models.ApiResponse with []models.InstallersItem data and
// an error if there was an issue with the request or response.
// Get List of Device Profiles
func (i *Installer) ListInstallerDeviceProfiles(
    ctx context.Context,
    orgId uuid.UUID,
    mType *models.DeviceTypeEnum) (
    models.ApiResponse[[]models.InstallersItem],
    error) {
    req := i.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/installer/orgs/%v/deviceprofiles", orgId),
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
    
    var result []models.InstallersItem
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.InstallersItem](decoder)
    return models.NewApiResponse(result, resp), err
}

// ListInstallerListOfRenctlyClaimedDevices takes context, orgId, model, siteName, siteId, limit, page as parameters and
// returns an models.ApiResponse with []models.InstallerDevice data and
// an error if there was an issue with the request or response.
// Get List of recently claimed devices
func (i *Installer) ListInstallerListOfRenctlyClaimedDevices(
    ctx context.Context,
    orgId uuid.UUID,
    model *string,
    siteName *string,
    siteId *uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.InstallerDevice],
    error) {
    req := i.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/installer/orgs/%v/devices", orgId),
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
    if model != nil {
        req.QueryParam("model", *model)
    }
    if siteName != nil {
        req.QueryParam("site_name", *siteName)
    }
    if siteId != nil {
        req.QueryParam("site_id", *siteId)
    }
    if limit != nil {
        req.QueryParam("limit", *limit)
    }
    if page != nil {
        req.QueryParam("page", *page)
    }
    
    var result []models.InstallerDevice
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.InstallerDevice](decoder)
    return models.NewApiResponse(result, resp), err
}

// ClaimInstallerDevices takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.ResponseInventory data and
// an error if there was an issue with the request or response.
// This mirrors `POST /api/v1/orgs/{org_id}/inventory` (see Inventory API)
func (i *Installer) ClaimInstallerDevices(
    ctx context.Context,
    orgId uuid.UUID,
    body []string) (
    models.ApiResponse[models.ResponseInventory],
    error) {
    req := i.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/installer/orgs/%v/devices", orgId),
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
        "400": {Message: "Bad Request if none of the entries are valid"},
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

// UnassignInstallerRecentlyClaimedDevice takes context, orgId, deviceMac as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Unassign recently claimed devices
func (i *Installer) UnassignInstallerRecentlyClaimedDevice(
    ctx context.Context,
    orgId uuid.UUID,
    deviceMac string) (
    *http.Response,
    error) {
    req := i.prepareRequest(
      ctx,
      "DELETE",
      fmt.Sprintf("/api/v1/installer/orgs/%v/devices/%v", orgId, deviceMac),
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

// ProvisionInstallerDevices takes context, orgId, deviceMac, body as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Provision or Replace a device 
// If replacing_mac is in the request payload, other attributes are ignored, we attempt to replace existing device (with mac replacing_mac) with the inventory device being configured. The replacement device must be in the inventory but not assigned, and the replacing_mac device must be assigned to a site, and satisfy grace period requirements. The Device replaced will become unassigned.
func (i *Installer) ProvisionInstallerDevices(
    ctx context.Context,
    orgId uuid.UUID,
    deviceMac string,
    body *models.InstallerProvisionDevice) (
    *http.Response,
    error) {
    req := i.prepareRequest(
      ctx,
      "PUT",
      fmt.Sprintf("/api/v1/installer/orgs/%v/devices/%v", orgId, deviceMac),
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
        "400": {Message: "Bad Request", Unmarshaller: errors.NewResponseDetailString},
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
        "404": {Message: "Not Found", Unmarshaller: errors.NewResponseDetailString},
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

// StartInstallerLocateDevice takes context, orgId, deviceMac as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Locate a Device by blinking it’s LED, it’s a persisted state that has to be stopped by calling Stop Locating API
func (i *Installer) StartInstallerLocateDevice(
    ctx context.Context,
    orgId uuid.UUID,
    deviceMac string) (
    *http.Response,
    error) {
    req := i.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/installer/orgs/%v/devices/%v/locate", orgId, deviceMac),
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

// StopInstallerLocateDevice takes context, orgId, deviceMac as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Stop it
func (i *Installer) StopInstallerLocateDevice(
    ctx context.Context,
    orgId uuid.UUID,
    deviceMac string) (
    *http.Response,
    error) {
    req := i.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/installer/orgs/%v/devices/%v/unlocate", orgId, deviceMac),
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

// DeleteInstallerDeviceImage takes context, orgId, imageName, deviceMac as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// delete image
func (i *Installer) DeleteInstallerDeviceImage(
    ctx context.Context,
    orgId uuid.UUID,
    imageName string,
    deviceMac string) (
    *http.Response,
    error) {
    req := i.prepareRequest(
      ctx,
      "DELETE",
      fmt.Sprintf("/api/v1/installer/orgs/%v/devices/%v/%v", orgId, deviceMac, imageName),
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

// AddInstallerDeviceImage takes context, orgId, imageName, deviceMac, autoDeviceprofileAssignment, csv, file, json as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Add image
func (i *Installer) AddInstallerDeviceImage(
    ctx context.Context,
    orgId uuid.UUID,
    imageName string,
    deviceMac string,
    autoDeviceprofileAssignment *bool,
    csv *models.FileWrapper,
    file *models.FileWrapper,
    json *models.MapImportJson) (
    *http.Response,
    error) {
    req := i.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/installer/orgs/%v/devices/%v/%v", orgId, deviceMac, imageName),
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
    req.Header("Content-Type", "multipart/form_data")
    formFields := []https.FormParam{}
    if autoDeviceprofileAssignment != nil {
        auto_deviceprofile_assignmentParam := https.FormParam{Key: "auto_deviceprofile_assignment", Value: *autoDeviceprofileAssignment, Headers: http.Header{}}
        formFields = append(formFields, auto_deviceprofile_assignmentParam)
    }
    if csv != nil {
        csvParam := https.FormParam{Key: "csv", Value: *csv, Headers: http.Header{}}
        formFields = append(formFields, csvParam)
    }
    if file != nil {
        fileParam := https.FormParam{Key: "file", Value: *file, Headers: http.Header{}}
        formFields = append(formFields, fileParam)
    }
    if json != nil {
        jsonParam := https.FormParam{Key: "json", Value: *json, Headers: http.Header{}}
        formFields = append(formFields, jsonParam)
    }
    req.FormData(formFields)
    
    context, err := req.Call()
    if err != nil {
        return context.Response, err
    }
    return context.Response, err
}

// GetInstallerDeviceVirtualChassis takes context, orgId, fpc0Mac as parameters and
// returns an models.ApiResponse with models.ResponseVirtualChassisConfig data and
// an error if there was an issue with the request or response.
// Get VC Status
// The API returns a combined view of the VC status which includes topology and stats
func (i *Installer) GetInstallerDeviceVirtualChassis(
    ctx context.Context,
    orgId uuid.UUID,
    fpc0Mac string) (
    models.ApiResponse[models.ResponseVirtualChassisConfig],
    error) {
    req := i.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/installer/orgs/%v/devices/%v/vc", orgId, fpc0Mac),
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

// CreateInstallerVirtualChassis takes context, orgId, fpc0Mac, body as parameters and
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
func (i *Installer) CreateInstallerVirtualChassis(
    ctx context.Context,
    orgId uuid.UUID,
    fpc0Mac string,
    body *models.VirtualChassisConfig) (
    models.ApiResponse[models.ResponseVirtualChassisConfig],
    error) {
    req := i.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/installer/orgs/%v/devices/%v/vc", orgId, fpc0Mac),
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

// UpdateInstallerVirtualChassisMember takes context, orgId, fpc0Mac, body as parameters and
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
func (i *Installer) UpdateInstallerVirtualChassisMember(
    ctx context.Context,
    orgId uuid.UUID,
    fpc0Mac string,
    body *models.VirtualChassisUpdate) (
    models.ApiResponse[models.ResponseVirtualChassisConfig],
    error) {
    req := i.prepareRequest(
      ctx,
      "PUT",
      fmt.Sprintf("/api/v1/installer/orgs/%v/devices/%v/vc", orgId, fpc0Mac),
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

// ListInstallerRfTemplatesNames takes context, orgId as parameters and
// returns an models.ApiResponse with []models.InstallersItem data and
// an error if there was an issue with the request or response.
// Get List of RF Templates
func (i *Installer) ListInstallerRfTemplatesNames(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[[]models.InstallersItem],
    error) {
    req := i.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/installer/orgs/%v/rftemplates", orgId),
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
    
    var result []models.InstallersItem
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.InstallersItem](decoder)
    return models.NewApiResponse(result, resp), err
}

// ListInstallerSiteGroups takes context, orgId as parameters and
// returns an models.ApiResponse with []models.InstallersItem data and
// an error if there was an issue with the request or response.
// Get List of Site Groups
func (i *Installer) ListInstallerSiteGroups(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[[]models.InstallersItem],
    error) {
    req := i.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/installer/orgs/%v/sitegroups", orgId),
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
    
    var result []models.InstallersItem
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.InstallersItem](decoder)
    return models.NewApiResponse(result, resp), err
}

// ListInstallerSites takes context, orgId as parameters and
// returns an models.ApiResponse with []models.InstallerSite data and
// an error if there was an issue with the request or response.
// Get List of Sites
func (i *Installer) ListInstallerSites(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[[]models.InstallerSite],
    error) {
    req := i.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/installer/orgs/%v/sites", orgId),
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
    
    var result []models.InstallerSite
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.InstallerSite](decoder)
    return models.NewApiResponse(result, resp), err
}

// CreateOrUpdateInstallerSites takes context, orgId, siteName, body as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Often the Installers are asked to assign Devices to Sites. The Sites can either be pre-created or created/modified by the Installer. If this is an update, the same grace period also applies.
func (i *Installer) CreateOrUpdateInstallerSites(
    ctx context.Context,
    orgId uuid.UUID,
    siteName string,
    body *models.InstallerSite) (
    *http.Response,
    error) {
    req := i.prepareRequest(
      ctx,
      "PUT",
      fmt.Sprintf("/api/v1/installer/orgs/%v/sites/%v", orgId, siteName),
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

// ListInstallerMaps takes context, orgId, siteName as parameters and
// returns an models.ApiResponse with []models.Map data and
// an error if there was an issue with the request or response.
// Get List of Maps
func (i *Installer) ListInstallerMaps(
    ctx context.Context,
    orgId uuid.UUID,
    siteName string) (
    models.ApiResponse[[]models.Map],
    error) {
    req := i.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/installer/orgs/%v/sites/%v/maps", orgId, siteName),
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
    
    var result []models.Map
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.Map](decoder)
    return models.NewApiResponse(result, resp), err
}

// ImportInstallerMap takes context, orgId, siteName, autoDeviceprofileAssignment, csv, file, json as parameters and
// returns an models.ApiResponse with models.ResponseMapImport data and
// an error if there was an issue with the request or response.
// Import data from files is a multipart POST which has an file, an optional json, and an optional csv, to create floorplan, assign & place ap if name or mac matches
func (i *Installer) ImportInstallerMap(
    ctx context.Context,
    orgId uuid.UUID,
    siteName string,
    autoDeviceprofileAssignment *bool,
    csv *models.FileWrapper,
    file *models.FileWrapper,
    json *models.MapImportJson) (
    models.ApiResponse[models.ResponseMapImport],
    error) {
    req := i.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/installer/orgs/%v/sites/%v/maps/import", orgId, siteName),
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
    req.Header("Content-Type", "multipart/form_data")
    formFields := []https.FormParam{}
    if autoDeviceprofileAssignment != nil {
        auto_deviceprofile_assignmentParam := https.FormParam{Key: "auto_deviceprofile_assignment", Value: *autoDeviceprofileAssignment, Headers: http.Header{}}
        formFields = append(formFields, auto_deviceprofile_assignmentParam)
    }
    if csv != nil {
        csvParam := https.FormParam{Key: "csv", Value: *csv, Headers: http.Header{}}
        formFields = append(formFields, csvParam)
    }
    if file != nil {
        fileParam := https.FormParam{Key: "file", Value: *file, Headers: http.Header{}}
        formFields = append(formFields, fileParam)
    }
    if json != nil {
        jsonParam := https.FormParam{Key: "json", Value: *json, Headers: http.Header{}}
        formFields = append(formFields, jsonParam)
    }
    req.FormData(formFields)
    
    var result models.ResponseMapImport
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseMapImport](decoder)
    return models.NewApiResponse(result, resp), err
}

// DeleteInstallerMap takes context, orgId, siteName, mapId as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Delete Map
func (i *Installer) DeleteInstallerMap(
    ctx context.Context,
    orgId uuid.UUID,
    siteName string,
    mapId uuid.UUID) (
    *http.Response,
    error) {
    req := i.prepareRequest(
      ctx,
      "DELETE",
      fmt.Sprintf("/api/v1/installer/orgs/%v/sites/%v/maps/%v", orgId, siteName, mapId),
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

// CreateInstallerMap takes context, orgId, siteName, mapId, body as parameters and
// returns an models.ApiResponse with models.Map data and
// an error if there was an issue with the request or response.
// Create a MAP
func (i *Installer) CreateInstallerMap(
    ctx context.Context,
    orgId uuid.UUID,
    siteName string,
    mapId uuid.UUID,
    body *models.Map) (
    models.ApiResponse[models.Map],
    error) {
    req := i.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/installer/orgs/%v/sites/%v/maps/%v", orgId, siteName, mapId),
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
    
    var result models.Map
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Map](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateInstallerMap takes context, orgId, siteName, mapId, body as parameters and
// returns an models.ApiResponse with models.Map data and
// an error if there was an issue with the request or response.
// Update map
func (i *Installer) UpdateInstallerMap(
    ctx context.Context,
    orgId uuid.UUID,
    siteName string,
    mapId uuid.UUID,
    body *models.Map) (
    models.ApiResponse[models.Map],
    error) {
    req := i.prepareRequest(
      ctx,
      "PUT",
      fmt.Sprintf("/api/v1/installer/orgs/%v/sites/%v/maps/%v", orgId, siteName, mapId),
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
    
    var result models.Map
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Map](decoder)
    return models.NewApiResponse(result, resp), err
}

// OptimizeInstallerRrm takes context, siteName as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// After installation is considered complete (APs are placed on maps, all powered up), you can trigger an optimize operation where RRM will kick in (and maybe other things in the future) before it’s automatically scheduled.
func (i *Installer) OptimizeInstallerRrm(
    ctx context.Context,
    siteName string) (
    *http.Response,
    error) {
    req := i.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/installer/sites/%v/optimize", siteName),
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
