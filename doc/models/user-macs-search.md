
# User Macs Search

## Structure

`UserMacsSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Limit` | `*int` | Optional | - |
| `Page` | `*int` | Optional | - |
| `Results` | [`[]models.UserMac`](../../doc/models/user-mac.md) | Optional | - |
| `Total` | `*int` | Optional | - |

## Example (as JSON)

```json
{
  "limit": 100,
  "page": 1,
  "total": 1,
  "results": [
    {
      "id": "000023ba-0000-0000-0000-000000000000",
      "labels": [
        "labels2",
        "labels1"
      ],
      "mac": "mac0",
      "name": "name6",
      "notes": "notes6",
      "radius_group": "radius_group8"
    },
    {
      "id": "000023ba-0000-0000-0000-000000000000",
      "labels": [
        "labels2",
        "labels1"
      ],
      "mac": "mac0",
      "name": "name6",
      "notes": "notes6",
      "radius_group": "radius_group8"
    },
    {
      "id": "000023ba-0000-0000-0000-000000000000",
      "labels": [
        "labels2",
        "labels1"
      ],
      "mac": "mac0",
      "name": "name6",
      "notes": "notes6",
      "radius_group": "radius_group8"
    }
  ]
}
```

