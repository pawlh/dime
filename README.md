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

## 260 Rubric Items
- [x] Calls third party service endpoints
  - About page, project information is loaded from a service that shows repo information
- [x] Provides service endpoints running under ~~Node.js~~ Go
  - The frontend is served through a static endpoint
  - Logging in, registering, uploading files are all handled by GET/POST endpoints
- [x] Stores data in MongoDB
  - User and transaction data is stored in MongoDB
- [x] Provides authenticated login with securely stored credentials in MongoDB
  - User passwords are hashed and stored in MongoDB
- [x] Peer communication using WebSockets
  - Although the use case is a little contrived, the transactions page is updated in real time when a new transaction is uploaded via WebSockets
- [x] Multiple Git commits with meaningful comments.
  - 247 at the time of this writing. 248 by the time this is committed.
- [x] Notes in your start up Git repository README.md file documenting what you have learned using services, node.js, mongodb, authentication, and webSockets.
  - This file

## Development Setup
Go, Node, and Yarn are need to build and run the application. To simplify deployment, a `Dockerfile` is provided along with a `docker-compose.yml` file. If Docker is not used, the the environmental variable `MONGO_HOST` must be set to the host of the MongoDB instance.

## Authentication
Ultimately authentication should be handled via an HttpOnly cookie containing a JWT, however for now the HttpOnly attribute is not set due to some issues with cross-origin settings in the current development environment. This will be fixed in the future.

## What I've learned
### WebSockets
Getting WebSockets going was pretty easy with Gorilla, but I had some issues with sending the `Ping` message. For some reason, the client wasn't receiving messages with the Ping type, so I had to send a message with a different type. I have not been able to resolve this, but I did observe that the socket connections were staying open even after 30 minutes of inactivity on Chrome and Firefox.

### MongoDB
I am most familiar with relational databases, so I had to do some research to figure out how to store the data in MongoDB. I ended up using a single collection for users and another for transactions. Without a schema, I found that I kept leaning on programmatic solutions rather than using MongoDB native features. I think I would have been better off using a schema and using the native features of MongoDB. I found the [operator documentation](https://www.mongodb.com/docs/v6.0/reference/operator/) to be very helpful.

### Authentication
Functional and secure authentication is tough. After researching lots of auth solutions, I found that there is usually a trade-off between security and ease of use. I decided to go with a simple username/password solution for now that relies on non-HttpOnly cookie so that the client doesn't need to ping a `/api/me` endpoint, but there are definitely much more secure options that aren't too much more difficult to implement that I plan to look into in the near future.