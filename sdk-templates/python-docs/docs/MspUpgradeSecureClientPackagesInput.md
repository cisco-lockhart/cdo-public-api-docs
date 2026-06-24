# MspUpgradeSecureClientPackagesInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**device_uids** | **List[str]** | &lt;p&gt;List of device identifiers to upgrade in Security Cloud Control.&lt;/p&gt; &lt;p&gt;Each entry must meet all of the following:&lt;/p&gt; &lt;ul&gt;   &lt;li&gt;The identifier is a UUID (unique identifier).&lt;/li&gt;   &lt;li&gt;The device is managed by the MSSP portal.&lt;/li&gt;   &lt;li&gt;The device type is &lt;code&gt;ASA&lt;/code&gt;.&lt;/li&gt;   &lt;li&gt;The device has the &lt;code&gt;ASA_RA_VPN_HEADEND&lt;/code&gt; role.&lt;/li&gt;   &lt;li&gt;The connectivity state is &lt;code&gt;ONLINE&lt;/code&gt;.&lt;/li&gt;   &lt;li&gt;The configuration state is &lt;code&gt;SYNCED&lt;/code&gt;.&lt;/li&gt;   &lt;li&gt;The device has a known software version.&lt;/li&gt; &lt;/ul&gt; &lt;p&gt;Device-level validation is performed asynchronously; any device that fails these checks causes the upgrade run to fail with the offending devices reported in the error message.&lt;/p&gt; &lt;p&gt;A maximum of 50 devices per managed tenant is allowed.&lt;/p&gt; | 
**name** | **str** | Specify a human-readable name for the upgrade run. | [optional] 
**platforms** | [**List[MspSecureClientPlatform]**](MspSecureClientPlatform.md) | The OS and CPU architecture combinations to upgrade. If omitted, all platforms available for the specified version are uploaded to each device. | [optional] 
**version** | **str** | Specify the Secure Client package version to which all devices will be upgraded. | 

## Example

```python
from scc_firewall_manager_sdk.models.msp_upgrade_secure_client_packages_input import MspUpgradeSecureClientPackagesInput

# TODO update the JSON string below
json = "{}"
# create an instance of MspUpgradeSecureClientPackagesInput from a JSON string
msp_upgrade_secure_client_packages_input_instance = MspUpgradeSecureClientPackagesInput.from_json(json)
# print the JSON string representation of the object
print(MspUpgradeSecureClientPackagesInput.to_json())

# convert the object into a dict
msp_upgrade_secure_client_packages_input_dict = msp_upgrade_secure_client_packages_input_instance.to_dict()
# create an instance of MspUpgradeSecureClientPackagesInput from a dict
msp_upgrade_secure_client_packages_input_form_dict = msp_upgrade_secure_client_packages_input.from_dict(msp_upgrade_secure_client_packages_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


