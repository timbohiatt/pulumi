package project

import (
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Settings struct {
	run_tests bool
}

type Args struct {
	billing_account string
	name            string
	parent_folder   string
	parent_org      string
	prefix          string
	services        []string
}

type Factory struct {
	Args     *Args
	Settings *Settings
	Project  *organizations.Project
}

func (factory *Factory) Create(ctx *pulumi.Context) (err error) {

	// Create Pulumi Google Cloud Project Arguments Object
	args := &organizations.ProjectArgs{}

	// Confgure Google Cloud Project - Pulumi Arguments
	args.Name = pulumi.String(factory.Args.name)
	args.ProjectId = pulumi.String(factory.Args.name)
	args.BillingAccount = pulumi.String(factory.Args.name)
	args.AutoCreateNetwork = pulumi.Bool(false)
	if factory.Args.parent_folder != "" {
		args.FolderId = pulumi.String(factory.Args.parent_folder)
	} else {
		args.OrgId = pulumi.String(factory.Args.parent_org)
	}

	//args.Labels = [],
	project, err := organizations.NewProject(ctx, "project", args)

	ctx.Export("project", project)
	factory.Project = project

	// Return the Factory & Error
	return err

}
