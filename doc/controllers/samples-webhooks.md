# Samples Webhooks

```go
samplesWebhooks := client.SamplesWebhooks()
```

## Class Name

`SamplesWebhooks`

## Methods

* [Alarms](../../doc/controllers/samples-webhooks.md#alarms)
* [Audits](../../doc/controllers/samples-webhooks.md#audits)
* [Client Info](../../doc/controllers/samples-webhooks.md#client-info)
* [Client Join](../../doc/controllers/samples-webhooks.md#client-join)
* [Client Latency](../../doc/controllers/samples-webhooks.md#client-latency)
* [Client Sessions](../../doc/controllers/samples-webhooks.md#client-sessions)
* [Device Events](../../doc/controllers/samples-webhooks.md#device-events)
* [Device up Down](../../doc/controllers/samples-webhooks.md#device-up-down)
* [Discovered-Raw-Rssi](../../doc/controllers/samples-webhooks.md#discovered-raw-rssi)
* [Guest Authorization](../../doc/controllers/samples-webhooks.md#guest-authorization)
* [Location](../../doc/controllers/samples-webhooks.md#location)
* [Location Asset](../../doc/controllers/samples-webhooks.md#location-asset)
* [Location Centrak](../../doc/controllers/samples-webhooks.md#location-centrak)
* [Location Client](../../doc/controllers/samples-webhooks.md#location-client)
* [Location Sdk](../../doc/controllers/samples-webhooks.md#location-sdk)
* [Location Unclient](../../doc/controllers/samples-webhooks.md#location-unclient)
* [Nac Accounting](../../doc/controllers/samples-webhooks.md#nac-accounting)
* [Nac Events](../../doc/controllers/samples-webhooks.md#nac-events)
* [Occupancy Alerts](../../doc/controllers/samples-webhooks.md#occupancy-alerts)
* [Ping](../../doc/controllers/samples-webhooks.md#ping)
* [Sdkclient Scan Data](../../doc/controllers/samples-webhooks.md#sdkclient-scan-data)
* [Site Sle](../../doc/controllers/samples-webhooks.md#site-sle)
* [Zone](../../doc/controllers/samples-webhooks.md#zone)


# Alarms

Webhook sample for `alarm` topic

**Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages

```go
Alarms(
    ctx context.Context,
    body *models.WebhookAlarms) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`*models.WebhookAlarms`](../../doc/models/webhook-alarms.md) | Body, Optional | **N.B.**: Fields like `aps`, `bssids`, `ssids` are event specific. They are relevant to this event type ( rogue-ap-detected). For a different event type, different fields may be sent. These don’t contain all affected entities and are representative samples of entities (capped at 10). For marvis action related events, we expose `details` to include more event specific details.<br><br>Events specific fields for other alarm event type can be found with API [List Alarm Definitions]($e/Events%20Definitions/listAlarmDefinitions), under "fields" array of /alarm_defs response object. |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

body := models.WebhookAlarms{
    Events:               []models.WebhookAlarmEvent{
        models.WebhookAlarmEvent{
            Id:                   uuid.MustParse(""),
            LastSeen:             nil,
            OrgId:                uuid.MustParse(""),
            SiteId:               uuid.MustParse(""),
            Timestamp:            0.0,
            Type:                 "",
        },
    },
    Topic:                "alarms",
}

resp, err := samplesWebhooks.Alarms(ctx, &body)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```


# Audits

Webhook sample for `audit` topic

**Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages

```go
Audits(
    ctx context.Context,
    body *models.WebhookAudits) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`*models.WebhookAudits`](../../doc/models/webhook-audits.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

body := models.WebhookAudits{
    Events:               []models.WebhookAuditEvent{
        models.WebhookAuditEvent{
            AdminName:            "admin_name8",
            DeviceId:             uuid.MustParse(""),
            Id:                   uuid.MustParse(""),
            Message:              "message0",
            OrgId:                uuid.MustParse(""),
            SiteId:               uuid.MustParse(""),
            SrcIp:                "src_ip6",
            Timestamp:            0.0,
        },
    },
    Topic:                "audits",
}

resp, err := samplesWebhooks.Audits(ctx, &body)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```


# Client Info

Webhook sample for `client-info` topic

**Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages

```go
ClientInfo(
    ctx context.Context,
    body *models.WebhookClientInfo) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`*models.WebhookClientInfo`](../../doc/models/webhook-client-info.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

body := models.WebhookClientInfo{
    Events:               []models.WebhookClientInfoEvent{
        models.WebhookClientInfoEvent{
            Hostname:             models.ToPointer("service.company.net"),
            Ip:                   models.ToPointer("21.0.128.151"),
            Mac:                  models.ToPointer("6ebaa47a3fd4"),
        },
    },
    Topic:                models.ToPointer(models.WebhookClientInfoTopicEnum_CLIENTINFO),
}

resp, err := samplesWebhooks.ClientInfo(ctx, &body)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```


# Client Join

Webhook sample for `client_join` topic

**Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages

```go
ClientJoin(
    ctx context.Context,
    body *models.WebhookClientJoin) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`*models.WebhookClientJoin`](../../doc/models/webhook-client-join.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

body := models.WebhookClientJoin{
    Events:               []models.WebhookClientJoinEvent{
        models.WebhookClientJoinEvent{
            Ap:                   "string",
            ApName:               "string",
            Band:                 "string",
            Bssid:                "string",
            Connect:              0,
            ConnectFloat:         float64(0),
            Mac:                  "string",
            OrgId:                uuid.MustParse(""),
            Rssi:                 float64(0),
            SiteId:               uuid.MustParse(""),
            SiteName:             "string",
            Ssid:                 "string",
            Timestamp:            0.0,
            Version:              float64(0),
            WlanId:               uuid.MustParse("5028e92b-fc59-4056-91d1-ea4b4ca1617a"),
        },
    },
    Topic:                "client-join",
}

resp, err := samplesWebhooks.ClientJoin(ctx, &body)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```


# Client Latency

Webhook sample for `client-latency` topic

**Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages

```go
ClientLatency(
    ctx context.Context,
    body *models.WebhookClientLatency) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`*models.WebhookClientLatency`](../../doc/models/webhook-client-latency.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

body := models.WebhookClientLatency{
    Events:               []models.WebhookClientLatencyEvent{
        models.WebhookClientLatencyEvent{
            AvgAuth:              models.ToPointer(float64(0.17170219)),
            AvgDhcp:              models.ToPointer(float64(0.017828934)),
            AvgDns:               models.ToPointer(float64(0.024532124)),
            MaxAuth:              models.ToPointer(float64(0.18170219)),
            MaxDhcp:              models.ToPointer(float64(0.027828934)),
            MaxDns:               models.ToPointer(float64(0.029532124)),
            MinAuth:              models.ToPointer(float64(0.16050219)),
            MinDhcp:              models.ToPointer(float64(0.015828934)),
            MinDns:               models.ToPointer(float64(0.022532124)),
        },
    },
    Topic:                models.ToPointer("client-latency"),
}

resp, err := samplesWebhooks.ClientLatency(ctx, &body)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```


# Client Sessions

Webhook sample for `client_sessions` topic

**Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages

```go
ClientSessions(
    ctx context.Context,
    body *models.WebhookClientSessions) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`*models.WebhookClientSessions`](../../doc/models/webhook-client-sessions.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

body := models.WebhookClientSessions{
    Events:               []models.WebhookClientSessionsEvent{
        models.WebhookClientSessionsEvent{
            Ap:                   "string",
            ApName:               "string",
            Band:                 "string",
            Bssid:                "string",
            ClientFamily:         "string",
            ClientManufacture:    "string",
            ClientModel:          "string",
            ClientOs:             "string",
            Connect:              0,
            ConnectFloat:         float64(0),
            Disconnect:           0,
            DisconnectFloat:      float64(0),
            Duration:             0,
            Mac:                  "string",
            NextAp:               "string",
            OrgId:                uuid.MustParse(""),
            Rssi:                 float64(0),
            SiteId:               uuid.MustParse(""),
            SiteName:             "string",
            Ssid:                 "string",
            TerminationReason:    0,
            Timestamp:            0.0,
            Version:              float64(0),
            WlanId:               uuid.MustParse("5028e92b-fc59-4056-91d1-ea4b4ca1617a"),
        },
    },
    Topic:                "client-sessions",
}

resp, err := samplesWebhooks.ClientSessions(ctx, &body)
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


# Device Events

Webhook sample for `device_events` topic

**Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages

```go
DeviceEvents(
    ctx context.Context,
    body *models.WebhookDeviceEvents) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`*models.WebhookDeviceEvents`](../../doc/models/webhook-device-events.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

body := models.WebhookDeviceEvents{
    Events:               []models.WebhookDeviceEventsEvent{
        models.WebhookDeviceEventsEvent{
            Ap:                   models.ToPointer("string"),
            ApName:               models.ToPointer("string"),
            AuditId:              models.ToPointer(uuid.MustParse("78c04fa6-cfb4-46a0-9aa5-3681ba4f3897")),
            DeviceName:           "string",
            DeviceType:           models.DeviceTypeEnum_AP,
            EvType:               models.WebhookDeviceEventsEventEvTypeEnum_NOTICE,
            Mac:                  "string",
            OrgId:                uuid.MustParse(""),
            Reason:               models.ToPointer("string"),
            SiteName:             models.ToPointer("string"),
            Text:                 models.ToPointer("string"),
            Timestamp:            0.0,
            Type:                 "string",
        },
    },
    Topic:                "device_events",
}

resp, err := samplesWebhooks.DeviceEvents(ctx, &body)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```


# Device up Down

Webhook sample for `device_updowns` topic

**Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages

```go
DeviceUpDown(
    ctx context.Context,
    body *models.WebhookDeviceUpdowns) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`*models.WebhookDeviceUpdowns`](../../doc/models/webhook-device-updowns.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

body := models.WebhookDeviceUpdowns{
    Events:               []models.WebhookDeviceUpdownsEvent{
        models.WebhookDeviceUpdownsEvent{
            Ap:                   "",
            ApName:               "",
            OrgId:                uuid.MustParse(""),
            SiteId:               uuid.MustParse(""),
            SiteName:             "",
            Timestamp:            0.0,
            Type:                 "",
        },
    },
    Topic:                "device-updowns",
}

resp, err := samplesWebhooks.DeviceUpDown(ctx, &body)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```


# Discovered-Raw-Rssi

Webhook sample for `discovered-raw-rssi` topic

**Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages

```go
DiscoveredRawRssi(
    ctx context.Context,
    body *models.WebhookDiscoveredRawRssi) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`*models.WebhookDiscoveredRawRssi`](../../doc/models/webhook-discovered-raw-rssi.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

body := models.WebhookDiscoveredRawRssi{
    Events:               []models.WebhookDiscoveredRawRssiEvent{
        models.WebhookDiscoveredRawRssiEvent{
            ApLoc:                []float64{
                float64(0),
            },
            Beam:                 0,
            DeviceId:             uuid.MustParse("3bafab7b-4400-4bcf-8e6e-09f954699940"),
            IbeaconMajor:         models.ToPointer(0),
            IbeaconMinor:         models.ToPointer(0),
            IbeaconUuid:          models.ToPointer(uuid.MustParse("1f89bc00-d0af-481b-82fe-a6629259a39f")),
            IsAsset:              models.ToPointer(true),
            Mac:                  "string",
            MapId:                uuid.MustParse("09d2b626-2e4e-45ef-a3c4-e6aeb6c83db1"),
            MfgCompanyId:         models.ToPointer("string"),
            MfgData:              models.ToPointer("string"),
            OrgId:                uuid.MustParse(""),
            Rssi:                 float64(0),
            ServicePackets:       []models.ServicePacket{
                models.ServicePacket{
                    ServiceData:          models.ToPointer("string"),
                    ServiceUuid:          models.ToPointer("7138cc00-c611-4dec-a05e-5c4b1cae13c0"),
                },
            },
            SiteId:               uuid.MustParse(""),
        },
    },
    Topic:                "string",
}

resp, err := samplesWebhooks.DiscoveredRawRssi(ctx, &body)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```


# Guest Authorization

Webhook sample for `guest-authorizations` topic

**Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages

```go
GuestAuthorization(
    ctx context.Context,
    body *models.WebhookGuestAuthorizations) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`*models.WebhookGuestAuthorizations`](../../doc/models/webhook-guest-authorizations.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

body := models.WebhookGuestAuthorizations{
    Events:               []models.WebhookGuestAuthorizationsEvent{
        models.WebhookGuestAuthorizationsEvent{
            Ap:                     models.ToPointer("5c5b350e55c8"),
            AuthMethod:             models.ToPointer("passphrase"),
            AuthorizedExpiringTime: models.ToPointer(1677076639),
            AuthorizedTime:         models.ToPointer(1677076519),
            Carrier:                models.ToPointer("docomo"),
            Client:                 models.ToPointer("ac2316eca70a"),
            Company:                models.ToPointer("MIST"),
            Email:                  models.ToPointer("abcd@abcd.com"),
            Field1:                 models.ToPointer("field1 value"),
            Field2:                 models.ToPointer("field2 value"),
            Field3:                 models.ToPointer("field3 value"),
            Field4:                 models.ToPointer("field4 value"),
            Mobile:                 models.ToPointer("+0123456789"),
            Name:                   models.ToPointer("Dr Strange"),
            SmsGateway:             models.ToPointer("Telstra"),
            SponsorEmail:           models.ToPointer("sponsor@gmail.com"),
            Ssid:                   models.ToPointer("Portal Auth"),
            WlanId:                 models.ToPointer("7681be9a-044a-4622-90cf-3accde5ad853"),
        },
    },
    Topic:                models.ToPointer("guest-authorizations"),
}

resp, err := samplesWebhooks.GuestAuthorization(ctx, &body)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```


# Location

Webhook sample for `location` topic

**Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages

```go
Location(
    ctx context.Context,
    body *models.WebhookLocation) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`*models.WebhookLocation`](../../doc/models/webhook-location.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

body := models.WebhookLocation{
    Events:               []models.WebhookLocationEvent{
        models.WebhookLocationEvent{
            BatteryVoltage:         models.ToPointer(0),
            EddystoneUidInstance:   models.ToPointer("string"),
            EddystoneUidNamespace:  models.ToPointer("string"),
            EddystoneUrlUrl:        models.ToPointer("string"),
            IbeaconMajor:           models.ToPointer(0),
            IbeaconMinor:           models.ToPointer(0),
            IbeaconUuid:            models.ToPointer(uuid.MustParse("1f89bc00-d0af-481b-82fe-a6629259a39f")),
            Id:                     uuid.MustParse(""),
            Mac:                    models.ToPointer("string"),
            MapId:                  uuid.MustParse("09d2b626-2e4e-45ef-a3c4-e6aeb6c83db1"),
            MfgCompanyId:           models.ToPointer(0),
            MfgData:                models.ToPointer("string"),
            Name:                   models.ToPointer("string"),
            SiteId:                 uuid.MustParse(""),
            Timestamp:              0.0,
            Type:                   "string",
            X:                      float64(0),
            Y:                      float64(0),
        },
    },
    Topic:                "location",
}

resp, err := samplesWebhooks.Location(ctx, &body)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```


# Location Asset

Webhook sample for `location_asset` topic

**Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages

```go
LocationAsset(
    ctx context.Context,
    body *models.WebhookLocationAsset) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`*models.WebhookLocationAsset`](../../doc/models/webhook-location-asset.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

body := models.WebhookLocationAsset{
    Events:               []models.WebhookLocationAssetEvent{
        models.WebhookLocationAssetEvent{
            BatteryVoltage:        models.ToPointer(3370),
            EddystoneUidInstance:  models.ToPointer("5c5b35000001"),
            EddystoneUidNamespace: models.ToPointer("2818e3868dec25629ede"),
            EddystoneUrlUrl:       models.ToPointer("https://www.abc.com"),
            IbeaconMajor:          models.ToPointer(13),
            IbeaconMinor:          models.ToPointer(138),
            IbeaconUuid:           models.ToPointer(uuid.MustParse("f3f17139-704a-f03a-2786-0400279e37c3")),
            Mac:                   models.ToPointer("7fc2936fd243"),
            MapId:                 models.ToPointer(uuid.MustParse("845a23bf-bed9-e43c-4c86-6fa474be7ae5")),
            MfgCompanyId:          models.ToPointer(935),
            MfgData:               models.ToPointer("648520a1020000"),
            Type:                  models.ToPointer("asset"),
            X:                     models.ToPointer(float64(13.5)),
            Y:                     models.ToPointer(float64(3.2)),
        },
    },
    Topic:                "location_asset",
}

resp, err := samplesWebhooks.LocationAsset(ctx, &body)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```


# Location Centrak

Webhook sample for `location_centrak` topic

**Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages

```go
LocationCentrak(
    ctx context.Context,
    body *models.WebhookLocationCentrak) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`*models.WebhookLocationCentrak`](../../doc/models/webhook-location-centrak.md) | Body, Optional | **N.B.**: Fields like `aps`, `bssids`, `ssids` are event specific. They are relevant to this event type ( rogue-ap-detected). For a different event type, different fields may be sent. These don’t contain all affected entities and are representative samples of entities (capped at 10). For marvis action related events, we expose `details` to include more event specific details.<br><br>Events specific fields for other alarm event type can be found with API [List Alarm Definitions]($e/Events%20Definitions/listAlarmDefinitions), under "fields" array of /alarm_defs response object. |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

body := models.WebhookLocationCentrak{
    Events:               []models.WebhookLocationCentrakEvent{
        models.WebhookLocationCentrakEvent{
            MapId:                  models.ToPointer("845a23bf-bed9-e43c-4c86-6fa474be7ae5"),
            WifiBeaconExtendedInfo: []models.WifiBeaconExtendedInfoItems{
                models.WifiBeaconExtendedInfoItems{
                    FrameCtrl:            models.ToPointer(776),
                    Payload:              models.ToPointer("............"),
                    SeqCtrl:              models.ToPointer(772),
                },
            },
            X:                      models.ToPointer(float64(13.5)),
            Y:                      models.ToPointer(float64(3.2)),
            AdditionalProperties:   map[string]interface{}{
                "mac": interface{}("5684dae9ac8b"),
                "site_id": interface{}("4ac1dcf4-9d8b-7211-65c4-057819f0862b"),
                "type": interface{}("wifi"),
            },
        },
    },
    Topic:                "location_centrak",
}

resp, err := samplesWebhooks.LocationCentrak(ctx, &body)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```


# Location Client

Webhook sample for `location_client` topic

**Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages

```go
LocationClient(
    ctx context.Context,
    body *models.WebhookLocationClient) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`*models.WebhookLocationClient`](../../doc/models/webhook-location-client.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

body := models.WebhookLocationClient{
    Events:               []models.WebhookLocationClientEvent{
        models.WebhookLocationClientEvent{
            Mac:                    models.ToPointer("5684dae9ac8b"),
            MapId:                  models.ToPointer(uuid.MustParse("845a23bf-bed9-e43c-4c86-6fa474be7ae5")),
            Type:                   models.ToPointer("wifi"),
            WifiBeaconExtendedInfo: []models.WifiBeaconExtendedInfoItems{
                models.WifiBeaconExtendedInfoItems{
                    FrameCtrl:            models.ToPointer(776),
                    Payload:              models.ToPointer("............"),
                    SeqCtrl:              models.ToPointer(772),
                },
            },
            X:                      models.ToPointer(float64(13.5)),
            Y:                      models.ToPointer(float64(3.2)),
        },
    },
    Topic:                "location_client",
}

resp, err := samplesWebhooks.LocationClient(ctx, &body)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```


# Location Sdk

Webhook sample for `location_sdk` topic

**Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages

```go
LocationSdk(
    ctx context.Context,
    body *models.WebhookLocationSdk) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`*models.WebhookLocationSdk`](../../doc/models/webhook-location-sdk.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

body := models.WebhookLocationSdk{
    Events:               []models.WebhookLocationSdkEvent{
        models.WebhookLocationSdkEvent{
            MapId:                models.ToPointer(uuid.MustParse("845a23bf-bed9-e43c-4c86-6fa474be7ae5")),
            Name:                 models.ToPointer("optional"),
            Type:                 models.ToPointer("sdk"),
            X:                    models.ToPointer(float64(13.5)),
            Y:                    models.ToPointer(float64(3.2)),
        },
    },
    Topic:                "location_sdk",
}

resp, err := samplesWebhooks.LocationSdk(ctx, &body)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```


# Location Unclient

Webhook sample for `location_unclient` topic

**Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages

```go
LocationUnclient(
    ctx context.Context,
    body *models.WebhookLocationUnclient) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`*models.WebhookLocationUnclient`](../../doc/models/webhook-location-unclient.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

body := models.WebhookLocationUnclient{
    Events:               []models.WebhookLocationUnclientEvent{
        models.WebhookLocationUnclientEvent{
            Mac:                    models.ToPointer("5684dae9ac8b"),
            MapId:                  models.ToPointer(uuid.MustParse("845a23bf-bed9-e43c-4c86-6fa474be7ae5")),
            Type:                   models.ToPointer("wifi"),
            WifiBeaconExtendedInfo: []models.WifiBeaconExtendedInfoItems{
                models.WifiBeaconExtendedInfoItems{
                    FrameCtrl:            models.ToPointer(776),
                    Payload:              models.ToPointer("............"),
                    SeqCtrl:              models.ToPointer(772),
                },
            },
            X:                      models.ToPointer(float64(13.5)),
            Y:                      models.ToPointer(float64(3.2)),
        },
    },
    Topic:                "location_unclient",
}

resp, err := samplesWebhooks.LocationUnclient(ctx, &body)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```


# Nac Accounting

Webhook sample for `nac-accounting` topic

**Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages

```go
NacAccounting(
    ctx context.Context,
    body *models.WebhookNacAccounting) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`*models.WebhookNacAccounting`](../../doc/models/webhook-nac-accounting.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

body := models.WebhookNacAccounting{
    Events:               []models.WebhookNacAccountingEvent{
        models.WebhookNacAccountingEvent{
            Ap:                   models.ToPointer("5c5b355005be"),
            AuthType:             models.ToPointer(models.NacAuthTypeEnum_EAPTLS),
            Bssid:                models.ToPointer("5c5b35546bb4"),
            ClientIp:             models.ToPointer("172.16.87.4"),
            ClientType:           models.ToPointer("wireless"),
            Mac:                  models.ToPointer("6e795836d5f9"),
            NasVendor:            models.ToPointer("juniper-mist"),
            Ssid:                 models.ToPointer("Test-CMR SSID"),
            Type:                 models.ToPointer("NAC_ACCOUNTING_STOP"),
            Username:             models.ToPointer("hi"),
        },
    },
    Topic:                models.ToPointer("nac-accounting"),
}

resp, err := samplesWebhooks.NacAccounting(ctx, &body)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```


# Nac Events

Example Delivery of nac_events

```go
NacEvents(
    ctx context.Context,
    body *models.WebhookNacEvents) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`*models.WebhookNacEvents`](../../doc/models/webhook-nac-events.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

body := models.WebhookNacEvents{
    Events:               []models.WebhookNacEventsEvent{
        models.WebhookNacEventsEvent{
            AuthType:             models.ToPointer(models.NacAuthTypeEnum_EAPTEAP),
            Bssid:                models.ToPointer("5c5b355fafcc"),
            IdpRole:              []string{
                "itsuperusers",
                "vip",
            },
            RandomMac:            models.ToPointer(true),
            RespAttrs:            []string{
                "Tunnel-Type=VLAN",
                "Tunnel-Medium-Type=IEEE-802",
                "Tunnel-Private-Group-Id=750",
                "User-Name=anonymous",
            },
        },
    },
    Topic:                models.ToPointer("string"),
}

resp, err := samplesWebhooks.NacEvents(ctx, &body)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```


# Occupancy Alerts

Webhook sample for `occupancy_alerts` topic

**Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages

```go
OccupancyAlerts(
    ctx context.Context,
    body *models.WebhookOccupancyAlerts) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`*models.WebhookOccupancyAlerts`](../../doc/models/webhook-occupancy-alerts.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

body := models.WebhookOccupancyAlerts{
    Events:               []models.WebhookOccupancyAlertsEvent{
        models.WebhookOccupancyAlertsEvent{
            SiteId:               uuid.MustParse(""),
            SiteName:             "",
        },
    },
    Topic:                "occupancy-alerts",
}

resp, err := samplesWebhooks.OccupancyAlerts(ctx, &body)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```


# Ping

Webhook sample for `ping` topic

**Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages

```go
Ping(
    ctx context.Context,
    body *models.WebhookPing) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`*models.WebhookPing`](../../doc/models/webhook-ping.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

body := models.WebhookPing{
    Events:               []models.WebhookPingEvent{
        models.WebhookPingEvent{
            Id:                   uuid.MustParse(""),
            Name:                 "string",
            SiteId:               uuid.MustParse(""),
            Timestamp:            0.0,
        },
    },
    Topic:                "ping",
}

resp, err := samplesWebhooks.Ping(ctx, &body)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```


# Sdkclient Scan Data

Webhook sample for `sdkclient_scan_data` topic

**Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages

```go
SdkclientScanData(
    ctx context.Context,
    body *models.WebhookSdkclientScanData) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`*models.WebhookSdkclientScanData`](../../doc/models/webhook-sdkclient-scan-data.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

body := models.WebhookSdkclientScanData{
    Events:               []models.WebhookSdkclientScanDataEvent{
        models.WebhookSdkclientScanDataEvent{
            ConnectionAp:         "5c5b352f587e",
            ConnectionBand:       "2.4",
            ConnectionBssid:      "5c5b352b51b4",
            ConnectionChannel:    11,
            ConnectionRssi:       float64(-87),
            LastSeen:             nil,
            Mac:                  "70ef0071535f",
            ScanData:             []models.WebhookSdkclientScanDataEventScanDataItem{
                models.WebhookSdkclientScanDataEventScanDataItem{
                    Ap:                   "5c5b352f587e",
                    Band:                 models.ScanDataItemBandEnum_ENUM24,
                    Bssid:                "5c5b352b51b4",
                    Channel:              11,
                    Rssi:                 float64(-87),
                    Ssid:                 "mist-wifi",
                    Timestamp:            0.0,
                },
                models.WebhookSdkclientScanDataEventScanDataItem{
                    Ap:                   "5c5b352f587e",
                    Band:                 models.ScanDataItemBandEnum_ENUM5,
                    Bssid:                "5c5b352b51b8",
                    Channel:              36,
                    Rssi:                 float64(-75),
                    Ssid:                 "mist-wifi",
                    Timestamp:            0.0,
                },
            },
            SiteId:               uuid.MustParse(""),
        },
    },
    Topic:                "sdkclient_scan_data",
}

resp, err := samplesWebhooks.SdkclientScanData(ctx, &body)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```


# Site Sle

Webhook sample for `site_sle` topic

**Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages

```go
SiteSle(
    ctx context.Context,
    body *models.WebhookSiteSle) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`*models.WebhookSiteSle`](../../doc/models/webhook-site-sle.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

body := models.WebhookSiteSle{
    Events:               []models.WebhookSiteSleEvent{
        models.WebhookSiteSleEvent{
            Sle:                  models.ToPointer(models.WebhookSiteSleEventSle{
                ApAvailability:       models.ToPointer(float64(0.6)),
                SuccessfulConnect:    models.ToPointer(float64(0.7)),
                TimeToConnect:        models.ToPointer(float64(0.9)),
            }),
        },
    },
    Topic:                models.ToPointer("site_sle"),
}

resp, err := samplesWebhooks.SiteSle(ctx, &body)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```


# Zone

Webhook sample for `zone` topic

**Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages

```go
Zone(
    ctx context.Context,
    body *models.WebhookZone) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`*models.WebhookZone`](../../doc/models/webhook-zone.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

body := models.WebhookZone{
    Events:               []models.WebhookZoneEvent{
        models.WebhookZoneEvent{
            AssetId:              models.ToPointer(uuid.MustParse("b4695157-0d1d-4da0-8f9e-5c53149389e4")),
            MapId:                uuid.MustParse("09d2b626-2e4e-45ef-a3c4-e6aeb6c83db1"),
            Name:                 models.ToPointer("asset_name"),
            SiteId:               uuid.MustParse(""),
            Timestamp:            0.0,
            Trigger:              models.WebhookZoneEventTriggerEnum_ENTER,
            Type:                 models.WebhookZoneEventTypeEnum_ASSET,
            ZoneId:               uuid.MustParse("4495020a-236f-46e0-9453-e3f9cc6476f4"),
        },
    },
    Topic:                "zone",
}

resp, err := samplesWebhooks.Zone(ctx, &body)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```

