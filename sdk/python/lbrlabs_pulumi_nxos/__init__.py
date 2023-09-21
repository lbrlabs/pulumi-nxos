# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .provider import *
from ._inputs import *

# Make subpackages available:
if typing.TYPE_CHECKING:
    import lbrlabs_pulumi_nxos.config as __config
    config = __config
    import lbrlabs_pulumi_nxos.nxos as __nxos
    nxos = __nxos
else:
    config = _utilities.lazy_import('lbrlabs_pulumi_nxos.config')
    nxos = _utilities.lazy_import('lbrlabs_pulumi_nxos.nxos')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "nxos",
  "mod": "nxos/bgp",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/bgp:Bgp": "Bgp"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/bgpAddressFamily",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/bgpAddressFamily:BgpAddressFamily": "BgpAddressFamily"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/bgpAdvertisedPrefix",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/bgpAdvertisedPrefix:BgpAdvertisedPrefix": "BgpAdvertisedPrefix"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/bgpGracefulRestart",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/bgpGracefulRestart:BgpGracefulRestart": "BgpGracefulRestart"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/bgpInstance",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/bgpInstance:BgpInstance": "BgpInstance"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/bgpPeer",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/bgpPeer:BgpPeer": "BgpPeer"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/bgpPeerAddressFamily",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/bgpPeerAddressFamily:BgpPeerAddressFamily": "BgpPeerAddressFamily"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/bgpPeerAddressFamilyPrefixListControl",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/bgpPeerAddressFamilyPrefixListControl:BgpPeerAddressFamilyPrefixListControl": "BgpPeerAddressFamilyPrefixListControl"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/bgpPeerAddressFamilyRouteControl",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/bgpPeerAddressFamilyRouteControl:BgpPeerAddressFamilyRouteControl": "BgpPeerAddressFamilyRouteControl"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/bgpPeerTemplate",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/bgpPeerTemplate:BgpPeerTemplate": "BgpPeerTemplate"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/bgpPeerTemplateAddressFamily",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/bgpPeerTemplateAddressFamily:BgpPeerTemplateAddressFamily": "BgpPeerTemplateAddressFamily"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/bgpPeerTemplateMaxPrefix",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/bgpPeerTemplateMaxPrefix:BgpPeerTemplateMaxPrefix": "BgpPeerTemplateMaxPrefix"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/bgpRouteControl",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/bgpRouteControl:BgpRouteControl": "BgpRouteControl"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/bgpVrf",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/bgpVrf:BgpVrf": "BgpVrf"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/bridgeDomain",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/bridgeDomain:BridgeDomain": "BridgeDomain"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/defaultQosClassMap",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/defaultQosClassMap:DefaultQosClassMap": "DefaultQosClassMap"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/defaultQosClassMapDscp",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/defaultQosClassMapDscp:DefaultQosClassMapDscp": "DefaultQosClassMapDscp"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/defaultQosPolicyInterfaceIn",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/defaultQosPolicyInterfaceIn:DefaultQosPolicyInterfaceIn": "DefaultQosPolicyInterfaceIn"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/defaultQosPolicyInterfaceInPolicyMap",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/defaultQosPolicyInterfaceInPolicyMap:DefaultQosPolicyInterfaceInPolicyMap": "DefaultQosPolicyInterfaceInPolicyMap"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/defaultQosPolicyMap",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/defaultQosPolicyMap:DefaultQosPolicyMap": "DefaultQosPolicyMap"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/defaultQosPolicyMapMatchClassMap",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/defaultQosPolicyMapMatchClassMap:DefaultQosPolicyMapMatchClassMap": "DefaultQosPolicyMapMatchClassMap"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/defaultQosPolicyMapMatchClassMapPolice",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/defaultQosPolicyMapMatchClassMapPolice:DefaultQosPolicyMapMatchClassMapPolice": "DefaultQosPolicyMapMatchClassMapPolice"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/defaultQosPolicyMapMatchClassMapSetQosGroup",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/defaultQosPolicyMapMatchClassMapSetQosGroup:DefaultQosPolicyMapMatchClassMapSetQosGroup": "DefaultQosPolicyMapMatchClassMapSetQosGroup"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/dhcpRelayAddress",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/dhcpRelayAddress:DhcpRelayAddress": "DhcpRelayAddress"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/dhcpRelayInterface",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/dhcpRelayInterface:DhcpRelayInterface": "DhcpRelayInterface"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/ethernet",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/ethernet:Ethernet": "Ethernet"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/evpn",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/evpn:Evpn": "Evpn"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/evpnVni",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/evpnVni:EvpnVni": "EvpnVni"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/evpnVniRouteTarget",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/evpnVniRouteTarget:EvpnVniRouteTarget": "EvpnVniRouteTarget"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/evpnVniRouteTargetDirection",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/evpnVniRouteTargetDirection:EvpnVniRouteTargetDirection": "EvpnVniRouteTargetDirection"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/featureBfd",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/featureBfd:FeatureBfd": "FeatureBfd"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/featureBgp",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/featureBgp:FeatureBgp": "FeatureBgp"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/featureDhcp",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/featureDhcp:FeatureDhcp": "FeatureDhcp"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/featureEvpn",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/featureEvpn:FeatureEvpn": "FeatureEvpn"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/featureHmm",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/featureHmm:FeatureHmm": "FeatureHmm"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/featureHsrp",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/featureHsrp:FeatureHsrp": "FeatureHsrp"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/featureInterfaceVlan",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/featureInterfaceVlan:FeatureInterfaceVlan": "FeatureInterfaceVlan"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/featureIsis",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/featureIsis:FeatureIsis": "FeatureIsis"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/featureLacp",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/featureLacp:FeatureLacp": "FeatureLacp"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/featureLldp",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/featureLldp:FeatureLldp": "FeatureLldp"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/featureMacsec",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/featureMacsec:FeatureMacsec": "FeatureMacsec"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/featureNetflow",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/featureNetflow:FeatureNetflow": "FeatureNetflow"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/featureNvOverlay",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/featureNvOverlay:FeatureNvOverlay": "FeatureNvOverlay"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/featureOspf",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/featureOspf:FeatureOspf": "FeatureOspf"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/featureOspfv3",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/featureOspfv3:FeatureOspfv3": "FeatureOspfv3"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/featurePim",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/featurePim:FeaturePim": "FeaturePim"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/featurePtp",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/featurePtp:FeaturePtp": "FeaturePtp"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/featurePvlan",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/featurePvlan:FeaturePvlan": "FeaturePvlan"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/featureSsh",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/featureSsh:FeatureSsh": "FeatureSsh"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/featureTacacs",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/featureTacacs:FeatureTacacs": "FeatureTacacs"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/featureTelnet",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/featureTelnet:FeatureTelnet": "FeatureTelnet"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/featureUdld",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/featureUdld:FeatureUdld": "FeatureUdld"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/featureVnSegment",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/featureVnSegment:FeatureVnSegment": "FeatureVnSegment"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/featureVpc",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/featureVpc:FeatureVpc": "FeatureVpc"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/hmm",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/hmm:Hmm": "Hmm"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/hmmInstance",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/hmmInstance:HmmInstance": "HmmInstance"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/hmmInterface",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/hmmInterface:HmmInterface": "HmmInterface"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/ipv4AccessList",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/ipv4AccessList:Ipv4AccessList": "Ipv4AccessList"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/ipv4AccessListEntry",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/ipv4AccessListEntry:Ipv4AccessListEntry": "Ipv4AccessListEntry"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/ipv4AccessListPolicyEgressInterface",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/ipv4AccessListPolicyEgressInterface:Ipv4AccessListPolicyEgressInterface": "Ipv4AccessListPolicyEgressInterface"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/ipv4AccessListPolicyIngressInterface",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/ipv4AccessListPolicyIngressInterface:Ipv4AccessListPolicyIngressInterface": "Ipv4AccessListPolicyIngressInterface"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/ipv4Interface",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/ipv4Interface:Ipv4Interface": "Ipv4Interface"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/ipv4InterfaceAddress",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/ipv4InterfaceAddress:Ipv4InterfaceAddress": "Ipv4InterfaceAddress"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/ipv4PrefixListRule",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/ipv4PrefixListRule:Ipv4PrefixListRule": "Ipv4PrefixListRule"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/ipv4PrefixListRuleEntry",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/ipv4PrefixListRuleEntry:Ipv4PrefixListRuleEntry": "Ipv4PrefixListRuleEntry"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/ipv4StaticRoute",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/ipv4StaticRoute:Ipv4StaticRoute": "Ipv4StaticRoute"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/ipv4Vrf",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/ipv4Vrf:Ipv4Vrf": "Ipv4Vrf"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/isis",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/isis:Isis": "Isis"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/isisInstance",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/isisInstance:IsisInstance": "IsisInstance"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/isisInterface",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/isisInterface:IsisInterface": "IsisInterface"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/isisVrf",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/isisVrf:IsisVrf": "IsisVrf"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/loopbackInterface",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/loopbackInterface:LoopbackInterface": "LoopbackInterface"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/loopbackInterfaceVrf",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/loopbackInterfaceVrf:LoopbackInterfaceVrf": "LoopbackInterfaceVrf"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/ntpServer",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/ntpServer:NtpServer": "NtpServer"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/nveInterface",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/nveInterface:NveInterface": "NveInterface"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/nveVni",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/nveVni:NveVni": "NveVni"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/nveVniContainer",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/nveVniContainer:NveVniContainer": "NveVniContainer"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/nveVniIngressReplication",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/nveVniIngressReplication:NveVniIngressReplication": "NveVniIngressReplication"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/ospf",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/ospf:Ospf": "Ospf"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/ospfArea",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/ospfArea:OspfArea": "OspfArea"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/ospfAuthentication",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/ospfAuthentication:OspfAuthentication": "OspfAuthentication"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/ospfInstance",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/ospfInstance:OspfInstance": "OspfInstance"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/ospfInterface",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/ospfInterface:OspfInterface": "OspfInterface"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/ospfVrf",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/ospfVrf:OspfVrf": "OspfVrf"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/physicalInterface",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/physicalInterface:PhysicalInterface": "PhysicalInterface"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/physicalInterfaceVrf",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/physicalInterfaceVrf:PhysicalInterfaceVrf": "PhysicalInterfaceVrf"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/pim",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/pim:Pim": "Pim"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/pimAnycastRp",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/pimAnycastRp:PimAnycastRp": "PimAnycastRp"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/pimAnycastRpPeer",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/pimAnycastRpPeer:PimAnycastRpPeer": "PimAnycastRpPeer"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/pimInstance",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/pimInstance:PimInstance": "PimInstance"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/pimInterface",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/pimInterface:PimInterface": "PimInterface"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/pimSsmPolicy",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/pimSsmPolicy:PimSsmPolicy": "PimSsmPolicy"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/pimSsmRange",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/pimSsmRange:PimSsmRange": "PimSsmRange"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/pimStaticRp",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/pimStaticRp:PimStaticRp": "PimStaticRp"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/pimStaticRpGroupList",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/pimStaticRpGroupList:PimStaticRpGroupList": "PimStaticRpGroupList"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/pimStaticRpPolicy",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/pimStaticRpPolicy:PimStaticRpPolicy": "PimStaticRpPolicy"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/pimVrf",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/pimVrf:PimVrf": "PimVrf"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/portChannelInterface",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/portChannelInterface:PortChannelInterface": "PortChannelInterface"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/portChannelInterfaceMember",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/portChannelInterfaceMember:PortChannelInterfaceMember": "PortChannelInterfaceMember"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/portChannelInterfaceVrf",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/portChannelInterfaceVrf:PortChannelInterfaceVrf": "PortChannelInterfaceVrf"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/queuingQosPolicyMap",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/queuingQosPolicyMap:QueuingQosPolicyMap": "QueuingQosPolicyMap"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/queuingQosPolicyMapMatchClassMap",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/queuingQosPolicyMapMatchClassMap:QueuingQosPolicyMapMatchClassMap": "QueuingQosPolicyMapMatchClassMap"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/queuingQosPolicyMapMatchClassMapPriority",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/queuingQosPolicyMapMatchClassMapPriority:QueuingQosPolicyMapMatchClassMapPriority": "QueuingQosPolicyMapMatchClassMapPriority"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/queuingQosPolicyMapMatchClassMapRemainingBandwidth",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/queuingQosPolicyMapMatchClassMapRemainingBandwidth:QueuingQosPolicyMapMatchClassMapRemainingBandwidth": "QueuingQosPolicyMapMatchClassMapRemainingBandwidth"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/queuingQosPolicySystemOut",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/queuingQosPolicySystemOut:QueuingQosPolicySystemOut": "QueuingQosPolicySystemOut"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/queuingQosPolicySystemOutPolicyMap",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/queuingQosPolicySystemOutPolicyMap:QueuingQosPolicySystemOutPolicyMap": "QueuingQosPolicySystemOutPolicyMap"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/rest",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/rest:Rest": "Rest"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/routeMapRule",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/routeMapRule:RouteMapRule": "RouteMapRule"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/routeMapRuleEntry",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/routeMapRuleEntry:RouteMapRuleEntry": "RouteMapRuleEntry"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/routeMapRuleEntryMatchRoute",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/routeMapRuleEntryMatchRoute:RouteMapRuleEntryMatchRoute": "RouteMapRuleEntryMatchRoute"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/routeMapRuleEntryMatchRoutePrefixList",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/routeMapRuleEntryMatchRoutePrefixList:RouteMapRuleEntryMatchRoutePrefixList": "RouteMapRuleEntryMatchRoutePrefixList"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/routeMapRuleEntrySetRegularCommunity",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/routeMapRuleEntrySetRegularCommunity:RouteMapRuleEntrySetRegularCommunity": "RouteMapRuleEntrySetRegularCommunity"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/routeMapRuleEntrySetRegularCommunityItem",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/routeMapRuleEntrySetRegularCommunityItem:RouteMapRuleEntrySetRegularCommunityItem": "RouteMapRuleEntrySetRegularCommunityItem"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/spanningTreeInterface",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/spanningTreeInterface:SpanningTreeInterface": "SpanningTreeInterface"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/subinterface",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/subinterface:Subinterface": "Subinterface"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/subinterfaceVrf",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/subinterfaceVrf:SubinterfaceVrf": "SubinterfaceVrf"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/sviInterface",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/sviInterface:SviInterface": "SviInterface"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/sviInterfaceVrf",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/sviInterfaceVrf:SviInterfaceVrf": "SviInterfaceVrf"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/system",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/system:System": "System"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/vpcDomain",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/vpcDomain:VpcDomain": "VpcDomain"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/vpcInstance",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/vpcInstance:VpcInstance": "VpcInstance"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/vpcInterface",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/vpcInterface:VpcInterface": "VpcInterface"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/vrf",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/vrf:Vrf": "Vrf"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/vrfAddressFamily",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/vrfAddressFamily:VrfAddressFamily": "VrfAddressFamily"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/vrfRouteTarget",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/vrfRouteTarget:VrfRouteTarget": "VrfRouteTarget"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/vrfRouteTargetAddressFamily",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/vrfRouteTargetAddressFamily:VrfRouteTargetAddressFamily": "VrfRouteTargetAddressFamily"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/vrfRouteTargetDirection",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/vrfRouteTargetDirection:VrfRouteTargetDirection": "VrfRouteTargetDirection"
  }
 },
 {
  "pkg": "nxos",
  "mod": "nxos/vrfRouting",
  "fqn": "lbrlabs_pulumi_nxos.nxos",
  "classes": {
   "nxos:nxos/vrfRouting:VrfRouting": "VrfRouting"
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