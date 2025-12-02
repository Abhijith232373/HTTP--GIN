# <mark>Web Development in Go (Gin Framework)</mark>

***

# ***Day 1: HTTP & REST API Fundamentals***

- [x] Create a simple Go server that handles `GET` and `POST`.

- [x] Implement `/api/hello` returning JSON (`{"message": "Hello, World!"}`).

- [x] Create `/api/user` endpoint accepting `POST` JSON data (`name`, `age`).

## ***Day 2: Gin Framework & Routing***

- [x] Build a Gin server that serves a homepage.

- [x] Add a `/api` route group for REST APIs.

- [x] Implement `/users` API supporting `GET` & `POST`.

## ***Day 3: Sessions, Cookies & Middleware***

- [x] Implement a login endpoint validating hardcoded username/password.

- [x] Store a session cookie on login, and keep user logged in until sign out.

- [ ] Create logging middleware to log all requests to login/logout routes.

- [ ] Implement authentication middleware to protect certain routes

## ***Day 4: Validation, Authentication & Security***

- [ ] Validate user input (`empty fields`, `invalid format`).

- [ ] Hash passwords before storing (use `golang.org/x/crypto/bcrypt`).

- [ ] Implement a **protected dashboard route** accessible only to logged-in users

## ***Day 5: Testing***

- [ ] Write unit tests for session & cookie handling.

- [ ] Write tests for authentication logic.

## ***Project***

- Build a complete login system using Gin
- Add middleware for logging & authentication
- Implement session management using cookies
- Create a protected dashboard
- Push the project to GitHub with README
