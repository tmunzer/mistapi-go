
# Gateway Mgmt

Gateway settings

## Structure

`GatewayMgmt`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ConfigRevertTimer` | `*int` | Optional | Rollback timer for commit confirmed<br><br>**Default**: `10`<br><br>**Constraints**: `>= 1`, `<= 30` |

## Example (as JSON)

```json
{
  "config_revert_timer": 10
}
```

