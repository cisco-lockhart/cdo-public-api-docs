# MspVirtualAccountDto


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**compliance_status** | **str** | The compliance status of the virtual account. | [optional] 
**last_updated_time** | **datetime** | The time, in RFC 3339 format, when information on this Virtual Account was last updated | [optional] 
**licenses_in_compliance_count** | **int** | Number of licenses in this Virtual Account with in compliance status | [optional] 
**licenses_out_of_compliance_count** | **int** | Number of licenses in this Virtual Account with out of compliance status | [optional] 
**managed_tenants** | [**List[ManagedTenantDto]**](ManagedTenantDto.md) | The list of managed tenants associated with this Virtual Account. | [optional] 
**name** | **str** | The name of the virtual account. | [optional] 
**smart_account_uid** | **str** | The unique identifier, represented as a UUID, of the smart account associated with this virtual account. | [optional] 
**uid** | **str** | The unique identifier, represented as a UUID, of the virtual account in Security Cloud Control. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.msp_virtual_account_dto import MspVirtualAccountDto

# TODO update the JSON string below
json = "{}"
# create an instance of MspVirtualAccountDto from a JSON string
msp_virtual_account_dto_instance = MspVirtualAccountDto.from_json(json)
# print the JSON string representation of the object
print(MspVirtualAccountDto.to_json())

# convert the object into a dict
msp_virtual_account_dto_dict = msp_virtual_account_dto_instance.to_dict()
# create an instance of MspVirtualAccountDto from a dict
msp_virtual_account_dto_form_dict = msp_virtual_account_dto.from_dict(msp_virtual_account_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


