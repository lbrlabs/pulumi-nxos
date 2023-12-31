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

// This resource can manage the queuing QoS policy system out policy map configuration.
//
// - API Documentation: [ipqosInst](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Qos/ipqos:Inst/)
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
//			_, err := nxos.NewQueuingQosPolicySystemOutPolicyMap(ctx, "example", &nxos.QueuingQosPolicySystemOutPolicyMapArgs{
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
//	$ pulumi import nxos:index/queuingQosPolicySystemOutPolicyMap:QueuingQosPolicySystemOutPolicyMap example "sys/ipqos/queuing/policy/out/sys/pmap"
//
// ```
type QueuingQosPolicySystemOutPolicyMap struct {
	pulumi.CustomResourceState

	// A device name from the provider configuration.
	Device pulumi.StringPtrOutput `pulumi:"device"`
	// Policy map name.
	PolicyMapName pulumi.StringOutput `pulumi:"policyMapName"`
}

// NewQueuingQosPolicySystemOutPolicyMap registers a new resource with the given unique name, arguments, and options.
func NewQueuingQosPolicySystemOutPolicyMap(ctx *pulumi.Context,
	name string, args *QueuingQosPolicySystemOutPolicyMapArgs, opts ...pulumi.ResourceOption) (*QueuingQosPolicySystemOutPolicyMap, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyMapName == nil {
		return nil, errors.New("invalid value for required argument 'PolicyMapName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource QueuingQosPolicySystemOutPolicyMap
	err := ctx.RegisterResource("nxos:index/queuingQosPolicySystemOutPolicyMap:QueuingQosPolicySystemOutPolicyMap", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQueuingQosPolicySystemOutPolicyMap gets an existing QueuingQosPolicySystemOutPolicyMap resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQueuingQosPolicySystemOutPolicyMap(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QueuingQosPolicySystemOutPolicyMapState, opts ...pulumi.ResourceOption) (*QueuingQosPolicySystemOutPolicyMap, error) {
	var resource QueuingQosPolicySystemOutPolicyMap
	err := ctx.ReadResource("nxos:index/queuingQosPolicySystemOutPolicyMap:QueuingQosPolicySystemOutPolicyMap", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering QueuingQosPolicySystemOutPolicyMap resources.
type queuingQosPolicySystemOutPolicyMapState struct {
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// Policy map name.
	PolicyMapName *string `pulumi:"policyMapName"`
}

type QueuingQosPolicySystemOutPolicyMapState struct {
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
	// Policy map name.
	PolicyMapName pulumi.StringPtrInput
}

func (QueuingQosPolicySystemOutPolicyMapState) ElementType() reflect.Type {
	return reflect.TypeOf((*queuingQosPolicySystemOutPolicyMapState)(nil)).Elem()
}

type queuingQosPolicySystemOutPolicyMapArgs struct {
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// Policy map name.
	PolicyMapName string `pulumi:"policyMapName"`
}

// The set of arguments for constructing a QueuingQosPolicySystemOutPolicyMap resource.
type QueuingQosPolicySystemOutPolicyMapArgs struct {
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
	// Policy map name.
	PolicyMapName pulumi.StringInput
}

func (QueuingQosPolicySystemOutPolicyMapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*queuingQosPolicySystemOutPolicyMapArgs)(nil)).Elem()
}

type QueuingQosPolicySystemOutPolicyMapInput interface {
	pulumi.Input

	ToQueuingQosPolicySystemOutPolicyMapOutput() QueuingQosPolicySystemOutPolicyMapOutput
	ToQueuingQosPolicySystemOutPolicyMapOutputWithContext(ctx context.Context) QueuingQosPolicySystemOutPolicyMapOutput
}

func (*QueuingQosPolicySystemOutPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((**QueuingQosPolicySystemOutPolicyMap)(nil)).Elem()
}

func (i *QueuingQosPolicySystemOutPolicyMap) ToQueuingQosPolicySystemOutPolicyMapOutput() QueuingQosPolicySystemOutPolicyMapOutput {
	return i.ToQueuingQosPolicySystemOutPolicyMapOutputWithContext(context.Background())
}

func (i *QueuingQosPolicySystemOutPolicyMap) ToQueuingQosPolicySystemOutPolicyMapOutputWithContext(ctx context.Context) QueuingQosPolicySystemOutPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueuingQosPolicySystemOutPolicyMapOutput)
}

func (i *QueuingQosPolicySystemOutPolicyMap) ToOutput(ctx context.Context) pulumix.Output[*QueuingQosPolicySystemOutPolicyMap] {
	return pulumix.Output[*QueuingQosPolicySystemOutPolicyMap]{
		OutputState: i.ToQueuingQosPolicySystemOutPolicyMapOutputWithContext(ctx).OutputState,
	}
}

// QueuingQosPolicySystemOutPolicyMapArrayInput is an input type that accepts QueuingQosPolicySystemOutPolicyMapArray and QueuingQosPolicySystemOutPolicyMapArrayOutput values.
// You can construct a concrete instance of `QueuingQosPolicySystemOutPolicyMapArrayInput` via:
//
//	QueuingQosPolicySystemOutPolicyMapArray{ QueuingQosPolicySystemOutPolicyMapArgs{...} }
type QueuingQosPolicySystemOutPolicyMapArrayInput interface {
	pulumi.Input

	ToQueuingQosPolicySystemOutPolicyMapArrayOutput() QueuingQosPolicySystemOutPolicyMapArrayOutput
	ToQueuingQosPolicySystemOutPolicyMapArrayOutputWithContext(context.Context) QueuingQosPolicySystemOutPolicyMapArrayOutput
}

type QueuingQosPolicySystemOutPolicyMapArray []QueuingQosPolicySystemOutPolicyMapInput

func (QueuingQosPolicySystemOutPolicyMapArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*QueuingQosPolicySystemOutPolicyMap)(nil)).Elem()
}

func (i QueuingQosPolicySystemOutPolicyMapArray) ToQueuingQosPolicySystemOutPolicyMapArrayOutput() QueuingQosPolicySystemOutPolicyMapArrayOutput {
	return i.ToQueuingQosPolicySystemOutPolicyMapArrayOutputWithContext(context.Background())
}

func (i QueuingQosPolicySystemOutPolicyMapArray) ToQueuingQosPolicySystemOutPolicyMapArrayOutputWithContext(ctx context.Context) QueuingQosPolicySystemOutPolicyMapArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueuingQosPolicySystemOutPolicyMapArrayOutput)
}

func (i QueuingQosPolicySystemOutPolicyMapArray) ToOutput(ctx context.Context) pulumix.Output[[]*QueuingQosPolicySystemOutPolicyMap] {
	return pulumix.Output[[]*QueuingQosPolicySystemOutPolicyMap]{
		OutputState: i.ToQueuingQosPolicySystemOutPolicyMapArrayOutputWithContext(ctx).OutputState,
	}
}

// QueuingQosPolicySystemOutPolicyMapMapInput is an input type that accepts QueuingQosPolicySystemOutPolicyMapMap and QueuingQosPolicySystemOutPolicyMapMapOutput values.
// You can construct a concrete instance of `QueuingQosPolicySystemOutPolicyMapMapInput` via:
//
//	QueuingQosPolicySystemOutPolicyMapMap{ "key": QueuingQosPolicySystemOutPolicyMapArgs{...} }
type QueuingQosPolicySystemOutPolicyMapMapInput interface {
	pulumi.Input

	ToQueuingQosPolicySystemOutPolicyMapMapOutput() QueuingQosPolicySystemOutPolicyMapMapOutput
	ToQueuingQosPolicySystemOutPolicyMapMapOutputWithContext(context.Context) QueuingQosPolicySystemOutPolicyMapMapOutput
}

type QueuingQosPolicySystemOutPolicyMapMap map[string]QueuingQosPolicySystemOutPolicyMapInput

func (QueuingQosPolicySystemOutPolicyMapMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*QueuingQosPolicySystemOutPolicyMap)(nil)).Elem()
}

func (i QueuingQosPolicySystemOutPolicyMapMap) ToQueuingQosPolicySystemOutPolicyMapMapOutput() QueuingQosPolicySystemOutPolicyMapMapOutput {
	return i.ToQueuingQosPolicySystemOutPolicyMapMapOutputWithContext(context.Background())
}

func (i QueuingQosPolicySystemOutPolicyMapMap) ToQueuingQosPolicySystemOutPolicyMapMapOutputWithContext(ctx context.Context) QueuingQosPolicySystemOutPolicyMapMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueuingQosPolicySystemOutPolicyMapMapOutput)
}

func (i QueuingQosPolicySystemOutPolicyMapMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*QueuingQosPolicySystemOutPolicyMap] {
	return pulumix.Output[map[string]*QueuingQosPolicySystemOutPolicyMap]{
		OutputState: i.ToQueuingQosPolicySystemOutPolicyMapMapOutputWithContext(ctx).OutputState,
	}
}

type QueuingQosPolicySystemOutPolicyMapOutput struct{ *pulumi.OutputState }

func (QueuingQosPolicySystemOutPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueuingQosPolicySystemOutPolicyMap)(nil)).Elem()
}

func (o QueuingQosPolicySystemOutPolicyMapOutput) ToQueuingQosPolicySystemOutPolicyMapOutput() QueuingQosPolicySystemOutPolicyMapOutput {
	return o
}

func (o QueuingQosPolicySystemOutPolicyMapOutput) ToQueuingQosPolicySystemOutPolicyMapOutputWithContext(ctx context.Context) QueuingQosPolicySystemOutPolicyMapOutput {
	return o
}

func (o QueuingQosPolicySystemOutPolicyMapOutput) ToOutput(ctx context.Context) pulumix.Output[*QueuingQosPolicySystemOutPolicyMap] {
	return pulumix.Output[*QueuingQosPolicySystemOutPolicyMap]{
		OutputState: o.OutputState,
	}
}

// A device name from the provider configuration.
func (o QueuingQosPolicySystemOutPolicyMapOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueuingQosPolicySystemOutPolicyMap) pulumi.StringPtrOutput { return v.Device }).(pulumi.StringPtrOutput)
}

// Policy map name.
func (o QueuingQosPolicySystemOutPolicyMapOutput) PolicyMapName() pulumi.StringOutput {
	return o.ApplyT(func(v *QueuingQosPolicySystemOutPolicyMap) pulumi.StringOutput { return v.PolicyMapName }).(pulumi.StringOutput)
}

type QueuingQosPolicySystemOutPolicyMapArrayOutput struct{ *pulumi.OutputState }

func (QueuingQosPolicySystemOutPolicyMapArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*QueuingQosPolicySystemOutPolicyMap)(nil)).Elem()
}

func (o QueuingQosPolicySystemOutPolicyMapArrayOutput) ToQueuingQosPolicySystemOutPolicyMapArrayOutput() QueuingQosPolicySystemOutPolicyMapArrayOutput {
	return o
}

func (o QueuingQosPolicySystemOutPolicyMapArrayOutput) ToQueuingQosPolicySystemOutPolicyMapArrayOutputWithContext(ctx context.Context) QueuingQosPolicySystemOutPolicyMapArrayOutput {
	return o
}

func (o QueuingQosPolicySystemOutPolicyMapArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*QueuingQosPolicySystemOutPolicyMap] {
	return pulumix.Output[[]*QueuingQosPolicySystemOutPolicyMap]{
		OutputState: o.OutputState,
	}
}

func (o QueuingQosPolicySystemOutPolicyMapArrayOutput) Index(i pulumi.IntInput) QueuingQosPolicySystemOutPolicyMapOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *QueuingQosPolicySystemOutPolicyMap {
		return vs[0].([]*QueuingQosPolicySystemOutPolicyMap)[vs[1].(int)]
	}).(QueuingQosPolicySystemOutPolicyMapOutput)
}

type QueuingQosPolicySystemOutPolicyMapMapOutput struct{ *pulumi.OutputState }

func (QueuingQosPolicySystemOutPolicyMapMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*QueuingQosPolicySystemOutPolicyMap)(nil)).Elem()
}

func (o QueuingQosPolicySystemOutPolicyMapMapOutput) ToQueuingQosPolicySystemOutPolicyMapMapOutput() QueuingQosPolicySystemOutPolicyMapMapOutput {
	return o
}

func (o QueuingQosPolicySystemOutPolicyMapMapOutput) ToQueuingQosPolicySystemOutPolicyMapMapOutputWithContext(ctx context.Context) QueuingQosPolicySystemOutPolicyMapMapOutput {
	return o
}

func (o QueuingQosPolicySystemOutPolicyMapMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*QueuingQosPolicySystemOutPolicyMap] {
	return pulumix.Output[map[string]*QueuingQosPolicySystemOutPolicyMap]{
		OutputState: o.OutputState,
	}
}

func (o QueuingQosPolicySystemOutPolicyMapMapOutput) MapIndex(k pulumi.StringInput) QueuingQosPolicySystemOutPolicyMapOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *QueuingQosPolicySystemOutPolicyMap {
		return vs[0].(map[string]*QueuingQosPolicySystemOutPolicyMap)[vs[1].(string)]
	}).(QueuingQosPolicySystemOutPolicyMapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*QueuingQosPolicySystemOutPolicyMapInput)(nil)).Elem(), &QueuingQosPolicySystemOutPolicyMap{})
	pulumi.RegisterInputType(reflect.TypeOf((*QueuingQosPolicySystemOutPolicyMapArrayInput)(nil)).Elem(), QueuingQosPolicySystemOutPolicyMapArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*QueuingQosPolicySystemOutPolicyMapMapInput)(nil)).Elem(), QueuingQosPolicySystemOutPolicyMapMap{})
	pulumi.RegisterOutputType(QueuingQosPolicySystemOutPolicyMapOutput{})
	pulumi.RegisterOutputType(QueuingQosPolicySystemOutPolicyMapArrayOutput{})
	pulumi.RegisterOutputType(QueuingQosPolicySystemOutPolicyMapMapOutput{})
}
