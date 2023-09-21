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

type Rest struct {
	pulumi.CustomResourceState

	// List of children.
	Childrens RestChildrenArrayOutput `pulumi:"childrens"`
	// Which class object is being created. (Make sure there is no colon in the classname)
	ClassName pulumi.StringOutput `pulumi:"className"`
	// Map of key-value pairs that need to be passed to the Model object as parameters.
	Content pulumi.StringMapOutput `pulumi:"content"`
	// Delete object during destroy operation. Default value is `true`.
	Delete pulumi.BoolOutput `pulumi:"delete"`
	// A device name from the provider configuration.
	Device pulumi.StringPtrOutput `pulumi:"device"`
	// Distinguished name of object being managed including its relative name, e.g. sys/intf/phys-[eth1/1].
	Dn pulumi.StringOutput `pulumi:"dn"`
}

// NewRest registers a new resource with the given unique name, arguments, and options.
func NewRest(ctx *pulumi.Context,
	name string, args *RestArgs, opts ...pulumi.ResourceOption) (*Rest, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClassName == nil {
		return nil, errors.New("invalid value for required argument 'ClassName'")
	}
	if args.Dn == nil {
		return nil, errors.New("invalid value for required argument 'Dn'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Rest
	err := ctx.RegisterResource("nxos:nxos/rest:Rest", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRest gets an existing Rest resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRest(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RestState, opts ...pulumi.ResourceOption) (*Rest, error) {
	var resource Rest
	err := ctx.ReadResource("nxos:nxos/rest:Rest", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Rest resources.
type restState struct {
	// List of children.
	Childrens []RestChildren `pulumi:"childrens"`
	// Which class object is being created. (Make sure there is no colon in the classname)
	ClassName *string `pulumi:"className"`
	// Map of key-value pairs that need to be passed to the Model object as parameters.
	Content map[string]string `pulumi:"content"`
	// Delete object during destroy operation. Default value is `true`.
	Delete *bool `pulumi:"delete"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// Distinguished name of object being managed including its relative name, e.g. sys/intf/phys-[eth1/1].
	Dn *string `pulumi:"dn"`
}

type RestState struct {
	// List of children.
	Childrens RestChildrenArrayInput
	// Which class object is being created. (Make sure there is no colon in the classname)
	ClassName pulumi.StringPtrInput
	// Map of key-value pairs that need to be passed to the Model object as parameters.
	Content pulumi.StringMapInput
	// Delete object during destroy operation. Default value is `true`.
	Delete pulumi.BoolPtrInput
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
	// Distinguished name of object being managed including its relative name, e.g. sys/intf/phys-[eth1/1].
	Dn pulumi.StringPtrInput
}

func (RestState) ElementType() reflect.Type {
	return reflect.TypeOf((*restState)(nil)).Elem()
}

type restArgs struct {
	// List of children.
	Childrens []RestChildren `pulumi:"childrens"`
	// Which class object is being created. (Make sure there is no colon in the classname)
	ClassName string `pulumi:"className"`
	// Map of key-value pairs that need to be passed to the Model object as parameters.
	Content map[string]string `pulumi:"content"`
	// Delete object during destroy operation. Default value is `true`.
	Delete *bool `pulumi:"delete"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// Distinguished name of object being managed including its relative name, e.g. sys/intf/phys-[eth1/1].
	Dn string `pulumi:"dn"`
}

// The set of arguments for constructing a Rest resource.
type RestArgs struct {
	// List of children.
	Childrens RestChildrenArrayInput
	// Which class object is being created. (Make sure there is no colon in the classname)
	ClassName pulumi.StringInput
	// Map of key-value pairs that need to be passed to the Model object as parameters.
	Content pulumi.StringMapInput
	// Delete object during destroy operation. Default value is `true`.
	Delete pulumi.BoolPtrInput
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
	// Distinguished name of object being managed including its relative name, e.g. sys/intf/phys-[eth1/1].
	Dn pulumi.StringInput
}

func (RestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*restArgs)(nil)).Elem()
}

type RestInput interface {
	pulumi.Input

	ToRestOutput() RestOutput
	ToRestOutputWithContext(ctx context.Context) RestOutput
}

func (*Rest) ElementType() reflect.Type {
	return reflect.TypeOf((**Rest)(nil)).Elem()
}

func (i *Rest) ToRestOutput() RestOutput {
	return i.ToRestOutputWithContext(context.Background())
}

func (i *Rest) ToRestOutputWithContext(ctx context.Context) RestOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestOutput)
}

func (i *Rest) ToOutput(ctx context.Context) pulumix.Output[*Rest] {
	return pulumix.Output[*Rest]{
		OutputState: i.ToRestOutputWithContext(ctx).OutputState,
	}
}

// RestArrayInput is an input type that accepts RestArray and RestArrayOutput values.
// You can construct a concrete instance of `RestArrayInput` via:
//
//	RestArray{ RestArgs{...} }
type RestArrayInput interface {
	pulumi.Input

	ToRestArrayOutput() RestArrayOutput
	ToRestArrayOutputWithContext(context.Context) RestArrayOutput
}

type RestArray []RestInput

func (RestArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Rest)(nil)).Elem()
}

func (i RestArray) ToRestArrayOutput() RestArrayOutput {
	return i.ToRestArrayOutputWithContext(context.Background())
}

func (i RestArray) ToRestArrayOutputWithContext(ctx context.Context) RestArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestArrayOutput)
}

func (i RestArray) ToOutput(ctx context.Context) pulumix.Output[[]*Rest] {
	return pulumix.Output[[]*Rest]{
		OutputState: i.ToRestArrayOutputWithContext(ctx).OutputState,
	}
}

// RestMapInput is an input type that accepts RestMap and RestMapOutput values.
// You can construct a concrete instance of `RestMapInput` via:
//
//	RestMap{ "key": RestArgs{...} }
type RestMapInput interface {
	pulumi.Input

	ToRestMapOutput() RestMapOutput
	ToRestMapOutputWithContext(context.Context) RestMapOutput
}

type RestMap map[string]RestInput

func (RestMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Rest)(nil)).Elem()
}

func (i RestMap) ToRestMapOutput() RestMapOutput {
	return i.ToRestMapOutputWithContext(context.Background())
}

func (i RestMap) ToRestMapOutputWithContext(ctx context.Context) RestMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestMapOutput)
}

func (i RestMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Rest] {
	return pulumix.Output[map[string]*Rest]{
		OutputState: i.ToRestMapOutputWithContext(ctx).OutputState,
	}
}

type RestOutput struct{ *pulumi.OutputState }

func (RestOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Rest)(nil)).Elem()
}

func (o RestOutput) ToRestOutput() RestOutput {
	return o
}

func (o RestOutput) ToRestOutputWithContext(ctx context.Context) RestOutput {
	return o
}

func (o RestOutput) ToOutput(ctx context.Context) pulumix.Output[*Rest] {
	return pulumix.Output[*Rest]{
		OutputState: o.OutputState,
	}
}

// List of children.
func (o RestOutput) Childrens() RestChildrenArrayOutput {
	return o.ApplyT(func(v *Rest) RestChildrenArrayOutput { return v.Childrens }).(RestChildrenArrayOutput)
}

// Which class object is being created. (Make sure there is no colon in the classname)
func (o RestOutput) ClassName() pulumi.StringOutput {
	return o.ApplyT(func(v *Rest) pulumi.StringOutput { return v.ClassName }).(pulumi.StringOutput)
}

// Map of key-value pairs that need to be passed to the Model object as parameters.
func (o RestOutput) Content() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Rest) pulumi.StringMapOutput { return v.Content }).(pulumi.StringMapOutput)
}

// Delete object during destroy operation. Default value is `true`.
func (o RestOutput) Delete() pulumi.BoolOutput {
	return o.ApplyT(func(v *Rest) pulumi.BoolOutput { return v.Delete }).(pulumi.BoolOutput)
}

// A device name from the provider configuration.
func (o RestOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rest) pulumi.StringPtrOutput { return v.Device }).(pulumi.StringPtrOutput)
}

// Distinguished name of object being managed including its relative name, e.g. sys/intf/phys-[eth1/1].
func (o RestOutput) Dn() pulumi.StringOutput {
	return o.ApplyT(func(v *Rest) pulumi.StringOutput { return v.Dn }).(pulumi.StringOutput)
}

type RestArrayOutput struct{ *pulumi.OutputState }

func (RestArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Rest)(nil)).Elem()
}

func (o RestArrayOutput) ToRestArrayOutput() RestArrayOutput {
	return o
}

func (o RestArrayOutput) ToRestArrayOutputWithContext(ctx context.Context) RestArrayOutput {
	return o
}

func (o RestArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Rest] {
	return pulumix.Output[[]*Rest]{
		OutputState: o.OutputState,
	}
}

func (o RestArrayOutput) Index(i pulumi.IntInput) RestOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Rest {
		return vs[0].([]*Rest)[vs[1].(int)]
	}).(RestOutput)
}

type RestMapOutput struct{ *pulumi.OutputState }

func (RestMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Rest)(nil)).Elem()
}

func (o RestMapOutput) ToRestMapOutput() RestMapOutput {
	return o
}

func (o RestMapOutput) ToRestMapOutputWithContext(ctx context.Context) RestMapOutput {
	return o
}

func (o RestMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Rest] {
	return pulumix.Output[map[string]*Rest]{
		OutputState: o.OutputState,
	}
}

func (o RestMapOutput) MapIndex(k pulumi.StringInput) RestOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Rest {
		return vs[0].(map[string]*Rest)[vs[1].(string)]
	}).(RestOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RestInput)(nil)).Elem(), &Rest{})
	pulumi.RegisterInputType(reflect.TypeOf((*RestArrayInput)(nil)).Elem(), RestArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RestMapInput)(nil)).Elem(), RestMap{})
	pulumi.RegisterOutputType(RestOutput{})
	pulumi.RegisterOutputType(RestArrayOutput{})
	pulumi.RegisterOutputType(RestMapOutput{})
}
