
# Stats Device

Device statistics record for an AP, switch, or gateway

## Class Name

`StatsDevice`

## Cases

| Type | Factory Method |
|  --- | --- |
| [`models.StatsAp`](../../../doc/models/stats-ap.md) | models.StatsDeviceContainer.FromStatsAp(models.StatsAp statsAp) |
| [`models.StatsSwitch`](../../../doc/models/stats-switch.md) | models.StatsDeviceContainer.FromStatsSwitch(models.StatsSwitch statsSwitch) |
| [`models.StatsGateway`](../../../doc/models/stats-gateway.md) | models.StatsDeviceContainer.FromStatsGateway(models.StatsGateway statsGateway) |

## models.StatsAp

### Initialization Code

#### Example

```go
value := models.StatsDeviceContainer.FromStatsAp(models.StatsAp{
    AntennaSelect:        models.ToPointer(models.AntennaSelectEnum_EXTERNAL),
    CertExpiry:           models.NewOptional(models.ToPointer(float64(1534534392))),
    ExtIp:                models.NewOptional(models.ToPointer("73.92.124.103")),
    Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
    IotStat:              map[string]models.StatsApIotStatAdditionalProperties{
        "DI2": models.StatsApIotStatAdditionalProperties{
            Value:                models.NewOptional(models.ToPointer(0)),
        },
    },
    Ip:                   models.NewOptional(models.ToPointer("10.2.9.159")),
    LastSeen:             models.NewOptional(models.ToPointer(float64(1470417522))),
    Locating:             models.NewOptional(models.ToPointer(false)),
    Locked:               models.NewOptional(models.ToPointer(true)),
    Mac:                  models.NewOptional(models.ToPointer("5c5b35000010")),
    MapId:                models.NewOptional(models.ToPointer(uuid.MustParse("63eda950-c6da-11e4-a628-60f81dd250cc"))),
    Model:                models.NewOptional(models.ToPointer("AP200")),
    Mount:                models.NewOptional(models.ToPointer("faceup")),
    Name:                 models.NewOptional(models.ToPointer("conference room")),
    OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
    PowerBudget:          models.NewOptional(models.ToPointer(1000)),
    PowerConstrained:     models.NewOptional(models.ToPointer(false)),
    PowerOpmode:          models.NewOptional(models.ToPointer("[20] 6GHz(2x2) 5GHz(4x4) 2.4GHz(2x2).")),
    PowerSrc:             models.NewOptional(models.ToPointer("PoE 802.3af")),
    RxBps:                models.NewOptional(models.ToPointer(int64(60003))),
    RxBytes:              models.NewOptional(models.ToPointer(int64(8515104416))),
    RxPkts:               models.NewOptional(models.ToPointer(int64(57770567))),
    Serial:               models.NewOptional(models.ToPointer("FXLH2015170017")),
    SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
    TxBps:                models.NewOptional(models.ToPointer(int64(634301))),
    TxBytes:              models.NewOptional(models.ToPointer(int64(211217389682))),
    TxPkts:               models.NewOptional(models.ToPointer(int64(812204062))),
    Type:                 "ap",
    Uptime:               models.NewOptional(models.ToPointer(float64(13500))),
    Version:              models.NewOptional(models.ToPointer("0.14.12345")),
    X:                    models.NewOptional(models.ToPointer(float64(53.5))),
    Y:                    models.NewOptional(models.ToPointer(float64(173.1))),
})
```

## models.StatsSwitch

### Initialization Code

#### Example

```go
value := models.StatsDeviceContainer.FromStatsSwitch(models.StatsSwitch{
    HasPcap:              models.ToPointer(false),
    Hostname:             models.ToPointer("sj-sw1"),
    Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
    Ip:                   models.ToPointer("10.2.11.137"),
    LastSeen:             models.NewOptional(models.ToPointer(float64(1470417522))),
    Model:                models.ToPointer("EX4600"),
    Name:                 models.ToPointer("sj-sw1"),
    OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
    Serial:               models.ToPointer("TC3714190003"),
    SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
    Status:               models.ToPointer("connected"),
    Type:                 "switch",
    Uptime:               models.NewOptional(models.ToPointer(float64(13501))),
    Version:              models.NewOptional(models.ToPointer("18.4R1.8")),
})
```

## models.StatsGateway

### Initialization Code

#### Example

```go
value := models.StatsDeviceContainer.FromStatsGateway(models.StatsGateway{
    ExtIp:                models.NewOptional(models.ToPointer("66.129.234.224")),
    Hostname:             models.ToPointer("sj1"),
    Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
    Ip:                   models.NewOptional(models.ToPointer("10.2.11.137")),
    LastSeen:             models.NewOptional(models.ToPointer(float64(1470417522))),
    Mac:                  "dc38e1dbf3cd",
    Model:                models.ToPointer("SRX320"),
    Name:                 models.ToPointer("sj1"),
    NodeName:             models.ToPointer("node0"),
    OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
    RouterName:           models.ToPointer("sj1"),
    Serial:               models.ToPointer("TC3714190003"),
    SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
    Status:               models.ToPointer("connected"),
    Type:                 "gateway",
    Uptime:               models.NewOptional(models.ToPointer(float64(3671219))),
    Version:              models.NewOptional(models.ToPointer("18.4R1.8")),
})
```

