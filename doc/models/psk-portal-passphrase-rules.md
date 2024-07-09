
# Psk Portal Passphrase Rules

## Structure

`PskPortalPassphraseRules`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AlphabertsEnabled` | `*bool` | Optional | **Default**: `true` |
| `Length` | `*int` | Optional | **Constraints**: `>= 8`, `<= 63` |
| `MaxLength` | `*int` | Optional | for valid `max_length` and `min_length`, passphrase size is set randomly from that range.<br><br>- if `max_length` and/or `min_length` are invalid, passphrase size is equal to `length` parameter<br>- if `length` is not set or is invalid, default passphrase size is 8.<br>  valid `max_length`, `min_length`, `length` should be an integer between 8 to 63. Also, `max_length` > `min_length`<br>**Constraints**: `>= 8`, `<= 63` |
| `MinLength` | `*int` | Optional | for valid `max_length` and `min_length`, passphrase size is set randomly from that range.<br><br>- if `max_length` and/or `min_length` are invalid, passphrase size is equal to `length` parameter<br>- if `length` is not set or is invalid, default passphrase size is 8.<br>  valid `max_length`, `min_length`, `length` should be an integer between 8 to 63. Also, `max_length` > `min_length`<br>**Constraints**: `>= 8`, `<= 63` |
| `NumericsEnabled` | `*bool` | Optional | **Default**: `true` |
| `Symbols` | `*string` | Optional | - |
| `SymbolsEnabled` | `*bool` | Optional | **Default**: `true` |

## Example (as JSON)

```json
{
  "alphaberts_enabled": true,
  "numerics_enabled": true,
  "symbols": "()[]{}_%@#&$",
  "symbols_enabled": true,
  "length": 124,
  "max_length": 192,
  "min_length": 68
}
```

