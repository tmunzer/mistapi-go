
# Remote Syslog Archive Files

Number of archived syslog files to retain

## Class Name

`RemoteSyslogArchiveFiles`

## Cases

| Type | Factory Method |
|  --- | --- |
| `string` | models.RemoteSyslogArchiveFilesContainer.FromString(string mString) |
| `int` | models.RemoteSyslogArchiveFilesContainer.FromNumber(int number) |

## string

### Initialization Code

#### Example

```go
value := models.RemoteSyslogArchiveFilesContainer.FromString("String0")
```

## int

### Initialization Code

#### Example

```go
value := models.RemoteSyslogArchiveFilesContainer.FromNumber(0)
```

