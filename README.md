# isthemunibusy-go

This was originally a project to predict how crowded the muni trains would be, but given the current circumstances with COVID-19, I have decided to pivot this project to be a feature test admin tool.

## Features
- Create a new feature test
    - Name
    - Start date
    - End date
- Delete an existing feature test
    - Cascade to feature test variants
    - Cascade to users with those test variants
- Update an existing feature test
    - Name
    - End date
- Create a feature test variant
    - Name
    - Percentage
    - Is control?
- Update a feature test variant
    - Name
    - Percentage
    - Is control?
- Delete a feature test variant
    - Cascade to users with that test variant
- Assign users to a test variant
- Change test variants for a user
    
