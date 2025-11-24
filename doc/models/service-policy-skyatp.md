
# Service Policy Skyatp

SRX only

*This model accepts additional fields of type interface{}.*

## Structure

`ServicePolicySkyatp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DnsDgaDetection` | [`*models.DnsDgaDetectionEnum`](../../doc/models/dns-dga-detection-enum.md) | Optional | enum: `disabled`, `default`, `standard`, `strict`<br><br>**Default**: `"disabled"` |
| `DnsTunnelDetection` | [`*models.DnsTunnelDetectionEnum`](../../doc/models/dns-tunnel-detection-enum.md) | Optional | enum: `disabled`, `default`, `standard`, `strict`<br><br>**Default**: `"disabled"` |
| `HttpInspection` | [`*models.HttpInspectionEnum`](../../doc/models/http-inspection-enum.md) | Optional | enum: `disabled`, `standard`<br><br>**Default**: `"disabled"` |
| `IotDevicePolicy` | [`*models.IotDevicePolicyEnum`](../../doc/models/iot-device-policy-enum.md) | Optional | enum: `disabled`, `enabled`<br><br>**Default**: `"disabled"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "dns_dga_detection": "disabled",
  "dns_tunnel_detection": "disabled",
  "http_inspection": "disabled",
  "iot_device_policy": "disabled",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

