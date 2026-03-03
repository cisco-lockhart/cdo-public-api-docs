# MspManagedTenantVirtualAccountDto


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**compliance_status** | **str** | The compliance status of the virtual account. | [optional] 
**name** | **str** | The name of the virtual account. | [optional] 
**smart_account_uid** | **str** | The unique identifier, represented as a UUID, of the smart account associated with this virtual account. | [optional] 
**uid** | **str** | The unique identifier, represented as a UUID, of the virtual account in Security Cloud Control. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.msp_managed_tenant_virtual_account_dto import MspManagedTenantVirtualAccountDto

# TODO update the JSON string below
json = "{}"
# create an instance of MspManagedTenantVirtualAccountDto from a JSON string
msp_managed_tenant_virtual_account_dto_instance = MspManagedTenantVirtualAccountDto.from_json(json)
# print the JSON string representation of the object
print(MspManagedTenantVirtualAccountDto.to_json())

# convert the object into a dict
msp_managed_tenant_virtual_account_dto_dict = msp_managed_tenant_virtual_account_dto_instance.to_dict()
# create an instance of MspManagedTenantVirtualAccountDto from a dict
msp_managed_tenant_virtual_account_dto_form_dict = msp_managed_tenant_virtual_account_dto.from_dict(msp_managed_tenant_virtual_account_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


