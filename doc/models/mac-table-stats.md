
# Mac Table Stats

*This model accepts additional fields of type interface{}.*

## Structure

`MacTableStats`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MacTableCount` | `*int` | Optional | - |
| `MaxMacEntriesSupported` | `*int` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "mac_table_count": 26,
  "max_mac_entries_supported": 170,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

