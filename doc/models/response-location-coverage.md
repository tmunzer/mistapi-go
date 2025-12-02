
# Response Location Coverage

## Structure

`ResponseLocationCoverage`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BeamsMeans` | `[][]float64` | Required | List of [x, y, mean]s, x/y are in meters (UI would need to use map.ppm to calculate the pixel location from top-left). |
| `End` | `int` | Required | - |
| `Gridsize` | `float64` | Required | Size of grid, in meter |
| `ResultDef` | `[]string` | Required | List of names annotating the fields in results |
| `Results` | `[][]float64` | Required | List of results, see result_def. |
| `Start` | `int` | Required | - |

## Example (as JSON)

```json
{
  "beams_means": [
    [
      153.16,
      153.17
    ],
    [
      153.16,
      153.17
    ]
  ],
  "end": 148,
  "gridsize": 202.46,
  "result_def": [
    "result_def4"
  ],
  "results": [
    [
      31.85,
      31.84
    ],
    [
      31.85,
      31.84
    ]
  ],
  "start": 106
}
```

