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
- Update an existing feature test (TBD)
    - Name
    - End date
- Create a feature test variant (TBD)
    - Name
    - Percentage
    - Is control?
- Update a feature test variant (TBD)
    - Name
    - Percentage
    - Is control?
- Delete a feature test variant (TBD)
    - Cascade to users with that test variant (TBD)
- Assign users to a test variant (TBD)
- Change test variants for a user (TBD)
    
