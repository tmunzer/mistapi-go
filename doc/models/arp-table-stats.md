
# Arp Table Stats

*This model accepts additional fields of type interface{}.*

## Structure

`ArpTableStats`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ArpTableCount` | `*int` | Optional | - |
| `MaxEntriesSupported` | `*int` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "arp_table_count": 136,
  "max_entries_supported": 8,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

