
# Evpn Options Vs Instance

*This model accepts additional fields of type interface{}.*

## Structure

`EvpnOptionsVsInstance`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Networks` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "networks": [
    "networks0",
    "networks1"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

