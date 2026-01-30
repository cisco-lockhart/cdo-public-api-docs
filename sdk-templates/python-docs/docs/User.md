# User


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**api_only_user** | **bool** | Whether the user is API-only, an API-only user cannot access Security Cloud Control in the UI. | [optional] 
**email_address** | **str** | The e-mail address in Security Cloud Control. | [optional] 
**first_name** | **str** | The first name of the user in Security Cloud Control. | [optional] 
**last_name** | **str** | The last name of the user in Security Cloud Control. | [optional] 
**last_successful_login** | **datetime** | The time (UTC; represented using the RFC-3339 standard) that indicate the last time the user successfully logged in to Security Cloud Control. | [optional] 
**name** | **str** | The username (e-mail address) in Security Cloud Control. | [optional] 
**roles** | [**List[UserRole]**](UserRole.md) | Roles associated with this user in Security Cloud Control. | [optional] 
**uid** | **str** | The unique identifier of the user in Security Cloud Control. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.user import User

# TODO update the JSON string below
json = "{}"
# create an instance of User from a JSON string
user_instance = User.from_json(json)
# print the JSON string representation of the object
print(User.to_json())

# convert the object into a dict
user_dict = user_instance.to_dict()
# create an instance of User from a dict
user_form_dict = user.from_dict(user_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


