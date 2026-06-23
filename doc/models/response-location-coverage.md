
# Response Location Coverage

Beam coverage overview for a site location map

## Structure

`ResponseLocationCoverage`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BeamsMeans` | `[][]float64` | Required | List of [x, y, mean]s, x/y are in meters (UI would need to use map.ppm to calculate the pixel location from top-left). |
| `End` | `int` | Required | Epoch timestamp for the end of the coverage analysis window |
| `Gridsize` | `float64` | Required | Grid cell size in meters for coverage samples |
| `ResultDef` | `[]string` | Required | List of names annotating the fields in results |
| `Results` | `[][]float64` | Required | List of results, see result_def. |
| `Start` | `int` | Required | Epoch timestamp for the start of the coverage analysis window |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseLocationCoverage := models.ResponseLocationCoverage{
        BeamsMeans:           [][]float64{
            []float64{
                float64(97.46),
            },
        },
        End:                  210,
        Gridsize:             float64(146.76),
        ResultDef:            []string{
            "result_def4",
            "result_def5",
            "result_def6",
        },
        Results:              [][]float64{
            []float64{
                float64(168.45),
            },
        },
        Start:                168,
    }

}
```

