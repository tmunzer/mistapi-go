
# Evpn Config

EVPN Junos settings

## Structure

`EvpnConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | - |
| `Role` | [`*models.EvpnConfigRoleEnum`](../../doc/models/evpn-config-role-enum.md) | Optional | enum: `access`, `collapsed-core`, `core`, `distribution`, `esilag-access`, `none` |

## Example (as JSON)

```json
{
  "enabled": false,
  "role": "esilag-access"
}
```

