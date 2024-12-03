
# Route Summary Stats

*This model accepts additional fields of type interface{}.*

## Structure

`RouteSummaryStats`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `FibRoutes` | `*int` | Optional | - |
| `MaxUnicastRoutesSupported` | `*int` | Optional | - |
| `RibRoutes` | `*int` | Optional | - |
| `TotalRoutes` | `*int` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "fib_routes": 100,
  "max_unicast_routes_supported": 94,
  "rib_routes": 218,
  "total_routes": 170,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

