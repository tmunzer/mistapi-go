
# Const State

State or territory definition returned by the constants API

## Structure

`ConstState`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `IsoCode` | `*string` | Optional | ISO subdivision code for the state or territory |
| `Name` | `*string` | Optional | State or territory display name |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    constState := models.ConstState{
        IsoCode:              models.ToPointer("AK"),
        Name:                 models.ToPointer("Alaska"),
    }

}
```

