
# Const Fingerprint Types

*This model accepts additional fields of type interface{}.*

## Structure

`ConstFingerprintTypes`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Family` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `Mfg` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `Model` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `Os` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "family": [
    "family9",
    "family0"
  ],
  "mfg": [
    "mfg2"
  ],
  "model": [
    "model8",
    "model9"
  ],
  "os": [
    "os3"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

