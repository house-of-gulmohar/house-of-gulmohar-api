# House Of Gulmohar API
APIs for House of Gulmohar Ecommerce app is written in golang and uses these modules 
- **go-chi** - [github.com/go-chi/chi/v5](https://pkg.go.dev/github.com/go-chi/chi/v5) (routing)
- **cors** - [github.com/go-chi/cors](https://pkg.go.dev/github.com/go-chi/cors) (cors)
- **UUID** - [github.com/gofrs/uuid/v3](https://pkg.go.dev/github.com/gofrs/uuid/v3) (uuid)
- **pgx** - [github.com/jackc/pgx/v4](https://pkg.go.dev/github.com/jackc/pgx/v4) (postgres)
- **godotenv** - [github.com/joho/godotenv](https://pkg.go.dev/github.com/joho/godotenv) (.env)
- **logrus** - [github.com/sirupsen/logrus](https://pkg.go.dev/github.com/sirupsen/logrus) (logging)

## sample env
> APP_PORT=  
> DB_USERNAME=  
> DB_PASSWORD=   
> DB_HOST=  
> DB_PORT=  
> DB_DATABASE= 

*Note: you need these environmental variables with values present in a .env file to run this app*