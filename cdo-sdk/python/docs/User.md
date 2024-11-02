# User


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**uid** | **str** | The unique identifier, represented as a UUID, of the user in SCC. | [optional] 
**name** | **str** | The name of the user in SCC. | [optional] 
**roles** | [**List[UserRole]**](UserRole.md) | Roles associated with this user in SCC. | [optional] 
**api_only_user** | **bool** | Whether the user is API-only, an API-only user cannot access SCC in the UI. | [optional] 
**last_successful_login** | **datetime** | The time (UTC; represented using the RFC-3339 standard) that indicate the last time the user successfully logged in to SCC. | [optional] 

## Example

```python
from cdo_sdk_python.models.user import User

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


