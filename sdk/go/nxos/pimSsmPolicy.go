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

// This resource can manage the PIM SSM policy configuration.
//
// - API Documentation: [pimSSMPatP](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/pim:SSMPatP/)
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
//			_, err := nxos.NewPimSsmPolicy(ctx, "example", &nxos.PimSsmPolicyArgs{
//				VrfName: pulumi.String("default"),
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
//	$ pulumi import nxos:index/pimSsmPolicy:PimSsmPolicy example "sys/pim/inst/dom-[default]/ssm"
//
// ```
type PimSsmPolicy struct {
	pulumi.CustomResourceState

	// A device name from the provider configuration.
	Device pulumi.StringPtrOutput `pulumi:"device"`
	// Policy name.
	Name pulumi.StringOutput `pulumi:"name"`
	// VRF name.
	VrfName pulumi.StringOutput `pulumi:"vrfName"`
}

// NewPimSsmPolicy registers a new resource with the given unique name, arguments, and options.
func NewPimSsmPolicy(ctx *pulumi.Context,
	name string, args *PimSsmPolicyArgs, opts ...pulumi.ResourceOption) (*PimSsmPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.VrfName == nil {
		return nil, errors.New("invalid value for required argument 'VrfName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PimSsmPolicy
	err := ctx.RegisterResource("nxos:index/pimSsmPolicy:PimSsmPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPimSsmPolicy gets an existing PimSsmPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPimSsmPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PimSsmPolicyState, opts ...pulumi.ResourceOption) (*PimSsmPolicy, error) {
	var resource PimSsmPolicy
	err := ctx.ReadResource("nxos:index/pimSsmPolicy:PimSsmPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PimSsmPolicy resources.
type pimSsmPolicyState struct {
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// Policy name.
	Name *string `pulumi:"name"`
	// VRF name.
	VrfName *string `pulumi:"vrfName"`
}

type PimSsmPolicyState struct {
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
	// Policy name.
	Name pulumi.StringPtrInput
	// VRF name.
	VrfName pulumi.StringPtrInput
}

func (PimSsmPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*pimSsmPolicyState)(nil)).Elem()
}

type pimSsmPolicyArgs struct {
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// Policy name.
	Name *string `pulumi:"name"`
	// VRF name.
	VrfName string `pulumi:"vrfName"`
}

// The set of arguments for constructing a PimSsmPolicy resource.
type PimSsmPolicyArgs struct {
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
	// Policy name.
	Name pulumi.StringPtrInput
	// VRF name.
	VrfName pulumi.StringInput
}

func (PimSsmPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pimSsmPolicyArgs)(nil)).Elem()
}

type PimSsmPolicyInput interface {
	pulumi.Input

	ToPimSsmPolicyOutput() PimSsmPolicyOutput
	ToPimSsmPolicyOutputWithContext(ctx context.Context) PimSsmPolicyOutput
}

func (*PimSsmPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**PimSsmPolicy)(nil)).Elem()
}

func (i *PimSsmPolicy) ToPimSsmPolicyOutput() PimSsmPolicyOutput {
	return i.ToPimSsmPolicyOutputWithContext(context.Background())
}

func (i *PimSsmPolicy) ToPimSsmPolicyOutputWithContext(ctx context.Context) PimSsmPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PimSsmPolicyOutput)
}

func (i *PimSsmPolicy) ToOutput(ctx context.Context) pulumix.Output[*PimSsmPolicy] {
	return pulumix.Output[*PimSsmPolicy]{
		OutputState: i.ToPimSsmPolicyOutputWithContext(ctx).OutputState,
	}
}

// PimSsmPolicyArrayInput is an input type that accepts PimSsmPolicyArray and PimSsmPolicyArrayOutput values.
// You can construct a concrete instance of `PimSsmPolicyArrayInput` via:
//
//	PimSsmPolicyArray{ PimSsmPolicyArgs{...} }
type PimSsmPolicyArrayInput interface {
	pulumi.Input

	ToPimSsmPolicyArrayOutput() PimSsmPolicyArrayOutput
	ToPimSsmPolicyArrayOutputWithContext(context.Context) PimSsmPolicyArrayOutput
}

type PimSsmPolicyArray []PimSsmPolicyInput

func (PimSsmPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PimSsmPolicy)(nil)).Elem()
}

func (i PimSsmPolicyArray) ToPimSsmPolicyArrayOutput() PimSsmPolicyArrayOutput {
	return i.ToPimSsmPolicyArrayOutputWithContext(context.Background())
}

func (i PimSsmPolicyArray) ToPimSsmPolicyArrayOutputWithContext(ctx context.Context) PimSsmPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PimSsmPolicyArrayOutput)
}

func (i PimSsmPolicyArray) ToOutput(ctx context.Context) pulumix.Output[[]*PimSsmPolicy] {
	return pulumix.Output[[]*PimSsmPolicy]{
		OutputState: i.ToPimSsmPolicyArrayOutputWithContext(ctx).OutputState,
	}
}

// PimSsmPolicyMapInput is an input type that accepts PimSsmPolicyMap and PimSsmPolicyMapOutput values.
// You can construct a concrete instance of `PimSsmPolicyMapInput` via:
//
//	PimSsmPolicyMap{ "key": PimSsmPolicyArgs{...} }
type PimSsmPolicyMapInput interface {
	pulumi.Input

	ToPimSsmPolicyMapOutput() PimSsmPolicyMapOutput
	ToPimSsmPolicyMapOutputWithContext(context.Context) PimSsmPolicyMapOutput
}

type PimSsmPolicyMap map[string]PimSsmPolicyInput

func (PimSsmPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PimSsmPolicy)(nil)).Elem()
}

func (i PimSsmPolicyMap) ToPimSsmPolicyMapOutput() PimSsmPolicyMapOutput {
	return i.ToPimSsmPolicyMapOutputWithContext(context.Background())
}

func (i PimSsmPolicyMap) ToPimSsmPolicyMapOutputWithContext(ctx context.Context) PimSsmPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PimSsmPolicyMapOutput)
}

func (i PimSsmPolicyMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*PimSsmPolicy] {
	return pulumix.Output[map[string]*PimSsmPolicy]{
		OutputState: i.ToPimSsmPolicyMapOutputWithContext(ctx).OutputState,
	}
}

type PimSsmPolicyOutput struct{ *pulumi.OutputState }

func (PimSsmPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PimSsmPolicy)(nil)).Elem()
}

func (o PimSsmPolicyOutput) ToPimSsmPolicyOutput() PimSsmPolicyOutput {
	return o
}

func (o PimSsmPolicyOutput) ToPimSsmPolicyOutputWithContext(ctx context.Context) PimSsmPolicyOutput {
	return o
}

func (o PimSsmPolicyOutput) ToOutput(ctx context.Context) pulumix.Output[*PimSsmPolicy] {
	return pulumix.Output[*PimSsmPolicy]{
		OutputState: o.OutputState,
	}
}

// A device name from the provider configuration.
func (o PimSsmPolicyOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PimSsmPolicy) pulumi.StringPtrOutput { return v.Device }).(pulumi.StringPtrOutput)
}

// Policy name.
func (o PimSsmPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PimSsmPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// VRF name.
func (o PimSsmPolicyOutput) VrfName() pulumi.StringOutput {
	return o.ApplyT(func(v *PimSsmPolicy) pulumi.StringOutput { return v.VrfName }).(pulumi.StringOutput)
}

type PimSsmPolicyArrayOutput struct{ *pulumi.OutputState }

func (PimSsmPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PimSsmPolicy)(nil)).Elem()
}

func (o PimSsmPolicyArrayOutput) ToPimSsmPolicyArrayOutput() PimSsmPolicyArrayOutput {
	return o
}

func (o PimSsmPolicyArrayOutput) ToPimSsmPolicyArrayOutputWithContext(ctx context.Context) PimSsmPolicyArrayOutput {
	return o
}

func (o PimSsmPolicyArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*PimSsmPolicy] {
	return pulumix.Output[[]*PimSsmPolicy]{
		OutputState: o.OutputState,
	}
}

func (o PimSsmPolicyArrayOutput) Index(i pulumi.IntInput) PimSsmPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PimSsmPolicy {
		return vs[0].([]*PimSsmPolicy)[vs[1].(int)]
	}).(PimSsmPolicyOutput)
}

type PimSsmPolicyMapOutput struct{ *pulumi.OutputState }

func (PimSsmPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PimSsmPolicy)(nil)).Elem()
}

func (o PimSsmPolicyMapOutput) ToPimSsmPolicyMapOutput() PimSsmPolicyMapOutput {
	return o
}

func (o PimSsmPolicyMapOutput) ToPimSsmPolicyMapOutputWithContext(ctx context.Context) PimSsmPolicyMapOutput {
	return o
}

func (o PimSsmPolicyMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*PimSsmPolicy] {
	return pulumix.Output[map[string]*PimSsmPolicy]{
		OutputState: o.OutputState,
	}
}

func (o PimSsmPolicyMapOutput) MapIndex(k pulumi.StringInput) PimSsmPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PimSsmPolicy {
		return vs[0].(map[string]*PimSsmPolicy)[vs[1].(string)]
	}).(PimSsmPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PimSsmPolicyInput)(nil)).Elem(), &PimSsmPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*PimSsmPolicyArrayInput)(nil)).Elem(), PimSsmPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PimSsmPolicyMapInput)(nil)).Elem(), PimSsmPolicyMap{})
	pulumi.RegisterOutputType(PimSsmPolicyOutput{})
	pulumi.RegisterOutputType(PimSsmPolicyArrayOutput{})
	pulumi.RegisterOutputType(PimSsmPolicyMapOutput{})
}
