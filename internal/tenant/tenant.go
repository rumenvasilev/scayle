package tenant

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Create(cmd *cobra.Command) error {
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		return err
	}
	size, err := cmd.Flags().GetString("size")
	if err != nil {
		return err
	}

	fmt.Printf("Creating a new environment %q, with size %q\n", name, size)

	fmt.Println(`
	An example workflow execution:
	1. Create a new configuration repository for the environment (call VCS API and provision new repo from template)
	2. Populate the new repository with correct settings
	3. Link the repository to a relevant CI/CD system
	4. Call the deployment job
	`)
	return nil
}
