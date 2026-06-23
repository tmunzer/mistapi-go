
# Psk Portal Passphrase Rules

Passphrase generation rules for PSKs created through a portal

## Structure

`PskPortalPassphraseRules`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AlphabetsEnabled` | `*bool` | Optional | Whether generated passphrases may include alphabetic characters<br><br>**Default**: `true` |
| `Length` | `*int` | Optional | Fixed generated passphrase length used when min and max length are not both valid<br><br>**Constraints**: `>= 8`, `<= 63` |
| `MaxLength` | `*int` | Optional | Maximum generated passphrase length when paired with a valid `min_length`. If `max_length` or `min_length` is invalid, the portal uses `length`; if `length` is unset or invalid, it uses 8. Valid values are integers from 8 through 63, and `max_length` must be greater than `min_length`<br><br>**Constraints**: `>= 8`, `<= 63` |
| `MinLength` | `*int` | Optional | Minimum generated passphrase length when paired with a valid `max_length`. If `max_length` or `min_length` is invalid, the portal uses `length`; if `length` is unset or invalid, it uses 8. Valid values are integers from 8 through 63, and `max_length` must be greater than `min_length`<br><br>**Constraints**: `>= 8`, `<= 63` |
| `NumericsEnabled` | `*bool` | Optional | Whether generated passphrases may include numeric characters<br><br>**Default**: `true` |
| `Symbols` | `*string` | Optional | Allowed symbol characters for generated passphrases |
| `SymbolsEnabled` | `*bool` | Optional | Whether generated passphrases may include symbols<br><br>**Default**: `true` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    pskPortalPassphraseRules := models.PskPortalPassphraseRules{
        AlphabetsEnabled:     models.ToPointer(true),
        Length:               models.ToPointer(63),
        MaxLength:            models.ToPointer(63),
        MinLength:            models.ToPointer(63),
        NumericsEnabled:      models.ToPointer(true),
        Symbols:              models.ToPointer("()[]{}_%@#&$"),
        SymbolsEnabled:       models.ToPointer(true),
    }

}
```

