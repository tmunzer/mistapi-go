
# Map Micello

*This model accepts additional fields of type interface{}.*

## Structure

`MapMicello`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AccountKey` | `string` | Required | Account key that has access to the map |
| `DefaultLevelId` | `int` | Required | Micello floor/level id |
| `MapId` | `uuid.UUID` | Required | Micello map id |
| `VendorName` | `string` | Required, Constant | The vendor ‘micello’. enum: `micello`<br><br>**Value**: `"micello"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "account_key": null,
  "default_level_id": 5,
  "map_id": "6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9",
  "vendor_name": "micello",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

