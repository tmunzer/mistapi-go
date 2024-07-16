
# Vpn

## Structure

`Vpn`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | - |
| `ModifiedTime` | `*float64` | Optional | - |
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
  "created_time": 55.36,
  "id": "0000213a-0000-0000-0000-000000000000",
  "modified_time": 23.6
}
```

