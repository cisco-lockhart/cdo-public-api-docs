# MspSmartAccountDto


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**domain** | **str** | The domain of the smart account. | [optional] 
**last_updated_time** | **datetime** | The time, in RFC 3339 format, when information on this Smart Account was last updated | [optional] 
**managed_tenants** | [**List[ManagedTenantDto]**](ManagedTenantDto.md) | The list of managed tenants associated with this Smart Account. | [optional] 
**smart_account_id** | **str** | The identifier of the smart account, represented as a String, returned by the Smart Account API. | [optional] 
**uid** | **str** | The unique identifier, represented as a UUID, of the smart account in Security Cloud Control. | [optional] 
**virtual_accounts_in_compliance_count** | **int** | Number of Virtual Accounts in this Smart Account with IN_COMPLIANCE status | [optional] 
**virtual_accounts_out_of_compliance_count** | **int** | Number of Virtual Accounts in this Smart Account with OUT_OF_COMPLIANCE status | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.msp_smart_account_dto import MspSmartAccountDto

# TODO update the JSON string below
json = "{}"
# create an instance of MspSmartAccountDto from a JSON string
msp_smart_account_dto_instance = MspSmartAccountDto.from_json(json)
# print the JSON string representation of the object
print(MspSmartAccountDto.to_json())

# convert the object into a dict
msp_smart_account_dto_dict = msp_smart_account_dto_instance.to_dict()
# create an instance of MspSmartAccountDto from a dict
msp_smart_account_dto_form_dict = msp_smart_account_dto.from_dict(msp_smart_account_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


