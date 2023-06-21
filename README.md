# Dime
#### Self-hosted budgeting and transaction management

Dime is a self-hostable web application that allows users to track their spending. It is built with Go in the backend and VueJS in the frontend. Users can upload .csv files from their bank and view their spending in a variety of ways. Budgeting goals can be set and tracked. Users can also create and manage multiple accounts.

This is in early development and is some time away from having many of the basic features implemented.

## Milestones
- [x] Basic login/register
- [x] Upload .csv files
- [x] View uploaded transactions
- [ ] Customize meta info for csv processing
- [ ] Transaction account organization
- [ ] User account management
- [ ] Budgeting
- [ ] Graphs and other reporting
- [ ] Mobile-first design
- [ ] A lot of other things I haven't thought of yet

## Development Setup
Go, Node, and Yarn are need to build and run the application. To simplify deployment, a `Dockerfile` is provided along with a `docker-compose.yml` file. If Docker is not used, the the environmental variable `MONGO_HOST` must be set to the host of the MongoDB instance.

## Authentication
Ultimately authentication should be handled via an HttpOnly cookie containing a JWT, however for now the HttpOnly attribute is not set due to some issues with cross-origin settings in the current development environment. This will be fixed in the future.

