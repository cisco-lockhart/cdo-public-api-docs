# MspManagedTenantDistinctAttributeValues


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**compliance_statuses** | **List[str]** | The compliance statuses of the tenants that are managed by the MSP Portal. | [optional] 
**regions** | **List[str]** | The regions of the tenants that are managed by the MSP Portal. | [optional] 
**tenant_pay_types** | **List[str]** | The pay types of the tenants that are managed by the MSP Portal. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.msp_managed_tenant_distinct_attribute_values import MspManagedTenantDistinctAttributeValues

# TODO update the JSON string below
json = "{}"
# create an instance of MspManagedTenantDistinctAttributeValues from a JSON string
msp_managed_tenant_distinct_attribute_values_instance = MspManagedTenantDistinctAttributeValues.from_json(json)
# print the JSON string representation of the object
print(MspManagedTenantDistinctAttributeValues.to_json())

# convert the object into a dict
msp_managed_tenant_distinct_attribute_values_dict = msp_managed_tenant_distinct_attribute_values_instance.to_dict()
# create an instance of MspManagedTenantDistinctAttributeValues from a dict
msp_managed_tenant_distinct_attribute_values_form_dict = msp_managed_tenant_distinct_attribute_values.from_dict(msp_managed_tenant_distinct_attribute_values_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


