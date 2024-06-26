# coding: utf-8

"""
    Cisco Defense Orchestrator API

    Use the interactive documentation to explore the endpoints CDO has to offer

    The version of the OpenAPI document: 0.0.1
    Contact: cdo.tac@cisco.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import unittest

from cdo_sdk_python.models.service_object_value_content import ServiceObjectValueContent

class TestServiceObjectValueContent(unittest.TestCase):
    """ServiceObjectValueContent unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> ServiceObjectValueContent:
        """Test ServiceObjectValueContent
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `ServiceObjectValueContent`
        """
        model = ServiceObjectValueContent()
        if include_optional:
            return ServiceObjectValueContent(
                icmp4_type = 'DESTINATION_UNREACHABLE',
                icmp4_code = 'NET_UNREACHABLE',
                icmp6_type = 'DESTINATION_UNREACHABLE',
                icmp6_code = 'ADDRESS_UNREACHABLE',
                op = 'RANGE',
                ports = [80, 81]
            )
        else:
            return ServiceObjectValueContent(
        )
        """

    def testServiceObjectValueContent(self):
        """Test ServiceObjectValueContent"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
