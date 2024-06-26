// Copyright 2016-2023, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package konnect

import (
	"fmt"
	"path"

	// Allow embedding bridge-metadata.json in the provider.
	_ "embed"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"

	// Replace this provider with the provider you are bridging.
	konnect "github.com/csechrist/terraform-provider-konnect/konnect"

	"github.com/csechrist/pulumi-konnect/provider/pkg/version"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "konnect"
	// modules:
	mainMod = "index" // the konnect module
)

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(resource.PropertyMap, shim.ResourceConfig) error {
	return nil
}

//go:embed cmd/pulumi-resource-konnect/bridge-metadata.json
var metadata []byte

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		// Instantiate the Terraform provider
		P:       shimv2.NewProvider(konnect.Provider()),
		Name:    "konnect",
		Version: version.Version,
		// DisplayName is a way to be able to change the casing of the provider
		// name when being displayed on the Pulumi registry
		DisplayName: "Kong Konnect",
		// The default publisher for all packages is Pulumi.
		// Change this to your personal name (or a company name) that you
		// would like to be shown in the Pulumi Registry if this package is published
		// there.
		Publisher: "Cameron Sechrist",
		// LogoURL is optional but useful to help identify your package in the Pulumi Registry
		// if this package is published there.
		//
		// You may host a logo on a domain you control or add an SVG logo for your package
		// in your repository and use the raw content URL for that file as your logo URL.
		LogoURL: "",
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		// e.g https://github.com/org/pulumi-provider-name/releases/
		PluginDownloadURL: "github://api.github.com/csechrist/pulumi-konnect",
		Description:       "A Pulumi package for creating and managing Kong konnect cloud resources.",
		// category/cloud tag helps with categorizing the package in the Pulumi Registry.
		// For all available categories, see `Keywords` in
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords:   []string{"pulumi", "konnect", "category/cloud"},
		License:    "Apache-2.0",
		Homepage:   "https://www.pulumi.com",
		Repository: "https://github.com/csechrist/pulumi-konnect",
		// The GitHub Org for the provider - defaults to `terraform-providers`. Note that this
		// should match the TF provider module's require directive, not any replace directives.
		GitHubOrg:    "csechrist",
		MetadataInfo: tfbridge.NewProviderMetadata(metadata),
		Config:       map[string]*tfbridge.SchemaInfo{
			// Add any required configuration here, or remove the example below if
			// no additional points are required.
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			"konnect_control_plane":           {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ControlPlane")},
			"konnect_authentication_settings": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AuthenticationSettings")},
			"konnect_identity_provider":       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "IdentityProvider")},
			"konnect_user":                    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "User")},
			"konnect_team":                    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Team")},
			"konnect_team_user":               {Tok: tfbridge.MakeResource(mainPkg, mainMod, "TeamUser")},
			"konnect_team_role":               {Tok: tfbridge.MakeResource(mainPkg, mainMod, "TeamRole")},
			"konnect_team_mappings":           {Tok: tfbridge.MakeResource(mainPkg, mainMod, "TeamMappings")},
			"konnect_user_role":               {Tok: tfbridge.MakeResource(mainPkg, mainMod, "UserRole")},
			"konnect_service":                 {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Service")},
			"konnect_route":                   {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Route")},
			"konnect_consumer":                {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Consumer")},
			"konnect_consumer_key":            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ConsumerKey")},
			"konnect_consumer_acl":            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ConsumerACL")},
			"konnect_consumer_basic":          {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ConsumerBasic")},
			"konnect_consumer_hmac":           {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ConsumerHMAC")},
			"konnect_consumer_jwt":            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ConsumerJWT")},
			"konnect_plugin":                  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Plugin")},
			"konnect_certificate":             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Certificate")},
			"konnect_dp_certificate":          {Tok: tfbridge.MakeResource(mainPkg, mainMod, "DpCertificate")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"konnect_control_plane": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getControlPlane")},
			"konnect_user":          {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getUser")},
			"konnect_team":          {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getTeam")},
			"konnect_role":          {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getRole")},
			"konnect_team_role":     {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getTeamRole")},
			"konnect_user_role":     {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getUserRole")},
			"konnect_nodes":         {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getNodes")},
			"konnect_consumer":      {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getConsumer")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@camsechrist/pulumi-konnect",
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: path.Join(
				fmt.Sprintf("github.com/csechrist/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	// These are new API's that you may opt to use to automatically compute resource
	// tokens, and apply auto aliasing for full backwards compatibility.  For more
	// information, please reference:
	// https://pkg.go.dev/github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge#ProviderInfo.ComputeTokens
	prov.MustComputeTokens(tokens.SingleModule("konnect_", mainMod,
		tokens.MakeStandard(mainPkg)))
	prov.MustApplyAutoAliases()
	prov.SetAutonaming(255, "-")

	return prov
}
