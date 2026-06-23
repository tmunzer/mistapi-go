
# Remote Syslog Server Port

Syslog Service Port, value from 1 to 65535

## Class Name

`RemoteSyslogServerPort`

## Cases

| Type | Factory Method |
|  --- | --- |
| `int` | models.RemoteSyslogServerPortContainer.FromNumber(int number) |
| `string` | models.RemoteSyslogServerPortContainer.FromString(string mString) |

## int

### Initialization Code

#### Example

```go
value := models.RemoteSyslogServerPortContainer.FromNumber(514)
```

## string

### Initialization Code

#### Example

```go
value := models.RemoteSyslogServerPortContainer.FromString("String0")
```

