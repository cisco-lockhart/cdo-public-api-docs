# VirtualAccountDto


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**compliance_status** | **str** | Indicates whether any of the licenses used in this tenant from this Virtual Account are out of compliance. | [optional] 
**last_updated_time** | **datetime** | The time, in RFC 3339 format, when information on this Virtual Account was last updated | [optional] 
**licenses_in_compliance_count** | **int** | Number of licenses in this Virtual Account with in compliance status | [optional] 
**licenses_out_of_compliance_count** | **int** | Number of licenses in this Virtual Account with out of compliance status | [optional] 
**name** | **str** | Name of the Virtual Account. | [optional] 
**uid** | **str** | Unique identifier, represented as a UUID, for your Virtual Account | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.virtual_account_dto import VirtualAccountDto

# TODO update the JSON string below
json = "{}"
# create an instance of VirtualAccountDto from a JSON string
virtual_account_dto_instance = VirtualAccountDto.from_json(json)
# print the JSON string representation of the object
print(VirtualAccountDto.to_json())

# convert the object into a dict
virtual_account_dto_dict = virtual_account_dto_instance.to_dict()
# create an instance of VirtualAccountDto from a dict
virtual_account_dto_form_dict = virtual_account_dto.from_dict(virtual_account_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


