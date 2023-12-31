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

// This resource can manage the LLDP feature configuration.
//
// - API Documentation: [fmLldp](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Feature%20Management/fm:Lldp/)
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
//			_, err := nxos.NewFeatureLldp(ctx, "example", &nxos.FeatureLldpArgs{
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
//	$ pulumi import nxos:index/featureLldp:FeatureLldp example "sys/fm/lldp"
//
// ```
type FeatureLldp struct {
	pulumi.CustomResourceState

	// Administrative state. - Choices: `enabled`, `disabled`
	AdminState pulumi.StringOutput `pulumi:"adminState"`
	// A device name from the provider configuration.
	Device pulumi.StringPtrOutput `pulumi:"device"`
}

// NewFeatureLldp registers a new resource with the given unique name, arguments, and options.
func NewFeatureLldp(ctx *pulumi.Context,
	name string, args *FeatureLldpArgs, opts ...pulumi.ResourceOption) (*FeatureLldp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AdminState == nil {
		return nil, errors.New("invalid value for required argument 'AdminState'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FeatureLldp
	err := ctx.RegisterResource("nxos:index/featureLldp:FeatureLldp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFeatureLldp gets an existing FeatureLldp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFeatureLldp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FeatureLldpState, opts ...pulumi.ResourceOption) (*FeatureLldp, error) {
	var resource FeatureLldp
	err := ctx.ReadResource("nxos:index/featureLldp:FeatureLldp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FeatureLldp resources.
type featureLldpState struct {
	// Administrative state. - Choices: `enabled`, `disabled`
	AdminState *string `pulumi:"adminState"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
}

type FeatureLldpState struct {
	// Administrative state. - Choices: `enabled`, `disabled`
	AdminState pulumi.StringPtrInput
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
}

func (FeatureLldpState) ElementType() reflect.Type {
	return reflect.TypeOf((*featureLldpState)(nil)).Elem()
}

type featureLldpArgs struct {
	// Administrative state. - Choices: `enabled`, `disabled`
	AdminState string `pulumi:"adminState"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
}

// The set of arguments for constructing a FeatureLldp resource.
type FeatureLldpArgs struct {
	// Administrative state. - Choices: `enabled`, `disabled`
	AdminState pulumi.StringInput
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
}

func (FeatureLldpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*featureLldpArgs)(nil)).Elem()
}

type FeatureLldpInput interface {
	pulumi.Input

	ToFeatureLldpOutput() FeatureLldpOutput
	ToFeatureLldpOutputWithContext(ctx context.Context) FeatureLldpOutput
}

func (*FeatureLldp) ElementType() reflect.Type {
	return reflect.TypeOf((**FeatureLldp)(nil)).Elem()
}

func (i *FeatureLldp) ToFeatureLldpOutput() FeatureLldpOutput {
	return i.ToFeatureLldpOutputWithContext(context.Background())
}

func (i *FeatureLldp) ToFeatureLldpOutputWithContext(ctx context.Context) FeatureLldpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FeatureLldpOutput)
}

func (i *FeatureLldp) ToOutput(ctx context.Context) pulumix.Output[*FeatureLldp] {
	return pulumix.Output[*FeatureLldp]{
		OutputState: i.ToFeatureLldpOutputWithContext(ctx).OutputState,
	}
}

// FeatureLldpArrayInput is an input type that accepts FeatureLldpArray and FeatureLldpArrayOutput values.
// You can construct a concrete instance of `FeatureLldpArrayInput` via:
//
//	FeatureLldpArray{ FeatureLldpArgs{...} }
type FeatureLldpArrayInput interface {
	pulumi.Input

	ToFeatureLldpArrayOutput() FeatureLldpArrayOutput
	ToFeatureLldpArrayOutputWithContext(context.Context) FeatureLldpArrayOutput
}

type FeatureLldpArray []FeatureLldpInput

func (FeatureLldpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FeatureLldp)(nil)).Elem()
}

func (i FeatureLldpArray) ToFeatureLldpArrayOutput() FeatureLldpArrayOutput {
	return i.ToFeatureLldpArrayOutputWithContext(context.Background())
}

func (i FeatureLldpArray) ToFeatureLldpArrayOutputWithContext(ctx context.Context) FeatureLldpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FeatureLldpArrayOutput)
}

func (i FeatureLldpArray) ToOutput(ctx context.Context) pulumix.Output[[]*FeatureLldp] {
	return pulumix.Output[[]*FeatureLldp]{
		OutputState: i.ToFeatureLldpArrayOutputWithContext(ctx).OutputState,
	}
}

// FeatureLldpMapInput is an input type that accepts FeatureLldpMap and FeatureLldpMapOutput values.
// You can construct a concrete instance of `FeatureLldpMapInput` via:
//
//	FeatureLldpMap{ "key": FeatureLldpArgs{...} }
type FeatureLldpMapInput interface {
	pulumi.Input

	ToFeatureLldpMapOutput() FeatureLldpMapOutput
	ToFeatureLldpMapOutputWithContext(context.Context) FeatureLldpMapOutput
}

type FeatureLldpMap map[string]FeatureLldpInput

func (FeatureLldpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FeatureLldp)(nil)).Elem()
}

func (i FeatureLldpMap) ToFeatureLldpMapOutput() FeatureLldpMapOutput {
	return i.ToFeatureLldpMapOutputWithContext(context.Background())
}

func (i FeatureLldpMap) ToFeatureLldpMapOutputWithContext(ctx context.Context) FeatureLldpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FeatureLldpMapOutput)
}

func (i FeatureLldpMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*FeatureLldp] {
	return pulumix.Output[map[string]*FeatureLldp]{
		OutputState: i.ToFeatureLldpMapOutputWithContext(ctx).OutputState,
	}
}

type FeatureLldpOutput struct{ *pulumi.OutputState }

func (FeatureLldpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FeatureLldp)(nil)).Elem()
}

func (o FeatureLldpOutput) ToFeatureLldpOutput() FeatureLldpOutput {
	return o
}

func (o FeatureLldpOutput) ToFeatureLldpOutputWithContext(ctx context.Context) FeatureLldpOutput {
	return o
}

func (o FeatureLldpOutput) ToOutput(ctx context.Context) pulumix.Output[*FeatureLldp] {
	return pulumix.Output[*FeatureLldp]{
		OutputState: o.OutputState,
	}
}

// Administrative state. - Choices: `enabled`, `disabled`
func (o FeatureLldpOutput) AdminState() pulumi.StringOutput {
	return o.ApplyT(func(v *FeatureLldp) pulumi.StringOutput { return v.AdminState }).(pulumi.StringOutput)
}

// A device name from the provider configuration.
func (o FeatureLldpOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FeatureLldp) pulumi.StringPtrOutput { return v.Device }).(pulumi.StringPtrOutput)
}

type FeatureLldpArrayOutput struct{ *pulumi.OutputState }

func (FeatureLldpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FeatureLldp)(nil)).Elem()
}

func (o FeatureLldpArrayOutput) ToFeatureLldpArrayOutput() FeatureLldpArrayOutput {
	return o
}

func (o FeatureLldpArrayOutput) ToFeatureLldpArrayOutputWithContext(ctx context.Context) FeatureLldpArrayOutput {
	return o
}

func (o FeatureLldpArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*FeatureLldp] {
	return pulumix.Output[[]*FeatureLldp]{
		OutputState: o.OutputState,
	}
}

func (o FeatureLldpArrayOutput) Index(i pulumi.IntInput) FeatureLldpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FeatureLldp {
		return vs[0].([]*FeatureLldp)[vs[1].(int)]
	}).(FeatureLldpOutput)
}

type FeatureLldpMapOutput struct{ *pulumi.OutputState }

func (FeatureLldpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FeatureLldp)(nil)).Elem()
}

func (o FeatureLldpMapOutput) ToFeatureLldpMapOutput() FeatureLldpMapOutput {
	return o
}

func (o FeatureLldpMapOutput) ToFeatureLldpMapOutputWithContext(ctx context.Context) FeatureLldpMapOutput {
	return o
}

func (o FeatureLldpMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*FeatureLldp] {
	return pulumix.Output[map[string]*FeatureLldp]{
		OutputState: o.OutputState,
	}
}

func (o FeatureLldpMapOutput) MapIndex(k pulumi.StringInput) FeatureLldpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FeatureLldp {
		return vs[0].(map[string]*FeatureLldp)[vs[1].(string)]
	}).(FeatureLldpOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FeatureLldpInput)(nil)).Elem(), &FeatureLldp{})
	pulumi.RegisterInputType(reflect.TypeOf((*FeatureLldpArrayInput)(nil)).Elem(), FeatureLldpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FeatureLldpMapInput)(nil)).Elem(), FeatureLldpMap{})
	pulumi.RegisterOutputType(FeatureLldpOutput{})
	pulumi.RegisterOutputType(FeatureLldpArrayOutput{})
	pulumi.RegisterOutputType(FeatureLldpMapOutput{})
}
