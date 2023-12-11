# scayle
dx poc skeleton

## Usage

```
demo@user scayle % ./scayle tenant create-environment
Error: required flag(s) "name" not set
Usage:
  scayle tenant create-environment [flags]

Flags:
  -h, --help          help for create-environment
      --name string   Name of the environment
      --size string   Size of the environment (default "M")

demo@user scayle % ./scayle tenant create-environment --name "demo"
Creating a new environment "demo", with size "M"

	An example workflow execution:
	1. Create a new configuration repository for the environment (call VCS API and provision new repo from template)
	2. Populate the new repository with correct settings
	3. Link the repository to a relevant CI/CD system
	4. Call the deployment job
```

## Building it

```
go build .
```