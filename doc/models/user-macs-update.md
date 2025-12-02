
# User Macs Update

## Structure

`UserMacsUpdate`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Errors` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `Updated` | `[]uuid.UUID` | Optional | - |

## Example (as JSON)

```json
{
  "errors": [
    "errors5"
  ],
  "updated": [
    "00001e9d-0000-0000-0000-000000000000",
    "00001e9e-0000-0000-0000-000000000000"
  ]
}
```

