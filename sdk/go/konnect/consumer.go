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

// Represents a consumer within a control plane
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
//			_, err = konnect.NewConsumer(ctx, "example", &konnect.ConsumerArgs{
//				ControlPlaneId: *pulumi.String(controlPlane.Id),
//				Username:       pulumi.String("testuser"),
//				CustomId:       pulumi.String("testcustom"),
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
// Consumers can be imported using a proper value of `id` as described above
type Consumer struct {
	pulumi.CustomResourceState

	// **(String)** Id of the consumer alone
	ConsumerId pulumi.StringOutput `pulumi:"consumerId"`
	// **(Required, String)** The id of the control plane.
	ControlPlaneId pulumi.StringOutput `pulumi:"controlPlaneId"`
	// **(Optional, String)** Field for storing an existing unique ID for the Consumer.
	CustomId pulumi.StringPtrOutput `pulumi:"customId"`
	// **(Optional, String)** The unique username of the Consumer.
	Username pulumi.StringPtrOutput `pulumi:"username"`
}

// NewConsumer registers a new resource with the given unique name, arguments, and options.
func NewConsumer(ctx *pulumi.Context,
	name string, args *ConsumerArgs, opts ...pulumi.ResourceOption) (*Consumer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ControlPlaneId == nil {
		return nil, errors.New("invalid value for required argument 'ControlPlaneId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Consumer
	err := ctx.RegisterResource("konnect:index/consumer:Consumer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConsumer gets an existing Consumer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConsumer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConsumerState, opts ...pulumi.ResourceOption) (*Consumer, error) {
	var resource Consumer
	err := ctx.ReadResource("konnect:index/consumer:Consumer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Consumer resources.
type consumerState struct {
	// **(String)** Id of the consumer alone
	ConsumerId *string `pulumi:"consumerId"`
	// **(Required, String)** The id of the control plane.
	ControlPlaneId *string `pulumi:"controlPlaneId"`
	// **(Optional, String)** Field for storing an existing unique ID for the Consumer.
	CustomId *string `pulumi:"customId"`
	// **(Optional, String)** The unique username of the Consumer.
	Username *string `pulumi:"username"`
}

type ConsumerState struct {
	// **(String)** Id of the consumer alone
	ConsumerId pulumi.StringPtrInput
	// **(Required, String)** The id of the control plane.
	ControlPlaneId pulumi.StringPtrInput
	// **(Optional, String)** Field for storing an existing unique ID for the Consumer.
	CustomId pulumi.StringPtrInput
	// **(Optional, String)** The unique username of the Consumer.
	Username pulumi.StringPtrInput
}

func (ConsumerState) ElementType() reflect.Type {
	return reflect.TypeOf((*consumerState)(nil)).Elem()
}

type consumerArgs struct {
	// **(Required, String)** The id of the control plane.
	ControlPlaneId string `pulumi:"controlPlaneId"`
	// **(Optional, String)** Field for storing an existing unique ID for the Consumer.
	CustomId *string `pulumi:"customId"`
	// **(Optional, String)** The unique username of the Consumer.
	Username *string `pulumi:"username"`
}

// The set of arguments for constructing a Consumer resource.
type ConsumerArgs struct {
	// **(Required, String)** The id of the control plane.
	ControlPlaneId pulumi.StringInput
	// **(Optional, String)** Field for storing an existing unique ID for the Consumer.
	CustomId pulumi.StringPtrInput
	// **(Optional, String)** The unique username of the Consumer.
	Username pulumi.StringPtrInput
}

func (ConsumerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*consumerArgs)(nil)).Elem()
}

type ConsumerInput interface {
	pulumi.Input

	ToConsumerOutput() ConsumerOutput
	ToConsumerOutputWithContext(ctx context.Context) ConsumerOutput
}

func (*Consumer) ElementType() reflect.Type {
	return reflect.TypeOf((**Consumer)(nil)).Elem()
}

func (i *Consumer) ToConsumerOutput() ConsumerOutput {
	return i.ToConsumerOutputWithContext(context.Background())
}

func (i *Consumer) ToConsumerOutputWithContext(ctx context.Context) ConsumerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsumerOutput)
}

// ConsumerArrayInput is an input type that accepts ConsumerArray and ConsumerArrayOutput values.
// You can construct a concrete instance of `ConsumerArrayInput` via:
//
//	ConsumerArray{ ConsumerArgs{...} }
type ConsumerArrayInput interface {
	pulumi.Input

	ToConsumerArrayOutput() ConsumerArrayOutput
	ToConsumerArrayOutputWithContext(context.Context) ConsumerArrayOutput
}

type ConsumerArray []ConsumerInput

func (ConsumerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Consumer)(nil)).Elem()
}

func (i ConsumerArray) ToConsumerArrayOutput() ConsumerArrayOutput {
	return i.ToConsumerArrayOutputWithContext(context.Background())
}

func (i ConsumerArray) ToConsumerArrayOutputWithContext(ctx context.Context) ConsumerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsumerArrayOutput)
}

// ConsumerMapInput is an input type that accepts ConsumerMap and ConsumerMapOutput values.
// You can construct a concrete instance of `ConsumerMapInput` via:
//
//	ConsumerMap{ "key": ConsumerArgs{...} }
type ConsumerMapInput interface {
	pulumi.Input

	ToConsumerMapOutput() ConsumerMapOutput
	ToConsumerMapOutputWithContext(context.Context) ConsumerMapOutput
}

type ConsumerMap map[string]ConsumerInput

func (ConsumerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Consumer)(nil)).Elem()
}

func (i ConsumerMap) ToConsumerMapOutput() ConsumerMapOutput {
	return i.ToConsumerMapOutputWithContext(context.Background())
}

func (i ConsumerMap) ToConsumerMapOutputWithContext(ctx context.Context) ConsumerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsumerMapOutput)
}

type ConsumerOutput struct{ *pulumi.OutputState }

func (ConsumerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Consumer)(nil)).Elem()
}

func (o ConsumerOutput) ToConsumerOutput() ConsumerOutput {
	return o
}

func (o ConsumerOutput) ToConsumerOutputWithContext(ctx context.Context) ConsumerOutput {
	return o
}

// **(String)** Id of the consumer alone
func (o ConsumerOutput) ConsumerId() pulumi.StringOutput {
	return o.ApplyT(func(v *Consumer) pulumi.StringOutput { return v.ConsumerId }).(pulumi.StringOutput)
}

// **(Required, String)** The id of the control plane.
func (o ConsumerOutput) ControlPlaneId() pulumi.StringOutput {
	return o.ApplyT(func(v *Consumer) pulumi.StringOutput { return v.ControlPlaneId }).(pulumi.StringOutput)
}

// **(Optional, String)** Field for storing an existing unique ID for the Consumer.
func (o ConsumerOutput) CustomId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Consumer) pulumi.StringPtrOutput { return v.CustomId }).(pulumi.StringPtrOutput)
}

// **(Optional, String)** The unique username of the Consumer.
func (o ConsumerOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Consumer) pulumi.StringPtrOutput { return v.Username }).(pulumi.StringPtrOutput)
}

type ConsumerArrayOutput struct{ *pulumi.OutputState }

func (ConsumerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Consumer)(nil)).Elem()
}

func (o ConsumerArrayOutput) ToConsumerArrayOutput() ConsumerArrayOutput {
	return o
}

func (o ConsumerArrayOutput) ToConsumerArrayOutputWithContext(ctx context.Context) ConsumerArrayOutput {
	return o
}

func (o ConsumerArrayOutput) Index(i pulumi.IntInput) ConsumerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Consumer {
		return vs[0].([]*Consumer)[vs[1].(int)]
	}).(ConsumerOutput)
}

type ConsumerMapOutput struct{ *pulumi.OutputState }

func (ConsumerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Consumer)(nil)).Elem()
}

func (o ConsumerMapOutput) ToConsumerMapOutput() ConsumerMapOutput {
	return o
}

func (o ConsumerMapOutput) ToConsumerMapOutputWithContext(ctx context.Context) ConsumerMapOutput {
	return o
}

func (o ConsumerMapOutput) MapIndex(k pulumi.StringInput) ConsumerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Consumer {
		return vs[0].(map[string]*Consumer)[vs[1].(string)]
	}).(ConsumerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConsumerInput)(nil)).Elem(), &Consumer{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConsumerArrayInput)(nil)).Elem(), ConsumerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConsumerMapInput)(nil)).Elem(), ConsumerMap{})
	pulumi.RegisterOutputType(ConsumerOutput{})
	pulumi.RegisterOutputType(ConsumerArrayOutput{})
	pulumi.RegisterOutputType(ConsumerMapOutput{})
}
