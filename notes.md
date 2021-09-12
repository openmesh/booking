# Entities To Create

- Auth
- Booking
- Resource
- Unavailability
- User

# Create user flow

- Organization is created either via signup or through default organization in config
- User is created through invite or through default user

- User has API key

# User stories

- [ ] Users should be able to create a resource

# Create organization payload

- Name
- OwnerEmail
- OwnerPassword
- OwnerName


# Layers written out

## Domain
- Entity definitions

## Service 
- Implementation of business logic

## Transport
- API layer that hands requests to service layer

