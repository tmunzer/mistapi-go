
# Mxcluster Radsec Tls

TLS settings for RadSec on a Mist Edge cluster

## Structure

`MxclusterRadsecTls`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Keypair` | `*string` | Optional | Name or identifier of the TLS keypair used by RadSec |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    mxclusterRadsecTls := models.MxclusterRadsecTls{
        Keypair:              models.ToPointer("keypair8"),
    }

}
```

