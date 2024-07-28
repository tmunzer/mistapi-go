
# Service

Applications used for the Gateway configurations

## Structure

`Service`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Addresses` | `[]string` | Optional | if `type`==`custom`, ip subnets (e.g. 10.0.0.0/8) |
| `AppCategories` | `[]string` | Optional | when `type`==`app_categories`, list of application categories are available through /api/v1/const/app_categories |
| `AppSubcategories` | `[]string` | Optional | when `type`==`app_categories`, list of application categories are available through /api/v1/const/app_subcategories |
| `Apps` | `[]string` | Optional | when `type`==`apps`, list of applications are available through:<br><br>* /api/v1/const/applications<br>* /api/v1/const/gateway_applications<br>* /insight/top_app_by-bytes?wired=true |
| `CreatedTime` | `*float64` | Optional | - |
| `Description` | `*string` | Optional | - |
| `Dscp` | [`*models.ServiceDscp`](../../doc/models/containers/service-dscp.md) | Optional | This is a container for one-of cases. |
| `FailoverPolicy` | [`*models.ServiceFailoverPolicyEnum`](../../doc/models/service-failover-policy-enum.md) | Optional | **Default**: `"revertable"` |
| `Hostnames` | `[]string` | Optional | if `type`==`custom`, web filtering |
| `Id` | `*uuid.UUID` | Optional | - |
| `MaxJitter` | [`*models.ServiceMaxJitter`](../../doc/models/containers/service-max-jitter.md) | Optional | This is a container for one-of cases. |
| `MaxLatency` | [`*models.ServiceMaxLatency`](../../doc/models/containers/service-max-latency.md) | Optional | This is a container for one-of cases. |
| `MaxLoss` | [`*models.ServiceMaxLoss`](../../doc/models/containers/service-max-loss.md) | Optional | This is a container for one-of cases. |
| `ModifiedTime` | `*float64` | Optional | - |
| `Name` | `*string` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `SleEnabled` | `*bool` | Optional | whether to enable measure SLE<br>**Default**: `false` |
| `Specs` | [`[]models.ServiceSpec`](../../doc/models/service-spec.md) | Optional | when `type`==`custom`, optional, if it doesn't exist, http and https is assumed |
| `SsrRelaxedTcpStateEnforcement` | `*bool` | Optional | **Default**: `false` |
| `TrafficClass` | [`*models.ServiceTrafficClassEnum`](../../doc/models/service-traffic-class-enum.md) | Optional | when `traffic_type`==`custom`<br>**Default**: `"best_effort"` |
| `TrafficType` | `*string` | Optional | values from `/api/v1/consts/traffic_types`<br><br>* when `type`==`apps`, we''ll choose traffic_type automatically<br>* when `type`==`addresses` or `type`==`hostnames`, you can provide your own settings (optional)<br>**Default**: `"data_best_effort"` |
| `Type` | [`*models.ServiceTypeEnum`](../../doc/models/service-type-enum.md) | Optional | **Default**: `"custom"` |
| `Urls` | `[]string` | Optional | when `type`==`urls`, no need for spec as URL can encode the ports being used |

## Example (as JSON)

```json
{
  "app_categories": [
    "Sports"
  ],
  "app_subcategories": [
    "Shopping"
  ],
  "apps": [
    "office365",
    "okta"
  ],
  "failover_policy": "revertable",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "sle_enabled": false,
  "ssr_relaxed_tcp_state_enforcement": false,
  "traffic_class": "best_effort",
  "traffic_type": "data_best_effort",
  "type": "custom",
  "addresses": [
    "addresses8",
    "addresses9",
    "addresses0"
  ],
  "created_time": 40.82
}
```

