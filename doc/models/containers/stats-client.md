
# Stats Client

Client statistics record for a wireless or wired client

## Class Name

`StatsClient`

## Cases

| Type | Factory Method |
|  --- | --- |
| [`models.StatsWirelessClient`](../../../doc/models/stats-wireless-client.md) | models.StatsClientContainer.FromStatsWirelessClient(models.StatsWirelessClient statsWirelessClient) |
| [`models.StatsWiredClient`](../../../doc/models/stats-wired-client.md) | models.StatsClientContainer.FromStatsWiredClient(models.StatsWiredClient statsWiredClient) |

## models.StatsWirelessClient

### Initialization Code

#### Example

```go
value := models.StatsClientContainer.FromStatsWirelessClient(models.StatsWirelessClient{
    ApId:                 uuid.MustParse("00001902-0000-0000-0000-000000000000"),
    ApMac:                "ap_mac2",
    Band:                 models.Dot11BandEnum_ENUM5,
    Channel:              152,
    IsGuest:              false,
    KeyMgmt:              "key_mgmt2",
    LastSeen:             models.NewOptional(models.ToPointer(float64(1470417522))),
    Mac:                  "mac4",
    Proto:                models.Dot11ProtoEnum_B,
    Rssi:                 float64(66.02),
    RxBps:                models.NewOptional(models.ToPointer(int64(60003))),
    RxBytes:              models.NewOptional(models.ToPointer(int64(8515104416))),
    RxPkts:               models.NewOptional(models.ToPointer(int64(57770567))),
    SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
    Snr:                  float64(64.24),
    Ssid:                 "ssid8",
    TxBps:                models.NewOptional(models.ToPointer(int64(634301))),
    TxBytes:              models.NewOptional(models.ToPointer(int64(211217389682))),
    TxPkts:               models.NewOptional(models.ToPointer(int64(812204062))),
    WlanId:               uuid.MustParse("000004a8-0000-0000-0000-000000000000"),
})
```

## models.StatsWiredClient

### Initialization Code

#### Example

```go
value := models.StatsClientContainer.FromStatsWiredClient(models.StatsWiredClient{
    Mac:                  "mac0",
    RxBytes:              models.NewOptional(models.ToPointer(int64(8515104416))),
    RxPkts:               models.NewOptional(models.ToPointer(int64(57770567))),
    SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
    TxBytes:              models.NewOptional(models.ToPointer(int64(211217389682))),
    TxPkts:               models.NewOptional(models.ToPointer(int64(812204062))),
})
```

