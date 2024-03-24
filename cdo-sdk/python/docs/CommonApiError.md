# CommonApiError


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**error_code** | **str** | A unique error code that describes the error. | [optional] 
**error_msg** | **str** | A human-readable error description in English. | [optional] 
**details** | **Dict[str, object]** | Additional details, if any, on the error | [optional] 

## Example

```python
from cdo_sdk_python.models.common_api_error import CommonApiError

# TODO update the JSON string below
json = "{}"
# create an instance of CommonApiError from a JSON string
common_api_error_instance = CommonApiError.from_json(json)
# print the JSON string representation of the object
print(CommonApiError.to_json())

# convert the object into a dict
common_api_error_dict = common_api_error_instance.to_dict()
# create an instance of CommonApiError from a dict
common_api_error_form_dict = common_api_error.from_dict(common_api_error_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


