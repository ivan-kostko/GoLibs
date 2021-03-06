# HttpService
--
#### Install

    go get "github.com/ivan-kostko/GoLibs/Repository/DataSource/Implementations/HttpService"

#### Import

    "github.com/ivan-kostko/GoLibs/Repository/DataSource/Implementations"
    _ "github.com/ivan-kostko/GoLibs/Repository/DataSource/Implementations/HttpService"

#### Description

    The package represents generic Http service data source

### Usage

```go
const (
	ERR_IVALIDCONFTYPE    = "HttpService: Invalid configuration type"
	ERR_WONTBUILDREQUEST  = "HttpService: Wont build http.Request due to internal error: "
	ERR_WONTINVOKEREQUEST = "HttpService: Wont invoke the Request due to internal error: "
	ERR_INVALIDRESPONSE   = "HttpService: Returned invalid response: "
	ERR_WONTREADRESPONSE  = "HttpService: Wont read the Response due to internal error: "
)
```

##### func  GetNewHttpServiceDataSource

```go
func GetNewHttpServiceDataSource(conf interface{}) (*ds.DataSource, *Error)
```
Generic HTTP service data source factory NB!: conf should be of type
HttpService.Configuration. Otherwise it returns InvalidArgument,
ERR_IVALIDCONFTYPE

##### type Configuration

```go
type Configuration struct {
	// The url for the service which will be followed by slash "/" and parametrised instructions
	MainUrl string

	Headers map[string][]string
	// Checks response for http errors
	CheckResponse func(*http.Response) *Error
}
```

The type represents configuration for HTTP service data source
