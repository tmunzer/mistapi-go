
# Gateway Wan Probe Override

if `usage`==`wan`

## Structure

`GatewayWanProbeOverride`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ips` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `ProbeProfile` | [`*models.GatewayWanProbeOverrideProbeProfileEnum`](../../doc/models/gateway-wan-probe-override-probe-profile-enum.md) | Optional | enum: `broadband`, `lte`<br>**Default**: `"broadband"` |

## Example (as JSON)

```json
{
  "probe_profile": "broadband",
  "ips": [
    "ips0",
    "ips1"
  ]
}
```

