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
- Better error messages for constraint violations (TBD)
- Only allow changes to feature test variant id for user feature test variant update (TBD)

DB:
- check constraint on test variant percentage sum <= 100
- check constraint on test variants for a test only having one row is_control = true

## Notes
- graphql DateTime field expects an input like "2020-05-05T00:00:00Z". It's pretty picky about the time formatting.