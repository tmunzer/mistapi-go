
# Service Policy

Site-level service policy that allows or denies traffic for tenants and services

## Structure

`ServicePolicy`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Action` | [`*models.AllowDenyEnum`](../../doc/models/allow-deny-enum.md) | Optional | Policy action value that either allows or denies matching traffic. enum: `allow`, `deny` |
| `Antivirus` | [`*models.ServicePolicyAntivirus`](../../doc/models/service-policy-antivirus.md) | Optional | SRX antivirus inspection settings for a service policy |
| `Appqoe` | [`*models.ServicePolicyAppqoe`](../../doc/models/service-policy-appqoe.md) | Optional | SRX application QoE settings for a service policy |
| `Ewf` | [`[]models.ServicePolicyEwfRule`](../../doc/models/service-policy-ewf-rule.md) | Optional | Enhanced web filtering rules applied by a service policy |
| `Idp` | [`*models.IdpConfig`](../../doc/models/idp-config.md) | Optional | Intrusion detection and prevention settings for a service policy |
| `LocalRouting` | `*bool` | Optional | Whether the policy permits access within the same VRF |
| `Name` | `*string` | Optional | Display name of the service policy |
| `PathPreference` | `*string` | Optional | By default, we derive all paths available and use them. Optionally, you can customize by using `path_preference` |
| `Secintel` | [`*models.ServicePolicySecintel`](../../doc/models/service-policy-secintel.md) | Optional | SRX SecIntel settings for a service policy |
| `ServicepolicyId` | `*uuid.UUID` | Optional | Organization-level service policy identifier used to link and override selected attributes |
| `Services` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `Skyatp` | [`*models.ServicePolicySkyatp`](../../doc/models/service-policy-skyatp.md) | Optional | SRX Sky ATP threat inspection settings for a service policy |
| `SslProxy` | [`*models.ServicePolicySslProxy`](../../doc/models/service-policy-ssl-proxy.md) | Optional | SRX SSL proxy inspection settings for a service policy |
| `Syslog` | [`*models.ServicePolicySyslog`](../../doc/models/service-policy-syslog.md) | Optional | Syslog logging settings for a service policy |
| `Tenants` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    servicePolicy := models.ServicePolicy{
        Action:               models.ToPointer(models.AllowDenyEnum_ALLOW),
        Antivirus:            models.ToPointer(models.ServicePolicyAntivirus{
            AvprofileId:          models.ToPointer(uuid.MustParse("00000282-0000-0000-0000-000000000000")),
            Enabled:              models.ToPointer(false),
            Profile:              models.ToPointer("profile4"),
        }),
        Appqoe:               models.ToPointer(models.ServicePolicyAppqoe{
            Enabled:              models.ToPointer(false),
        }),
        Ewf:                  []models.ServicePolicyEwfRule{
            models.ServicePolicyEwfRule{
                AlertOnly:            models.ToPointer(false),
                BlockMessage:         models.ToPointer("block_message0"),
                Enabled:              models.ToPointer(false),
                Profile:              models.ToPointer(models.ServicePolicyEwfRuleProfileEnum_STANDARD),
            },
            models.ServicePolicyEwfRule{
                AlertOnly:            models.ToPointer(false),
                BlockMessage:         models.ToPointer("block_message0"),
                Enabled:              models.ToPointer(false),
                Profile:              models.ToPointer(models.ServicePolicyEwfRuleProfileEnum_STANDARD),
            },
            models.ServicePolicyEwfRule{
                AlertOnly:            models.ToPointer(false),
                BlockMessage:         models.ToPointer("block_message0"),
                Enabled:              models.ToPointer(false),
                Profile:              models.ToPointer(models.ServicePolicyEwfRuleProfileEnum_STANDARD),
            },
        },
        Idp:                  models.ToPointer(models.IdpConfig{
            AlertOnly:            models.ToPointer(false),
            Enabled:              models.ToPointer(false),
            IdpprofileId:         models.ToPointer(uuid.MustParse("00000e94-0000-0000-0000-000000000000")),
            Profile:              models.ToPointer("profile8"),
        }),
    }

}
```

