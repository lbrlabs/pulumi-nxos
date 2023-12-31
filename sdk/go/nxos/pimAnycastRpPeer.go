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

// This resource can manage the PIM Anycast RP peer configuration.
//
// - API Documentation: [pimAcastRPPeer](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/pim:AcastRPPeer/)
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
//			_, err := nxos.NewPimAnycastRpPeer(ctx, "example", &nxos.PimAnycastRpPeerArgs{
//				Address:      pulumi.String("10.1.1.1/32"),
//				RpSetAddress: pulumi.String("20.1.1.1/32"),
//				VrfName:      pulumi.String("default"),
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
//	$ pulumi import nxos:index/pimAnycastRpPeer:PimAnycastRpPeer example "sys/pim/inst/dom-[default]/acastrpfunc/peer-[10.1.1.1/32]-peer-[20.1.1.1/32]"
//
// ```
type PimAnycastRpPeer struct {
	pulumi.CustomResourceState

	// Anycast RP address.
	Address pulumi.StringOutput `pulumi:"address"`
	// A device name from the provider configuration.
	Device pulumi.StringPtrOutput `pulumi:"device"`
	// RP set address.
	RpSetAddress pulumi.StringOutput `pulumi:"rpSetAddress"`
	// VRF name.
	VrfName pulumi.StringOutput `pulumi:"vrfName"`
}

// NewPimAnycastRpPeer registers a new resource with the given unique name, arguments, and options.
func NewPimAnycastRpPeer(ctx *pulumi.Context,
	name string, args *PimAnycastRpPeerArgs, opts ...pulumi.ResourceOption) (*PimAnycastRpPeer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Address == nil {
		return nil, errors.New("invalid value for required argument 'Address'")
	}
	if args.RpSetAddress == nil {
		return nil, errors.New("invalid value for required argument 'RpSetAddress'")
	}
	if args.VrfName == nil {
		return nil, errors.New("invalid value for required argument 'VrfName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PimAnycastRpPeer
	err := ctx.RegisterResource("nxos:index/pimAnycastRpPeer:PimAnycastRpPeer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPimAnycastRpPeer gets an existing PimAnycastRpPeer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPimAnycastRpPeer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PimAnycastRpPeerState, opts ...pulumi.ResourceOption) (*PimAnycastRpPeer, error) {
	var resource PimAnycastRpPeer
	err := ctx.ReadResource("nxos:index/pimAnycastRpPeer:PimAnycastRpPeer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PimAnycastRpPeer resources.
type pimAnycastRpPeerState struct {
	// Anycast RP address.
	Address *string `pulumi:"address"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// RP set address.
	RpSetAddress *string `pulumi:"rpSetAddress"`
	// VRF name.
	VrfName *string `pulumi:"vrfName"`
}

type PimAnycastRpPeerState struct {
	// Anycast RP address.
	Address pulumi.StringPtrInput
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
	// RP set address.
	RpSetAddress pulumi.StringPtrInput
	// VRF name.
	VrfName pulumi.StringPtrInput
}

func (PimAnycastRpPeerState) ElementType() reflect.Type {
	return reflect.TypeOf((*pimAnycastRpPeerState)(nil)).Elem()
}

type pimAnycastRpPeerArgs struct {
	// Anycast RP address.
	Address string `pulumi:"address"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// RP set address.
	RpSetAddress string `pulumi:"rpSetAddress"`
	// VRF name.
	VrfName string `pulumi:"vrfName"`
}

// The set of arguments for constructing a PimAnycastRpPeer resource.
type PimAnycastRpPeerArgs struct {
	// Anycast RP address.
	Address pulumi.StringInput
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
	// RP set address.
	RpSetAddress pulumi.StringInput
	// VRF name.
	VrfName pulumi.StringInput
}

func (PimAnycastRpPeerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pimAnycastRpPeerArgs)(nil)).Elem()
}

type PimAnycastRpPeerInput interface {
	pulumi.Input

	ToPimAnycastRpPeerOutput() PimAnycastRpPeerOutput
	ToPimAnycastRpPeerOutputWithContext(ctx context.Context) PimAnycastRpPeerOutput
}

func (*PimAnycastRpPeer) ElementType() reflect.Type {
	return reflect.TypeOf((**PimAnycastRpPeer)(nil)).Elem()
}

func (i *PimAnycastRpPeer) ToPimAnycastRpPeerOutput() PimAnycastRpPeerOutput {
	return i.ToPimAnycastRpPeerOutputWithContext(context.Background())
}

func (i *PimAnycastRpPeer) ToPimAnycastRpPeerOutputWithContext(ctx context.Context) PimAnycastRpPeerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PimAnycastRpPeerOutput)
}

func (i *PimAnycastRpPeer) ToOutput(ctx context.Context) pulumix.Output[*PimAnycastRpPeer] {
	return pulumix.Output[*PimAnycastRpPeer]{
		OutputState: i.ToPimAnycastRpPeerOutputWithContext(ctx).OutputState,
	}
}

// PimAnycastRpPeerArrayInput is an input type that accepts PimAnycastRpPeerArray and PimAnycastRpPeerArrayOutput values.
// You can construct a concrete instance of `PimAnycastRpPeerArrayInput` via:
//
//	PimAnycastRpPeerArray{ PimAnycastRpPeerArgs{...} }
type PimAnycastRpPeerArrayInput interface {
	pulumi.Input

	ToPimAnycastRpPeerArrayOutput() PimAnycastRpPeerArrayOutput
	ToPimAnycastRpPeerArrayOutputWithContext(context.Context) PimAnycastRpPeerArrayOutput
}

type PimAnycastRpPeerArray []PimAnycastRpPeerInput

func (PimAnycastRpPeerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PimAnycastRpPeer)(nil)).Elem()
}

func (i PimAnycastRpPeerArray) ToPimAnycastRpPeerArrayOutput() PimAnycastRpPeerArrayOutput {
	return i.ToPimAnycastRpPeerArrayOutputWithContext(context.Background())
}

func (i PimAnycastRpPeerArray) ToPimAnycastRpPeerArrayOutputWithContext(ctx context.Context) PimAnycastRpPeerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PimAnycastRpPeerArrayOutput)
}

func (i PimAnycastRpPeerArray) ToOutput(ctx context.Context) pulumix.Output[[]*PimAnycastRpPeer] {
	return pulumix.Output[[]*PimAnycastRpPeer]{
		OutputState: i.ToPimAnycastRpPeerArrayOutputWithContext(ctx).OutputState,
	}
}

// PimAnycastRpPeerMapInput is an input type that accepts PimAnycastRpPeerMap and PimAnycastRpPeerMapOutput values.
// You can construct a concrete instance of `PimAnycastRpPeerMapInput` via:
//
//	PimAnycastRpPeerMap{ "key": PimAnycastRpPeerArgs{...} }
type PimAnycastRpPeerMapInput interface {
	pulumi.Input

	ToPimAnycastRpPeerMapOutput() PimAnycastRpPeerMapOutput
	ToPimAnycastRpPeerMapOutputWithContext(context.Context) PimAnycastRpPeerMapOutput
}

type PimAnycastRpPeerMap map[string]PimAnycastRpPeerInput

func (PimAnycastRpPeerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PimAnycastRpPeer)(nil)).Elem()
}

func (i PimAnycastRpPeerMap) ToPimAnycastRpPeerMapOutput() PimAnycastRpPeerMapOutput {
	return i.ToPimAnycastRpPeerMapOutputWithContext(context.Background())
}

func (i PimAnycastRpPeerMap) ToPimAnycastRpPeerMapOutputWithContext(ctx context.Context) PimAnycastRpPeerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PimAnycastRpPeerMapOutput)
}

func (i PimAnycastRpPeerMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*PimAnycastRpPeer] {
	return pulumix.Output[map[string]*PimAnycastRpPeer]{
		OutputState: i.ToPimAnycastRpPeerMapOutputWithContext(ctx).OutputState,
	}
}

type PimAnycastRpPeerOutput struct{ *pulumi.OutputState }

func (PimAnycastRpPeerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PimAnycastRpPeer)(nil)).Elem()
}

func (o PimAnycastRpPeerOutput) ToPimAnycastRpPeerOutput() PimAnycastRpPeerOutput {
	return o
}

func (o PimAnycastRpPeerOutput) ToPimAnycastRpPeerOutputWithContext(ctx context.Context) PimAnycastRpPeerOutput {
	return o
}

func (o PimAnycastRpPeerOutput) ToOutput(ctx context.Context) pulumix.Output[*PimAnycastRpPeer] {
	return pulumix.Output[*PimAnycastRpPeer]{
		OutputState: o.OutputState,
	}
}

// Anycast RP address.
func (o PimAnycastRpPeerOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v *PimAnycastRpPeer) pulumi.StringOutput { return v.Address }).(pulumi.StringOutput)
}

// A device name from the provider configuration.
func (o PimAnycastRpPeerOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PimAnycastRpPeer) pulumi.StringPtrOutput { return v.Device }).(pulumi.StringPtrOutput)
}

// RP set address.
func (o PimAnycastRpPeerOutput) RpSetAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *PimAnycastRpPeer) pulumi.StringOutput { return v.RpSetAddress }).(pulumi.StringOutput)
}

// VRF name.
func (o PimAnycastRpPeerOutput) VrfName() pulumi.StringOutput {
	return o.ApplyT(func(v *PimAnycastRpPeer) pulumi.StringOutput { return v.VrfName }).(pulumi.StringOutput)
}

type PimAnycastRpPeerArrayOutput struct{ *pulumi.OutputState }

func (PimAnycastRpPeerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PimAnycastRpPeer)(nil)).Elem()
}

func (o PimAnycastRpPeerArrayOutput) ToPimAnycastRpPeerArrayOutput() PimAnycastRpPeerArrayOutput {
	return o
}

func (o PimAnycastRpPeerArrayOutput) ToPimAnycastRpPeerArrayOutputWithContext(ctx context.Context) PimAnycastRpPeerArrayOutput {
	return o
}

func (o PimAnycastRpPeerArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*PimAnycastRpPeer] {
	return pulumix.Output[[]*PimAnycastRpPeer]{
		OutputState: o.OutputState,
	}
}

func (o PimAnycastRpPeerArrayOutput) Index(i pulumi.IntInput) PimAnycastRpPeerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PimAnycastRpPeer {
		return vs[0].([]*PimAnycastRpPeer)[vs[1].(int)]
	}).(PimAnycastRpPeerOutput)
}

type PimAnycastRpPeerMapOutput struct{ *pulumi.OutputState }

func (PimAnycastRpPeerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PimAnycastRpPeer)(nil)).Elem()
}

func (o PimAnycastRpPeerMapOutput) ToPimAnycastRpPeerMapOutput() PimAnycastRpPeerMapOutput {
	return o
}

func (o PimAnycastRpPeerMapOutput) ToPimAnycastRpPeerMapOutputWithContext(ctx context.Context) PimAnycastRpPeerMapOutput {
	return o
}

func (o PimAnycastRpPeerMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*PimAnycastRpPeer] {
	return pulumix.Output[map[string]*PimAnycastRpPeer]{
		OutputState: o.OutputState,
	}
}

func (o PimAnycastRpPeerMapOutput) MapIndex(k pulumi.StringInput) PimAnycastRpPeerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PimAnycastRpPeer {
		return vs[0].(map[string]*PimAnycastRpPeer)[vs[1].(string)]
	}).(PimAnycastRpPeerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PimAnycastRpPeerInput)(nil)).Elem(), &PimAnycastRpPeer{})
	pulumi.RegisterInputType(reflect.TypeOf((*PimAnycastRpPeerArrayInput)(nil)).Elem(), PimAnycastRpPeerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PimAnycastRpPeerMapInput)(nil)).Elem(), PimAnycastRpPeerMap{})
	pulumi.RegisterOutputType(PimAnycastRpPeerOutput{})
	pulumi.RegisterOutputType(PimAnycastRpPeerArrayOutput{})
	pulumi.RegisterOutputType(PimAnycastRpPeerMapOutput{})
}
