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

type Evpn struct {
	pulumi.CustomResourceState

	// Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
	AdminState pulumi.StringOutput `pulumi:"adminState"`
	// A device name from the provider configuration.
	Device pulumi.StringPtrOutput `pulumi:"device"`
}

// NewEvpn registers a new resource with the given unique name, arguments, and options.
func NewEvpn(ctx *pulumi.Context,
	name string, args *EvpnArgs, opts ...pulumi.ResourceOption) (*Evpn, error) {
	if args == nil {
		args = &EvpnArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Evpn
	err := ctx.RegisterResource("nxos:nxos/evpn:Evpn", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEvpn gets an existing Evpn resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEvpn(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EvpnState, opts ...pulumi.ResourceOption) (*Evpn, error) {
	var resource Evpn
	err := ctx.ReadResource("nxos:nxos/evpn:Evpn", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Evpn resources.
type evpnState struct {
	// Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
	AdminState *string `pulumi:"adminState"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
}

type EvpnState struct {
	// Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
	AdminState pulumi.StringPtrInput
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
}

func (EvpnState) ElementType() reflect.Type {
	return reflect.TypeOf((*evpnState)(nil)).Elem()
}

type evpnArgs struct {
	// Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
	AdminState *string `pulumi:"adminState"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
}

// The set of arguments for constructing a Evpn resource.
type EvpnArgs struct {
	// Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
	AdminState pulumi.StringPtrInput
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
}

func (EvpnArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*evpnArgs)(nil)).Elem()
}

type EvpnInput interface {
	pulumi.Input

	ToEvpnOutput() EvpnOutput
	ToEvpnOutputWithContext(ctx context.Context) EvpnOutput
}

func (*Evpn) ElementType() reflect.Type {
	return reflect.TypeOf((**Evpn)(nil)).Elem()
}

func (i *Evpn) ToEvpnOutput() EvpnOutput {
	return i.ToEvpnOutputWithContext(context.Background())
}

func (i *Evpn) ToEvpnOutputWithContext(ctx context.Context) EvpnOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EvpnOutput)
}

func (i *Evpn) ToOutput(ctx context.Context) pulumix.Output[*Evpn] {
	return pulumix.Output[*Evpn]{
		OutputState: i.ToEvpnOutputWithContext(ctx).OutputState,
	}
}

// EvpnArrayInput is an input type that accepts EvpnArray and EvpnArrayOutput values.
// You can construct a concrete instance of `EvpnArrayInput` via:
//
//	EvpnArray{ EvpnArgs{...} }
type EvpnArrayInput interface {
	pulumi.Input

	ToEvpnArrayOutput() EvpnArrayOutput
	ToEvpnArrayOutputWithContext(context.Context) EvpnArrayOutput
}

type EvpnArray []EvpnInput

func (EvpnArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Evpn)(nil)).Elem()
}

func (i EvpnArray) ToEvpnArrayOutput() EvpnArrayOutput {
	return i.ToEvpnArrayOutputWithContext(context.Background())
}

func (i EvpnArray) ToEvpnArrayOutputWithContext(ctx context.Context) EvpnArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EvpnArrayOutput)
}

func (i EvpnArray) ToOutput(ctx context.Context) pulumix.Output[[]*Evpn] {
	return pulumix.Output[[]*Evpn]{
		OutputState: i.ToEvpnArrayOutputWithContext(ctx).OutputState,
	}
}

// EvpnMapInput is an input type that accepts EvpnMap and EvpnMapOutput values.
// You can construct a concrete instance of `EvpnMapInput` via:
//
//	EvpnMap{ "key": EvpnArgs{...} }
type EvpnMapInput interface {
	pulumi.Input

	ToEvpnMapOutput() EvpnMapOutput
	ToEvpnMapOutputWithContext(context.Context) EvpnMapOutput
}

type EvpnMap map[string]EvpnInput

func (EvpnMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Evpn)(nil)).Elem()
}

func (i EvpnMap) ToEvpnMapOutput() EvpnMapOutput {
	return i.ToEvpnMapOutputWithContext(context.Background())
}

func (i EvpnMap) ToEvpnMapOutputWithContext(ctx context.Context) EvpnMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EvpnMapOutput)
}

func (i EvpnMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Evpn] {
	return pulumix.Output[map[string]*Evpn]{
		OutputState: i.ToEvpnMapOutputWithContext(ctx).OutputState,
	}
}

type EvpnOutput struct{ *pulumi.OutputState }

func (EvpnOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Evpn)(nil)).Elem()
}

func (o EvpnOutput) ToEvpnOutput() EvpnOutput {
	return o
}

func (o EvpnOutput) ToEvpnOutputWithContext(ctx context.Context) EvpnOutput {
	return o
}

func (o EvpnOutput) ToOutput(ctx context.Context) pulumix.Output[*Evpn] {
	return pulumix.Output[*Evpn]{
		OutputState: o.OutputState,
	}
}

// Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
func (o EvpnOutput) AdminState() pulumi.StringOutput {
	return o.ApplyT(func(v *Evpn) pulumi.StringOutput { return v.AdminState }).(pulumi.StringOutput)
}

// A device name from the provider configuration.
func (o EvpnOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Evpn) pulumi.StringPtrOutput { return v.Device }).(pulumi.StringPtrOutput)
}

type EvpnArrayOutput struct{ *pulumi.OutputState }

func (EvpnArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Evpn)(nil)).Elem()
}

func (o EvpnArrayOutput) ToEvpnArrayOutput() EvpnArrayOutput {
	return o
}

func (o EvpnArrayOutput) ToEvpnArrayOutputWithContext(ctx context.Context) EvpnArrayOutput {
	return o
}

func (o EvpnArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Evpn] {
	return pulumix.Output[[]*Evpn]{
		OutputState: o.OutputState,
	}
}

func (o EvpnArrayOutput) Index(i pulumi.IntInput) EvpnOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Evpn {
		return vs[0].([]*Evpn)[vs[1].(int)]
	}).(EvpnOutput)
}

type EvpnMapOutput struct{ *pulumi.OutputState }

func (EvpnMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Evpn)(nil)).Elem()
}

func (o EvpnMapOutput) ToEvpnMapOutput() EvpnMapOutput {
	return o
}

func (o EvpnMapOutput) ToEvpnMapOutputWithContext(ctx context.Context) EvpnMapOutput {
	return o
}

func (o EvpnMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Evpn] {
	return pulumix.Output[map[string]*Evpn]{
		OutputState: o.OutputState,
	}
}

func (o EvpnMapOutput) MapIndex(k pulumi.StringInput) EvpnOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Evpn {
		return vs[0].(map[string]*Evpn)[vs[1].(string)]
	}).(EvpnOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EvpnInput)(nil)).Elem(), &Evpn{})
	pulumi.RegisterInputType(reflect.TypeOf((*EvpnArrayInput)(nil)).Elem(), EvpnArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EvpnMapInput)(nil)).Elem(), EvpnMap{})
	pulumi.RegisterOutputType(EvpnOutput{})
	pulumi.RegisterOutputType(EvpnArrayOutput{})
	pulumi.RegisterOutputType(EvpnMapOutput{})
}
