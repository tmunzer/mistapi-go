
# Gateway Wan Probe Override

Only if `usage`==`wan`

## Structure

`GatewayWanProbeOverride`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ip6s` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `Ips` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `ProbeProfile` | [`*models.GatewayWanProbeOverrideProbeProfileEnum`](../../doc/models/gateway-wan-probe-override-probe-profile-enum.md) | Optional | enum: `broadband`, `lte`<br><br>**Default**: `"broadband"` |

## Example (as JSON)

```json
{
  "probe_profile": "broadband",
  "ip6s": [
    "ip6s8",
    "ip6s9"
  ],
  "ips": [
    "ips0",
    "ips1"
  ]
}
```

