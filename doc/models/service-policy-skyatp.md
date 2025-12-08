
# Service Policy Skyatp

SRX only

## Structure

`ServicePolicySkyatp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DnsDgaDetection` | [`*models.ServicePolicySkyatpDnsDgaDetectionEnum`](../../doc/models/service-policy-skyatp-dns-dga-detection-enum.md) | Optional | enum: `disabled`, `default`, `standard`, `strict`<br><br>**Default**: `"disabled"` |
| `DnsTunnelDetection` | [`*models.ServicePolicySkyatpDnsTunnelDetectionEnum`](../../doc/models/service-policy-skyatp-dns-tunnel-detection-enum.md) | Optional | enum: `disabled`, `default`, `standard`, `strict`<br><br>**Default**: `"disabled"` |
| `HttpInspection` | [`*models.ServicePolicySkyatpHttpInspectionEnum`](../../doc/models/service-policy-skyatp-http-inspection-enum.md) | Optional | enum: `disabled`, `standard`<br><br>**Default**: `"disabled"` |
| `IotDevicePolicy` | [`*models.ServicePolicySkyatpIotDevicePolicyEnum`](../../doc/models/service-policy-skyatp-iot-device-policy-enum.md) | Optional | enum: `disabled`, `enabled`<br><br>**Default**: `"disabled"` |

## Example (as JSON)

```json
{
  "dns_dga_detection": "disabled",
  "dns_tunnel_detection": "disabled",
  "http_inspection": "disabled",
  "iot_device_policy": "disabled"
}
```

