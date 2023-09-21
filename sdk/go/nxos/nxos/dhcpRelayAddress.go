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

type DhcpRelayAddress struct {
	pulumi.CustomResourceState

	// IPv4 or IPv6 address.
	Address pulumi.StringOutput `pulumi:"address"`
	// A device name from the provider configuration.
	Device pulumi.StringPtrOutput `pulumi:"device"`
	// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
	InterfaceId pulumi.StringOutput `pulumi:"interfaceId"`
	// VRF name.
	Vrf pulumi.StringOutput `pulumi:"vrf"`
}

// NewDhcpRelayAddress registers a new resource with the given unique name, arguments, and options.
func NewDhcpRelayAddress(ctx *pulumi.Context,
	name string, args *DhcpRelayAddressArgs, opts ...pulumi.ResourceOption) (*DhcpRelayAddress, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Address == nil {
		return nil, errors.New("invalid value for required argument 'Address'")
	}
	if args.InterfaceId == nil {
		return nil, errors.New("invalid value for required argument 'InterfaceId'")
	}
	if args.Vrf == nil {
		return nil, errors.New("invalid value for required argument 'Vrf'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DhcpRelayAddress
	err := ctx.RegisterResource("nxos:nxos/dhcpRelayAddress:DhcpRelayAddress", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDhcpRelayAddress gets an existing DhcpRelayAddress resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDhcpRelayAddress(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DhcpRelayAddressState, opts ...pulumi.ResourceOption) (*DhcpRelayAddress, error) {
	var resource DhcpRelayAddress
	err := ctx.ReadResource("nxos:nxos/dhcpRelayAddress:DhcpRelayAddress", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DhcpRelayAddress resources.
type dhcpRelayAddressState struct {
	// IPv4 or IPv6 address.
	Address *string `pulumi:"address"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
	InterfaceId *string `pulumi:"interfaceId"`
	// VRF name.
	Vrf *string `pulumi:"vrf"`
}

type DhcpRelayAddressState struct {
	// IPv4 or IPv6 address.
	Address pulumi.StringPtrInput
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
	// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
	InterfaceId pulumi.StringPtrInput
	// VRF name.
	Vrf pulumi.StringPtrInput
}

func (DhcpRelayAddressState) ElementType() reflect.Type {
	return reflect.TypeOf((*dhcpRelayAddressState)(nil)).Elem()
}

type dhcpRelayAddressArgs struct {
	// IPv4 or IPv6 address.
	Address string `pulumi:"address"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
	InterfaceId string `pulumi:"interfaceId"`
	// VRF name.
	Vrf string `pulumi:"vrf"`
}

// The set of arguments for constructing a DhcpRelayAddress resource.
type DhcpRelayAddressArgs struct {
	// IPv4 or IPv6 address.
	Address pulumi.StringInput
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
	// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
	InterfaceId pulumi.StringInput
	// VRF name.
	Vrf pulumi.StringInput
}

func (DhcpRelayAddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dhcpRelayAddressArgs)(nil)).Elem()
}

type DhcpRelayAddressInput interface {
	pulumi.Input

	ToDhcpRelayAddressOutput() DhcpRelayAddressOutput
	ToDhcpRelayAddressOutputWithContext(ctx context.Context) DhcpRelayAddressOutput
}

func (*DhcpRelayAddress) ElementType() reflect.Type {
	return reflect.TypeOf((**DhcpRelayAddress)(nil)).Elem()
}

func (i *DhcpRelayAddress) ToDhcpRelayAddressOutput() DhcpRelayAddressOutput {
	return i.ToDhcpRelayAddressOutputWithContext(context.Background())
}

func (i *DhcpRelayAddress) ToDhcpRelayAddressOutputWithContext(ctx context.Context) DhcpRelayAddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DhcpRelayAddressOutput)
}

func (i *DhcpRelayAddress) ToOutput(ctx context.Context) pulumix.Output[*DhcpRelayAddress] {
	return pulumix.Output[*DhcpRelayAddress]{
		OutputState: i.ToDhcpRelayAddressOutputWithContext(ctx).OutputState,
	}
}

// DhcpRelayAddressArrayInput is an input type that accepts DhcpRelayAddressArray and DhcpRelayAddressArrayOutput values.
// You can construct a concrete instance of `DhcpRelayAddressArrayInput` via:
//
//	DhcpRelayAddressArray{ DhcpRelayAddressArgs{...} }
type DhcpRelayAddressArrayInput interface {
	pulumi.Input

	ToDhcpRelayAddressArrayOutput() DhcpRelayAddressArrayOutput
	ToDhcpRelayAddressArrayOutputWithContext(context.Context) DhcpRelayAddressArrayOutput
}

type DhcpRelayAddressArray []DhcpRelayAddressInput

func (DhcpRelayAddressArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DhcpRelayAddress)(nil)).Elem()
}

func (i DhcpRelayAddressArray) ToDhcpRelayAddressArrayOutput() DhcpRelayAddressArrayOutput {
	return i.ToDhcpRelayAddressArrayOutputWithContext(context.Background())
}

func (i DhcpRelayAddressArray) ToDhcpRelayAddressArrayOutputWithContext(ctx context.Context) DhcpRelayAddressArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DhcpRelayAddressArrayOutput)
}

func (i DhcpRelayAddressArray) ToOutput(ctx context.Context) pulumix.Output[[]*DhcpRelayAddress] {
	return pulumix.Output[[]*DhcpRelayAddress]{
		OutputState: i.ToDhcpRelayAddressArrayOutputWithContext(ctx).OutputState,
	}
}

// DhcpRelayAddressMapInput is an input type that accepts DhcpRelayAddressMap and DhcpRelayAddressMapOutput values.
// You can construct a concrete instance of `DhcpRelayAddressMapInput` via:
//
//	DhcpRelayAddressMap{ "key": DhcpRelayAddressArgs{...} }
type DhcpRelayAddressMapInput interface {
	pulumi.Input

	ToDhcpRelayAddressMapOutput() DhcpRelayAddressMapOutput
	ToDhcpRelayAddressMapOutputWithContext(context.Context) DhcpRelayAddressMapOutput
}

type DhcpRelayAddressMap map[string]DhcpRelayAddressInput

func (DhcpRelayAddressMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DhcpRelayAddress)(nil)).Elem()
}

func (i DhcpRelayAddressMap) ToDhcpRelayAddressMapOutput() DhcpRelayAddressMapOutput {
	return i.ToDhcpRelayAddressMapOutputWithContext(context.Background())
}

func (i DhcpRelayAddressMap) ToDhcpRelayAddressMapOutputWithContext(ctx context.Context) DhcpRelayAddressMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DhcpRelayAddressMapOutput)
}

func (i DhcpRelayAddressMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*DhcpRelayAddress] {
	return pulumix.Output[map[string]*DhcpRelayAddress]{
		OutputState: i.ToDhcpRelayAddressMapOutputWithContext(ctx).OutputState,
	}
}

type DhcpRelayAddressOutput struct{ *pulumi.OutputState }

func (DhcpRelayAddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DhcpRelayAddress)(nil)).Elem()
}

func (o DhcpRelayAddressOutput) ToDhcpRelayAddressOutput() DhcpRelayAddressOutput {
	return o
}

func (o DhcpRelayAddressOutput) ToDhcpRelayAddressOutputWithContext(ctx context.Context) DhcpRelayAddressOutput {
	return o
}

func (o DhcpRelayAddressOutput) ToOutput(ctx context.Context) pulumix.Output[*DhcpRelayAddress] {
	return pulumix.Output[*DhcpRelayAddress]{
		OutputState: o.OutputState,
	}
}

// IPv4 or IPv6 address.
func (o DhcpRelayAddressOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v *DhcpRelayAddress) pulumi.StringOutput { return v.Address }).(pulumi.StringOutput)
}

// A device name from the provider configuration.
func (o DhcpRelayAddressOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DhcpRelayAddress) pulumi.StringPtrOutput { return v.Device }).(pulumi.StringPtrOutput)
}

// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
func (o DhcpRelayAddressOutput) InterfaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DhcpRelayAddress) pulumi.StringOutput { return v.InterfaceId }).(pulumi.StringOutput)
}

// VRF name.
func (o DhcpRelayAddressOutput) Vrf() pulumi.StringOutput {
	return o.ApplyT(func(v *DhcpRelayAddress) pulumi.StringOutput { return v.Vrf }).(pulumi.StringOutput)
}

type DhcpRelayAddressArrayOutput struct{ *pulumi.OutputState }

func (DhcpRelayAddressArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DhcpRelayAddress)(nil)).Elem()
}

func (o DhcpRelayAddressArrayOutput) ToDhcpRelayAddressArrayOutput() DhcpRelayAddressArrayOutput {
	return o
}

func (o DhcpRelayAddressArrayOutput) ToDhcpRelayAddressArrayOutputWithContext(ctx context.Context) DhcpRelayAddressArrayOutput {
	return o
}

func (o DhcpRelayAddressArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*DhcpRelayAddress] {
	return pulumix.Output[[]*DhcpRelayAddress]{
		OutputState: o.OutputState,
	}
}

func (o DhcpRelayAddressArrayOutput) Index(i pulumi.IntInput) DhcpRelayAddressOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DhcpRelayAddress {
		return vs[0].([]*DhcpRelayAddress)[vs[1].(int)]
	}).(DhcpRelayAddressOutput)
}

type DhcpRelayAddressMapOutput struct{ *pulumi.OutputState }

func (DhcpRelayAddressMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DhcpRelayAddress)(nil)).Elem()
}

func (o DhcpRelayAddressMapOutput) ToDhcpRelayAddressMapOutput() DhcpRelayAddressMapOutput {
	return o
}

func (o DhcpRelayAddressMapOutput) ToDhcpRelayAddressMapOutputWithContext(ctx context.Context) DhcpRelayAddressMapOutput {
	return o
}

func (o DhcpRelayAddressMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*DhcpRelayAddress] {
	return pulumix.Output[map[string]*DhcpRelayAddress]{
		OutputState: o.OutputState,
	}
}

func (o DhcpRelayAddressMapOutput) MapIndex(k pulumi.StringInput) DhcpRelayAddressOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DhcpRelayAddress {
		return vs[0].(map[string]*DhcpRelayAddress)[vs[1].(string)]
	}).(DhcpRelayAddressOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DhcpRelayAddressInput)(nil)).Elem(), &DhcpRelayAddress{})
	pulumi.RegisterInputType(reflect.TypeOf((*DhcpRelayAddressArrayInput)(nil)).Elem(), DhcpRelayAddressArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DhcpRelayAddressMapInput)(nil)).Elem(), DhcpRelayAddressMap{})
	pulumi.RegisterOutputType(DhcpRelayAddressOutput{})
	pulumi.RegisterOutputType(DhcpRelayAddressArrayOutput{})
	pulumi.RegisterOutputType(DhcpRelayAddressMapOutput{})
}
