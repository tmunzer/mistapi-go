
# Account Oauth Add

Account configuration payload for adding a supported OAuth-linked application

## Class Name

`AccountOauthAdd`

## Cases

| Type | Factory Method |
|  --- | --- |
| [`models.AccountJamfConfig`](../../../doc/models/account-jamf-config.md) | models.AccountOauthAddContainer.FromAccountJamfConfig(models.AccountJamfConfig accountJamfConfig) |
| [`models.AccountVmwareConfig`](../../../doc/models/account-vmware-config.md) | models.AccountOauthAddContainer.FromAccountVmwareConfig(models.AccountVmwareConfig accountVmwareConfig) |
| [`models.AccountMobicontrolConfig`](../../../doc/models/account-mobicontrol-config.md) | models.AccountOauthAddContainer.FromAccountMobicontrolConfig(models.AccountMobicontrolConfig accountMobicontrolConfig) |
| [`models.AccountZdxConfig`](../../../doc/models/account-zdx-config.md) | models.AccountOauthAddContainer.FromAccountZdxConfig(models.AccountZdxConfig accountZdxConfig) |
| [`models.AccountCrowdstrikeConfig`](../../../doc/models/account-crowdstrike-config.md) | models.AccountOauthAddContainer.FromAccountCrowdstrikeConfig(models.AccountCrowdstrikeConfig accountCrowdstrikeConfig) |
| [`models.AccountPrismaConfig`](../../../doc/models/account-prisma-config.md) | models.AccountOauthAddContainer.FromAccountPrismaConfig(models.AccountPrismaConfig accountPrismaConfig) |
| [`models.AccountSentineloneConfig`](../../../doc/models/account-sentinelone-config.md) | models.AccountOauthAddContainer.FromAccountSentineloneConfig(models.AccountSentineloneConfig accountSentineloneConfig) |

## models.AccountJamfConfig

### Initialization Code

#### Example

```go
value := models.AccountOauthAddContainer.FromAccountJamfConfig(models.AccountJamfConfig{
    ClientId:             "client_id0",
    ClientSecret:         "client_secret6",
    InstanceUrl:          "junipertest.jamfcloud.com",
    SmartgroupName:       "CompliantGroup1",
})
```

## models.AccountVmwareConfig

### Initialization Code

#### Example

```go
value := models.AccountOauthAddContainer.FromAccountVmwareConfig(models.AccountVmwareConfig{
    ClientId:             "client_id8",
    ClientSecret:         "client_secret4",
    InstanceUrl:          "instance_url4",
    WebhookEnabled:       false,
})
```

## models.AccountMobicontrolConfig

### Initialization Code

#### Example

```go
value := models.AccountOauthAddContainer.FromAccountMobicontrolConfig(models.AccountMobicontrolConfig{
    ClientId:             "client_id0",
    ClientSecret:         "client_secret6",
    InstanceUrl:          "instance_url2",
    Password:             "password2",
    Username:             "username2",
})
```

## models.AccountZdxConfig

### Initialization Code

#### Example

```go
value := models.AccountOauthAddContainer.FromAccountZdxConfig(models.AccountZdxConfig{
    CloudName:            models.ToPointer("zdxcloud.net"),
    KeyId:                "K35vrZcK3JvrZc",
    KeySecret:            "K35vrZcK3JvrZcjjswpp2eii1oo100",
    ZdxOrgId:             "123456",
})
```

## models.AccountCrowdstrikeConfig

### Initialization Code

#### Example

```go
value := models.AccountOauthAddContainer.FromAccountCrowdstrikeConfig(models.AccountCrowdstrikeConfig{
    ClientId:             "client_id4",
    ClientSecret:         "client_secret0",
    CustomerId:           "customer_id0",
})
```

## models.AccountPrismaConfig

### Initialization Code

#### Example

```go
value := models.AccountOauthAddContainer.FromAccountPrismaConfig(models.AccountPrismaConfig{
    AutoProbeSubnet:      models.ToPointer("11.0.0.0/8"),
    ClientId:             "client_id2",
    ClientSecret:         "client_secret8",
    EnableProbe:          models.ToPointer(false),
    TsgId:                "tsg_id6",
})
```

## models.AccountSentineloneConfig

### Initialization Code

#### Example

```go
value := models.AccountOauthAddContainer.FromAccountSentineloneConfig(models.AccountSentineloneConfig{
    ApiToken:             "api_token6",
    InstanceUrl:          "instance_url4",
})
```

