
# Webhook Minis Reachability Event

## Structure

`WebhookMinisReachabilityEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AvgLatency` | `*float64` | Optional | Average latency in milliseconds |
| `DeviceMac` | `*string` | Optional | MAC address of the device performing the test |
| `LossPercentage` | `*float64` | Optional | Packet loss percentage |
| `MaxLatency` | `*float64` | Optional | Maximum latency in milliseconds |
| `MinLatency` | `*float64` | Optional | Minimum latency in milliseconds |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `ProbeName` | `*string` | Optional | Name of the probe |
| `ProbeTarget` | `*string` | Optional | Target host or IP for the probe |
| `ProbeType` | `*string` | Optional | Type of probe |
| `Protocol` | `*string` | Optional | Protocol used for the test |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `TestType` | `*string` | Optional | Type of test performed |
| `Timestamp` | `*float64` | Optional | Epoch (seconds) |
| `Vlan` | `*int` | Optional | VLAN ID used for the test |

## Example (as JSON)

```json
{
  "device_mac": "7cb68d8f0440",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "probe_name": "google ping",
  "probe_target": "google.com",
  "probe_type": "reachability",
  "protocol": "icmp",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "test_type": "ping",
  "vlan": 12,
  "avg_latency": 229.9,
  "loss_percentage": 135.74,
  "max_latency": 93.62,
  "min_latency": 4.72
}
```

