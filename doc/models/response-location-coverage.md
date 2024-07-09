
# Response Location Coverage

## Structure

`ResponseLocationCoverage`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BeamsMeans` | `[][]float64` | Required | list of [x, y, mean]s, x/y are in meters (UI would need to use map.ppm to calulate the pixel location from top-left). |
| `End` | `int` | Required | - |
| `Gridsize` | `float64` | Required | the size of grid, in meter |
| `ResultDef` | `[]string` | Required | list of names annotating the fields in results |
| `Results` | `[][]float64` | Required | list of results, see result_def. |
| `Start` | `int` | Required | - |

## Example (as JSON)

```json
{
  "beams_means": [
    [
      251.76,
      251.77,
      251.78
    ],
    [
      251.76,
      251.77,
      251.78
    ],
    [
      251.76,
      251.77,
      251.78
    ]
  ],
  "end": 24,
  "gridsize": 45.06,
  "result_def": [
    "result_def4",
    "result_def5"
  ],
  "results": [
    [
      57.79,
      57.8
    ],
    [
      57.79,
      57.8
    ]
  ],
  "start": 238
}
```

