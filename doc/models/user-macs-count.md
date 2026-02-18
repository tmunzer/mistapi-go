
# User Macs Count

User MACs count response

## Structure

`UserMacsCount`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | End timestamp |
| `Limit` | `*int` | Optional | Number of results to return |
| `Results` | [`[]models.UserMac`](../../doc/models/user-mac.md) | Optional | List of user MAC entries |
| `Start` | `*int` | Optional | Start timestamp |
| `Total` | `*int` | Optional | Total number of results |

## Example (as JSON)

```json
{
  "end": 200,
  "limit": 226,
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
      "radius_group": "radius_group8",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "start": 158,
  "total": 64
}
```

