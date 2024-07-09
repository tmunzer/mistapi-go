
# Sle Impact Summary Ap Item

## Structure

`SleImpactSummaryApItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApMac` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `Degraded` | `float64` | Required | - |
| `Duration` | `float64` | Required | - |
| `Name` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `Total` | `float64` | Required | - |

## Example (as JSON)

```json
{
  "ap_mac": "ap_mac4",
  "degraded": 71.52,
  "duration": 200.58,
  "name": "name2",
  "total": 156.48
}
```

