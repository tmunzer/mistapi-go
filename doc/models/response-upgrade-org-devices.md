
# Response Upgrade Org Devices

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseUpgradeOrgDevices`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `EnableP2p` | `*bool` | Optional | whether to allow local AP-to-AP FW upgrade |
| `Force` | `*bool` | Optional | whether to force upgrade when requested version is same as running version |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organnization |
| `Strategy` | [`*models.UpgradeDeviceStrategyEnum`](../../doc/models/upgrade-device-strategy-enum.md) | Optional | For APs only. enum: `big_bang` (upgrade all at once), `canary`, `rrm`, `serial` (one at a time)<br>**Default**: `"big_bang"` |
| `TargetVersion` | `*string` | Optional | version to upgrade to |
| `Upgrades` | [`[]models.UpgradeOrgDevicesUpgrade`](../../doc/models/upgrade-org-devices-upgrade.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "strategy": "big_bang",
  "target_version": "0.14.29411",
  "enable_p2p": false,
  "force": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

