
# Ssr Upgrade Multi

Request to upgrade multiple SSR devices

*This model accepts additional fields of type interface{}.*

## Structure

`SsrUpgradeMulti`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Channel` | [`*models.SsrUpgradeChannelEnum`](../../doc/models/ssr-upgrade-channel-enum.md) | Optional | upgrade channel to follow. enum: `alpha`, `beta`, `stable`<br><br>**Default**: `"stable"` |
| `DeviceIds` | `[]uuid.UUID` | Required | List of 128T device IDs to upgrade |
| `RebootAt` | `*int` | Optional | Reboot start time in epoch seconds, default is start_time, -1 disables reboot |
| `StartTime` | `*int` | Optional | 128T firmware download start time in epoch seconds, default is now, -1 disables download |
| `Strategy` | [`*models.SsrUpgradeStrategyEnum`](../../doc/models/ssr-upgrade-strategy-enum.md) | Optional | enum:<br><br>* `big_bang`: upgrade all at once<br>* `serial`: one at a time<br><br>**Default**: `"big_bang"` |
| `Version` | `*string` | Optional | 128T firmware version to upgrade (e.g. 5.3.0-93)<br><br>**Constraints**: *Minimum Length*: `1` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    ssrUpgradeMulti := models.SsrUpgradeMulti{
        Channel:              models.ToPointer(models.SsrUpgradeChannelEnum_STABLE),
        DeviceIds:            []uuid.UUID{
            uuid.MustParse("000002a9-0000-0000-0000-000000000000"),
            uuid.MustParse("000002aa-0000-0000-0000-000000000000"),
            uuid.MustParse("000002ab-0000-0000-0000-000000000000"),
        },
        RebootAt:             models.ToPointer(112),
        StartTime:            models.ToPointer(182),
        Strategy:             models.ToPointer(models.SsrUpgradeStrategyEnum_BIGBANG),
        Version:              models.ToPointer("version8"),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

