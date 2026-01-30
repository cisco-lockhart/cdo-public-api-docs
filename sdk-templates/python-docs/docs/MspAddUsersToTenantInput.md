# MspAddUsersToTenantInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**users** | [**List[UserInput]**](UserInput.md) | The list of users to be added to the tenant. You can add a maximum of 50 users at a time. | 

## Example

```python
from scc_firewall_manager_sdk.models.msp_add_users_to_tenant_input import MspAddUsersToTenantInput

# TODO update the JSON string below
json = "{}"
# create an instance of MspAddUsersToTenantInput from a JSON string
msp_add_users_to_tenant_input_instance = MspAddUsersToTenantInput.from_json(json)
# print the JSON string representation of the object
print(MspAddUsersToTenantInput.to_json())

# convert the object into a dict
msp_add_users_to_tenant_input_dict = msp_add_users_to_tenant_input_instance.to_dict()
# create an instance of MspAddUsersToTenantInput from a dict
msp_add_users_to_tenant_input_form_dict = msp_add_users_to_tenant_input.from_dict(msp_add_users_to_tenant_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


