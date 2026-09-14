# MslaEntitlementDto

The list of items retrieved.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**entitlement** | **str** | The unique identifier of the MSLA entitlement available on the tenant. | [optional] 
**uid** | **str** | The unique identifier, represented as a UUID, of the entitlement assignment row. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.msla_entitlement_dto import MslaEntitlementDto

# TODO update the JSON string below
json = "{}"
# create an instance of MslaEntitlementDto from a JSON string
msla_entitlement_dto_instance = MslaEntitlementDto.from_json(json)
# print the JSON string representation of the object
print(MslaEntitlementDto.to_json())

# convert the object into a dict
msla_entitlement_dto_dict = msla_entitlement_dto_instance.to_dict()
# create an instance of MslaEntitlementDto from a dict
msla_entitlement_dto_form_dict = msla_entitlement_dto.from_dict(msla_entitlement_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


