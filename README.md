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

### Developing
To develop the application, you must have gotten to a stage where you can run the application. The best way to develop is to start a UI-less instance of the golang server with 
```
DISABLE_UI=true go run server.go
``` 
and then in a new terminal run
```
cd ui/
yarn dev
```
Connect to the vite server from `yarn dev`, this gives you hot-reloading so when you make the change to part of the ui, the web browsers instantly reflects this change. The backend server is running and is getting the `/api` requests from the vite server because vite has been configured to proxy api requests.

## Environment Variables
Environment variables are used to store configuration data. Heres a list of variables that can be modified. This variables can be set with a `.env` file in the same directory as the executable.

All environment variables are required unless explicitly mentioned otherwise, if you do not include a required environment variable, things might not work as expected.

### General
- `ENVIRONMENT` *(optional, default is production)*: This is used to determin the environment the application should be running in. Allowed options are `production` and `development`.

- `ALLOWED_ORIGINS` *(optional, default is all origins)*: This is a comma seperated list of allowed origins (this is also used to define allowed origins for the webauthn relying party). e.g `http://localhost:8080,http://localhost:5173`.

- `DB_PATH`: This is the path for your db, this project uses sqlite3. Example value: `./ent/dev.db`.

- `SERVER_ADDRESS`: This how Go binds the server, its expected in a format `HOSTNAME:PORT`. Example value: `0.0.0.0:8001` (Allow connection from any interface on port 8001).

- `DISABLE_API` *(optional, default is false)*: This is used to disable the API section of the webserver.

- `DISABLE_UI` *(optional, default is false)*: This is used to disable the UI section of the webserver.

### Email
- `SMTP_HOST`: This is the host for the SMTP server to use to send email addresses. Example value: `smtp.example.org`.

- `SMTP_PORT`: This is the port for the SMTP server to send email addresses. Example value: `587`.

- `SMTP_FROM`: This is the email address that the message should be from. Example value: `noreply@example.org`.

- `SMTP_USERNAME`: This is the username to access your SMTP server.

- `SMTP_PASSWORD`: This is the password to access your SMTP server.

### WebAuthn
- `RP_DISPLAY_NAME`: This is the display name for a relying party..

- `RP_ID`: This is the relying partys origin domain. Example value: `localhost`.

- `RP_ICON`: This is a URL to the icon for a relying party. This is optional, a blank value can be set.

### Crypto
- `CRYPTO_PEPPER` *(optional, default is blank)*: This is a random string that if set will be added to every `masterHash` before being strengthened for added security. It can technically be any string however a random string will be most beneficial.