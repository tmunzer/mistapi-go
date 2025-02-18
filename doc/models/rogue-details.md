
# Rogue Details

*This model accepts additional fields of type interface{}.*

## Structure

`RogueDetails`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Manufacture` | `string` | Required | - |
| `SeenAsClient` | `bool` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "manufacture": "manufacture8",
  "seen_as_client": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

