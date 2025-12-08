
# Switch Stp Config

## Structure

`SwitchStpConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BridgePriority` | `*string` | Optional | Switch STP priority. Range [0, 4k, 8k.. 60k] in steps of 4k. Bridge priority applies to both VSTP and RSTP.<br><br>**Default**: `"32k"` |

## Example (as JSON)

```json
{
  "bridge_priority": "40k"
}
```

