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

// This resource can manage the PIM interface configuration.
//
// - API Documentation: [pimIf](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/pim:If/)
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
//			_, err := nxos.NewPimInterface(ctx, "example", &nxos.PimInterfaceArgs{
//				AdminState:  pulumi.String("enabled"),
//				Bfd:         pulumi.String("enabled"),
//				DrPriority:  pulumi.Int(10),
//				InterfaceId: pulumi.String("eth1/10"),
//				Passive:     pulumi.Bool(false),
//				SparseMode:  pulumi.Bool(true),
//				VrfName:     pulumi.String("default"),
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
//	$ pulumi import nxos:index/pimInterface:PimInterface example "sys/pim/inst/dom-[default]/if-[eth1/10]"
//
// ```
type PimInterface struct {
	pulumi.CustomResourceState

	// Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
	AdminState pulumi.StringOutput `pulumi:"adminState"`
	// BFD. - Choices: `none`, `enabled`, `disabled` - Default value: `none`
	Bfd pulumi.StringOutput `pulumi:"bfd"`
	// A device name from the provider configuration.
	Device pulumi.StringPtrOutput `pulumi:"device"`
	// Designated Router priority level. - Range: `1`-`4294967295` - Default value: `1`
	DrPriority pulumi.IntOutput `pulumi:"drPriority"`
	// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
	InterfaceId pulumi.StringOutput `pulumi:"interfaceId"`
	// Passive interface. - Default value: `false`
	Passive pulumi.BoolOutput `pulumi:"passive"`
	// Sparse mode. - Default value: `false`
	SparseMode pulumi.BoolOutput `pulumi:"sparseMode"`
	// VRF name.
	VrfName pulumi.StringOutput `pulumi:"vrfName"`
}

// NewPimInterface registers a new resource with the given unique name, arguments, and options.
func NewPimInterface(ctx *pulumi.Context,
	name string, args *PimInterfaceArgs, opts ...pulumi.ResourceOption) (*PimInterface, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InterfaceId == nil {
		return nil, errors.New("invalid value for required argument 'InterfaceId'")
	}
	if args.VrfName == nil {
		return nil, errors.New("invalid value for required argument 'VrfName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PimInterface
	err := ctx.RegisterResource("nxos:index/pimInterface:PimInterface", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPimInterface gets an existing PimInterface resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPimInterface(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PimInterfaceState, opts ...pulumi.ResourceOption) (*PimInterface, error) {
	var resource PimInterface
	err := ctx.ReadResource("nxos:index/pimInterface:PimInterface", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PimInterface resources.
type pimInterfaceState struct {
	// Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
	AdminState *string `pulumi:"adminState"`
	// BFD. - Choices: `none`, `enabled`, `disabled` - Default value: `none`
	Bfd *string `pulumi:"bfd"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// Designated Router priority level. - Range: `1`-`4294967295` - Default value: `1`
	DrPriority *int `pulumi:"drPriority"`
	// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
	InterfaceId *string `pulumi:"interfaceId"`
	// Passive interface. - Default value: `false`
	Passive *bool `pulumi:"passive"`
	// Sparse mode. - Default value: `false`
	SparseMode *bool `pulumi:"sparseMode"`
	// VRF name.
	VrfName *string `pulumi:"vrfName"`
}

type PimInterfaceState struct {
	// Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
	AdminState pulumi.StringPtrInput
	// BFD. - Choices: `none`, `enabled`, `disabled` - Default value: `none`
	Bfd pulumi.StringPtrInput
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
	// Designated Router priority level. - Range: `1`-`4294967295` - Default value: `1`
	DrPriority pulumi.IntPtrInput
	// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
	InterfaceId pulumi.StringPtrInput
	// Passive interface. - Default value: `false`
	Passive pulumi.BoolPtrInput
	// Sparse mode. - Default value: `false`
	SparseMode pulumi.BoolPtrInput
	// VRF name.
	VrfName pulumi.StringPtrInput
}

func (PimInterfaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*pimInterfaceState)(nil)).Elem()
}

type pimInterfaceArgs struct {
	// Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
	AdminState *string `pulumi:"adminState"`
	// BFD. - Choices: `none`, `enabled`, `disabled` - Default value: `none`
	Bfd *string `pulumi:"bfd"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// Designated Router priority level. - Range: `1`-`4294967295` - Default value: `1`
	DrPriority *int `pulumi:"drPriority"`
	// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
	InterfaceId string `pulumi:"interfaceId"`
	// Passive interface. - Default value: `false`
	Passive *bool `pulumi:"passive"`
	// Sparse mode. - Default value: `false`
	SparseMode *bool `pulumi:"sparseMode"`
	// VRF name.
	VrfName string `pulumi:"vrfName"`
}

// The set of arguments for constructing a PimInterface resource.
type PimInterfaceArgs struct {
	// Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
	AdminState pulumi.StringPtrInput
	// BFD. - Choices: `none`, `enabled`, `disabled` - Default value: `none`
	Bfd pulumi.StringPtrInput
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
	// Designated Router priority level. - Range: `1`-`4294967295` - Default value: `1`
	DrPriority pulumi.IntPtrInput
	// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
	InterfaceId pulumi.StringInput
	// Passive interface. - Default value: `false`
	Passive pulumi.BoolPtrInput
	// Sparse mode. - Default value: `false`
	SparseMode pulumi.BoolPtrInput
	// VRF name.
	VrfName pulumi.StringInput
}

func (PimInterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pimInterfaceArgs)(nil)).Elem()
}

type PimInterfaceInput interface {
	pulumi.Input

	ToPimInterfaceOutput() PimInterfaceOutput
	ToPimInterfaceOutputWithContext(ctx context.Context) PimInterfaceOutput
}

func (*PimInterface) ElementType() reflect.Type {
	return reflect.TypeOf((**PimInterface)(nil)).Elem()
}

func (i *PimInterface) ToPimInterfaceOutput() PimInterfaceOutput {
	return i.ToPimInterfaceOutputWithContext(context.Background())
}

func (i *PimInterface) ToPimInterfaceOutputWithContext(ctx context.Context) PimInterfaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PimInterfaceOutput)
}

func (i *PimInterface) ToOutput(ctx context.Context) pulumix.Output[*PimInterface] {
	return pulumix.Output[*PimInterface]{
		OutputState: i.ToPimInterfaceOutputWithContext(ctx).OutputState,
	}
}

// PimInterfaceArrayInput is an input type that accepts PimInterfaceArray and PimInterfaceArrayOutput values.
// You can construct a concrete instance of `PimInterfaceArrayInput` via:
//
//	PimInterfaceArray{ PimInterfaceArgs{...} }
type PimInterfaceArrayInput interface {
	pulumi.Input

	ToPimInterfaceArrayOutput() PimInterfaceArrayOutput
	ToPimInterfaceArrayOutputWithContext(context.Context) PimInterfaceArrayOutput
}

type PimInterfaceArray []PimInterfaceInput

func (PimInterfaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PimInterface)(nil)).Elem()
}

func (i PimInterfaceArray) ToPimInterfaceArrayOutput() PimInterfaceArrayOutput {
	return i.ToPimInterfaceArrayOutputWithContext(context.Background())
}

func (i PimInterfaceArray) ToPimInterfaceArrayOutputWithContext(ctx context.Context) PimInterfaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PimInterfaceArrayOutput)
}

func (i PimInterfaceArray) ToOutput(ctx context.Context) pulumix.Output[[]*PimInterface] {
	return pulumix.Output[[]*PimInterface]{
		OutputState: i.ToPimInterfaceArrayOutputWithContext(ctx).OutputState,
	}
}

// PimInterfaceMapInput is an input type that accepts PimInterfaceMap and PimInterfaceMapOutput values.
// You can construct a concrete instance of `PimInterfaceMapInput` via:
//
//	PimInterfaceMap{ "key": PimInterfaceArgs{...} }
type PimInterfaceMapInput interface {
	pulumi.Input

	ToPimInterfaceMapOutput() PimInterfaceMapOutput
	ToPimInterfaceMapOutputWithContext(context.Context) PimInterfaceMapOutput
}

type PimInterfaceMap map[string]PimInterfaceInput

func (PimInterfaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PimInterface)(nil)).Elem()
}

func (i PimInterfaceMap) ToPimInterfaceMapOutput() PimInterfaceMapOutput {
	return i.ToPimInterfaceMapOutputWithContext(context.Background())
}

func (i PimInterfaceMap) ToPimInterfaceMapOutputWithContext(ctx context.Context) PimInterfaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PimInterfaceMapOutput)
}

func (i PimInterfaceMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*PimInterface] {
	return pulumix.Output[map[string]*PimInterface]{
		OutputState: i.ToPimInterfaceMapOutputWithContext(ctx).OutputState,
	}
}

type PimInterfaceOutput struct{ *pulumi.OutputState }

func (PimInterfaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PimInterface)(nil)).Elem()
}

func (o PimInterfaceOutput) ToPimInterfaceOutput() PimInterfaceOutput {
	return o
}

func (o PimInterfaceOutput) ToPimInterfaceOutputWithContext(ctx context.Context) PimInterfaceOutput {
	return o
}

func (o PimInterfaceOutput) ToOutput(ctx context.Context) pulumix.Output[*PimInterface] {
	return pulumix.Output[*PimInterface]{
		OutputState: o.OutputState,
	}
}

// Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
func (o PimInterfaceOutput) AdminState() pulumi.StringOutput {
	return o.ApplyT(func(v *PimInterface) pulumi.StringOutput { return v.AdminState }).(pulumi.StringOutput)
}

// BFD. - Choices: `none`, `enabled`, `disabled` - Default value: `none`
func (o PimInterfaceOutput) Bfd() pulumi.StringOutput {
	return o.ApplyT(func(v *PimInterface) pulumi.StringOutput { return v.Bfd }).(pulumi.StringOutput)
}

// A device name from the provider configuration.
func (o PimInterfaceOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PimInterface) pulumi.StringPtrOutput { return v.Device }).(pulumi.StringPtrOutput)
}

// Designated Router priority level. - Range: `1`-`4294967295` - Default value: `1`
func (o PimInterfaceOutput) DrPriority() pulumi.IntOutput {
	return o.ApplyT(func(v *PimInterface) pulumi.IntOutput { return v.DrPriority }).(pulumi.IntOutput)
}

// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
func (o PimInterfaceOutput) InterfaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *PimInterface) pulumi.StringOutput { return v.InterfaceId }).(pulumi.StringOutput)
}

// Passive interface. - Default value: `false`
func (o PimInterfaceOutput) Passive() pulumi.BoolOutput {
	return o.ApplyT(func(v *PimInterface) pulumi.BoolOutput { return v.Passive }).(pulumi.BoolOutput)
}

// Sparse mode. - Default value: `false`
func (o PimInterfaceOutput) SparseMode() pulumi.BoolOutput {
	return o.ApplyT(func(v *PimInterface) pulumi.BoolOutput { return v.SparseMode }).(pulumi.BoolOutput)
}

// VRF name.
func (o PimInterfaceOutput) VrfName() pulumi.StringOutput {
	return o.ApplyT(func(v *PimInterface) pulumi.StringOutput { return v.VrfName }).(pulumi.StringOutput)
}

type PimInterfaceArrayOutput struct{ *pulumi.OutputState }

func (PimInterfaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PimInterface)(nil)).Elem()
}

func (o PimInterfaceArrayOutput) ToPimInterfaceArrayOutput() PimInterfaceArrayOutput {
	return o
}

func (o PimInterfaceArrayOutput) ToPimInterfaceArrayOutputWithContext(ctx context.Context) PimInterfaceArrayOutput {
	return o
}

func (o PimInterfaceArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*PimInterface] {
	return pulumix.Output[[]*PimInterface]{
		OutputState: o.OutputState,
	}
}

func (o PimInterfaceArrayOutput) Index(i pulumi.IntInput) PimInterfaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PimInterface {
		return vs[0].([]*PimInterface)[vs[1].(int)]
	}).(PimInterfaceOutput)
}

type PimInterfaceMapOutput struct{ *pulumi.OutputState }

func (PimInterfaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PimInterface)(nil)).Elem()
}

func (o PimInterfaceMapOutput) ToPimInterfaceMapOutput() PimInterfaceMapOutput {
	return o
}

func (o PimInterfaceMapOutput) ToPimInterfaceMapOutputWithContext(ctx context.Context) PimInterfaceMapOutput {
	return o
}

func (o PimInterfaceMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*PimInterface] {
	return pulumix.Output[map[string]*PimInterface]{
		OutputState: o.OutputState,
	}
}

func (o PimInterfaceMapOutput) MapIndex(k pulumi.StringInput) PimInterfaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PimInterface {
		return vs[0].(map[string]*PimInterface)[vs[1].(string)]
	}).(PimInterfaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PimInterfaceInput)(nil)).Elem(), &PimInterface{})
	pulumi.RegisterInputType(reflect.TypeOf((*PimInterfaceArrayInput)(nil)).Elem(), PimInterfaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PimInterfaceMapInput)(nil)).Elem(), PimInterfaceMap{})
	pulumi.RegisterOutputType(PimInterfaceOutput{})
	pulumi.RegisterOutputType(PimInterfaceArrayOutput{})
	pulumi.RegisterOutputType(PimInterfaceMapOutput{})
}
