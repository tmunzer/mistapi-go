
# Bgp Config Community

## Structure

`BgpConfigCommunity`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organnization |
| `LocalPreference` | `*int` | Optional | - |
| `VpnName` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "local_preference": 14,
  "vpn_name": "vpn_name0"
}
```

