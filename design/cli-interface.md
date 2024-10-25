

| Category      | Command                           | Description                           |
|---------------|-----------------------------------|---------------------------------------|
| Workspace     | maestro init                      | Initialize a new workspace            |
|               | maestro clone                     | Clone a workspace                     |
| ------------- | --------------------------------- | ------------------------------------- |
| Configuration | maestro config view               | Show the configuration                |
|               | maestro config                    | Show the configuration                |
| ------------- | --------------------------------- | ------------------------------------- |
| Repos         | maestro list                      | List all repositories                 |
|               | maestro status                    | Show the status of all repositories   |
|               | maestro pull all                  | Pull all repositories                 |
|               | maestro clean                     | Clean all repositories                |
|               | maestro describe                  | Describe a repository                 |
|               | maestro remove                    | Remove a repository                   |
| ------------- | --------------------------------  | ------------------------------------- |
| Groups        | maestro add URL                   | List all repositories                 |
|               | maestro list                      | List all repositories                 |
|               | maestro status                    | Show the status of all repositories   |
|               | maestro pull                      | Pull all repositories                 |
|               | maestro clean                     | Clean all repositories                |
|               | maestro describe                  | Describe a repository                 |
|               | maestro remove                    | Remove a repository                   |
| ------------- | --------------------------------- | ------------------------------------- |
| Repo          | maestro add REPO                  | Add a repository to the workspace     |
|               | maestro pull                      | Pull a repository                     |
|               | maestro remove REPO               | Remove a repository                   |
|               | maestro describe                  | Describe a repository                 |
| ------------- | --------------------------------  | ------------------------------------- |
| Commands      | maestro run CMD                   | Run a command against all repos       |
|               | maestro run CMD --group GROUP     | Run a command against a set of repos  |
|               | maestro run CMD --filter GROUP    | Run a command against a group         |
|               | maestro run CMD --repo REPO       | Run a command against a specific repo |
| ------------- | --------------------------------  | ------------------------------------- |

### flags
  -f, --filter FILTER
  -p, --profile PROFILE
  --output JSON
  --save-alias ALIAS=CMD
  --save-cmd
  
