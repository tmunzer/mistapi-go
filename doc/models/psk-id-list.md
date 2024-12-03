
# Psk Id List

*This model accepts additional fields of type interface{}.*

## Structure

`PskIdList`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PskIds` | `[]uuid.UUID` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "psk_ids": [
    "0039c16c-ca87-4d3f-bb94-b97c58199f18",
    "6562cc8e-5893-418a-acaa-4d7c1af8084f"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

