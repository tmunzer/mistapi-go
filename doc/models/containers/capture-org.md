
# Capture Org

Packet capture request payload for organization-level captures

## Class Name

`CaptureOrg`

## Cases

| Type | Factory Method |
|  --- | --- |
| [`models.CaptureMxedge`](../../../doc/models/capture-mxedge.md) | models.CaptureOrgContainer.FromCaptureMxedge(models.CaptureMxedge captureMxedge) |

## models.CaptureMxedge

### Initialization Code

#### Example

```go
value := models.CaptureOrgContainer.FromCaptureMxedge(models.CaptureMxedge{
    Duration:             models.ToPointer(600),
    Format:               models.ToPointer(models.CaptureMxedgeFormatEnum_STREAM),
    MaxPktLen:            models.ToPointer(512),
    NumPackets:           models.ToPointer(100),
    Type:                 "mxedge",
    TzspHost:             models.ToPointer("192.168.1.2"),
    TzspPort:             models.ToPointer(37008),
})
```

