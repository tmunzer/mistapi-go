
# Utils Zigbee Join

## Structure

`UtilsZigbeeJoin`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Duration` | `*int` | Optional | Duration in seconds for which new Zigbee end device joins are permitted. Range is 30–3600<br><br>**Default**: `600`<br><br>**Constraints**: `>= 30`, `<= 3600` |

## Example (as JSON)

```json
{
  "duration": 600
}
```

