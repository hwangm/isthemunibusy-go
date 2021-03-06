# isthemunibusy-go

This was originally a project to predict how crowded the muni trains would be, but given the current circumstances with COVID-19, I have decided to pivot this project to be a feature test admin tool.

## Features
- Create a new feature test (DONE)
    - Name
    - Start date
    - End date
    - Variants[]
        - Name
        - Percentage
        - Is control?
- Delete an existing feature test (DONE)
    - Cascade to feature test variants 
    - Cascade to users with those test variants 
- Update an existing feature test (DONE)
    - Name
    - End date
- Create a feature test variant (DONE)
    - Name
    - Percentage
    - Is control?
- Update a feature test variant (DONE)
    - Name
    - Percentage
    - Is control?
- Delete a feature test variant (DONE)
    - Cascade to users with that test variant
- List user test variants (DONE)
- List feature tests (DONE)
- List feature test variants (DONE)
- Assign users to a test variant (create user test variant) (DONE)
- Change test variants for a user (update user test variant) (DONE)
- Delete user test variant (DONE)
- Only allow changes to feature test variant id for user feature test variant update (DONE)
- Better error messages for constraint violations (DONE)

DB:
- check constraint on test variant percentage sum <= 100 (DONE)
- check constraint on test variants for a test only having one row is_control = true
    - insert (DONE)
    - update (leaving alone for now)

## How to run locally
Navigate to the go repo directory (`go\src\github.com\hwangm\isthemunibusy-go`)
In a terminal window, run `go run .\main.go`.

To restart, Ctrl-C and re-run.

## Notes
- graphql DateTime field expects an input like "2020-05-05T00:00:00Z". It's pretty picky about the time formatting.

- 2/15/2021: added a websocket handler at "ws://localhost:8080/websocket" to test notification service. React app currently configured to auto-connect to this WS endpoint. Incoming messages will be printed in stdout, then sleep for 5 seconds and send a message to the client. 

