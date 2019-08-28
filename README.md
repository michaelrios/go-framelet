# Go API Template

# Layered Approach, and general rules
* Router > Middleware > Controller > Domain > (Repository, Gateway)
* Models should be accessible anywhere
* Config should only be used in main

# Router
* I used HTTPRouter since it the the fastest option, with the most features.
* Should only route http requests

# Middleware
* I setup a custom Middleware solution, since it is simple and will do what 80% of us will need
* Middleware should have very limited access to dependencies
* The Access Middleware is a good example of doing stuff before and after the rest of the application logic runs

# Controller
* Controllers should only be able to call the Domain layer. This protects you from overly complex controller functions, and makes testing a lot easier
* Think of your controller methods as a conductor, they are the shot callers, and should not be doing much work

# Domain
* This is where your business logic goes
* Do not do any direct calls to 3rd party APIs or data sources, that is what Gateways and Repositories are for.

# Repository
* This is where your DB calls go
* These functions should only retrieve data, fill an object, and return it with any errors
* No logic should go here

# Gateway
* See Repository, but instead of a DB think of other sources on the internet

# Models
* These should be accessible anywhere
* Models should be relatively dumb, and only know about themselves
* A DB Model should likely have it's own Repository    
