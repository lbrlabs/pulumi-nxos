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

// This data source can read the OSPF instance configuration.
//
// - API Documentation: [ospfInst](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/ospf:Inst/)
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
//			_, err := nxos.LookupOspfInstance(ctx, &nxos.LookupOspfInstanceArgs{
//				Name: "OSPF1",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupOspfInstance(ctx *pulumi.Context, args *LookupOspfInstanceArgs, opts ...pulumi.InvokeOption) (*LookupOspfInstanceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupOspfInstanceResult
	err := ctx.Invoke("nxos:index/getOspfInstance:getOspfInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOspfInstance.
type LookupOspfInstanceArgs struct {
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// OSPF instance name.
	Name string `pulumi:"name"`
}

// A collection of values returned by getOspfInstance.
type LookupOspfInstanceResult struct {
	// Administrative state.
	AdminState string `pulumi:"adminState"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// The distinguished name of the object.
	Id string `pulumi:"id"`
	// OSPF instance name.
	Name string `pulumi:"name"`
}

func LookupOspfInstanceOutput(ctx *pulumi.Context, args LookupOspfInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupOspfInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOspfInstanceResult, error) {
			args := v.(LookupOspfInstanceArgs)
			r, err := LookupOspfInstance(ctx, &args, opts...)
			var s LookupOspfInstanceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOspfInstanceResultOutput)
}

// A collection of arguments for invoking getOspfInstance.
type LookupOspfInstanceOutputArgs struct {
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput `pulumi:"device"`
	// OSPF instance name.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupOspfInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOspfInstanceArgs)(nil)).Elem()
}

// A collection of values returned by getOspfInstance.
type LookupOspfInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupOspfInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOspfInstanceResult)(nil)).Elem()
}

func (o LookupOspfInstanceResultOutput) ToLookupOspfInstanceResultOutput() LookupOspfInstanceResultOutput {
	return o
}

func (o LookupOspfInstanceResultOutput) ToLookupOspfInstanceResultOutputWithContext(ctx context.Context) LookupOspfInstanceResultOutput {
	return o
}

func (o LookupOspfInstanceResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupOspfInstanceResult] {
	return pulumix.Output[LookupOspfInstanceResult]{
		OutputState: o.OutputState,
	}
}

// Administrative state.
func (o LookupOspfInstanceResultOutput) AdminState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOspfInstanceResult) string { return v.AdminState }).(pulumi.StringOutput)
}

// A device name from the provider configuration.
func (o LookupOspfInstanceResultOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOspfInstanceResult) *string { return v.Device }).(pulumi.StringPtrOutput)
}

// The distinguished name of the object.
func (o LookupOspfInstanceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOspfInstanceResult) string { return v.Id }).(pulumi.StringOutput)
}

// OSPF instance name.
func (o LookupOspfInstanceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOspfInstanceResult) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOspfInstanceResultOutput{})
}