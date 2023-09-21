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

type FeatureOspfv3 struct {
	pulumi.CustomResourceState

	// Administrative state. - Choices: `enabled`, `disabled`
	AdminState pulumi.StringOutput `pulumi:"adminState"`
	// A device name from the provider configuration.
	Device pulumi.StringPtrOutput `pulumi:"device"`
}

// NewFeatureOspfv3 registers a new resource with the given unique name, arguments, and options.
func NewFeatureOspfv3(ctx *pulumi.Context,
	name string, args *FeatureOspfv3Args, opts ...pulumi.ResourceOption) (*FeatureOspfv3, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AdminState == nil {
		return nil, errors.New("invalid value for required argument 'AdminState'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FeatureOspfv3
	err := ctx.RegisterResource("nxos:nxos/featureOspfv3:FeatureOspfv3", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFeatureOspfv3 gets an existing FeatureOspfv3 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFeatureOspfv3(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FeatureOspfv3State, opts ...pulumi.ResourceOption) (*FeatureOspfv3, error) {
	var resource FeatureOspfv3
	err := ctx.ReadResource("nxos:nxos/featureOspfv3:FeatureOspfv3", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FeatureOspfv3 resources.
type featureOspfv3State struct {
	// Administrative state. - Choices: `enabled`, `disabled`
	AdminState *string `pulumi:"adminState"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
}

type FeatureOspfv3State struct {
	// Administrative state. - Choices: `enabled`, `disabled`
	AdminState pulumi.StringPtrInput
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
}

func (FeatureOspfv3State) ElementType() reflect.Type {
	return reflect.TypeOf((*featureOspfv3State)(nil)).Elem()
}

type featureOspfv3Args struct {
	// Administrative state. - Choices: `enabled`, `disabled`
	AdminState string `pulumi:"adminState"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
}

// The set of arguments for constructing a FeatureOspfv3 resource.
type FeatureOspfv3Args struct {
	// Administrative state. - Choices: `enabled`, `disabled`
	AdminState pulumi.StringInput
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
}

func (FeatureOspfv3Args) ElementType() reflect.Type {
	return reflect.TypeOf((*featureOspfv3Args)(nil)).Elem()
}

type FeatureOspfv3Input interface {
	pulumi.Input

	ToFeatureOspfv3Output() FeatureOspfv3Output
	ToFeatureOspfv3OutputWithContext(ctx context.Context) FeatureOspfv3Output
}

func (*FeatureOspfv3) ElementType() reflect.Type {
	return reflect.TypeOf((**FeatureOspfv3)(nil)).Elem()
}

func (i *FeatureOspfv3) ToFeatureOspfv3Output() FeatureOspfv3Output {
	return i.ToFeatureOspfv3OutputWithContext(context.Background())
}

func (i *FeatureOspfv3) ToFeatureOspfv3OutputWithContext(ctx context.Context) FeatureOspfv3Output {
	return pulumi.ToOutputWithContext(ctx, i).(FeatureOspfv3Output)
}

func (i *FeatureOspfv3) ToOutput(ctx context.Context) pulumix.Output[*FeatureOspfv3] {
	return pulumix.Output[*FeatureOspfv3]{
		OutputState: i.ToFeatureOspfv3OutputWithContext(ctx).OutputState,
	}
}

// FeatureOspfv3ArrayInput is an input type that accepts FeatureOspfv3Array and FeatureOspfv3ArrayOutput values.
// You can construct a concrete instance of `FeatureOspfv3ArrayInput` via:
//
//	FeatureOspfv3Array{ FeatureOspfv3Args{...} }
type FeatureOspfv3ArrayInput interface {
	pulumi.Input

	ToFeatureOspfv3ArrayOutput() FeatureOspfv3ArrayOutput
	ToFeatureOspfv3ArrayOutputWithContext(context.Context) FeatureOspfv3ArrayOutput
}

type FeatureOspfv3Array []FeatureOspfv3Input

func (FeatureOspfv3Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FeatureOspfv3)(nil)).Elem()
}

func (i FeatureOspfv3Array) ToFeatureOspfv3ArrayOutput() FeatureOspfv3ArrayOutput {
	return i.ToFeatureOspfv3ArrayOutputWithContext(context.Background())
}

func (i FeatureOspfv3Array) ToFeatureOspfv3ArrayOutputWithContext(ctx context.Context) FeatureOspfv3ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FeatureOspfv3ArrayOutput)
}

func (i FeatureOspfv3Array) ToOutput(ctx context.Context) pulumix.Output[[]*FeatureOspfv3] {
	return pulumix.Output[[]*FeatureOspfv3]{
		OutputState: i.ToFeatureOspfv3ArrayOutputWithContext(ctx).OutputState,
	}
}

// FeatureOspfv3MapInput is an input type that accepts FeatureOspfv3Map and FeatureOspfv3MapOutput values.
// You can construct a concrete instance of `FeatureOspfv3MapInput` via:
//
//	FeatureOspfv3Map{ "key": FeatureOspfv3Args{...} }
type FeatureOspfv3MapInput interface {
	pulumi.Input

	ToFeatureOspfv3MapOutput() FeatureOspfv3MapOutput
	ToFeatureOspfv3MapOutputWithContext(context.Context) FeatureOspfv3MapOutput
}

type FeatureOspfv3Map map[string]FeatureOspfv3Input

func (FeatureOspfv3Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FeatureOspfv3)(nil)).Elem()
}

func (i FeatureOspfv3Map) ToFeatureOspfv3MapOutput() FeatureOspfv3MapOutput {
	return i.ToFeatureOspfv3MapOutputWithContext(context.Background())
}

func (i FeatureOspfv3Map) ToFeatureOspfv3MapOutputWithContext(ctx context.Context) FeatureOspfv3MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FeatureOspfv3MapOutput)
}

func (i FeatureOspfv3Map) ToOutput(ctx context.Context) pulumix.Output[map[string]*FeatureOspfv3] {
	return pulumix.Output[map[string]*FeatureOspfv3]{
		OutputState: i.ToFeatureOspfv3MapOutputWithContext(ctx).OutputState,
	}
}

type FeatureOspfv3Output struct{ *pulumi.OutputState }

func (FeatureOspfv3Output) ElementType() reflect.Type {
	return reflect.TypeOf((**FeatureOspfv3)(nil)).Elem()
}

func (o FeatureOspfv3Output) ToFeatureOspfv3Output() FeatureOspfv3Output {
	return o
}

func (o FeatureOspfv3Output) ToFeatureOspfv3OutputWithContext(ctx context.Context) FeatureOspfv3Output {
	return o
}

func (o FeatureOspfv3Output) ToOutput(ctx context.Context) pulumix.Output[*FeatureOspfv3] {
	return pulumix.Output[*FeatureOspfv3]{
		OutputState: o.OutputState,
	}
}

// Administrative state. - Choices: `enabled`, `disabled`
func (o FeatureOspfv3Output) AdminState() pulumi.StringOutput {
	return o.ApplyT(func(v *FeatureOspfv3) pulumi.StringOutput { return v.AdminState }).(pulumi.StringOutput)
}

// A device name from the provider configuration.
func (o FeatureOspfv3Output) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FeatureOspfv3) pulumi.StringPtrOutput { return v.Device }).(pulumi.StringPtrOutput)
}

type FeatureOspfv3ArrayOutput struct{ *pulumi.OutputState }

func (FeatureOspfv3ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FeatureOspfv3)(nil)).Elem()
}

func (o FeatureOspfv3ArrayOutput) ToFeatureOspfv3ArrayOutput() FeatureOspfv3ArrayOutput {
	return o
}

func (o FeatureOspfv3ArrayOutput) ToFeatureOspfv3ArrayOutputWithContext(ctx context.Context) FeatureOspfv3ArrayOutput {
	return o
}

func (o FeatureOspfv3ArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*FeatureOspfv3] {
	return pulumix.Output[[]*FeatureOspfv3]{
		OutputState: o.OutputState,
	}
}

func (o FeatureOspfv3ArrayOutput) Index(i pulumi.IntInput) FeatureOspfv3Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FeatureOspfv3 {
		return vs[0].([]*FeatureOspfv3)[vs[1].(int)]
	}).(FeatureOspfv3Output)
}

type FeatureOspfv3MapOutput struct{ *pulumi.OutputState }

func (FeatureOspfv3MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FeatureOspfv3)(nil)).Elem()
}

func (o FeatureOspfv3MapOutput) ToFeatureOspfv3MapOutput() FeatureOspfv3MapOutput {
	return o
}

func (o FeatureOspfv3MapOutput) ToFeatureOspfv3MapOutputWithContext(ctx context.Context) FeatureOspfv3MapOutput {
	return o
}

func (o FeatureOspfv3MapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*FeatureOspfv3] {
	return pulumix.Output[map[string]*FeatureOspfv3]{
		OutputState: o.OutputState,
	}
}

func (o FeatureOspfv3MapOutput) MapIndex(k pulumi.StringInput) FeatureOspfv3Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FeatureOspfv3 {
		return vs[0].(map[string]*FeatureOspfv3)[vs[1].(string)]
	}).(FeatureOspfv3Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FeatureOspfv3Input)(nil)).Elem(), &FeatureOspfv3{})
	pulumi.RegisterInputType(reflect.TypeOf((*FeatureOspfv3ArrayInput)(nil)).Elem(), FeatureOspfv3Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*FeatureOspfv3MapInput)(nil)).Elem(), FeatureOspfv3Map{})
	pulumi.RegisterOutputType(FeatureOspfv3Output{})
	pulumi.RegisterOutputType(FeatureOspfv3ArrayOutput{})
	pulumi.RegisterOutputType(FeatureOspfv3MapOutput{})
}
