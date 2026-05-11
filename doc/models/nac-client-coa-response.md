
# Nac Client Coa Response

## Structure

`NacClientCoaResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DeviceMac` | `*string` | Optional | MAC address of the target device (AP or switch MAC) |
| `DeviceType` | [`*models.DeviceTypeEnum`](../../doc/models/device-type-enum.md) | Optional | enum: `ap`, `gateway`, `switch` |

## Example (as JSON)

```json
{
  "device_mac": "device_mac2",
  "device_type": "gateway"
}
```

