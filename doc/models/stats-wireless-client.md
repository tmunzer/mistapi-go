
# Stats Wireless Client

Wireless client connection, traffic, and location statistics

*This model accepts additional fields of type interface{}.*

## Structure

`StatsWirelessClient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Accuracy` | `*int` | Optional | Estimated client location accuracy, in meter |
| `AirespaceIfname` | `*string` | Optional | RADIUS Airespace interface name reported for the wireless client, when available |
| `Airwatch` | [`*models.StatsWirelessClientAirwatch`](../../doc/models/stats-wireless-client-airwatch.md) | Optional | AirWatch authorization information reported for a wireless client |
| `Annotation` | `*string` | Optional | User-visible annotation label applied to the wireless client |
| `ApId` | `uuid.UUID` | Required | AP ID the client is connected to |
| `ApMac` | `string` | Required | AP the client is connected to |
| `AssocTime` | `*int` | Optional | Time when the wireless client associated to the AP, in epoch seconds |
| `Band` | [`models.Dot11BandEnum`](../../doc/models/dot-11-band-enum.md) | Required | enum: `24`, `5`, `5-dedicated`, `5-selectable`, `6`, `6-dedicated`, `6-selectable` |
| `Bssid` | `*string` | Optional | AP radio BSSID serving the wireless client connection |
| `Channel` | `int` | Required | Radio channel used by the wireless client connection |
| `DualBand` | `*bool` | Optional | Whether the client is dual_band capable (determined by whether we’ve seen probe requests from both bands) |
| `Family` | `*string` | Optional | Device family, through fingerprinting. iPod / Nexus Galaxy / Windows Mobile or CE … |
| `Group` | `*string` | Optional | Client group label reported for the wireless client |
| `Guest` | [`*models.Guest`](../../doc/models/guest.md) | Optional | Guest authorization record at site scope |
| `Hostname` | `*string` | Optional | DHCP hostname learned for the wireless client |
| `IdleTime` | `*float64` | Optional | How long, in seconds, has the client been idle (since the last RX packet) |
| `Ip` | `*string` | Optional | Current IP address reported for the wireless client |
| `IsGuest` | `bool` | Required | Whether this is a guest<br><br>**Default**: `false` |
| `KeyMgmt` | `string` | Required | Security key-management and cipher suite used by the wireless client |
| `LastSeen` | `models.Optional[float64]` | Optional, Read-only | Timestamp indicating when the entity was last seen |
| `Mac` | `string` | Required | Wireless client MAC address observed by Mist |
| `Manufacture` | `*string` | Optional | Device manufacture, through fingerprinting or OUI |
| `MapId` | `*uuid.UUID` | Optional | Estimated client location - map_id |
| `Model` | `*string` | Optional | Device model, may be available if we can identify them |
| `NumLocatingAps` | `*int` | Optional | Number of APs used to locate this client |
| `Os` | `*string` | Optional | Device os, through fingerprinting |
| `PowerSaving` | `*bool` | Optional | If it’s currently in power-save mode |
| `Proto` | [`models.Dot11ProtoEnum`](../../doc/models/dot-11-proto-enum.md) | Required | enum: `a`, `ac`, `ax`, `b`, `be`, `g`, `n` |
| `PskId` | `*uuid.UUID` | Optional | PSK id (if multi-psk is used) |
| `Rssi` | `float64` | Required | Received signal strength indicator for the wireless client, in dBm |
| `Rssizones` | [`[]models.StatsWirelessClientRssiZone`](../../doc/models/stats-wireless-client-rssi-zone.md) | Optional | RSSI zone memberships for the wireless client, including when each membership began |
| `RxBps` | `models.Optional[int64]` | Optional, Read-only | Rate of receiving traffic, bits/seconds, last known |
| `RxBytes` | `models.Optional[int64]` | Optional, Read-only | Amount of traffic received since connection |
| `RxPkts` | `models.Optional[int64]` | Optional, Read-only | Amount of packets received since connection |
| `RxRate` | `models.Optional[float64]` | Optional, Read-only | Receive data rate reported for a wireless or mesh link, in Mbps |
| `RxRetries` | `models.Optional[int]` | Optional, Read-only | Amount of rx retries |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `Snr` | `float64` | Required | Signal-to-noise ratio for the wireless client connection |
| `Ssid` | `string` | Required | Wireless SSID used by the client connection |
| `TxBps` | `models.Optional[int64]` | Optional, Read-only | Rate of transmitting traffic, bits/seconds, last known |
| `TxBytes` | `models.Optional[int64]` | Optional, Read-only | Amount of traffic sent since connection |
| `TxPkts` | `models.Optional[int64]` | Optional, Read-only | Amount of packets sent since connection |
| `TxRate` | `models.Optional[float64]` | Optional, Read-only | Transmit data rate reported for a wireless or mesh link, in Mbps |
| `TxRetries` | `models.Optional[int]` | Optional, Read-only | Amount of tx retries |
| `Type` | `*string` | Optional | Client’s type, regular / vip / resource / blocked (if client object is created) |
| `Uptime` | `*float64` | Optional | How long, in seconds, has the client been connected |
| `Username` | `*string` | Optional | User identity learned from 802.1X, per-user PSK, or user portal authentication |
| `Vbeacons` | [`[]models.StatsWirelessClientVbeacon`](../../doc/models/stats-wireless-client-vbeacon.md) | Optional | Virtual beacon associations for the wireless client, including when each began |
| `VlanId` | `*string` | Optional | VLAN ID, could be empty (from older AP) |
| `WlanId` | `uuid.UUID` | Required | WLAN ID the client is connected to |
| `WxruleId` | `*uuid.UUID` | Optional | Current WxlanRule using for a Client or an authorized Guest (portal user). null if default rule is matched. |
| `WxruleUsage` | [`[]models.StatsWirelessClientWxruleUsage`](../../doc/models/stats-wireless-client-wxrule-usage.md) | Optional | Current WxLAN rule usage counters keyed by wireless tag |
| `X` | `*float64` | Optional | Estimated client location in pixels |
| `XM` | `*float64` | Optional | Estimated client location in meter |
| `Y` | `*float64` | Optional | Estimated client location in pixels |
| `YM` | `*float64` | Optional | Estimated client location in meter |
| `Zones` | [`[]models.StatsWirelessClientZone`](../../doc/models/stats-wireless-client-zone.md) | Optional | Zone memberships for the wireless client, including when each membership began |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    statsWirelessClient := models.StatsWirelessClient{
        Accuracy:             models.ToPointer(48),
        AirespaceIfname:      models.ToPointer("airespace_ifname8"),
        Airwatch:             models.ToPointer(models.StatsWirelessClientAirwatch{
            Authorized:           false,
        }),
        Annotation:           models.ToPointer("annotation8"),
        ApId:                 uuid.MustParse("00001388-0000-0000-0000-000000000000"),
        ApMac:                "ap_mac0",
        AssocTime:            models.ToPointer(6),
        Band:                 models.Dot11BandEnum_ENUM5SELECTABLE,
        Channel:              2,
        IsGuest:              false,
        KeyMgmt:              "key_mgmt0",
        LastSeen:             models.NewOptional(models.ToPointer(float64(1470417522))),
        Mac:                  "mac2",
        Proto:                models.Dot11ProtoEnum_AC,
        Rssi:                 float64(152),
        RxBps:                models.NewOptional(models.ToPointer(int64(60003))),
        RxBytes:              models.NewOptional(models.ToPointer(int64(8515104416))),
        RxPkts:               models.NewOptional(models.ToPointer(int64(57770567))),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        Snr:                  float64(21.74),
        Ssid:                 "ssid6",
        TxBps:                models.NewOptional(models.ToPointer(int64(634301))),
        TxBytes:              models.NewOptional(models.ToPointer(int64(211217389682))),
        TxPkts:               models.NewOptional(models.ToPointer(int64(812204062))),
        WlanId:               uuid.MustParse("0000263e-0000-0000-0000-000000000000"),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

