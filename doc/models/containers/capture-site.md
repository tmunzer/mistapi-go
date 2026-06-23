
# Capture Site

Packet capture request payload for site-level captures

## Class Name

`CaptureSite`

## Cases

| Type | Factory Method |
|  --- | --- |
| [`models.CaptureClient`](../../../doc/models/capture-client.md) | models.CaptureSiteContainer.FromCaptureClient(models.CaptureClient captureClient) |
| [`models.CaptureGateway`](../../../doc/models/capture-gateway.md) | models.CaptureSiteContainer.FromCaptureGateway(models.CaptureGateway captureGateway) |
| [`models.CaptureNewAssoc`](../../../doc/models/capture-new-assoc.md) | models.CaptureSiteContainer.FromCaptureNewAssoc(models.CaptureNewAssoc captureNewAssoc) |
| [`models.CaptureRadiotap`](../../../doc/models/capture-radiotap.md) | models.CaptureSiteContainer.FromCaptureRadiotap(models.CaptureRadiotap captureRadiotap) |
| [`models.CaptureRadiotapwired`](../../../doc/models/capture-radiotapwired.md) | models.CaptureSiteContainer.FromCaptureRadiotapwired(models.CaptureRadiotapwired captureRadiotapwired) |
| [`models.CaptureScan`](../../../doc/models/capture-scan.md) | models.CaptureSiteContainer.FromCaptureScan(models.CaptureScan captureScan) |
| [`models.CaptureSwitch`](../../../doc/models/capture-switch.md) | models.CaptureSiteContainer.FromCaptureSwitch(models.CaptureSwitch captureSwitch) |
| [`models.CaptureWired`](../../../doc/models/capture-wired.md) | models.CaptureSiteContainer.FromCaptureWired(models.CaptureWired captureWired) |
| [`models.CaptureWireless`](../../../doc/models/capture-wireless.md) | models.CaptureSiteContainer.FromCaptureWireless(models.CaptureWireless captureWireless) |

## models.CaptureClient

### Initialization Code

#### Example

```go
value := models.CaptureSiteContainer.FromCaptureClient(models.CaptureClient{
    ClientMac:            models.NewOptional(models.ToPointer("60a10a773412")),
    Duration:             models.NewOptional(models.ToPointer(300)),
    IncludesMcast:        models.ToPointer(false),
    MaxPktLen:            models.NewOptional(models.ToPointer(128)),
    NumPackets:           models.NewOptional(models.ToPointer(1000)),
    Type:                 "client",
})
```

## models.CaptureGateway

### Initialization Code

#### Example

```go
value := models.CaptureSiteContainer.FromCaptureGateway(models.CaptureGateway{
    Duration:             models.NewOptional(models.ToPointer(300)),
    Format:               models.ToPointer(models.CaptureGatewayFormatEnum_STREAM),
    Gateways:             map[string]models.CaptureGatewayGateways{
        "key0": models.CaptureGatewayGateways{
        },
    },
    MaxPktLen:            models.NewOptional(models.ToPointer(128)),
    NumPackets:           models.NewOptional(models.ToPointer(1000)),
    Type:                 "gateway",
})
```

## models.CaptureNewAssoc

### Initialization Code

#### Example

```go
value := models.CaptureSiteContainer.FromCaptureNewAssoc(models.CaptureNewAssoc{
    ApMac:                models.ToPointer("a83a79a947ee"),
    ClientMac:            models.ToPointer("60a10a773412"),
    Duration:             models.NewOptional(models.ToPointer(300)),
    IncludesMcast:        models.ToPointer(false),
    MaxPktLen:            models.NewOptional(models.ToPointer(128)),
    NumPackets:           models.NewOptional(models.ToPointer(1000)),
    Type:                 "new_assoc",
})
```

## models.CaptureRadiotap

### Initialization Code

#### Example

```go
value := models.CaptureSiteContainer.FromCaptureRadiotap(models.CaptureRadiotap{
    ApMac:                models.ToPointer("a83a79a947ee"),
    Band:                 models.ToPointer(models.CaptureRadiotapBandEnum_ENUM24),
    ClientMac:            models.ToPointer("38f9d3972ff1"),
    Duration:             models.NewOptional(models.ToPointer(300)),
    Format:               models.ToPointer(models.CaptureRadiotapFormatEnum_STREAM),
    MaxPktLen:            models.NewOptional(models.ToPointer(128)),
    NumPackets:           models.NewOptional(models.ToPointer(1000)),
    Ssid:                 models.ToPointer("test"),
    TcpdumpExpression:    models.NewOptional(models.ToPointer("tcp port 80")),
    Type:                 "radiotap",
    WlanId:               models.ToPointer(uuid.MustParse("fac8e973-feb9-421a-b381-aabbc4b61f5a")),
})
```

## models.CaptureRadiotapwired

### Initialization Code

#### Example

```go
value := models.CaptureSiteContainer.FromCaptureRadiotapwired(models.CaptureRadiotapwired{
    Band:                      models.ToPointer(models.CaptureRadiotapwiredBandEnum_ENUM24),
    ClientMac:                 models.NewOptional(models.ToPointer("38f9d3972ff1")),
    Duration:                  models.NewOptional(models.ToPointer(300)),
    Format:                    models.ToPointer(models.CaptureRadiotapwiredFormatEnum_STREAM),
    MaxPktLen:                 models.NewOptional(models.ToPointer(128)),
    NumPackets:                models.NewOptional(models.ToPointer(1000)),
    RadiotapTcpdumpExpression: models.ToPointer("type"),
    Ssid:                      models.NewOptional(models.ToPointer("test")),
    TcpdumpExpression:         models.NewOptional(models.ToPointer("tcp port 80")),
    Type:                      "radiotap,wired",
    WiredTcpdumpExpression:    models.NewOptional(models.ToPointer("tcp port 80")),
    WlanId:                    models.NewOptional(models.ToPointer("fac8e973-feb9-421a-b381-aabbc4b61f5a")),
})
```

## models.CaptureScan

### Initialization Code

#### Example

```go
value := models.CaptureSiteContainer.FromCaptureScan(models.CaptureScan{
    Band:                 models.NewOptional(models.ToPointer(models.CaptureScanBandEnum_ENUM24)),
    Bandwidth:            models.ToPointer(models.Dot11BandwidthEnum_ENUM20),
    Channel:              models.ToPointer(1),
    ClientMac:            models.NewOptional(models.ToPointer("38f9d3972ff1")),
    Duration:             models.NewOptional(models.ToPointer(300)),
    Format:               models.ToPointer(models.CaptureScanFormatEnum_STREAM),
    MaxPktLen:            models.NewOptional(models.ToPointer(128)),
    NumPackets:           models.NewOptional(models.ToPointer(1000)),
    TcpdumpExpression:    models.ToPointer("tcp port 80"),
    Type:                 "scan",
})
```

## models.CaptureSwitch

### Initialization Code

#### Example

```go
value := models.CaptureSiteContainer.FromCaptureSwitch(models.CaptureSwitch{
    Duration:             models.NewOptional(models.ToPointer(300)),
    Format:               models.ToPointer(models.CaptureSwitchFormatEnum_STREAM),
    MaxPktLen:            models.NewOptional(models.ToPointer(128)),
    NumPackets:           models.NewOptional(models.ToPointer(1000)),
    Switches:             map[string]models.CaptureSwitchSwitches{
        "key0": models.CaptureSwitchSwitches{
        },
        "key1": models.CaptureSwitchSwitches{
        },
        "key2": models.CaptureSwitchSwitches{
        },
    },
    TcpdumpExpression:    models.ToPointer("port 443"),
    Type:                 "switch",
})
```

## models.CaptureWired

### Initialization Code

#### Example

```go
value := models.CaptureSiteContainer.FromCaptureWired(models.CaptureWired{
    Duration:             models.NewOptional(models.ToPointer(300)),
    Format:               models.ToPointer(models.CaptureWiredFormatEnum_PCAP),
    MaxPktLen:            models.NewOptional(models.ToPointer(128)),
    NumPackets:           models.NewOptional(models.ToPointer(1000)),
    TcpdumpExpression:    models.NewOptional(models.ToPointer("tcp port 80")),
    Type:                 "wired",
})
```

## models.CaptureWireless

### Initialization Code

#### Example

```go
value := models.CaptureSiteContainer.FromCaptureWireless(models.CaptureWireless{
    Band:                 models.ToPointer(models.CaptureWirelessBandEnum_ENUM24),
    Duration:             models.NewOptional(models.ToPointer(300)),
    Format:               models.ToPointer(models.CaptureWirelessFormatEnum_PCAP),
    MaxPktLen:            models.NewOptional(models.ToPointer(128)),
    NumPackets:           models.NewOptional(models.ToPointer(1000)),
    Type:                 "wireless",
})
```

