# HttpStatusCode


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**error** | **bool** |  | [optional] 
**is4xx_client_error** | **bool** |  | [optional] 
**is5xx_server_error** | **bool** |  | [optional] 
**is1xx_informational** | **bool** |  | [optional] 
**is2xx_successful** | **bool** |  | [optional] 
**is3xx_redirection** | **bool** |  | [optional] 

## Example

```python
from cdo_sdk_python.models.http_status_code import HttpStatusCode

# TODO update the JSON string below
json = "{}"
# create an instance of HttpStatusCode from a JSON string
http_status_code_instance = HttpStatusCode.from_json(json)
# print the JSON string representation of the object
print(HttpStatusCode.to_json())

# convert the object into a dict
http_status_code_dict = http_status_code_instance.to_dict()
# create an instance of HttpStatusCode from a dict
http_status_code_form_dict = http_status_code.from_dict(http_status_code_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


