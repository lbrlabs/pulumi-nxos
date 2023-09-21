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

type FeatureVnSegment struct {
	pulumi.CustomResourceState

	// Administrative state. - Choices: `enabled`, `disabled`
	AdminState pulumi.StringOutput `pulumi:"adminState"`
	// A device name from the provider configuration.
	Device pulumi.StringPtrOutput `pulumi:"device"`
}

// NewFeatureVnSegment registers a new resource with the given unique name, arguments, and options.
func NewFeatureVnSegment(ctx *pulumi.Context,
	name string, args *FeatureVnSegmentArgs, opts ...pulumi.ResourceOption) (*FeatureVnSegment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AdminState == nil {
		return nil, errors.New("invalid value for required argument 'AdminState'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FeatureVnSegment
	err := ctx.RegisterResource("nxos:nxos/featureVnSegment:FeatureVnSegment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFeatureVnSegment gets an existing FeatureVnSegment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFeatureVnSegment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FeatureVnSegmentState, opts ...pulumi.ResourceOption) (*FeatureVnSegment, error) {
	var resource FeatureVnSegment
	err := ctx.ReadResource("nxos:nxos/featureVnSegment:FeatureVnSegment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FeatureVnSegment resources.
type featureVnSegmentState struct {
	// Administrative state. - Choices: `enabled`, `disabled`
	AdminState *string `pulumi:"adminState"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
}

type FeatureVnSegmentState struct {
	// Administrative state. - Choices: `enabled`, `disabled`
	AdminState pulumi.StringPtrInput
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
}

func (FeatureVnSegmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*featureVnSegmentState)(nil)).Elem()
}

type featureVnSegmentArgs struct {
	// Administrative state. - Choices: `enabled`, `disabled`
	AdminState string `pulumi:"adminState"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
}

// The set of arguments for constructing a FeatureVnSegment resource.
type FeatureVnSegmentArgs struct {
	// Administrative state. - Choices: `enabled`, `disabled`
	AdminState pulumi.StringInput
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
}

func (FeatureVnSegmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*featureVnSegmentArgs)(nil)).Elem()
}

type FeatureVnSegmentInput interface {
	pulumi.Input

	ToFeatureVnSegmentOutput() FeatureVnSegmentOutput
	ToFeatureVnSegmentOutputWithContext(ctx context.Context) FeatureVnSegmentOutput
}

func (*FeatureVnSegment) ElementType() reflect.Type {
	return reflect.TypeOf((**FeatureVnSegment)(nil)).Elem()
}

func (i *FeatureVnSegment) ToFeatureVnSegmentOutput() FeatureVnSegmentOutput {
	return i.ToFeatureVnSegmentOutputWithContext(context.Background())
}

func (i *FeatureVnSegment) ToFeatureVnSegmentOutputWithContext(ctx context.Context) FeatureVnSegmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FeatureVnSegmentOutput)
}

func (i *FeatureVnSegment) ToOutput(ctx context.Context) pulumix.Output[*FeatureVnSegment] {
	return pulumix.Output[*FeatureVnSegment]{
		OutputState: i.ToFeatureVnSegmentOutputWithContext(ctx).OutputState,
	}
}

// FeatureVnSegmentArrayInput is an input type that accepts FeatureVnSegmentArray and FeatureVnSegmentArrayOutput values.
// You can construct a concrete instance of `FeatureVnSegmentArrayInput` via:
//
//	FeatureVnSegmentArray{ FeatureVnSegmentArgs{...} }
type FeatureVnSegmentArrayInput interface {
	pulumi.Input

	ToFeatureVnSegmentArrayOutput() FeatureVnSegmentArrayOutput
	ToFeatureVnSegmentArrayOutputWithContext(context.Context) FeatureVnSegmentArrayOutput
}

type FeatureVnSegmentArray []FeatureVnSegmentInput

func (FeatureVnSegmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FeatureVnSegment)(nil)).Elem()
}

func (i FeatureVnSegmentArray) ToFeatureVnSegmentArrayOutput() FeatureVnSegmentArrayOutput {
	return i.ToFeatureVnSegmentArrayOutputWithContext(context.Background())
}

func (i FeatureVnSegmentArray) ToFeatureVnSegmentArrayOutputWithContext(ctx context.Context) FeatureVnSegmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FeatureVnSegmentArrayOutput)
}

func (i FeatureVnSegmentArray) ToOutput(ctx context.Context) pulumix.Output[[]*FeatureVnSegment] {
	return pulumix.Output[[]*FeatureVnSegment]{
		OutputState: i.ToFeatureVnSegmentArrayOutputWithContext(ctx).OutputState,
	}
}

// FeatureVnSegmentMapInput is an input type that accepts FeatureVnSegmentMap and FeatureVnSegmentMapOutput values.
// You can construct a concrete instance of `FeatureVnSegmentMapInput` via:
//
//	FeatureVnSegmentMap{ "key": FeatureVnSegmentArgs{...} }
type FeatureVnSegmentMapInput interface {
	pulumi.Input

	ToFeatureVnSegmentMapOutput() FeatureVnSegmentMapOutput
	ToFeatureVnSegmentMapOutputWithContext(context.Context) FeatureVnSegmentMapOutput
}

type FeatureVnSegmentMap map[string]FeatureVnSegmentInput

func (FeatureVnSegmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FeatureVnSegment)(nil)).Elem()
}

func (i FeatureVnSegmentMap) ToFeatureVnSegmentMapOutput() FeatureVnSegmentMapOutput {
	return i.ToFeatureVnSegmentMapOutputWithContext(context.Background())
}

func (i FeatureVnSegmentMap) ToFeatureVnSegmentMapOutputWithContext(ctx context.Context) FeatureVnSegmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FeatureVnSegmentMapOutput)
}

func (i FeatureVnSegmentMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*FeatureVnSegment] {
	return pulumix.Output[map[string]*FeatureVnSegment]{
		OutputState: i.ToFeatureVnSegmentMapOutputWithContext(ctx).OutputState,
	}
}

type FeatureVnSegmentOutput struct{ *pulumi.OutputState }

func (FeatureVnSegmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FeatureVnSegment)(nil)).Elem()
}

func (o FeatureVnSegmentOutput) ToFeatureVnSegmentOutput() FeatureVnSegmentOutput {
	return o
}

func (o FeatureVnSegmentOutput) ToFeatureVnSegmentOutputWithContext(ctx context.Context) FeatureVnSegmentOutput {
	return o
}

func (o FeatureVnSegmentOutput) ToOutput(ctx context.Context) pulumix.Output[*FeatureVnSegment] {
	return pulumix.Output[*FeatureVnSegment]{
		OutputState: o.OutputState,
	}
}

// Administrative state. - Choices: `enabled`, `disabled`
func (o FeatureVnSegmentOutput) AdminState() pulumi.StringOutput {
	return o.ApplyT(func(v *FeatureVnSegment) pulumi.StringOutput { return v.AdminState }).(pulumi.StringOutput)
}

// A device name from the provider configuration.
func (o FeatureVnSegmentOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FeatureVnSegment) pulumi.StringPtrOutput { return v.Device }).(pulumi.StringPtrOutput)
}

type FeatureVnSegmentArrayOutput struct{ *pulumi.OutputState }

func (FeatureVnSegmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FeatureVnSegment)(nil)).Elem()
}

func (o FeatureVnSegmentArrayOutput) ToFeatureVnSegmentArrayOutput() FeatureVnSegmentArrayOutput {
	return o
}

func (o FeatureVnSegmentArrayOutput) ToFeatureVnSegmentArrayOutputWithContext(ctx context.Context) FeatureVnSegmentArrayOutput {
	return o
}

func (o FeatureVnSegmentArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*FeatureVnSegment] {
	return pulumix.Output[[]*FeatureVnSegment]{
		OutputState: o.OutputState,
	}
}

func (o FeatureVnSegmentArrayOutput) Index(i pulumi.IntInput) FeatureVnSegmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FeatureVnSegment {
		return vs[0].([]*FeatureVnSegment)[vs[1].(int)]
	}).(FeatureVnSegmentOutput)
}

type FeatureVnSegmentMapOutput struct{ *pulumi.OutputState }

func (FeatureVnSegmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FeatureVnSegment)(nil)).Elem()
}

func (o FeatureVnSegmentMapOutput) ToFeatureVnSegmentMapOutput() FeatureVnSegmentMapOutput {
	return o
}

func (o FeatureVnSegmentMapOutput) ToFeatureVnSegmentMapOutputWithContext(ctx context.Context) FeatureVnSegmentMapOutput {
	return o
}

func (o FeatureVnSegmentMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*FeatureVnSegment] {
	return pulumix.Output[map[string]*FeatureVnSegment]{
		OutputState: o.OutputState,
	}
}

func (o FeatureVnSegmentMapOutput) MapIndex(k pulumi.StringInput) FeatureVnSegmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FeatureVnSegment {
		return vs[0].(map[string]*FeatureVnSegment)[vs[1].(string)]
	}).(FeatureVnSegmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FeatureVnSegmentInput)(nil)).Elem(), &FeatureVnSegment{})
	pulumi.RegisterInputType(reflect.TypeOf((*FeatureVnSegmentArrayInput)(nil)).Elem(), FeatureVnSegmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FeatureVnSegmentMapInput)(nil)).Elem(), FeatureVnSegmentMap{})
	pulumi.RegisterOutputType(FeatureVnSegmentOutput{})
	pulumi.RegisterOutputType(FeatureVnSegmentArrayOutput{})
	pulumi.RegisterOutputType(FeatureVnSegmentMapOutput{})
}
