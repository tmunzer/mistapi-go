
# Alarm Template

Alarm template defining default delivery and per-alarm rules

*This model accepts additional fields of type interface{}.*

## Structure

`AlarmTemplate`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `Delivery` | [`models.Delivery`](../../doc/models/delivery.md) | Required | Delivery object to configure the alarm delivery |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `Name` | `*string` | Optional | Some string to name the alarm template |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `Rules` | [`map[string]models.AlarmTemplateRule`](../../doc/models/alarm-template-rule.md) | Required | Alarm Rules object to configure the individual alarm keys/types. Property key is the alarm name. |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    alarmTemplate := models.AlarmTemplate{
        CreatedTime:          models.ToPointer(float64(87.26)),
        Delivery:             models.Delivery{
            AdditionalEmails:     []string{
                "additional_emails9",
                "additional_emails0",
                "additional_emails1",
            },
            Enabled:              true,
            ToOrgAdmins:          models.ToPointer(true),
            ToSiteAdmins:         models.ToPointer(false),
        },
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        ModifiedTime:         models.ToPointer(float64(247.7)),
        Name:                 models.ToPointer("default"),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        Rules:                map[string]models.AlarmTemplateRule{
            "ap_offline": models.AlarmTemplateRule{
                Delivery:             models.ToPointer(models.Delivery{
                    AdditionalEmails:     []string{
                        "string",
                    },
                    Enabled:              true,
                    ToOrgAdmins:          models.ToPointer(true),
                    ToSiteAdmins:         models.ToPointer(true),
                }),
                Enabled:              models.ToPointer(true),
            },
            "bad_cable": models.AlarmTemplateRule{
                Delivery:             models.ToPointer(models.Delivery{
                    AdditionalEmails:     []string{
                        "string",
                    },
                    Enabled:              true,
                    ToOrgAdmins:          models.ToPointer(true),
                    ToSiteAdmins:         models.ToPointer(true),
                }),
                Enabled:              models.ToPointer(true),
            },
        },
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

