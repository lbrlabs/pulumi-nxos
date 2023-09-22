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

// This data source can read the PIM SSM range configuration.
//
// - API Documentation: [pimSSMRangeP](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/pim:SSMRangeP/)
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
//			_, err := nxos.LookupPimSsmRange(ctx, &nxos.LookupPimSsmRangeArgs{
//				VrfName: "default",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupPimSsmRange(ctx *pulumi.Context, args *LookupPimSsmRangeArgs, opts ...pulumi.InvokeOption) (*LookupPimSsmRangeResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPimSsmRangeResult
	err := ctx.Invoke("nxos:index/getPimSsmRange:getPimSsmRange", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPimSsmRange.
type LookupPimSsmRangeArgs struct {
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// VRF name.
	VrfName string `pulumi:"vrfName"`
}

// A collection of values returned by getPimSsmRange.
type LookupPimSsmRangeResult struct {
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// Group list 1.
	GroupList1 string `pulumi:"groupList1"`
	// Group list 2.
	GroupList2 string `pulumi:"groupList2"`
	// Group list 3.
	GroupList3 string `pulumi:"groupList3"`
	// Group list 4.
	GroupList4 string `pulumi:"groupList4"`
	// The distinguished name of the object.
	Id string `pulumi:"id"`
	// Prefix list name.
	PrefixList string `pulumi:"prefixList"`
	// Route map name.
	RouteMap string `pulumi:"routeMap"`
	// Exclude standard SSM range (232.0.0.0/8).
	SsmNone bool `pulumi:"ssmNone"`
	// VRF name.
	VrfName string `pulumi:"vrfName"`
}

func LookupPimSsmRangeOutput(ctx *pulumi.Context, args LookupPimSsmRangeOutputArgs, opts ...pulumi.InvokeOption) LookupPimSsmRangeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPimSsmRangeResult, error) {
			args := v.(LookupPimSsmRangeArgs)
			r, err := LookupPimSsmRange(ctx, &args, opts...)
			var s LookupPimSsmRangeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPimSsmRangeResultOutput)
}

// A collection of arguments for invoking getPimSsmRange.
type LookupPimSsmRangeOutputArgs struct {
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput `pulumi:"device"`
	// VRF name.
	VrfName pulumi.StringInput `pulumi:"vrfName"`
}

func (LookupPimSsmRangeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPimSsmRangeArgs)(nil)).Elem()
}

// A collection of values returned by getPimSsmRange.
type LookupPimSsmRangeResultOutput struct{ *pulumi.OutputState }

func (LookupPimSsmRangeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPimSsmRangeResult)(nil)).Elem()
}

func (o LookupPimSsmRangeResultOutput) ToLookupPimSsmRangeResultOutput() LookupPimSsmRangeResultOutput {
	return o
}

func (o LookupPimSsmRangeResultOutput) ToLookupPimSsmRangeResultOutputWithContext(ctx context.Context) LookupPimSsmRangeResultOutput {
	return o
}

func (o LookupPimSsmRangeResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupPimSsmRangeResult] {
	return pulumix.Output[LookupPimSsmRangeResult]{
		OutputState: o.OutputState,
	}
}

// A device name from the provider configuration.
func (o LookupPimSsmRangeResultOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPimSsmRangeResult) *string { return v.Device }).(pulumi.StringPtrOutput)
}

// Group list 1.
func (o LookupPimSsmRangeResultOutput) GroupList1() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPimSsmRangeResult) string { return v.GroupList1 }).(pulumi.StringOutput)
}

// Group list 2.
func (o LookupPimSsmRangeResultOutput) GroupList2() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPimSsmRangeResult) string { return v.GroupList2 }).(pulumi.StringOutput)
}

// Group list 3.
func (o LookupPimSsmRangeResultOutput) GroupList3() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPimSsmRangeResult) string { return v.GroupList3 }).(pulumi.StringOutput)
}

// Group list 4.
func (o LookupPimSsmRangeResultOutput) GroupList4() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPimSsmRangeResult) string { return v.GroupList4 }).(pulumi.StringOutput)
}

// The distinguished name of the object.
func (o LookupPimSsmRangeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPimSsmRangeResult) string { return v.Id }).(pulumi.StringOutput)
}

// Prefix list name.
func (o LookupPimSsmRangeResultOutput) PrefixList() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPimSsmRangeResult) string { return v.PrefixList }).(pulumi.StringOutput)
}

// Route map name.
func (o LookupPimSsmRangeResultOutput) RouteMap() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPimSsmRangeResult) string { return v.RouteMap }).(pulumi.StringOutput)
}

// Exclude standard SSM range (232.0.0.0/8).
func (o LookupPimSsmRangeResultOutput) SsmNone() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupPimSsmRangeResult) bool { return v.SsmNone }).(pulumi.BoolOutput)
}

// VRF name.
func (o LookupPimSsmRangeResultOutput) VrfName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPimSsmRangeResult) string { return v.VrfName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPimSsmRangeResultOutput{})
}