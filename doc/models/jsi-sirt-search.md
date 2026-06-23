
# Jsi Sirt Search

Juniper Security Intelligence SIRT search response with result metadata

## Structure

`JsiSirtSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | Upper bound timestamp for the SIRT search window |
| `Limit` | `*int` | Optional | Number of results to return |
| `Next` | `*string` | Optional | Pagination URL for the next page of SIRT advisories |
| `Results` | [`[]models.JsiSirtItem`](../../doc/models/jsi-sirt-item.md) | Optional | List of SIRT advisories |
| `Start` | `*int` | Optional | Lower bound timestamp for the SIRT search window |
| `Total` | `*int` | Optional | Count of SIRT advisories matching the search |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    jsiSirtSearch := models.JsiSirtSearch{
        End:                  models.ToPointer(58),
        Limit:                models.ToPointer(112),
        Next:                 models.ToPointer("next2"),
        Results:              []models.JsiSirtItem{
            models.JsiSirtItem{
                CvssScore:            models.ToPointer(float64(231.84)),
                Id:                   models.ToPointer("id6"),
                Models:               []string{
                    "models2",
                    "models3",
                },
                Problem:              models.ToPointer("problem6"),
                PublishedDate:        models.ToPointer(160),
            },
            models.JsiSirtItem{
                CvssScore:            models.ToPointer(float64(231.84)),
                Id:                   models.ToPointer("id6"),
                Models:               []string{
                    "models2",
                    "models3",
                },
                Problem:              models.ToPointer("problem6"),
                PublishedDate:        models.ToPointer(160),
            },
        },
        Start:                models.ToPointer(16),
    }

}
```

