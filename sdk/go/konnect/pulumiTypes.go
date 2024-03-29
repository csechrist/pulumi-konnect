// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package konnect

import (
	"context"
	"reflect"

	"github.com/csechrist123/pulumi-konnect/sdk/go/konnect/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type RouteHeader struct {
	// **(Required, String)** Name of header this route should require.
	Name string `pulumi:"name"`
	// **(Required, List of String)** Allowed values this header should equal.
	Values []string `pulumi:"values"`
}

// RouteHeaderInput is an input type that accepts RouteHeaderArgs and RouteHeaderOutput values.
// You can construct a concrete instance of `RouteHeaderInput` via:
//
//	RouteHeaderArgs{...}
type RouteHeaderInput interface {
	pulumi.Input

	ToRouteHeaderOutput() RouteHeaderOutput
	ToRouteHeaderOutputWithContext(context.Context) RouteHeaderOutput
}

type RouteHeaderArgs struct {
	// **(Required, String)** Name of header this route should require.
	Name pulumi.StringInput `pulumi:"name"`
	// **(Required, List of String)** Allowed values this header should equal.
	Values pulumi.StringArrayInput `pulumi:"values"`
}

func (RouteHeaderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteHeader)(nil)).Elem()
}

func (i RouteHeaderArgs) ToRouteHeaderOutput() RouteHeaderOutput {
	return i.ToRouteHeaderOutputWithContext(context.Background())
}

func (i RouteHeaderArgs) ToRouteHeaderOutputWithContext(ctx context.Context) RouteHeaderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteHeaderOutput)
}

// RouteHeaderArrayInput is an input type that accepts RouteHeaderArray and RouteHeaderArrayOutput values.
// You can construct a concrete instance of `RouteHeaderArrayInput` via:
//
//	RouteHeaderArray{ RouteHeaderArgs{...} }
type RouteHeaderArrayInput interface {
	pulumi.Input

	ToRouteHeaderArrayOutput() RouteHeaderArrayOutput
	ToRouteHeaderArrayOutputWithContext(context.Context) RouteHeaderArrayOutput
}

type RouteHeaderArray []RouteHeaderInput

func (RouteHeaderArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RouteHeader)(nil)).Elem()
}

func (i RouteHeaderArray) ToRouteHeaderArrayOutput() RouteHeaderArrayOutput {
	return i.ToRouteHeaderArrayOutputWithContext(context.Background())
}

func (i RouteHeaderArray) ToRouteHeaderArrayOutputWithContext(ctx context.Context) RouteHeaderArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteHeaderArrayOutput)
}

type RouteHeaderOutput struct{ *pulumi.OutputState }

func (RouteHeaderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteHeader)(nil)).Elem()
}

func (o RouteHeaderOutput) ToRouteHeaderOutput() RouteHeaderOutput {
	return o
}

func (o RouteHeaderOutput) ToRouteHeaderOutputWithContext(ctx context.Context) RouteHeaderOutput {
	return o
}

// **(Required, String)** Name of header this route should require.
func (o RouteHeaderOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RouteHeader) string { return v.Name }).(pulumi.StringOutput)
}

// **(Required, List of String)** Allowed values this header should equal.
func (o RouteHeaderOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RouteHeader) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type RouteHeaderArrayOutput struct{ *pulumi.OutputState }

func (RouteHeaderArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RouteHeader)(nil)).Elem()
}

func (o RouteHeaderArrayOutput) ToRouteHeaderArrayOutput() RouteHeaderArrayOutput {
	return o
}

func (o RouteHeaderArrayOutput) ToRouteHeaderArrayOutputWithContext(ctx context.Context) RouteHeaderArrayOutput {
	return o
}

func (o RouteHeaderArrayOutput) Index(i pulumi.IntInput) RouteHeaderOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RouteHeader {
		return vs[0].([]RouteHeader)[vs[1].(int)]
	}).(RouteHeaderOutput)
}

type TeamMappingsMapping struct {
	// **(Required, String)** Identifier of an IdP group that is contained with OIDC ID token for groups claim
	Group string `pulumi:"group"`
	// **(Required, List of String)** Konnect teams that should map to this group.
	TeamIds []string `pulumi:"teamIds"`
}

// TeamMappingsMappingInput is an input type that accepts TeamMappingsMappingArgs and TeamMappingsMappingOutput values.
// You can construct a concrete instance of `TeamMappingsMappingInput` via:
//
//	TeamMappingsMappingArgs{...}
type TeamMappingsMappingInput interface {
	pulumi.Input

	ToTeamMappingsMappingOutput() TeamMappingsMappingOutput
	ToTeamMappingsMappingOutputWithContext(context.Context) TeamMappingsMappingOutput
}

type TeamMappingsMappingArgs struct {
	// **(Required, String)** Identifier of an IdP group that is contained with OIDC ID token for groups claim
	Group pulumi.StringInput `pulumi:"group"`
	// **(Required, List of String)** Konnect teams that should map to this group.
	TeamIds pulumi.StringArrayInput `pulumi:"teamIds"`
}

func (TeamMappingsMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TeamMappingsMapping)(nil)).Elem()
}

func (i TeamMappingsMappingArgs) ToTeamMappingsMappingOutput() TeamMappingsMappingOutput {
	return i.ToTeamMappingsMappingOutputWithContext(context.Background())
}

func (i TeamMappingsMappingArgs) ToTeamMappingsMappingOutputWithContext(ctx context.Context) TeamMappingsMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamMappingsMappingOutput)
}

// TeamMappingsMappingArrayInput is an input type that accepts TeamMappingsMappingArray and TeamMappingsMappingArrayOutput values.
// You can construct a concrete instance of `TeamMappingsMappingArrayInput` via:
//
//	TeamMappingsMappingArray{ TeamMappingsMappingArgs{...} }
type TeamMappingsMappingArrayInput interface {
	pulumi.Input

	ToTeamMappingsMappingArrayOutput() TeamMappingsMappingArrayOutput
	ToTeamMappingsMappingArrayOutputWithContext(context.Context) TeamMappingsMappingArrayOutput
}

type TeamMappingsMappingArray []TeamMappingsMappingInput

func (TeamMappingsMappingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TeamMappingsMapping)(nil)).Elem()
}

func (i TeamMappingsMappingArray) ToTeamMappingsMappingArrayOutput() TeamMappingsMappingArrayOutput {
	return i.ToTeamMappingsMappingArrayOutputWithContext(context.Background())
}

func (i TeamMappingsMappingArray) ToTeamMappingsMappingArrayOutputWithContext(ctx context.Context) TeamMappingsMappingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamMappingsMappingArrayOutput)
}

type TeamMappingsMappingOutput struct{ *pulumi.OutputState }

func (TeamMappingsMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TeamMappingsMapping)(nil)).Elem()
}

func (o TeamMappingsMappingOutput) ToTeamMappingsMappingOutput() TeamMappingsMappingOutput {
	return o
}

func (o TeamMappingsMappingOutput) ToTeamMappingsMappingOutputWithContext(ctx context.Context) TeamMappingsMappingOutput {
	return o
}

// **(Required, String)** Identifier of an IdP group that is contained with OIDC ID token for groups claim
func (o TeamMappingsMappingOutput) Group() pulumi.StringOutput {
	return o.ApplyT(func(v TeamMappingsMapping) string { return v.Group }).(pulumi.StringOutput)
}

// **(Required, List of String)** Konnect teams that should map to this group.
func (o TeamMappingsMappingOutput) TeamIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TeamMappingsMapping) []string { return v.TeamIds }).(pulumi.StringArrayOutput)
}

type TeamMappingsMappingArrayOutput struct{ *pulumi.OutputState }

func (TeamMappingsMappingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TeamMappingsMapping)(nil)).Elem()
}

func (o TeamMappingsMappingArrayOutput) ToTeamMappingsMappingArrayOutput() TeamMappingsMappingArrayOutput {
	return o
}

func (o TeamMappingsMappingArrayOutput) ToTeamMappingsMappingArrayOutputWithContext(ctx context.Context) TeamMappingsMappingArrayOutput {
	return o
}

func (o TeamMappingsMappingArrayOutput) Index(i pulumi.IntInput) TeamMappingsMappingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TeamMappingsMapping {
		return vs[0].([]TeamMappingsMapping)[vs[1].(int)]
	}).(TeamMappingsMappingOutput)
}

type GetNodesNode struct {
	// **(String)** Hash of the current configuration state of node.
	ConfigHash string `pulumi:"configHash"`
	// **(String)** Id of certificate used in communication between node and control plane.
	DataPlaneCertId string `pulumi:"dataPlaneCertId"`
	// **(String)** Hostname of node.
	Hostname string `pulumi:"hostname"`
	// **(String)** Id of node.
	Id string `pulumi:"id"`
	// **(Integer)** Last time of ping of node.
	LastPing int `pulumi:"lastPing"`
	// **(String)** Type of node.
	Type string `pulumi:"type"`
	// **(String)** Version of node.
	Version string `pulumi:"version"`
}

// GetNodesNodeInput is an input type that accepts GetNodesNodeArgs and GetNodesNodeOutput values.
// You can construct a concrete instance of `GetNodesNodeInput` via:
//
//	GetNodesNodeArgs{...}
type GetNodesNodeInput interface {
	pulumi.Input

	ToGetNodesNodeOutput() GetNodesNodeOutput
	ToGetNodesNodeOutputWithContext(context.Context) GetNodesNodeOutput
}

type GetNodesNodeArgs struct {
	// **(String)** Hash of the current configuration state of node.
	ConfigHash pulumi.StringInput `pulumi:"configHash"`
	// **(String)** Id of certificate used in communication between node and control plane.
	DataPlaneCertId pulumi.StringInput `pulumi:"dataPlaneCertId"`
	// **(String)** Hostname of node.
	Hostname pulumi.StringInput `pulumi:"hostname"`
	// **(String)** Id of node.
	Id pulumi.StringInput `pulumi:"id"`
	// **(Integer)** Last time of ping of node.
	LastPing pulumi.IntInput `pulumi:"lastPing"`
	// **(String)** Type of node.
	Type pulumi.StringInput `pulumi:"type"`
	// **(String)** Version of node.
	Version pulumi.StringInput `pulumi:"version"`
}

func (GetNodesNodeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNodesNode)(nil)).Elem()
}

func (i GetNodesNodeArgs) ToGetNodesNodeOutput() GetNodesNodeOutput {
	return i.ToGetNodesNodeOutputWithContext(context.Background())
}

func (i GetNodesNodeArgs) ToGetNodesNodeOutputWithContext(ctx context.Context) GetNodesNodeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetNodesNodeOutput)
}

// GetNodesNodeArrayInput is an input type that accepts GetNodesNodeArray and GetNodesNodeArrayOutput values.
// You can construct a concrete instance of `GetNodesNodeArrayInput` via:
//
//	GetNodesNodeArray{ GetNodesNodeArgs{...} }
type GetNodesNodeArrayInput interface {
	pulumi.Input

	ToGetNodesNodeArrayOutput() GetNodesNodeArrayOutput
	ToGetNodesNodeArrayOutputWithContext(context.Context) GetNodesNodeArrayOutput
}

type GetNodesNodeArray []GetNodesNodeInput

func (GetNodesNodeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetNodesNode)(nil)).Elem()
}

func (i GetNodesNodeArray) ToGetNodesNodeArrayOutput() GetNodesNodeArrayOutput {
	return i.ToGetNodesNodeArrayOutputWithContext(context.Background())
}

func (i GetNodesNodeArray) ToGetNodesNodeArrayOutputWithContext(ctx context.Context) GetNodesNodeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetNodesNodeArrayOutput)
}

type GetNodesNodeOutput struct{ *pulumi.OutputState }

func (GetNodesNodeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNodesNode)(nil)).Elem()
}

func (o GetNodesNodeOutput) ToGetNodesNodeOutput() GetNodesNodeOutput {
	return o
}

func (o GetNodesNodeOutput) ToGetNodesNodeOutputWithContext(ctx context.Context) GetNodesNodeOutput {
	return o
}

// **(String)** Hash of the current configuration state of node.
func (o GetNodesNodeOutput) ConfigHash() pulumi.StringOutput {
	return o.ApplyT(func(v GetNodesNode) string { return v.ConfigHash }).(pulumi.StringOutput)
}

// **(String)** Id of certificate used in communication between node and control plane.
func (o GetNodesNodeOutput) DataPlaneCertId() pulumi.StringOutput {
	return o.ApplyT(func(v GetNodesNode) string { return v.DataPlaneCertId }).(pulumi.StringOutput)
}

// **(String)** Hostname of node.
func (o GetNodesNodeOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v GetNodesNode) string { return v.Hostname }).(pulumi.StringOutput)
}

// **(String)** Id of node.
func (o GetNodesNodeOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetNodesNode) string { return v.Id }).(pulumi.StringOutput)
}

// **(Integer)** Last time of ping of node.
func (o GetNodesNodeOutput) LastPing() pulumi.IntOutput {
	return o.ApplyT(func(v GetNodesNode) int { return v.LastPing }).(pulumi.IntOutput)
}

// **(String)** Type of node.
func (o GetNodesNodeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetNodesNode) string { return v.Type }).(pulumi.StringOutput)
}

// **(String)** Version of node.
func (o GetNodesNodeOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v GetNodesNode) string { return v.Version }).(pulumi.StringOutput)
}

type GetNodesNodeArrayOutput struct{ *pulumi.OutputState }

func (GetNodesNodeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetNodesNode)(nil)).Elem()
}

func (o GetNodesNodeArrayOutput) ToGetNodesNodeArrayOutput() GetNodesNodeArrayOutput {
	return o
}

func (o GetNodesNodeArrayOutput) ToGetNodesNodeArrayOutputWithContext(ctx context.Context) GetNodesNodeArrayOutput {
	return o
}

func (o GetNodesNodeArrayOutput) Index(i pulumi.IntInput) GetNodesNodeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetNodesNode {
		return vs[0].([]GetNodesNode)[vs[1].(int)]
	}).(GetNodesNodeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RouteHeaderInput)(nil)).Elem(), RouteHeaderArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouteHeaderArrayInput)(nil)).Elem(), RouteHeaderArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TeamMappingsMappingInput)(nil)).Elem(), TeamMappingsMappingArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TeamMappingsMappingArrayInput)(nil)).Elem(), TeamMappingsMappingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetNodesNodeInput)(nil)).Elem(), GetNodesNodeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetNodesNodeArrayInput)(nil)).Elem(), GetNodesNodeArray{})
	pulumi.RegisterOutputType(RouteHeaderOutput{})
	pulumi.RegisterOutputType(RouteHeaderArrayOutput{})
	pulumi.RegisterOutputType(TeamMappingsMappingOutput{})
	pulumi.RegisterOutputType(TeamMappingsMappingArrayOutput{})
	pulumi.RegisterOutputType(GetNodesNodeOutput{})
	pulumi.RegisterOutputType(GetNodesNodeArrayOutput{})
}
