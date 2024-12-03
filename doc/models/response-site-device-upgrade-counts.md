
# Response Site Device Upgrade Counts

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseSiteDeviceUpgradeCounts`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DownloadRequested` | `*int` | Optional | count of devices which cloud has requested to download firmware |
| `Downloaded` | `*int` | Optional | count of ap’s which have the firmware downloaded |
| `Failed` | `*int` | Optional | count of devices which have failed to upgrade |
| `RebootInProgress` | `*int` | Optional | count of devices which are rebooting |
| `Rebooted` | `*int` | Optional | count of devices which have rebooted successfully |
| `Scheduled` | `*int` | Optional | count of devices which cloud has scheduled an upgrade for |
| `Skipped` | `*int` | Optional | count of devices which skipped upgrade since requested version was same as running version. Use force to always upgrade |
| `Total` | `*int` | Optional | count of devices part of this upgrade |
| `Upgraded` | `*int` | Optional | count of devices which have upgraded successfully |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "download_requested": 102,
  "downloaded": 222,
  "failed": 130,
  "reboot_in_progress": 52,
  "rebooted": 112,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

