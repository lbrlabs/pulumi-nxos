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

// This resource can manage the queuing QoS policy map match class map priority configuration.
//
// - API Documentation: [ipqosPriority](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Qos/ipqos:Priority/)
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
//			_, err := nxos.NewQueuingQosPolicyMapMatchClassMapPriority(ctx, "example", &nxos.QueuingQosPolicyMapMatchClassMapPriorityArgs{
//				ClassMapName:  pulumi.String("c-out-q1"),
//				Level:         pulumi.Int(1),
//				PolicyMapName: pulumi.String("PM1"),
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
//	$ pulumi import nxos:index/queuingQosPolicyMapMatchClassMapPriority:QueuingQosPolicyMapMatchClassMapPriority example "sys/ipqos/queuing/p/name-[PM1]/cmap-[c-out-q1]/prio"
//
// ```
type QueuingQosPolicyMapMatchClassMapPriority struct {
	pulumi.CustomResourceState

	// Class map name.
	ClassMapName pulumi.StringOutput `pulumi:"classMapName"`
	// A device name from the provider configuration.
	Device pulumi.StringPtrOutput `pulumi:"device"`
	// Priority level. - Range: `1`-`8`
	Level pulumi.IntOutput `pulumi:"level"`
	// Policy map name.
	PolicyMapName pulumi.StringOutput `pulumi:"policyMapName"`
}

// NewQueuingQosPolicyMapMatchClassMapPriority registers a new resource with the given unique name, arguments, and options.
func NewQueuingQosPolicyMapMatchClassMapPriority(ctx *pulumi.Context,
	name string, args *QueuingQosPolicyMapMatchClassMapPriorityArgs, opts ...pulumi.ResourceOption) (*QueuingQosPolicyMapMatchClassMapPriority, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClassMapName == nil {
		return nil, errors.New("invalid value for required argument 'ClassMapName'")
	}
	if args.Level == nil {
		return nil, errors.New("invalid value for required argument 'Level'")
	}
	if args.PolicyMapName == nil {
		return nil, errors.New("invalid value for required argument 'PolicyMapName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource QueuingQosPolicyMapMatchClassMapPriority
	err := ctx.RegisterResource("nxos:index/queuingQosPolicyMapMatchClassMapPriority:QueuingQosPolicyMapMatchClassMapPriority", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQueuingQosPolicyMapMatchClassMapPriority gets an existing QueuingQosPolicyMapMatchClassMapPriority resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQueuingQosPolicyMapMatchClassMapPriority(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QueuingQosPolicyMapMatchClassMapPriorityState, opts ...pulumi.ResourceOption) (*QueuingQosPolicyMapMatchClassMapPriority, error) {
	var resource QueuingQosPolicyMapMatchClassMapPriority
	err := ctx.ReadResource("nxos:index/queuingQosPolicyMapMatchClassMapPriority:QueuingQosPolicyMapMatchClassMapPriority", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering QueuingQosPolicyMapMatchClassMapPriority resources.
type queuingQosPolicyMapMatchClassMapPriorityState struct {
	// Class map name.
	ClassMapName *string `pulumi:"classMapName"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// Priority level. - Range: `1`-`8`
	Level *int `pulumi:"level"`
	// Policy map name.
	PolicyMapName *string `pulumi:"policyMapName"`
}

type QueuingQosPolicyMapMatchClassMapPriorityState struct {
	// Class map name.
	ClassMapName pulumi.StringPtrInput
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
	// Priority level. - Range: `1`-`8`
	Level pulumi.IntPtrInput
	// Policy map name.
	PolicyMapName pulumi.StringPtrInput
}

func (QueuingQosPolicyMapMatchClassMapPriorityState) ElementType() reflect.Type {
	return reflect.TypeOf((*queuingQosPolicyMapMatchClassMapPriorityState)(nil)).Elem()
}

type queuingQosPolicyMapMatchClassMapPriorityArgs struct {
	// Class map name.
	ClassMapName string `pulumi:"classMapName"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// Priority level. - Range: `1`-`8`
	Level int `pulumi:"level"`
	// Policy map name.
	PolicyMapName string `pulumi:"policyMapName"`
}

// The set of arguments for constructing a QueuingQosPolicyMapMatchClassMapPriority resource.
type QueuingQosPolicyMapMatchClassMapPriorityArgs struct {
	// Class map name.
	ClassMapName pulumi.StringInput
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
	// Priority level. - Range: `1`-`8`
	Level pulumi.IntInput
	// Policy map name.
	PolicyMapName pulumi.StringInput
}

func (QueuingQosPolicyMapMatchClassMapPriorityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*queuingQosPolicyMapMatchClassMapPriorityArgs)(nil)).Elem()
}

type QueuingQosPolicyMapMatchClassMapPriorityInput interface {
	pulumi.Input

	ToQueuingQosPolicyMapMatchClassMapPriorityOutput() QueuingQosPolicyMapMatchClassMapPriorityOutput
	ToQueuingQosPolicyMapMatchClassMapPriorityOutputWithContext(ctx context.Context) QueuingQosPolicyMapMatchClassMapPriorityOutput
}

func (*QueuingQosPolicyMapMatchClassMapPriority) ElementType() reflect.Type {
	return reflect.TypeOf((**QueuingQosPolicyMapMatchClassMapPriority)(nil)).Elem()
}

func (i *QueuingQosPolicyMapMatchClassMapPriority) ToQueuingQosPolicyMapMatchClassMapPriorityOutput() QueuingQosPolicyMapMatchClassMapPriorityOutput {
	return i.ToQueuingQosPolicyMapMatchClassMapPriorityOutputWithContext(context.Background())
}

func (i *QueuingQosPolicyMapMatchClassMapPriority) ToQueuingQosPolicyMapMatchClassMapPriorityOutputWithContext(ctx context.Context) QueuingQosPolicyMapMatchClassMapPriorityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueuingQosPolicyMapMatchClassMapPriorityOutput)
}

func (i *QueuingQosPolicyMapMatchClassMapPriority) ToOutput(ctx context.Context) pulumix.Output[*QueuingQosPolicyMapMatchClassMapPriority] {
	return pulumix.Output[*QueuingQosPolicyMapMatchClassMapPriority]{
		OutputState: i.ToQueuingQosPolicyMapMatchClassMapPriorityOutputWithContext(ctx).OutputState,
	}
}

// QueuingQosPolicyMapMatchClassMapPriorityArrayInput is an input type that accepts QueuingQosPolicyMapMatchClassMapPriorityArray and QueuingQosPolicyMapMatchClassMapPriorityArrayOutput values.
// You can construct a concrete instance of `QueuingQosPolicyMapMatchClassMapPriorityArrayInput` via:
//
//	QueuingQosPolicyMapMatchClassMapPriorityArray{ QueuingQosPolicyMapMatchClassMapPriorityArgs{...} }
type QueuingQosPolicyMapMatchClassMapPriorityArrayInput interface {
	pulumi.Input

	ToQueuingQosPolicyMapMatchClassMapPriorityArrayOutput() QueuingQosPolicyMapMatchClassMapPriorityArrayOutput
	ToQueuingQosPolicyMapMatchClassMapPriorityArrayOutputWithContext(context.Context) QueuingQosPolicyMapMatchClassMapPriorityArrayOutput
}

type QueuingQosPolicyMapMatchClassMapPriorityArray []QueuingQosPolicyMapMatchClassMapPriorityInput

func (QueuingQosPolicyMapMatchClassMapPriorityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*QueuingQosPolicyMapMatchClassMapPriority)(nil)).Elem()
}

func (i QueuingQosPolicyMapMatchClassMapPriorityArray) ToQueuingQosPolicyMapMatchClassMapPriorityArrayOutput() QueuingQosPolicyMapMatchClassMapPriorityArrayOutput {
	return i.ToQueuingQosPolicyMapMatchClassMapPriorityArrayOutputWithContext(context.Background())
}

func (i QueuingQosPolicyMapMatchClassMapPriorityArray) ToQueuingQosPolicyMapMatchClassMapPriorityArrayOutputWithContext(ctx context.Context) QueuingQosPolicyMapMatchClassMapPriorityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueuingQosPolicyMapMatchClassMapPriorityArrayOutput)
}

func (i QueuingQosPolicyMapMatchClassMapPriorityArray) ToOutput(ctx context.Context) pulumix.Output[[]*QueuingQosPolicyMapMatchClassMapPriority] {
	return pulumix.Output[[]*QueuingQosPolicyMapMatchClassMapPriority]{
		OutputState: i.ToQueuingQosPolicyMapMatchClassMapPriorityArrayOutputWithContext(ctx).OutputState,
	}
}

// QueuingQosPolicyMapMatchClassMapPriorityMapInput is an input type that accepts QueuingQosPolicyMapMatchClassMapPriorityMap and QueuingQosPolicyMapMatchClassMapPriorityMapOutput values.
// You can construct a concrete instance of `QueuingQosPolicyMapMatchClassMapPriorityMapInput` via:
//
//	QueuingQosPolicyMapMatchClassMapPriorityMap{ "key": QueuingQosPolicyMapMatchClassMapPriorityArgs{...} }
type QueuingQosPolicyMapMatchClassMapPriorityMapInput interface {
	pulumi.Input

	ToQueuingQosPolicyMapMatchClassMapPriorityMapOutput() QueuingQosPolicyMapMatchClassMapPriorityMapOutput
	ToQueuingQosPolicyMapMatchClassMapPriorityMapOutputWithContext(context.Context) QueuingQosPolicyMapMatchClassMapPriorityMapOutput
}

type QueuingQosPolicyMapMatchClassMapPriorityMap map[string]QueuingQosPolicyMapMatchClassMapPriorityInput

func (QueuingQosPolicyMapMatchClassMapPriorityMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*QueuingQosPolicyMapMatchClassMapPriority)(nil)).Elem()
}

func (i QueuingQosPolicyMapMatchClassMapPriorityMap) ToQueuingQosPolicyMapMatchClassMapPriorityMapOutput() QueuingQosPolicyMapMatchClassMapPriorityMapOutput {
	return i.ToQueuingQosPolicyMapMatchClassMapPriorityMapOutputWithContext(context.Background())
}

func (i QueuingQosPolicyMapMatchClassMapPriorityMap) ToQueuingQosPolicyMapMatchClassMapPriorityMapOutputWithContext(ctx context.Context) QueuingQosPolicyMapMatchClassMapPriorityMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueuingQosPolicyMapMatchClassMapPriorityMapOutput)
}

func (i QueuingQosPolicyMapMatchClassMapPriorityMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*QueuingQosPolicyMapMatchClassMapPriority] {
	return pulumix.Output[map[string]*QueuingQosPolicyMapMatchClassMapPriority]{
		OutputState: i.ToQueuingQosPolicyMapMatchClassMapPriorityMapOutputWithContext(ctx).OutputState,
	}
}

type QueuingQosPolicyMapMatchClassMapPriorityOutput struct{ *pulumi.OutputState }

func (QueuingQosPolicyMapMatchClassMapPriorityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueuingQosPolicyMapMatchClassMapPriority)(nil)).Elem()
}

func (o QueuingQosPolicyMapMatchClassMapPriorityOutput) ToQueuingQosPolicyMapMatchClassMapPriorityOutput() QueuingQosPolicyMapMatchClassMapPriorityOutput {
	return o
}

func (o QueuingQosPolicyMapMatchClassMapPriorityOutput) ToQueuingQosPolicyMapMatchClassMapPriorityOutputWithContext(ctx context.Context) QueuingQosPolicyMapMatchClassMapPriorityOutput {
	return o
}

func (o QueuingQosPolicyMapMatchClassMapPriorityOutput) ToOutput(ctx context.Context) pulumix.Output[*QueuingQosPolicyMapMatchClassMapPriority] {
	return pulumix.Output[*QueuingQosPolicyMapMatchClassMapPriority]{
		OutputState: o.OutputState,
	}
}

// Class map name.
func (o QueuingQosPolicyMapMatchClassMapPriorityOutput) ClassMapName() pulumi.StringOutput {
	return o.ApplyT(func(v *QueuingQosPolicyMapMatchClassMapPriority) pulumi.StringOutput { return v.ClassMapName }).(pulumi.StringOutput)
}

// A device name from the provider configuration.
func (o QueuingQosPolicyMapMatchClassMapPriorityOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueuingQosPolicyMapMatchClassMapPriority) pulumi.StringPtrOutput { return v.Device }).(pulumi.StringPtrOutput)
}

// Priority level. - Range: `1`-`8`
func (o QueuingQosPolicyMapMatchClassMapPriorityOutput) Level() pulumi.IntOutput {
	return o.ApplyT(func(v *QueuingQosPolicyMapMatchClassMapPriority) pulumi.IntOutput { return v.Level }).(pulumi.IntOutput)
}

// Policy map name.
func (o QueuingQosPolicyMapMatchClassMapPriorityOutput) PolicyMapName() pulumi.StringOutput {
	return o.ApplyT(func(v *QueuingQosPolicyMapMatchClassMapPriority) pulumi.StringOutput { return v.PolicyMapName }).(pulumi.StringOutput)
}

type QueuingQosPolicyMapMatchClassMapPriorityArrayOutput struct{ *pulumi.OutputState }

func (QueuingQosPolicyMapMatchClassMapPriorityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*QueuingQosPolicyMapMatchClassMapPriority)(nil)).Elem()
}

func (o QueuingQosPolicyMapMatchClassMapPriorityArrayOutput) ToQueuingQosPolicyMapMatchClassMapPriorityArrayOutput() QueuingQosPolicyMapMatchClassMapPriorityArrayOutput {
	return o
}

func (o QueuingQosPolicyMapMatchClassMapPriorityArrayOutput) ToQueuingQosPolicyMapMatchClassMapPriorityArrayOutputWithContext(ctx context.Context) QueuingQosPolicyMapMatchClassMapPriorityArrayOutput {
	return o
}

func (o QueuingQosPolicyMapMatchClassMapPriorityArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*QueuingQosPolicyMapMatchClassMapPriority] {
	return pulumix.Output[[]*QueuingQosPolicyMapMatchClassMapPriority]{
		OutputState: o.OutputState,
	}
}

func (o QueuingQosPolicyMapMatchClassMapPriorityArrayOutput) Index(i pulumi.IntInput) QueuingQosPolicyMapMatchClassMapPriorityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *QueuingQosPolicyMapMatchClassMapPriority {
		return vs[0].([]*QueuingQosPolicyMapMatchClassMapPriority)[vs[1].(int)]
	}).(QueuingQosPolicyMapMatchClassMapPriorityOutput)
}

type QueuingQosPolicyMapMatchClassMapPriorityMapOutput struct{ *pulumi.OutputState }

func (QueuingQosPolicyMapMatchClassMapPriorityMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*QueuingQosPolicyMapMatchClassMapPriority)(nil)).Elem()
}

func (o QueuingQosPolicyMapMatchClassMapPriorityMapOutput) ToQueuingQosPolicyMapMatchClassMapPriorityMapOutput() QueuingQosPolicyMapMatchClassMapPriorityMapOutput {
	return o
}

func (o QueuingQosPolicyMapMatchClassMapPriorityMapOutput) ToQueuingQosPolicyMapMatchClassMapPriorityMapOutputWithContext(ctx context.Context) QueuingQosPolicyMapMatchClassMapPriorityMapOutput {
	return o
}

func (o QueuingQosPolicyMapMatchClassMapPriorityMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*QueuingQosPolicyMapMatchClassMapPriority] {
	return pulumix.Output[map[string]*QueuingQosPolicyMapMatchClassMapPriority]{
		OutputState: o.OutputState,
	}
}

func (o QueuingQosPolicyMapMatchClassMapPriorityMapOutput) MapIndex(k pulumi.StringInput) QueuingQosPolicyMapMatchClassMapPriorityOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *QueuingQosPolicyMapMatchClassMapPriority {
		return vs[0].(map[string]*QueuingQosPolicyMapMatchClassMapPriority)[vs[1].(string)]
	}).(QueuingQosPolicyMapMatchClassMapPriorityOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*QueuingQosPolicyMapMatchClassMapPriorityInput)(nil)).Elem(), &QueuingQosPolicyMapMatchClassMapPriority{})
	pulumi.RegisterInputType(reflect.TypeOf((*QueuingQosPolicyMapMatchClassMapPriorityArrayInput)(nil)).Elem(), QueuingQosPolicyMapMatchClassMapPriorityArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*QueuingQosPolicyMapMatchClassMapPriorityMapInput)(nil)).Elem(), QueuingQosPolicyMapMatchClassMapPriorityMap{})
	pulumi.RegisterOutputType(QueuingQosPolicyMapMatchClassMapPriorityOutput{})
	pulumi.RegisterOutputType(QueuingQosPolicyMapMatchClassMapPriorityArrayOutput{})
	pulumi.RegisterOutputType(QueuingQosPolicyMapMatchClassMapPriorityMapOutput{})
}
