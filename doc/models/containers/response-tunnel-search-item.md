
# Response Tunnel Search Item

Tunnel statistics record; shape depends on the requested tunnel type

## Class Name

`ResponseTunnelSearchItem`

## Cases

| Type | Factory Method |
|  --- | --- |
| [`models.StatsMxtunnel`](../../../doc/models/stats-mxtunnel.md) | models.ResponseTunnelSearchItemContainer.FromStatsMxtunnel(models.StatsMxtunnel statsMxtunnel) |
| [`models.StatsWanTunnel`](../../../doc/models/stats-wan-tunnel.md) | models.ResponseTunnelSearchItemContainer.FromStatsWanTunnel(models.StatsWanTunnel statsWanTunnel) |

## models.StatsMxtunnel

### Initialization Code

#### Example

```go
value := models.ResponseTunnelSearchItemContainer.FromStatsMxtunnel(models.StatsMxtunnel{
    LastSeen:             models.NewOptional(models.ToPointer(float64(1470417522))),
    OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
    RemoteIp:             "",
    SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
})
```

## models.StatsWanTunnel

### Initialization Code

#### Example

```go
value := models.ResponseTunnelSearchItemContainer.FromStatsWanTunnel(models.StatsWanTunnel{
    OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
    PeerIp:               "peer_ip6",
    RxBytes:              models.NewOptional(models.ToPointer(int64(8515104416))),
    RxPkts:               models.NewOptional(models.ToPointer(int64(57770567))),
    SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
    TxBytes:              models.NewOptional(models.ToPointer(int64(211217389682))),
    TxPkts:               models.NewOptional(models.ToPointer(int64(812204062))),
    WanName:              models.ToPointer("wan"),
})
```

