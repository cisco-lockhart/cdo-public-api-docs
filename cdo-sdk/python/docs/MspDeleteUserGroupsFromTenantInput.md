# MspDeleteUserGroupsFromTenantInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**user_group_uids** | **List[str]** | The list of unique user group identifiers, represented as UUIDs, to be removed from the MSP-managed tenant. You can remove a maximum of 50 user groups at a time. | 

## Example

```python
from cdo_sdk_python.models.msp_delete_user_groups_from_tenant_input import MspDeleteUserGroupsFromTenantInput

# TODO update the JSON string below
json = "{}"
# create an instance of MspDeleteUserGroupsFromTenantInput from a JSON string
msp_delete_user_groups_from_tenant_input_instance = MspDeleteUserGroupsFromTenantInput.from_json(json)
# print the JSON string representation of the object
print(MspDeleteUserGroupsFromTenantInput.to_json())

# convert the object into a dict
msp_delete_user_groups_from_tenant_input_dict = msp_delete_user_groups_from_tenant_input_instance.to_dict()
# create an instance of MspDeleteUserGroupsFromTenantInput from a dict
msp_delete_user_groups_from_tenant_input_form_dict = msp_delete_user_groups_from_tenant_input.from_dict(msp_delete_user_groups_from_tenant_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


