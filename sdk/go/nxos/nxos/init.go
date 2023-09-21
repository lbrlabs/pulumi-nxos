// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nxos

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/lbrlabs/pulumi-nxos/sdk/go/nxos/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "nxos:nxos/bgp:Bgp":
		r = &Bgp{}
	case "nxos:nxos/bgpAddressFamily:BgpAddressFamily":
		r = &BgpAddressFamily{}
	case "nxos:nxos/bgpAdvertisedPrefix:BgpAdvertisedPrefix":
		r = &BgpAdvertisedPrefix{}
	case "nxos:nxos/bgpGracefulRestart:BgpGracefulRestart":
		r = &BgpGracefulRestart{}
	case "nxos:nxos/bgpInstance:BgpInstance":
		r = &BgpInstance{}
	case "nxos:nxos/bgpPeer:BgpPeer":
		r = &BgpPeer{}
	case "nxos:nxos/bgpPeerAddressFamily:BgpPeerAddressFamily":
		r = &BgpPeerAddressFamily{}
	case "nxos:nxos/bgpPeerAddressFamilyPrefixListControl:BgpPeerAddressFamilyPrefixListControl":
		r = &BgpPeerAddressFamilyPrefixListControl{}
	case "nxos:nxos/bgpPeerAddressFamilyRouteControl:BgpPeerAddressFamilyRouteControl":
		r = &BgpPeerAddressFamilyRouteControl{}
	case "nxos:nxos/bgpPeerTemplate:BgpPeerTemplate":
		r = &BgpPeerTemplate{}
	case "nxos:nxos/bgpPeerTemplateAddressFamily:BgpPeerTemplateAddressFamily":
		r = &BgpPeerTemplateAddressFamily{}
	case "nxos:nxos/bgpPeerTemplateMaxPrefix:BgpPeerTemplateMaxPrefix":
		r = &BgpPeerTemplateMaxPrefix{}
	case "nxos:nxos/bgpRouteControl:BgpRouteControl":
		r = &BgpRouteControl{}
	case "nxos:nxos/bgpVrf:BgpVrf":
		r = &BgpVrf{}
	case "nxos:nxos/bridgeDomain:BridgeDomain":
		r = &BridgeDomain{}
	case "nxos:nxos/defaultQosClassMap:DefaultQosClassMap":
		r = &DefaultQosClassMap{}
	case "nxos:nxos/defaultQosClassMapDscp:DefaultQosClassMapDscp":
		r = &DefaultQosClassMapDscp{}
	case "nxos:nxos/defaultQosPolicyInterfaceIn:DefaultQosPolicyInterfaceIn":
		r = &DefaultQosPolicyInterfaceIn{}
	case "nxos:nxos/defaultQosPolicyInterfaceInPolicyMap:DefaultQosPolicyInterfaceInPolicyMap":
		r = &DefaultQosPolicyInterfaceInPolicyMap{}
	case "nxos:nxos/defaultQosPolicyMap:DefaultQosPolicyMap":
		r = &DefaultQosPolicyMap{}
	case "nxos:nxos/defaultQosPolicyMapMatchClassMap:DefaultQosPolicyMapMatchClassMap":
		r = &DefaultQosPolicyMapMatchClassMap{}
	case "nxos:nxos/defaultQosPolicyMapMatchClassMapPolice:DefaultQosPolicyMapMatchClassMapPolice":
		r = &DefaultQosPolicyMapMatchClassMapPolice{}
	case "nxos:nxos/defaultQosPolicyMapMatchClassMapSetQosGroup:DefaultQosPolicyMapMatchClassMapSetQosGroup":
		r = &DefaultQosPolicyMapMatchClassMapSetQosGroup{}
	case "nxos:nxos/dhcpRelayAddress:DhcpRelayAddress":
		r = &DhcpRelayAddress{}
	case "nxos:nxos/dhcpRelayInterface:DhcpRelayInterface":
		r = &DhcpRelayInterface{}
	case "nxos:nxos/ethernet:Ethernet":
		r = &Ethernet{}
	case "nxos:nxos/evpn:Evpn":
		r = &Evpn{}
	case "nxos:nxos/evpnVni:EvpnVni":
		r = &EvpnVni{}
	case "nxos:nxos/evpnVniRouteTarget:EvpnVniRouteTarget":
		r = &EvpnVniRouteTarget{}
	case "nxos:nxos/evpnVniRouteTargetDirection:EvpnVniRouteTargetDirection":
		r = &EvpnVniRouteTargetDirection{}
	case "nxos:nxos/featureBfd:FeatureBfd":
		r = &FeatureBfd{}
	case "nxos:nxos/featureBgp:FeatureBgp":
		r = &FeatureBgp{}
	case "nxos:nxos/featureDhcp:FeatureDhcp":
		r = &FeatureDhcp{}
	case "nxos:nxos/featureEvpn:FeatureEvpn":
		r = &FeatureEvpn{}
	case "nxos:nxos/featureHmm:FeatureHmm":
		r = &FeatureHmm{}
	case "nxos:nxos/featureHsrp:FeatureHsrp":
		r = &FeatureHsrp{}
	case "nxos:nxos/featureInterfaceVlan:FeatureInterfaceVlan":
		r = &FeatureInterfaceVlan{}
	case "nxos:nxos/featureIsis:FeatureIsis":
		r = &FeatureIsis{}
	case "nxos:nxos/featureLacp:FeatureLacp":
		r = &FeatureLacp{}
	case "nxos:nxos/featureLldp:FeatureLldp":
		r = &FeatureLldp{}
	case "nxos:nxos/featureMacsec:FeatureMacsec":
		r = &FeatureMacsec{}
	case "nxos:nxos/featureNetflow:FeatureNetflow":
		r = &FeatureNetflow{}
	case "nxos:nxos/featureNvOverlay:FeatureNvOverlay":
		r = &FeatureNvOverlay{}
	case "nxos:nxos/featureOspf:FeatureOspf":
		r = &FeatureOspf{}
	case "nxos:nxos/featureOspfv3:FeatureOspfv3":
		r = &FeatureOspfv3{}
	case "nxos:nxos/featurePim:FeaturePim":
		r = &FeaturePim{}
	case "nxos:nxos/featurePtp:FeaturePtp":
		r = &FeaturePtp{}
	case "nxos:nxos/featurePvlan:FeaturePvlan":
		r = &FeaturePvlan{}
	case "nxos:nxos/featureSsh:FeatureSsh":
		r = &FeatureSsh{}
	case "nxos:nxos/featureTacacs:FeatureTacacs":
		r = &FeatureTacacs{}
	case "nxos:nxos/featureTelnet:FeatureTelnet":
		r = &FeatureTelnet{}
	case "nxos:nxos/featureUdld:FeatureUdld":
		r = &FeatureUdld{}
	case "nxos:nxos/featureVnSegment:FeatureVnSegment":
		r = &FeatureVnSegment{}
	case "nxos:nxos/featureVpc:FeatureVpc":
		r = &FeatureVpc{}
	case "nxos:nxos/hmm:Hmm":
		r = &Hmm{}
	case "nxos:nxos/hmmInstance:HmmInstance":
		r = &HmmInstance{}
	case "nxos:nxos/hmmInterface:HmmInterface":
		r = &HmmInterface{}
	case "nxos:nxos/ipv4AccessList:Ipv4AccessList":
		r = &Ipv4AccessList{}
	case "nxos:nxos/ipv4AccessListEntry:Ipv4AccessListEntry":
		r = &Ipv4AccessListEntry{}
	case "nxos:nxos/ipv4AccessListPolicyEgressInterface:Ipv4AccessListPolicyEgressInterface":
		r = &Ipv4AccessListPolicyEgressInterface{}
	case "nxos:nxos/ipv4AccessListPolicyIngressInterface:Ipv4AccessListPolicyIngressInterface":
		r = &Ipv4AccessListPolicyIngressInterface{}
	case "nxos:nxos/ipv4Interface:Ipv4Interface":
		r = &Ipv4Interface{}
	case "nxos:nxos/ipv4InterfaceAddress:Ipv4InterfaceAddress":
		r = &Ipv4InterfaceAddress{}
	case "nxos:nxos/ipv4PrefixListRule:Ipv4PrefixListRule":
		r = &Ipv4PrefixListRule{}
	case "nxos:nxos/ipv4PrefixListRuleEntry:Ipv4PrefixListRuleEntry":
		r = &Ipv4PrefixListRuleEntry{}
	case "nxos:nxos/ipv4StaticRoute:Ipv4StaticRoute":
		r = &Ipv4StaticRoute{}
	case "nxos:nxos/ipv4Vrf:Ipv4Vrf":
		r = &Ipv4Vrf{}
	case "nxos:nxos/isis:Isis":
		r = &Isis{}
	case "nxos:nxos/isisInstance:IsisInstance":
		r = &IsisInstance{}
	case "nxos:nxos/isisInterface:IsisInterface":
		r = &IsisInterface{}
	case "nxos:nxos/isisVrf:IsisVrf":
		r = &IsisVrf{}
	case "nxos:nxos/loopbackInterface:LoopbackInterface":
		r = &LoopbackInterface{}
	case "nxos:nxos/loopbackInterfaceVrf:LoopbackInterfaceVrf":
		r = &LoopbackInterfaceVrf{}
	case "nxos:nxos/ntpServer:NtpServer":
		r = &NtpServer{}
	case "nxos:nxos/nveInterface:NveInterface":
		r = &NveInterface{}
	case "nxos:nxos/nveVni:NveVni":
		r = &NveVni{}
	case "nxos:nxos/nveVniContainer:NveVniContainer":
		r = &NveVniContainer{}
	case "nxos:nxos/nveVniIngressReplication:NveVniIngressReplication":
		r = &NveVniIngressReplication{}
	case "nxos:nxos/ospf:Ospf":
		r = &Ospf{}
	case "nxos:nxos/ospfArea:OspfArea":
		r = &OspfArea{}
	case "nxos:nxos/ospfAuthentication:OspfAuthentication":
		r = &OspfAuthentication{}
	case "nxos:nxos/ospfInstance:OspfInstance":
		r = &OspfInstance{}
	case "nxos:nxos/ospfInterface:OspfInterface":
		r = &OspfInterface{}
	case "nxos:nxos/ospfVrf:OspfVrf":
		r = &OspfVrf{}
	case "nxos:nxos/physicalInterface:PhysicalInterface":
		r = &PhysicalInterface{}
	case "nxos:nxos/physicalInterfaceVrf:PhysicalInterfaceVrf":
		r = &PhysicalInterfaceVrf{}
	case "nxos:nxos/pim:Pim":
		r = &Pim{}
	case "nxos:nxos/pimAnycastRp:PimAnycastRp":
		r = &PimAnycastRp{}
	case "nxos:nxos/pimAnycastRpPeer:PimAnycastRpPeer":
		r = &PimAnycastRpPeer{}
	case "nxos:nxos/pimInstance:PimInstance":
		r = &PimInstance{}
	case "nxos:nxos/pimInterface:PimInterface":
		r = &PimInterface{}
	case "nxos:nxos/pimSsmPolicy:PimSsmPolicy":
		r = &PimSsmPolicy{}
	case "nxos:nxos/pimSsmRange:PimSsmRange":
		r = &PimSsmRange{}
	case "nxos:nxos/pimStaticRp:PimStaticRp":
		r = &PimStaticRp{}
	case "nxos:nxos/pimStaticRpGroupList:PimStaticRpGroupList":
		r = &PimStaticRpGroupList{}
	case "nxos:nxos/pimStaticRpPolicy:PimStaticRpPolicy":
		r = &PimStaticRpPolicy{}
	case "nxos:nxos/pimVrf:PimVrf":
		r = &PimVrf{}
	case "nxos:nxos/portChannelInterface:PortChannelInterface":
		r = &PortChannelInterface{}
	case "nxos:nxos/portChannelInterfaceMember:PortChannelInterfaceMember":
		r = &PortChannelInterfaceMember{}
	case "nxos:nxos/portChannelInterfaceVrf:PortChannelInterfaceVrf":
		r = &PortChannelInterfaceVrf{}
	case "nxos:nxos/queuingQosPolicyMap:QueuingQosPolicyMap":
		r = &QueuingQosPolicyMap{}
	case "nxos:nxos/queuingQosPolicyMapMatchClassMap:QueuingQosPolicyMapMatchClassMap":
		r = &QueuingQosPolicyMapMatchClassMap{}
	case "nxos:nxos/queuingQosPolicyMapMatchClassMapPriority:QueuingQosPolicyMapMatchClassMapPriority":
		r = &QueuingQosPolicyMapMatchClassMapPriority{}
	case "nxos:nxos/queuingQosPolicyMapMatchClassMapRemainingBandwidth:QueuingQosPolicyMapMatchClassMapRemainingBandwidth":
		r = &QueuingQosPolicyMapMatchClassMapRemainingBandwidth{}
	case "nxos:nxos/queuingQosPolicySystemOut:QueuingQosPolicySystemOut":
		r = &QueuingQosPolicySystemOut{}
	case "nxos:nxos/queuingQosPolicySystemOutPolicyMap:QueuingQosPolicySystemOutPolicyMap":
		r = &QueuingQosPolicySystemOutPolicyMap{}
	case "nxos:nxos/rest:Rest":
		r = &Rest{}
	case "nxos:nxos/routeMapRule:RouteMapRule":
		r = &RouteMapRule{}
	case "nxos:nxos/routeMapRuleEntry:RouteMapRuleEntry":
		r = &RouteMapRuleEntry{}
	case "nxos:nxos/routeMapRuleEntryMatchRoute:RouteMapRuleEntryMatchRoute":
		r = &RouteMapRuleEntryMatchRoute{}
	case "nxos:nxos/routeMapRuleEntryMatchRoutePrefixList:RouteMapRuleEntryMatchRoutePrefixList":
		r = &RouteMapRuleEntryMatchRoutePrefixList{}
	case "nxos:nxos/routeMapRuleEntrySetRegularCommunity:RouteMapRuleEntrySetRegularCommunity":
		r = &RouteMapRuleEntrySetRegularCommunity{}
	case "nxos:nxos/routeMapRuleEntrySetRegularCommunityItem:RouteMapRuleEntrySetRegularCommunityItem":
		r = &RouteMapRuleEntrySetRegularCommunityItem{}
	case "nxos:nxos/spanningTreeInterface:SpanningTreeInterface":
		r = &SpanningTreeInterface{}
	case "nxos:nxos/subinterface:Subinterface":
		r = &Subinterface{}
	case "nxos:nxos/subinterfaceVrf:SubinterfaceVrf":
		r = &SubinterfaceVrf{}
	case "nxos:nxos/sviInterface:SviInterface":
		r = &SviInterface{}
	case "nxos:nxos/sviInterfaceVrf:SviInterfaceVrf":
		r = &SviInterfaceVrf{}
	case "nxos:nxos/system:System":
		r = &System{}
	case "nxos:nxos/vpcDomain:VpcDomain":
		r = &VpcDomain{}
	case "nxos:nxos/vpcInstance:VpcInstance":
		r = &VpcInstance{}
	case "nxos:nxos/vpcInterface:VpcInterface":
		r = &VpcInterface{}
	case "nxos:nxos/vrf:Vrf":
		r = &Vrf{}
	case "nxos:nxos/vrfAddressFamily:VrfAddressFamily":
		r = &VrfAddressFamily{}
	case "nxos:nxos/vrfRouteTarget:VrfRouteTarget":
		r = &VrfRouteTarget{}
	case "nxos:nxos/vrfRouteTargetAddressFamily:VrfRouteTargetAddressFamily":
		r = &VrfRouteTargetAddressFamily{}
	case "nxos:nxos/vrfRouteTargetDirection:VrfRouteTargetDirection":
		r = &VrfRouteTargetDirection{}
	case "nxos:nxos/vrfRouting:VrfRouting":
		r = &VrfRouting{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/bgp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/bgpAddressFamily",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/bgpAdvertisedPrefix",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/bgpGracefulRestart",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/bgpInstance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/bgpPeer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/bgpPeerAddressFamily",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/bgpPeerAddressFamilyPrefixListControl",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/bgpPeerAddressFamilyRouteControl",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/bgpPeerTemplate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/bgpPeerTemplateAddressFamily",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/bgpPeerTemplateMaxPrefix",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/bgpRouteControl",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/bgpVrf",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/bridgeDomain",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/defaultQosClassMap",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/defaultQosClassMapDscp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/defaultQosPolicyInterfaceIn",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/defaultQosPolicyInterfaceInPolicyMap",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/defaultQosPolicyMap",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/defaultQosPolicyMapMatchClassMap",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/defaultQosPolicyMapMatchClassMapPolice",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/defaultQosPolicyMapMatchClassMapSetQosGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/dhcpRelayAddress",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/dhcpRelayInterface",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/ethernet",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/evpn",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/evpnVni",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/evpnVniRouteTarget",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/evpnVniRouteTargetDirection",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/featureBfd",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/featureBgp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/featureDhcp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/featureEvpn",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/featureHmm",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/featureHsrp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/featureInterfaceVlan",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/featureIsis",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/featureLacp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/featureLldp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/featureMacsec",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/featureNetflow",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/featureNvOverlay",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/featureOspf",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/featureOspfv3",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/featurePim",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/featurePtp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/featurePvlan",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/featureSsh",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/featureTacacs",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/featureTelnet",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/featureUdld",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/featureVnSegment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/featureVpc",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/hmm",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/hmmInstance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/hmmInterface",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/ipv4AccessList",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/ipv4AccessListEntry",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/ipv4AccessListPolicyEgressInterface",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/ipv4AccessListPolicyIngressInterface",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/ipv4Interface",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/ipv4InterfaceAddress",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/ipv4PrefixListRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/ipv4PrefixListRuleEntry",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/ipv4StaticRoute",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/ipv4Vrf",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/isis",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/isisInstance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/isisInterface",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/isisVrf",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/loopbackInterface",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/loopbackInterfaceVrf",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/ntpServer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/nveInterface",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/nveVni",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/nveVniContainer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/nveVniIngressReplication",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/ospf",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/ospfArea",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/ospfAuthentication",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/ospfInstance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/ospfInterface",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/ospfVrf",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/physicalInterface",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/physicalInterfaceVrf",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/pim",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/pimAnycastRp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/pimAnycastRpPeer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/pimInstance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/pimInterface",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/pimSsmPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/pimSsmRange",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/pimStaticRp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/pimStaticRpGroupList",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/pimStaticRpPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/pimVrf",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/portChannelInterface",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/portChannelInterfaceMember",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/portChannelInterfaceVrf",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/queuingQosPolicyMap",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/queuingQosPolicyMapMatchClassMap",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/queuingQosPolicyMapMatchClassMapPriority",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/queuingQosPolicyMapMatchClassMapRemainingBandwidth",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/queuingQosPolicySystemOut",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/queuingQosPolicySystemOutPolicyMap",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/rest",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/routeMapRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/routeMapRuleEntry",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/routeMapRuleEntryMatchRoute",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/routeMapRuleEntryMatchRoutePrefixList",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/routeMapRuleEntrySetRegularCommunity",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/routeMapRuleEntrySetRegularCommunityItem",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/spanningTreeInterface",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/subinterface",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/subinterfaceVrf",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/sviInterface",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/sviInterfaceVrf",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/system",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/vpcDomain",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/vpcInstance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/vpcInterface",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/vrf",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/vrfAddressFamily",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/vrfRouteTarget",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/vrfRouteTargetAddressFamily",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/vrfRouteTargetDirection",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"nxos",
		"nxos/vrfRouting",
		&module{version},
	)
}