# This application is solution for jumpcloud coding assignment.

### Instructions to Run
* Clone this project in your local machine at $HOME/go/src/github.com/<your_github_username>
* cd to directory jumpcloud-assignment # *cd jumpcloud-assignment*
* get all the packages # *go get .*
#### On local machine
* run your application # *go run .* Application is now ready to serve the requests.

#### With docker
* run *docker-compose up --build* from the application directory.

### APIs:
APIs can run on any host and port. Can be configured from config.txt file (for local) and docker-compose.yml file.
* POST http://localhost:8080/hash sends the request with form data. ex: password=anypassword. It can handle json payload as well.
  * returns the reference Id
* GET http://localhost:8080/hash/{id} Retrieves the hashed password from storage. Returns 404 if the requested id is not found
* GET http://localhost:8080/stats returns the performance stats of all post call
* GET http://localhost:8080/shutdown shutdowns the application gracefully sending interrupted signal
