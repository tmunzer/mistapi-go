
# Proxy

Proxy Configuration to talk to Mist

## Structure

`Proxy`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Disabled` | `*bool` | Optional | Whether this proxy configuration is disabled<br><br>**Default**: `false` |
| `Url` | `*string` | Optional | Proxy URL used to reach Mist |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    proxy := models.Proxy{
        Disabled:             models.ToPointer(true),
        Url:                  models.ToPointer("https://proxy.corp.com:8080/"),
    }

}
```

