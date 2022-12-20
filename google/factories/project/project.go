package project

import (
	"fmt"

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

type Factory struct {
	Args     *Args
	Settings *Settings
	Project  *organizations.Project
}

func (factory *Factory) Create(ctx *pulumi.Context) (err error) {

	// Create Pulumi Google Cloud Project Arguments Object
	args := &organizations.ProjectArgs{}

	// Confgure Google Cloud Project - Pulumi Arguments
	args.Name = pulumi.String(factory.Args.Name)
	args.ProjectId = pulumi.String(factory.Args.Name)
	args.BillingAccount = pulumi.String(factory.Args.Name)
	args.AutoCreateNetwork = pulumi.Bool(false)
	if factory.Args.ParentFolder != "" {
		args.FolderId = pulumi.String(factory.Args.ParentFolder)
	} else {
		args.OrgId = pulumi.String(factory.Args.ParentOrg)
	}

	//args.Labels = [],
	project, err := organizations.NewProject(ctx, "project", args)

	ctx.Export("project", project)
	factory.Project = project

	// Return the Factory & Error
	return err

}

func main() {
	fmt.Println("Running...")
}
