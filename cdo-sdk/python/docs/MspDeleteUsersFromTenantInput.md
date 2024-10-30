# MspDeleteUsersFromTenantInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**usernames** | **List[str]** | The list of usernames to be removed from the MSP-managed tenant. You can remove a maximum of 50 users at a time. | 

## Example

```python
from cdo_sdk_python.models.msp_delete_users_from_tenant_input import MspDeleteUsersFromTenantInput

# TODO update the JSON string below
json = "{}"
# create an instance of MspDeleteUsersFromTenantInput from a JSON string
msp_delete_users_from_tenant_input_instance = MspDeleteUsersFromTenantInput.from_json(json)
# print the JSON string representation of the object
print(MspDeleteUsersFromTenantInput.to_json())

# convert the object into a dict
msp_delete_users_from_tenant_input_dict = msp_delete_users_from_tenant_input_instance.to_dict()
# create an instance of MspDeleteUsersFromTenantInput from a dict
msp_delete_users_from_tenant_input_form_dict = msp_delete_users_from_tenant_input.from_dict(msp_delete_users_from_tenant_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


