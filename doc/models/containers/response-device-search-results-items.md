
# Response Device Search Results Items

Device search record for an AP, switch, or gateway

## Class Name

`ResponseDeviceSearchResultsItems`

## Cases

| Type | Factory Method |
|  --- | --- |
| [`models.ApSearch`](../../../doc/models/ap-search.md) | models.ResponseDeviceSearchResultsItemsContainer.FromApSearch(models.ApSearch apSearch) |
| [`models.SwitchSearch`](../../../doc/models/switch-search.md) | models.ResponseDeviceSearchResultsItemsContainer.FromSwitchSearch(models.SwitchSearch switchSearch) |
| [`models.GatewaySearch`](../../../doc/models/gateway-search.md) | models.ResponseDeviceSearchResultsItemsContainer.FromGatewaySearch(models.GatewaySearch gatewaySearch) |

## models.ApSearch

### Initialization Code

#### Example

```go
value := models.ResponseDeviceSearchResultsItemsContainer.FromApSearch(models.ApSearch{
    MxtunnelStatus:       "mxtunnel_status4",
    OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
    PowerConstrained:     false,
    PowerOpmode:          "power_opmode0",
    SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
    Wlans:                []models.ApSearchWlan{
        models.ApSearchWlan{
        },
    },
})
```

## models.SwitchSearch

### Initialization Code

#### Example

```go
value := models.ResponseDeviceSearchResultsItemsContainer.FromSwitchSearch(models.SwitchSearch{
    OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
    SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
    Type:                 "switch",
})
```

## models.GatewaySearch

### Initialization Code

#### Example

```go
value := models.ResponseDeviceSearchResultsItemsContainer.FromGatewaySearch(models.GatewaySearch{
    OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
    SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
    Type:                 "gateway",
})
```

