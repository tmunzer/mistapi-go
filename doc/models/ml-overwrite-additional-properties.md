
# Ml Overwrite Additional Properties

Location machine learning parameter overwrite values for one client model

## Structure

`MlOverwriteAdditionalProperties`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Int` | `*int` | Optional | Location ML intercept value to use for the model |
| `Ple` | `*int` | Optional | Path-loss estimate value to use for the location machine learning model |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    mlOverwriteAdditionalProperties := models.MlOverwriteAdditionalProperties{
        Int:                  models.ToPointer(214),
        Ple:                  models.ToPointer(210),
    }

}
```

