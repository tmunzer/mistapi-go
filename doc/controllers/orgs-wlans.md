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

[`models.Wlan`](../../doc/models/wlan.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Wlan{
    AcctImmediateUpdate:                models.ToPointer(false),
    AcctInterimInterval:                models.ToPointer(0),
    AllowIpv6Ndp:                       models.ToPointer(true),
    AllowMdns:                          models.ToPointer(false),
    AllowSsdp:                          models.ToPointer(false),
    ArpFilter:                          models.ToPointer(false),
    AuthServerSelection:                models.ToPointer(models.WlanAuthServerSelectionEnum("ordered")),
    AuthServersNasId:                   models.NewOptional(models.ToPointer("5c5b350e0101-nas")),
    AuthServersNasIp:                   models.NewOptional(models.ToPointer("15.3.1.5")),
    AuthServersRetries:                 models.ToPointer(5),
    AuthServersTimeout:                 models.ToPointer(5),
    BandSteer:                          models.ToPointer(false),
    BandSteerForceBand5:                models.ToPointer(false),
    BlockBlacklistClients:              models.ToPointer(false),
    ClientLimitDownEnabled:             models.ToPointer(false),
    ClientLimitUpEnabled:               models.ToPointer(false),
    Disable11ax:                        models.ToPointer(false),
    DisableHtVhtRates:                  models.ToPointer(false),
    DisableUapsd:                       models.ToPointer(false),
    DisableV1RoamNotify:                models.ToPointer(false),
    DisableV2RoamNotify:                models.ToPointer(false),
    DisableWmm:                         models.ToPointer(false),
    Dtim:                               models.ToPointer(2),
    EnableLocalKeycaching:              models.ToPointer(false),
    EnableWirelessBridging:             models.ToPointer(false),
    EnableWirelessBridgingDhcpTracking: models.ToPointer(false),
    Enabled:                            models.ToPointer(true),
    FastDot1xTimers:                    models.ToPointer(false),
    HideSsid:                           models.ToPointer(false),
    HostnameIe:                         models.ToPointer(false),
    Interface:                          models.ToPointer(models.WlanInterfaceEnum("all")),
    Isolation:                          models.ToPointer(false),
    L2Isolation:                        models.ToPointer(false),
    LegacyOverds:                       models.ToPointer(false),
    LimitBcast:                         models.ToPointer(false),
    LimitProbeResponse:                 models.ToPointer(false),
    MaxIdletime:                        models.ToPointer(1800),
    NoStaticDns:                        models.ToPointer(false),
    NoStaticIp:                         models.ToPointer(false),
    OrgId:                              models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
    PortalAllowedHostnames:             []string{
        "snapchat.com",
        "ibm.com",
    },
    PortalAllowedSubnets:               []string{
        "63.5.3.0/24",
    },
    PortalApiSecret:                    models.NewOptional(models.ToPointer("EIfPMOykI3lMlDdNPub2WcbqT6dNOtWwmYHAd6bY")),
    PortalDeniedHostnames:              []string{
        "msg.snapchat.com",
    },
    PortalImage:                        models.NewOptional(models.ToPointer("https://url/to/image.png")),
    RoamMode:                           models.ToPointer(models.WlanRoamModeEnum("none")),
    SiteId:                             models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
    SleExcluded:                        models.ToPointer(false),
    Ssid:                               "corporate",
    UseEapolV1:                         models.ToPointer(false),
    VlanEnabled:                        models.ToPointer(false),
    VlanIds:                            []models.WlanVlanIds{
        models.WlanVlanIdsContainer.FromNumber(3),
        models.WlanVlanIdsContainer.FromNumber(4),
        models.WlanVlanIdsContainer.FromNumber(5),
    },
    VlanPooling:                        models.ToPointer(false),
    WlanLimitDownEnabled:               models.ToPointer(false),
    WlanLimitUpEnabled:                 models.ToPointer(false),
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

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


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

``

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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


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

``

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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


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

[`models.Wlan`](../../doc/models/wlan.md)

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

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# List Org Wlans

Get List of Org Wlans

```go
ListOrgWlans(
    ctx context.Context,
    orgId uuid.UUID,
    page *int,
    limit *int) (
    models.ApiResponse[[]models.Wlan],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `page` | `*int` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |

## Response Type

[`[]models.Wlan`](../../doc/models/wlan.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

page := 1

limit := 100

apiResponse, err := orgsWlans.ListOrgWlans(ctx, orgId, &page, &limit)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


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

[`models.Wlan`](../../doc/models/wlan.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

wlanId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Wlan{
    AcctImmediateUpdate:                models.ToPointer(false),
    AcctInterimInterval:                models.ToPointer(0),
    AllowIpv6Ndp:                       models.ToPointer(true),
    AllowMdns:                          models.ToPointer(false),
    AllowSsdp:                          models.ToPointer(false),
    ArpFilter:                          models.ToPointer(false),
    AuthServerSelection:                models.ToPointer(models.WlanAuthServerSelectionEnum("ordered")),
    AuthServersNasId:                   models.NewOptional(models.ToPointer("5c5b350e0101-nas")),
    AuthServersNasIp:                   models.NewOptional(models.ToPointer("15.3.1.5")),
    AuthServersRetries:                 models.ToPointer(5),
    AuthServersTimeout:                 models.ToPointer(5),
    BandSteer:                          models.ToPointer(false),
    BandSteerForceBand5:                models.ToPointer(false),
    BlockBlacklistClients:              models.ToPointer(false),
    ClientLimitDownEnabled:             models.ToPointer(false),
    ClientLimitUpEnabled:               models.ToPointer(false),
    Disable11ax:                        models.ToPointer(false),
    DisableHtVhtRates:                  models.ToPointer(false),
    DisableUapsd:                       models.ToPointer(false),
    DisableV1RoamNotify:                models.ToPointer(false),
    DisableV2RoamNotify:                models.ToPointer(false),
    DisableWmm:                         models.ToPointer(false),
    Dtim:                               models.ToPointer(2),
    EnableLocalKeycaching:              models.ToPointer(false),
    EnableWirelessBridging:             models.ToPointer(false),
    EnableWirelessBridgingDhcpTracking: models.ToPointer(false),
    Enabled:                            models.ToPointer(true),
    FastDot1xTimers:                    models.ToPointer(false),
    HideSsid:                           models.ToPointer(false),
    HostnameIe:                         models.ToPointer(false),
    Interface:                          models.ToPointer(models.WlanInterfaceEnum("all")),
    Isolation:                          models.ToPointer(false),
    L2Isolation:                        models.ToPointer(false),
    LegacyOverds:                       models.ToPointer(false),
    LimitBcast:                         models.ToPointer(false),
    LimitProbeResponse:                 models.ToPointer(false),
    MaxIdletime:                        models.ToPointer(1800),
    NoStaticDns:                        models.ToPointer(false),
    NoStaticIp:                         models.ToPointer(false),
    OrgId:                              models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
    PortalAllowedHostnames:             []string{
        "snapchat.com",
        "ibm.com",
    },
    PortalAllowedSubnets:               []string{
        "63.5.3.0/24",
    },
    PortalApiSecret:                    models.NewOptional(models.ToPointer("EIfPMOykI3lMlDdNPub2WcbqT6dNOtWwmYHAd6bY")),
    PortalDeniedHostnames:              []string{
        "msg.snapchat.com",
    },
    PortalImage:                        models.NewOptional(models.ToPointer("https://url/to/image.png")),
    RoamMode:                           models.ToPointer(models.WlanRoamModeEnum("none")),
    SiteId:                             models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
    SleExcluded:                        models.ToPointer(false),
    Ssid:                               "corporate",
    UseEapolV1:                         models.ToPointer(false),
    VlanEnabled:                        models.ToPointer(false),
    VlanIds:                            []models.WlanVlanIds{
        models.WlanVlanIdsContainer.FromNumber(3),
        models.WlanVlanIdsContainer.FromNumber(4),
        models.WlanVlanIdsContainer.FromNumber(5),
    },
    VlanPooling:                        models.ToPointer(false),
    WlanLimitDownEnabled:               models.ToPointer(false),
    WlanLimitUpEnabled:                 models.ToPointer(false),
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

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


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
    models.ApiResponse[models.PortalTemplate],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `wlanId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.WlanPortalTemplate`](../../doc/models/wlan-portal-template.md) | Body, Optional | Request Body |

## Response Type

[`models.PortalTemplate`](../../doc/models/portal-template.md)

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
  "accessCodeAlternateEmail": "string",
  "alignment": "string",
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
  "created_time": 0,
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
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "message": "string",
  "modified_time": 0,
  "name": true,
  "nameError": "string",
  "nameLabel": "string",
  "optout": true,
  "optoutLabel": "string",
  "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
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
  "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
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
  "smsValidityDuration": 0,
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
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Upload Org Wlan Portal Image

Upload Org WLAN Portal Image

```go
UploadOrgWlanPortalImage(
    ctx context.Context,
    orgId uuid.UUID,
    wlanId uuid.UUID,
    file models.FileWrapper,
    json *string) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `wlanId` | `uuid.UUID` | Template, Required | - |
| `file` | `models.FileWrapper` | Form, Required | binary file |
| `json` | `*string` | Form, Optional | JSON string describing your upload |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

wlanId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

file := getFile("dummy_file", func(err error) { log.Fatalln(err) })



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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |

