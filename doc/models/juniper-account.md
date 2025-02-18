
# Juniper Account

*This model accepts additional fields of type interface{}.*

## Structure

`JuniperAccount`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `LinkedBy` | `*string` | Optional | - |
| `Name` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "linked_by": "John Smith (john@abccorp.com)",
  "name": "ABC Corp",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

