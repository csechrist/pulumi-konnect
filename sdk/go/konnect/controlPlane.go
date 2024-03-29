// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package konnect

import (
	"context"
	"reflect"

	"github.com/csechrist/pulumi-konnect/sdk/go/konnect/internal"
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
//	"github.com/csechrist/pulumi-konnect/sdk/go/konnect"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := konnect.NewControlPlane(ctx, "example", &konnect.ControlPlaneArgs{
//				Description: pulumi.String("TestControlPlane"),
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
// Control planes can be imported using a proper value of `id` as described above
type ControlPlane struct {
	pulumi.CustomResourceState

	// **(String)** The cluster type of the control plane.
	ClusterType pulumi.StringOutput `pulumi:"clusterType"`
	// **(String)** The control plane endpoint URL of the control plane.
	ControlPlaneEndpoint pulumi.StringOutput `pulumi:"controlPlaneEndpoint"`
	// **(Optional, String)** The description of the control plane.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// **(Required, String)** The name of the control plane.
	Name pulumi.StringOutput `pulumi:"name"`
	// **(String)** The telemetry endpoint URL of the control plane.
	TelemetryEndpoint pulumi.StringOutput `pulumi:"telemetryEndpoint"`
}

// NewControlPlane registers a new resource with the given unique name, arguments, and options.
func NewControlPlane(ctx *pulumi.Context,
	name string, args *ControlPlaneArgs, opts ...pulumi.ResourceOption) (*ControlPlane, error) {
	if args == nil {
		args = &ControlPlaneArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ControlPlane
	err := ctx.RegisterResource("konnect:index/controlPlane:ControlPlane", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetControlPlane gets an existing ControlPlane resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetControlPlane(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ControlPlaneState, opts ...pulumi.ResourceOption) (*ControlPlane, error) {
	var resource ControlPlane
	err := ctx.ReadResource("konnect:index/controlPlane:ControlPlane", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ControlPlane resources.
type controlPlaneState struct {
	// **(String)** The cluster type of the control plane.
	ClusterType *string `pulumi:"clusterType"`
	// **(String)** The control plane endpoint URL of the control plane.
	ControlPlaneEndpoint *string `pulumi:"controlPlaneEndpoint"`
	// **(Optional, String)** The description of the control plane.
	Description *string `pulumi:"description"`
	// **(Required, String)** The name of the control plane.
	Name *string `pulumi:"name"`
	// **(String)** The telemetry endpoint URL of the control plane.
	TelemetryEndpoint *string `pulumi:"telemetryEndpoint"`
}

type ControlPlaneState struct {
	// **(String)** The cluster type of the control plane.
	ClusterType pulumi.StringPtrInput
	// **(String)** The control plane endpoint URL of the control plane.
	ControlPlaneEndpoint pulumi.StringPtrInput
	// **(Optional, String)** The description of the control plane.
	Description pulumi.StringPtrInput
	// **(Required, String)** The name of the control plane.
	Name pulumi.StringPtrInput
	// **(String)** The telemetry endpoint URL of the control plane.
	TelemetryEndpoint pulumi.StringPtrInput
}

func (ControlPlaneState) ElementType() reflect.Type {
	return reflect.TypeOf((*controlPlaneState)(nil)).Elem()
}

type controlPlaneArgs struct {
	// **(Optional, String)** The description of the control plane.
	Description *string `pulumi:"description"`
	// **(Required, String)** The name of the control plane.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a ControlPlane resource.
type ControlPlaneArgs struct {
	// **(Optional, String)** The description of the control plane.
	Description pulumi.StringPtrInput
	// **(Required, String)** The name of the control plane.
	Name pulumi.StringPtrInput
}

func (ControlPlaneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*controlPlaneArgs)(nil)).Elem()
}

type ControlPlaneInput interface {
	pulumi.Input

	ToControlPlaneOutput() ControlPlaneOutput
	ToControlPlaneOutputWithContext(ctx context.Context) ControlPlaneOutput
}

func (*ControlPlane) ElementType() reflect.Type {
	return reflect.TypeOf((**ControlPlane)(nil)).Elem()
}

func (i *ControlPlane) ToControlPlaneOutput() ControlPlaneOutput {
	return i.ToControlPlaneOutputWithContext(context.Background())
}

func (i *ControlPlane) ToControlPlaneOutputWithContext(ctx context.Context) ControlPlaneOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ControlPlaneOutput)
}

// ControlPlaneArrayInput is an input type that accepts ControlPlaneArray and ControlPlaneArrayOutput values.
// You can construct a concrete instance of `ControlPlaneArrayInput` via:
//
//	ControlPlaneArray{ ControlPlaneArgs{...} }
type ControlPlaneArrayInput interface {
	pulumi.Input

	ToControlPlaneArrayOutput() ControlPlaneArrayOutput
	ToControlPlaneArrayOutputWithContext(context.Context) ControlPlaneArrayOutput
}

type ControlPlaneArray []ControlPlaneInput

func (ControlPlaneArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ControlPlane)(nil)).Elem()
}

func (i ControlPlaneArray) ToControlPlaneArrayOutput() ControlPlaneArrayOutput {
	return i.ToControlPlaneArrayOutputWithContext(context.Background())
}

func (i ControlPlaneArray) ToControlPlaneArrayOutputWithContext(ctx context.Context) ControlPlaneArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ControlPlaneArrayOutput)
}

// ControlPlaneMapInput is an input type that accepts ControlPlaneMap and ControlPlaneMapOutput values.
// You can construct a concrete instance of `ControlPlaneMapInput` via:
//
//	ControlPlaneMap{ "key": ControlPlaneArgs{...} }
type ControlPlaneMapInput interface {
	pulumi.Input

	ToControlPlaneMapOutput() ControlPlaneMapOutput
	ToControlPlaneMapOutputWithContext(context.Context) ControlPlaneMapOutput
}

type ControlPlaneMap map[string]ControlPlaneInput

func (ControlPlaneMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ControlPlane)(nil)).Elem()
}

func (i ControlPlaneMap) ToControlPlaneMapOutput() ControlPlaneMapOutput {
	return i.ToControlPlaneMapOutputWithContext(context.Background())
}

func (i ControlPlaneMap) ToControlPlaneMapOutputWithContext(ctx context.Context) ControlPlaneMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ControlPlaneMapOutput)
}

type ControlPlaneOutput struct{ *pulumi.OutputState }

func (ControlPlaneOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ControlPlane)(nil)).Elem()
}

func (o ControlPlaneOutput) ToControlPlaneOutput() ControlPlaneOutput {
	return o
}

func (o ControlPlaneOutput) ToControlPlaneOutputWithContext(ctx context.Context) ControlPlaneOutput {
	return o
}

// **(String)** The cluster type of the control plane.
func (o ControlPlaneOutput) ClusterType() pulumi.StringOutput {
	return o.ApplyT(func(v *ControlPlane) pulumi.StringOutput { return v.ClusterType }).(pulumi.StringOutput)
}

// **(String)** The control plane endpoint URL of the control plane.
func (o ControlPlaneOutput) ControlPlaneEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *ControlPlane) pulumi.StringOutput { return v.ControlPlaneEndpoint }).(pulumi.StringOutput)
}

// **(Optional, String)** The description of the control plane.
func (o ControlPlaneOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ControlPlane) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// **(Required, String)** The name of the control plane.
func (o ControlPlaneOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ControlPlane) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// **(String)** The telemetry endpoint URL of the control plane.
func (o ControlPlaneOutput) TelemetryEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *ControlPlane) pulumi.StringOutput { return v.TelemetryEndpoint }).(pulumi.StringOutput)
}

type ControlPlaneArrayOutput struct{ *pulumi.OutputState }

func (ControlPlaneArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ControlPlane)(nil)).Elem()
}

func (o ControlPlaneArrayOutput) ToControlPlaneArrayOutput() ControlPlaneArrayOutput {
	return o
}

func (o ControlPlaneArrayOutput) ToControlPlaneArrayOutputWithContext(ctx context.Context) ControlPlaneArrayOutput {
	return o
}

func (o ControlPlaneArrayOutput) Index(i pulumi.IntInput) ControlPlaneOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ControlPlane {
		return vs[0].([]*ControlPlane)[vs[1].(int)]
	}).(ControlPlaneOutput)
}

type ControlPlaneMapOutput struct{ *pulumi.OutputState }

func (ControlPlaneMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ControlPlane)(nil)).Elem()
}

func (o ControlPlaneMapOutput) ToControlPlaneMapOutput() ControlPlaneMapOutput {
	return o
}

func (o ControlPlaneMapOutput) ToControlPlaneMapOutputWithContext(ctx context.Context) ControlPlaneMapOutput {
	return o
}

func (o ControlPlaneMapOutput) MapIndex(k pulumi.StringInput) ControlPlaneOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ControlPlane {
		return vs[0].(map[string]*ControlPlane)[vs[1].(string)]
	}).(ControlPlaneOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ControlPlaneInput)(nil)).Elem(), &ControlPlane{})
	pulumi.RegisterInputType(reflect.TypeOf((*ControlPlaneArrayInput)(nil)).Elem(), ControlPlaneArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ControlPlaneMapInput)(nil)).Elem(), ControlPlaneMap{})
	pulumi.RegisterOutputType(ControlPlaneOutput{})
	pulumi.RegisterOutputType(ControlPlaneArrayOutput{})
	pulumi.RegisterOutputType(ControlPlaneMapOutput{})
}
