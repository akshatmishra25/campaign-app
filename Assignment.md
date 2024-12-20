# Assignment

Consider a social media campaign with the following fields:
1. Name
2. Description 
3. Status (ACTIVE, INACTIVE, DELETED) 
4. Leads (Array of social media link of people to reach out to)
5. AccountIDs (Array of account ids to use to send messages)

Example of a social media campaign:
```json
{
  "name": "Campaign 1",
  "description": "This is a campaign to reach out to people",
  "status": "active",
  "leads": ["https://www.facebook.com/123", "https://www.twitter.com/123"],
  "accountIDs": ["123", "456"]
}
```

---
Implement the following methods for the social media campaign - 
1. GET /campaigns - Get all campaigns which are not DELETED 
2. GET /campaigns/:id - Get a campaign by id
3. POST /campaigns - Create a campaign
4. PUT /campaigns/:id - Update a campaign (change status to active or inactive)
5. DELETE /campaigns/:id - Delete a campaign (change status to deleted)

---
Note - 
1. Backend should be in Go
2. Use gin router for handling requests 
3. Use MongoDB for storing the data


