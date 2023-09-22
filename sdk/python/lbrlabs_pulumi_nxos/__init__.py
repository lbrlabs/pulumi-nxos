# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .bgp import *
from .bgp_address_family import *
from .bgp_advertised_prefix import *
from .bgp_graceful_restart import *
from .bgp_instance import *
from .bgp_peer import *
from .bgp_peer_address_family import *
from .bgp_peer_address_family_prefix_list_control import *
from .bgp_peer_address_family_route_control import *
from .bgp_peer_template import *
from .bgp_peer_template_address_family import *
from .bgp_peer_template_max_prefix import *
from .bgp_route_control import *
from .bgp_vrf import *
from .bridge_domain import *
from .default_qos_class_map import *
from .default_qos_class_map_dscp import *
from .default_qos_policy_interface_in import *
from .default_qos_policy_interface_in_policy_map import *
from .default_qos_policy_map import *
from .default_qos_policy_map_match_class_map import *
from .default_qos_policy_map_match_class_map_police import *
from .default_qos_policy_map_match_class_map_set_qos_group import *
from .dhcp_relay_address import *
from .dhcp_relay_interface import *
from .ethernet import *
from .evpn import *
from .evpn_vni import *
from .evpn_vni_route_target import *
from .evpn_vni_route_target_direction import *
from .feature_bfd import *
from .feature_bgp import *
from .feature_dhcp import *
from .feature_evpn import *
from .feature_hmm import *
from .feature_hsrp import *
from .feature_interface_vlan import *
from .feature_isis import *
from .feature_lacp import *
from .feature_lldp import *
from .feature_macsec import *
from .feature_netflow import *
from .feature_nv_overlay import *
from .feature_ospf import *
from .feature_ospfv3 import *
from .feature_pim import *
from .feature_ptp import *
from .feature_pvlan import *
from .feature_ssh import *
from .feature_tacacs import *
from .feature_telnet import *
from .feature_udld import *
from .feature_vn_segment import *
from .feature_vpc import *
from .get_bgp import *
from .get_bgp_address_family import *
from .get_bgp_advertised_prefix import *
from .get_bgp_graceful_restart import *
from .get_bgp_instance import *
from .get_bgp_peer import *
from .get_bgp_peer_address_family import *
from .get_bgp_peer_address_family_prefix_list_control import *
from .get_bgp_peer_address_family_route_control import *
from .get_bgp_peer_template import *
from .get_bgp_peer_template_address_family import *
from .get_bgp_peer_template_max_prefix import *
from .get_bgp_route_control import *
from .get_bgp_vrf import *
from .get_bridge_domain import *
from .get_default_qos_class_map import *
from .get_default_qos_class_map_dscp import *
from .get_default_qos_policy_interface_in import *
from .get_default_qos_policy_interface_in_policy_map import *
from .get_default_qos_policy_map import *
from .get_default_qos_policy_map_match_class_map import *
from .get_default_qos_policy_map_match_class_map_police import *
from .get_default_qos_policy_map_match_class_map_set_qos_group import *
from .get_dhcp_relay_address import *
from .get_dhcp_relay_interface import *
from .get_ethernet import *
from .get_evpn import *
from .get_evpn_vni import *
from .get_evpn_vni_route_target import *
from .get_evpn_vni_route_target_direction import *
from .get_feature_bfd import *
from .get_feature_bgp import *
from .get_feature_dhcp import *
from .get_feature_evpn import *
from .get_feature_hmm import *
from .get_feature_hsrp import *
from .get_feature_interface_vlan import *
from .get_feature_isis import *
from .get_feature_lacp import *
from .get_feature_lldp import *
from .get_feature_macsec import *
from .get_feature_netflow import *
from .get_feature_nv_overlay import *
from .get_feature_ospf import *
from .get_feature_ospfv3 import *
from .get_feature_pim import *
from .get_feature_ptp import *
from .get_feature_pvlan import *
from .get_feature_ssh import *
from .get_feature_tacacs import *
from .get_feature_telnet import *
from .get_feature_udld import *
from .get_feature_vn_segment import *
from .get_feature_vpc import *
from .get_hmm import *
from .get_hmm_instance import *
from .get_hmm_interface import *
from .get_ipv4_access_list import *
from .get_ipv4_access_list_entry import *
from .get_ipv4_access_list_policy_egress_interface import *
from .get_ipv4_access_list_policy_ingress_interface import *
from .get_ipv4_interface import *
from .get_ipv4_interface_address import *
from .get_ipv4_prefix_list_rule import *
from .get_ipv4_prefix_list_rule_entry import *
from .get_ipv4_static_route import *
from .get_ipv4_vrf import *
from .get_isis import *
from .get_isis_instance import *
from .get_isis_interface import *
from .get_isis_vrf import *
from .get_loopback_interface import *
from .get_loopback_interface_vrf import *
from .get_ntp_server import *
from .get_nve_interface import *
from .get_nve_vni import *
from .get_nve_vni_container import *
from .get_nve_vni_ingress_replication import *
from .get_ospf import *
from .get_ospf_area import *
from .get_ospf_authentication import *
from .get_ospf_instance import *
from .get_ospf_interface import *
from .get_ospf_vrf import *
from .get_physical_interface import *
from .get_physical_interface_vrf import *
from .get_pim import *
from .get_pim_anycast_rp import *
from .get_pim_anycast_rp_peer import *
from .get_pim_instance import *
from .get_pim_interface import *
from .get_pim_ssm_policy import *
from .get_pim_ssm_range import *
from .get_pim_static_rp import *
from .get_pim_static_rp_group_list import *
from .get_pim_static_rp_policy import *
from .get_pim_vrf import *
from .get_port_channel_interface import *
from .get_port_channel_interface_member import *
from .get_port_channel_interface_vrf import *
from .get_queuing_qos_policy_map import *
from .get_queuing_qos_policy_map_match_class_map import *
from .get_queuing_qos_policy_map_match_class_map_priority import *
from .get_queuing_qos_policy_map_match_class_map_remaining_bandwidth import *
from .get_queuing_qos_policy_system_out import *
from .get_queuing_qos_policy_system_out_policy_map import *
from .get_rest import *
from .get_route_map_rule import *
from .get_route_map_rule_entry import *
from .get_route_map_rule_entry_match_route import *
from .get_route_map_rule_entry_match_route_prefix_list import *
from .get_route_map_rule_entry_set_regular_community import *
from .get_route_map_rule_entry_set_regular_community_item import *
from .get_spanning_tree_interface import *
from .get_subinterface import *
from .get_subinterface_vrf import *
from .get_svi_interface import *
from .get_svi_interface_vrf import *
from .get_system import *
from .get_vpc_domain import *
from .get_vpc_instance import *
from .get_vpc_interface import *
from .get_vrf import *
from .get_vrf_address_family import *
from .get_vrf_route_target import *
from .get_vrf_route_target_address_family import *
from .get_vrf_route_target_direction import *
from .get_vrf_routing import *
from .hmm import *
from .hmm_instance import *
from .hmm_interface import *
from .ipv4_access_list import *
from .ipv4_access_list_entry import *
from .ipv4_access_list_policy_egress_interface import *
from .ipv4_access_list_policy_ingress_interface import *
from .ipv4_interface import *
from .ipv4_interface_address import *
from .ipv4_prefix_list_rule import *
from .ipv4_prefix_list_rule_entry import *
from .ipv4_static_route import *
from .ipv4_vrf import *
from .isis import *
from .isis_instance import *
from .isis_interface import *
from .isis_vrf import *
from .loopback_interface import *
from .loopback_interface_vrf import *
from .ntp_server import *
from .nve_interface import *
from .nve_vni import *
from .nve_vni_container import *
from .nve_vni_ingress_replication import *
from .ospf import *
from .ospf_area import *
from .ospf_authentication import *
from .ospf_instance import *
from .ospf_interface import *
from .ospf_vrf import *
from .physical_interface import *
from .physical_interface_vrf import *
from .pim import *
from .pim_anycast_rp import *
from .pim_anycast_rp_peer import *
from .pim_instance import *
from .pim_interface import *
from .pim_ssm_policy import *
from .pim_ssm_range import *
from .pim_static_rp import *
from .pim_static_rp_group_list import *
from .pim_static_rp_policy import *
from .pim_vrf import *
from .port_channel_interface import *
from .port_channel_interface_member import *
from .port_channel_interface_vrf import *
from .provider import *
from .queuing_qos_policy_map import *
from .queuing_qos_policy_map_match_class_map import *
from .queuing_qos_policy_map_match_class_map_priority import *
from .queuing_qos_policy_map_match_class_map_remaining_bandwidth import *
from .queuing_qos_policy_system_out import *
from .queuing_qos_policy_system_out_policy_map import *
from .rest import *
from .route_map_rule import *
from .route_map_rule_entry import *
from .route_map_rule_entry_match_route import *
from .route_map_rule_entry_match_route_prefix_list import *
from .route_map_rule_entry_set_regular_community import *
from .route_map_rule_entry_set_regular_community_item import *
from .spanning_tree_interface import *
from .subinterface import *
from .subinterface_vrf import *
from .svi_interface import *
from .svi_interface_vrf import *
from .system import *
from .vpc_domain import *
from .vpc_instance import *
from .vpc_interface import *
from .vrf import *
from .vrf_address_family import *
from .vrf_route_target import *
from .vrf_route_target_address_family import *
from .vrf_route_target_direction import *
from .vrf_routing import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import lbrlabs_pulumi_nxos.config as __config
    config = __config
else:
    config = _utilities.lazy_import('lbrlabs_pulumi_nxos.config')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "nxos",
  "mod": "index/bgp",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/bgp:Bgp": "Bgp"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/bgpAddressFamily",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/bgpAddressFamily:BgpAddressFamily": "BgpAddressFamily"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/bgpAdvertisedPrefix",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/bgpAdvertisedPrefix:BgpAdvertisedPrefix": "BgpAdvertisedPrefix"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/bgpGracefulRestart",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/bgpGracefulRestart:BgpGracefulRestart": "BgpGracefulRestart"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/bgpInstance",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/bgpInstance:BgpInstance": "BgpInstance"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/bgpPeer",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/bgpPeer:BgpPeer": "BgpPeer"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/bgpPeerAddressFamily",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/bgpPeerAddressFamily:BgpPeerAddressFamily": "BgpPeerAddressFamily"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/bgpPeerAddressFamilyPrefixListControl",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/bgpPeerAddressFamilyPrefixListControl:BgpPeerAddressFamilyPrefixListControl": "BgpPeerAddressFamilyPrefixListControl"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/bgpPeerAddressFamilyRouteControl",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/bgpPeerAddressFamilyRouteControl:BgpPeerAddressFamilyRouteControl": "BgpPeerAddressFamilyRouteControl"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/bgpPeerTemplate",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/bgpPeerTemplate:BgpPeerTemplate": "BgpPeerTemplate"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/bgpPeerTemplateAddressFamily",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/bgpPeerTemplateAddressFamily:BgpPeerTemplateAddressFamily": "BgpPeerTemplateAddressFamily"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/bgpPeerTemplateMaxPrefix",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/bgpPeerTemplateMaxPrefix:BgpPeerTemplateMaxPrefix": "BgpPeerTemplateMaxPrefix"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/bgpRouteControl",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/bgpRouteControl:BgpRouteControl": "BgpRouteControl"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/bgpVrf",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/bgpVrf:BgpVrf": "BgpVrf"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/bridgeDomain",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/bridgeDomain:BridgeDomain": "BridgeDomain"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/defaultQosClassMap",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/defaultQosClassMap:DefaultQosClassMap": "DefaultQosClassMap"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/defaultQosClassMapDscp",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/defaultQosClassMapDscp:DefaultQosClassMapDscp": "DefaultQosClassMapDscp"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/defaultQosPolicyInterfaceIn",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/defaultQosPolicyInterfaceIn:DefaultQosPolicyInterfaceIn": "DefaultQosPolicyInterfaceIn"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/defaultQosPolicyInterfaceInPolicyMap",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/defaultQosPolicyInterfaceInPolicyMap:DefaultQosPolicyInterfaceInPolicyMap": "DefaultQosPolicyInterfaceInPolicyMap"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/defaultQosPolicyMap",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/defaultQosPolicyMap:DefaultQosPolicyMap": "DefaultQosPolicyMap"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/defaultQosPolicyMapMatchClassMap",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/defaultQosPolicyMapMatchClassMap:DefaultQosPolicyMapMatchClassMap": "DefaultQosPolicyMapMatchClassMap"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/defaultQosPolicyMapMatchClassMapPolice",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/defaultQosPolicyMapMatchClassMapPolice:DefaultQosPolicyMapMatchClassMapPolice": "DefaultQosPolicyMapMatchClassMapPolice"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/defaultQosPolicyMapMatchClassMapSetQosGroup",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/defaultQosPolicyMapMatchClassMapSetQosGroup:DefaultQosPolicyMapMatchClassMapSetQosGroup": "DefaultQosPolicyMapMatchClassMapSetQosGroup"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/dhcpRelayAddress",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/dhcpRelayAddress:DhcpRelayAddress": "DhcpRelayAddress"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/dhcpRelayInterface",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/dhcpRelayInterface:DhcpRelayInterface": "DhcpRelayInterface"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/ethernet",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/ethernet:Ethernet": "Ethernet"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/evpn",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/evpn:Evpn": "Evpn"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/evpnVni",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/evpnVni:EvpnVni": "EvpnVni"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/evpnVniRouteTarget",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/evpnVniRouteTarget:EvpnVniRouteTarget": "EvpnVniRouteTarget"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/evpnVniRouteTargetDirection",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/evpnVniRouteTargetDirection:EvpnVniRouteTargetDirection": "EvpnVniRouteTargetDirection"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/featureBfd",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/featureBfd:FeatureBfd": "FeatureBfd"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/featureBgp",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/featureBgp:FeatureBgp": "FeatureBgp"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/featureDhcp",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/featureDhcp:FeatureDhcp": "FeatureDhcp"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/featureEvpn",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/featureEvpn:FeatureEvpn": "FeatureEvpn"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/featureHmm",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/featureHmm:FeatureHmm": "FeatureHmm"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/featureHsrp",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/featureHsrp:FeatureHsrp": "FeatureHsrp"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/featureInterfaceVlan",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/featureInterfaceVlan:FeatureInterfaceVlan": "FeatureInterfaceVlan"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/featureIsis",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/featureIsis:FeatureIsis": "FeatureIsis"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/featureLacp",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/featureLacp:FeatureLacp": "FeatureLacp"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/featureLldp",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/featureLldp:FeatureLldp": "FeatureLldp"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/featureMacsec",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/featureMacsec:FeatureMacsec": "FeatureMacsec"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/featureNetflow",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/featureNetflow:FeatureNetflow": "FeatureNetflow"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/featureNvOverlay",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/featureNvOverlay:FeatureNvOverlay": "FeatureNvOverlay"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/featureOspf",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/featureOspf:FeatureOspf": "FeatureOspf"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/featureOspfv3",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/featureOspfv3:FeatureOspfv3": "FeatureOspfv3"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/featurePim",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/featurePim:FeaturePim": "FeaturePim"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/featurePtp",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/featurePtp:FeaturePtp": "FeaturePtp"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/featurePvlan",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/featurePvlan:FeaturePvlan": "FeaturePvlan"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/featureSsh",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/featureSsh:FeatureSsh": "FeatureSsh"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/featureTacacs",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/featureTacacs:FeatureTacacs": "FeatureTacacs"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/featureTelnet",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/featureTelnet:FeatureTelnet": "FeatureTelnet"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/featureUdld",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/featureUdld:FeatureUdld": "FeatureUdld"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/featureVnSegment",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/featureVnSegment:FeatureVnSegment": "FeatureVnSegment"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/featureVpc",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/featureVpc:FeatureVpc": "FeatureVpc"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/hmm",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/hmm:Hmm": "Hmm"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/hmmInstance",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/hmmInstance:HmmInstance": "HmmInstance"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/hmmInterface",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/hmmInterface:HmmInterface": "HmmInterface"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/ipv4AccessList",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/ipv4AccessList:Ipv4AccessList": "Ipv4AccessList"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/ipv4AccessListEntry",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/ipv4AccessListEntry:Ipv4AccessListEntry": "Ipv4AccessListEntry"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/ipv4AccessListPolicyEgressInterface",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/ipv4AccessListPolicyEgressInterface:Ipv4AccessListPolicyEgressInterface": "Ipv4AccessListPolicyEgressInterface"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/ipv4AccessListPolicyIngressInterface",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/ipv4AccessListPolicyIngressInterface:Ipv4AccessListPolicyIngressInterface": "Ipv4AccessListPolicyIngressInterface"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/ipv4Interface",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/ipv4Interface:Ipv4Interface": "Ipv4Interface"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/ipv4InterfaceAddress",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/ipv4InterfaceAddress:Ipv4InterfaceAddress": "Ipv4InterfaceAddress"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/ipv4PrefixListRule",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/ipv4PrefixListRule:Ipv4PrefixListRule": "Ipv4PrefixListRule"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/ipv4PrefixListRuleEntry",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/ipv4PrefixListRuleEntry:Ipv4PrefixListRuleEntry": "Ipv4PrefixListRuleEntry"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/ipv4StaticRoute",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/ipv4StaticRoute:Ipv4StaticRoute": "Ipv4StaticRoute"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/ipv4Vrf",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/ipv4Vrf:Ipv4Vrf": "Ipv4Vrf"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/isis",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/isis:Isis": "Isis"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/isisInstance",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/isisInstance:IsisInstance": "IsisInstance"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/isisInterface",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/isisInterface:IsisInterface": "IsisInterface"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/isisVrf",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/isisVrf:IsisVrf": "IsisVrf"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/loopbackInterface",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/loopbackInterface:LoopbackInterface": "LoopbackInterface"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/loopbackInterfaceVrf",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/loopbackInterfaceVrf:LoopbackInterfaceVrf": "LoopbackInterfaceVrf"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/ntpServer",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/ntpServer:NtpServer": "NtpServer"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/nveInterface",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/nveInterface:NveInterface": "NveInterface"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/nveVni",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/nveVni:NveVni": "NveVni"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/nveVniContainer",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/nveVniContainer:NveVniContainer": "NveVniContainer"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/nveVniIngressReplication",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/nveVniIngressReplication:NveVniIngressReplication": "NveVniIngressReplication"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/ospf",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/ospf:Ospf": "Ospf"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/ospfArea",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/ospfArea:OspfArea": "OspfArea"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/ospfAuthentication",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/ospfAuthentication:OspfAuthentication": "OspfAuthentication"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/ospfInstance",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/ospfInstance:OspfInstance": "OspfInstance"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/ospfInterface",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/ospfInterface:OspfInterface": "OspfInterface"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/ospfVrf",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/ospfVrf:OspfVrf": "OspfVrf"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/physicalInterface",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/physicalInterface:PhysicalInterface": "PhysicalInterface"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/physicalInterfaceVrf",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/physicalInterfaceVrf:PhysicalInterfaceVrf": "PhysicalInterfaceVrf"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/pim",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/pim:Pim": "Pim"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/pimAnycastRp",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/pimAnycastRp:PimAnycastRp": "PimAnycastRp"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/pimAnycastRpPeer",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/pimAnycastRpPeer:PimAnycastRpPeer": "PimAnycastRpPeer"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/pimInstance",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/pimInstance:PimInstance": "PimInstance"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/pimInterface",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/pimInterface:PimInterface": "PimInterface"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/pimSsmPolicy",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/pimSsmPolicy:PimSsmPolicy": "PimSsmPolicy"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/pimSsmRange",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/pimSsmRange:PimSsmRange": "PimSsmRange"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/pimStaticRp",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/pimStaticRp:PimStaticRp": "PimStaticRp"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/pimStaticRpGroupList",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/pimStaticRpGroupList:PimStaticRpGroupList": "PimStaticRpGroupList"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/pimStaticRpPolicy",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/pimStaticRpPolicy:PimStaticRpPolicy": "PimStaticRpPolicy"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/pimVrf",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/pimVrf:PimVrf": "PimVrf"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/portChannelInterface",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/portChannelInterface:PortChannelInterface": "PortChannelInterface"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/portChannelInterfaceMember",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/portChannelInterfaceMember:PortChannelInterfaceMember": "PortChannelInterfaceMember"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/portChannelInterfaceVrf",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/portChannelInterfaceVrf:PortChannelInterfaceVrf": "PortChannelInterfaceVrf"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/queuingQosPolicyMap",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/queuingQosPolicyMap:QueuingQosPolicyMap": "QueuingQosPolicyMap"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/queuingQosPolicyMapMatchClassMap",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/queuingQosPolicyMapMatchClassMap:QueuingQosPolicyMapMatchClassMap": "QueuingQosPolicyMapMatchClassMap"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/queuingQosPolicyMapMatchClassMapPriority",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/queuingQosPolicyMapMatchClassMapPriority:QueuingQosPolicyMapMatchClassMapPriority": "QueuingQosPolicyMapMatchClassMapPriority"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/queuingQosPolicyMapMatchClassMapRemainingBandwidth",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/queuingQosPolicyMapMatchClassMapRemainingBandwidth:QueuingQosPolicyMapMatchClassMapRemainingBandwidth": "QueuingQosPolicyMapMatchClassMapRemainingBandwidth"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/queuingQosPolicySystemOut",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/queuingQosPolicySystemOut:QueuingQosPolicySystemOut": "QueuingQosPolicySystemOut"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/queuingQosPolicySystemOutPolicyMap",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/queuingQosPolicySystemOutPolicyMap:QueuingQosPolicySystemOutPolicyMap": "QueuingQosPolicySystemOutPolicyMap"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/rest",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/rest:Rest": "Rest"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/routeMapRule",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/routeMapRule:RouteMapRule": "RouteMapRule"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/routeMapRuleEntry",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/routeMapRuleEntry:RouteMapRuleEntry": "RouteMapRuleEntry"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/routeMapRuleEntryMatchRoute",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/routeMapRuleEntryMatchRoute:RouteMapRuleEntryMatchRoute": "RouteMapRuleEntryMatchRoute"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/routeMapRuleEntryMatchRoutePrefixList",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/routeMapRuleEntryMatchRoutePrefixList:RouteMapRuleEntryMatchRoutePrefixList": "RouteMapRuleEntryMatchRoutePrefixList"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/routeMapRuleEntrySetRegularCommunity",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/routeMapRuleEntrySetRegularCommunity:RouteMapRuleEntrySetRegularCommunity": "RouteMapRuleEntrySetRegularCommunity"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/routeMapRuleEntrySetRegularCommunityItem",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/routeMapRuleEntrySetRegularCommunityItem:RouteMapRuleEntrySetRegularCommunityItem": "RouteMapRuleEntrySetRegularCommunityItem"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/spanningTreeInterface",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/spanningTreeInterface:SpanningTreeInterface": "SpanningTreeInterface"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/subinterface",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/subinterface:Subinterface": "Subinterface"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/subinterfaceVrf",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/subinterfaceVrf:SubinterfaceVrf": "SubinterfaceVrf"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/sviInterface",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/sviInterface:SviInterface": "SviInterface"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/sviInterfaceVrf",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/sviInterfaceVrf:SviInterfaceVrf": "SviInterfaceVrf"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/system",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/system:System": "System"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/vpcDomain",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/vpcDomain:VpcDomain": "VpcDomain"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/vpcInstance",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/vpcInstance:VpcInstance": "VpcInstance"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/vpcInterface",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/vpcInterface:VpcInterface": "VpcInterface"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/vrf",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/vrf:Vrf": "Vrf"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/vrfAddressFamily",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/vrfAddressFamily:VrfAddressFamily": "VrfAddressFamily"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/vrfRouteTarget",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/vrfRouteTarget:VrfRouteTarget": "VrfRouteTarget"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/vrfRouteTargetAddressFamily",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/vrfRouteTargetAddressFamily:VrfRouteTargetAddressFamily": "VrfRouteTargetAddressFamily"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/vrfRouteTargetDirection",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/vrfRouteTargetDirection:VrfRouteTargetDirection": "VrfRouteTargetDirection"
  }
 },
 {
  "pkg": "nxos",
  "mod": "index/vrfRouting",
  "fqn": "lbrlabs_pulumi_nxos",
  "classes": {
   "nxos:index/vrfRouting:VrfRouting": "VrfRouting"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "nxos",
  "token": "pulumi:providers:nxos",
  "fqn": "lbrlabs_pulumi_nxos",
  "class": "Provider"
 }
]
"""
)
