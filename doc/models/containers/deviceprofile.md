
# Deviceprofile

Device profile configuration for an AP, switch, or gateway

## Class Name

`Deviceprofile`

## Cases

| Type | Factory Method |
|  --- | --- |
| [`models.DeviceprofileAp`](../../../doc/models/deviceprofile-ap.md) | models.DeviceprofileContainer.FromDeviceprofileAp(models.DeviceprofileAp deviceprofileAp) |
| [`models.DeviceprofileGateway`](../../../doc/models/deviceprofile-gateway.md) | models.DeviceprofileContainer.FromDeviceprofileGateway(models.DeviceprofileGateway deviceprofileGateway) |
| [`models.DeviceprofileSwitch`](../../../doc/models/deviceprofile-switch.md) | models.DeviceprofileContainer.FromDeviceprofileSwitch(models.DeviceprofileSwitch deviceprofileSwitch) |

## models.DeviceprofileAp

### Initialization Code

#### Example

```go
value := models.DeviceprofileContainer.FromDeviceprofileAp(models.DeviceprofileAp{
    DisableEth1:          models.ToPointer(false),
    DisableEth2:          models.ToPointer(false),
    DisableEth3:          models.ToPointer(false),
    DisableModule:        models.ToPointer(false),
    Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
    OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
    PoePassthrough:       models.ToPointer(false),
    SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
    Type:                 "ap",
    Vars:                 map[string]string{
        "RADIUS_IP1": "172.31.2.5",
        "RADIUS_SECRET": "11s64632d",
    },
})
```

## models.DeviceprofileGateway

### Initialization Code

#### Example

```go
value := models.DeviceprofileContainer.FromDeviceprofileGateway(models.DeviceprofileGateway{
    DnsOverride:             models.ToPointer(false),
    ExtraRoutes6:            map[string]models.GatewayExtraRoute6{
        "2a02:1234:420a:10c9::/64": models.GatewayExtraRoute6{
            Via:                  models.ToPointer("2a02:1234:200a::100"),
        },
    },
    Id:                      models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
    Name:                    "gw_template",
    NtpOverride:             models.ToPointer(false),
    OrgId:                   models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
    RouterId:                models.ToPointer("10.2.1.10"),
    Type:                    "gateway",
    UrlFilteringDenyMsg:     models.ToPointer("Access to this URL Category has been blocked"),
    VrfInstances:            map[string]models.GatewayVrfInstance{
        "CORP_VRF": models.GatewayVrfInstance{
            Networks:             []string{
                "CORP_NET",
                "MGMT_NET",
            },
        },
    },
})
```

## models.DeviceprofileSwitch

### Initialization Code

#### Example

```go
value := models.DeviceprofileContainer.FromDeviceprofileSwitch(models.DeviceprofileSwitch{
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
    Name:                  "name4",
    OrgId:                 models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
    SiteId:                models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
    Type:                  "switch",
    UseRouterIdAsSourceIp: models.ToPointer(false),
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
})
```

