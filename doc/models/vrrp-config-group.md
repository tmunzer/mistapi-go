
# Vrrp Config Group

## Structure

`VrrpConfigGroup`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Preempt` | `*bool` | Optional | If `true`, allow preemption (a backup router can preempt a primary router)<br><br>**Default**: `false` |
| `Priority` | `*int` | Optional | - |

## Example (as JSON)

```json
{
  "preempt": false,
  "priority": 24
}
```

