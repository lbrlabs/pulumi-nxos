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

// This data source can read the IS-IS instance configuration.
//
// - API Documentation: [isisInst](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/isis:Inst/)
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
//			_, err := nxos.LookupIsisInstance(ctx, &nxos.LookupIsisInstanceArgs{
//				Name: "ISIS1",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupIsisInstance(ctx *pulumi.Context, args *LookupIsisInstanceArgs, opts ...pulumi.InvokeOption) (*LookupIsisInstanceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupIsisInstanceResult
	err := ctx.Invoke("nxos:index/getIsisInstance:getIsisInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIsisInstance.
type LookupIsisInstanceArgs struct {
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// IS-IS instance name.
	Name string `pulumi:"name"`
}

// A collection of values returned by getIsisInstance.
type LookupIsisInstanceResult struct {
	// Administrative state.
	AdminState string `pulumi:"adminState"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// The distinguished name of the object.
	Id string `pulumi:"id"`
	// IS-IS instance name.
	Name string `pulumi:"name"`
}

func LookupIsisInstanceOutput(ctx *pulumi.Context, args LookupIsisInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupIsisInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIsisInstanceResult, error) {
			args := v.(LookupIsisInstanceArgs)
			r, err := LookupIsisInstance(ctx, &args, opts...)
			var s LookupIsisInstanceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIsisInstanceResultOutput)
}

// A collection of arguments for invoking getIsisInstance.
type LookupIsisInstanceOutputArgs struct {
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput `pulumi:"device"`
	// IS-IS instance name.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupIsisInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIsisInstanceArgs)(nil)).Elem()
}

// A collection of values returned by getIsisInstance.
type LookupIsisInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupIsisInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIsisInstanceResult)(nil)).Elem()
}

func (o LookupIsisInstanceResultOutput) ToLookupIsisInstanceResultOutput() LookupIsisInstanceResultOutput {
	return o
}

func (o LookupIsisInstanceResultOutput) ToLookupIsisInstanceResultOutputWithContext(ctx context.Context) LookupIsisInstanceResultOutput {
	return o
}

func (o LookupIsisInstanceResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupIsisInstanceResult] {
	return pulumix.Output[LookupIsisInstanceResult]{
		OutputState: o.OutputState,
	}
}

// Administrative state.
func (o LookupIsisInstanceResultOutput) AdminState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIsisInstanceResult) string { return v.AdminState }).(pulumi.StringOutput)
}

// A device name from the provider configuration.
func (o LookupIsisInstanceResultOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIsisInstanceResult) *string { return v.Device }).(pulumi.StringPtrOutput)
}

// The distinguished name of the object.
func (o LookupIsisInstanceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIsisInstanceResult) string { return v.Id }).(pulumi.StringOutput)
}

// IS-IS instance name.
func (o LookupIsisInstanceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIsisInstanceResult) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIsisInstanceResultOutput{})
}
