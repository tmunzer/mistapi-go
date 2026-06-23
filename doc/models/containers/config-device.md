
# Config Device

Device configuration object for an AP, switch, or gateway

## Class Name

`ConfigDevice`

## Cases

| Type | Factory Method |
|  --- | --- |
| [`models.DeviceAp`](../../../doc/models/device-ap.md) | models.ConfigDeviceContainer.FromDeviceAp(models.DeviceAp deviceAp) |
| [`models.DeviceSwitch`](../../../doc/models/device-switch.md) | models.ConfigDeviceContainer.FromDeviceSwitch(models.DeviceSwitch deviceSwitch) |
| [`models.DeviceGateway`](../../../doc/models/device-gateway.md) | models.ConfigDeviceContainer.FromDeviceGateway(models.DeviceGateway deviceGateway) |

## models.DeviceAp

### Initialization Code

#### Example

```go
value := models.ConfigDeviceContainer.FromDeviceAp(models.DeviceAp{
    DeviceprofileId:      models.NewOptional(models.ToPointer(uuid.MustParse("6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9"))),
    DisableEth1:          models.ToPointer(false),
    DisableEth2:          models.ToPointer(false),
    DisableEth3:          models.ToPointer(false),
    DisableModule:        models.ToPointer(false),
    FlowControl:          models.ToPointer(false),
    Height:               models.ToPointer(float64(2.75)),
    Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
    MapId:                models.ToPointer(uuid.MustParse("63eda950-c6da-11e4-a628-60f81dd250cc")),
    Name:                 models.ToPointer("conference room"),
    Notes:                models.ToPointer("slightly off center"),
    OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
    Orientation:          models.ToPointer(45),
    PoePassthrough:       models.ToPointer(false),
    SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
    Type:                 "ap",
    Vars:                 map[string]string{
        "RADIUS_IP1": "172.31.2.5",
        "RADIUS_SECRET": "11s64632d",
    },
    X:                    models.ToPointer(float64(53.5)),
    Y:                    models.ToPointer(float64(173.1)),
})
```

## models.DeviceSwitch

### Initialization Code

#### Example

```go
value := models.ConfigDeviceContainer.FromDeviceSwitch(models.DeviceSwitch{
    AggregateRoutes:       map[string]models.AggregateRoute{
        "172.16.3.0/24": models.AggregateRoute{
            Discard:              models.ToPointer(false),
            Metric:               models.NewOptional[int](nil),
            Preference:           models.NewOptional(models.ToPointer(30)),
        },
    },
    AggregateRoutes6:      map[string]models.AggregateRoute{
        "2a02:1234:420a:10c9::/64": models.AggregateRoute{
            Discard:              models.ToPointer(false),
            Metric:               models.NewOptional[int](nil),
            Preference:           models.NewOptional(models.ToPointer(30)),
        },
    },
    DefaultPortUsage:      models.ToPointer("default"),
    DisableAutoConfig:     models.ToPointer(false),
    ExtraRoutes:           map[string]models.ExtraRoute{
        "0.0.0.0/0": models.ExtraRoute{
            Via:                  models.ToPointer(models.NextHopViaContainer.FromString("192.168.1.10")),
        },
    },
    ExtraRoutes6:          map[string]models.ExtraRoute6{
        "2a02:1234:420a:10c9::/64": models.ExtraRoute6{
            Via:                  models.ToPointer(models.NextHopViaContainer.FromString("2a02:1234:200a::100")),
        },
    },
    Id:                    models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
    Managed:               models.ToPointer(false),
    MapId:                 models.ToPointer(uuid.MustParse("63eda950-c6da-11e4-a628-60f81dd250cc")),
    OrgId:                 models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
    RouterId:              models.ToPointer("10.2.1.10"),
    SiteId:                models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
    Type:                  "switch",
    UseRouterIdAsSourceIp: models.ToPointer(false),
    Vars:                  map[string]string{
        "RADIUS_IP1": "172.31.2.5",
        "RADIUS_SECRET": "11s64632d",
    },
    VrfInstances:          map[string]models.SwitchVrfInstance{
        "guest": models.SwitchVrfInstance{
            ExtraRoutes:             map[string]models.VrfExtraRoute{
                "0.0.0.0/0": models.VrfExtraRoute{
                    Via:                  models.ToPointer("192.168.31.1"),
                },
            },
            Networks:                []string{
                "guest",
            },
        },
    },
    X:                     models.ToPointer(float64(53.5)),
    Y:                     models.ToPointer(float64(173.1)),
})
```

## models.DeviceGateway

### Initialization Code

#### Example

```go
value := models.ConfigDeviceContainer.FromDeviceGateway(models.DeviceGateway{
    ExtraRoutes6:            map[string]models.GatewayExtraRoute6{
        "2a02:1234:420a:10c9::/64": models.GatewayExtraRoute6{
            Via:                  models.ToPointer("2a02:1234:200a::100"),
        },
    },
    Id:                      models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
    MapId:                   models.ToPointer(uuid.MustParse("63eda950-c6da-11e4-a628-60f81dd250cc")),
    MspId:                   models.ToPointer(uuid.MustParse("b9d42c2e-88ee-41f8-b798-f009ce7fe909")),
    OrgId:                   models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
    RouterId:                models.ToPointer("10.2.1.10"),
    SiteId:                  models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
    Type:                    "gateway",
    UrlFilteringDenyMsg:     models.ToPointer("Access to this URL Category has been blocked"),
    Vars:                    map[string]string{
        "RADIUS_IP1": "172.31.2.5",
        "RADIUS_SECRET": "11s64632d",
    },
    VrfInstances:            map[string]models.GatewayVrfInstance{
        "CORP_VRF": models.GatewayVrfInstance{
            Networks:             []string{
                "CORP_NET",
                "MGMT_NET",
            },
        },
    },
    X:                       models.ToPointer(float64(53.5)),
    Y:                       models.ToPointer(float64(173.1)),
})
```

