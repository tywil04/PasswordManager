# Password Manager
A password manager written using Golang (Gin framework) and SvelteKit for the Backend and Frontend respectively.

This is for my Computer Science NEA, this code has not been security audited and shouldn't be used for any application which requires security in any form whatsoever.

## Building and Running
### Prerequisites
To build and run this project, you need to have the following installed:
- `nodejs`: This is required to compile Svelte Frontend.
- `yarn`: This is the tool used to install packages for the Svelte Frontend.
- `golang`: This is the library for Go which means you can build/run this project.
- `git`: Required to clone this repo.

### Building
```
git clone https://github.com/tywil04/PasswordManager      # Clone this repo
cd PasswordManager                                        # Enter the cloned repo 
go generate ./...                                         # Generate DB and compile Svelte Frontend
go build server.go                                        # Build the project into a single executable
```
Once the project has been built, you run it simply by running the executable (provided you are in the same directory you built the project from): `./server`.

### Running
Instead of building, its possible to run this Go project, run the `Building` commands but replace:
```
go generate ./...      # Generate DB and compile Svelte Frontend
go build server.go     # Build the project into a single executable
```
with
```
go run server.go
```
(this builds the project in a temporary location and runs it).

## Environment Variables
Environment variables are used to store configuration data. Heres a list of variables that can be modified. This variables can be set with a `.env` file in the same directory as the executable.

if you do not include a required environment variable, things might not work as expected.

### General (Required)
- `DB_PATH`: This is the path for your db, this project uses sqlite3. Example value: `./ent/dev.db`.
- `SERVER_ADDRESS`: This how Go binds the server, its expected in a format `HOSTNAME:PORT`. Example value: `0.0.0.0:8001` (Allow connection from any interface on port 8001).

### Email (Required)
- `SMTP_HOST`: This is the host for the SMTP server to use to send email addresses. Example value: `smtp.example.org`.

- `SMTP_PORT`: This is the port for the SMTP server to send email addresses. Example value: `587`.

- `SMTP_FROM`: This is the email address that the message should be from. Example value: `noreply@example.org`.

- `SMTP_USERNAME`: This is the username to access your SMTP server.

- `SMTP_PASSWORD`: This is the password to access your SMTP server.

### WebAuthn (Required)
- `RP_DISPLAY_NAME`: This is the display name for a relying party..

- `RP_ID`: This is the relying partys origin domain. Example value: `localhost`.

- `RP_ORIGINS`: This is a comma seperated list of allowed origins. Example value: `http://localhost:8080,https://example.org`.

- `RP_ICON`: This is a URL to the icon for a relying party. This is optional, a blank value can be set.

## Todo
Stuff that needs to be done

### Frontend
- [ ] Signin Page
    - Need to fix email verification code box
- [ ] Signup Page
    - Need to fix email verification code box
- [ ] Home Page
    - Needs to decrypt password names and display them in a list
- [ ] Settings Page
    - Needs to exist
        - Needs to allow for the registration and revoking on WebAuthn devices
        - Needs to allow for the registration and removal of TOTP secrets
        - Needs to allow the user to select accesibility features like a High-Contrast filter.

### Backend
- [ ] Document API
- [X] Auth API (/api/v1/auth/*)
- [ ] WebAuthn API (/api/v1/webauthn/*)
    - Need to implement the revoking of credentials
- [ ] TOTP API (/api/v1/totp/*)
    - Need to implement register endpoint
    - Need to implement signinChallenge endpoint
- [X] Email API (/api/v1/email/*)
- [ ] Passwords API
    - Need to implement updating passwords
- [X] Email
- [X] Email templating