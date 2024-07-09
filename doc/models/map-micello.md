
# Map Micello

## Structure

`MapMicello`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AccountKey` | `string` | Required | the account key that has access to the map |
| `DefaultLevelId` | `int` | Required | micello floor/level id |
| `MapId` | `uuid.UUID` | Required | micello map id |
| `VendorName` | `string` | Required, Constant | the vendor ‘micello’<br>**Default**: `"micello"` |

## Example (as JSON)

```json
{
  "account_key": "account_key6",
  "default_level_id": 5,
  "map_id": "6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9",
  "vendor_name": "micello"
}
```

