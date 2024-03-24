# ApiTokenInfo


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**api_token** | **str** | The API Access Token. | [optional] 

## Example

```python
from cdo_sdk_python.models.api_token_info import ApiTokenInfo

# TODO update the JSON string below
json = "{}"
# create an instance of ApiTokenInfo from a JSON string
api_token_info_instance = ApiTokenInfo.from_json(json)
# print the JSON string representation of the object
print(ApiTokenInfo.to_json())

# convert the object into a dict
api_token_info_dict = api_token_info_instance.to_dict()
# create an instance of ApiTokenInfo from a dict
api_token_info_form_dict = api_token_info.from_dict(api_token_info_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


