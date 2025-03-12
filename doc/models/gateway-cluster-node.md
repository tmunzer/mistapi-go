
# Gateway Cluster Node

*This model accepts additional fields of type interface{}.*

## Structure

`GatewayClusterNode`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mac` | `string` | Required | Gateway MAC Address. Format is `[0-9a-f]{12}` (e.g. "5684dae9ac8b") |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "mac": "mac6",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

