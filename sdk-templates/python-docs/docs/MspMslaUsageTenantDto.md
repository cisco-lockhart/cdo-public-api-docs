# MspMslaUsageTenantDto

A managed tenant with at least one device consuming an MSLA entitlement.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**display_name** | **str** | The human-readable display name of the managed tenant. | [optional] 
**name** | **str** | The name of the managed tenant (unique in CDO). | [optional] 
**region** | **str** | The CDO region of the managed tenant. | [optional] 
**uid** | **str** | The unique identifier, represented as a UUID, of the managed tenant in Security Cloud Control. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.msp_msla_usage_tenant_dto import MspMslaUsageTenantDto

# TODO update the JSON string below
json = "{}"
# create an instance of MspMslaUsageTenantDto from a JSON string
msp_msla_usage_tenant_dto_instance = MspMslaUsageTenantDto.from_json(json)
# print the JSON string representation of the object
print(MspMslaUsageTenantDto.to_json())

# convert the object into a dict
msp_msla_usage_tenant_dto_dict = msp_msla_usage_tenant_dto_instance.to_dict()
# create an instance of MspMslaUsageTenantDto from a dict
msp_msla_usage_tenant_dto_form_dict = msp_msla_usage_tenant_dto.from_dict(msp_msla_usage_tenant_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


