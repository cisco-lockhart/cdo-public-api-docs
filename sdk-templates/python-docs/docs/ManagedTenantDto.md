# ManagedTenantDto


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**display_name** | **str** | The display name of the managed tenant. | [optional] 
**name** | **str** | The name of the managed tenant. | [optional] 
**uid** | **str** | The unique identifier of the managed tenant. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.managed_tenant_dto import ManagedTenantDto

# TODO update the JSON string below
json = "{}"
# create an instance of ManagedTenantDto from a JSON string
managed_tenant_dto_instance = ManagedTenantDto.from_json(json)
# print the JSON string representation of the object
print(ManagedTenantDto.to_json())

# convert the object into a dict
managed_tenant_dto_dict = managed_tenant_dto_instance.to_dict()
# create an instance of ManagedTenantDto from a dict
managed_tenant_dto_form_dict = managed_tenant_dto.from_dict(managed_tenant_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


