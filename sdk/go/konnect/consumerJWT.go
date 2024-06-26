// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package konnect

import (
	"context"
	"reflect"

	"errors"
	"github.com/csechrist/pulumi-konnect/sdk/go/konnect/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a JWT credential for a consumer within a control plane
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
//			_, err = konnect.NewConsumerJWT(ctx, "example", &konnect.ConsumerJWTArgs{
//				ControlPlaneId: *pulumi.String(controlPlane.Id),
//				ConsumerId:     *pulumi.String(consumer.ConsumerId),
//				Key:            pulumi.String("my-key"),
//				Secret:         pulumi.String("my-secret"),
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
// Consumer JWTs can be imported using a proper value of `id` as described above
type ConsumerJWT struct {
	pulumi.CustomResourceState

	// **(Optional, String)** The algorithm for the JWT.  Allowed values: `HS256`, `HS384`, `HS512`, `RS256`, `RS384`, `RS512`, `ES256`, `ES384`. Default: `HS256`
	Algorithm pulumi.StringPtrOutput `pulumi:"algorithm"`
	// **(Required, String)** The id of the consumer.
	ConsumerId pulumi.StringOutput `pulumi:"consumerId"`
	// **(Required, String)** The id of the control plane.
	ControlPlaneId pulumi.StringOutput `pulumi:"controlPlaneId"`
	// **(String)** Id of the consumer JWT alone
	JwtId pulumi.StringOutput `pulumi:"jwtId"`
	// **(Optional/Computed, String)** The key value.  If left out, a key will be generated for you.
	Key pulumi.StringOutput `pulumi:"key"`
	// **(Optional, String)** The RSA public key in PEM format for the JWT.  Required if `algorithm` is `RS256`, `RS384`, `RS512`, `ES256`, or `ES384`.
	RsaPublicKey pulumi.StringPtrOutput `pulumi:"rsaPublicKey"`
	// **(Optional/Computed, String)** The secret value.  If left out, a key will be generated for you.
	Secret pulumi.StringOutput `pulumi:"secret"`
}

// NewConsumerJWT registers a new resource with the given unique name, arguments, and options.
func NewConsumerJWT(ctx *pulumi.Context,
	name string, args *ConsumerJWTArgs, opts ...pulumi.ResourceOption) (*ConsumerJWT, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConsumerId == nil {
		return nil, errors.New("invalid value for required argument 'ConsumerId'")
	}
	if args.ControlPlaneId == nil {
		return nil, errors.New("invalid value for required argument 'ControlPlaneId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ConsumerJWT
	err := ctx.RegisterResource("konnect:index/consumerJWT:ConsumerJWT", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConsumerJWT gets an existing ConsumerJWT resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConsumerJWT(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConsumerJWTState, opts ...pulumi.ResourceOption) (*ConsumerJWT, error) {
	var resource ConsumerJWT
	err := ctx.ReadResource("konnect:index/consumerJWT:ConsumerJWT", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConsumerJWT resources.
type consumerJWTState struct {
	// **(Optional, String)** The algorithm for the JWT.  Allowed values: `HS256`, `HS384`, `HS512`, `RS256`, `RS384`, `RS512`, `ES256`, `ES384`. Default: `HS256`
	Algorithm *string `pulumi:"algorithm"`
	// **(Required, String)** The id of the consumer.
	ConsumerId *string `pulumi:"consumerId"`
	// **(Required, String)** The id of the control plane.
	ControlPlaneId *string `pulumi:"controlPlaneId"`
	// **(String)** Id of the consumer JWT alone
	JwtId *string `pulumi:"jwtId"`
	// **(Optional/Computed, String)** The key value.  If left out, a key will be generated for you.
	Key *string `pulumi:"key"`
	// **(Optional, String)** The RSA public key in PEM format for the JWT.  Required if `algorithm` is `RS256`, `RS384`, `RS512`, `ES256`, or `ES384`.
	RsaPublicKey *string `pulumi:"rsaPublicKey"`
	// **(Optional/Computed, String)** The secret value.  If left out, a key will be generated for you.
	Secret *string `pulumi:"secret"`
}

type ConsumerJWTState struct {
	// **(Optional, String)** The algorithm for the JWT.  Allowed values: `HS256`, `HS384`, `HS512`, `RS256`, `RS384`, `RS512`, `ES256`, `ES384`. Default: `HS256`
	Algorithm pulumi.StringPtrInput
	// **(Required, String)** The id of the consumer.
	ConsumerId pulumi.StringPtrInput
	// **(Required, String)** The id of the control plane.
	ControlPlaneId pulumi.StringPtrInput
	// **(String)** Id of the consumer JWT alone
	JwtId pulumi.StringPtrInput
	// **(Optional/Computed, String)** The key value.  If left out, a key will be generated for you.
	Key pulumi.StringPtrInput
	// **(Optional, String)** The RSA public key in PEM format for the JWT.  Required if `algorithm` is `RS256`, `RS384`, `RS512`, `ES256`, or `ES384`.
	RsaPublicKey pulumi.StringPtrInput
	// **(Optional/Computed, String)** The secret value.  If left out, a key will be generated for you.
	Secret pulumi.StringPtrInput
}

func (ConsumerJWTState) ElementType() reflect.Type {
	return reflect.TypeOf((*consumerJWTState)(nil)).Elem()
}

type consumerJWTArgs struct {
	// **(Optional, String)** The algorithm for the JWT.  Allowed values: `HS256`, `HS384`, `HS512`, `RS256`, `RS384`, `RS512`, `ES256`, `ES384`. Default: `HS256`
	Algorithm *string `pulumi:"algorithm"`
	// **(Required, String)** The id of the consumer.
	ConsumerId string `pulumi:"consumerId"`
	// **(Required, String)** The id of the control plane.
	ControlPlaneId string `pulumi:"controlPlaneId"`
	// **(Optional/Computed, String)** The key value.  If left out, a key will be generated for you.
	Key *string `pulumi:"key"`
	// **(Optional, String)** The RSA public key in PEM format for the JWT.  Required if `algorithm` is `RS256`, `RS384`, `RS512`, `ES256`, or `ES384`.
	RsaPublicKey *string `pulumi:"rsaPublicKey"`
	// **(Optional/Computed, String)** The secret value.  If left out, a key will be generated for you.
	Secret *string `pulumi:"secret"`
}

// The set of arguments for constructing a ConsumerJWT resource.
type ConsumerJWTArgs struct {
	// **(Optional, String)** The algorithm for the JWT.  Allowed values: `HS256`, `HS384`, `HS512`, `RS256`, `RS384`, `RS512`, `ES256`, `ES384`. Default: `HS256`
	Algorithm pulumi.StringPtrInput
	// **(Required, String)** The id of the consumer.
	ConsumerId pulumi.StringInput
	// **(Required, String)** The id of the control plane.
	ControlPlaneId pulumi.StringInput
	// **(Optional/Computed, String)** The key value.  If left out, a key will be generated for you.
	Key pulumi.StringPtrInput
	// **(Optional, String)** The RSA public key in PEM format for the JWT.  Required if `algorithm` is `RS256`, `RS384`, `RS512`, `ES256`, or `ES384`.
	RsaPublicKey pulumi.StringPtrInput
	// **(Optional/Computed, String)** The secret value.  If left out, a key will be generated for you.
	Secret pulumi.StringPtrInput
}

func (ConsumerJWTArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*consumerJWTArgs)(nil)).Elem()
}

type ConsumerJWTInput interface {
	pulumi.Input

	ToConsumerJWTOutput() ConsumerJWTOutput
	ToConsumerJWTOutputWithContext(ctx context.Context) ConsumerJWTOutput
}

func (*ConsumerJWT) ElementType() reflect.Type {
	return reflect.TypeOf((**ConsumerJWT)(nil)).Elem()
}

func (i *ConsumerJWT) ToConsumerJWTOutput() ConsumerJWTOutput {
	return i.ToConsumerJWTOutputWithContext(context.Background())
}

func (i *ConsumerJWT) ToConsumerJWTOutputWithContext(ctx context.Context) ConsumerJWTOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsumerJWTOutput)
}

// ConsumerJWTArrayInput is an input type that accepts ConsumerJWTArray and ConsumerJWTArrayOutput values.
// You can construct a concrete instance of `ConsumerJWTArrayInput` via:
//
//	ConsumerJWTArray{ ConsumerJWTArgs{...} }
type ConsumerJWTArrayInput interface {
	pulumi.Input

	ToConsumerJWTArrayOutput() ConsumerJWTArrayOutput
	ToConsumerJWTArrayOutputWithContext(context.Context) ConsumerJWTArrayOutput
}

type ConsumerJWTArray []ConsumerJWTInput

func (ConsumerJWTArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConsumerJWT)(nil)).Elem()
}

func (i ConsumerJWTArray) ToConsumerJWTArrayOutput() ConsumerJWTArrayOutput {
	return i.ToConsumerJWTArrayOutputWithContext(context.Background())
}

func (i ConsumerJWTArray) ToConsumerJWTArrayOutputWithContext(ctx context.Context) ConsumerJWTArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsumerJWTArrayOutput)
}

// ConsumerJWTMapInput is an input type that accepts ConsumerJWTMap and ConsumerJWTMapOutput values.
// You can construct a concrete instance of `ConsumerJWTMapInput` via:
//
//	ConsumerJWTMap{ "key": ConsumerJWTArgs{...} }
type ConsumerJWTMapInput interface {
	pulumi.Input

	ToConsumerJWTMapOutput() ConsumerJWTMapOutput
	ToConsumerJWTMapOutputWithContext(context.Context) ConsumerJWTMapOutput
}

type ConsumerJWTMap map[string]ConsumerJWTInput

func (ConsumerJWTMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConsumerJWT)(nil)).Elem()
}

func (i ConsumerJWTMap) ToConsumerJWTMapOutput() ConsumerJWTMapOutput {
	return i.ToConsumerJWTMapOutputWithContext(context.Background())
}

func (i ConsumerJWTMap) ToConsumerJWTMapOutputWithContext(ctx context.Context) ConsumerJWTMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsumerJWTMapOutput)
}

type ConsumerJWTOutput struct{ *pulumi.OutputState }

func (ConsumerJWTOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConsumerJWT)(nil)).Elem()
}

func (o ConsumerJWTOutput) ToConsumerJWTOutput() ConsumerJWTOutput {
	return o
}

func (o ConsumerJWTOutput) ToConsumerJWTOutputWithContext(ctx context.Context) ConsumerJWTOutput {
	return o
}

// **(Optional, String)** The algorithm for the JWT.  Allowed values: `HS256`, `HS384`, `HS512`, `RS256`, `RS384`, `RS512`, `ES256`, `ES384`. Default: `HS256`
func (o ConsumerJWTOutput) Algorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConsumerJWT) pulumi.StringPtrOutput { return v.Algorithm }).(pulumi.StringPtrOutput)
}

// **(Required, String)** The id of the consumer.
func (o ConsumerJWTOutput) ConsumerId() pulumi.StringOutput {
	return o.ApplyT(func(v *ConsumerJWT) pulumi.StringOutput { return v.ConsumerId }).(pulumi.StringOutput)
}

// **(Required, String)** The id of the control plane.
func (o ConsumerJWTOutput) ControlPlaneId() pulumi.StringOutput {
	return o.ApplyT(func(v *ConsumerJWT) pulumi.StringOutput { return v.ControlPlaneId }).(pulumi.StringOutput)
}

// **(String)** Id of the consumer JWT alone
func (o ConsumerJWTOutput) JwtId() pulumi.StringOutput {
	return o.ApplyT(func(v *ConsumerJWT) pulumi.StringOutput { return v.JwtId }).(pulumi.StringOutput)
}

// **(Optional/Computed, String)** The key value.  If left out, a key will be generated for you.
func (o ConsumerJWTOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *ConsumerJWT) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// **(Optional, String)** The RSA public key in PEM format for the JWT.  Required if `algorithm` is `RS256`, `RS384`, `RS512`, `ES256`, or `ES384`.
func (o ConsumerJWTOutput) RsaPublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConsumerJWT) pulumi.StringPtrOutput { return v.RsaPublicKey }).(pulumi.StringPtrOutput)
}

// **(Optional/Computed, String)** The secret value.  If left out, a key will be generated for you.
func (o ConsumerJWTOutput) Secret() pulumi.StringOutput {
	return o.ApplyT(func(v *ConsumerJWT) pulumi.StringOutput { return v.Secret }).(pulumi.StringOutput)
}

type ConsumerJWTArrayOutput struct{ *pulumi.OutputState }

func (ConsumerJWTArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConsumerJWT)(nil)).Elem()
}

func (o ConsumerJWTArrayOutput) ToConsumerJWTArrayOutput() ConsumerJWTArrayOutput {
	return o
}

func (o ConsumerJWTArrayOutput) ToConsumerJWTArrayOutputWithContext(ctx context.Context) ConsumerJWTArrayOutput {
	return o
}

func (o ConsumerJWTArrayOutput) Index(i pulumi.IntInput) ConsumerJWTOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ConsumerJWT {
		return vs[0].([]*ConsumerJWT)[vs[1].(int)]
	}).(ConsumerJWTOutput)
}

type ConsumerJWTMapOutput struct{ *pulumi.OutputState }

func (ConsumerJWTMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConsumerJWT)(nil)).Elem()
}

func (o ConsumerJWTMapOutput) ToConsumerJWTMapOutput() ConsumerJWTMapOutput {
	return o
}

func (o ConsumerJWTMapOutput) ToConsumerJWTMapOutputWithContext(ctx context.Context) ConsumerJWTMapOutput {
	return o
}

func (o ConsumerJWTMapOutput) MapIndex(k pulumi.StringInput) ConsumerJWTOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ConsumerJWT {
		return vs[0].(map[string]*ConsumerJWT)[vs[1].(string)]
	}).(ConsumerJWTOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConsumerJWTInput)(nil)).Elem(), &ConsumerJWT{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConsumerJWTArrayInput)(nil)).Elem(), ConsumerJWTArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConsumerJWTMapInput)(nil)).Elem(), ConsumerJWTMap{})
	pulumi.RegisterOutputType(ConsumerJWTOutput{})
	pulumi.RegisterOutputType(ConsumerJWTArrayOutput{})
	pulumi.RegisterOutputType(ConsumerJWTMapOutput{})
}
