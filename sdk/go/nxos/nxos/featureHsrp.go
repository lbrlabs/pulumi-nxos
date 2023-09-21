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

type FeatureHsrp struct {
	pulumi.CustomResourceState

	// Administrative state. - Choices: `enabled`, `disabled`
	AdminState pulumi.StringOutput `pulumi:"adminState"`
	// A device name from the provider configuration.
	Device pulumi.StringPtrOutput `pulumi:"device"`
}

// NewFeatureHsrp registers a new resource with the given unique name, arguments, and options.
func NewFeatureHsrp(ctx *pulumi.Context,
	name string, args *FeatureHsrpArgs, opts ...pulumi.ResourceOption) (*FeatureHsrp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AdminState == nil {
		return nil, errors.New("invalid value for required argument 'AdminState'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FeatureHsrp
	err := ctx.RegisterResource("nxos:nxos/featureHsrp:FeatureHsrp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFeatureHsrp gets an existing FeatureHsrp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFeatureHsrp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FeatureHsrpState, opts ...pulumi.ResourceOption) (*FeatureHsrp, error) {
	var resource FeatureHsrp
	err := ctx.ReadResource("nxos:nxos/featureHsrp:FeatureHsrp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FeatureHsrp resources.
type featureHsrpState struct {
	// Administrative state. - Choices: `enabled`, `disabled`
	AdminState *string `pulumi:"adminState"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
}

type FeatureHsrpState struct {
	// Administrative state. - Choices: `enabled`, `disabled`
	AdminState pulumi.StringPtrInput
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
}

func (FeatureHsrpState) ElementType() reflect.Type {
	return reflect.TypeOf((*featureHsrpState)(nil)).Elem()
}

type featureHsrpArgs struct {
	// Administrative state. - Choices: `enabled`, `disabled`
	AdminState string `pulumi:"adminState"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
}

// The set of arguments for constructing a FeatureHsrp resource.
type FeatureHsrpArgs struct {
	// Administrative state. - Choices: `enabled`, `disabled`
	AdminState pulumi.StringInput
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
}

func (FeatureHsrpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*featureHsrpArgs)(nil)).Elem()
}

type FeatureHsrpInput interface {
	pulumi.Input

	ToFeatureHsrpOutput() FeatureHsrpOutput
	ToFeatureHsrpOutputWithContext(ctx context.Context) FeatureHsrpOutput
}

func (*FeatureHsrp) ElementType() reflect.Type {
	return reflect.TypeOf((**FeatureHsrp)(nil)).Elem()
}

func (i *FeatureHsrp) ToFeatureHsrpOutput() FeatureHsrpOutput {
	return i.ToFeatureHsrpOutputWithContext(context.Background())
}

func (i *FeatureHsrp) ToFeatureHsrpOutputWithContext(ctx context.Context) FeatureHsrpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FeatureHsrpOutput)
}

func (i *FeatureHsrp) ToOutput(ctx context.Context) pulumix.Output[*FeatureHsrp] {
	return pulumix.Output[*FeatureHsrp]{
		OutputState: i.ToFeatureHsrpOutputWithContext(ctx).OutputState,
	}
}

// FeatureHsrpArrayInput is an input type that accepts FeatureHsrpArray and FeatureHsrpArrayOutput values.
// You can construct a concrete instance of `FeatureHsrpArrayInput` via:
//
//	FeatureHsrpArray{ FeatureHsrpArgs{...} }
type FeatureHsrpArrayInput interface {
	pulumi.Input

	ToFeatureHsrpArrayOutput() FeatureHsrpArrayOutput
	ToFeatureHsrpArrayOutputWithContext(context.Context) FeatureHsrpArrayOutput
}

type FeatureHsrpArray []FeatureHsrpInput

func (FeatureHsrpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FeatureHsrp)(nil)).Elem()
}

func (i FeatureHsrpArray) ToFeatureHsrpArrayOutput() FeatureHsrpArrayOutput {
	return i.ToFeatureHsrpArrayOutputWithContext(context.Background())
}

func (i FeatureHsrpArray) ToFeatureHsrpArrayOutputWithContext(ctx context.Context) FeatureHsrpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FeatureHsrpArrayOutput)
}

func (i FeatureHsrpArray) ToOutput(ctx context.Context) pulumix.Output[[]*FeatureHsrp] {
	return pulumix.Output[[]*FeatureHsrp]{
		OutputState: i.ToFeatureHsrpArrayOutputWithContext(ctx).OutputState,
	}
}

// FeatureHsrpMapInput is an input type that accepts FeatureHsrpMap and FeatureHsrpMapOutput values.
// You can construct a concrete instance of `FeatureHsrpMapInput` via:
//
//	FeatureHsrpMap{ "key": FeatureHsrpArgs{...} }
type FeatureHsrpMapInput interface {
	pulumi.Input

	ToFeatureHsrpMapOutput() FeatureHsrpMapOutput
	ToFeatureHsrpMapOutputWithContext(context.Context) FeatureHsrpMapOutput
}

type FeatureHsrpMap map[string]FeatureHsrpInput

func (FeatureHsrpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FeatureHsrp)(nil)).Elem()
}

func (i FeatureHsrpMap) ToFeatureHsrpMapOutput() FeatureHsrpMapOutput {
	return i.ToFeatureHsrpMapOutputWithContext(context.Background())
}

func (i FeatureHsrpMap) ToFeatureHsrpMapOutputWithContext(ctx context.Context) FeatureHsrpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FeatureHsrpMapOutput)
}

func (i FeatureHsrpMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*FeatureHsrp] {
	return pulumix.Output[map[string]*FeatureHsrp]{
		OutputState: i.ToFeatureHsrpMapOutputWithContext(ctx).OutputState,
	}
}

type FeatureHsrpOutput struct{ *pulumi.OutputState }

func (FeatureHsrpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FeatureHsrp)(nil)).Elem()
}

func (o FeatureHsrpOutput) ToFeatureHsrpOutput() FeatureHsrpOutput {
	return o
}

func (o FeatureHsrpOutput) ToFeatureHsrpOutputWithContext(ctx context.Context) FeatureHsrpOutput {
	return o
}

func (o FeatureHsrpOutput) ToOutput(ctx context.Context) pulumix.Output[*FeatureHsrp] {
	return pulumix.Output[*FeatureHsrp]{
		OutputState: o.OutputState,
	}
}

// Administrative state. - Choices: `enabled`, `disabled`
func (o FeatureHsrpOutput) AdminState() pulumi.StringOutput {
	return o.ApplyT(func(v *FeatureHsrp) pulumi.StringOutput { return v.AdminState }).(pulumi.StringOutput)
}

// A device name from the provider configuration.
func (o FeatureHsrpOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FeatureHsrp) pulumi.StringPtrOutput { return v.Device }).(pulumi.StringPtrOutput)
}

type FeatureHsrpArrayOutput struct{ *pulumi.OutputState }

func (FeatureHsrpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FeatureHsrp)(nil)).Elem()
}

func (o FeatureHsrpArrayOutput) ToFeatureHsrpArrayOutput() FeatureHsrpArrayOutput {
	return o
}

func (o FeatureHsrpArrayOutput) ToFeatureHsrpArrayOutputWithContext(ctx context.Context) FeatureHsrpArrayOutput {
	return o
}

func (o FeatureHsrpArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*FeatureHsrp] {
	return pulumix.Output[[]*FeatureHsrp]{
		OutputState: o.OutputState,
	}
}

func (o FeatureHsrpArrayOutput) Index(i pulumi.IntInput) FeatureHsrpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FeatureHsrp {
		return vs[0].([]*FeatureHsrp)[vs[1].(int)]
	}).(FeatureHsrpOutput)
}

type FeatureHsrpMapOutput struct{ *pulumi.OutputState }

func (FeatureHsrpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FeatureHsrp)(nil)).Elem()
}

func (o FeatureHsrpMapOutput) ToFeatureHsrpMapOutput() FeatureHsrpMapOutput {
	return o
}

func (o FeatureHsrpMapOutput) ToFeatureHsrpMapOutputWithContext(ctx context.Context) FeatureHsrpMapOutput {
	return o
}

func (o FeatureHsrpMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*FeatureHsrp] {
	return pulumix.Output[map[string]*FeatureHsrp]{
		OutputState: o.OutputState,
	}
}

func (o FeatureHsrpMapOutput) MapIndex(k pulumi.StringInput) FeatureHsrpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FeatureHsrp {
		return vs[0].(map[string]*FeatureHsrp)[vs[1].(string)]
	}).(FeatureHsrpOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FeatureHsrpInput)(nil)).Elem(), &FeatureHsrp{})
	pulumi.RegisterInputType(reflect.TypeOf((*FeatureHsrpArrayInput)(nil)).Elem(), FeatureHsrpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FeatureHsrpMapInput)(nil)).Elem(), FeatureHsrpMap{})
	pulumi.RegisterOutputType(FeatureHsrpOutput{})
	pulumi.RegisterOutputType(FeatureHsrpArrayOutput{})
	pulumi.RegisterOutputType(FeatureHsrpMapOutput{})
}
