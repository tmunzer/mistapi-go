
# Ssr Proxy

SSR proxy configuration to talk to Mist

## Structure

`SsrProxy`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Disabled` | `*bool` | Optional | Whether the SSR proxy configuration is disabled<br><br>**Default**: `false` |
| `Url` | `*string` | Optional | Proxy URL that SSR devices use to reach Mist |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    ssrProxy := models.SsrProxy{
        Disabled:             models.ToPointer(true),
        Url:                  models.ToPointer("https://proxy.corp.com:8080/"),
    }

}
```

