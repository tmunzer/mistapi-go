
# Org Service Policy

Organization-level service policy that allows or denies traffic for tenants and services

*This model accepts additional fields of type interface{}.*

## Structure

`OrgServicePolicy`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Aamw` | [`*models.ServicePolicyAamw`](../../doc/models/service-policy-aamw.md) | Optional | SRX advanced anti-malware settings for a service policy |
| `Action` | [`*models.AllowDenyEnum`](../../doc/models/allow-deny-enum.md) | Optional | Policy action value that either allows or denies matching traffic. enum: `allow`, `deny` |
| `Antivirus` | [`*models.ServicePolicyAntivirus`](../../doc/models/service-policy-antivirus.md) | Optional | SRX antivirus inspection settings for a service policy |
| `Appqoe` | [`*models.ServicePolicyAppqoe`](../../doc/models/service-policy-appqoe.md) | Optional | SRX application QoE settings for a service policy |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `Ewf` | [`[]models.ServicePolicyEwfRule`](../../doc/models/service-policy-ewf-rule.md) | Optional | Enhanced web filtering rules applied by a service policy |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `Idp` | [`*models.IdpConfig`](../../doc/models/idp-config.md) | Optional | Intrusion detection and prevention settings for a service policy |
| `LocalRouting` | `*bool` | Optional | Whether the policy permits access within the same VRF |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `Name` | `*string` | Optional | Display name of the service policy |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `PathPreference` | `*string` | Optional | By default, we derive all paths available and use them, optionally, you can customize by using `path_preference` |
| `Secintel` | [`*models.ServicePolicySecintel`](../../doc/models/service-policy-secintel.md) | Optional | SRX SecIntel settings for a service policy |
| `Services` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `SslProxy` | [`*models.ServicePolicySslProxy`](../../doc/models/service-policy-ssl-proxy.md) | Optional | SRX SSL proxy inspection settings for a service policy |
| `Tenants` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    orgServicePolicy := models.OrgServicePolicy{
        Aamw:                 models.ToPointer(models.ServicePolicyAamw{
            AamwprofileId:        models.ToPointer(uuid.MustParse("0000066a-0000-0000-0000-000000000000")),
            Enabled:              models.ToPointer(false),
            Profile:              models.ToPointer(models.ServicePolicyAamwProfileEnum_EXECUTABLES),
        }),
        Action:               models.ToPointer(models.AllowDenyEnum_ALLOW),
        Antivirus:            models.ToPointer(models.ServicePolicyAntivirus{
            AvprofileId:          models.ToPointer(uuid.MustParse("00000282-0000-0000-0000-000000000000")),
            Enabled:              models.ToPointer(false),
            Profile:              models.ToPointer("profile4"),
        }),
        Appqoe:               models.ToPointer(models.ServicePolicyAppqoe{
            Enabled:              models.ToPointer(false),
        }),
        CreatedTime:          models.ToPointer(float64(13.04)),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

