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

// UtilitiesPCAPs represents a controller struct.
type UtilitiesPCAPs struct {
    baseController
}

// NewUtilitiesPCAPs creates a new instance of UtilitiesPCAPs.
// It takes a baseController as a parameter and returns a pointer to the UtilitiesPCAPs.
func NewUtilitiesPCAPs(baseController baseController) *UtilitiesPCAPs {
    utilitiesPCAPs := UtilitiesPCAPs{baseController: baseController}
    return &utilitiesPCAPs
}

// ListOrgPacketCaptures takes context, orgId, page, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.ResponsePcapSearch data and
// an error if there was an issue with the request or response.
// Get List of Org  Packet Captures
func (u *UtilitiesPCAPs) ListOrgPacketCaptures(
    ctx context.Context,
    orgId uuid.UUID,
    page *int,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.ResponsePcapSearch],
    error) {
    req := u.prepareRequest(ctx, "GET", fmt.Sprintf("/api/v1/orgs/%v/pcaps", orgId))
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
    
    var result models.ResponsePcapSearch
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponsePcapSearch](decoder)
    return models.NewApiResponse(result, resp), err
}

// StopOrgPacketCapture takes context, orgId as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Stop current Org capture
func (u *UtilitiesPCAPs) StopOrgPacketCapture(
    ctx context.Context,
    orgId uuid.UUID) (
    *http.Response,
    error) {
    req := u.prepareRequest(
      ctx,
      "DELETE",
      fmt.Sprintf("/api/v1/orgs/%v/pcaps/capture", orgId),
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
    
    context, err := req.Call()
    if err != nil {
        return context.Response, err
    }
    return context.Response, err
}

// GetOrgCapturingStatus takes context, orgId as parameters and
// returns an models.ApiResponse with models.ResponsePcapStatus data and
// an error if there was an issue with the request or response.
// Get Org Capturing status
func (u *UtilitiesPCAPs) GetOrgCapturingStatus(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[models.ResponsePcapStatus],
    error) {
    req := u.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/orgs/%v/pcaps/capture", orgId),
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
    
    var result models.ResponsePcapStatus
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponsePcapStatus](decoder)
    return models.NewApiResponse(result, resp), err
}

// StartOrgPacketCapture takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.ResponsePcapStart data and
// an error if there was an issue with the request or response.
// Initiate a Packet Capture
// The output will be available through websocket. As there can be multiple command issued against the same AP at the same time and the output all goes through the same websocket stream, session is introduced for demux.
// #### Subscribe to Device Command outputs
// `WS /api-ws/v1/stream`
// ```json
// {
// subscribe: "/sites/{site_id}/pcaps"
// }
// ```
// #### Response (Wireless/RadioTap)
// ```json
// {
// "event": "data"
// "channel": "/orgs/67970e46-4e12-11e6-9188-0242ac110007/pcaps"
// "data": {
// "capture_id": "f039b1b4-a23e-48b2-906a-0da40524de73", 
// "pcap_dict": {
// "dst_mac": "68:ec:c5:09:2e:87",
// "src_mac": "8c:3b:ad:e0:47:40", 
// "vlan": 1, 
// "src_ip": "34.224.147.117", 
// "dst_ip": "192.168.1.55",
// "dst_port": 51635, 
// "src_port": 443,
// "protocol": "TCP", 
// "mxedge_id": "00000000-0000-0000-1000-001122334455",
// "direction": "tx", 
// "timestamp": 1652247615, 
// "length": 159.0, 
// "interface": "port0",
// "info": "1652247616.007409 IP ec2-34-224-147-117.compute-1.amazonaws.com.https > ip-192-168-1-55.ec2.internal.51635: Flags [P.], seq 
// 2192123968:2192124057, ack 4035166782, win 12, options [nop,nop,TS val 597467050 ecr 740580660], length 89\\n",
// }, 
// "pcap_raw": "1MOyoQIABAAAAAAAAAAAAP//AAABAAAAQEx7YhMzAACfAAAAnwAAAGjsxQkuh4w7reBHQIEAAAEIAEUAAI1bLEAAKAZ/CiLgk3XAqAE3AbvJs4KpKEDwg8I+gBgADFf9AAABAQgKI5yfqiwkXTQXAwMAVKY5JopoKQrVEn0/3ld4YntctGEH/rTZuwtCvzSncFw71QJveJi9uxHs57KC8w9Apph3YvXJrmWg7M37+o+YV0KH/xmr626s5Bkhb3QhKOu+NoNEmA==\"
// }
// }
// ```
func (u *UtilitiesPCAPs) StartOrgPacketCapture(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.CaptureOrg) (
    models.ApiResponse[models.ResponsePcapStart],
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/orgs/%v/pcaps/capture", orgId),
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
    
    var result models.ResponsePcapStart
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponsePcapStart](decoder)
    return models.NewApiResponse(result, resp), err
}

// ListSitePacketCaptures takes context, siteId, page, limit, start, end, duration, clientMac as parameters and
// returns an models.ApiResponse with models.ResponsePcapSearch data and
// an error if there was an issue with the request or response.
// Get List of Site Packet Captures
func (u *UtilitiesPCAPs) ListSitePacketCaptures(
    ctx context.Context,
    siteId uuid.UUID,
    page *int,
    limit *int,
    start *int,
    end *int,
    duration *string,
    clientMac *string) (
    models.ApiResponse[models.ResponsePcapSearch],
    error) {
    req := u.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/pcaps", siteId),
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
    if clientMac != nil {
        req.QueryParam("client_mac", *clientMac)
    }
    
    var result models.ResponsePcapSearch
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponsePcapSearch](decoder)
    return models.NewApiResponse(result, resp), err
}

// StopSitePacketCapture takes context, siteId as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Stop current capture
func (u *UtilitiesPCAPs) StopSitePacketCapture(
    ctx context.Context,
    siteId uuid.UUID) (
    *http.Response,
    error) {
    req := u.prepareRequest(
      ctx,
      "DELETE",
      fmt.Sprintf("/api/v1/sites/%v/pcaps/capture", siteId),
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
    
    context, err := req.Call()
    if err != nil {
        return context.Response, err
    }
    return context.Response, err
}

// GetSiteCapturingStatus takes context, siteId as parameters and
// returns an models.ApiResponse with models.ResponsePcapStatus data and
// an error if there was an issue with the request or response.
// Get Capturing status
func (u *UtilitiesPCAPs) GetSiteCapturingStatus(
    ctx context.Context,
    siteId uuid.UUID) (
    models.ApiResponse[models.ResponsePcapStatus],
    error) {
    req := u.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/pcaps/capture", siteId),
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
    
    var result models.ResponsePcapStatus
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponsePcapStatus](decoder)
    return models.NewApiResponse(result, resp), err
}

// StartSitePacketCapture takes context, siteId, body as parameters and
// returns an models.ApiResponse with models.ResponsePcapStart data and
// an error if there was an issue with the request or response.
// Initiate a Site Packet Capture
// The output will be available through websocket. As there can be multiple command issued against the same AP at the same time and the output all goes through the same websocket stream, session is introduced for demux.
// #### Subscribe to Device Command outputs
// `WS /api-ws/v1/stream`
// ```json
// {
// subscribe: "/sites/{site_id}/pcaps"
// }
// ```
// #### Response (MxEdge)
// ```json
// {
// "event": "data"
// "channel": "/sites/:site_id/pcaps"
// "data": {
// "capture_id": "6b1be4fb-b239-44d9-9d3b-cb1ff3af1721",
// "lost_messages": 0
// "pcap_dict": {
// "channel_frequency": 2412,
// "channel": "1",
// "datarate": "1.0 Mbps",
// "rssi": -75, 
// "dst": "78:bd:bc:ca:0b:0a",
// "src": "18:b8:1f:4c:91:c0",
// "bssid": "18:b8:1f:4c:91:c0",
// "frame_type": "Management", 
// "frame_subtype": "Probe Response", 
// "proto": "802.11", 
// "ap_mac": "d4:20:b0:81:99:2e", 
// "direction": "tx", 
// "timestamp": 1652246543, 
// "length": 416.0,
// "interface": "radiotap",
// "info": "1652246544.467733 1683216786us tsft 1.0 Mb/s 2412 MHz 11g -75dBm signal -82dBm noise antenna 0 Probe Response (ATTKmsWiVS) [1.0* 2.0* 5.5* 11.0* 18.0 24.0 36.0 54.0 Mbit] CH: 2, PRIVACY\\n",
// }, 
// "pcap_raw": "1MOyoQIABAAAAAAAAAAAAP//AAABAAAAEEh7Yh5VBwCgAQAAoAEAAAAAKwBvCADAAQAAAIw7reCS2VNkAAAAABACbAmABLWuAAEAEBgAAwACAABQADoBeL28ygsKGLgfTJHAGLgfTJHAcIZ2WDlBJQAAAGQAERUACkFUVEttc1dpVlMBCIKEi5YkMEhsAwECBwZVUyABCx4gAQAjAhkAKgEEMgQMEhhgMBQBAAAPrAQBAAAPrAQBAAAPrAIMAAsFAQAbAABGBTIIAQAALRqtCR////8AAAAAAAAAAAAAAAAAAAAAAAAAAD0WAggVAAAAAAAAAAAAAAAAAAAAAAAAAH8IBAAIAAAAAEDdkwBQ8gQQSgABEBBEAAECEDsAAQMQRwAQn2481frn3KT+uGod2ERx+RAhAAtBcnJpcywgSW5jLhAjAApCR1cyMTAtNzAwECQACkJHVzIxMC03MDAQQgAKQkdXMjEwLTcwMBBUAAgABgBQ8gQAARARAA5BcnJpcyBXaXJlbGVzcxAIAAIgCBA8AAEBEEkABgA3KgABIN0JABAYAgEQHAAA3RgAUPICAQGEAAOkAAAnpAAAQkNeAGIyLwAzjakr"
// }
// ```
// #### vResponse (Wired)
// ```json
// {
// "event": "data"
// "channel": "/sites/67970e46-4e12-11e6-9188-0242ac110007/pcaps"
// "data": {
// "capture_id": "f039b1b4-a23e-48b2-906a-0da40524de73", 
// "pcap_dict": {
// "dst_mac": "68:ec:c5:09:2e:87",
// "src_mac": "8c:3b:ad:e0:47:40", 
// "vlan": 1, 
// "src_ip": "34.224.147.117", 
// "dst_ip": "192.168.1.55",
// "dst_port": 51635, 
// "src_port": 443,
// "proto": "TCP", 
// "ap_mac": "d4:20:b0:81:99:2e",
// "direction": "tx", 
// "timestamp": 1652247615, 
// "length": 159.0, 
// "interface": "wired",
// "info": "1652247616.007409 IP ec2-34-224-147-117.compute-1.amazonaws.com.https > ip-192-168-1-55.ec2.internal.51635: Flags [P.], seq 2192123968:2192124057, ack 4035166782, win 12, options [nop,nop,TS val 597467050 ecr 740580660], length 89\\n",
// }, 
// "pcap_raw": "1MOyoQIABAAAAAAAAAAAAP//AAABAAAAQEx7YhMzAACfAAAAnwAAAGjsxQkuh4w7reBHQIEAAAEIAEUAAI1bLEAAKAZ/CiLgk3XAqAE3AbvJs4KpKEDwg8I+gBgADFf9AAABAQgKI5yfqiwkXTQXAwMAVKY5JopoKQrVEn0/3ld4YntctGEH/rTZuwtCvzSncFw71QJveJi9uxHs57KC8w9Apph3YvXJrmWg7M37+o+YV0KH/xmr626s5Bkhb3QhKOu+NoNEmA=="
// }
// }
// ```
// #### Stop Response (Wired/Wireless)
// ```json
// {
// "event": "data"
// "channel": "/sites/67970e46-4e12-11e6-9188-0242ac110007/pcaps"
// "data": {
// "capture_id": "a2f7374d-6a70-41fd-8a3f-71e42573baaf", 
// "lost_messages": 0,
// "pcap_dict": null
// }
// }
// ```
func (u *UtilitiesPCAPs) StartSitePacketCapture(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.CaptureSite) (
    models.ApiResponse[models.ResponsePcapStart],
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/pcaps/capture", siteId),
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
    
    var result models.ResponsePcapStart
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponsePcapStart](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateSitePacketCapture takes context, siteId, pcapId, body as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Update or add notes to a completed packet capture
func (u *UtilitiesPCAPs) UpdateSitePacketCapture(
    ctx context.Context,
    siteId uuid.UUID,
    pcapId uuid.UUID,
    body *models.NotesString) (
    *http.Response,
    error) {
    req := u.prepareRequest(
      ctx,
      "PUT",
      fmt.Sprintf("/api/v1/sites/%v/pcaps/%v", siteId, pcapId),
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
    
    context, err := req.Call()
    if err != nil {
        return context.Response, err
    }
    return context.Response, err
}
