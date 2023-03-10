package project

import (
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/serviceaccount"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/timbohiatt/pulumi/google/modules/project"
	"github.com/timbohiatt/pulumi/google/modules/serviceAccount"
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
	Args            *Args
	Settings        *Settings
	Project         *organizations.Project
	ServiceAccounts []*serviceaccount.Account
}

func (factory *Factory) Create(ctx *pulumi.Context) (err error) {

	// Create Project

	// Create Pulumi Google Cloud Project Arguments Object
	prj := &project.Module{}

	// Confgure Google Cloud Project - Pulumi Arguments
	prj.Args.Name = factory.Args.Name
	prj.Args.ProjectId = factory.Args.Name
	prj.Args.BillingAccount = factory.Args.BillingAccount
	prj.Args.AutoCreateNetwork = false
	if factory.Args.ParentFolder != "" {
		prj.Args.FolderId = factory.Args.ParentFolder
	} else {
		prj.Args.OrgId = factory.Args.ParentOrg
	}

	// Create the Project from Module
	err = prj.Create(ctx)
	if err != nil {
		return err
	}
	factory.Project = prj.Project

	// Create Service Accounts

	// Create Pulumi Google Cloud Project Arguments Object
	sa := &serviceAccount.Module{}
	// Confgure Google Cloud Project - Pulumi Arguments
	sa.Args.ProjectId = factory.Args.Name
	sa.Args.AccountId = factory.Args.Name
	sa.Args.DisplayName = factory.Args.Name
	sa.Args.Description = factory.Args.Name
	sa.Args.Disabled = false

	// Create the Service Account from Module
	err = sa.Create(ctx)
	if err != nil {
		return err
	}
	
	// Add Service Account to Factory Collection of Service Accounts
	append(factory.ServiceAccounts, sa.ServiceAccount)

	// Return the Factory & Error
	return err

}
