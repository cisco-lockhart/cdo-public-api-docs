# AuthenticationError


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**error** | **str** | A human-readable error description in English. | [optional] 
**error_description** | **str** | A human-readable error description in English. | [optional] 

## Example

```python
from cdo_sdk_python.models.authentication_error import AuthenticationError

# TODO update the JSON string below
json = "{}"
# create an instance of AuthenticationError from a JSON string
authentication_error_instance = AuthenticationError.from_json(json)
# print the JSON string representation of the object
print(AuthenticationError.to_json())

# convert the object into a dict
authentication_error_dict = authentication_error_instance.to_dict()
# create an instance of AuthenticationError from a dict
authentication_error_form_dict = authentication_error.from_dict(authentication_error_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


