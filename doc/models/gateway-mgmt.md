
# Gateway Mgmt

Gateway settings

*This model accepts additional fields of type interface{}.*

## Structure

`GatewayMgmt`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ConfigRevertTimer` | `*int` | Optional | Rollback timer for commit confirmed<br><br>**Default**: `10`<br><br>**Constraints**: `>= 1`, `<= 30` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "config_revert_timer": 10,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

