
# Synthetictest Config Custom Probe

Custom probe definition for synthetic tests

## Structure

`SynthetictestConfigCustomProbe`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Aggressiveness` | [`*models.SynthetictestConfigAggressivenessEnum`](../../doc/models/synthetictest-config-aggressiveness-enum.md) | Optional | Aggressiveness level for a synthetic test. enum: `auto`, `high`, `med`, `low`<br><br>**Default**: `"auto"` |
| `Target` | `*string` | Optional | Can be URL (e.g. http://x.com, https://x.com:8080/path/to/resource), IP address, or IP:port combination |
| `Threshold` | `*int` | Optional | Response-time threshold for this custom probe, in milliseconds |
| `Type` | [`*models.SynthetictestConfigCustomProbeTypeEnum`](../../doc/models/synthetictest-config-custom-probe-type-enum.md) | Optional | enum: `application`, `curl`, `icmp`, `reachability`, `tcp`<br><br>**Default**: `"icmp"` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    synthetictestConfigCustomProbe := models.SynthetictestConfigCustomProbe{
        Aggressiveness:       models.ToPointer(models.SynthetictestConfigAggressivenessEnum_AUTO),
        Target:               models.ToPointer("10.3.5.3:8080"),
        Threshold:            models.ToPointer(100),
        Type:                 models.ToPointer(models.SynthetictestConfigCustomProbeTypeEnum_ICMP),
    }

}
```

