
# Version String

JSI device upgrade target version request

*This model accepts additional fields of type interface{}.*

## Structure

`VersionString`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Version` | `*string` | Optional | Target software version for the JSI device upgrade |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    versionString := models.VersionString{
        Version:              models.ToPointer("version8"),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

