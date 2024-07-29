
# Response Upgrade Org Devices

## Structure

`ResponseUpgradeOrgDevices`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `EnableP2p` | `*bool` | Optional | whether to allow local AP-to-AP FW upgrade |
| `Force` | `*bool` | Optional | whether to force upgrade when requested version is same as running version |
| `Id` | `*uuid.UUID` | Optional | - |
| `Strategy` | [`*models.DeviceUpgradeStrategyEnum`](../../doc/models/device-upgrade-strategy-enum.md) | Optional | enum: `big_bang` (upgrade all at once), `canary`, `rrm`, `serial` (one at a time)<br>**Default**: `"big_bang"` |
| `TargetVersion` | `*string` | Optional | version to upgrade to |
| `Upgrades` | [`[]models.ResponseUpgradeOrgDevice`](../../doc/models/response-upgrade-org-device.md) | Optional | - |

## Example (as JSON)

```json
{
  "id": "31223085-405d-4b64-8aea-9c5b98098b4b",
  "strategy": "big_bang",
  "target_version": "0.14.29411",
  "enable_p2p": false,
  "force": false
}
```

