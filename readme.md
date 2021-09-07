# This application is solution for jumpcloud coding assignment.

#### Instructions to Run
* Checkout this project in your local machine at $HOME/go/src/github.com/<your_github_username>
* cd to directory jumpcloud-assignment # *cd jumpcloud-assignment*
* get all the packages # *go get .*
* make sure you have Postgres db with configuration defined on config.yaml file. You can change the config values.
* run your application # *go run .* Application is now ready to serve the requests.

#### APIs:
APIs can run on any host and port. Can be configured from config.yaml file.
* POST http://localhost:8080/hash sends the request with payload. ex: password=anypassword
  * returns the reference Id
* GET http://localhost:8080/hash/{id} Retrieves the hashed password from storage
* GET http://localhost:8080/stats returns the performance stats of post call
* GET http://localhost:8080/shutdown shutdowns the application gracefully sending interrupted signal

#### Library Dependencies:
* "github.com/gin-gonic/gin" to define server and apis
* "github.com/go-pg/pg/v10" for handling postgres connection
* "github.com/go-pg/pg/v10/orm" for postgres orm
* "github.com/spf13/viper" to load and use configuration
* "golang.org/x/crypto/bcrypt" to encrypt the password
* "context"
* "log"
* "math"
* "net/http"
* "os"
* "os/signal"
* "strings"
* "syscall"
* "time"
* "github.com/thoas/stats" to manage stats

#### Application Dependencies
* Postgres db
