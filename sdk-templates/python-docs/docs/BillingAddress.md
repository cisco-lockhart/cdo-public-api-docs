# BillingAddress

The billing address for the tenant's MSLA subscription. If omitted, no billing address is recorded.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**city** | **str** |  | [optional] 
**country** | **str** |  | [optional] 
**country_code** | **str** |  | [optional] 
**postal_code** | **str** |  | [optional] 
**state** | **str** |  | [optional] 
**street_address_line1** | **str** |  | [optional] 
**street_address_line2** | **str** |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.billing_address import BillingAddress

# TODO update the JSON string below
json = "{}"
# create an instance of BillingAddress from a JSON string
billing_address_instance = BillingAddress.from_json(json)
# print the JSON string representation of the object
print(BillingAddress.to_json())

# convert the object into a dict
billing_address_dict = billing_address_instance.to_dict()
# create an instance of BillingAddress from a dict
billing_address_form_dict = billing_address.from_dict(billing_address_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


