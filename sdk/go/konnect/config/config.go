// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/csechrist/pulumi-konnect/sdk/go/konnect/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

var _ = internal.GetEnvOrDefault

func GetPat(ctx *pulumi.Context) string {
	return config.Get(ctx, "konnect:pat")
}
func GetRegion(ctx *pulumi.Context) string {
	return config.Get(ctx, "konnect:region")
}
