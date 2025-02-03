
# Service

APplications used for the Gateway configurations

*This model accepts additional fields of type interface{}.*

## Structure

`Service`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Addresses` | `[]string` | Optional | If `type`==`custom`, ip subnets (e.g. 10.0.0.0/8) |
| `AppCategories` | `[]string` | Optional | When `type`==`app_categories`, list of application categories are available through /api/v1/const/app_categories |
| `AppSubcategories` | `[]string` | Optional | When `type`==`app_categories`, list of application categories are available through /api/v1/const/app_subcategories |
| `Apps` | `[]string` | Optional | When `type`==`apps`, list of applications are available through:<br><br>* /api/v1/const/applications<br>* /api/v1/const/gateway_applications<br>* /insight/top_app_by-bytes?wired=true |
| `ClientLimitDown` | `*int` | Optional | 0 means unlimited<br>**Default**: `0`<br>**Constraints**: `>= 0`, `<= 107374182` |
| `ClientLimitUp` | `*int` | Optional | 0 means unlimited<br>**Default**: `0`<br>**Constraints**: `>= 0`, `<= 107374182` |
| `CreatedTime` | `*float64` | Optional | When the object has been created, in epoch |
| `Description` | `*string` | Optional | - |
| `Dscp` | [`*models.ServiceDscp`](../../doc/models/containers/service-dscp.md) | Optional | For SSR only, when `traffic_type`==`custom`. 0-63 or variable |
| `FailoverPolicy` | [`*models.ServiceFailoverPolicyEnum`](../../doc/models/service-failover-policy-enum.md) | Optional | enum: `non_revertable`, `none`, `revertable`<br>**Default**: `"revertable"` |
| `Hostnames` | `[]string` | Optional | If `type`==`custom`, web filtering |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organnization |
| `MaxJitter` | [`*models.ServiceMaxJitter`](../../doc/models/containers/service-max-jitter.md) | Optional | For SSR only, when `traffic_type`==`custom`, for uplink selection. 0-2147483647 or variable |
| `MaxLatency` | [`*models.ServiceMaxLatency`](../../doc/models/containers/service-max-latency.md) | Optional | For SSR only, when `traffic_type`==`custom`, for uplink selection. 0-2147483647 or variable |
| `MaxLoss` | [`*models.ServiceMaxLoss`](../../doc/models/containers/service-max-loss.md) | Optional | For SSR only, when `traffic_type`==`custom`, for uplink selection. 0-100 or variable |
| `ModifiedTime` | `*float64` | Optional | When the object has been modified for the last time, in epoch |
| `Name` | `*string` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `ServiceLimitDown` | `*int` | Optional | 0 means unlimited<br>**Default**: `0`<br>**Constraints**: `>= 0`, `<= 107374182` |
| `ServiceLimitUp` | `*int` | Optional | 0 means unlimited<br>**Default**: `0`<br>**Constraints**: `>= 0`, `<= 107374182` |
| `SleEnabled` | `*bool` | Optional | Whether to enable measure SLE<br>**Default**: `false` |
| `Specs` | [`[]models.ServiceSpec`](../../doc/models/service-spec.md) | Optional | When `type`==`custom`, optional, if it doesn't exist, http and https is assumed |
| `SsrRelaxedTcpStateEnforcement` | `*bool` | Optional | **Default**: `false` |
| `TrafficClass` | [`*models.ServiceTrafficClassEnum`](../../doc/models/service-traffic-class-enum.md) | Optional | when `traffic_type`==`custom`. enum: `best_effort`, `high`, `low`, `medium`<br>**Default**: `"best_effort"` |
| `TrafficType` | `*string` | Optional | values from `/api/v1/consts/traffic_types`<br>**Default**: `"data_best_effort"` |
| `Type` | [`*models.ServiceTypeEnum`](../../doc/models/service-type-enum.md) | Optional | enum: `app_categories`, `apps`, `custom`, `urls`<br>**Default**: `"custom"` |
| `Urls` | `[]string` | Optional | When `type`==`urls`, no need for spec as URL can encode the ports being used |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

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
  "client_limit_down": 300000,
  "client_limit_up": 300000,
  "failover_policy": "revertable",
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "service_limit_down": 300000,
  "service_limit_up": 300000,
  "sle_enabled": false,
  "ssr_relaxed_tcp_state_enforcement": false,
  "traffic_class": "best_effort",
  "traffic_type": "data_best_effort",
  "type": "custom",
  "addresses": [
    "addresses6",
    "addresses7",
    "addresses8"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

