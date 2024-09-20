# CdoTokenInfo


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**uid** | **str** | The unique identifier, represented as a UUID, of the user this token belongs to. The user can be API-only or a human. | [optional] 
**name** | **str** | The name of the user this token belongs to. The user can be API-only or a human. | [optional] 
**roles** | [**List[UserRole]**](UserRole.md) | The role of the user this token belongs to. The user can be API-only or a human. | [optional] 
**expires_at** | **datetime** | The time (UTC; represented using the RFC-3339 standard) the token expires. If this field is missing, the token will never expire. | [optional] 
**tenant_uid** | **str** | The unique identifier, represented as a UUID, of the tenant this token is associated with. While a user can be associated with multiple tenants, a token is associated with a single tenant. | [optional] 

## Example

```python
from cdo_sdk_python.models.cdo_token_info import CdoTokenInfo

# TODO update the JSON string below
json = "{}"
# create an instance of CdoTokenInfo from a JSON string
cdo_token_info_instance = CdoTokenInfo.from_json(json)
# print the JSON string representation of the object
print(CdoTokenInfo.to_json())

# convert the object into a dict
cdo_token_info_dict = cdo_token_info_instance.to_dict()
# create an instance of CdoTokenInfo from a dict
cdo_token_info_form_dict = cdo_token_info.from_dict(cdo_token_info_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


