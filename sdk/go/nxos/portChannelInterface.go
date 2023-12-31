// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nxos

import (
	"context"
	"reflect"

	"errors"
	"github.com/lbrlabs/pulumi-nxos/sdk/go/nxos/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// This resource can manage a port-channel interface.
//
// - API Documentation: [pcAggrIf](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Interfaces/pc:AggrIf/)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/lbrlabs/pulumi-nxos/sdk/go/nxos"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := nxos.NewPortChannelInterface(ctx, "example", &nxos.PortChannelInterfaceArgs{
//				AccessVlan:          pulumi.String("vlan-1"),
//				AdminState:          pulumi.String("up"),
//				AutoNegotiation:     pulumi.String("on"),
//				Bandwidth:           pulumi.Int(0),
//				Delay:               pulumi.Int(1),
//				Description:         pulumi.String("My Description"),
//				Duplex:              pulumi.String("auto"),
//				InterfaceId:         pulumi.String("po1"),
//				Layer:               pulumi.String("Layer2"),
//				LinkLogging:         pulumi.String("enable"),
//				MaximumLinks:        pulumi.Int(10),
//				Medium:              pulumi.String("broadcast"),
//				MinimumLinks:        pulumi.Int(2),
//				Mode:                pulumi.String("access"),
//				Mtu:                 pulumi.Int(1500),
//				NativeVlan:          pulumi.String("unknown"),
//				PortChannelMode:     pulumi.String("active"),
//				Speed:               pulumi.String("auto"),
//				SuspendIndividual:   pulumi.String("disable"),
//				TrunkVlans:          pulumi.String("1-4094"),
//				UserConfiguredFlags: pulumi.String("admin_layer,admin_mtu,admin_state"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// ```sh
//
//	$ pulumi import nxos:index/portChannelInterface:PortChannelInterface example "sys/intf/aggr-[po1]"
//
// ```
type PortChannelInterface struct {
	pulumi.CustomResourceState

	// Access VLAN. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`. - Default value: `vlan-1`
	AccessVlan pulumi.StringOutput `pulumi:"accessVlan"`
	// Administrative port state. - Choices: `up`, `down` - Default value: `up`
	AdminState pulumi.StringOutput `pulumi:"adminState"`
	// Administrative port auto-negotiation. - Choices: `on`, `off`, `25G` - Default value: `on`
	AutoNegotiation pulumi.StringOutput `pulumi:"autoNegotiation"`
	// The bandwidth parameter for a routed interface, port channel, or subinterface. - Range: `0`-`3200000000` - Default
	// value: `0`
	Bandwidth pulumi.IntOutput `pulumi:"bandwidth"`
	// The administrative port delay time. - Range: `1`-`16777215` - Default value: `1`
	Delay pulumi.IntOutput `pulumi:"delay"`
	// Interface description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A device name from the provider configuration.
	Device pulumi.StringPtrOutput `pulumi:"device"`
	// Duplex. - Choices: `auto`, `full`, `half` - Default value: `auto`
	Duplex pulumi.StringOutput `pulumi:"duplex"`
	// Must match first field in the output of `show intf brief`. Example: `po1`.
	InterfaceId pulumi.StringOutput `pulumi:"interfaceId"`
	// Administrative port layer. - Choices: `Layer2`, `Layer3` - Default value: `Layer2`
	Layer pulumi.StringOutput `pulumi:"layer"`
	// Administrative link logging. - Choices: `default`, `enable`, `disable` - Default value: `default`
	LinkLogging pulumi.StringOutput `pulumi:"linkLogging"`
	// Maximum links. - Range: `1`-`32` - Default value: `32`
	MaximumLinks pulumi.IntOutput `pulumi:"maximumLinks"`
	// The administrative port medium type. - Choices: `broadcast`, `p2p` - Default value: `broadcast`
	Medium pulumi.StringOutput `pulumi:"medium"`
	// Minimum links. - Range: `1`-`32` - Default value: `1`
	MinimumLinks pulumi.IntOutput `pulumi:"minimumLinks"`
	// Administrative port mode. - Choices: `access`, `trunk`, `fex-fabric`, `dot1q-tunnel`, `promiscuous`, `host`,
	// `trunk_secondary`, `trunk_promiscuous`, `vntag` - Default value: `access`
	Mode pulumi.StringOutput `pulumi:"mode"`
	// Administrative port MTU. - Range: `576`-`9216` - Default value: `1500`
	Mtu pulumi.IntOutput `pulumi:"mtu"`
	// Native VLAN. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`. - Default value: `vlan-1`
	NativeVlan pulumi.StringOutput `pulumi:"nativeVlan"`
	// The aggregated interface protocol channel mode. - Choices: `on`, `static`, `active`, `passive`, `mac-pin` - Default
	// value: `on`
	PortChannelMode pulumi.StringOutput `pulumi:"portChannelMode"`
	// Administrative port speed. - Choices: `unknown`, `100M`, `1G`, `10G`, `40G`, `auto`, `auto 100M`, `auto 100M 1G`,
	// `100G`, `25G`, `10M`, `50G`, `200G`, `400G`, `2.5G`, `5G`, `auto 2.5G 5G 10G`, `auto 100M 1G 2.5G 5G` - Default value:
	// `auto`
	Speed pulumi.StringOutput `pulumi:"speed"`
	// Suspend Individual Port. - Choices: `enable`, `disable` - Default value: `enable`
	SuspendIndividual pulumi.StringOutput `pulumi:"suspendIndividual"`
	// List of trunk VLANs. - Default value: `1-4094`
	TrunkVlans pulumi.StringOutput `pulumi:"trunkVlans"`
	// Port User Config Flags.
	UserConfiguredFlags pulumi.StringPtrOutput `pulumi:"userConfiguredFlags"`
}

// NewPortChannelInterface registers a new resource with the given unique name, arguments, and options.
func NewPortChannelInterface(ctx *pulumi.Context,
	name string, args *PortChannelInterfaceArgs, opts ...pulumi.ResourceOption) (*PortChannelInterface, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InterfaceId == nil {
		return nil, errors.New("invalid value for required argument 'InterfaceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PortChannelInterface
	err := ctx.RegisterResource("nxos:index/portChannelInterface:PortChannelInterface", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPortChannelInterface gets an existing PortChannelInterface resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPortChannelInterface(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PortChannelInterfaceState, opts ...pulumi.ResourceOption) (*PortChannelInterface, error) {
	var resource PortChannelInterface
	err := ctx.ReadResource("nxos:index/portChannelInterface:PortChannelInterface", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PortChannelInterface resources.
type portChannelInterfaceState struct {
	// Access VLAN. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`. - Default value: `vlan-1`
	AccessVlan *string `pulumi:"accessVlan"`
	// Administrative port state. - Choices: `up`, `down` - Default value: `up`
	AdminState *string `pulumi:"adminState"`
	// Administrative port auto-negotiation. - Choices: `on`, `off`, `25G` - Default value: `on`
	AutoNegotiation *string `pulumi:"autoNegotiation"`
	// The bandwidth parameter for a routed interface, port channel, or subinterface. - Range: `0`-`3200000000` - Default
	// value: `0`
	Bandwidth *int `pulumi:"bandwidth"`
	// The administrative port delay time. - Range: `1`-`16777215` - Default value: `1`
	Delay *int `pulumi:"delay"`
	// Interface description.
	Description *string `pulumi:"description"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// Duplex. - Choices: `auto`, `full`, `half` - Default value: `auto`
	Duplex *string `pulumi:"duplex"`
	// Must match first field in the output of `show intf brief`. Example: `po1`.
	InterfaceId *string `pulumi:"interfaceId"`
	// Administrative port layer. - Choices: `Layer2`, `Layer3` - Default value: `Layer2`
	Layer *string `pulumi:"layer"`
	// Administrative link logging. - Choices: `default`, `enable`, `disable` - Default value: `default`
	LinkLogging *string `pulumi:"linkLogging"`
	// Maximum links. - Range: `1`-`32` - Default value: `32`
	MaximumLinks *int `pulumi:"maximumLinks"`
	// The administrative port medium type. - Choices: `broadcast`, `p2p` - Default value: `broadcast`
	Medium *string `pulumi:"medium"`
	// Minimum links. - Range: `1`-`32` - Default value: `1`
	MinimumLinks *int `pulumi:"minimumLinks"`
	// Administrative port mode. - Choices: `access`, `trunk`, `fex-fabric`, `dot1q-tunnel`, `promiscuous`, `host`,
	// `trunk_secondary`, `trunk_promiscuous`, `vntag` - Default value: `access`
	Mode *string `pulumi:"mode"`
	// Administrative port MTU. - Range: `576`-`9216` - Default value: `1500`
	Mtu *int `pulumi:"mtu"`
	// Native VLAN. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`. - Default value: `vlan-1`
	NativeVlan *string `pulumi:"nativeVlan"`
	// The aggregated interface protocol channel mode. - Choices: `on`, `static`, `active`, `passive`, `mac-pin` - Default
	// value: `on`
	PortChannelMode *string `pulumi:"portChannelMode"`
	// Administrative port speed. - Choices: `unknown`, `100M`, `1G`, `10G`, `40G`, `auto`, `auto 100M`, `auto 100M 1G`,
	// `100G`, `25G`, `10M`, `50G`, `200G`, `400G`, `2.5G`, `5G`, `auto 2.5G 5G 10G`, `auto 100M 1G 2.5G 5G` - Default value:
	// `auto`
	Speed *string `pulumi:"speed"`
	// Suspend Individual Port. - Choices: `enable`, `disable` - Default value: `enable`
	SuspendIndividual *string `pulumi:"suspendIndividual"`
	// List of trunk VLANs. - Default value: `1-4094`
	TrunkVlans *string `pulumi:"trunkVlans"`
	// Port User Config Flags.
	UserConfiguredFlags *string `pulumi:"userConfiguredFlags"`
}

type PortChannelInterfaceState struct {
	// Access VLAN. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`. - Default value: `vlan-1`
	AccessVlan pulumi.StringPtrInput
	// Administrative port state. - Choices: `up`, `down` - Default value: `up`
	AdminState pulumi.StringPtrInput
	// Administrative port auto-negotiation. - Choices: `on`, `off`, `25G` - Default value: `on`
	AutoNegotiation pulumi.StringPtrInput
	// The bandwidth parameter for a routed interface, port channel, or subinterface. - Range: `0`-`3200000000` - Default
	// value: `0`
	Bandwidth pulumi.IntPtrInput
	// The administrative port delay time. - Range: `1`-`16777215` - Default value: `1`
	Delay pulumi.IntPtrInput
	// Interface description.
	Description pulumi.StringPtrInput
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
	// Duplex. - Choices: `auto`, `full`, `half` - Default value: `auto`
	Duplex pulumi.StringPtrInput
	// Must match first field in the output of `show intf brief`. Example: `po1`.
	InterfaceId pulumi.StringPtrInput
	// Administrative port layer. - Choices: `Layer2`, `Layer3` - Default value: `Layer2`
	Layer pulumi.StringPtrInput
	// Administrative link logging. - Choices: `default`, `enable`, `disable` - Default value: `default`
	LinkLogging pulumi.StringPtrInput
	// Maximum links. - Range: `1`-`32` - Default value: `32`
	MaximumLinks pulumi.IntPtrInput
	// The administrative port medium type. - Choices: `broadcast`, `p2p` - Default value: `broadcast`
	Medium pulumi.StringPtrInput
	// Minimum links. - Range: `1`-`32` - Default value: `1`
	MinimumLinks pulumi.IntPtrInput
	// Administrative port mode. - Choices: `access`, `trunk`, `fex-fabric`, `dot1q-tunnel`, `promiscuous`, `host`,
	// `trunk_secondary`, `trunk_promiscuous`, `vntag` - Default value: `access`
	Mode pulumi.StringPtrInput
	// Administrative port MTU. - Range: `576`-`9216` - Default value: `1500`
	Mtu pulumi.IntPtrInput
	// Native VLAN. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`. - Default value: `vlan-1`
	NativeVlan pulumi.StringPtrInput
	// The aggregated interface protocol channel mode. - Choices: `on`, `static`, `active`, `passive`, `mac-pin` - Default
	// value: `on`
	PortChannelMode pulumi.StringPtrInput
	// Administrative port speed. - Choices: `unknown`, `100M`, `1G`, `10G`, `40G`, `auto`, `auto 100M`, `auto 100M 1G`,
	// `100G`, `25G`, `10M`, `50G`, `200G`, `400G`, `2.5G`, `5G`, `auto 2.5G 5G 10G`, `auto 100M 1G 2.5G 5G` - Default value:
	// `auto`
	Speed pulumi.StringPtrInput
	// Suspend Individual Port. - Choices: `enable`, `disable` - Default value: `enable`
	SuspendIndividual pulumi.StringPtrInput
	// List of trunk VLANs. - Default value: `1-4094`
	TrunkVlans pulumi.StringPtrInput
	// Port User Config Flags.
	UserConfiguredFlags pulumi.StringPtrInput
}

func (PortChannelInterfaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*portChannelInterfaceState)(nil)).Elem()
}

type portChannelInterfaceArgs struct {
	// Access VLAN. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`. - Default value: `vlan-1`
	AccessVlan *string `pulumi:"accessVlan"`
	// Administrative port state. - Choices: `up`, `down` - Default value: `up`
	AdminState *string `pulumi:"adminState"`
	// Administrative port auto-negotiation. - Choices: `on`, `off`, `25G` - Default value: `on`
	AutoNegotiation *string `pulumi:"autoNegotiation"`
	// The bandwidth parameter for a routed interface, port channel, or subinterface. - Range: `0`-`3200000000` - Default
	// value: `0`
	Bandwidth *int `pulumi:"bandwidth"`
	// The administrative port delay time. - Range: `1`-`16777215` - Default value: `1`
	Delay *int `pulumi:"delay"`
	// Interface description.
	Description *string `pulumi:"description"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// Duplex. - Choices: `auto`, `full`, `half` - Default value: `auto`
	Duplex *string `pulumi:"duplex"`
	// Must match first field in the output of `show intf brief`. Example: `po1`.
	InterfaceId string `pulumi:"interfaceId"`
	// Administrative port layer. - Choices: `Layer2`, `Layer3` - Default value: `Layer2`
	Layer *string `pulumi:"layer"`
	// Administrative link logging. - Choices: `default`, `enable`, `disable` - Default value: `default`
	LinkLogging *string `pulumi:"linkLogging"`
	// Maximum links. - Range: `1`-`32` - Default value: `32`
	MaximumLinks *int `pulumi:"maximumLinks"`
	// The administrative port medium type. - Choices: `broadcast`, `p2p` - Default value: `broadcast`
	Medium *string `pulumi:"medium"`
	// Minimum links. - Range: `1`-`32` - Default value: `1`
	MinimumLinks *int `pulumi:"minimumLinks"`
	// Administrative port mode. - Choices: `access`, `trunk`, `fex-fabric`, `dot1q-tunnel`, `promiscuous`, `host`,
	// `trunk_secondary`, `trunk_promiscuous`, `vntag` - Default value: `access`
	Mode *string `pulumi:"mode"`
	// Administrative port MTU. - Range: `576`-`9216` - Default value: `1500`
	Mtu *int `pulumi:"mtu"`
	// Native VLAN. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`. - Default value: `vlan-1`
	NativeVlan *string `pulumi:"nativeVlan"`
	// The aggregated interface protocol channel mode. - Choices: `on`, `static`, `active`, `passive`, `mac-pin` - Default
	// value: `on`
	PortChannelMode *string `pulumi:"portChannelMode"`
	// Administrative port speed. - Choices: `unknown`, `100M`, `1G`, `10G`, `40G`, `auto`, `auto 100M`, `auto 100M 1G`,
	// `100G`, `25G`, `10M`, `50G`, `200G`, `400G`, `2.5G`, `5G`, `auto 2.5G 5G 10G`, `auto 100M 1G 2.5G 5G` - Default value:
	// `auto`
	Speed *string `pulumi:"speed"`
	// Suspend Individual Port. - Choices: `enable`, `disable` - Default value: `enable`
	SuspendIndividual *string `pulumi:"suspendIndividual"`
	// List of trunk VLANs. - Default value: `1-4094`
	TrunkVlans *string `pulumi:"trunkVlans"`
	// Port User Config Flags.
	UserConfiguredFlags *string `pulumi:"userConfiguredFlags"`
}

// The set of arguments for constructing a PortChannelInterface resource.
type PortChannelInterfaceArgs struct {
	// Access VLAN. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`. - Default value: `vlan-1`
	AccessVlan pulumi.StringPtrInput
	// Administrative port state. - Choices: `up`, `down` - Default value: `up`
	AdminState pulumi.StringPtrInput
	// Administrative port auto-negotiation. - Choices: `on`, `off`, `25G` - Default value: `on`
	AutoNegotiation pulumi.StringPtrInput
	// The bandwidth parameter for a routed interface, port channel, or subinterface. - Range: `0`-`3200000000` - Default
	// value: `0`
	Bandwidth pulumi.IntPtrInput
	// The administrative port delay time. - Range: `1`-`16777215` - Default value: `1`
	Delay pulumi.IntPtrInput
	// Interface description.
	Description pulumi.StringPtrInput
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
	// Duplex. - Choices: `auto`, `full`, `half` - Default value: `auto`
	Duplex pulumi.StringPtrInput
	// Must match first field in the output of `show intf brief`. Example: `po1`.
	InterfaceId pulumi.StringInput
	// Administrative port layer. - Choices: `Layer2`, `Layer3` - Default value: `Layer2`
	Layer pulumi.StringPtrInput
	// Administrative link logging. - Choices: `default`, `enable`, `disable` - Default value: `default`
	LinkLogging pulumi.StringPtrInput
	// Maximum links. - Range: `1`-`32` - Default value: `32`
	MaximumLinks pulumi.IntPtrInput
	// The administrative port medium type. - Choices: `broadcast`, `p2p` - Default value: `broadcast`
	Medium pulumi.StringPtrInput
	// Minimum links. - Range: `1`-`32` - Default value: `1`
	MinimumLinks pulumi.IntPtrInput
	// Administrative port mode. - Choices: `access`, `trunk`, `fex-fabric`, `dot1q-tunnel`, `promiscuous`, `host`,
	// `trunk_secondary`, `trunk_promiscuous`, `vntag` - Default value: `access`
	Mode pulumi.StringPtrInput
	// Administrative port MTU. - Range: `576`-`9216` - Default value: `1500`
	Mtu pulumi.IntPtrInput
	// Native VLAN. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`. - Default value: `vlan-1`
	NativeVlan pulumi.StringPtrInput
	// The aggregated interface protocol channel mode. - Choices: `on`, `static`, `active`, `passive`, `mac-pin` - Default
	// value: `on`
	PortChannelMode pulumi.StringPtrInput
	// Administrative port speed. - Choices: `unknown`, `100M`, `1G`, `10G`, `40G`, `auto`, `auto 100M`, `auto 100M 1G`,
	// `100G`, `25G`, `10M`, `50G`, `200G`, `400G`, `2.5G`, `5G`, `auto 2.5G 5G 10G`, `auto 100M 1G 2.5G 5G` - Default value:
	// `auto`
	Speed pulumi.StringPtrInput
	// Suspend Individual Port. - Choices: `enable`, `disable` - Default value: `enable`
	SuspendIndividual pulumi.StringPtrInput
	// List of trunk VLANs. - Default value: `1-4094`
	TrunkVlans pulumi.StringPtrInput
	// Port User Config Flags.
	UserConfiguredFlags pulumi.StringPtrInput
}

func (PortChannelInterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*portChannelInterfaceArgs)(nil)).Elem()
}

type PortChannelInterfaceInput interface {
	pulumi.Input

	ToPortChannelInterfaceOutput() PortChannelInterfaceOutput
	ToPortChannelInterfaceOutputWithContext(ctx context.Context) PortChannelInterfaceOutput
}

func (*PortChannelInterface) ElementType() reflect.Type {
	return reflect.TypeOf((**PortChannelInterface)(nil)).Elem()
}

func (i *PortChannelInterface) ToPortChannelInterfaceOutput() PortChannelInterfaceOutput {
	return i.ToPortChannelInterfaceOutputWithContext(context.Background())
}

func (i *PortChannelInterface) ToPortChannelInterfaceOutputWithContext(ctx context.Context) PortChannelInterfaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortChannelInterfaceOutput)
}

func (i *PortChannelInterface) ToOutput(ctx context.Context) pulumix.Output[*PortChannelInterface] {
	return pulumix.Output[*PortChannelInterface]{
		OutputState: i.ToPortChannelInterfaceOutputWithContext(ctx).OutputState,
	}
}

// PortChannelInterfaceArrayInput is an input type that accepts PortChannelInterfaceArray and PortChannelInterfaceArrayOutput values.
// You can construct a concrete instance of `PortChannelInterfaceArrayInput` via:
//
//	PortChannelInterfaceArray{ PortChannelInterfaceArgs{...} }
type PortChannelInterfaceArrayInput interface {
	pulumi.Input

	ToPortChannelInterfaceArrayOutput() PortChannelInterfaceArrayOutput
	ToPortChannelInterfaceArrayOutputWithContext(context.Context) PortChannelInterfaceArrayOutput
}

type PortChannelInterfaceArray []PortChannelInterfaceInput

func (PortChannelInterfaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PortChannelInterface)(nil)).Elem()
}

func (i PortChannelInterfaceArray) ToPortChannelInterfaceArrayOutput() PortChannelInterfaceArrayOutput {
	return i.ToPortChannelInterfaceArrayOutputWithContext(context.Background())
}

func (i PortChannelInterfaceArray) ToPortChannelInterfaceArrayOutputWithContext(ctx context.Context) PortChannelInterfaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortChannelInterfaceArrayOutput)
}

func (i PortChannelInterfaceArray) ToOutput(ctx context.Context) pulumix.Output[[]*PortChannelInterface] {
	return pulumix.Output[[]*PortChannelInterface]{
		OutputState: i.ToPortChannelInterfaceArrayOutputWithContext(ctx).OutputState,
	}
}

// PortChannelInterfaceMapInput is an input type that accepts PortChannelInterfaceMap and PortChannelInterfaceMapOutput values.
// You can construct a concrete instance of `PortChannelInterfaceMapInput` via:
//
//	PortChannelInterfaceMap{ "key": PortChannelInterfaceArgs{...} }
type PortChannelInterfaceMapInput interface {
	pulumi.Input

	ToPortChannelInterfaceMapOutput() PortChannelInterfaceMapOutput
	ToPortChannelInterfaceMapOutputWithContext(context.Context) PortChannelInterfaceMapOutput
}

type PortChannelInterfaceMap map[string]PortChannelInterfaceInput

func (PortChannelInterfaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PortChannelInterface)(nil)).Elem()
}

func (i PortChannelInterfaceMap) ToPortChannelInterfaceMapOutput() PortChannelInterfaceMapOutput {
	return i.ToPortChannelInterfaceMapOutputWithContext(context.Background())
}

func (i PortChannelInterfaceMap) ToPortChannelInterfaceMapOutputWithContext(ctx context.Context) PortChannelInterfaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortChannelInterfaceMapOutput)
}

func (i PortChannelInterfaceMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*PortChannelInterface] {
	return pulumix.Output[map[string]*PortChannelInterface]{
		OutputState: i.ToPortChannelInterfaceMapOutputWithContext(ctx).OutputState,
	}
}

type PortChannelInterfaceOutput struct{ *pulumi.OutputState }

func (PortChannelInterfaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PortChannelInterface)(nil)).Elem()
}

func (o PortChannelInterfaceOutput) ToPortChannelInterfaceOutput() PortChannelInterfaceOutput {
	return o
}

func (o PortChannelInterfaceOutput) ToPortChannelInterfaceOutputWithContext(ctx context.Context) PortChannelInterfaceOutput {
	return o
}

func (o PortChannelInterfaceOutput) ToOutput(ctx context.Context) pulumix.Output[*PortChannelInterface] {
	return pulumix.Output[*PortChannelInterface]{
		OutputState: o.OutputState,
	}
}

// Access VLAN. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`. - Default value: `vlan-1`
func (o PortChannelInterfaceOutput) AccessVlan() pulumi.StringOutput {
	return o.ApplyT(func(v *PortChannelInterface) pulumi.StringOutput { return v.AccessVlan }).(pulumi.StringOutput)
}

// Administrative port state. - Choices: `up`, `down` - Default value: `up`
func (o PortChannelInterfaceOutput) AdminState() pulumi.StringOutput {
	return o.ApplyT(func(v *PortChannelInterface) pulumi.StringOutput { return v.AdminState }).(pulumi.StringOutput)
}

// Administrative port auto-negotiation. - Choices: `on`, `off`, `25G` - Default value: `on`
func (o PortChannelInterfaceOutput) AutoNegotiation() pulumi.StringOutput {
	return o.ApplyT(func(v *PortChannelInterface) pulumi.StringOutput { return v.AutoNegotiation }).(pulumi.StringOutput)
}

// The bandwidth parameter for a routed interface, port channel, or subinterface. - Range: `0`-`3200000000` - Default
// value: `0`
func (o PortChannelInterfaceOutput) Bandwidth() pulumi.IntOutput {
	return o.ApplyT(func(v *PortChannelInterface) pulumi.IntOutput { return v.Bandwidth }).(pulumi.IntOutput)
}

// The administrative port delay time. - Range: `1`-`16777215` - Default value: `1`
func (o PortChannelInterfaceOutput) Delay() pulumi.IntOutput {
	return o.ApplyT(func(v *PortChannelInterface) pulumi.IntOutput { return v.Delay }).(pulumi.IntOutput)
}

// Interface description.
func (o PortChannelInterfaceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PortChannelInterface) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// A device name from the provider configuration.
func (o PortChannelInterfaceOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PortChannelInterface) pulumi.StringPtrOutput { return v.Device }).(pulumi.StringPtrOutput)
}

// Duplex. - Choices: `auto`, `full`, `half` - Default value: `auto`
func (o PortChannelInterfaceOutput) Duplex() pulumi.StringOutput {
	return o.ApplyT(func(v *PortChannelInterface) pulumi.StringOutput { return v.Duplex }).(pulumi.StringOutput)
}

// Must match first field in the output of `show intf brief`. Example: `po1`.
func (o PortChannelInterfaceOutput) InterfaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *PortChannelInterface) pulumi.StringOutput { return v.InterfaceId }).(pulumi.StringOutput)
}

// Administrative port layer. - Choices: `Layer2`, `Layer3` - Default value: `Layer2`
func (o PortChannelInterfaceOutput) Layer() pulumi.StringOutput {
	return o.ApplyT(func(v *PortChannelInterface) pulumi.StringOutput { return v.Layer }).(pulumi.StringOutput)
}

// Administrative link logging. - Choices: `default`, `enable`, `disable` - Default value: `default`
func (o PortChannelInterfaceOutput) LinkLogging() pulumi.StringOutput {
	return o.ApplyT(func(v *PortChannelInterface) pulumi.StringOutput { return v.LinkLogging }).(pulumi.StringOutput)
}

// Maximum links. - Range: `1`-`32` - Default value: `32`
func (o PortChannelInterfaceOutput) MaximumLinks() pulumi.IntOutput {
	return o.ApplyT(func(v *PortChannelInterface) pulumi.IntOutput { return v.MaximumLinks }).(pulumi.IntOutput)
}

// The administrative port medium type. - Choices: `broadcast`, `p2p` - Default value: `broadcast`
func (o PortChannelInterfaceOutput) Medium() pulumi.StringOutput {
	return o.ApplyT(func(v *PortChannelInterface) pulumi.StringOutput { return v.Medium }).(pulumi.StringOutput)
}

// Minimum links. - Range: `1`-`32` - Default value: `1`
func (o PortChannelInterfaceOutput) MinimumLinks() pulumi.IntOutput {
	return o.ApplyT(func(v *PortChannelInterface) pulumi.IntOutput { return v.MinimumLinks }).(pulumi.IntOutput)
}

// Administrative port mode. - Choices: `access`, `trunk`, `fex-fabric`, `dot1q-tunnel`, `promiscuous`, `host`,
// `trunk_secondary`, `trunk_promiscuous`, `vntag` - Default value: `access`
func (o PortChannelInterfaceOutput) Mode() pulumi.StringOutput {
	return o.ApplyT(func(v *PortChannelInterface) pulumi.StringOutput { return v.Mode }).(pulumi.StringOutput)
}

// Administrative port MTU. - Range: `576`-`9216` - Default value: `1500`
func (o PortChannelInterfaceOutput) Mtu() pulumi.IntOutput {
	return o.ApplyT(func(v *PortChannelInterface) pulumi.IntOutput { return v.Mtu }).(pulumi.IntOutput)
}

// Native VLAN. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`. - Default value: `vlan-1`
func (o PortChannelInterfaceOutput) NativeVlan() pulumi.StringOutput {
	return o.ApplyT(func(v *PortChannelInterface) pulumi.StringOutput { return v.NativeVlan }).(pulumi.StringOutput)
}

// The aggregated interface protocol channel mode. - Choices: `on`, `static`, `active`, `passive`, `mac-pin` - Default
// value: `on`
func (o PortChannelInterfaceOutput) PortChannelMode() pulumi.StringOutput {
	return o.ApplyT(func(v *PortChannelInterface) pulumi.StringOutput { return v.PortChannelMode }).(pulumi.StringOutput)
}

// Administrative port speed. - Choices: `unknown`, `100M`, `1G`, `10G`, `40G`, `auto`, `auto 100M`, `auto 100M 1G`,
// `100G`, `25G`, `10M`, `50G`, `200G`, `400G`, `2.5G`, `5G`, `auto 2.5G 5G 10G`, `auto 100M 1G 2.5G 5G` - Default value:
// `auto`
func (o PortChannelInterfaceOutput) Speed() pulumi.StringOutput {
	return o.ApplyT(func(v *PortChannelInterface) pulumi.StringOutput { return v.Speed }).(pulumi.StringOutput)
}

// Suspend Individual Port. - Choices: `enable`, `disable` - Default value: `enable`
func (o PortChannelInterfaceOutput) SuspendIndividual() pulumi.StringOutput {
	return o.ApplyT(func(v *PortChannelInterface) pulumi.StringOutput { return v.SuspendIndividual }).(pulumi.StringOutput)
}

// List of trunk VLANs. - Default value: `1-4094`
func (o PortChannelInterfaceOutput) TrunkVlans() pulumi.StringOutput {
	return o.ApplyT(func(v *PortChannelInterface) pulumi.StringOutput { return v.TrunkVlans }).(pulumi.StringOutput)
}

// Port User Config Flags.
func (o PortChannelInterfaceOutput) UserConfiguredFlags() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PortChannelInterface) pulumi.StringPtrOutput { return v.UserConfiguredFlags }).(pulumi.StringPtrOutput)
}

type PortChannelInterfaceArrayOutput struct{ *pulumi.OutputState }

func (PortChannelInterfaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PortChannelInterface)(nil)).Elem()
}

func (o PortChannelInterfaceArrayOutput) ToPortChannelInterfaceArrayOutput() PortChannelInterfaceArrayOutput {
	return o
}

func (o PortChannelInterfaceArrayOutput) ToPortChannelInterfaceArrayOutputWithContext(ctx context.Context) PortChannelInterfaceArrayOutput {
	return o
}

func (o PortChannelInterfaceArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*PortChannelInterface] {
	return pulumix.Output[[]*PortChannelInterface]{
		OutputState: o.OutputState,
	}
}

func (o PortChannelInterfaceArrayOutput) Index(i pulumi.IntInput) PortChannelInterfaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PortChannelInterface {
		return vs[0].([]*PortChannelInterface)[vs[1].(int)]
	}).(PortChannelInterfaceOutput)
}

type PortChannelInterfaceMapOutput struct{ *pulumi.OutputState }

func (PortChannelInterfaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PortChannelInterface)(nil)).Elem()
}

func (o PortChannelInterfaceMapOutput) ToPortChannelInterfaceMapOutput() PortChannelInterfaceMapOutput {
	return o
}

func (o PortChannelInterfaceMapOutput) ToPortChannelInterfaceMapOutputWithContext(ctx context.Context) PortChannelInterfaceMapOutput {
	return o
}

func (o PortChannelInterfaceMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*PortChannelInterface] {
	return pulumix.Output[map[string]*PortChannelInterface]{
		OutputState: o.OutputState,
	}
}

func (o PortChannelInterfaceMapOutput) MapIndex(k pulumi.StringInput) PortChannelInterfaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PortChannelInterface {
		return vs[0].(map[string]*PortChannelInterface)[vs[1].(string)]
	}).(PortChannelInterfaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PortChannelInterfaceInput)(nil)).Elem(), &PortChannelInterface{})
	pulumi.RegisterInputType(reflect.TypeOf((*PortChannelInterfaceArrayInput)(nil)).Elem(), PortChannelInterfaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PortChannelInterfaceMapInput)(nil)).Elem(), PortChannelInterfaceMap{})
	pulumi.RegisterOutputType(PortChannelInterfaceOutput{})
	pulumi.RegisterOutputType(PortChannelInterfaceArrayOutput{})
	pulumi.RegisterOutputType(PortChannelInterfaceMapOutput{})
}
