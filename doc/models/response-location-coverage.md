
# Response Location Coverage

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseLocationCoverage`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BeamsMeans` | `[][]float64` | Required | List of [x, y, mean]s, x/y are in meters (UI would need to use map.ppm to calulate the pixel location from top-left). |
| `End` | `int` | Required | - |
| `Gridsize` | `float64` | Required | Size of grid, in meter |
| `ResultDef` | `[]string` | Required | List of names annotating the fields in results |
| `Results` | `[][]float64` | Required | List of results, see result_def. |
| `Start` | `int` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

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
  "start": 106,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

