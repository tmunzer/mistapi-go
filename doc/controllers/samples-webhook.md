# Samples-Webhook

```go
samplesWebhook := client.SamplesWebhook()
```

## Class Name

`SamplesWebhook`

## Methods

* [Alarms](../../doc/controllers/samples-webhook.md#alarms)
* [Asset Raw](../../doc/controllers/samples-webhook.md#asset-raw)
* [Audits](../../doc/controllers/samples-webhook.md#audits)
* [Client Info](../../doc/controllers/samples-webhook.md#client-info)
* [Client Join](../../doc/controllers/samples-webhook.md#client-join)
* [Client Latency](../../doc/controllers/samples-webhook.md#client-latency)
* [Client Sessions](../../doc/controllers/samples-webhook.md#client-sessions)
* [Device Events](../../doc/controllers/samples-webhook.md#device-events)
* [Device up Down](../../doc/controllers/samples-webhook.md#device-up-down)
* [Discovered-Raw-Rssi](../../doc/controllers/samples-webhook.md#discovered-raw-rssi)
* [Location](../../doc/controllers/samples-webhook.md#location)
* [Location Asset](../../doc/controllers/samples-webhook.md#location-asset)
* [Location Centrak](../../doc/controllers/samples-webhook.md#location-centrak)
* [Location Client](../../doc/controllers/samples-webhook.md#location-client)
* [Location Sdk](../../doc/controllers/samples-webhook.md#location-sdk)
* [Location Unclient](../../doc/controllers/samples-webhook.md#location-unclient)
* [Nac Accounting](../../doc/controllers/samples-webhook.md#nac-accounting)
* [Nac Events](../../doc/controllers/samples-webhook.md#nac-events)
* [Occupancy Alerts](../../doc/controllers/samples-webhook.md#occupancy-alerts)
* [Ping](../../doc/controllers/samples-webhook.md#ping)
* [Sdkclient Scan Data](../../doc/controllers/samples-webhook.md#sdkclient-scan-data)
* [Site Sle](../../doc/controllers/samples-webhook.md#site-sle)
* [Zone](../../doc/controllers/samples-webhook.md#zone)


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
| `body` | [`*models.WebhookAlarms`](../../doc/models/webhook-alarms.md) | Body, Optional | **N.B.**: Fields like `aps`, `bssids`, `ssids` are event specific. They are relevant to this event type ( rogue-ap-detected). For a different event type, different fields may be sent. These don’t contain all affected entities and are representative samples of entities (capped at 10). For marvis action related events, we expose `details` to include more event specific details.<br><br>Events specific fields for other alarm event type can be found with API https://api.mist.com/api/v1/const/alarm_defs, under “fields” array of /alarm_defs response object. |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

body := models.WebhookAlarms{
    Events: []models.WebhookAlarmEvent{
        models.WebhookAlarmEvent{
            Aps:       []string{
                "string",
            },
            Bssids:    []string{
                "string",
            },
            Count:     models.ToPointer(0),
            EventId:   models.ToPointer(uuid.MustParse("a7a26ff2-e851-45b6-9634-d595f45458b7")),
            ForSite:   models.ToPointer(true),
            Id:        uuid.MustParse("489f6eca-6276-4993-bfeb-c3cbbbba1f08"),
            LastSeen:  float64(0),
            OrgId:     uuid.MustParse("a40f5d1f-d889-42e9-94ea-b9b33585fc6b"),
            SiteId:    uuid.MustParse("72771e6a-6f5e-4de4-a5b9-1266c4197811"),
            Ssids:     []string{
                "string",
            },
            Timestamp: 0,
            Type:      "string",
            Update:    models.ToPointer(true),
        },
    },
    Topic:  "alarms",
}

resp, err := samplesWebhook.Alarms(ctx, &body)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```


# Asset Raw

Webhook sample for `asset_raw` topic

**Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages

**will be deprecated after 06/30/2024**

```go
AssetRaw(
    ctx context.Context,
    body *models.WebhookAssetRaw) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`*models.WebhookAssetRaw`](../../doc/models/webhook-asset-raw.md) | Body, Optional | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

body := models.WebhookAssetRaw{
    Events: []models.WebhookAssetRawEvent{
        models.WebhookAssetRawEvent{
            AssetId:               uuid.MustParse("b4695157-0d1d-4da0-8f9e-5c53149389e4"),
            Beam:                  0,
            DeviceId:              uuid.MustParse("3bafab7b-4400-4bcf-8e6e-09f954699940"),
            IbeaconMajor:          models.ToPointer(0),
            IbeaconMinor:          models.ToPointer(0),
            IbeaconUuid:           models.ToPointer(uuid.MustParse("1f89bc00-d0af-481b-82fe-a6629259a39f")),
            Mac:                   "string",
            MapId:                 uuid.MustParse("09d2b626-2e4e-45ef-a3c4-e6aeb6c83db1"),
            MfgCompanyId:          float64(0),
            MfgData:               "string",
            Rssi:                  float64(0),
            SiteId:                uuid.MustParse("72771e6a-6f5e-4de4-a5b9-1266c4197811"),
            Timestamp:             float64(0),
        },
    },
    Topic:  "asset-raw",
}

resp, err := samplesWebhook.AssetRaw(ctx, &body)
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

``

## Example Usage

```go
ctx := context.Background()

body := models.WebhookAudits{
    Events: []models.WebhookAuditEvent{
        models.WebhookAuditEvent{
            AdminName: "admin_name8",
            DeviceId:  uuid.MustParse("00000380-0000-0000-0000-000000000000"),
            Id:        uuid.MustParse("0000122a-0000-0000-0000-000000000000"),
            Message:   "message0",
            OrgId:     uuid.MustParse("00001302-0000-0000-0000-000000000000"),
            SiteId:    uuid.MustParse("00000290-0000-0000-0000-000000000000"),
            SrcIp:     "src_ip6",
            Timestamp: float64(157.68),
        },
    },
    Topic:  "audits",
}

resp, err := samplesWebhook.Audits(ctx, &body)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```


# Client Info

Webhook sample for `client_info` topic

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

``

## Example Usage

```go
ctx := context.Background()

body := models.WebhookClientInfo{
    Events: []models.WebhookClientInfoEvent{
        models.WebhookClientInfoEvent{
            Mac:       models.ToPointer("string"),
            OrgId:     models.ToPointer(uuid.MustParse("a40f5d1f-d889-42e9-94ea-b9b33585fc6b")),
            SiteId:    models.ToPointer(uuid.MustParse("72771e6a-6f5e-4de4-a5b9-1266c4197811")),
            Timestamp: models.ToPointer(float64(0)),
        },
    },
    Topic:  models.ToPointer(models.WebhookClientInfoTopicEnum("client-info")),
}

resp, err := samplesWebhook.ClientInfo(ctx, &body)
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

``

## Example Usage

```go
ctx := context.Background()

body := models.WebhookClientJoin{
    Events: []models.WebhookClientJoinEvent{
        models.WebhookClientJoinEvent{
            Ap:           "string",
            ApName:       "string",
            Band:         "string",
            Bssid:        "string",
            Connect:      0,
            ConnectFloat: float64(0),
            Mac:          "string",
            OrgId:        uuid.MustParse("a40f5d1f-d889-42e9-94ea-b9b33585fc6b"),
            Rssi:         float64(0),
            SiteId:       uuid.MustParse("72771e6a-6f5e-4de4-a5b9-1266c4197811"),
            SiteName:     "string",
            Ssid:         "string",
            Timestamp:    float64(0),
            Version:      float64(0),
            WlanId:       uuid.MustParse("5028e92b-fc59-4056-91d1-ea4b4ca1617a"),
        },
    },
    Topic:  "client-join",
}

resp, err := samplesWebhook.ClientJoin(ctx, &body)
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

``

## Example Usage

```go
ctx := context.Background()

body := models.WebhookClientLatency{
    Events: []models.WebhookClientLatencyEvent{
        models.WebhookClientLatencyEvent{
            AvgAuth:   models.ToPointer(float64(0.17170219)),
            AvgDhcp:   models.ToPointer(float64(0.017828934)),
            AvgDns:    models.ToPointer(float64(0.024532124)),
            MaxAuth:   models.ToPointer(float64(0.18170219)),
            MaxDhcp:   models.ToPointer(float64(0.027828934)),
            MaxDns:    models.ToPointer(float64(0.029532124)),
            MinAuth:   models.ToPointer(float64(0.16050219)),
            MinDhcp:   models.ToPointer(float64(0.015828934)),
            MinDns:    models.ToPointer(float64(0.022532124)),
            OrgId:     models.ToPointer(uuid.MustParse("2818e386-8dec-2562-9ede-5b8a0fbbdc71")),
            SiteId:    models.ToPointer(uuid.MustParse("4ac1dcf4-9d8b-7211-65c4-057819f0862b")),
            Timestamp: models.ToPointer(1696401600),
        },
    },
    Topic:  models.ToPointer("client-latency"),
}

resp, err := samplesWebhook.ClientLatency(ctx, &body)
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

``

## Example Usage

```go
ctx := context.Background()

body := models.WebhookClientSessions{
    Events: []models.WebhookClientSessionsEvent{
        models.WebhookClientSessionsEvent{
            Ap:                "string",
            ApName:            "string",
            Band:              "string",
            Bssid:             "string",
            ClientFamily:      "string",
            ClientManufacture: "string",
            ClientModel:       "string",
            ClientOs:          "string",
            Connect:           0,
            ConnectFloat:      float64(0),
            Disconnect:        0,
            DisconnectFloat:   float64(0),
            Duration:          0,
            Mac:               "string",
            NextAp:            "string",
            OrgId:             uuid.MustParse("a40f5d1f-d889-42e9-94ea-b9b33585fc6b"),
            Rssi:              float64(0),
            SiteId:            uuid.MustParse("72771e6a-6f5e-4de4-a5b9-1266c4197811"),
            SiteName:          "string",
            Ssid:              "string",
            TerminationReason: 0,
            Timestamp:         float64(0),
            Version:           float64(0),
            WlanId:            uuid.MustParse("5028e92b-fc59-4056-91d1-ea4b4ca1617a"),
        },
    },
    Topic:  "client-sessions",
}

resp, err := samplesWebhook.ClientSessions(ctx, &body)
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

``

## Example Usage

```go
ctx := context.Background()

body := models.WebhookDeviceEvents{
    Events: []models.WebhookDeviceEventsEvent{
        models.WebhookDeviceEventsEvent{
            Ap:         models.ToPointer("string"),
            ApName:     models.ToPointer("string"),
            AuditId:    models.ToPointer(uuid.MustParse("78c04fa6-cfb4-46a0-9aa5-3681ba4f3897")),
            DeviceName: "string",
            DeviceType: models.WebhookDeviceEventsEventDeviceTypeEnum("ap"),
            EvType:     models.WebhookDeviceEventsEventEvTypeEnum("notice"),
            Mac:        "string",
            OrgId:      uuid.MustParse("a40f5d1f-d889-42e9-94ea-b9b33585fc6b"),
            Reason:     models.ToPointer("string"),
            SiteId:     models.ToPointer(uuid.MustParse("72771e6a-6f5e-4de4-a5b9-1266c4197811")),
            SiteName:   models.ToPointer("string"),
            Text:       models.ToPointer("string"),
            Timestamp:  0,
            Type:       "string",
        },
    },
    Topic:  "device_events",
}

resp, err := samplesWebhook.DeviceEvents(ctx, &body)
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

``

## Example Usage

```go
ctx := context.Background()

body := models.WebhookDeviceUpdowns{
    Events: []models.WebhookDeviceUpdownsEvent{
        models.WebhookDeviceUpdownsEvent{
            Ap:        "string",
            ApName:    "string",
            ForSite:   models.ToPointer(true),
            OrgId:     uuid.MustParse("a40f5d1f-d889-42e9-94ea-b9b33585fc6b"),
            SiteId:    uuid.MustParse("72771e6a-6f5e-4de4-a5b9-1266c4197811"),
            SiteName:  "string",
            Timestamp: float64(0),
            Type:      "string",
        },
    },
    Topic:  "device-updowns",
}

resp, err := samplesWebhook.DeviceUpDown(ctx, &body)
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

``

## Example Usage

```go
ctx := context.Background()

body := models.WebhookDiscoveredRawRssi{
    Events: []models.WebhookDiscoveredRawRssiEvent{
        models.WebhookDiscoveredRawRssiEvent{
            ApLoc:          []float64{
                float64(0),
            },
            Beam:           0,
            DeviceId:       uuid.MustParse("3bafab7b-4400-4bcf-8e6e-09f954699940"),
            IbeaconMajor:   models.ToPointer(0),
            IbeaconMinor:   models.ToPointer(0),
            IbeaconUuid:    models.ToPointer(uuid.MustParse("1f89bc00-d0af-481b-82fe-a6629259a39f")),
            IsAsset:        models.ToPointer(true),
            Mac:            "string",
            MapId:          uuid.MustParse("09d2b626-2e4e-45ef-a3c4-e6aeb6c83db1"),
            MfgCompanyId:   models.ToPointer("string"),
            MfgData:        models.ToPointer("string"),
            OrgId:          uuid.MustParse("a40f5d1f-d889-42e9-94ea-b9b33585fc6b"),
            Rssi:           float64(0),
            ServicePackets: []models.ServicePacket{
                models.ServicePacket{
                    ServiceData: models.ToPointer("string"),
                    ServiceUuid: models.ToPointer("7138cc00-c611-4dec-a05e-5c4b1cae13c0"),
                },
            },
            SiteId:         uuid.MustParse("72771e6a-6f5e-4de4-a5b9-1266c4197811"),
            Timestamp:      models.ToPointer(0),
        },
    },
    Topic:  "string",
}

resp, err := samplesWebhook.DiscoveredRawRssi(ctx, &body)
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

``

## Example Usage

```go
ctx := context.Background()

body := models.WebhookLocation{
    Events: []models.WebhookLocationEvent{
        models.WebhookLocationEvent{
            BatteryVoltage:         models.ToPointer(0),
            EddystoneUidInstance:   models.ToPointer("string"),
            EddystoneUidNamespace:  models.ToPointer("string"),
            EddystoneUrlUrl:        models.ToPointer("string"),
            IbeaconMajor:           models.ToPointer(0),
            IbeaconMinor:           models.ToPointer(0),
            IbeaconUuid:            models.ToPointer(uuid.MustParse("1f89bc00-d0af-481b-82fe-a6629259a39f")),
            Id:                     uuid.MustParse("487f6eca-6276-4993-bfeb-e3cbbbba3f08"),
            Mac:                    models.ToPointer("string"),
            MapId:                  uuid.MustParse("09d2b626-2e4e-45ef-a3c4-e6aeb6c83db1"),
            MfgCompanyId:           models.ToPointer(0),
            MfgData:                models.ToPointer("string"),
            Name:                   models.ToPointer("string"),
            SiteId:                 uuid.MustParse("72771e6a-6f5e-4de4-a5b9-1266c4197811"),
            Timestamp:              0,
            Type:                   "string",
            X:                      float64(0),
            Y:                      float64(0),
        },
    },
    Topic:  "location",
}

resp, err := samplesWebhook.Location(ctx, &body)
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

``

## Example Usage

```go
ctx := context.Background()

body := models.WebhookLocationAsset{
    Events: []models.WebhookLocationAssetEvent{
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
            SiteId:                models.ToPointer(uuid.MustParse("4ac1dcf4-9d8b-7211-65c4-057819f0862b")),
            Timestamp:             models.ToPointer(1461220784),
            Type:                  models.ToPointer("asset"),
            X:                     models.ToPointer(float64(13.5)),
            Y:                     models.ToPointer(float64(3.2)),
        },
    },
    Topic:  "location_asset",
}

resp, err := samplesWebhook.LocationAsset(ctx, &body)
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
| `body` | [`*models.WebhookLocationCentrak`](../../doc/models/webhook-location-centrak.md) | Body, Optional | **N.B.**: Fields like `aps`, `bssids`, `ssids` are event specific. They are relevant to this event type ( rogue-ap-detected). For a different event type, different fields may be sent. These don’t contain all affected entities and are representative samples of entities (capped at 10). For marvis action related events, we expose `details` to include more event specific details.<br><br>Events specific fields for other alarm event type can be found with API https://api.mist.com/api/v1/const/alarm_defs, under “fields” array of /alarm_defs response object. |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

body := models.WebhookLocationCentrak{
    Events: []models.WebhookLocationCentrakEvent{
        models.WebhookLocationCentrakEvent{
            MapId:                  models.ToPointer("845a23bf-bed9-e43c-4c86-6fa474be7ae5"),
            Timestamp:              models.ToPointer(1461220784),
            WifiBeaconExtendedInfo: []models.WifiBeaconExtendedInfoItems{
                models.WifiBeaconExtendedInfoItems{
                    FrameCtrl: models.ToPointer(776),
                    Payload:   models.ToPointer("............"),
                    SeqCtrl:   models.ToPointer(772),
                },
            },
            X:                      models.ToPointer(float64(13.5)),
            Y:                      models.ToPointer(float64(3.2)),
        },
    },
    Topic:  "location_centrak",
}

resp, err := samplesWebhook.LocationCentrak(ctx, &body)
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

``

## Example Usage

```go
ctx := context.Background()

body := models.WebhookLocationClient{
    Events: []models.WebhookLocationClientEvent{
        models.WebhookLocationClientEvent{
            Mac:                    models.ToPointer("5684dae9ac8b"),
            MapId:                  models.ToPointer(uuid.MustParse("845a23bf-bed9-e43c-4c86-6fa474be7ae5")),
            SiteId:                 models.ToPointer(uuid.MustParse("4ac1dcf4-9d8b-7211-65c4-057819f0862b")),
            Timestamp:              models.ToPointer(1461220784),
            Type:                   models.ToPointer("wifi"),
            WifiBeaconExtendedInfo: []models.WifiBeaconExtendedInfoItems{
                models.WifiBeaconExtendedInfoItems{
                    FrameCtrl: models.ToPointer(776),
                    Payload:   models.ToPointer("............"),
                    SeqCtrl:   models.ToPointer(772),
                },
            },
            X:                      models.ToPointer(float64(13.5)),
            Y:                      models.ToPointer(float64(3.2)),
        },
    },
    Topic:  "location_client",
}

resp, err := samplesWebhook.LocationClient(ctx, &body)
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

``

## Example Usage

```go
ctx := context.Background()

body := models.WebhookLocationSdk{
    Events: []models.WebhookLocationSdkEvent{
        models.WebhookLocationSdkEvent{
            Id:        models.ToPointer(uuid.MustParse("de87bf9d-183f-e383-cc68-6ba43947d403")),
            MapId:     models.ToPointer(uuid.MustParse("845a23bf-bed9-e43c-4c86-6fa474be7ae5")),
            Name:      models.ToPointer("optional"),
            SiteId:    models.ToPointer(uuid.MustParse("4ac1dcf4-9d8b-7211-65c4-057819f0862b")),
            Timestamp: models.ToPointer(1461220784),
            Type:      models.ToPointer("sdk"),
            X:         models.ToPointer(float64(13.5)),
            Y:         models.ToPointer(float64(3.2)),
        },
    },
    Topic:  "location_sdk",
}

resp, err := samplesWebhook.LocationSdk(ctx, &body)
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

``

## Example Usage

```go
ctx := context.Background()

body := models.WebhookLocationUnclient{
    Events: []models.WebhookLocationUnclientEvent{
        models.WebhookLocationUnclientEvent{
            Mac:                    models.ToPointer("5684dae9ac8b"),
            MapId:                  models.ToPointer(uuid.MustParse("845a23bf-bed9-e43c-4c86-6fa474be7ae5")),
            SiteId:                 models.ToPointer(uuid.MustParse("4ac1dcf4-9d8b-7211-65c4-057819f0862b")),
            Timestamp:              models.ToPointer(1461220784),
            Type:                   models.ToPointer("wifi"),
            WifiBeaconExtendedInfo: []models.WifiBeaconExtendedInfoItems{
                models.WifiBeaconExtendedInfoItems{
                    FrameCtrl: models.ToPointer(776),
                    Payload:   models.ToPointer("............"),
                    SeqCtrl:   models.ToPointer(772),
                },
            },
            X:                      models.ToPointer(float64(13.5)),
            Y:                      models.ToPointer(float64(3.2)),
        },
    },
    Topic:  "location_unclient",
}

resp, err := samplesWebhook.LocationUnclient(ctx, &body)
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

``

## Example Usage

```go
ctx := context.Background()

body := models.WebhookNacAccounting{
    Events: []models.WebhookNacAccountingEvent{
        models.WebhookNacAccountingEvent{
            Ap:         models.ToPointer("5c5b355005be"),
            AuthType:   models.ToPointer("eap-tls"),
            Bssid:      models.ToPointer("5c5b35546bb4"),
            ClientIp:   models.ToPointer("172.16.87.4"),
            ClientType: models.ToPointer("wireless"),
            Mac:        models.ToPointer("6e795836d5f9"),
            NasVendor:  models.ToPointer("juniper-mist"),
            OrgId:      models.ToPointer(uuid.MustParse("625aba64-fe72-4b14-8985-cbf31ec3d78a")),
            RxPkts:     models.ToPointer(9523),
            SiteId:     models.ToPointer(uuid.MustParse("ec9d1e85-af24-43f9-8d65-d620580e8631")),
            Ssid:       models.ToPointer("Test-CMR SSID"),
            Timestamp:  models.ToPointer(float64(1547235620.89)),
            TxPkts:     models.ToPointer(5233),
            Type:       models.ToPointer("NAC_ACCOUNTING_STOP"),
            Username:   models.ToPointer("hi"),
        },
    },
    Topic:  models.ToPointer("nac-accounting"),
}

resp, err := samplesWebhook.NacAccounting(ctx, &body)
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

``

## Example Usage

```go
ctx := context.Background()

body := models.WebhookNacEvents{
    Events: []models.WebhookNacEventsEvent{
        models.WebhookNacEventsEvent{
            Ap:                   models.ToPointer("5c5b35513227"),
            AuthType:             models.ToPointer("802.1X"),
            Bssid:                models.ToPointer("5c5b355fafcc"),
            DryrunNacruleId:      models.ToPointer(uuid.MustParse("32f27e7d-ff26-4a9b-b3d1-ff9bcb264012")),
            DryrunNacruleMatched: models.ToPointer(true),
            IdpId:                models.ToPointer(uuid.MustParse("912ef72e-2239-4996-b81e-469e87a27cd6")),
            IdpRole:              []string{
                "itsuperusers",
                "vip",
            },
            Mac:                  models.ToPointer("ac3eb179e535"),
            NacruleId:            models.ToPointer(uuid.MustParse("32f27e7d-ff26-4a9b-b3d1-ff9bcb264c62")),
            NacruleMatched:       models.ToPointer(true),
            NasVendor:            models.ToPointer("juniper-mist"),
            OrgId:                models.ToPointer(uuid.MustParse("27547ac2-d114-4e04-beb1-f3f1e6e81ec6")),
            RandomMac:            models.ToPointer(true),
            RespAttrs:            []string{
                "Tunnel-Type=VLAN",
                "Tunnel-Medium-Type=IEEE-802",
                "Tunnel-Private-Group-Id=750",
                "User-Name=anonymous",
            },
            SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
            Ssid:                 models.ToPointer("##mist_nac"),
            Timestamp:            models.ToPointer(float64(1691512031.358188)),
            Type:                 models.ToPointer("NAC_CLIENT_PERMIT"),
            Username:             models.ToPointer("user@deaflyz.net"),
            Vlan:                 models.ToPointer("750"),
        },
    },
    Topic:  models.ToPointer("string"),
}

resp, err := samplesWebhook.NacEvents(ctx, &body)
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

``

## Example Usage

```go
ctx := context.Background()

body := models.WebhookOccupancyAlerts{
    Events: []models.WebhookOccupancyAlertsEvent{
        models.WebhookOccupancyAlertsEvent{
            AlertEvents: []models.WebhookOccupancyAlertsEventAlertEventsItems{
                models.WebhookOccupancyAlertsEventAlertEventsItems{
                    CurrentOccupancy: 0,
                    MapId:            uuid.MustParse("09d2b626-2e4e-45ef-a3c4-e6aeb6c83db1"),
                    OccupancyLimit:   0,
                    OrgId:            uuid.MustParse("a40f5d1f-d889-42e9-94ea-b9b33585fc6b"),
                    Timestamp:        float64(0),
                    Type:             models.WebhookOccupancyAlertTypeEnum("COMPLIANCE-VIOLATION"),
                    ZoneId:           uuid.MustParse("4495020a-236f-46e0-9453-e3f9cc6476f4"),
                    ZoneName:         "string",
                },
            },
            ForSite:     models.ToPointer(true),
            SiteId:      uuid.MustParse("72771e6a-6f5e-4de4-a5b9-1266c4197811"),
            SiteName:    "string",
        },
    },
    Topic:  "occupancy-alerts",
}

resp, err := samplesWebhook.OccupancyAlerts(ctx, &body)
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

``

## Example Usage

```go
ctx := context.Background()

body := models.WebhookPing{
    Events: []models.WebhookPingEvent{
        models.WebhookPingEvent{
            Id:        uuid.MustParse("487f6eca-6276-4993-bfeb-f3cbbbba4f08"),
            Name:      "string",
            SiteId:    uuid.MustParse("72771e6a-6f5e-4de4-a5b9-1266c4197811"),
            Timestamp: float64(0),
        },
    },
    Topic:  "ping",
}

resp, err := samplesWebhook.Ping(ctx, &body)
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

``

## Example Usage

```go
ctx := context.Background()

body := models.WebhookSdkclientScanData{
    Events: []models.WebhookSdkclientScanDataEvent{
        models.WebhookSdkclientScanDataEvent{
            ConnectionAp:      "5c5b352f587e",
            ConnectionBand:    "2.4",
            ConnectionBssid:   "5c5b352b51b4",
            ConnectionChannel: 11,
            ConnectionRssi:    float64(-87),
            LastSeen:          float64(1592333828),
            Mac:               "70ef0071535f",
            ScanData:          []models.WebhookSdkclientScanDataEventScanDataItem{
                models.WebhookSdkclientScanDataEventScanDataItem{
                    Ap:        "5c5b352f587e",
                    Band:      models.ScanDataItemBandEnum("2.4"),
                    Bssid:     "5c5b352b51b4",
                    Channel:   11,
                    Rssi:      float64(-87),
                    Ssid:      "mist-wifi",
                    Timestamp: float64(1592333828),
                },
                models.WebhookSdkclientScanDataEventScanDataItem{
                    Ap:        "5c5b352f587e",
                    Band:      models.ScanDataItemBandEnum("5"),
                    Bssid:     "5c5b352b51b8",
                    Channel:   36,
                    Rssi:      float64(-75),
                    Ssid:      "mist-wifi",
                    Timestamp: float64(1592333828),
                },
            },
            SiteId:            uuid.MustParse("93986f10-773b-42be-9438-8d3e6d127f1a"),
        },
    },
    Topic:  "sdkclient_scan_data",
}

resp, err := samplesWebhook.SdkclientScanData(ctx, &body)
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

``

## Example Usage

```go
ctx := context.Background()

body := models.WebhookSiteSle{
    Events: []models.WebhookSiteSleEvent{
        models.WebhookSiteSleEvent{
            OrgId:     models.ToPointer(uuid.MustParse("2818e386-8dec-2562-9ede-5b8a0fbbdc71")),
            SiteId:    models.ToPointer(uuid.MustParse("4ac1dcf4-9d8b-7211-65c4-057819f0862b")),
            Sle:       models.ToPointer(models.WebhookSiteSleEventSle{
                ApAvailability:    models.ToPointer(float64(0.6)),
                SuccessfulConnect: models.ToPointer(float64(0.7)),
                TimeToConnect:     models.ToPointer(float64(0.9)),
            }),
            Timestamp: models.ToPointer(1694620800),
        },
    },
    Topic:  models.ToPointer("site_sle"),
}

resp, err := samplesWebhook.SiteSle(ctx, &body)
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

``

## Example Usage

```go
ctx := context.Background()

body := models.WebhookZone{
    Events: []models.WebhookZoneEvent{
        models.WebhookZoneEvent{
            AssetId:   models.ToPointer(uuid.MustParse("b4695157-0d1d-4da0-8f9e-5c53149389e4")),
            Id:        uuid.MustParse("485f6eca-6276-4993-bfeb-54cbbbba5f08"),
            Mac:       models.ToPointer("string"),
            MapId:     uuid.MustParse("09d2b626-2e4e-45ef-a3c4-e6aeb6c83db1"),
            Name:      models.ToPointer("string"),
            SiteId:    uuid.MustParse("72771e6a-6f5e-4de4-a5b9-1266c4197811"),
            Timestamp: 0,
            Trigger:   models.WebhookZoneEventTriggerEnum("enter"),
            Type:      "string",
            ZoneId:    uuid.MustParse("4495020a-236f-46e0-9453-e3f9cc6476f4"),
        },
    },
    Topic:  "zone",
}

resp, err := samplesWebhook.Zone(ctx, &body)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```

