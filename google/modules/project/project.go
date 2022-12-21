package Project

import (
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Settings struct {
	RunTests bool
}

type Args struct {
	BillingAccount string
	Name           string
	ParentFolder   string
	ParentOrg      string
	Prefix         string
	Services       []string
}

type Module struct {
	Args     *Args
	Settings *Settings
	Project  *organizations.Project
}

func (module *Module) Create(ctx *pulumi.Context) (err error) {

	// Create Pulumi Google Cloud Project Arguments Object
	args := &organizations.ProjectArgs{}

	// Confgure Google Cloud Project - Pulumi Arguments
	args.Name = pulumi.String(module.Args.Name)
	args.ProjectId = pulumi.String(module.Args.Name)
	args.BillingAccount = pulumi.String(module.Args.Name)
	args.AutoCreateNetwork = pulumi.Bool(false)
	if module.Args.ParentFolder != "" {
		args.FolderId = pulumi.String(module.Args.ParentFolder)
	} else {
		args.OrgId = pulumi.String(module.Args.ParentOrg)
	}

	//args.Labels = [],
	project, err := organizations.NewProject(ctx, "project", args)

	ctx.Export("project", project)
	module.Project = project

	// Return the MOdule & Error
	return err

}
