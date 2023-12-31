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

// This resource can manage the PIM feature configuration.
//
// - API Documentation: [fmPim](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Feature%20Management/fm:Pim/)
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
//			_, err := nxos.NewFeaturePim(ctx, "example", &nxos.FeaturePimArgs{
//				AdminState: pulumi.String("enabled"),
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
//	$ pulumi import nxos:index/featurePim:FeaturePim example "sys/fm/pim"
//
// ```
type FeaturePim struct {
	pulumi.CustomResourceState

	// Administrative state. - Choices: `enabled`, `disabled`
	AdminState pulumi.StringOutput `pulumi:"adminState"`
	// A device name from the provider configuration.
	Device pulumi.StringPtrOutput `pulumi:"device"`
}

// NewFeaturePim registers a new resource with the given unique name, arguments, and options.
func NewFeaturePim(ctx *pulumi.Context,
	name string, args *FeaturePimArgs, opts ...pulumi.ResourceOption) (*FeaturePim, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AdminState == nil {
		return nil, errors.New("invalid value for required argument 'AdminState'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FeaturePim
	err := ctx.RegisterResource("nxos:index/featurePim:FeaturePim", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFeaturePim gets an existing FeaturePim resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFeaturePim(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FeaturePimState, opts ...pulumi.ResourceOption) (*FeaturePim, error) {
	var resource FeaturePim
	err := ctx.ReadResource("nxos:index/featurePim:FeaturePim", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FeaturePim resources.
type featurePimState struct {
	// Administrative state. - Choices: `enabled`, `disabled`
	AdminState *string `pulumi:"adminState"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
}

type FeaturePimState struct {
	// Administrative state. - Choices: `enabled`, `disabled`
	AdminState pulumi.StringPtrInput
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
}

func (FeaturePimState) ElementType() reflect.Type {
	return reflect.TypeOf((*featurePimState)(nil)).Elem()
}

type featurePimArgs struct {
	// Administrative state. - Choices: `enabled`, `disabled`
	AdminState string `pulumi:"adminState"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
}

// The set of arguments for constructing a FeaturePim resource.
type FeaturePimArgs struct {
	// Administrative state. - Choices: `enabled`, `disabled`
	AdminState pulumi.StringInput
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
}

func (FeaturePimArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*featurePimArgs)(nil)).Elem()
}

type FeaturePimInput interface {
	pulumi.Input

	ToFeaturePimOutput() FeaturePimOutput
	ToFeaturePimOutputWithContext(ctx context.Context) FeaturePimOutput
}

func (*FeaturePim) ElementType() reflect.Type {
	return reflect.TypeOf((**FeaturePim)(nil)).Elem()
}

func (i *FeaturePim) ToFeaturePimOutput() FeaturePimOutput {
	return i.ToFeaturePimOutputWithContext(context.Background())
}

func (i *FeaturePim) ToFeaturePimOutputWithContext(ctx context.Context) FeaturePimOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FeaturePimOutput)
}

func (i *FeaturePim) ToOutput(ctx context.Context) pulumix.Output[*FeaturePim] {
	return pulumix.Output[*FeaturePim]{
		OutputState: i.ToFeaturePimOutputWithContext(ctx).OutputState,
	}
}

// FeaturePimArrayInput is an input type that accepts FeaturePimArray and FeaturePimArrayOutput values.
// You can construct a concrete instance of `FeaturePimArrayInput` via:
//
//	FeaturePimArray{ FeaturePimArgs{...} }
type FeaturePimArrayInput interface {
	pulumi.Input

	ToFeaturePimArrayOutput() FeaturePimArrayOutput
	ToFeaturePimArrayOutputWithContext(context.Context) FeaturePimArrayOutput
}

type FeaturePimArray []FeaturePimInput

func (FeaturePimArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FeaturePim)(nil)).Elem()
}

func (i FeaturePimArray) ToFeaturePimArrayOutput() FeaturePimArrayOutput {
	return i.ToFeaturePimArrayOutputWithContext(context.Background())
}

func (i FeaturePimArray) ToFeaturePimArrayOutputWithContext(ctx context.Context) FeaturePimArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FeaturePimArrayOutput)
}

func (i FeaturePimArray) ToOutput(ctx context.Context) pulumix.Output[[]*FeaturePim] {
	return pulumix.Output[[]*FeaturePim]{
		OutputState: i.ToFeaturePimArrayOutputWithContext(ctx).OutputState,
	}
}

// FeaturePimMapInput is an input type that accepts FeaturePimMap and FeaturePimMapOutput values.
// You can construct a concrete instance of `FeaturePimMapInput` via:
//
//	FeaturePimMap{ "key": FeaturePimArgs{...} }
type FeaturePimMapInput interface {
	pulumi.Input

	ToFeaturePimMapOutput() FeaturePimMapOutput
	ToFeaturePimMapOutputWithContext(context.Context) FeaturePimMapOutput
}

type FeaturePimMap map[string]FeaturePimInput

func (FeaturePimMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FeaturePim)(nil)).Elem()
}

func (i FeaturePimMap) ToFeaturePimMapOutput() FeaturePimMapOutput {
	return i.ToFeaturePimMapOutputWithContext(context.Background())
}

func (i FeaturePimMap) ToFeaturePimMapOutputWithContext(ctx context.Context) FeaturePimMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FeaturePimMapOutput)
}

func (i FeaturePimMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*FeaturePim] {
	return pulumix.Output[map[string]*FeaturePim]{
		OutputState: i.ToFeaturePimMapOutputWithContext(ctx).OutputState,
	}
}

type FeaturePimOutput struct{ *pulumi.OutputState }

func (FeaturePimOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FeaturePim)(nil)).Elem()
}

func (o FeaturePimOutput) ToFeaturePimOutput() FeaturePimOutput {
	return o
}

func (o FeaturePimOutput) ToFeaturePimOutputWithContext(ctx context.Context) FeaturePimOutput {
	return o
}

func (o FeaturePimOutput) ToOutput(ctx context.Context) pulumix.Output[*FeaturePim] {
	return pulumix.Output[*FeaturePim]{
		OutputState: o.OutputState,
	}
}

// Administrative state. - Choices: `enabled`, `disabled`
func (o FeaturePimOutput) AdminState() pulumi.StringOutput {
	return o.ApplyT(func(v *FeaturePim) pulumi.StringOutput { return v.AdminState }).(pulumi.StringOutput)
}

// A device name from the provider configuration.
func (o FeaturePimOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FeaturePim) pulumi.StringPtrOutput { return v.Device }).(pulumi.StringPtrOutput)
}

type FeaturePimArrayOutput struct{ *pulumi.OutputState }

func (FeaturePimArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FeaturePim)(nil)).Elem()
}

func (o FeaturePimArrayOutput) ToFeaturePimArrayOutput() FeaturePimArrayOutput {
	return o
}

func (o FeaturePimArrayOutput) ToFeaturePimArrayOutputWithContext(ctx context.Context) FeaturePimArrayOutput {
	return o
}

func (o FeaturePimArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*FeaturePim] {
	return pulumix.Output[[]*FeaturePim]{
		OutputState: o.OutputState,
	}
}

func (o FeaturePimArrayOutput) Index(i pulumi.IntInput) FeaturePimOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FeaturePim {
		return vs[0].([]*FeaturePim)[vs[1].(int)]
	}).(FeaturePimOutput)
}

type FeaturePimMapOutput struct{ *pulumi.OutputState }

func (FeaturePimMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FeaturePim)(nil)).Elem()
}

func (o FeaturePimMapOutput) ToFeaturePimMapOutput() FeaturePimMapOutput {
	return o
}

func (o FeaturePimMapOutput) ToFeaturePimMapOutputWithContext(ctx context.Context) FeaturePimMapOutput {
	return o
}

func (o FeaturePimMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*FeaturePim] {
	return pulumix.Output[map[string]*FeaturePim]{
		OutputState: o.OutputState,
	}
}

func (o FeaturePimMapOutput) MapIndex(k pulumi.StringInput) FeaturePimOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FeaturePim {
		return vs[0].(map[string]*FeaturePim)[vs[1].(string)]
	}).(FeaturePimOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FeaturePimInput)(nil)).Elem(), &FeaturePim{})
	pulumi.RegisterInputType(reflect.TypeOf((*FeaturePimArrayInput)(nil)).Elem(), FeaturePimArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FeaturePimMapInput)(nil)).Elem(), FeaturePimMap{})
	pulumi.RegisterOutputType(FeaturePimOutput{})
	pulumi.RegisterOutputType(FeaturePimArrayOutput{})
	pulumi.RegisterOutputType(FeaturePimMapOutput{})
}
