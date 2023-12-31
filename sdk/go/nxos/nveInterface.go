// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nxos

import (
	"context"
	"reflect"

	"github.com/lbrlabs/pulumi-nxos/sdk/go/nxos/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// This resource can manage the NVE interface configuration.
//
// - API Documentation: [nvoEp](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Network%20Virtualization/nvo:Ep/)
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
//			_, err := nxos.NewNveInterface(ctx, "example", &nxos.NveInterfaceArgs{
//				AdminState:                    pulumi.String("enabled"),
//				AdvertiseVirtualMac:           pulumi.Bool(true),
//				HoldDownTime:                  pulumi.Int(60),
//				HostReachabilityProtocol:      pulumi.String("bgp"),
//				IngressReplicationProtocolBgp: pulumi.Bool(true),
//				MulticastGroupL2:              pulumi.String("0.0.0.0"),
//				MulticastGroupL3:              pulumi.String("0.0.0.0"),
//				MultisiteSourceInterface:      pulumi.String("unspecified"),
//				SourceInterface:               pulumi.String("lo0"),
//				SuppressArp:                   pulumi.Bool(true),
//				SuppressMacRoute:              pulumi.Bool(false),
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
//	$ pulumi import nxos:index/nveInterface:NveInterface example "sys/eps/epId-[1]"
//
// ```
type NveInterface struct {
	pulumi.CustomResourceState

	// Administrative state. - Choices: `enabled`, `disabled` - Default value: `disabled`
	AdminState pulumi.StringOutput `pulumi:"adminState"`
	// Enable or disable Virtual MAC Advertisement in VPC mode. - Default value: `false`
	AdvertiseVirtualMac pulumi.BoolOutput `pulumi:"advertiseVirtualMac"`
	// A device name from the provider configuration.
	Device pulumi.StringPtrOutput `pulumi:"device"`
	// Hold Down Time. - Range: `1`-`1500` - Default value: `180`
	HoldDownTime pulumi.IntOutput `pulumi:"holdDownTime"`
	// Host Reachability Protocol. - Choices: `Flood-and-learn`, `bgp`, `controller`, `openflow`, `openflowIR` - Default value:
	// `Flood-and-learn`
	HostReachabilityProtocol pulumi.StringOutput `pulumi:"hostReachabilityProtocol"`
	// VxLAN Ingress Replication Protocol BGP. - Default value: `false`
	IngressReplicationProtocolBgp pulumi.BoolOutput `pulumi:"ingressReplicationProtocolBgp"`
	// Base multicast group address for L2. - Default value: `0.0.0.0`
	MulticastGroupL2 pulumi.StringOutput `pulumi:"multicastGroupL2"`
	// Base multicast group address for L3. - Default value: `0.0.0.0`
	MulticastGroupL3 pulumi.StringOutput `pulumi:"multicastGroupL3"`
	// Interface representing the Multisite Border Gateway. Must match first field in the output of `show int brief`. - Default
	// value: `unspecified`
	MultisiteSourceInterface pulumi.StringOutput `pulumi:"multisiteSourceInterface"`
	// Source Interface associated with the NVE. Must match first field in the output of `show int brief`. - Default value:
	// `unspecified`
	SourceInterface pulumi.StringOutput `pulumi:"sourceInterface"`
	// Suppress ARP. - Default value: `false`
	SuppressArp pulumi.BoolOutput `pulumi:"suppressArp"`
	// Suppress MAC Route. - Default value: `false`
	SuppressMacRoute pulumi.BoolOutput `pulumi:"suppressMacRoute"`
}

// NewNveInterface registers a new resource with the given unique name, arguments, and options.
func NewNveInterface(ctx *pulumi.Context,
	name string, args *NveInterfaceArgs, opts ...pulumi.ResourceOption) (*NveInterface, error) {
	if args == nil {
		args = &NveInterfaceArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NveInterface
	err := ctx.RegisterResource("nxos:index/nveInterface:NveInterface", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNveInterface gets an existing NveInterface resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNveInterface(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NveInterfaceState, opts ...pulumi.ResourceOption) (*NveInterface, error) {
	var resource NveInterface
	err := ctx.ReadResource("nxos:index/nveInterface:NveInterface", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NveInterface resources.
type nveInterfaceState struct {
	// Administrative state. - Choices: `enabled`, `disabled` - Default value: `disabled`
	AdminState *string `pulumi:"adminState"`
	// Enable or disable Virtual MAC Advertisement in VPC mode. - Default value: `false`
	AdvertiseVirtualMac *bool `pulumi:"advertiseVirtualMac"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// Hold Down Time. - Range: `1`-`1500` - Default value: `180`
	HoldDownTime *int `pulumi:"holdDownTime"`
	// Host Reachability Protocol. - Choices: `Flood-and-learn`, `bgp`, `controller`, `openflow`, `openflowIR` - Default value:
	// `Flood-and-learn`
	HostReachabilityProtocol *string `pulumi:"hostReachabilityProtocol"`
	// VxLAN Ingress Replication Protocol BGP. - Default value: `false`
	IngressReplicationProtocolBgp *bool `pulumi:"ingressReplicationProtocolBgp"`
	// Base multicast group address for L2. - Default value: `0.0.0.0`
	MulticastGroupL2 *string `pulumi:"multicastGroupL2"`
	// Base multicast group address for L3. - Default value: `0.0.0.0`
	MulticastGroupL3 *string `pulumi:"multicastGroupL3"`
	// Interface representing the Multisite Border Gateway. Must match first field in the output of `show int brief`. - Default
	// value: `unspecified`
	MultisiteSourceInterface *string `pulumi:"multisiteSourceInterface"`
	// Source Interface associated with the NVE. Must match first field in the output of `show int brief`. - Default value:
	// `unspecified`
	SourceInterface *string `pulumi:"sourceInterface"`
	// Suppress ARP. - Default value: `false`
	SuppressArp *bool `pulumi:"suppressArp"`
	// Suppress MAC Route. - Default value: `false`
	SuppressMacRoute *bool `pulumi:"suppressMacRoute"`
}

type NveInterfaceState struct {
	// Administrative state. - Choices: `enabled`, `disabled` - Default value: `disabled`
	AdminState pulumi.StringPtrInput
	// Enable or disable Virtual MAC Advertisement in VPC mode. - Default value: `false`
	AdvertiseVirtualMac pulumi.BoolPtrInput
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
	// Hold Down Time. - Range: `1`-`1500` - Default value: `180`
	HoldDownTime pulumi.IntPtrInput
	// Host Reachability Protocol. - Choices: `Flood-and-learn`, `bgp`, `controller`, `openflow`, `openflowIR` - Default value:
	// `Flood-and-learn`
	HostReachabilityProtocol pulumi.StringPtrInput
	// VxLAN Ingress Replication Protocol BGP. - Default value: `false`
	IngressReplicationProtocolBgp pulumi.BoolPtrInput
	// Base multicast group address for L2. - Default value: `0.0.0.0`
	MulticastGroupL2 pulumi.StringPtrInput
	// Base multicast group address for L3. - Default value: `0.0.0.0`
	MulticastGroupL3 pulumi.StringPtrInput
	// Interface representing the Multisite Border Gateway. Must match first field in the output of `show int brief`. - Default
	// value: `unspecified`
	MultisiteSourceInterface pulumi.StringPtrInput
	// Source Interface associated with the NVE. Must match first field in the output of `show int brief`. - Default value:
	// `unspecified`
	SourceInterface pulumi.StringPtrInput
	// Suppress ARP. - Default value: `false`
	SuppressArp pulumi.BoolPtrInput
	// Suppress MAC Route. - Default value: `false`
	SuppressMacRoute pulumi.BoolPtrInput
}

func (NveInterfaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*nveInterfaceState)(nil)).Elem()
}

type nveInterfaceArgs struct {
	// Administrative state. - Choices: `enabled`, `disabled` - Default value: `disabled`
	AdminState *string `pulumi:"adminState"`
	// Enable or disable Virtual MAC Advertisement in VPC mode. - Default value: `false`
	AdvertiseVirtualMac *bool `pulumi:"advertiseVirtualMac"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// Hold Down Time. - Range: `1`-`1500` - Default value: `180`
	HoldDownTime *int `pulumi:"holdDownTime"`
	// Host Reachability Protocol. - Choices: `Flood-and-learn`, `bgp`, `controller`, `openflow`, `openflowIR` - Default value:
	// `Flood-and-learn`
	HostReachabilityProtocol *string `pulumi:"hostReachabilityProtocol"`
	// VxLAN Ingress Replication Protocol BGP. - Default value: `false`
	IngressReplicationProtocolBgp *bool `pulumi:"ingressReplicationProtocolBgp"`
	// Base multicast group address for L2. - Default value: `0.0.0.0`
	MulticastGroupL2 *string `pulumi:"multicastGroupL2"`
	// Base multicast group address for L3. - Default value: `0.0.0.0`
	MulticastGroupL3 *string `pulumi:"multicastGroupL3"`
	// Interface representing the Multisite Border Gateway. Must match first field in the output of `show int brief`. - Default
	// value: `unspecified`
	MultisiteSourceInterface *string `pulumi:"multisiteSourceInterface"`
	// Source Interface associated with the NVE. Must match first field in the output of `show int brief`. - Default value:
	// `unspecified`
	SourceInterface *string `pulumi:"sourceInterface"`
	// Suppress ARP. - Default value: `false`
	SuppressArp *bool `pulumi:"suppressArp"`
	// Suppress MAC Route. - Default value: `false`
	SuppressMacRoute *bool `pulumi:"suppressMacRoute"`
}

// The set of arguments for constructing a NveInterface resource.
type NveInterfaceArgs struct {
	// Administrative state. - Choices: `enabled`, `disabled` - Default value: `disabled`
	AdminState pulumi.StringPtrInput
	// Enable or disable Virtual MAC Advertisement in VPC mode. - Default value: `false`
	AdvertiseVirtualMac pulumi.BoolPtrInput
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
	// Hold Down Time. - Range: `1`-`1500` - Default value: `180`
	HoldDownTime pulumi.IntPtrInput
	// Host Reachability Protocol. - Choices: `Flood-and-learn`, `bgp`, `controller`, `openflow`, `openflowIR` - Default value:
	// `Flood-and-learn`
	HostReachabilityProtocol pulumi.StringPtrInput
	// VxLAN Ingress Replication Protocol BGP. - Default value: `false`
	IngressReplicationProtocolBgp pulumi.BoolPtrInput
	// Base multicast group address for L2. - Default value: `0.0.0.0`
	MulticastGroupL2 pulumi.StringPtrInput
	// Base multicast group address for L3. - Default value: `0.0.0.0`
	MulticastGroupL3 pulumi.StringPtrInput
	// Interface representing the Multisite Border Gateway. Must match first field in the output of `show int brief`. - Default
	// value: `unspecified`
	MultisiteSourceInterface pulumi.StringPtrInput
	// Source Interface associated with the NVE. Must match first field in the output of `show int brief`. - Default value:
	// `unspecified`
	SourceInterface pulumi.StringPtrInput
	// Suppress ARP. - Default value: `false`
	SuppressArp pulumi.BoolPtrInput
	// Suppress MAC Route. - Default value: `false`
	SuppressMacRoute pulumi.BoolPtrInput
}

func (NveInterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nveInterfaceArgs)(nil)).Elem()
}

type NveInterfaceInput interface {
	pulumi.Input

	ToNveInterfaceOutput() NveInterfaceOutput
	ToNveInterfaceOutputWithContext(ctx context.Context) NveInterfaceOutput
}

func (*NveInterface) ElementType() reflect.Type {
	return reflect.TypeOf((**NveInterface)(nil)).Elem()
}

func (i *NveInterface) ToNveInterfaceOutput() NveInterfaceOutput {
	return i.ToNveInterfaceOutputWithContext(context.Background())
}

func (i *NveInterface) ToNveInterfaceOutputWithContext(ctx context.Context) NveInterfaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NveInterfaceOutput)
}

func (i *NveInterface) ToOutput(ctx context.Context) pulumix.Output[*NveInterface] {
	return pulumix.Output[*NveInterface]{
		OutputState: i.ToNveInterfaceOutputWithContext(ctx).OutputState,
	}
}

// NveInterfaceArrayInput is an input type that accepts NveInterfaceArray and NveInterfaceArrayOutput values.
// You can construct a concrete instance of `NveInterfaceArrayInput` via:
//
//	NveInterfaceArray{ NveInterfaceArgs{...} }
type NveInterfaceArrayInput interface {
	pulumi.Input

	ToNveInterfaceArrayOutput() NveInterfaceArrayOutput
	ToNveInterfaceArrayOutputWithContext(context.Context) NveInterfaceArrayOutput
}

type NveInterfaceArray []NveInterfaceInput

func (NveInterfaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NveInterface)(nil)).Elem()
}

func (i NveInterfaceArray) ToNveInterfaceArrayOutput() NveInterfaceArrayOutput {
	return i.ToNveInterfaceArrayOutputWithContext(context.Background())
}

func (i NveInterfaceArray) ToNveInterfaceArrayOutputWithContext(ctx context.Context) NveInterfaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NveInterfaceArrayOutput)
}

func (i NveInterfaceArray) ToOutput(ctx context.Context) pulumix.Output[[]*NveInterface] {
	return pulumix.Output[[]*NveInterface]{
		OutputState: i.ToNveInterfaceArrayOutputWithContext(ctx).OutputState,
	}
}

// NveInterfaceMapInput is an input type that accepts NveInterfaceMap and NveInterfaceMapOutput values.
// You can construct a concrete instance of `NveInterfaceMapInput` via:
//
//	NveInterfaceMap{ "key": NveInterfaceArgs{...} }
type NveInterfaceMapInput interface {
	pulumi.Input

	ToNveInterfaceMapOutput() NveInterfaceMapOutput
	ToNveInterfaceMapOutputWithContext(context.Context) NveInterfaceMapOutput
}

type NveInterfaceMap map[string]NveInterfaceInput

func (NveInterfaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NveInterface)(nil)).Elem()
}

func (i NveInterfaceMap) ToNveInterfaceMapOutput() NveInterfaceMapOutput {
	return i.ToNveInterfaceMapOutputWithContext(context.Background())
}

func (i NveInterfaceMap) ToNveInterfaceMapOutputWithContext(ctx context.Context) NveInterfaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NveInterfaceMapOutput)
}

func (i NveInterfaceMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*NveInterface] {
	return pulumix.Output[map[string]*NveInterface]{
		OutputState: i.ToNveInterfaceMapOutputWithContext(ctx).OutputState,
	}
}

type NveInterfaceOutput struct{ *pulumi.OutputState }

func (NveInterfaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NveInterface)(nil)).Elem()
}

func (o NveInterfaceOutput) ToNveInterfaceOutput() NveInterfaceOutput {
	return o
}

func (o NveInterfaceOutput) ToNveInterfaceOutputWithContext(ctx context.Context) NveInterfaceOutput {
	return o
}

func (o NveInterfaceOutput) ToOutput(ctx context.Context) pulumix.Output[*NveInterface] {
	return pulumix.Output[*NveInterface]{
		OutputState: o.OutputState,
	}
}

// Administrative state. - Choices: `enabled`, `disabled` - Default value: `disabled`
func (o NveInterfaceOutput) AdminState() pulumi.StringOutput {
	return o.ApplyT(func(v *NveInterface) pulumi.StringOutput { return v.AdminState }).(pulumi.StringOutput)
}

// Enable or disable Virtual MAC Advertisement in VPC mode. - Default value: `false`
func (o NveInterfaceOutput) AdvertiseVirtualMac() pulumi.BoolOutput {
	return o.ApplyT(func(v *NveInterface) pulumi.BoolOutput { return v.AdvertiseVirtualMac }).(pulumi.BoolOutput)
}

// A device name from the provider configuration.
func (o NveInterfaceOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NveInterface) pulumi.StringPtrOutput { return v.Device }).(pulumi.StringPtrOutput)
}

// Hold Down Time. - Range: `1`-`1500` - Default value: `180`
func (o NveInterfaceOutput) HoldDownTime() pulumi.IntOutput {
	return o.ApplyT(func(v *NveInterface) pulumi.IntOutput { return v.HoldDownTime }).(pulumi.IntOutput)
}

// Host Reachability Protocol. - Choices: `Flood-and-learn`, `bgp`, `controller`, `openflow`, `openflowIR` - Default value:
// `Flood-and-learn`
func (o NveInterfaceOutput) HostReachabilityProtocol() pulumi.StringOutput {
	return o.ApplyT(func(v *NveInterface) pulumi.StringOutput { return v.HostReachabilityProtocol }).(pulumi.StringOutput)
}

// VxLAN Ingress Replication Protocol BGP. - Default value: `false`
func (o NveInterfaceOutput) IngressReplicationProtocolBgp() pulumi.BoolOutput {
	return o.ApplyT(func(v *NveInterface) pulumi.BoolOutput { return v.IngressReplicationProtocolBgp }).(pulumi.BoolOutput)
}

// Base multicast group address for L2. - Default value: `0.0.0.0`
func (o NveInterfaceOutput) MulticastGroupL2() pulumi.StringOutput {
	return o.ApplyT(func(v *NveInterface) pulumi.StringOutput { return v.MulticastGroupL2 }).(pulumi.StringOutput)
}

// Base multicast group address for L3. - Default value: `0.0.0.0`
func (o NveInterfaceOutput) MulticastGroupL3() pulumi.StringOutput {
	return o.ApplyT(func(v *NveInterface) pulumi.StringOutput { return v.MulticastGroupL3 }).(pulumi.StringOutput)
}

// Interface representing the Multisite Border Gateway. Must match first field in the output of `show int brief`. - Default
// value: `unspecified`
func (o NveInterfaceOutput) MultisiteSourceInterface() pulumi.StringOutput {
	return o.ApplyT(func(v *NveInterface) pulumi.StringOutput { return v.MultisiteSourceInterface }).(pulumi.StringOutput)
}

// Source Interface associated with the NVE. Must match first field in the output of `show int brief`. - Default value:
// `unspecified`
func (o NveInterfaceOutput) SourceInterface() pulumi.StringOutput {
	return o.ApplyT(func(v *NveInterface) pulumi.StringOutput { return v.SourceInterface }).(pulumi.StringOutput)
}

// Suppress ARP. - Default value: `false`
func (o NveInterfaceOutput) SuppressArp() pulumi.BoolOutput {
	return o.ApplyT(func(v *NveInterface) pulumi.BoolOutput { return v.SuppressArp }).(pulumi.BoolOutput)
}

// Suppress MAC Route. - Default value: `false`
func (o NveInterfaceOutput) SuppressMacRoute() pulumi.BoolOutput {
	return o.ApplyT(func(v *NveInterface) pulumi.BoolOutput { return v.SuppressMacRoute }).(pulumi.BoolOutput)
}

type NveInterfaceArrayOutput struct{ *pulumi.OutputState }

func (NveInterfaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NveInterface)(nil)).Elem()
}

func (o NveInterfaceArrayOutput) ToNveInterfaceArrayOutput() NveInterfaceArrayOutput {
	return o
}

func (o NveInterfaceArrayOutput) ToNveInterfaceArrayOutputWithContext(ctx context.Context) NveInterfaceArrayOutput {
	return o
}

func (o NveInterfaceArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*NveInterface] {
	return pulumix.Output[[]*NveInterface]{
		OutputState: o.OutputState,
	}
}

func (o NveInterfaceArrayOutput) Index(i pulumi.IntInput) NveInterfaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NveInterface {
		return vs[0].([]*NveInterface)[vs[1].(int)]
	}).(NveInterfaceOutput)
}

type NveInterfaceMapOutput struct{ *pulumi.OutputState }

func (NveInterfaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NveInterface)(nil)).Elem()
}

func (o NveInterfaceMapOutput) ToNveInterfaceMapOutput() NveInterfaceMapOutput {
	return o
}

func (o NveInterfaceMapOutput) ToNveInterfaceMapOutputWithContext(ctx context.Context) NveInterfaceMapOutput {
	return o
}

func (o NveInterfaceMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*NveInterface] {
	return pulumix.Output[map[string]*NveInterface]{
		OutputState: o.OutputState,
	}
}

func (o NveInterfaceMapOutput) MapIndex(k pulumi.StringInput) NveInterfaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NveInterface {
		return vs[0].(map[string]*NveInterface)[vs[1].(string)]
	}).(NveInterfaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NveInterfaceInput)(nil)).Elem(), &NveInterface{})
	pulumi.RegisterInputType(reflect.TypeOf((*NveInterfaceArrayInput)(nil)).Elem(), NveInterfaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NveInterfaceMapInput)(nil)).Elem(), NveInterfaceMap{})
	pulumi.RegisterOutputType(NveInterfaceOutput{})
	pulumi.RegisterOutputType(NveInterfaceArrayOutput{})
	pulumi.RegisterOutputType(NveInterfaceMapOutput{})
}
