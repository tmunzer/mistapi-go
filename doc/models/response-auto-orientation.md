
# Response Auto Orientation

## Structure

`ResponseAutoOrientation`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `State` | [`*models.AutoOrientationStateEnum`](../../doc/models/auto-orientation-state-enum.md) | Optional | The state of auto orient for a given map derived from an Enum. enum: `Enqueued`, `Not Started`, `Oriented` |
| `TimeQueued` | `*float64` | Optional | Time when auto orient process was last queued for this map |

## Example (as JSON)

```json
{
  "state": "Oriented",
  "time_queued": 82.0
}
```

