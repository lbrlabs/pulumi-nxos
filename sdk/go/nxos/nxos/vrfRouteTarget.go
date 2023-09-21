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

type VrfRouteTarget struct {
	pulumi.CustomResourceState

	// Address family. - Choices: `ipv4-ucast`, `ipv6-ucast`
	AddressFamily pulumi.StringOutput `pulumi:"addressFamily"`
	// A device name from the provider configuration.
	Device pulumi.StringPtrOutput `pulumi:"device"`
	// Route Target direction. - Choices: `import`, `export`
	Direction pulumi.StringOutput `pulumi:"direction"`
	// Route Target in NX-OS DME format.
	RouteTarget pulumi.StringOutput `pulumi:"routeTarget"`
	// Route Target Address Family. - Choices: `ipv4-ucast`, `ipv6-ucast`, `l2vpn-evpn`
	RouteTargetAddressFamily pulumi.StringOutput `pulumi:"routeTargetAddressFamily"`
	// VRF name.
	Vrf pulumi.StringOutput `pulumi:"vrf"`
}

// NewVrfRouteTarget registers a new resource with the given unique name, arguments, and options.
func NewVrfRouteTarget(ctx *pulumi.Context,
	name string, args *VrfRouteTargetArgs, opts ...pulumi.ResourceOption) (*VrfRouteTarget, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AddressFamily == nil {
		return nil, errors.New("invalid value for required argument 'AddressFamily'")
	}
	if args.Direction == nil {
		return nil, errors.New("invalid value for required argument 'Direction'")
	}
	if args.RouteTarget == nil {
		return nil, errors.New("invalid value for required argument 'RouteTarget'")
	}
	if args.RouteTargetAddressFamily == nil {
		return nil, errors.New("invalid value for required argument 'RouteTargetAddressFamily'")
	}
	if args.Vrf == nil {
		return nil, errors.New("invalid value for required argument 'Vrf'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VrfRouteTarget
	err := ctx.RegisterResource("nxos:nxos/vrfRouteTarget:VrfRouteTarget", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVrfRouteTarget gets an existing VrfRouteTarget resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVrfRouteTarget(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VrfRouteTargetState, opts ...pulumi.ResourceOption) (*VrfRouteTarget, error) {
	var resource VrfRouteTarget
	err := ctx.ReadResource("nxos:nxos/vrfRouteTarget:VrfRouteTarget", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VrfRouteTarget resources.
type vrfRouteTargetState struct {
	// Address family. - Choices: `ipv4-ucast`, `ipv6-ucast`
	AddressFamily *string `pulumi:"addressFamily"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// Route Target direction. - Choices: `import`, `export`
	Direction *string `pulumi:"direction"`
	// Route Target in NX-OS DME format.
	RouteTarget *string `pulumi:"routeTarget"`
	// Route Target Address Family. - Choices: `ipv4-ucast`, `ipv6-ucast`, `l2vpn-evpn`
	RouteTargetAddressFamily *string `pulumi:"routeTargetAddressFamily"`
	// VRF name.
	Vrf *string `pulumi:"vrf"`
}

type VrfRouteTargetState struct {
	// Address family. - Choices: `ipv4-ucast`, `ipv6-ucast`
	AddressFamily pulumi.StringPtrInput
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
	// Route Target direction. - Choices: `import`, `export`
	Direction pulumi.StringPtrInput
	// Route Target in NX-OS DME format.
	RouteTarget pulumi.StringPtrInput
	// Route Target Address Family. - Choices: `ipv4-ucast`, `ipv6-ucast`, `l2vpn-evpn`
	RouteTargetAddressFamily pulumi.StringPtrInput
	// VRF name.
	Vrf pulumi.StringPtrInput
}

func (VrfRouteTargetState) ElementType() reflect.Type {
	return reflect.TypeOf((*vrfRouteTargetState)(nil)).Elem()
}

type vrfRouteTargetArgs struct {
	// Address family. - Choices: `ipv4-ucast`, `ipv6-ucast`
	AddressFamily string `pulumi:"addressFamily"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// Route Target direction. - Choices: `import`, `export`
	Direction string `pulumi:"direction"`
	// Route Target in NX-OS DME format.
	RouteTarget string `pulumi:"routeTarget"`
	// Route Target Address Family. - Choices: `ipv4-ucast`, `ipv6-ucast`, `l2vpn-evpn`
	RouteTargetAddressFamily string `pulumi:"routeTargetAddressFamily"`
	// VRF name.
	Vrf string `pulumi:"vrf"`
}

// The set of arguments for constructing a VrfRouteTarget resource.
type VrfRouteTargetArgs struct {
	// Address family. - Choices: `ipv4-ucast`, `ipv6-ucast`
	AddressFamily pulumi.StringInput
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
	// Route Target direction. - Choices: `import`, `export`
	Direction pulumi.StringInput
	// Route Target in NX-OS DME format.
	RouteTarget pulumi.StringInput
	// Route Target Address Family. - Choices: `ipv4-ucast`, `ipv6-ucast`, `l2vpn-evpn`
	RouteTargetAddressFamily pulumi.StringInput
	// VRF name.
	Vrf pulumi.StringInput
}

func (VrfRouteTargetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vrfRouteTargetArgs)(nil)).Elem()
}

type VrfRouteTargetInput interface {
	pulumi.Input

	ToVrfRouteTargetOutput() VrfRouteTargetOutput
	ToVrfRouteTargetOutputWithContext(ctx context.Context) VrfRouteTargetOutput
}

func (*VrfRouteTarget) ElementType() reflect.Type {
	return reflect.TypeOf((**VrfRouteTarget)(nil)).Elem()
}

func (i *VrfRouteTarget) ToVrfRouteTargetOutput() VrfRouteTargetOutput {
	return i.ToVrfRouteTargetOutputWithContext(context.Background())
}

func (i *VrfRouteTarget) ToVrfRouteTargetOutputWithContext(ctx context.Context) VrfRouteTargetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VrfRouteTargetOutput)
}

func (i *VrfRouteTarget) ToOutput(ctx context.Context) pulumix.Output[*VrfRouteTarget] {
	return pulumix.Output[*VrfRouteTarget]{
		OutputState: i.ToVrfRouteTargetOutputWithContext(ctx).OutputState,
	}
}

// VrfRouteTargetArrayInput is an input type that accepts VrfRouteTargetArray and VrfRouteTargetArrayOutput values.
// You can construct a concrete instance of `VrfRouteTargetArrayInput` via:
//
//	VrfRouteTargetArray{ VrfRouteTargetArgs{...} }
type VrfRouteTargetArrayInput interface {
	pulumi.Input

	ToVrfRouteTargetArrayOutput() VrfRouteTargetArrayOutput
	ToVrfRouteTargetArrayOutputWithContext(context.Context) VrfRouteTargetArrayOutput
}

type VrfRouteTargetArray []VrfRouteTargetInput

func (VrfRouteTargetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VrfRouteTarget)(nil)).Elem()
}

func (i VrfRouteTargetArray) ToVrfRouteTargetArrayOutput() VrfRouteTargetArrayOutput {
	return i.ToVrfRouteTargetArrayOutputWithContext(context.Background())
}

func (i VrfRouteTargetArray) ToVrfRouteTargetArrayOutputWithContext(ctx context.Context) VrfRouteTargetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VrfRouteTargetArrayOutput)
}

func (i VrfRouteTargetArray) ToOutput(ctx context.Context) pulumix.Output[[]*VrfRouteTarget] {
	return pulumix.Output[[]*VrfRouteTarget]{
		OutputState: i.ToVrfRouteTargetArrayOutputWithContext(ctx).OutputState,
	}
}

// VrfRouteTargetMapInput is an input type that accepts VrfRouteTargetMap and VrfRouteTargetMapOutput values.
// You can construct a concrete instance of `VrfRouteTargetMapInput` via:
//
//	VrfRouteTargetMap{ "key": VrfRouteTargetArgs{...} }
type VrfRouteTargetMapInput interface {
	pulumi.Input

	ToVrfRouteTargetMapOutput() VrfRouteTargetMapOutput
	ToVrfRouteTargetMapOutputWithContext(context.Context) VrfRouteTargetMapOutput
}

type VrfRouteTargetMap map[string]VrfRouteTargetInput

func (VrfRouteTargetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VrfRouteTarget)(nil)).Elem()
}

func (i VrfRouteTargetMap) ToVrfRouteTargetMapOutput() VrfRouteTargetMapOutput {
	return i.ToVrfRouteTargetMapOutputWithContext(context.Background())
}

func (i VrfRouteTargetMap) ToVrfRouteTargetMapOutputWithContext(ctx context.Context) VrfRouteTargetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VrfRouteTargetMapOutput)
}

func (i VrfRouteTargetMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*VrfRouteTarget] {
	return pulumix.Output[map[string]*VrfRouteTarget]{
		OutputState: i.ToVrfRouteTargetMapOutputWithContext(ctx).OutputState,
	}
}

type VrfRouteTargetOutput struct{ *pulumi.OutputState }

func (VrfRouteTargetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VrfRouteTarget)(nil)).Elem()
}

func (o VrfRouteTargetOutput) ToVrfRouteTargetOutput() VrfRouteTargetOutput {
	return o
}

func (o VrfRouteTargetOutput) ToVrfRouteTargetOutputWithContext(ctx context.Context) VrfRouteTargetOutput {
	return o
}

func (o VrfRouteTargetOutput) ToOutput(ctx context.Context) pulumix.Output[*VrfRouteTarget] {
	return pulumix.Output[*VrfRouteTarget]{
		OutputState: o.OutputState,
	}
}

// Address family. - Choices: `ipv4-ucast`, `ipv6-ucast`
func (o VrfRouteTargetOutput) AddressFamily() pulumi.StringOutput {
	return o.ApplyT(func(v *VrfRouteTarget) pulumi.StringOutput { return v.AddressFamily }).(pulumi.StringOutput)
}

// A device name from the provider configuration.
func (o VrfRouteTargetOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VrfRouteTarget) pulumi.StringPtrOutput { return v.Device }).(pulumi.StringPtrOutput)
}

// Route Target direction. - Choices: `import`, `export`
func (o VrfRouteTargetOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v *VrfRouteTarget) pulumi.StringOutput { return v.Direction }).(pulumi.StringOutput)
}

// Route Target in NX-OS DME format.
func (o VrfRouteTargetOutput) RouteTarget() pulumi.StringOutput {
	return o.ApplyT(func(v *VrfRouteTarget) pulumi.StringOutput { return v.RouteTarget }).(pulumi.StringOutput)
}

// Route Target Address Family. - Choices: `ipv4-ucast`, `ipv6-ucast`, `l2vpn-evpn`
func (o VrfRouteTargetOutput) RouteTargetAddressFamily() pulumi.StringOutput {
	return o.ApplyT(func(v *VrfRouteTarget) pulumi.StringOutput { return v.RouteTargetAddressFamily }).(pulumi.StringOutput)
}

// VRF name.
func (o VrfRouteTargetOutput) Vrf() pulumi.StringOutput {
	return o.ApplyT(func(v *VrfRouteTarget) pulumi.StringOutput { return v.Vrf }).(pulumi.StringOutput)
}

type VrfRouteTargetArrayOutput struct{ *pulumi.OutputState }

func (VrfRouteTargetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VrfRouteTarget)(nil)).Elem()
}

func (o VrfRouteTargetArrayOutput) ToVrfRouteTargetArrayOutput() VrfRouteTargetArrayOutput {
	return o
}

func (o VrfRouteTargetArrayOutput) ToVrfRouteTargetArrayOutputWithContext(ctx context.Context) VrfRouteTargetArrayOutput {
	return o
}

func (o VrfRouteTargetArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*VrfRouteTarget] {
	return pulumix.Output[[]*VrfRouteTarget]{
		OutputState: o.OutputState,
	}
}

func (o VrfRouteTargetArrayOutput) Index(i pulumi.IntInput) VrfRouteTargetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VrfRouteTarget {
		return vs[0].([]*VrfRouteTarget)[vs[1].(int)]
	}).(VrfRouteTargetOutput)
}

type VrfRouteTargetMapOutput struct{ *pulumi.OutputState }

func (VrfRouteTargetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VrfRouteTarget)(nil)).Elem()
}

func (o VrfRouteTargetMapOutput) ToVrfRouteTargetMapOutput() VrfRouteTargetMapOutput {
	return o
}

func (o VrfRouteTargetMapOutput) ToVrfRouteTargetMapOutputWithContext(ctx context.Context) VrfRouteTargetMapOutput {
	return o
}

func (o VrfRouteTargetMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*VrfRouteTarget] {
	return pulumix.Output[map[string]*VrfRouteTarget]{
		OutputState: o.OutputState,
	}
}

func (o VrfRouteTargetMapOutput) MapIndex(k pulumi.StringInput) VrfRouteTargetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VrfRouteTarget {
		return vs[0].(map[string]*VrfRouteTarget)[vs[1].(string)]
	}).(VrfRouteTargetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VrfRouteTargetInput)(nil)).Elem(), &VrfRouteTarget{})
	pulumi.RegisterInputType(reflect.TypeOf((*VrfRouteTargetArrayInput)(nil)).Elem(), VrfRouteTargetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VrfRouteTargetMapInput)(nil)).Elem(), VrfRouteTargetMap{})
	pulumi.RegisterOutputType(VrfRouteTargetOutput{})
	pulumi.RegisterOutputType(VrfRouteTargetArrayOutput{})
	pulumi.RegisterOutputType(VrfRouteTargetMapOutput{})
}
