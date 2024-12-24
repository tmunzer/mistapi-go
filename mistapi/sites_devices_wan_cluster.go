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

// SitesDevicesWANCluster represents a controller struct.
type SitesDevicesWANCluster struct {
    baseController
}

// NewSitesDevicesWANCluster creates a new instance of SitesDevicesWANCluster.
// It takes a baseController as a parameter and returns a pointer to the SitesDevicesWANCluster.
func NewSitesDevicesWANCluster(baseController baseController) *SitesDevicesWANCluster {
    sitesDevicesWANCluster := SitesDevicesWANCluster{baseController: baseController}
    return &sitesDevicesWANCluster
}

// DeleteSiteDeviceHaCluster takes context, siteId, deviceId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete HA Cluster
func (s *SitesDevicesWANCluster) DeleteSiteDeviceHaCluster(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID) (
    *http.Response,
    error) {
    req := s.prepareRequest(ctx, "DELETE", "/api/v1/sites/%v/devices/%v/ha")
    req.AppendTemplateParams(siteId, deviceId)
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

// GetSiteDeviceHaClusterNode takes context, siteId, deviceId as parameters and
// returns an models.ApiResponse with models.GatewayCluster data and
// an error if there was an issue with the request or response.
// Delete HA Cluster
func (s *SitesDevicesWANCluster) GetSiteDeviceHaClusterNode(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID) (
    models.ApiResponse[models.GatewayCluster],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/devices/%v/ha")
    req.AppendTemplateParams(siteId, deviceId)
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
    
    var result models.GatewayCluster
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.GatewayCluster](decoder)
    return models.NewApiResponse(result, resp), err
}

// CreateSiteDeviceHaCluster takes context, siteId, deviceId, body as parameters and
// returns an models.ApiResponse with models.GatewayCluster data and
// an error if there was an issue with the request or response.
// ## Create HA Cluster
// Both nodes has to be in the same site. We expect the user to configure ha_sync / ha_data port in port_configs already
// ### SRX cabling
// see [Chassis Cluster User Guide for SRX Series Devices](https://www.juniper.net/documentation/us/en/software/junos/chassis-cluster-security-devices/topics/concept/chassis-cluster-srx-series-node-interface-understanding.html) Here’s the recommended cabling.
// #### SRX300
// From ZTP / default state, ge-0/0/0 and ge-0/0/7 (SFP) are default WAN ports and will get DHCP IP. However, ge-0/0/0 becomes OOB/fxp0 after cluster is enabled (i.e. using it for reach Mist is not recommended)
// 1.  form cluster in UI
// 2.  configure ge-0/0/7,ge-1/0/7 for WAN (reth0)
// 3.  configure ge-0/0/2,ge-1/0/2 for ha_data
// 4.  configure ge-0/0/3- for LAN or additional WAN e.g.
// ``` json
// {
// "port_config": {
// "ge-0/0/2,ge-1/0/2": {
// "usage": "ha_data"
// },
// "ge-0/0/7,ge-1/0/7": {
// "usage": "wan",
// "redundant": true,
// "reth_idx": 0,
// "ip_config": {"type": "dhcp"}
// },
// }
// }
// ```
// 1.  connect ge-0/0/1 back to back for ha_control
// 2.  connect ge-0/0/2 back to back for ha_data
// 3.  connect both ge-0/0/7 to uplink switch to WAN and to reach Mist
// 4.  power up both devices
// 5.  it takes about 30 minutes for the cluster to form
// #### SRX320
// From ZTP / default state, ge-0/0/0, ge-0/0/7 (SFP) and cl-1/0/0 (LTE) are default WAN ports and will get DHCP IP. However, ge-0/0/0 becomes OOB/fxp0 after cluster is enabled (i.e. using it for reach Mist is not recommended)
// ##### ZTP via ge-0/0/7
// Similar to SRX300
// ##### ZTP via cl-1/0/0 (LTE)
// 1.  form cluster in UI
// 2.  configure cl-1/0/0, cl-3/0/0 as WAN (reth0)
// 3.  configure ge-0/0/2,ge-3/0/2 for ha_data
// 4.  same as above
// #### SRX340 / SRX345 / SRX380
// SRX340/SRX345 has dedicated OOB/fxp0 ports
// 1.  form cluster in UI
// 2.  configure ge-0/0/0,ge-5/0/0 for WAN (reth0)
// 3.  configure ge-0/0/2,ge-5/0/2 for ha_data
// 4.  configure ge-0/0/3- for LAN or additional WAN
// 5.  connect ge-0/0/0 to uplink switch to WAN and to reach Mist
// 6.  connect ge-0/0/1 back-to-back for ha_control
// 7.  connect ge-0/0/2 back-to-back for ha_data (fabric); or for SRX380, xe-0/0/16 if 10G SFP+ is used
// 8.  connect ge-0/0/3- to LAN or additional WANs
// #### SRX550
// ge-0/0/0 becomes OOB/fxp0 after cluster is enabled, make suenable oob_ip_config as dhcp to maintain cloud connectivity
// 1.  connect ge-0/0/0 to reach Mist (after cluster is fully up, this port becomes OOB/fxp0)
// 2.  connect ge-0/0/1 back-to-back for ha_control
// 3.  connect ge-0/0/2 back-to-back for ha_data (fabric)
// 4.  connect ge-0/0/3 to WAN (after cluster is up, intended to be used for reth0)
// 5.  connect ge-0/0/4- to LAN or additional WANs
// #### SRX1500
// SRX1500 has, additionally, dedicated HA Control port
// 1.  form cluster in UI
// 2.  configure ge-0/0/0,ge-5/0/0 for WAN (reth0)
// 3.  configure ge-0/0/1,ge-5/0/1 for ha_data
// 4.  configure ge-0/0/2- for LAN or additional WAN
// 5.  connect dedicated ha_control back-to-back
// 6.  connect ge-0/0/0 to uplink switch to WAN and to reach Mist
// 7.  connect ge-0/0/1 back-to-back for ha_data
// 8.  connect ge-0/0/2- to LAN or additional WANs
// #### SRX4100
// SRX4100 has dedicated ha_control and ha_data (fabric) ports
// 1.  connect dedicated ha_control back-to-back
// 2.  connect dedicated ha_data back-to-back
// 3.  connect xe-0/0/0 to WAN to reach Mist
// 4.  connect xe-0/0/1- to LAN or additional WANs
// #### VSRX
// When standalone, VSRX has fxp0 as first Network Adapter, then ge-0/0/0-N When clustered, VSRX has fxp0, em0, then ge-0/0/0-N
// 1.  connect net0 (fxp0) to WAN to reach Mist
// 2.  connect net1 back-to-back for ha_control
// 3.  connect net2 (ge-0/0/0) back-to-back for ha_data (fab0/fab1)
// 4.  connect net3 (ge-0/0/1) to WAN, intended to be used for reth0
// 5.  connect net4 (ge-0/0/2) to LAN
// SRX340/SRX345 has dedicated OOB/fxp0 ports VSRX has fxp0 as first Network Adapter, then ge-0/0/0-N
// 1.  connect ge-0/0/0 to WAN to reach Mist
// 2.  connect ge-0/0/1 back-to-back for ha_control
// 3.  connect ge-0/0/2 back-to-back for ha_data (fabric); or for SRX380, xe-0/0/16 if 10G SFP+ is used
// 4.  connect ge-0/0/3- to LAN or additional WANs
// #### SRX550
// ge-0/0/0 becomes OOB/fxp0 after cluster is enabled, make suenable oob_ip_config as dhcp to maintain cloud connectivity
// 1.  connect ge-0/0/0 to reach Mist (after cluster is fully up, this port becomes OOB/fxp0)
// 2.  connect ge-0/0/1 back-to-back for ha_control
// 3.  connect ge-0/0/2 back-to-back for ha_data (fabric)
// 4.  connect ge-0/0/3 to WAN (after cluster is up, intended to be used for reth0)
// 5.  connect ge-0/0/4- to LAN or additional WANs
// #### SRX1500
// SRX1500 has, additionally, dedicated HA Control port
// 1. form cluster in UI
// 2. configure ge-0/0/0,ge-7/0/0 for WAN (reth0)
// 3. configure ge-0/0/1,ge-7/0/1 for ha_data
// 4. configure ge-0/0/2- for LAN or additional WAN
// 5. connect dedicated ha_control back-to-back
// 6. connect ge-0/0/0 to uplink switch to WAN and to reach Mist
// 7. connect ge-0/0/1 back-to-back for ha_data
// 8. connect ge-0/0/2- to LAN or additional WANs
// #### SRX1600
// SRX1600 has, additionally, two dedicated HA Control port
// 1. form cluster in UI
// 2. configure ge-0/0/0,ge-7/0/0 for WAN (reth0)
// 3. configure ge-0/0/1,ge-7/0/1 for ha_data
// 4. configure ge-0/0/2- for LAN or additional WAN
// 5. connect dedicated both ha_control back-to-back
// 6. connect ge-0/0/0 to uplink switch to WAN and to reach Mist
// 7. connect ge-0/0/1 back-to-back for ha_data
// 8. connect ge-0/0/2- to LAN or additional WANs
// #### SRX4100
// SRX4100 has dedicated ha_control and ha_data (fabric) ports
// 1.  connect dedicated ha_control back-to-back
// 2.  connect dedicated ha_data back-to-back
// 3.  connect xe-0/0/0 to WAN to reach Mist
// 4.  connect xe-0/0/1- to LAN or additional WANs
// ## Replace a Node in a HA Cluster
// Usually Device Replacement is done by Device Replacement API. For a HA cluster, you can also replace a node by another device in the same site.
func (s *SitesDevicesWANCluster) CreateSiteDeviceHaCluster(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.GatewayCluster) (
    models.ApiResponse[models.GatewayCluster],
    error) {
    req := s.prepareRequest(ctx, "POST", "/api/v1/sites/%v/devices/%v/ha")
    req.AppendTemplateParams(siteId, deviceId)
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
    
    var result models.GatewayCluster
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.GatewayCluster](decoder)
    return models.NewApiResponse(result, resp), err
}
