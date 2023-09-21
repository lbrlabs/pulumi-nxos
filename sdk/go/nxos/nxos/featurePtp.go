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

type FeaturePtp struct {
	pulumi.CustomResourceState

	// Administrative state. - Choices: `enabled`, `disabled`
	AdminState pulumi.StringOutput `pulumi:"adminState"`
	// A device name from the provider configuration.
	Device pulumi.StringPtrOutput `pulumi:"device"`
}

// NewFeaturePtp registers a new resource with the given unique name, arguments, and options.
func NewFeaturePtp(ctx *pulumi.Context,
	name string, args *FeaturePtpArgs, opts ...pulumi.ResourceOption) (*FeaturePtp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AdminState == nil {
		return nil, errors.New("invalid value for required argument 'AdminState'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FeaturePtp
	err := ctx.RegisterResource("nxos:nxos/featurePtp:FeaturePtp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFeaturePtp gets an existing FeaturePtp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFeaturePtp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FeaturePtpState, opts ...pulumi.ResourceOption) (*FeaturePtp, error) {
	var resource FeaturePtp
	err := ctx.ReadResource("nxos:nxos/featurePtp:FeaturePtp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FeaturePtp resources.
type featurePtpState struct {
	// Administrative state. - Choices: `enabled`, `disabled`
	AdminState *string `pulumi:"adminState"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
}

type FeaturePtpState struct {
	// Administrative state. - Choices: `enabled`, `disabled`
	AdminState pulumi.StringPtrInput
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
}

func (FeaturePtpState) ElementType() reflect.Type {
	return reflect.TypeOf((*featurePtpState)(nil)).Elem()
}

type featurePtpArgs struct {
	// Administrative state. - Choices: `enabled`, `disabled`
	AdminState string `pulumi:"adminState"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
}

// The set of arguments for constructing a FeaturePtp resource.
type FeaturePtpArgs struct {
	// Administrative state. - Choices: `enabled`, `disabled`
	AdminState pulumi.StringInput
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
}

func (FeaturePtpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*featurePtpArgs)(nil)).Elem()
}

type FeaturePtpInput interface {
	pulumi.Input

	ToFeaturePtpOutput() FeaturePtpOutput
	ToFeaturePtpOutputWithContext(ctx context.Context) FeaturePtpOutput
}

func (*FeaturePtp) ElementType() reflect.Type {
	return reflect.TypeOf((**FeaturePtp)(nil)).Elem()
}

func (i *FeaturePtp) ToFeaturePtpOutput() FeaturePtpOutput {
	return i.ToFeaturePtpOutputWithContext(context.Background())
}

func (i *FeaturePtp) ToFeaturePtpOutputWithContext(ctx context.Context) FeaturePtpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FeaturePtpOutput)
}

func (i *FeaturePtp) ToOutput(ctx context.Context) pulumix.Output[*FeaturePtp] {
	return pulumix.Output[*FeaturePtp]{
		OutputState: i.ToFeaturePtpOutputWithContext(ctx).OutputState,
	}
}

// FeaturePtpArrayInput is an input type that accepts FeaturePtpArray and FeaturePtpArrayOutput values.
// You can construct a concrete instance of `FeaturePtpArrayInput` via:
//
//	FeaturePtpArray{ FeaturePtpArgs{...} }
type FeaturePtpArrayInput interface {
	pulumi.Input

	ToFeaturePtpArrayOutput() FeaturePtpArrayOutput
	ToFeaturePtpArrayOutputWithContext(context.Context) FeaturePtpArrayOutput
}

type FeaturePtpArray []FeaturePtpInput

func (FeaturePtpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FeaturePtp)(nil)).Elem()
}

func (i FeaturePtpArray) ToFeaturePtpArrayOutput() FeaturePtpArrayOutput {
	return i.ToFeaturePtpArrayOutputWithContext(context.Background())
}

func (i FeaturePtpArray) ToFeaturePtpArrayOutputWithContext(ctx context.Context) FeaturePtpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FeaturePtpArrayOutput)
}

func (i FeaturePtpArray) ToOutput(ctx context.Context) pulumix.Output[[]*FeaturePtp] {
	return pulumix.Output[[]*FeaturePtp]{
		OutputState: i.ToFeaturePtpArrayOutputWithContext(ctx).OutputState,
	}
}

// FeaturePtpMapInput is an input type that accepts FeaturePtpMap and FeaturePtpMapOutput values.
// You can construct a concrete instance of `FeaturePtpMapInput` via:
//
//	FeaturePtpMap{ "key": FeaturePtpArgs{...} }
type FeaturePtpMapInput interface {
	pulumi.Input

	ToFeaturePtpMapOutput() FeaturePtpMapOutput
	ToFeaturePtpMapOutputWithContext(context.Context) FeaturePtpMapOutput
}

type FeaturePtpMap map[string]FeaturePtpInput

func (FeaturePtpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FeaturePtp)(nil)).Elem()
}

func (i FeaturePtpMap) ToFeaturePtpMapOutput() FeaturePtpMapOutput {
	return i.ToFeaturePtpMapOutputWithContext(context.Background())
}

func (i FeaturePtpMap) ToFeaturePtpMapOutputWithContext(ctx context.Context) FeaturePtpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FeaturePtpMapOutput)
}

func (i FeaturePtpMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*FeaturePtp] {
	return pulumix.Output[map[string]*FeaturePtp]{
		OutputState: i.ToFeaturePtpMapOutputWithContext(ctx).OutputState,
	}
}

type FeaturePtpOutput struct{ *pulumi.OutputState }

func (FeaturePtpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FeaturePtp)(nil)).Elem()
}

func (o FeaturePtpOutput) ToFeaturePtpOutput() FeaturePtpOutput {
	return o
}

func (o FeaturePtpOutput) ToFeaturePtpOutputWithContext(ctx context.Context) FeaturePtpOutput {
	return o
}

func (o FeaturePtpOutput) ToOutput(ctx context.Context) pulumix.Output[*FeaturePtp] {
	return pulumix.Output[*FeaturePtp]{
		OutputState: o.OutputState,
	}
}

// Administrative state. - Choices: `enabled`, `disabled`
func (o FeaturePtpOutput) AdminState() pulumi.StringOutput {
	return o.ApplyT(func(v *FeaturePtp) pulumi.StringOutput { return v.AdminState }).(pulumi.StringOutput)
}

// A device name from the provider configuration.
func (o FeaturePtpOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FeaturePtp) pulumi.StringPtrOutput { return v.Device }).(pulumi.StringPtrOutput)
}

type FeaturePtpArrayOutput struct{ *pulumi.OutputState }

func (FeaturePtpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FeaturePtp)(nil)).Elem()
}

func (o FeaturePtpArrayOutput) ToFeaturePtpArrayOutput() FeaturePtpArrayOutput {
	return o
}

func (o FeaturePtpArrayOutput) ToFeaturePtpArrayOutputWithContext(ctx context.Context) FeaturePtpArrayOutput {
	return o
}

func (o FeaturePtpArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*FeaturePtp] {
	return pulumix.Output[[]*FeaturePtp]{
		OutputState: o.OutputState,
	}
}

func (o FeaturePtpArrayOutput) Index(i pulumi.IntInput) FeaturePtpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FeaturePtp {
		return vs[0].([]*FeaturePtp)[vs[1].(int)]
	}).(FeaturePtpOutput)
}

type FeaturePtpMapOutput struct{ *pulumi.OutputState }

func (FeaturePtpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FeaturePtp)(nil)).Elem()
}

func (o FeaturePtpMapOutput) ToFeaturePtpMapOutput() FeaturePtpMapOutput {
	return o
}

func (o FeaturePtpMapOutput) ToFeaturePtpMapOutputWithContext(ctx context.Context) FeaturePtpMapOutput {
	return o
}

func (o FeaturePtpMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*FeaturePtp] {
	return pulumix.Output[map[string]*FeaturePtp]{
		OutputState: o.OutputState,
	}
}

func (o FeaturePtpMapOutput) MapIndex(k pulumi.StringInput) FeaturePtpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FeaturePtp {
		return vs[0].(map[string]*FeaturePtp)[vs[1].(string)]
	}).(FeaturePtpOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FeaturePtpInput)(nil)).Elem(), &FeaturePtp{})
	pulumi.RegisterInputType(reflect.TypeOf((*FeaturePtpArrayInput)(nil)).Elem(), FeaturePtpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FeaturePtpMapInput)(nil)).Elem(), FeaturePtpMap{})
	pulumi.RegisterOutputType(FeaturePtpOutput{})
	pulumi.RegisterOutputType(FeaturePtpArrayOutput{})
	pulumi.RegisterOutputType(FeaturePtpMapOutput{})
}