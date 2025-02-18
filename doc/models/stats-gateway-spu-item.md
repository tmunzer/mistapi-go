
# Stats Gateway Spu Item

*This model accepts additional fields of type interface{}.*

## Structure

`StatsGatewaySpuItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SpuCpu` | `*int` | Optional | - |
| `SpuCurrentSession` | `*int` | Optional | - |
| `SpuMaxSession` | `*int` | Optional | - |
| `SpuMemory` | `*int` | Optional | - |
| `SpuPendingSession` | `*int` | Optional | - |
| `SpuValidSession` | `*int` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "spu_cpu": 3670632,
  "spu_current_session": 215,
  "spu_max_session": 131072,
  "spu_memory": 46,
  "spu_pending_session": 0,
  "spu_valid_session": 0,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

