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

type FeatureVpc struct {
	pulumi.CustomResourceState

	// Administrative state. - Choices: `enabled`, `disabled`
	AdminState pulumi.StringOutput `pulumi:"adminState"`
	// A device name from the provider configuration.
	Device pulumi.StringPtrOutput `pulumi:"device"`
}

// NewFeatureVpc registers a new resource with the given unique name, arguments, and options.
func NewFeatureVpc(ctx *pulumi.Context,
	name string, args *FeatureVpcArgs, opts ...pulumi.ResourceOption) (*FeatureVpc, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AdminState == nil {
		return nil, errors.New("invalid value for required argument 'AdminState'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FeatureVpc
	err := ctx.RegisterResource("nxos:nxos/featureVpc:FeatureVpc", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFeatureVpc gets an existing FeatureVpc resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFeatureVpc(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FeatureVpcState, opts ...pulumi.ResourceOption) (*FeatureVpc, error) {
	var resource FeatureVpc
	err := ctx.ReadResource("nxos:nxos/featureVpc:FeatureVpc", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FeatureVpc resources.
type featureVpcState struct {
	// Administrative state. - Choices: `enabled`, `disabled`
	AdminState *string `pulumi:"adminState"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
}

type FeatureVpcState struct {
	// Administrative state. - Choices: `enabled`, `disabled`
	AdminState pulumi.StringPtrInput
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
}

func (FeatureVpcState) ElementType() reflect.Type {
	return reflect.TypeOf((*featureVpcState)(nil)).Elem()
}

type featureVpcArgs struct {
	// Administrative state. - Choices: `enabled`, `disabled`
	AdminState string `pulumi:"adminState"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
}

// The set of arguments for constructing a FeatureVpc resource.
type FeatureVpcArgs struct {
	// Administrative state. - Choices: `enabled`, `disabled`
	AdminState pulumi.StringInput
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
}

func (FeatureVpcArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*featureVpcArgs)(nil)).Elem()
}

type FeatureVpcInput interface {
	pulumi.Input

	ToFeatureVpcOutput() FeatureVpcOutput
	ToFeatureVpcOutputWithContext(ctx context.Context) FeatureVpcOutput
}

func (*FeatureVpc) ElementType() reflect.Type {
	return reflect.TypeOf((**FeatureVpc)(nil)).Elem()
}

func (i *FeatureVpc) ToFeatureVpcOutput() FeatureVpcOutput {
	return i.ToFeatureVpcOutputWithContext(context.Background())
}

func (i *FeatureVpc) ToFeatureVpcOutputWithContext(ctx context.Context) FeatureVpcOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FeatureVpcOutput)
}

func (i *FeatureVpc) ToOutput(ctx context.Context) pulumix.Output[*FeatureVpc] {
	return pulumix.Output[*FeatureVpc]{
		OutputState: i.ToFeatureVpcOutputWithContext(ctx).OutputState,
	}
}

// FeatureVpcArrayInput is an input type that accepts FeatureVpcArray and FeatureVpcArrayOutput values.
// You can construct a concrete instance of `FeatureVpcArrayInput` via:
//
//	FeatureVpcArray{ FeatureVpcArgs{...} }
type FeatureVpcArrayInput interface {
	pulumi.Input

	ToFeatureVpcArrayOutput() FeatureVpcArrayOutput
	ToFeatureVpcArrayOutputWithContext(context.Context) FeatureVpcArrayOutput
}

type FeatureVpcArray []FeatureVpcInput

func (FeatureVpcArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FeatureVpc)(nil)).Elem()
}

func (i FeatureVpcArray) ToFeatureVpcArrayOutput() FeatureVpcArrayOutput {
	return i.ToFeatureVpcArrayOutputWithContext(context.Background())
}

func (i FeatureVpcArray) ToFeatureVpcArrayOutputWithContext(ctx context.Context) FeatureVpcArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FeatureVpcArrayOutput)
}

func (i FeatureVpcArray) ToOutput(ctx context.Context) pulumix.Output[[]*FeatureVpc] {
	return pulumix.Output[[]*FeatureVpc]{
		OutputState: i.ToFeatureVpcArrayOutputWithContext(ctx).OutputState,
	}
}

// FeatureVpcMapInput is an input type that accepts FeatureVpcMap and FeatureVpcMapOutput values.
// You can construct a concrete instance of `FeatureVpcMapInput` via:
//
//	FeatureVpcMap{ "key": FeatureVpcArgs{...} }
type FeatureVpcMapInput interface {
	pulumi.Input

	ToFeatureVpcMapOutput() FeatureVpcMapOutput
	ToFeatureVpcMapOutputWithContext(context.Context) FeatureVpcMapOutput
}

type FeatureVpcMap map[string]FeatureVpcInput

func (FeatureVpcMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FeatureVpc)(nil)).Elem()
}

func (i FeatureVpcMap) ToFeatureVpcMapOutput() FeatureVpcMapOutput {
	return i.ToFeatureVpcMapOutputWithContext(context.Background())
}

func (i FeatureVpcMap) ToFeatureVpcMapOutputWithContext(ctx context.Context) FeatureVpcMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FeatureVpcMapOutput)
}

func (i FeatureVpcMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*FeatureVpc] {
	return pulumix.Output[map[string]*FeatureVpc]{
		OutputState: i.ToFeatureVpcMapOutputWithContext(ctx).OutputState,
	}
}

type FeatureVpcOutput struct{ *pulumi.OutputState }

func (FeatureVpcOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FeatureVpc)(nil)).Elem()
}

func (o FeatureVpcOutput) ToFeatureVpcOutput() FeatureVpcOutput {
	return o
}

func (o FeatureVpcOutput) ToFeatureVpcOutputWithContext(ctx context.Context) FeatureVpcOutput {
	return o
}

func (o FeatureVpcOutput) ToOutput(ctx context.Context) pulumix.Output[*FeatureVpc] {
	return pulumix.Output[*FeatureVpc]{
		OutputState: o.OutputState,
	}
}

// Administrative state. - Choices: `enabled`, `disabled`
func (o FeatureVpcOutput) AdminState() pulumi.StringOutput {
	return o.ApplyT(func(v *FeatureVpc) pulumi.StringOutput { return v.AdminState }).(pulumi.StringOutput)
}

// A device name from the provider configuration.
func (o FeatureVpcOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FeatureVpc) pulumi.StringPtrOutput { return v.Device }).(pulumi.StringPtrOutput)
}

type FeatureVpcArrayOutput struct{ *pulumi.OutputState }

func (FeatureVpcArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FeatureVpc)(nil)).Elem()
}

func (o FeatureVpcArrayOutput) ToFeatureVpcArrayOutput() FeatureVpcArrayOutput {
	return o
}

func (o FeatureVpcArrayOutput) ToFeatureVpcArrayOutputWithContext(ctx context.Context) FeatureVpcArrayOutput {
	return o
}

func (o FeatureVpcArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*FeatureVpc] {
	return pulumix.Output[[]*FeatureVpc]{
		OutputState: o.OutputState,
	}
}

func (o FeatureVpcArrayOutput) Index(i pulumi.IntInput) FeatureVpcOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FeatureVpc {
		return vs[0].([]*FeatureVpc)[vs[1].(int)]
	}).(FeatureVpcOutput)
}

type FeatureVpcMapOutput struct{ *pulumi.OutputState }

func (FeatureVpcMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FeatureVpc)(nil)).Elem()
}

func (o FeatureVpcMapOutput) ToFeatureVpcMapOutput() FeatureVpcMapOutput {
	return o
}

func (o FeatureVpcMapOutput) ToFeatureVpcMapOutputWithContext(ctx context.Context) FeatureVpcMapOutput {
	return o
}

func (o FeatureVpcMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*FeatureVpc] {
	return pulumix.Output[map[string]*FeatureVpc]{
		OutputState: o.OutputState,
	}
}

func (o FeatureVpcMapOutput) MapIndex(k pulumi.StringInput) FeatureVpcOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FeatureVpc {
		return vs[0].(map[string]*FeatureVpc)[vs[1].(string)]
	}).(FeatureVpcOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FeatureVpcInput)(nil)).Elem(), &FeatureVpc{})
	pulumi.RegisterInputType(reflect.TypeOf((*FeatureVpcArrayInput)(nil)).Elem(), FeatureVpcArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FeatureVpcMapInput)(nil)).Elem(), FeatureVpcMap{})
	pulumi.RegisterOutputType(FeatureVpcOutput{})
	pulumi.RegisterOutputType(FeatureVpcArrayOutput{})
	pulumi.RegisterOutputType(FeatureVpcMapOutput{})
}