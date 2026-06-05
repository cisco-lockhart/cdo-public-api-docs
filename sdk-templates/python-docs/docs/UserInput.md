# UserInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**api_only_user** | **bool** | Whether the user is an API-only user | [optional] 
**first_name** | **str** | The first name of the user in Security Cloud Control. If this is not set, the first name is set to the username. Note: This field must not be specified for API-only users. | [optional] 
**last_name** | **str** | The last name of the user in Security Cloud Control. If this is not set, the last name is set to the username. Note: This field must not be specified for API-only users. | [optional] 
**role** | [**UserRole**](UserRole.md) |  | 
**username** | **str** | The name of the user in Security Cloud Control. This must be a valid e-mail address if the user is not an API-only user. | 

## Example

```python
from scc_firewall_manager_sdk.models.user_input import UserInput

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


