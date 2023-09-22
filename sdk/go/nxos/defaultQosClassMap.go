// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nxos

import (
	"context"
	"reflect"

	"github.com/lbrlabs/pulumi-nxos/sdk/go/nxos/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// This resource can manage the default QoS class map configuration.
//
// - API Documentation: [ipqosCMapInst](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Qos/ipqos:CMapInst/)
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
//			_, err := nxos.NewDefaultQosClassMap(ctx, "example", &nxos.DefaultQosClassMapArgs{
//				MatchType: pulumi.String("match-any"),
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
//	$ pulumi import nxos:index/defaultQosClassMap:DefaultQosClassMap example "sys/ipqos/dflt/c/name-[Voice]"
//
// ```
type DefaultQosClassMap struct {
	pulumi.CustomResourceState

	// A device name from the provider configuration.
	Device pulumi.StringPtrOutput `pulumi:"device"`
	// Match type. - Choices: `match-any`, `match-all`, `match-first` - Default value: `match-all`
	MatchType pulumi.StringOutput `pulumi:"matchType"`
	// Class map name.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewDefaultQosClassMap registers a new resource with the given unique name, arguments, and options.
func NewDefaultQosClassMap(ctx *pulumi.Context,
	name string, args *DefaultQosClassMapArgs, opts ...pulumi.ResourceOption) (*DefaultQosClassMap, error) {
	if args == nil {
		args = &DefaultQosClassMapArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DefaultQosClassMap
	err := ctx.RegisterResource("nxos:index/defaultQosClassMap:DefaultQosClassMap", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDefaultQosClassMap gets an existing DefaultQosClassMap resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDefaultQosClassMap(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DefaultQosClassMapState, opts ...pulumi.ResourceOption) (*DefaultQosClassMap, error) {
	var resource DefaultQosClassMap
	err := ctx.ReadResource("nxos:index/defaultQosClassMap:DefaultQosClassMap", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DefaultQosClassMap resources.
type defaultQosClassMapState struct {
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// Match type. - Choices: `match-any`, `match-all`, `match-first` - Default value: `match-all`
	MatchType *string `pulumi:"matchType"`
	// Class map name.
	Name *string `pulumi:"name"`
}

type DefaultQosClassMapState struct {
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
	// Match type. - Choices: `match-any`, `match-all`, `match-first` - Default value: `match-all`
	MatchType pulumi.StringPtrInput
	// Class map name.
	Name pulumi.StringPtrInput
}

func (DefaultQosClassMapState) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultQosClassMapState)(nil)).Elem()
}

type defaultQosClassMapArgs struct {
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// Match type. - Choices: `match-any`, `match-all`, `match-first` - Default value: `match-all`
	MatchType *string `pulumi:"matchType"`
	// Class map name.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a DefaultQosClassMap resource.
type DefaultQosClassMapArgs struct {
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
	// Match type. - Choices: `match-any`, `match-all`, `match-first` - Default value: `match-all`
	MatchType pulumi.StringPtrInput
	// Class map name.
	Name pulumi.StringPtrInput
}

func (DefaultQosClassMapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultQosClassMapArgs)(nil)).Elem()
}

type DefaultQosClassMapInput interface {
	pulumi.Input

	ToDefaultQosClassMapOutput() DefaultQosClassMapOutput
	ToDefaultQosClassMapOutputWithContext(ctx context.Context) DefaultQosClassMapOutput
}

func (*DefaultQosClassMap) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultQosClassMap)(nil)).Elem()
}

func (i *DefaultQosClassMap) ToDefaultQosClassMapOutput() DefaultQosClassMapOutput {
	return i.ToDefaultQosClassMapOutputWithContext(context.Background())
}

func (i *DefaultQosClassMap) ToDefaultQosClassMapOutputWithContext(ctx context.Context) DefaultQosClassMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultQosClassMapOutput)
}

func (i *DefaultQosClassMap) ToOutput(ctx context.Context) pulumix.Output[*DefaultQosClassMap] {
	return pulumix.Output[*DefaultQosClassMap]{
		OutputState: i.ToDefaultQosClassMapOutputWithContext(ctx).OutputState,
	}
}

// DefaultQosClassMapArrayInput is an input type that accepts DefaultQosClassMapArray and DefaultQosClassMapArrayOutput values.
// You can construct a concrete instance of `DefaultQosClassMapArrayInput` via:
//
//	DefaultQosClassMapArray{ DefaultQosClassMapArgs{...} }
type DefaultQosClassMapArrayInput interface {
	pulumi.Input

	ToDefaultQosClassMapArrayOutput() DefaultQosClassMapArrayOutput
	ToDefaultQosClassMapArrayOutputWithContext(context.Context) DefaultQosClassMapArrayOutput
}

type DefaultQosClassMapArray []DefaultQosClassMapInput

func (DefaultQosClassMapArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DefaultQosClassMap)(nil)).Elem()
}

func (i DefaultQosClassMapArray) ToDefaultQosClassMapArrayOutput() DefaultQosClassMapArrayOutput {
	return i.ToDefaultQosClassMapArrayOutputWithContext(context.Background())
}

func (i DefaultQosClassMapArray) ToDefaultQosClassMapArrayOutputWithContext(ctx context.Context) DefaultQosClassMapArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultQosClassMapArrayOutput)
}

func (i DefaultQosClassMapArray) ToOutput(ctx context.Context) pulumix.Output[[]*DefaultQosClassMap] {
	return pulumix.Output[[]*DefaultQosClassMap]{
		OutputState: i.ToDefaultQosClassMapArrayOutputWithContext(ctx).OutputState,
	}
}

// DefaultQosClassMapMapInput is an input type that accepts DefaultQosClassMapMap and DefaultQosClassMapMapOutput values.
// You can construct a concrete instance of `DefaultQosClassMapMapInput` via:
//
//	DefaultQosClassMapMap{ "key": DefaultQosClassMapArgs{...} }
type DefaultQosClassMapMapInput interface {
	pulumi.Input

	ToDefaultQosClassMapMapOutput() DefaultQosClassMapMapOutput
	ToDefaultQosClassMapMapOutputWithContext(context.Context) DefaultQosClassMapMapOutput
}

type DefaultQosClassMapMap map[string]DefaultQosClassMapInput

func (DefaultQosClassMapMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DefaultQosClassMap)(nil)).Elem()
}

func (i DefaultQosClassMapMap) ToDefaultQosClassMapMapOutput() DefaultQosClassMapMapOutput {
	return i.ToDefaultQosClassMapMapOutputWithContext(context.Background())
}

func (i DefaultQosClassMapMap) ToDefaultQosClassMapMapOutputWithContext(ctx context.Context) DefaultQosClassMapMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultQosClassMapMapOutput)
}

func (i DefaultQosClassMapMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*DefaultQosClassMap] {
	return pulumix.Output[map[string]*DefaultQosClassMap]{
		OutputState: i.ToDefaultQosClassMapMapOutputWithContext(ctx).OutputState,
	}
}

type DefaultQosClassMapOutput struct{ *pulumi.OutputState }

func (DefaultQosClassMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultQosClassMap)(nil)).Elem()
}

func (o DefaultQosClassMapOutput) ToDefaultQosClassMapOutput() DefaultQosClassMapOutput {
	return o
}

func (o DefaultQosClassMapOutput) ToDefaultQosClassMapOutputWithContext(ctx context.Context) DefaultQosClassMapOutput {
	return o
}

func (o DefaultQosClassMapOutput) ToOutput(ctx context.Context) pulumix.Output[*DefaultQosClassMap] {
	return pulumix.Output[*DefaultQosClassMap]{
		OutputState: o.OutputState,
	}
}

// A device name from the provider configuration.
func (o DefaultQosClassMapOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefaultQosClassMap) pulumi.StringPtrOutput { return v.Device }).(pulumi.StringPtrOutput)
}

// Match type. - Choices: `match-any`, `match-all`, `match-first` - Default value: `match-all`
func (o DefaultQosClassMapOutput) MatchType() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultQosClassMap) pulumi.StringOutput { return v.MatchType }).(pulumi.StringOutput)
}

// Class map name.
func (o DefaultQosClassMapOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultQosClassMap) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type DefaultQosClassMapArrayOutput struct{ *pulumi.OutputState }

func (DefaultQosClassMapArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DefaultQosClassMap)(nil)).Elem()
}

func (o DefaultQosClassMapArrayOutput) ToDefaultQosClassMapArrayOutput() DefaultQosClassMapArrayOutput {
	return o
}

func (o DefaultQosClassMapArrayOutput) ToDefaultQosClassMapArrayOutputWithContext(ctx context.Context) DefaultQosClassMapArrayOutput {
	return o
}

func (o DefaultQosClassMapArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*DefaultQosClassMap] {
	return pulumix.Output[[]*DefaultQosClassMap]{
		OutputState: o.OutputState,
	}
}

func (o DefaultQosClassMapArrayOutput) Index(i pulumi.IntInput) DefaultQosClassMapOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DefaultQosClassMap {
		return vs[0].([]*DefaultQosClassMap)[vs[1].(int)]
	}).(DefaultQosClassMapOutput)
}

type DefaultQosClassMapMapOutput struct{ *pulumi.OutputState }

func (DefaultQosClassMapMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DefaultQosClassMap)(nil)).Elem()
}

func (o DefaultQosClassMapMapOutput) ToDefaultQosClassMapMapOutput() DefaultQosClassMapMapOutput {
	return o
}

func (o DefaultQosClassMapMapOutput) ToDefaultQosClassMapMapOutputWithContext(ctx context.Context) DefaultQosClassMapMapOutput {
	return o
}

func (o DefaultQosClassMapMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*DefaultQosClassMap] {
	return pulumix.Output[map[string]*DefaultQosClassMap]{
		OutputState: o.OutputState,
	}
}

func (o DefaultQosClassMapMapOutput) MapIndex(k pulumi.StringInput) DefaultQosClassMapOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DefaultQosClassMap {
		return vs[0].(map[string]*DefaultQosClassMap)[vs[1].(string)]
	}).(DefaultQosClassMapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultQosClassMapInput)(nil)).Elem(), &DefaultQosClassMap{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultQosClassMapArrayInput)(nil)).Elem(), DefaultQosClassMapArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultQosClassMapMapInput)(nil)).Elem(), DefaultQosClassMapMap{})
	pulumi.RegisterOutputType(DefaultQosClassMapOutput{})
	pulumi.RegisterOutputType(DefaultQosClassMapArrayOutput{})
	pulumi.RegisterOutputType(DefaultQosClassMapMapOutput{})
}