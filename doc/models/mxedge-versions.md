
# Mxedge Versions

Read-only Mist Edge service versions

## Structure

`MxedgeVersions`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mxagent` | `*string` | Optional, Read-only | Reported version of the mxagent service |
| `Tunterm` | `*string` | Optional, Read-only | Reported version of the tunnel termination service |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    mxedgeVersions := models.MxedgeVersions{
        Mxagent:              models.ToPointer("mxagent0"),
        Tunterm:              models.ToPointer("tunterm8"),
    }

}
```

