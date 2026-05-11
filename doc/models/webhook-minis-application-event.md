
# Webhook Minis Application Event

## Structure

`WebhookMinisApplicationEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DeviceMac` | `*string` | Optional | MAC address of the device |
| `Ip` | `*string` | Optional | IP address test was performed to |
| `Latency` | `*int` | Optional | latency in milliseconds |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `ProbeName` | `*string` | Optional | Name of the probe |
| `ProbeType` | `*string` | Optional | Type of probe |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `SrcIp` | `*string` | Optional | Source IP address of the test |
| `Success` | `*bool` | Optional | Whether the test was successful |
| `TestType` | [`*models.SynthetictestConfigCustomProbeTypeEnum`](../../doc/models/synthetictest-config-custom-probe-type-enum.md) | Optional | enum: `application`, `curl`, `icmp`, `reachability`, `tcp`<br><br>**Default**: `"icmp"` |
| `Timestamp` | `*float64` | Optional | Epoch (seconds) |
| `Vlan` | `*int` | Optional | VLAN ID used for the test |

## Example (as JSON)

```json
{
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "probe_name": "connectivitycheck.gstatic.com",
  "probe_type": "application",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "test_type": "icmp",
  "vlan": 12,
  "device_mac": "device_mac2",
  "ip": "ip2",
  "latency": 36
}
```

