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

type Ipv4Vrf struct {
	pulumi.CustomResourceState

	// A device name from the provider configuration.
	Device pulumi.StringPtrOutput `pulumi:"device"`
	// VRF name.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewIpv4Vrf registers a new resource with the given unique name, arguments, and options.
func NewIpv4Vrf(ctx *pulumi.Context,
	name string, args *Ipv4VrfArgs, opts ...pulumi.ResourceOption) (*Ipv4Vrf, error) {
	if args == nil {
		args = &Ipv4VrfArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Ipv4Vrf
	err := ctx.RegisterResource("nxos:nxos/ipv4Vrf:Ipv4Vrf", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpv4Vrf gets an existing Ipv4Vrf resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpv4Vrf(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Ipv4VrfState, opts ...pulumi.ResourceOption) (*Ipv4Vrf, error) {
	var resource Ipv4Vrf
	err := ctx.ReadResource("nxos:nxos/ipv4Vrf:Ipv4Vrf", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ipv4Vrf resources.
type ipv4VrfState struct {
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// VRF name.
	Name *string `pulumi:"name"`
}

type Ipv4VrfState struct {
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
	// VRF name.
	Name pulumi.StringPtrInput
}

func (Ipv4VrfState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipv4VrfState)(nil)).Elem()
}

type ipv4VrfArgs struct {
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// VRF name.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a Ipv4Vrf resource.
type Ipv4VrfArgs struct {
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
	// VRF name.
	Name pulumi.StringPtrInput
}

func (Ipv4VrfArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipv4VrfArgs)(nil)).Elem()
}

type Ipv4VrfInput interface {
	pulumi.Input

	ToIpv4VrfOutput() Ipv4VrfOutput
	ToIpv4VrfOutputWithContext(ctx context.Context) Ipv4VrfOutput
}

func (*Ipv4Vrf) ElementType() reflect.Type {
	return reflect.TypeOf((**Ipv4Vrf)(nil)).Elem()
}

func (i *Ipv4Vrf) ToIpv4VrfOutput() Ipv4VrfOutput {
	return i.ToIpv4VrfOutputWithContext(context.Background())
}

func (i *Ipv4Vrf) ToIpv4VrfOutputWithContext(ctx context.Context) Ipv4VrfOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ipv4VrfOutput)
}

func (i *Ipv4Vrf) ToOutput(ctx context.Context) pulumix.Output[*Ipv4Vrf] {
	return pulumix.Output[*Ipv4Vrf]{
		OutputState: i.ToIpv4VrfOutputWithContext(ctx).OutputState,
	}
}

// Ipv4VrfArrayInput is an input type that accepts Ipv4VrfArray and Ipv4VrfArrayOutput values.
// You can construct a concrete instance of `Ipv4VrfArrayInput` via:
//
//	Ipv4VrfArray{ Ipv4VrfArgs{...} }
type Ipv4VrfArrayInput interface {
	pulumi.Input

	ToIpv4VrfArrayOutput() Ipv4VrfArrayOutput
	ToIpv4VrfArrayOutputWithContext(context.Context) Ipv4VrfArrayOutput
}

type Ipv4VrfArray []Ipv4VrfInput

func (Ipv4VrfArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ipv4Vrf)(nil)).Elem()
}

func (i Ipv4VrfArray) ToIpv4VrfArrayOutput() Ipv4VrfArrayOutput {
	return i.ToIpv4VrfArrayOutputWithContext(context.Background())
}

func (i Ipv4VrfArray) ToIpv4VrfArrayOutputWithContext(ctx context.Context) Ipv4VrfArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ipv4VrfArrayOutput)
}

func (i Ipv4VrfArray) ToOutput(ctx context.Context) pulumix.Output[[]*Ipv4Vrf] {
	return pulumix.Output[[]*Ipv4Vrf]{
		OutputState: i.ToIpv4VrfArrayOutputWithContext(ctx).OutputState,
	}
}

// Ipv4VrfMapInput is an input type that accepts Ipv4VrfMap and Ipv4VrfMapOutput values.
// You can construct a concrete instance of `Ipv4VrfMapInput` via:
//
//	Ipv4VrfMap{ "key": Ipv4VrfArgs{...} }
type Ipv4VrfMapInput interface {
	pulumi.Input

	ToIpv4VrfMapOutput() Ipv4VrfMapOutput
	ToIpv4VrfMapOutputWithContext(context.Context) Ipv4VrfMapOutput
}

type Ipv4VrfMap map[string]Ipv4VrfInput

func (Ipv4VrfMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ipv4Vrf)(nil)).Elem()
}

func (i Ipv4VrfMap) ToIpv4VrfMapOutput() Ipv4VrfMapOutput {
	return i.ToIpv4VrfMapOutputWithContext(context.Background())
}

func (i Ipv4VrfMap) ToIpv4VrfMapOutputWithContext(ctx context.Context) Ipv4VrfMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ipv4VrfMapOutput)
}

func (i Ipv4VrfMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Ipv4Vrf] {
	return pulumix.Output[map[string]*Ipv4Vrf]{
		OutputState: i.ToIpv4VrfMapOutputWithContext(ctx).OutputState,
	}
}

type Ipv4VrfOutput struct{ *pulumi.OutputState }

func (Ipv4VrfOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ipv4Vrf)(nil)).Elem()
}

func (o Ipv4VrfOutput) ToIpv4VrfOutput() Ipv4VrfOutput {
	return o
}

func (o Ipv4VrfOutput) ToIpv4VrfOutputWithContext(ctx context.Context) Ipv4VrfOutput {
	return o
}

func (o Ipv4VrfOutput) ToOutput(ctx context.Context) pulumix.Output[*Ipv4Vrf] {
	return pulumix.Output[*Ipv4Vrf]{
		OutputState: o.OutputState,
	}
}

// A device name from the provider configuration.
func (o Ipv4VrfOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ipv4Vrf) pulumi.StringPtrOutput { return v.Device }).(pulumi.StringPtrOutput)
}

// VRF name.
func (o Ipv4VrfOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Ipv4Vrf) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type Ipv4VrfArrayOutput struct{ *pulumi.OutputState }

func (Ipv4VrfArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ipv4Vrf)(nil)).Elem()
}

func (o Ipv4VrfArrayOutput) ToIpv4VrfArrayOutput() Ipv4VrfArrayOutput {
	return o
}

func (o Ipv4VrfArrayOutput) ToIpv4VrfArrayOutputWithContext(ctx context.Context) Ipv4VrfArrayOutput {
	return o
}

func (o Ipv4VrfArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Ipv4Vrf] {
	return pulumix.Output[[]*Ipv4Vrf]{
		OutputState: o.OutputState,
	}
}

func (o Ipv4VrfArrayOutput) Index(i pulumi.IntInput) Ipv4VrfOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Ipv4Vrf {
		return vs[0].([]*Ipv4Vrf)[vs[1].(int)]
	}).(Ipv4VrfOutput)
}

type Ipv4VrfMapOutput struct{ *pulumi.OutputState }

func (Ipv4VrfMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ipv4Vrf)(nil)).Elem()
}

func (o Ipv4VrfMapOutput) ToIpv4VrfMapOutput() Ipv4VrfMapOutput {
	return o
}

func (o Ipv4VrfMapOutput) ToIpv4VrfMapOutputWithContext(ctx context.Context) Ipv4VrfMapOutput {
	return o
}

func (o Ipv4VrfMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Ipv4Vrf] {
	return pulumix.Output[map[string]*Ipv4Vrf]{
		OutputState: o.OutputState,
	}
}

func (o Ipv4VrfMapOutput) MapIndex(k pulumi.StringInput) Ipv4VrfOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Ipv4Vrf {
		return vs[0].(map[string]*Ipv4Vrf)[vs[1].(string)]
	}).(Ipv4VrfOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*Ipv4VrfInput)(nil)).Elem(), &Ipv4Vrf{})
	pulumi.RegisterInputType(reflect.TypeOf((*Ipv4VrfArrayInput)(nil)).Elem(), Ipv4VrfArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*Ipv4VrfMapInput)(nil)).Elem(), Ipv4VrfMap{})
	pulumi.RegisterOutputType(Ipv4VrfOutput{})
	pulumi.RegisterOutputType(Ipv4VrfArrayOutput{})
	pulumi.RegisterOutputType(Ipv4VrfMapOutput{})
}
