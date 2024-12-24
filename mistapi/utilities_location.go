package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/errors"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "net/http"
)

// UtilitiesLocation represents a controller struct.
type UtilitiesLocation struct {
    baseController
}

// NewUtilitiesLocation creates a new instance of UtilitiesLocation.
// It takes a baseController as a parameter and returns a pointer to the UtilitiesLocation.
func NewUtilitiesLocation(baseController baseController) *UtilitiesLocation {
    utilitiesLocation := UtilitiesLocation{baseController: baseController}
    return &utilitiesLocation
}

// SendSiteDevicesArbitratryBleBeacon takes context, siteId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Send arbitrary BLE Beacon for a period of time
// Note that only the devices that are connected will be restarted.
func (u *UtilitiesLocation) SendSiteDevicesArbitratryBleBeacon(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.UtilsSendBleBeacon) (
    *http.Response,
    error) {
    req := u.prepareRequest(ctx, "POST", "/api/v1/sites/%v/devices/send_ble_beacon")
    req.AppendTemplateParams(siteId)
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
