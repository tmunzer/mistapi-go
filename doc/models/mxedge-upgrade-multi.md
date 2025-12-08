
# Mxedge Upgrade Multi

*This model accepts additional fields of type interface{}.*

## Structure

`MxedgeUpgradeMulti`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AllowDowngrades` | [`*models.MxedgeUpgradeMultiAllowDowngrades`](../../doc/models/mxedge-upgrade-multi-allow-downgrades.md) | Optional | Whether downgrade is allowed when running version is higher than expected version for each service |
| `CanaryPhases` | `[]int` | Optional | Only if `strategy`==`canary`. Phases for canary deployment. Each phase represents percentage of devices that need to be upgraded in that phase. default is [1, 10, 50, 100] |
| `Channel` | [`*models.MxedgeUpgradeChannelEnum`](../../doc/models/mxedge-upgrade-channel-enum.md) | Optional | upgrade channel to follow. enum: `alpha`, `beta`, `stable`<br><br>**Default**: `"stable"` |
| `Distro` | `*string` | Optional | Distro upgrade, optional, to specific codename (e.g. bullseye) with highest qualified versions |
| `MaxFailurePercentage` | `*int` | Optional | Failure threshold before we stop the upgrade and mark it as failed<br><br>**Default**: `5` |
| `MxedgeIds` | `[]uuid.UUID` | Required | List of mxedge IDs to upgrade. If not specified, it means all the org mxedges. |
| `StartTime` | `*int` | Optional | Upgrade start time in epoch seconds, default is now |
| `Strategy` | [`*models.MxedgeUpgradeStrategyEnum`](../../doc/models/mxedge-upgrade-strategy-enum.md) | Optional | enum:<br><br>* `big_bang`: upgrade all at once, no orchestration<br>* `serial`: one at a time'<br>* `canary`: upgrade in phases<br><br>**Default**: `"big_bang"` |
| `Versions` | [`*models.MxedgeUpgradeVersion`](../../doc/models/mxedge-upgrade-version.md) | Optional | Version to upgrade for each service, `current` / `latest` / `default` / specific version (e.g. `2.5.100`).\nIgnored if distro upgrade, `tunterm`, `radsecproxy`, `mxagent`, `mxocproxy`, `mxdas` or `mxnacedge` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "canary_phases": null,
  "channel": "stable",
  "max_failure_percentage": 5,
  "mxedge_ids": [
    "00000bb2-0000-0000-0000-000000000000"
  ],
  "strategy": "big_bang",
  "allow_downgrades": {
    "mxagent": false,
    "mxdas": false,
    "mxocproxy": false,
    "radsecproxy": false,
    "tunterm": false
  },
  "distro": "distro4",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

