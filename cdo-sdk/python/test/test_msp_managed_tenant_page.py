# coding: utf-8

"""
    CDO API

    Use the documentation to explore the endpoints CDO has to offer

    The version of the OpenAPI document: 1.4.0
    Contact: cdo.tac@cisco.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import unittest

from cdo_sdk_python.models.msp_managed_tenant_page import MspManagedTenantPage

class TestMspManagedTenantPage(unittest.TestCase):
    """MspManagedTenantPage unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> MspManagedTenantPage:
        """Test MspManagedTenantPage
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `MspManagedTenantPage`
        """
        model = MspManagedTenantPage()
        if include_optional:
            return MspManagedTenantPage(
                count = 100,
                limit = 50,
                offset = 0,
                items = [
                    cdo_sdk_python.models.msp_managed_tenant.MspManagedTenant(
                        display_name = 'mytenantname', 
                        uid = '7131daad-e813-4b8f-8f42-be1e241e8cdb', 
                        name = 'my-example-device', 
                        region = 'US', )
                    ]
            )
        else:
            return MspManagedTenantPage(
        )
        """

    def testMspManagedTenantPage(self):
        """Test MspManagedTenantPage"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
