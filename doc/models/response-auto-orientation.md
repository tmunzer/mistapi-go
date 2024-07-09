
# Response Auto Orientation

## Structure

`ResponseAutoOrientation`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `State` | [`*models.AutoOrientationStateEnum`](../../doc/models/auto-orientation-state-enum.md) | Optional | The state of auto orient for a given map derived from an Enum |
| `TimeQueued` | `*float64` | Optional | Time when auto orient process was last queued for this map |

## Example (as JSON)

```json
{
  "state": "Not Started",
  "time_queued": 140.6
}
```

