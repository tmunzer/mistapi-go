
# Response Ssr Upgrade Status Targets

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseSsrUpgradeStatusTargets`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Failed` | `[]string` | Required | **Constraints**: *Unique Items Required* |
| `Queued` | `[]string` | Required | **Constraints**: *Unique Items Required* |
| `Success` | `[]string` | Required | **Constraints**: *Unique Items Required* |
| `Upgrading` | `[]string` | Required | **Constraints**: *Unique Items Required* |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "failed": [
    "failed6",
    "failed5"
  ],
  "queued": [
    "queued8",
    "queued9",
    "queued0"
  ],
  "success": [
    "success2",
    "success3"
  ],
  "upgrading": [
    "upgrading9",
    "upgrading0",
    "upgrading1"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

