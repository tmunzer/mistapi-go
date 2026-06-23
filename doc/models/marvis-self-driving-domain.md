
# Marvis Self Driving Domain

Self-driving automation settings for one Marvis domain

## Structure

`MarvisSelfDrivingDomain`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether self-driving automation is enabled for this domain<br><br>**Default**: `false` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    marvisSelfDrivingDomain := models.MarvisSelfDrivingDomain{
        Enabled:              models.ToPointer(false),
    }

}
```

