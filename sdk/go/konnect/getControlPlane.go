// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package konnect

import (
	"context"
	"reflect"

	"github.com/csechrist123/pulumi-konnect/sdk/go/konnect/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a control plane
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/csechrist123/pulumi-konnect/sdk/go/konnect"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := konnect.LookupControlPlane(ctx, &konnect.LookupControlPlaneArgs{
//				Name: pulumi.StringRef("TestControlPlane"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupControlPlane(ctx *pulumi.Context, args *LookupControlPlaneArgs, opts ...pulumi.InvokeOption) (*LookupControlPlaneResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupControlPlaneResult
	err := ctx.Invoke("konnect:index/getControlPlane:getControlPlane", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getControlPlane.
type LookupControlPlaneArgs struct {
	// **(Optional, String)** The filter string to apply to the name of the control plane. Uses equality.
	Name *string `pulumi:"name"`
	// **(Optional, String)** The search string to apply to the name of the control plane. Uses contains.
	SearchName *string `pulumi:"searchName"`
}

// A collection of values returned by getControlPlane.
type LookupControlPlaneResult struct {
	// **(String)** The cluster type of the control plane.
	ClusterType string `pulumi:"clusterType"`
	// **(String)** The control plane endpoint URL of the control plane.
	ControlPlaneEndpoint string `pulumi:"controlPlaneEndpoint"`
	// **(String)** The description of the control plane.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	Name       *string `pulumi:"name"`
	SearchName *string `pulumi:"searchName"`
	// **(String)** The telemetry endpoint URL of the control plane.
	TelemetryEndpoint string `pulumi:"telemetryEndpoint"`
}

func LookupControlPlaneOutput(ctx *pulumi.Context, args LookupControlPlaneOutputArgs, opts ...pulumi.InvokeOption) LookupControlPlaneResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupControlPlaneResult, error) {
			args := v.(LookupControlPlaneArgs)
			r, err := LookupControlPlane(ctx, &args, opts...)
			var s LookupControlPlaneResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupControlPlaneResultOutput)
}

// A collection of arguments for invoking getControlPlane.
type LookupControlPlaneOutputArgs struct {
	// **(Optional, String)** The filter string to apply to the name of the control plane. Uses equality.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// **(Optional, String)** The search string to apply to the name of the control plane. Uses contains.
	SearchName pulumi.StringPtrInput `pulumi:"searchName"`
}

func (LookupControlPlaneOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupControlPlaneArgs)(nil)).Elem()
}

// A collection of values returned by getControlPlane.
type LookupControlPlaneResultOutput struct{ *pulumi.OutputState }

func (LookupControlPlaneResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupControlPlaneResult)(nil)).Elem()
}

func (o LookupControlPlaneResultOutput) ToLookupControlPlaneResultOutput() LookupControlPlaneResultOutput {
	return o
}

func (o LookupControlPlaneResultOutput) ToLookupControlPlaneResultOutputWithContext(ctx context.Context) LookupControlPlaneResultOutput {
	return o
}

// **(String)** The cluster type of the control plane.
func (o LookupControlPlaneResultOutput) ClusterType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupControlPlaneResult) string { return v.ClusterType }).(pulumi.StringOutput)
}

// **(String)** The control plane endpoint URL of the control plane.
func (o LookupControlPlaneResultOutput) ControlPlaneEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupControlPlaneResult) string { return v.ControlPlaneEndpoint }).(pulumi.StringOutput)
}

// **(String)** The description of the control plane.
func (o LookupControlPlaneResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupControlPlaneResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupControlPlaneResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupControlPlaneResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupControlPlaneResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupControlPlaneResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupControlPlaneResultOutput) SearchName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupControlPlaneResult) *string { return v.SearchName }).(pulumi.StringPtrOutput)
}

// **(String)** The telemetry endpoint URL of the control plane.
func (o LookupControlPlaneResultOutput) TelemetryEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupControlPlaneResult) string { return v.TelemetryEndpoint }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupControlPlaneResultOutput{})
}