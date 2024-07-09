
# Vpn

## Structure

`Vpn`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*int` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | - |
| `ModifiedTime` | `*int` | Optional | - |
| `Name` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Paths` | [`map[string]models.VpnPath`](../../doc/models/vpn-path.md) | Required | - |

## Example (as JSON)

```json
{
  "name": "name6",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "paths": {
    "key0": {
      "bfd_profile": "broadband",
      "ip": "ip8",
      "pod": 146
    },
    "key1": {
      "bfd_profile": "broadband",
      "ip": "ip8",
      "pod": 146
    }
  },
  "created_time": 160,
  "id": "0000213a-0000-0000-0000-000000000000",
  "modified_time": 56
}
```

