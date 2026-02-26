# SmartAccountDto


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**domain** | **str** | Account domain for your Smart Account | [optional] 
**last_updated_time** | **datetime** | The time, in RFC 3339 format, when information on this Smart Account was last updated | [optional] 
**smart_account_id** | **str** | Smart Account Identifier | [optional] 
**uid** | **str** | Unique identifier, represented as a UUID, for your Smart Account | [optional] 
**virtual_accounts_in_compliance_count** | **int** | Number of Virtual Accounts in this Smart Account with in compliance status | [optional] 
**virtual_accounts_out_of_compliance_count** | **int** | Number of Virtual Accounts in this Smart Account with out of compliance status | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.smart_account_dto import SmartAccountDto

# TODO update the JSON string below
json = "{}"
# create an instance of SmartAccountDto from a JSON string
smart_account_dto_instance = SmartAccountDto.from_json(json)
# print the JSON string representation of the object
print(SmartAccountDto.to_json())

# convert the object into a dict
smart_account_dto_dict = smart_account_dto_instance.to_dict()
# create an instance of SmartAccountDto from a dict
smart_account_dto_form_dict = smart_account_dto.from_dict(smart_account_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


