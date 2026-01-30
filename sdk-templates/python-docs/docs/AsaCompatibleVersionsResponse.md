# AsaCompatibleVersionsResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**count** | **int** | The item count. | [optional] 
**items** | [**List[AsaCompatibleVersion]**](AsaCompatibleVersion.md) | The list of items retrieved. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.asa_compatible_versions_response import AsaCompatibleVersionsResponse

# TODO update the JSON string below
json = "{}"
# create an instance of AsaCompatibleVersionsResponse from a JSON string
asa_compatible_versions_response_instance = AsaCompatibleVersionsResponse.from_json(json)
# print the JSON string representation of the object
print(AsaCompatibleVersionsResponse.to_json())

# convert the object into a dict
asa_compatible_versions_response_dict = asa_compatible_versions_response_instance.to_dict()
# create an instance of AsaCompatibleVersionsResponse from a dict
asa_compatible_versions_response_form_dict = asa_compatible_versions_response.from_dict(asa_compatible_versions_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


