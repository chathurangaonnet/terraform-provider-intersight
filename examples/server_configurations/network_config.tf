resource "intersight_networkconfig_policy" "network_config1" {
  name                     = "network_config1"
  description              = "test policy"
  enable_dynamic_dns       = false
  preferred_ipv6dns_server = "::"
  enable_ipv6              = false
  enable_ipv6dns_from_dhcp = false
  preferred_ipv4dns_server = "171.70.168.183"
  alternate_ipv4dns_server = "173.36.131.10"
  alternate_ipv6dns_server = "::"
  dynamic_dns_domain       = ""
  enable_ipv4dns_from_dhcp = false
}

resource "intersight_networkconfig_policy" "network_config2" {
  name                     = "network_config2"
  description              = "test policy"
  enable_dynamic_dns       = true
  enable_ipv4dns_from_dhcp = false
  enable_ipv6              = false
  enable_ipv6dns_from_dhcp = false
  preferred_ipv4dns_server = "64.104.128.236"
  alternate_ipv4dns_server = "64.104.128.238"
  dynamic_dns_domain       = "qw12"
  alternate_ipv6dns_server = "::"
  preferred_ipv6dns_server = "::"
  profiles {
    moid        = intersight_server_profile.server1.id
    object_type = "server.Profile"
  }
}

/*
SAMPLE PAYLOAD
-----------------
NetworkconfigPolicyApi: {
    "Name": "AUTO_TEST_NETCONF_MASTER",
    "Description": "testing netconf policy",
    "Tags": [{"Key": "policy", "Value": "netconf"}],
    "EnableDynamicDns": True,
    "DynamicDnsDomain": "testdomain.com",
    "EnableIpv4dnsFromDhcp": False,
    "EnableIpv6": False,
    "PreferredIpv4dnsServer": "64.104.128.236",
    "AlternateIpv4dnsServer": ""
}
*/