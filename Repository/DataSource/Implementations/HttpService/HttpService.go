//   Copyright (c) 2016 Ivan A Kostko (github.com/ivan-kostko)

//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at

//       http://www.apache.org/licenses/LICENSE-2.0

//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.


package HttpService

import(
    "io/ioutil"
	"net/http"
    ds "github.com/ivan-kostko/GoLibs/Repository/DataSource"
    . "github.com/ivan-kostko/GoLibs/CustomErrors"
)

const(
    ERR_IVALIDCONFTYPE = "HttpService: Invalid configuration type"
    ERR_WONTBUILDREQUEST = "HttpService: Wont build http.Request due to internal error: "
    ERR_WONTINVOKEREQUEST = "HttpService: Wont invoke the Request due to internal error: "
    ERR_INVALIDRESPONSE = "HttpService: Returned invalid response: "
    ERR_WONTREADRESPONSE = "HttpService: Wont read the Response due to internal error: "
)

// The type represents configuration for HTTP service data source
type Configuration struct {
    // The url for the service which will be followed by slash "/" and parametrised instructions
    MainUrl        string
    // The header(s) which will applied to all requests

    Headers        map[string][]string
    // Checks response for http errors
    CheckResponse  func(*http.Response) *Error
}



// Generic HTTP service data source factory
// NB!: conf should be of type HttpService.Configuration. Otherwise it returns InvalidArgument, ERR_IVALIDCONFTYPE
func GetNewHttpServiceDataSource(conf interface{}) (*ds.DataSource, *Error) {
    // Try to assert input configuration to custom Cunfiguration type
    c, ok := conf.(Configuration)
    if !ok {
        return nil, NewError(InvalidArgument, ERR_IVALIDCONFTYPE)
    }
    client := http.DefaultClient
    mainUrl:= c.MainUrl
    header := http.Header(c.Headers)//{"Accept":{"application/vnd.sdmx.data+json;version=1.0.0-wd"}}
    var checkResponse func(*http.Response) *Error
    if c.CheckResponse != nil {
        checkResponse = c.CheckResponse
    }else{
        checkResponse = defaultCheckResponse
    }

    executeInstruction := ds.ExecuteInstruction(func(i ds.Instruction) (ds.Result, *Error) {
            url := mainUrl+"/"+string(i)
            request, err := http.NewRequest("GET", url, nil)
            if err != nil {
                NewError(InvalidOperation, ERR_WONTBUILDREQUEST+err.Error())
            }
            request.Header = header
            res, err := client.Do(request)
        	if err != nil {
                return nil, NewError(InvalidOperation, ERR_WONTINVOKEREQUEST+err.Error())
        	}
            err = checkResponse(res)
            if err != nil {
                return nil, NewError(InvalidOperation, ERR_INVALIDRESPONSE+err.Error())
            }
            result, err := ioutil.ReadAll(res.Body)
        	res.Body.Close()
        	if err != nil {
                return nil, NewError(InvalidOperation, ERR_WONTREADRESPONSE+err.Error())
        	}
            return result, nil
    })

    return ds.GetNewDataSource(executeInstruction), nil

}

//
func defaultCheckResponse (res *http.Response) *Error {
    return nil
}

