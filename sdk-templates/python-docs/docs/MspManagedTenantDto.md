# MspManagedTenantDto


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**api_token_valid** | **bool** | Indicates whether the tenant&#39;s API token is valid. | [optional] 
**cd_fmc_type** | **str** | Indicates whether the tenant&#39;s provisioned cdFMC instance is dedicated rather than shared. | [optional] 
**display_name** | **str** | A human-readable display name for the tenant. This is the tenant name displayed in the Security Cloud Control Web UI. | 
**name** | **str** | The name of the tenant in CDO. Tenant names are unique in Security Cloud Control. | 
**region** | **str** | The Security Cloud Control region the tenant exists in. | 
**scc_organization_uid** | **str** | The unique identifier of the organization in Security Cloud Control Firewall Platform. This is different to the unique identifier used in Security Cloud Control Firewall Manager. | [optional] 
**tenant_pay_type** | **str** | The pay type of the tenant. | [optional] 
**uid** | **str** | The unique identifier of the tenant in Security Cloud Control Firewall Manager. | 

## Example

```python
from scc_firewall_manager_sdk.models.msp_managed_tenant_dto import MspManagedTenantDto

# TODO update the JSON string below
json = "{}"
# create an instance of MspManagedTenantDto from a JSON string
msp_managed_tenant_dto_instance = MspManagedTenantDto.from_json(json)
# print the JSON string representation of the object
print(MspManagedTenantDto.to_json())

# convert the object into a dict
msp_managed_tenant_dto_dict = msp_managed_tenant_dto_instance.to_dict()
# create an instance of MspManagedTenantDto from a dict
msp_managed_tenant_dto_form_dict = msp_managed_tenant_dto.from_dict(msp_managed_tenant_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


