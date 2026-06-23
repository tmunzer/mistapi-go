
# Snmpv 3 Config

SNMPv3 notification, target, USM, and VACM configuration

## Structure

`Snmpv3Config`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Notify` | [`[]models.Snmpv3ConfigNotifyItems`](../../doc/models/snmpv-3-config-notify-items.md) | Optional | SNMPv3 notification definitions used for traps and informs |
| `NotifyFilter` | [`[]models.Snmpv3ConfigNotifyFilterItem`](../../doc/models/snmpv-3-config-notify-filter-item.md) | Optional | SNMPv3 notification filter profiles |
| `TargetAddress` | [`[]models.Snmpv3ConfigTargetAddressItem`](../../doc/models/snmpv-3-config-target-address-item.md) | Optional | SNMPv3 notification target addresses |
| `TargetParameters` | [`[]models.Snmpv3ConfigTargetParam`](../../doc/models/snmpv-3-config-target-param.md) | Optional | SNMPv3 target parameter profiles |
| `Usm` | [`[]models.SnmpUsm`](../../doc/models/snmp-usm.md) | Optional | SNMPv3 USM engine configurations |
| `Vacm` | [`*models.SnmpVacm`](../../doc/models/snmp-vacm.md) | Optional | SNMPv3 View-based Access Control Model configuration |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    snmpv3Config := models.Snmpv3Config{
        Notify:               []models.Snmpv3ConfigNotifyItems{
            models.Snmpv3ConfigNotifyItems{
                Name:                 models.ToPointer("name0"),
                Tag:                  models.ToPointer("tag4"),
                Type:                 models.ToPointer(models.Snmpv3ConfigNotifyTypeEnum_INFORM),
            },
            models.Snmpv3ConfigNotifyItems{
                Name:                 models.ToPointer("name0"),
                Tag:                  models.ToPointer("tag4"),
                Type:                 models.ToPointer(models.Snmpv3ConfigNotifyTypeEnum_INFORM),
            },
        },
        NotifyFilter:         []models.Snmpv3ConfigNotifyFilterItem{
            models.Snmpv3ConfigNotifyFilterItem{
                Contents:             []models.Snmpv3ConfigNotifyFilterItemContent{
                    models.Snmpv3ConfigNotifyFilterItemContent{
                        Include:              models.ToPointer(false),
                        Oid:                  models.ToPointer("oid4"),
                    },
                },
                ProfileName:          models.ToPointer("profile_name2"),
            },
        },
        TargetAddress:        []models.Snmpv3ConfigTargetAddressItem{
            models.Snmpv3ConfigTargetAddressItem{
                Address:              models.ToPointer("address8"),
                AddressMask:          models.ToPointer("address_mask0"),
                Port:                 models.NewOptional(models.ToPointer("port2")),
                TagList:              models.ToPointer("tag_list4"),
                TargetAddressName:    models.ToPointer("target_address_name6"),
            },
            models.Snmpv3ConfigTargetAddressItem{
                Address:              models.ToPointer("address8"),
                AddressMask:          models.ToPointer("address_mask0"),
                Port:                 models.NewOptional(models.ToPointer("port2")),
                TagList:              models.ToPointer("tag_list4"),
                TargetAddressName:    models.ToPointer("target_address_name6"),
            },
        },
        TargetParameters:     []models.Snmpv3ConfigTargetParam{
            models.Snmpv3ConfigTargetParam{
                MessageProcessingModel: models.ToPointer(models.Snmpv3ConfigTargetParamMessProcessModelEnum_V3),
                Name:                   models.ToPointer("name4"),
                NotifyFilter:           models.ToPointer("notify_filter6"),
                SecurityLevel:          models.ToPointer(models.Snmpv3ConfigTargetParamSecurityLevelEnum_NONE),
                SecurityModel:          models.ToPointer(models.Snmpv3ConfigTargetParamSecurityModelEnum_USM),
            },
            models.Snmpv3ConfigTargetParam{
                MessageProcessingModel: models.ToPointer(models.Snmpv3ConfigTargetParamMessProcessModelEnum_V3),
                Name:                   models.ToPointer("name4"),
                NotifyFilter:           models.ToPointer("notify_filter6"),
                SecurityLevel:          models.ToPointer(models.Snmpv3ConfigTargetParamSecurityLevelEnum_NONE),
                SecurityModel:          models.ToPointer(models.Snmpv3ConfigTargetParamSecurityModelEnum_USM),
            },
        },
        Usm:                  []models.SnmpUsm{
            models.SnmpUsm{
                EngineType:           models.ToPointer(models.SnmpUsmEngineTypeEnum_LOCALENGINE),
                RemoteEngineId:       models.ToPointer("remote_engine_id8"),
                Users:                []models.SnmpUsmUser{
                    models.SnmpUsmUser{
                        AuthenticationPassword: models.ToPointer("authentication_password0"),
                        AuthenticationType:     models.ToPointer(models.SnmpUsmUserAuthenticationTypeEnum_AUTHENTICATIONSHA384),
                        EncryptionPassword:     models.ToPointer("encryption_password4"),
                        EncryptionType:         models.ToPointer(models.SnmpUsmUserEncryptionTypeEnum_PRIVACY3DES),
                        Name:                   models.ToPointer("name6"),
                    },
                },
            },
            models.SnmpUsm{
                EngineType:           models.ToPointer(models.SnmpUsmEngineTypeEnum_LOCALENGINE),
                RemoteEngineId:       models.ToPointer("remote_engine_id8"),
                Users:                []models.SnmpUsmUser{
                    models.SnmpUsmUser{
                        AuthenticationPassword: models.ToPointer("authentication_password0"),
                        AuthenticationType:     models.ToPointer(models.SnmpUsmUserAuthenticationTypeEnum_AUTHENTICATIONSHA384),
                        EncryptionPassword:     models.ToPointer("encryption_password4"),
                        EncryptionType:         models.ToPointer(models.SnmpUsmUserEncryptionTypeEnum_PRIVACY3DES),
                        Name:                   models.ToPointer("name6"),
                    },
                },
            },
            models.SnmpUsm{
                EngineType:           models.ToPointer(models.SnmpUsmEngineTypeEnum_LOCALENGINE),
                RemoteEngineId:       models.ToPointer("remote_engine_id8"),
                Users:                []models.SnmpUsmUser{
                    models.SnmpUsmUser{
                        AuthenticationPassword: models.ToPointer("authentication_password0"),
                        AuthenticationType:     models.ToPointer(models.SnmpUsmUserAuthenticationTypeEnum_AUTHENTICATIONSHA384),
                        EncryptionPassword:     models.ToPointer("encryption_password4"),
                        EncryptionType:         models.ToPointer(models.SnmpUsmUserEncryptionTypeEnum_PRIVACY3DES),
                        Name:                   models.ToPointer("name6"),
                    },
                },
            },
        },
    }

}
```

