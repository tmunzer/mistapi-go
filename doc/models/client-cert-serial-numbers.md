
# Client Cert Serial Numbers

Client certificate serial numbers targeted by a certificate operation

*This model accepts additional fields of type interface{}.*

## Structure

`ClientCertSerialNumbers`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SerialNumbers` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    clientCertSerialNumbers := models.ClientCertSerialNumbers{
        SerialNumbers:        []string{
            "13 00 13 03 23 EE D5 84 01",
        },
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

