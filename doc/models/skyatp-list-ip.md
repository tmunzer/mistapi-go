
# Skyatp List Ip

IP address entry in a Sky ATP SecIntel list

## Structure

`SkyatpListIp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Comment` | `*string` | Optional | Optional note describing the IP address list entry |
| `Value` | `string` | Required | IP address included in the Sky ATP SecIntel list |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    skyatpListIp := models.SkyatpListIp{
        Comment:              models.ToPointer("nas"),
        Value:                "10.1.3.5",
    }

}
```

