
# Evpn Config

EVPN Junos settings

*This model accepts additional fields of type interface{}.*

## Structure

`EvpnConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | - |
| `Role` | [`*models.EvpnConfigRoleEnum`](../../doc/models/evpn-config-role-enum.md) | Optional | enum: `access`, `collapsed-core`, `core`, `distribution`, `esilag-access`, `none` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "enabled": false,
  "role": "esilag-access",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

