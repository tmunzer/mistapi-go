# Orgs Wlans

```go
orgsWlans := client.OrgsWlans()
```

## Class Name

`OrgsWlans`

## Methods

* [Create Org Wlan](../../doc/controllers/orgs-wlans.md#create-org-wlan)
* [Delete Org Wlan](../../doc/controllers/orgs-wlans.md#delete-org-wlan)
* [Delete Org Wlan Portal Image](../../doc/controllers/orgs-wlans.md#delete-org-wlan-portal-image)
* [Get Org WLAN](../../doc/controllers/orgs-wlans.md#get-org-wlan)
* [List Org Wlans](../../doc/controllers/orgs-wlans.md#list-org-wlans)
* [Update Org Wlan](../../doc/controllers/orgs-wlans.md#update-org-wlan)
* [Update Org Wlan Portal Template](../../doc/controllers/orgs-wlans.md#update-org-wlan-portal-template)
* [Upload Org Wlan Portal Image](../../doc/controllers/orgs-wlans.md#upload-org-wlan-portal-image)


# Create Org Wlan

Create Org Wlan

```go
CreateOrgWlan(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.Wlan) (
    models.ApiResponse[models.Wlan],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Wlan`](../../doc/models/wlan.md) | Body, Optional | Request Body |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Wlan](../../doc/models/wlan.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Wlan{
    AcctImmediateUpdate:                  models.ToPointer(false),
    AcctInterimInterval:                  models.ToPointer(0),
    AllowIpv6Ndp:                         models.ToPointer(true),
    AllowMdns:                            models.ToPointer(false),
    AllowSsdp:                            models.ToPointer(false),
    ArpFilter:                            models.ToPointer(false),
    AuthServerSelection:                  models.ToPointer(models.WlanAuthServerSelectionEnum_ORDERED),
    AuthServersNasId:                     models.NewOptional(models.ToPointer("5c5b350e0101-nas")),
    AuthServersNasIp:                     models.NewOptional(models.ToPointer("15.3.1.5")),
    AuthServersRetries:                   models.ToPointer(5),
    AuthServersTimeout:                   models.ToPointer(5),
    BandSteer:                            models.ToPointer(false),
    BandSteerForceBand5:                  models.ToPointer(false),
    BlockBlacklistClients:                models.ToPointer(false),
    ClientLimitDownEnabled:               models.ToPointer(false),
    ClientLimitUpEnabled:                 models.ToPointer(false),
    Disable11ax:                          models.ToPointer(false),
    Disable11be:                          models.ToPointer(false),
    DisableHtVhtRates:                    models.ToPointer(false),
    DisableUapsd:                         models.ToPointer(false),
    DisableV1RoamNotify:                  models.ToPointer(false),
    DisableV2RoamNotify:                  models.ToPointer(false),
    DisableWhenGatewayUnreachable:        models.ToPointer(false),
    DisableWhenMxtunnelDown:              models.ToPointer(false),
    DisableWmm:                           models.ToPointer(false),
    Dtim:                                 models.ToPointer(2),
    EnableLocalKeycaching:                models.ToPointer(false),
    EnableWirelessBridging:               models.ToPointer(false),
    EnableWirelessBridgingDhcpTracking:   models.ToPointer(false),
    Enabled:                              models.ToPointer(true),
    FastDot1xTimers:                      models.ToPointer(false),
    HideSsid:                             models.ToPointer(false),
    HostnameIe:                           models.ToPointer(false),
    Interface:                            models.ToPointer(models.WlanInterfaceEnum_ALL),
    Isolation:                            models.ToPointer(false),
    L2Isolation:                          models.ToPointer(false),
    LegacyOverds:                         models.ToPointer(false),
    LimitBcast:                           models.ToPointer(false),
    LimitProbeResponse:                   models.ToPointer(false),
    MaxIdletime:                          models.ToPointer(1800),
    MaxNumClients:                        models.ToPointer(0),
    NoStaticDns:                          models.ToPointer(false),
    NoStaticIp:                           models.ToPointer(false),
    PortalAllowedHostnames:               []string{
        "snapchat.com",
        "ibm.com",
    },
    PortalAllowedSubnets:                 []string{
        "63.5.3.0/24",
    },
    PortalApiSecret:                      models.NewOptional(models.ToPointer("EIfPMOykI3lMlDdNPub2WcbqT6dNOtWwmYHAd6bY")),
    PortalDeniedHostnames:                []string{
        "msg.snapchat.com",
    },
    ReconnectClientsWhenRoamingMxcluster: models.ToPointer(false),
    RoamMode:                             models.ToPointer(models.WlanRoamModeEnum_NONE),
    SleExcluded:                          models.ToPointer(false),
    Ssid:                                 "corporate",
    UseEapolV1:                           models.ToPointer(false),
    VlanEnabled:                          models.ToPointer(false),
    VlanPooling:                          models.ToPointer(false),
    WlanLimitDownEnabled:                 models.ToPointer(false),
    WlanLimitUpEnabled:                   models.ToPointer(false),
}

apiResponse, err := orgsWlans.CreateOrgWlan(ctx, orgId, &body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
{
  "allow_ipv6_ndp": true,
  "allow_mdns": false,
  "allow_ssdp": false,
  "arp_filter": false,
  "band_steer": false,
  "band_steer_force_band5": false,
  "bands": [
    "24",
    "5"
  ],
  "block_blacklist_clients": false,
  "bonjour": {
    "additional_vlan_ids": "10, 20",
    "enabled": false,
    "services": {
      "airplay": {
        "radius_groups": [
          "teachers"
        ],
        "scope": "same_ap"
      }
    }
  },
  "client_limit_down": 1000,
  "client_limit_down_enabled": false,
  "client_limit_up": 512,
  "client_limit_up_enabled": false,
  "disable_11ax": false,
  "disable_ht_vht_rates": false,
  "disable_uapsd": false,
  "disable_v1_roam_notify": false,
  "disable_v2_roam_notify": false,
  "disable_wmm": false,
  "dynamic_vlan": {
    "default_vlan_id": 999,
    "enabled": false,
    "local_vlan_ids": [
      1
    ],
    "type": "airespace-interface-name",
    "vlans": {
      "131": "default",
      "322": "fast,video"
    }
  },
  "enable_local_keycaching": false,
  "enable_wireless_bridging": false,
  "enabled": true,
  "fast_dot1x_timers": false,
  "hide_ssid": false,
  "hostname_ie": false,
  "ssid": "demo"
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Delete Org Wlan

Delete Org WLAN

```go
DeleteOrgWlan(
    ctx context.Context,
    orgId uuid.UUID,
    wlanId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `wlanId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

wlanId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsWlans.DeleteOrgWlan(ctx, orgId, wlanId)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Delete Org Wlan Portal Image

Delete Org WLAN Portal Image

```go
DeleteOrgWlanPortalImage(
    ctx context.Context,
    orgId uuid.UUID,
    wlanId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `wlanId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

wlanId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsWlans.DeleteOrgWlanPortalImage(ctx, orgId, wlanId)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Get Org WLAN

Get Org Wlan Detail

```go
GetOrgWLAN(
    ctx context.Context,
    orgId uuid.UUID,
    wlanId uuid.UUID) (
    models.ApiResponse[models.Wlan],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `wlanId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Wlan](../../doc/models/wlan.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

wlanId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsWlans.GetOrgWLAN(ctx, orgId, wlanId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
{
  "allow_ipv6_ndp": true,
  "allow_mdns": false,
  "allow_ssdp": false,
  "arp_filter": false,
  "band_steer": false,
  "band_steer_force_band5": false,
  "bands": [
    "24",
    "5"
  ],
  "block_blacklist_clients": false,
  "bonjour": {
    "additional_vlan_ids": "10, 20",
    "enabled": false,
    "services": {
      "airplay": {
        "radius_groups": [
          "teachers"
        ],
        "scope": "same_ap"
      }
    }
  },
  "client_limit_down": 1000,
  "client_limit_down_enabled": false,
  "client_limit_up": 512,
  "client_limit_up_enabled": false,
  "disable_11ax": false,
  "disable_ht_vht_rates": false,
  "disable_uapsd": false,
  "disable_v1_roam_notify": false,
  "disable_v2_roam_notify": false,
  "disable_wmm": false,
  "dynamic_vlan": {
    "default_vlan_id": 999,
    "enabled": false,
    "local_vlan_ids": [
      1
    ],
    "type": "airespace-interface-name",
    "vlans": {
      "131": "default",
      "322": "fast,video"
    }
  },
  "enable_local_keycaching": false,
  "enable_wireless_bridging": false,
  "enabled": true,
  "fast_dot1x_timers": false,
  "hide_ssid": false,
  "hostname_ie": false,
  "ssid": "demo"
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# List Org Wlans

Get List of Org Wlans

```go
ListOrgWlans(
    ctx context.Context,
    orgId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.Wlan],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br><br>**Constraints**: `>= 1` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.Wlan](../../doc/models/wlan.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

page := 1

apiResponse, err := orgsWlans.ListOrgWlans(ctx, orgId, &limit, &page)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Update Org Wlan

Update Org Wlan

```go
UpdateOrgWlan(
    ctx context.Context,
    orgId uuid.UUID,
    wlanId uuid.UUID,
    body *models.Wlan) (
    models.ApiResponse[models.Wlan],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `wlanId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Wlan`](../../doc/models/wlan.md) | Body, Optional | Request Body |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Wlan](../../doc/models/wlan.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

wlanId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Wlan{
    AllowIpv6Ndp:                         models.ToPointer(true),
    AllowMdns:                            models.ToPointer(false),
    AllowSsdp:                            models.ToPointer(false),
    ArpFilter:                            models.ToPointer(false),
    BandSteer:                            models.ToPointer(false),
    BandSteerForceBand5:                  models.ToPointer(false),
    Bands:                                []models.Dot11BandEnum{
        models.Dot11BandEnum_ENUM24,
        models.Dot11BandEnum_ENUM5,
    },
    BlockBlacklistClients:                models.ToPointer(false),
    Bonjour:                              models.ToPointer(models.WlanBonjour{
        AdditionalVlanIds:    models.ToPointer(models.AdditionalVlanIdsContainer.FromString("10,20")),
        Enabled:              models.ToPointer(false),
        Services:             map[string]models.WlanBonjourServiceProperties{
            "airplay": models.WlanBonjourServiceProperties{
                RadiusGroups:         []string{
                    "teachers",
                },
                Scope:                models.ToPointer(models.WlanBonjourServicePropertiesScopeEnum_SAMEAP),
            },
        },
    }),
    ClientLimitDown:                      models.ToPointer(models.WlanLimitContainer.FromNumber(1000)),
    ClientLimitDownEnabled:               models.ToPointer(false),
    ClientLimitUp:                        models.ToPointer(models.WlanLimitContainer.FromNumber(512)),
    ClientLimitUpEnabled:                 models.ToPointer(false),
    Disable11ax:                          models.ToPointer(false),
    DisableHtVhtRates:                    models.ToPointer(false),
    DisableUapsd:                         models.ToPointer(false),
    DisableV1RoamNotify:                  models.ToPointer(false),
    DisableV2RoamNotify:                  models.ToPointer(false),
    DisableWmm:                           models.ToPointer(false),
    DynamicVlan:                          models.NewOptional(models.ToPointer(models.WlanDynamicVlan{
        DefaultVlanId:        models.ToPointer(models.WlanDynamicVlanDefaultVlanIdDeprecatedContainer.FromNumber(999)),
        Enabled:              models.ToPointer(false),
        LocalVlanIds:         []models.VlanIdWithVariable{
            models.VlanIdWithVariableContainer.FromNumber(1),
        },
        Type:                 models.ToPointer(models.WlanDynamicVlanTypeEnum_AIRESPACEINTERFACENAME),
        Vlans:                map[string]string{
            "131": "default",
            "322": "fast,video",
        },
    })),
    EnableLocalKeycaching:                models.ToPointer(false),
    EnableWirelessBridging:               models.ToPointer(false),
    Enabled:                              models.ToPointer(true),
    FastDot1xTimers:                      models.ToPointer(false),
    HideSsid:                             models.ToPointer(false),
    HostnameIe:                           models.ToPointer(false),
    Ssid:                                 "demo",
}

apiResponse, err := orgsWlans.UpdateOrgWlan(ctx, orgId, wlanId, &body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
{
  "allow_ipv6_ndp": true,
  "allow_mdns": false,
  "allow_ssdp": false,
  "arp_filter": false,
  "band_steer": false,
  "band_steer_force_band5": false,
  "bands": [
    "24",
    "5"
  ],
  "block_blacklist_clients": false,
  "bonjour": {
    "additional_vlan_ids": "10, 20",
    "enabled": false,
    "services": {
      "airplay": {
        "radius_groups": [
          "teachers"
        ],
        "scope": "same_ap"
      }
    }
  },
  "client_limit_down": 1000,
  "client_limit_down_enabled": false,
  "client_limit_up": 512,
  "client_limit_up_enabled": false,
  "disable_11ax": false,
  "disable_ht_vht_rates": false,
  "disable_uapsd": false,
  "disable_v1_roam_notify": false,
  "disable_v2_roam_notify": false,
  "disable_wmm": false,
  "dynamic_vlan": {
    "default_vlan_id": 999,
    "enabled": false,
    "local_vlan_ids": [
      1
    ],
    "type": "airespace-interface-name",
    "vlans": {
      "131": "default",
      "322": "fast,video"
    }
  },
  "enable_local_keycaching": false,
  "enable_wireless_bridging": false,
  "enabled": true,
  "fast_dot1x_timers": false,
  "hide_ssid": false,
  "hostname_ie": false,
  "ssid": "demo"
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Update Org Wlan Portal Template

Update a Portal Template

#### Sponsor Email Template

Sponsor Email Template supports following template variables:

| **Name** | **Description** |
| --- | --- |
| approve_url | Renders URL to approve the request; optionally &minutes=N query param can be appended to change the Authorization period of the guest, where N is a valid integer denoting number of minutes a guest remains authorized |
| deny_url | Renders URL to reject the request |
| guest_email | Renders Email ID of the guest |
| guest_name | Renders Name of the guest |
| field1 | Renders value of the Custom Field 1 |
| field2 | Renders value of the Custom Field 2 |
| company | Renders value of the Company field |
| sponsor_link_validity_duration | Renders validity time of the request (i.e. Approve/Deny URL) |
| auth_expire_minutes | Renders Wlan-level configured Guest Authorization Expiration time period (in minutes), If not configured then default (1 day in minutes) |

```go
UpdateOrgWlanPortalTemplate(
    ctx context.Context,
    orgId uuid.UUID,
    wlanId uuid.UUID,
    body *models.WlanPortalTemplate) (
    models.ApiResponse[models.WlanPortalTemplate],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `wlanId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.WlanPortalTemplate`](../../doc/models/wlan-portal-template.md) | Body, Optional | Request Body |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.WlanPortalTemplate](../../doc/models/wlan-portal-template.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

wlanId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsWlans.UpdateOrgWlanPortalTemplate(ctx, orgId, wlanId, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
{
  "portal_template": {
    "accessCodeAlternateEmail": "string",
    "alignment": "left",
    "authButtonAmazon": "string",
    "authButtonAzure": "string",
    "authButtonEmail": "string",
    "authButtonFacebook": "string",
    "authButtonGoogle": "string",
    "authButtonMicrosoft": "string",
    "authButtonPassphrase": "string",
    "authButtonSms": "string",
    "authButtonSponsor": "string",
    "authLabel": "string",
    "backLink": "string",
    "color": "string",
    "colorDark": "string",
    "colorLight": "string",
    "company": true,
    "companyError": "string",
    "companyLabel": "string",
    "email": true,
    "emailAccessDomainError": "string",
    "emailCancel": "string",
    "emailCodeError": "string",
    "emailError": "string",
    "emailFieldLabel": "string",
    "emailLabel": "string",
    "emailMessage": "string",
    "emailSubmit": "string",
    "emailTitle": "string",
    "field1": true,
    "field1Error": "string",
    "field1Label": "string",
    "field1Required": true,
    "field2": true,
    "field2Error": "string",
    "field2Label": "string",
    "field2Required": true,
    "field3": true,
    "field3Error": "string",
    "field3Label": "string",
    "field3Required": true,
    "field4": true,
    "field4Error": "string",
    "field4Label": "string",
    "field4Required": true,
    "message": "string",
    "name": true,
    "nameError": "string",
    "nameLabel": "string",
    "optout": true,
    "optoutLabel": "string",
    "pageTitle": "string",
    "passphraseCancel": "string",
    "passphraseError": "string",
    "passphraseLabel": "string",
    "passphraseMessage": "string",
    "passphraseSubmit": "string",
    "passphraseTitle": "string",
    "poweredBy": true,
    "requiredFieldLabel": "string",
    "signInLabel": "string",
    "smsCarrierDefault": "string",
    "smsCarrierError": "string",
    "smsCarrierFieldLabel": "string",
    "smsCodeCancel": "string",
    "smsCodeError": "string",
    "smsCodeFieldLabel": "string",
    "smsCodeMessage": "string",
    "smsCodeSubmit": "string",
    "smsCodeTitle": "string",
    "smsCountryFieldLabel": "string",
    "smsCountryFormat": "string",
    "smsHaveAccessCode": "string",
    "smsMessageFormat": "string",
    "smsNumberCancel": "string",
    "smsNumberError": "string",
    "smsNumberFieldLabel": "string",
    "smsNumberFormat": "string",
    "smsNumberMessage": "string",
    "smsNumberSubmit": "string",
    "smsNumberTitle": "string",
    "smsUsernameFormat": "string",
    "smsValidityDuration": 5,
    "sponsorBackLink": "string",
    "sponsorCancel": "string",
    "sponsorEmail": "string",
    "sponsorEmailError": "string",
    "sponsorEmailTemplate": "string",
    "sponsorInfoApproved": "string",
    "sponsorInfoDenied": "string",
    "sponsorInfoPending": "string",
    "sponsorName": "string",
    "sponsorNameError": "string",
    "sponsorNotePending": "string",
    "sponsorStatusApproved": "string",
    "sponsorStatusDenied": "string",
    "sponsorStatusPending": "string",
    "sponsorSubmit": "string",
    "tos": true,
    "tosAcceptLabel": "string",
    "tosError": "string",
    "tosLink": "string",
    "tosText": "string"
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Upload Org Wlan Portal Image

Upload Org WLAN Portal Image

```go
UploadOrgWlanPortalImage(
    ctx context.Context,
    orgId uuid.UUID,
    wlanId uuid.UUID,
    file string,
    json *string) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `wlanId` | `uuid.UUID` | Template, Required | - |
| `file` | `string` | Form, Required | Binary file |
| `json` | `*string` | Form, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

wlanId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

file := "file0"

resp, err := orgsWlans.UploadOrgWlanPortalImage(ctx, orgId, wlanId, file, nil)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |

