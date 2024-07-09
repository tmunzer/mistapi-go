
# Response Ssr Upgrade Status Targets

## Structure

`ResponseSsrUpgradeStatusTargets`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Failed` | `[]string` | Required | **Constraints**: *Unique Items Required* |
| `Queued` | `[]string` | Required | **Constraints**: *Unique Items Required* |
| `Success` | `[]string` | Required | **Constraints**: *Unique Items Required* |
| `Upgrading` | `[]string` | Required | **Constraints**: *Unique Items Required* |

## Example (as JSON)

```json
{
  "failed": [
    "failed4"
  ],
  "queued": [
    "queued6"
  ],
  "success": [
    "success0",
    "success1",
    "success2"
  ],
  "upgrading": [
    "upgrading7"
  ]
}
```

