# UserCreateOrUpdateInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**api_only_user** | **bool** | Whether the user is API-only, an API-only user cannot access Security Cloud Control in the UI. | [optional] [default to False]
**first_name** | **str** | The first name of the user in Security Cloud Control. Note: This field must not be specified for API-only users. | [optional] 
**last_name** | **str** | The last name of the user in Security Cloud Control. Note: This field must not be specified for API-only users. | [optional] 
**name** | **str** | The name of the user in Security Cloud Control. This has to be a valid email if creating a human user, and must not be an email if creating an API-only user. | 
**role** | **str** | The user role in Security Cloud Control. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.user_create_or_update_input import UserCreateOrUpdateInput

# TODO update the JSON string below
json = "{}"
# create an instance of UserCreateOrUpdateInput from a JSON string
user_create_or_update_input_instance = UserCreateOrUpdateInput.from_json(json)
# print the JSON string representation of the object
print(UserCreateOrUpdateInput.to_json())

# convert the object into a dict
user_create_or_update_input_dict = user_create_or_update_input_instance.to_dict()
# create an instance of UserCreateOrUpdateInput from a dict
user_create_or_update_input_form_dict = user_create_or_update_input.from_dict(user_create_or_update_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


