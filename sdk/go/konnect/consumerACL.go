// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package konnect

import (
	"context"
	"reflect"

	"errors"
	"github.com/csechrist123/pulumi-konnect/sdk/go/konnect/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents an ACL credential for a consumer within a control plane
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
//			controlPlane, err := konnect.LookupControlPlane(ctx, &konnect.LookupControlPlaneArgs{
//				Name: pulumi.StringRef("TestControlPlane"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			consumer, err := konnect.LookupConsumer(ctx, &konnect.LookupConsumerArgs{
//				ControlPlaneId: controlPlane.Id,
//				SearchUsername: pulumi.StringRef("Bob"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = konnect.NewConsumerACL(ctx, "example", &konnect.ConsumerACLArgs{
//				ControlPlaneId: *pulumi.String(controlPlane.Id),
//				ConsumerId:     *pulumi.String(consumer.ConsumerId),
//				Group:          pulumi.String("my-acl-group"),
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
// Consumer ACLs can be imported using a proper value of `id` as described above
type ConsumerACL struct {
	pulumi.CustomResourceState

	// **(String)** Id of the consumer ACL alone
	AclId pulumi.StringOutput `pulumi:"aclId"`
	// **(Required, String)** The id of the consumer.
	ConsumerId pulumi.StringOutput `pulumi:"consumerId"`
	// **(Required, String)** The id of the control plane.
	ControlPlaneId pulumi.StringOutput `pulumi:"controlPlaneId"`
	// **(Required, String)** The ACL group value.
	Group pulumi.StringOutput `pulumi:"group"`
}

// NewConsumerACL registers a new resource with the given unique name, arguments, and options.
func NewConsumerACL(ctx *pulumi.Context,
	name string, args *ConsumerACLArgs, opts ...pulumi.ResourceOption) (*ConsumerACL, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConsumerId == nil {
		return nil, errors.New("invalid value for required argument 'ConsumerId'")
	}
	if args.ControlPlaneId == nil {
		return nil, errors.New("invalid value for required argument 'ControlPlaneId'")
	}
	if args.Group == nil {
		return nil, errors.New("invalid value for required argument 'Group'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ConsumerACL
	err := ctx.RegisterResource("konnect:index/consumerACL:ConsumerACL", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConsumerACL gets an existing ConsumerACL resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConsumerACL(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConsumerACLState, opts ...pulumi.ResourceOption) (*ConsumerACL, error) {
	var resource ConsumerACL
	err := ctx.ReadResource("konnect:index/consumerACL:ConsumerACL", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConsumerACL resources.
type consumerACLState struct {
	// **(String)** Id of the consumer ACL alone
	AclId *string `pulumi:"aclId"`
	// **(Required, String)** The id of the consumer.
	ConsumerId *string `pulumi:"consumerId"`
	// **(Required, String)** The id of the control plane.
	ControlPlaneId *string `pulumi:"controlPlaneId"`
	// **(Required, String)** The ACL group value.
	Group *string `pulumi:"group"`
}

type ConsumerACLState struct {
	// **(String)** Id of the consumer ACL alone
	AclId pulumi.StringPtrInput
	// **(Required, String)** The id of the consumer.
	ConsumerId pulumi.StringPtrInput
	// **(Required, String)** The id of the control plane.
	ControlPlaneId pulumi.StringPtrInput
	// **(Required, String)** The ACL group value.
	Group pulumi.StringPtrInput
}

func (ConsumerACLState) ElementType() reflect.Type {
	return reflect.TypeOf((*consumerACLState)(nil)).Elem()
}

type consumerACLArgs struct {
	// **(Required, String)** The id of the consumer.
	ConsumerId string `pulumi:"consumerId"`
	// **(Required, String)** The id of the control plane.
	ControlPlaneId string `pulumi:"controlPlaneId"`
	// **(Required, String)** The ACL group value.
	Group string `pulumi:"group"`
}

// The set of arguments for constructing a ConsumerACL resource.
type ConsumerACLArgs struct {
	// **(Required, String)** The id of the consumer.
	ConsumerId pulumi.StringInput
	// **(Required, String)** The id of the control plane.
	ControlPlaneId pulumi.StringInput
	// **(Required, String)** The ACL group value.
	Group pulumi.StringInput
}

func (ConsumerACLArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*consumerACLArgs)(nil)).Elem()
}

type ConsumerACLInput interface {
	pulumi.Input

	ToConsumerACLOutput() ConsumerACLOutput
	ToConsumerACLOutputWithContext(ctx context.Context) ConsumerACLOutput
}

func (*ConsumerACL) ElementType() reflect.Type {
	return reflect.TypeOf((**ConsumerACL)(nil)).Elem()
}

func (i *ConsumerACL) ToConsumerACLOutput() ConsumerACLOutput {
	return i.ToConsumerACLOutputWithContext(context.Background())
}

func (i *ConsumerACL) ToConsumerACLOutputWithContext(ctx context.Context) ConsumerACLOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsumerACLOutput)
}

// ConsumerACLArrayInput is an input type that accepts ConsumerACLArray and ConsumerACLArrayOutput values.
// You can construct a concrete instance of `ConsumerACLArrayInput` via:
//
//	ConsumerACLArray{ ConsumerACLArgs{...} }
type ConsumerACLArrayInput interface {
	pulumi.Input

	ToConsumerACLArrayOutput() ConsumerACLArrayOutput
	ToConsumerACLArrayOutputWithContext(context.Context) ConsumerACLArrayOutput
}

type ConsumerACLArray []ConsumerACLInput

func (ConsumerACLArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConsumerACL)(nil)).Elem()
}

func (i ConsumerACLArray) ToConsumerACLArrayOutput() ConsumerACLArrayOutput {
	return i.ToConsumerACLArrayOutputWithContext(context.Background())
}

func (i ConsumerACLArray) ToConsumerACLArrayOutputWithContext(ctx context.Context) ConsumerACLArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsumerACLArrayOutput)
}

// ConsumerACLMapInput is an input type that accepts ConsumerACLMap and ConsumerACLMapOutput values.
// You can construct a concrete instance of `ConsumerACLMapInput` via:
//
//	ConsumerACLMap{ "key": ConsumerACLArgs{...} }
type ConsumerACLMapInput interface {
	pulumi.Input

	ToConsumerACLMapOutput() ConsumerACLMapOutput
	ToConsumerACLMapOutputWithContext(context.Context) ConsumerACLMapOutput
}

type ConsumerACLMap map[string]ConsumerACLInput

func (ConsumerACLMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConsumerACL)(nil)).Elem()
}

func (i ConsumerACLMap) ToConsumerACLMapOutput() ConsumerACLMapOutput {
	return i.ToConsumerACLMapOutputWithContext(context.Background())
}

func (i ConsumerACLMap) ToConsumerACLMapOutputWithContext(ctx context.Context) ConsumerACLMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsumerACLMapOutput)
}

type ConsumerACLOutput struct{ *pulumi.OutputState }

func (ConsumerACLOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConsumerACL)(nil)).Elem()
}

func (o ConsumerACLOutput) ToConsumerACLOutput() ConsumerACLOutput {
	return o
}

func (o ConsumerACLOutput) ToConsumerACLOutputWithContext(ctx context.Context) ConsumerACLOutput {
	return o
}

// **(String)** Id of the consumer ACL alone
func (o ConsumerACLOutput) AclId() pulumi.StringOutput {
	return o.ApplyT(func(v *ConsumerACL) pulumi.StringOutput { return v.AclId }).(pulumi.StringOutput)
}

// **(Required, String)** The id of the consumer.
func (o ConsumerACLOutput) ConsumerId() pulumi.StringOutput {
	return o.ApplyT(func(v *ConsumerACL) pulumi.StringOutput { return v.ConsumerId }).(pulumi.StringOutput)
}

// **(Required, String)** The id of the control plane.
func (o ConsumerACLOutput) ControlPlaneId() pulumi.StringOutput {
	return o.ApplyT(func(v *ConsumerACL) pulumi.StringOutput { return v.ControlPlaneId }).(pulumi.StringOutput)
}

// **(Required, String)** The ACL group value.
func (o ConsumerACLOutput) Group() pulumi.StringOutput {
	return o.ApplyT(func(v *ConsumerACL) pulumi.StringOutput { return v.Group }).(pulumi.StringOutput)
}

type ConsumerACLArrayOutput struct{ *pulumi.OutputState }

func (ConsumerACLArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConsumerACL)(nil)).Elem()
}

func (o ConsumerACLArrayOutput) ToConsumerACLArrayOutput() ConsumerACLArrayOutput {
	return o
}

func (o ConsumerACLArrayOutput) ToConsumerACLArrayOutputWithContext(ctx context.Context) ConsumerACLArrayOutput {
	return o
}

func (o ConsumerACLArrayOutput) Index(i pulumi.IntInput) ConsumerACLOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ConsumerACL {
		return vs[0].([]*ConsumerACL)[vs[1].(int)]
	}).(ConsumerACLOutput)
}

type ConsumerACLMapOutput struct{ *pulumi.OutputState }

func (ConsumerACLMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConsumerACL)(nil)).Elem()
}

func (o ConsumerACLMapOutput) ToConsumerACLMapOutput() ConsumerACLMapOutput {
	return o
}

func (o ConsumerACLMapOutput) ToConsumerACLMapOutputWithContext(ctx context.Context) ConsumerACLMapOutput {
	return o
}

func (o ConsumerACLMapOutput) MapIndex(k pulumi.StringInput) ConsumerACLOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ConsumerACL {
		return vs[0].(map[string]*ConsumerACL)[vs[1].(string)]
	}).(ConsumerACLOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConsumerACLInput)(nil)).Elem(), &ConsumerACL{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConsumerACLArrayInput)(nil)).Elem(), ConsumerACLArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConsumerACLMapInput)(nil)).Elem(), ConsumerACLMap{})
	pulumi.RegisterOutputType(ConsumerACLOutput{})
	pulumi.RegisterOutputType(ConsumerACLArrayOutput{})
	pulumi.RegisterOutputType(ConsumerACLMapOutput{})
}
