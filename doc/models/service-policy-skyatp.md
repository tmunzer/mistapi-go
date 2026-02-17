
# Service Policy Skyatp

SRX only

## Structure

`ServicePolicySkyatp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DnsDgaDetection` | [`*models.ServicePolicySkyatpDnsDgaDetection`](../../doc/models/service-policy-skyatp-dns-dga-detection.md) | Optional | - |
| `DnsTunnelDetection` | [`*models.ServicePolicySkyatpDnsTunnelDetection`](../../doc/models/service-policy-skyatp-dns-tunnel-detection.md) | Optional | - |
| `HttpInspection` | [`*models.ServicePolicySkyatpHttpInspection`](../../doc/models/service-policy-skyatp-http-inspection.md) | Optional | - |
| `IotDevicePolicy` | [`*models.ServicePolicySkyatpIotDevicePolicy`](../../doc/models/service-policy-skyatp-iot-device-policy.md) | Optional | - |

## Example (as JSON)

```json
{
  "dns_dga_detection": {
    "enabled": false,
    "profile": "standard"
  },
  "dns_tunnel_detection": {
    "enabled": false,
    "profile": "default"
  },
  "http_inspection": {
    "enabled": false,
    "profile": "standard"
  },
  "iot_device_policy": {
    "enabled": false
  }
}
```

