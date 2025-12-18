## Webhooks Handler



## Events

Events available in this group. Subscribe to receive webhook notifications when these events occur.

| Name |
|  --- |
| [alarms](../../../doc/events/webhooks/webhooks/alarms.md) |
| [asset_raw_rssi](../../../doc/events/webhooks/webhooks/asset-raw-rssi.md) |
| [audits](../../../doc/events/webhooks/webhooks/audits.md) |
| [client_info](../../../doc/events/webhooks/webhooks/client-info.md) |
| [client_join](../../../doc/events/webhooks/webhooks/client-join.md) |
| [client_latency](../../../doc/events/webhooks/webhooks/client-latency.md) |
| [client_sessions](../../../doc/events/webhooks/webhooks/client-sessions.md) |
| [device_events](../../../doc/events/webhooks/webhooks/device-events.md) |
| [device_updowns](../../../doc/events/webhooks/webhooks/device-updowns.md) |
| [discovered_raw_rssi](../../../doc/events/webhooks/webhooks/discovered-raw-rssi.md) |
| [guest_authorizations](../../../doc/events/webhooks/webhooks/guest-authorizations.md) |
| [location](../../../doc/events/webhooks/webhooks/location.md) |
| [location_asset](../../../doc/events/webhooks/webhooks/location-asset.md) |
| [location_centrak](../../../doc/events/webhooks/webhooks/location-centrak.md) |
| [location_client](../../../doc/events/webhooks/webhooks/location-client.md) |
| [location_sdk](../../../doc/events/webhooks/webhooks/location-sdk.md) |
| [location_unclient](../../../doc/events/webhooks/webhooks/location-unclient.md) |
| [mxedge_events](../../../doc/events/webhooks/webhooks/mxedge-events.md) |
| [nac_accounting](../../../doc/events/webhooks/webhooks/nac-accounting.md) |
| [nac_events](../../../doc/events/webhooks/webhooks/nac-events.md) |
| [occupancy_alerts](../../../doc/events/webhooks/webhooks/occupancy-alerts.md) |
| [ping](../../../doc/events/webhooks/webhooks/ping.md) |
| [rssizone](../../../doc/events/webhooks/webhooks/rssizone.md) |
| [sdkclient_scan_data](../../../doc/events/webhooks/webhooks/sdkclient-scan-data.md) |
| [site_sle](../../../doc/events/webhooks/webhooks/site-sle.md) |
| [wifi_conn_raw](../../../doc/events/webhooks/webhooks/wifi-conn-raw.md) |
| [wifi_unconn_raw](../../../doc/events/webhooks/webhooks/wifi-unconn-raw.md) |
| [zone](../../../doc/events/webhooks/webhooks/zone.md) |

## SDK Usage Example

```go
package mistapi

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "mistapi/events/webhooks"
)

// WebhooksGinEventHandler receives and processes incoming events
func WebhooksGinEventHandler(c *gin.Context) {
    handler :=  webhooks.NewWebhooksHandler()
    parsingResult := handler.ParseEvent(c.Request)
    if event, ok := parsingResult.AsAlarms(); ok {
        c.JSON(200, map[string]any{
            "status":    "success",
            "eventInfo": fmt.Sprintf("Alarms event received %v", event),
        })
    } else if event, ok := parsingResult.AsAssetRawRssi(); ok {
        c.JSON(200, map[string]any{
            "status":    "success",
            "eventInfo": fmt.Sprintf("AssetRawRssi event received %v", event),
        })
    } else if event, ok := parsingResult.AsAudits(); ok {
        c.JSON(200, map[string]any{
            "status":    "success",
            "eventInfo": fmt.Sprintf("Audits event received %v", event),
        })
    } else if parsingResult.AsUnknown() {
        c.JSON(200, map[string]any{
            "status":    "success",
            "eventInfo": "Unknown event received",
        })
    } else {
        c.JSON(400, map[string]any{"status": "failure"})
    }
}
```

