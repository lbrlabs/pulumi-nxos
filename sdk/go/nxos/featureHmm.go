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

// This resource can manage the HMM feature (aka `feature fabric forwarding`) configuration.
//
// - API Documentation: [fmHmm](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Feature%20Management/fm:Hmm/)
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
//			_, err := nxos.NewFeatureHmm(ctx, "example", &nxos.FeatureHmmArgs{
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
//	$ pulumi import nxos:index/featureHmm:FeatureHmm example "sys/fm/hmm"
//
// ```
type FeatureHmm struct {
	pulumi.CustomResourceState

	// Administrative state. - Choices: `enabled`, `disabled`
	AdminState pulumi.StringOutput `pulumi:"adminState"`
	// A device name from the provider configuration.
	Device pulumi.StringPtrOutput `pulumi:"device"`
}

// NewFeatureHmm registers a new resource with the given unique name, arguments, and options.
func NewFeatureHmm(ctx *pulumi.Context,
	name string, args *FeatureHmmArgs, opts ...pulumi.ResourceOption) (*FeatureHmm, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AdminState == nil {
		return nil, errors.New("invalid value for required argument 'AdminState'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FeatureHmm
	err := ctx.RegisterResource("nxos:index/featureHmm:FeatureHmm", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFeatureHmm gets an existing FeatureHmm resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFeatureHmm(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FeatureHmmState, opts ...pulumi.ResourceOption) (*FeatureHmm, error) {
	var resource FeatureHmm
	err := ctx.ReadResource("nxos:index/featureHmm:FeatureHmm", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FeatureHmm resources.
type featureHmmState struct {
	// Administrative state. - Choices: `enabled`, `disabled`
	AdminState *string `pulumi:"adminState"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
}

type FeatureHmmState struct {
	// Administrative state. - Choices: `enabled`, `disabled`
	AdminState pulumi.StringPtrInput
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
}

func (FeatureHmmState) ElementType() reflect.Type {
	return reflect.TypeOf((*featureHmmState)(nil)).Elem()
}

type featureHmmArgs struct {
	// Administrative state. - Choices: `enabled`, `disabled`
	AdminState string `pulumi:"adminState"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
}

// The set of arguments for constructing a FeatureHmm resource.
type FeatureHmmArgs struct {
	// Administrative state. - Choices: `enabled`, `disabled`
	AdminState pulumi.StringInput
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
}

func (FeatureHmmArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*featureHmmArgs)(nil)).Elem()
}

type FeatureHmmInput interface {
	pulumi.Input

	ToFeatureHmmOutput() FeatureHmmOutput
	ToFeatureHmmOutputWithContext(ctx context.Context) FeatureHmmOutput
}

func (*FeatureHmm) ElementType() reflect.Type {
	return reflect.TypeOf((**FeatureHmm)(nil)).Elem()
}

func (i *FeatureHmm) ToFeatureHmmOutput() FeatureHmmOutput {
	return i.ToFeatureHmmOutputWithContext(context.Background())
}

func (i *FeatureHmm) ToFeatureHmmOutputWithContext(ctx context.Context) FeatureHmmOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FeatureHmmOutput)
}

func (i *FeatureHmm) ToOutput(ctx context.Context) pulumix.Output[*FeatureHmm] {
	return pulumix.Output[*FeatureHmm]{
		OutputState: i.ToFeatureHmmOutputWithContext(ctx).OutputState,
	}
}

// FeatureHmmArrayInput is an input type that accepts FeatureHmmArray and FeatureHmmArrayOutput values.
// You can construct a concrete instance of `FeatureHmmArrayInput` via:
//
//	FeatureHmmArray{ FeatureHmmArgs{...} }
type FeatureHmmArrayInput interface {
	pulumi.Input

	ToFeatureHmmArrayOutput() FeatureHmmArrayOutput
	ToFeatureHmmArrayOutputWithContext(context.Context) FeatureHmmArrayOutput
}

type FeatureHmmArray []FeatureHmmInput

func (FeatureHmmArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FeatureHmm)(nil)).Elem()
}

func (i FeatureHmmArray) ToFeatureHmmArrayOutput() FeatureHmmArrayOutput {
	return i.ToFeatureHmmArrayOutputWithContext(context.Background())
}

func (i FeatureHmmArray) ToFeatureHmmArrayOutputWithContext(ctx context.Context) FeatureHmmArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FeatureHmmArrayOutput)
}

func (i FeatureHmmArray) ToOutput(ctx context.Context) pulumix.Output[[]*FeatureHmm] {
	return pulumix.Output[[]*FeatureHmm]{
		OutputState: i.ToFeatureHmmArrayOutputWithContext(ctx).OutputState,
	}
}

// FeatureHmmMapInput is an input type that accepts FeatureHmmMap and FeatureHmmMapOutput values.
// You can construct a concrete instance of `FeatureHmmMapInput` via:
//
//	FeatureHmmMap{ "key": FeatureHmmArgs{...} }
type FeatureHmmMapInput interface {
	pulumi.Input

	ToFeatureHmmMapOutput() FeatureHmmMapOutput
	ToFeatureHmmMapOutputWithContext(context.Context) FeatureHmmMapOutput
}

type FeatureHmmMap map[string]FeatureHmmInput

func (FeatureHmmMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FeatureHmm)(nil)).Elem()
}

func (i FeatureHmmMap) ToFeatureHmmMapOutput() FeatureHmmMapOutput {
	return i.ToFeatureHmmMapOutputWithContext(context.Background())
}

func (i FeatureHmmMap) ToFeatureHmmMapOutputWithContext(ctx context.Context) FeatureHmmMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FeatureHmmMapOutput)
}

func (i FeatureHmmMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*FeatureHmm] {
	return pulumix.Output[map[string]*FeatureHmm]{
		OutputState: i.ToFeatureHmmMapOutputWithContext(ctx).OutputState,
	}
}

type FeatureHmmOutput struct{ *pulumi.OutputState }

func (FeatureHmmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FeatureHmm)(nil)).Elem()
}

func (o FeatureHmmOutput) ToFeatureHmmOutput() FeatureHmmOutput {
	return o
}

func (o FeatureHmmOutput) ToFeatureHmmOutputWithContext(ctx context.Context) FeatureHmmOutput {
	return o
}

func (o FeatureHmmOutput) ToOutput(ctx context.Context) pulumix.Output[*FeatureHmm] {
	return pulumix.Output[*FeatureHmm]{
		OutputState: o.OutputState,
	}
}

// Administrative state. - Choices: `enabled`, `disabled`
func (o FeatureHmmOutput) AdminState() pulumi.StringOutput {
	return o.ApplyT(func(v *FeatureHmm) pulumi.StringOutput { return v.AdminState }).(pulumi.StringOutput)
}

// A device name from the provider configuration.
func (o FeatureHmmOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FeatureHmm) pulumi.StringPtrOutput { return v.Device }).(pulumi.StringPtrOutput)
}

type FeatureHmmArrayOutput struct{ *pulumi.OutputState }

func (FeatureHmmArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FeatureHmm)(nil)).Elem()
}

func (o FeatureHmmArrayOutput) ToFeatureHmmArrayOutput() FeatureHmmArrayOutput {
	return o
}

func (o FeatureHmmArrayOutput) ToFeatureHmmArrayOutputWithContext(ctx context.Context) FeatureHmmArrayOutput {
	return o
}

func (o FeatureHmmArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*FeatureHmm] {
	return pulumix.Output[[]*FeatureHmm]{
		OutputState: o.OutputState,
	}
}

func (o FeatureHmmArrayOutput) Index(i pulumi.IntInput) FeatureHmmOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FeatureHmm {
		return vs[0].([]*FeatureHmm)[vs[1].(int)]
	}).(FeatureHmmOutput)
}

type FeatureHmmMapOutput struct{ *pulumi.OutputState }

func (FeatureHmmMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FeatureHmm)(nil)).Elem()
}

func (o FeatureHmmMapOutput) ToFeatureHmmMapOutput() FeatureHmmMapOutput {
	return o
}

func (o FeatureHmmMapOutput) ToFeatureHmmMapOutputWithContext(ctx context.Context) FeatureHmmMapOutput {
	return o
}

func (o FeatureHmmMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*FeatureHmm] {
	return pulumix.Output[map[string]*FeatureHmm]{
		OutputState: o.OutputState,
	}
}

func (o FeatureHmmMapOutput) MapIndex(k pulumi.StringInput) FeatureHmmOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FeatureHmm {
		return vs[0].(map[string]*FeatureHmm)[vs[1].(string)]
	}).(FeatureHmmOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FeatureHmmInput)(nil)).Elem(), &FeatureHmm{})
	pulumi.RegisterInputType(reflect.TypeOf((*FeatureHmmArrayInput)(nil)).Elem(), FeatureHmmArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FeatureHmmMapInput)(nil)).Elem(), FeatureHmmMap{})
	pulumi.RegisterOutputType(FeatureHmmOutput{})
	pulumi.RegisterOutputType(FeatureHmmArrayOutput{})
	pulumi.RegisterOutputType(FeatureHmmMapOutput{})
}
