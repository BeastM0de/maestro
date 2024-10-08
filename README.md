# maestro
Git Maestro is a powerful orchestration tool designed to take the pain out of managing complex workspaces.  Maestro can quickly create, sync, update, and perform other operations across an entire workspace.  Workspace state is stored as sharable configuration files that can be used to quickly initialize and sync workspaces.  

Maestro also allows you to view, analyze, and query against your workspace and the underlying repositories it contains.

Requirements

Installation
```bash
apt-get install maestro
```

## Setup
To initialize a new workspace
```bash
maestro init
```
Or to initialize from an existing configuration:
```bash
maestro clone [FILE URL]
```

## Usage 
```bash
# 
maestro pull
maestro import
maestro clean
maestro stat
maestro add REPO
maestro remove REPO
maestro list repos


maestro run CMD
maestro run CMD_ALIAS
```

### Options

`--filter -f FILTER`
Apply a filter to the command being run.

`--profile -p PROFILE`
The workspace profile to use.

`--output json`
Format output as JSON.

`--save-alias ALIAS=CMD`
Save an alias for a command

## Configuration 
Maestro has two levels of configuration, global and workspace.  Both are stored in `.maestro` directories in `$HOME` and the workspace root respectively.
```
$HOME/.maestro
credentials
settings 
workspace
workspace-{profle}
```

## Commands

### init
Initialize a new workspace




