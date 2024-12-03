
# Stats Wireless Client Wxrule Usage

*This model accepts additional fields of type interface{}.*

## Structure

`StatsWirelessClientWxruleUsage`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `TagId` | `*uuid.UUID` | Optional | - |
| `Usage` | `*int` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "tag_id": "0000094c-0000-0000-0000-000000000000",
  "usage": 238,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

