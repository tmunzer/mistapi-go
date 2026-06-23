
# Response Client Search

Paginated wireless client search response

## Structure

`ResponseClientSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `float64` | Required | Search window end timestamp for wireless clients, in epoch seconds |
| `Limit` | `int` | Required | Maximum number of wireless client results requested |
| `Next` | `*string` | Optional | URL for the next page of wireless client results |
| `Results` | [`[]models.ClientWireless`](../../doc/models/client-wireless.md) | Required | Wireless client records returned by a search response<br><br>**Constraints**: *Unique Items Required* |
| `Start` | `float64` | Required | Search window start timestamp for wireless clients, in epoch seconds |
| `Total` | `int` | Required | Number of wireless client records matching the search |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseClientSearch := models.ResponseClientSearch{
        End:                  float64(137.22),
        Limit:                16,
        Next:                 models.ToPointer("next8"),
        Results:              []models.ClientWireless{
            models.ClientWireless{
                Ap:                   []string{
                    "a83a79a947ee",
                    "003e73170b4c",
                },
                AppVersion:           []string{
                    "0.100.3",
                },
                Band:                 models.ToPointer("5"),
                Device:               []string{
                    "Mac",
                },
                Ftc:                  models.ToPointer(false),
                Hardware:             models.ToPointer("Apple Wi-Fi adapter"),
                Hostname:             []string{
                    "hostname-a",
                    "hostname-b",
                },
                Ip:                   []string{
                    "10.5.23.43",
                    "192.168.0.2",
                },
                LastAp:               models.ToPointer("a83a79a947ee"),
                LastDevice:           models.ToPointer("Zebra"),
                LastFirmware:         models.ToPointer("wl0: Jan 20 2024 04:08:41 version 20.103.12.0.8.7.171 FWID 01-e09d2675"),
                LastHostname:         models.ToPointer("hostname-a"),
                LastIp:               models.ToPointer("10.100.0.157"),
                LastModel:            models.ToPointer("MBP 16\\\" M1 2021"),
                LastOs:               models.ToPointer("Sonoma"),
                LastOsVersion:        models.ToPointer("14.4.1 (Build 23E224)"),
                LastPskId:            models.ToPointer(uuid.MustParse("abf7dc5c-bb51-4bb7-93b6-5547400ffe11")),
                LastPskName:          models.ToPointer("iot"),
                LastSsid:             models.ToPointer("john@mycorp.net"),
                LastVlan:             models.ToPointer(10),
                LastWlanId:           models.ToPointer(uuid.MustParse("e5d67b07-aae8-494b-8584-cbc20c8110aa")),
                Mac:                  models.ToPointer("bcd074000000"),
                Mfg:                  models.ToPointer("Apple"),
                Model:                models.ToPointer("MBP 16\\\" M1 2021"),
                OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
                Os:                   []string{
                    "Sonoma",
                },
                OsVersion:            []string{
                    "14.4.1 (Build 23E224)",
                },
                Protocol:             models.ToPointer("ax"),
                PskId:                []uuid.UUID{
                    uuid.MustParse("abf7dc5c-bb51-4bb7-93b6-5547400ffe11"),
                },
                PskName:              []string{
                    "iot",
                },
                SdkVersion:           []string{
                    "0.100.3",
                },
                SiteId:               models.ToPointer(uuid.MustParse("25ff5219-9be7-4db9-907d-0c9b60445147")),
                SiteIds:              []uuid.UUID{
                    uuid.MustParse("25ff5219-9be7-4db9-907d-0c9b60445147"),
                },
                Ssid:                 []string{
                    "IoT SSID",
                },
                Username:             []string{
                    "user@corp.com",
                },
                Vlan:                 []int{
                    10,
                },
                WlanId:               []uuid.UUID{
                    uuid.MustParse("e5d67b07-aae8-494b-8584-cbc20c8110aa"),
                },
            },
        },
        Start:                float64(93.28),
        Total:                110,
    }

}
```

