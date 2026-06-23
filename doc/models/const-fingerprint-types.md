
# Const Fingerprint Types

Supported client fingerprint values for NAC matching

## Structure

`ConstFingerprintTypes`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Family` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `Mfg` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `Model` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `Os` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    constFingerprintTypes := models.ConstFingerprintTypes{
        Family:               []string{
            "family9",
            "family0",
        },
        Mfg:                  []string{
            "mfg2",
        },
        Model:                []string{
            "model8",
            "model9",
        },
        Os:                   []string{
            "os3",
        },
    }

}
```

