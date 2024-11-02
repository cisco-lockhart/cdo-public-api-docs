# UserInput

The list of users to be added to the tenant. You can add a maximum of 50 users at a time.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**username** | **str** | The name of the user in SCC. This must be a valid e-mail address if the user is not an API-only user. | 
**role** | [**UserRole**](UserRole.md) |  | 
**api_only_user** | **bool** | Whether the user is an API-only user | [optional] 

## Example

```python
from cdo_sdk_python.models.user_input import UserInput

# TODO update the JSON string below
json = "{}"
# create an instance of UserInput from a JSON string
user_input_instance = UserInput.from_json(json)
# print the JSON string representation of the object
print(UserInput.to_json())

# convert the object into a dict
user_input_dict = user_input_instance.to_dict()
# create an instance of UserInput from a dict
user_input_form_dict = user_input.from_dict(user_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


