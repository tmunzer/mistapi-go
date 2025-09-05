
# Module Stat Item Network Resource

*This model accepts additional fields of type interface{}.*

## Structure

`ModuleStatItemNetworkResource`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Count` | `*int` | Optional | current usage of the network resource<br><br>**Constraints**: `>= 0` |
| `Limit` | `*int` | Optional | maximum usage of the network resource<br><br>**Constraints**: `>= 0` |
| `Type` | `*string` | Optional | type of the network resource (e.g. FIB, FLOW, ...) |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "count": 17,
  "limit": 768000,
  "type": "FIB",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

