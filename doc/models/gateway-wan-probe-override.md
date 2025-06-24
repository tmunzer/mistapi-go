
# Gateway Wan Probe Override

Only if `usage`==`wan`

*This model accepts additional fields of type interface{}.*

## Structure

`GatewayWanProbeOverride`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ips` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `ProbeProfile` | [`*models.GatewayWanProbeOverrideProbeProfileEnum`](../../doc/models/gateway-wan-probe-override-probe-profile-enum.md) | Optional | enum: `broadband`, `lte`<br><br>**Default**: `"broadband"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "probe_profile": "broadband",
  "ips": [
    "ips0",
    "ips1"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

