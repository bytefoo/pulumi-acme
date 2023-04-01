package acme

import (
	"fmt"
	"path/filepath"

	"github.com/bytefoo/pulumi-acme/provider/pkg/version"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/vancluever/terraform-provider-acme/v2/acme"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "acme"
	// modules:
	mainMod = "index" // the acme module
)

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(acme.Provider())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:                    p,
		Name:                 "acme",
		DisplayName:          "Acme",
		Publisher:            "bytefoo",
		LogoURL:              "",
		PluginDownloadURL:    "github://api.github.com/bytefoo/pulumi-acme",
		Description:          "A Pulumi package for creating and managing acme cloud resources.",
		Keywords:             []string{"pulumi", "acme", "bytefoo"},
		License:              "Apache-2.0",
		Homepage:             "https://www.pulumi.com",
		Repository:           "https://github.com/bytefoo/pulumi-acme",
		GitHubOrg:            "vancluever",
		Config:               map[string]*tfbridge.SchemaInfo{},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			"acme_certificate":  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Certificate")},
			"acme_registration": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Registration")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"acme_certificate":  {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getCertificate")},
			"acme_registration": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getRegistration")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@bytefoo/pulumi-acme",
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0",
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			PackageName: "bytefoo_pulumi_acme",
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/bytefoo/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RootNamespace: "ByteFoo.PulumiPackage",
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
