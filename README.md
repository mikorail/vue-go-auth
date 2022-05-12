# vue-golang-auth
## a simple app developed using Vue Js + Golang
Golang modules :
1. github.com/badoux/checkmail for mail checking
2. github.com/golang-jwt/jwt  for JWT tokens
3. github.com/gorilla/mux for HTTP routers
4. github.com/jinzhu/gorm for Object Relation Mapping (used in controllers)
5. github.com/joho/godotenv for importing .env files
6. golang.org/x/crypto for password encryption

Vue Js modules :
1. Vuetify for design pattern theme and components library
2. Vuex for state management
3. Vue-Router for routing and routes stuff

# Project Setup 
### Database
you can make new database and import the `.sql` file in `#importdb`
after that you can set up `.env` file. to the database you created or imported
if theres no `.env` file around try make `.env`file in the project folder and copy this code to your new `.env` file content

`API_SECRET=98hbun98h #Used when creating a JWT. It can be anything
DB_HOST=127.0.0.1
DB_DRIVER=mysql
DB_USER=root
DB_PASSWORD=
DB_NAME=golang
DB_PORT=3306`

### Go
setting up the go you can start by using this command in the terminal 
`go run main.go`

### Vue
navigate to `app` folder then you can start typing in your terminal using
`npm install`
wait for the process to be completed, and run 
`npm run serve`
