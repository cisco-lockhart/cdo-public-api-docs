# coding: utf-8

"""
    CDO API

    Use the documentation to explore the endpoints CDO has to offer

    The version of the OpenAPI document: 1.2.0
    Contact: cdo.tac@cisco.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import unittest

from cdo_sdk_python.models.fmc_health_metrics import FmcHealthMetrics

class TestFmcHealthMetrics(unittest.TestCase):
    """FmcHealthMetrics unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> FmcHealthMetrics:
        """Test FmcHealthMetrics
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `FmcHealthMetrics`
        """
        model = FmcHealthMetrics()
        if include_optional:
            return FmcHealthMetrics(
                device_uid = '7131daad-e813-4b8f-8f42-be1e241e8cdb',
                cpu_health_metrics = cdo_sdk_python.models.cpu_health_metrics.CpuHealthMetrics(
                    lina_usage_avg = 1.337, 
                    snort_usage_avg = 1.337, 
                    system_usage_avg = 1.337, )
            )
        else:
            return FmcHealthMetrics(
        )
        """

    def testFmcHealthMetrics(self):
        """Test FmcHealthMetrics"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()