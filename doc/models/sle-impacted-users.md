
# Sle Impacted Users

Paginated list of users impacted by an SLE metric

## Structure

`SleImpactedUsers`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Classifier` | `string` | Required | Requested SLE classifier filter applied to the query |
| `Clients` | [`[]models.SleImpactedUsersClient`](../../doc/models/sle-impacted-users-client.md) | Optional | Impacted client rows returned for an SLE query<br><br>**Constraints**: *Unique Items Required* |
| `End` | `float64` | Required | Last timestamp in the impacted users window |
| `Failure` | `string` | Required | Requested SLE failure filter applied to the query |
| `Limit` | `float64` | Required | Maximum number of impacted user rows returned per page |
| `Metric` | `string` | Required | SLE metric name used for the impacted users query<br><br>**Constraints**: *Minimum Length*: `1` |
| `Page` | `float64` | Required | Current page number for impacted user results |
| `Start` | `float64` | Required | First timestamp in the impacted users window |
| `TotalCount` | `float64` | Required | Number of impacted user rows matching the query |
| `Users` | [`[]models.SleImpactedUsersUser`](../../doc/models/sle-impacted-users-user.md) | Optional | Impacted user rows returned for an SLE query<br><br>**Constraints**: *Unique Items Required* |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    sleImpactedUsers := models.SleImpactedUsers{
        Classifier:           "classifier4",
        Clients:              []models.SleImpactedUsersClient{
            models.SleImpactedUsersClient{
                Degraded:             models.ToPointer(float64(166.58)),
                Duration:             models.ToPointer(float64(39.64)),
                Gateways:             []models.SleImpactedClientGateway{
                    models.SleImpactedClientGateway{
                        ChassisMac:           models.ToPointer("chassis_mac4"),
                        GatewayMac:           models.ToPointer("gateway_mac2"),
                        GatewayName:          models.ToPointer("gateway_name0"),
                        Interfaces:           []string{
                            "interfaces5",
                            "interfaces6",
                        },
                    },
                },
                Mac:                  models.ToPointer("mac2"),
                Name:                 models.ToPointer("name8"),
            },
            models.SleImpactedUsersClient{
                Degraded:             models.ToPointer(float64(166.58)),
                Duration:             models.ToPointer(float64(39.64)),
                Gateways:             []models.SleImpactedClientGateway{
                    models.SleImpactedClientGateway{
                        ChassisMac:           models.ToPointer("chassis_mac4"),
                        GatewayMac:           models.ToPointer("gateway_mac2"),
                        GatewayName:          models.ToPointer("gateway_name0"),
                        Interfaces:           []string{
                            "interfaces5",
                            "interfaces6",
                        },
                    },
                },
                Mac:                  models.ToPointer("mac2"),
                Name:                 models.ToPointer("name8"),
            },
            models.SleImpactedUsersClient{
                Degraded:             models.ToPointer(float64(166.58)),
                Duration:             models.ToPointer(float64(39.64)),
                Gateways:             []models.SleImpactedClientGateway{
                    models.SleImpactedClientGateway{
                        ChassisMac:           models.ToPointer("chassis_mac4"),
                        GatewayMac:           models.ToPointer("gateway_mac2"),
                        GatewayName:          models.ToPointer("gateway_name0"),
                        Interfaces:           []string{
                            "interfaces5",
                            "interfaces6",
                        },
                    },
                },
                Mac:                  models.ToPointer("mac2"),
                Name:                 models.ToPointer("name8"),
            },
        },
        End:                  float64(42.06),
        Failure:              "failure2",
        Limit:                float64(207.96),
        Metric:               "metric0",
        Page:                 float64(81.1),
        Start:                float64(254.12),
        TotalCount:           float64(112.24),
        Users:                []models.SleImpactedUsersUser{
            models.SleImpactedUsersUser{
                ApMac:                models.ToPointer("ap_mac8"),
                ApName:               models.ToPointer("ap_name4"),
                Degraded:             models.ToPointer(float64(137.76)),
                DeviceOs:             models.ToPointer("device_os6"),
                DeviceType:           models.ToPointer("device_type4"),
            },
        },
    }

}
```

