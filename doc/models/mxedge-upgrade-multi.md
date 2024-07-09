
# Mxedge Upgrade Multi

## Structure

`MxedgeUpgradeMulti`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AllowDowngrades` | [`*models.MxedgeUpgradeMultiAllowDowngrades`](../../doc/models/mxedge-upgrade-multi-allow-downgrades.md) | Optional | whether downgrade is allowed when running version is higher than expected version for each service |
| `Channel` | [`*models.MxedgeUpgradeChannelEnum`](../../doc/models/mxedge-upgrade-channel-enum.md) | Optional | upgrade channel to follow, stable (default) / beta / alpha<br>**Default**: `"stable"` |
| `Distro` | `*string` | Optional | distro upgrade, optional, to specific codename (e.g. bullseye) with highest qualified versions |
| `MxedgeIds` | `[]uuid.UUID` | Required | list of mxedge IDs to upgrade. If not specified, it means all the org mxedges. |
| `Strategy` | [`*models.MxedgeUpgradeStrategyEnum`](../../doc/models/mxedge-upgrade-strategy-enum.md) | Optional | * `big_bang`: upgrade all at once<br>* `serial`: one at a time<br>**Default**: `"big_bang"` |
| `Versions` | [`*models.MxedgeUpgradeVersion`](../../doc/models/mxedge-upgrade-version.md) | Optional | version to upgrade for each service, `current` / `latest` / `default` / specific version (e.g. `2.5.100`).\nignored if distro upgrade |

## Example (as JSON)

```json
{
  "channel": "stable",
  "mxedge_ids": [
    "00002230-0000-0000-0000-000000000000"
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
  "versions": {
    "mxagent": "mxagent2",
    "mxdas": "mxdas6",
    "mxocproxy": "mxocproxy0",
    "radsecproxy": "radsecproxy4",
    "tunterm": "tunterm0"
  }
}
```

