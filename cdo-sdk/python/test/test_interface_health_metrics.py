# coding: utf-8

"""
    CDO API

    Use the documentation to explore the endpoints CDO has to offer

    The version of the OpenAPI document: 1.3.0
    Contact: cdo.tac@cisco.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import unittest

from cdo_sdk_python.models.interface_health_metrics import InterfaceHealthMetrics

class TestInterfaceHealthMetrics(unittest.TestCase):
    """InterfaceHealthMetrics unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> InterfaceHealthMetrics:
        """Test InterfaceHealthMetrics
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `InterfaceHealthMetrics`
        """
        model = InterfaceHealthMetrics()
        if include_optional:
            return InterfaceHealthMetrics(
                status = '',
                buffer_underruns_avg = 1.337,
                buffer_overruns_avg = 1.337,
                interface = '',
                interface_name = ''
            )
        else:
            return InterfaceHealthMetrics(
        )
        """

    def testInterfaceHealthMetrics(self):
        """Test InterfaceHealthMetrics"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
