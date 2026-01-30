# FtdVersionsResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**count** | **int** | The item count. | [optional] 
**items** | [**List[FtdVersion]**](FtdVersion.md) | The list of items retrieved. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.ftd_versions_response import FtdVersionsResponse

# TODO update the JSON string below
json = "{}"
# create an instance of FtdVersionsResponse from a JSON string
ftd_versions_response_instance = FtdVersionsResponse.from_json(json)
# print the JSON string representation of the object
print(FtdVersionsResponse.to_json())

# convert the object into a dict
ftd_versions_response_dict = ftd_versions_response_instance.to_dict()
# create an instance of FtdVersionsResponse from a dict
ftd_versions_response_form_dict = ftd_versions_response.from_dict(ftd_versions_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


