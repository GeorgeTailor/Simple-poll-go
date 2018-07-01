# Simple-poll-go

Simple CRUD application written in Go. 

## Endpoints

### GET

`/polls` - returns all polls available.

### POST

`/poll_options` - returns specific poll by given PollId or returns empty struct when poll with such id does not exist in storage      
`/select_option` - increments the number of how many times the option by given ID was selected in the poll with given PollID        
`/add_poll` - adds Poll struct to the storage and returns its ID        
`/add_option` - adds Option struct to the poll with given ID in the storage returns poll ID       

### DELETE

`/remove_poll` - remove poll from storage by given ID, returns true when poll was found and removed or false when poll was not found         
`/remove_option` - remove option from poll by given ID, returns true when option was found and removed or false when option was not found         
### Postman collection

Postman collection containing all of the above endpoints can be found here: 
https://www.getpostman.com/collections/13b37bc0f2168203fda4       
or      
use file in the sources of this repo `Go polls.postman_collection.json`
