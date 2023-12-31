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

// This resource can manage the BGP domain (VRF) configuration.
//
// - API Documentation: [bgpDom](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/bgp:Dom/)
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
//			_, err := nxos.NewBgpVrf(ctx, "example", &nxos.BgpVrfArgs{
//				Asn:      pulumi.String("65001"),
//				RouterId: pulumi.String("1.1.1.1"),
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
//	$ pulumi import nxos:index/bgpVrf:BgpVrf example "sys/bgp/inst/dom-[default]"
//
// ```
type BgpVrf struct {
	pulumi.CustomResourceState

	// Autonomous system number.
	Asn pulumi.StringOutput `pulumi:"asn"`
	// A device name from the provider configuration.
	Device pulumi.StringPtrOutput `pulumi:"device"`
	// VRF name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Router ID.
	RouterId pulumi.StringPtrOutput `pulumi:"routerId"`
}

// NewBgpVrf registers a new resource with the given unique name, arguments, and options.
func NewBgpVrf(ctx *pulumi.Context,
	name string, args *BgpVrfArgs, opts ...pulumi.ResourceOption) (*BgpVrf, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Asn == nil {
		return nil, errors.New("invalid value for required argument 'Asn'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BgpVrf
	err := ctx.RegisterResource("nxos:index/bgpVrf:BgpVrf", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBgpVrf gets an existing BgpVrf resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBgpVrf(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BgpVrfState, opts ...pulumi.ResourceOption) (*BgpVrf, error) {
	var resource BgpVrf
	err := ctx.ReadResource("nxos:index/bgpVrf:BgpVrf", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BgpVrf resources.
type bgpVrfState struct {
	// Autonomous system number.
	Asn *string `pulumi:"asn"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// VRF name.
	Name *string `pulumi:"name"`
	// Router ID.
	RouterId *string `pulumi:"routerId"`
}

type BgpVrfState struct {
	// Autonomous system number.
	Asn pulumi.StringPtrInput
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
	// VRF name.
	Name pulumi.StringPtrInput
	// Router ID.
	RouterId pulumi.StringPtrInput
}

func (BgpVrfState) ElementType() reflect.Type {
	return reflect.TypeOf((*bgpVrfState)(nil)).Elem()
}

type bgpVrfArgs struct {
	// Autonomous system number.
	Asn string `pulumi:"asn"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// VRF name.
	Name *string `pulumi:"name"`
	// Router ID.
	RouterId *string `pulumi:"routerId"`
}

// The set of arguments for constructing a BgpVrf resource.
type BgpVrfArgs struct {
	// Autonomous system number.
	Asn pulumi.StringInput
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
	// VRF name.
	Name pulumi.StringPtrInput
	// Router ID.
	RouterId pulumi.StringPtrInput
}

func (BgpVrfArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bgpVrfArgs)(nil)).Elem()
}

type BgpVrfInput interface {
	pulumi.Input

	ToBgpVrfOutput() BgpVrfOutput
	ToBgpVrfOutputWithContext(ctx context.Context) BgpVrfOutput
}

func (*BgpVrf) ElementType() reflect.Type {
	return reflect.TypeOf((**BgpVrf)(nil)).Elem()
}

func (i *BgpVrf) ToBgpVrfOutput() BgpVrfOutput {
	return i.ToBgpVrfOutputWithContext(context.Background())
}

func (i *BgpVrf) ToBgpVrfOutputWithContext(ctx context.Context) BgpVrfOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BgpVrfOutput)
}

func (i *BgpVrf) ToOutput(ctx context.Context) pulumix.Output[*BgpVrf] {
	return pulumix.Output[*BgpVrf]{
		OutputState: i.ToBgpVrfOutputWithContext(ctx).OutputState,
	}
}

// BgpVrfArrayInput is an input type that accepts BgpVrfArray and BgpVrfArrayOutput values.
// You can construct a concrete instance of `BgpVrfArrayInput` via:
//
//	BgpVrfArray{ BgpVrfArgs{...} }
type BgpVrfArrayInput interface {
	pulumi.Input

	ToBgpVrfArrayOutput() BgpVrfArrayOutput
	ToBgpVrfArrayOutputWithContext(context.Context) BgpVrfArrayOutput
}

type BgpVrfArray []BgpVrfInput

func (BgpVrfArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BgpVrf)(nil)).Elem()
}

func (i BgpVrfArray) ToBgpVrfArrayOutput() BgpVrfArrayOutput {
	return i.ToBgpVrfArrayOutputWithContext(context.Background())
}

func (i BgpVrfArray) ToBgpVrfArrayOutputWithContext(ctx context.Context) BgpVrfArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BgpVrfArrayOutput)
}

func (i BgpVrfArray) ToOutput(ctx context.Context) pulumix.Output[[]*BgpVrf] {
	return pulumix.Output[[]*BgpVrf]{
		OutputState: i.ToBgpVrfArrayOutputWithContext(ctx).OutputState,
	}
}

// BgpVrfMapInput is an input type that accepts BgpVrfMap and BgpVrfMapOutput values.
// You can construct a concrete instance of `BgpVrfMapInput` via:
//
//	BgpVrfMap{ "key": BgpVrfArgs{...} }
type BgpVrfMapInput interface {
	pulumi.Input

	ToBgpVrfMapOutput() BgpVrfMapOutput
	ToBgpVrfMapOutputWithContext(context.Context) BgpVrfMapOutput
}

type BgpVrfMap map[string]BgpVrfInput

func (BgpVrfMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BgpVrf)(nil)).Elem()
}

func (i BgpVrfMap) ToBgpVrfMapOutput() BgpVrfMapOutput {
	return i.ToBgpVrfMapOutputWithContext(context.Background())
}

func (i BgpVrfMap) ToBgpVrfMapOutputWithContext(ctx context.Context) BgpVrfMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BgpVrfMapOutput)
}

func (i BgpVrfMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*BgpVrf] {
	return pulumix.Output[map[string]*BgpVrf]{
		OutputState: i.ToBgpVrfMapOutputWithContext(ctx).OutputState,
	}
}

type BgpVrfOutput struct{ *pulumi.OutputState }

func (BgpVrfOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BgpVrf)(nil)).Elem()
}

func (o BgpVrfOutput) ToBgpVrfOutput() BgpVrfOutput {
	return o
}

func (o BgpVrfOutput) ToBgpVrfOutputWithContext(ctx context.Context) BgpVrfOutput {
	return o
}

func (o BgpVrfOutput) ToOutput(ctx context.Context) pulumix.Output[*BgpVrf] {
	return pulumix.Output[*BgpVrf]{
		OutputState: o.OutputState,
	}
}

// Autonomous system number.
func (o BgpVrfOutput) Asn() pulumi.StringOutput {
	return o.ApplyT(func(v *BgpVrf) pulumi.StringOutput { return v.Asn }).(pulumi.StringOutput)
}

// A device name from the provider configuration.
func (o BgpVrfOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BgpVrf) pulumi.StringPtrOutput { return v.Device }).(pulumi.StringPtrOutput)
}

// VRF name.
func (o BgpVrfOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BgpVrf) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Router ID.
func (o BgpVrfOutput) RouterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BgpVrf) pulumi.StringPtrOutput { return v.RouterId }).(pulumi.StringPtrOutput)
}

type BgpVrfArrayOutput struct{ *pulumi.OutputState }

func (BgpVrfArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BgpVrf)(nil)).Elem()
}

func (o BgpVrfArrayOutput) ToBgpVrfArrayOutput() BgpVrfArrayOutput {
	return o
}

func (o BgpVrfArrayOutput) ToBgpVrfArrayOutputWithContext(ctx context.Context) BgpVrfArrayOutput {
	return o
}

func (o BgpVrfArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*BgpVrf] {
	return pulumix.Output[[]*BgpVrf]{
		OutputState: o.OutputState,
	}
}

func (o BgpVrfArrayOutput) Index(i pulumi.IntInput) BgpVrfOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BgpVrf {
		return vs[0].([]*BgpVrf)[vs[1].(int)]
	}).(BgpVrfOutput)
}

type BgpVrfMapOutput struct{ *pulumi.OutputState }

func (BgpVrfMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BgpVrf)(nil)).Elem()
}

func (o BgpVrfMapOutput) ToBgpVrfMapOutput() BgpVrfMapOutput {
	return o
}

func (o BgpVrfMapOutput) ToBgpVrfMapOutputWithContext(ctx context.Context) BgpVrfMapOutput {
	return o
}

func (o BgpVrfMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*BgpVrf] {
	return pulumix.Output[map[string]*BgpVrf]{
		OutputState: o.OutputState,
	}
}

func (o BgpVrfMapOutput) MapIndex(k pulumi.StringInput) BgpVrfOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BgpVrf {
		return vs[0].(map[string]*BgpVrf)[vs[1].(string)]
	}).(BgpVrfOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BgpVrfInput)(nil)).Elem(), &BgpVrf{})
	pulumi.RegisterInputType(reflect.TypeOf((*BgpVrfArrayInput)(nil)).Elem(), BgpVrfArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BgpVrfMapInput)(nil)).Elem(), BgpVrfMap{})
	pulumi.RegisterOutputType(BgpVrfOutput{})
	pulumi.RegisterOutputType(BgpVrfArrayOutput{})
	pulumi.RegisterOutputType(BgpVrfMapOutput{})
}
