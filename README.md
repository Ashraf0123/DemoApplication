# temprepo
PostGres Demo Application 
A demo application with rest apis to connect with a postgres cluster running in KloverCloud platform. The application dynamically loads required environment variables required to connect with postgres and and perform read/write actions.

Application details
The application established connection with each endpoint on postgres endpoint. If it fails to connect with any instance, it will throw a fatal log and exit. All the write operations (POST, PUT, DELETE) are executed through the master instance and all the GET operations are executed through slave instances (round-robin)

Deploying application and cache
create a new vpc with enough resources to deploy the go application and postgres server. Assuming you already have added personal access token of github/gitlab on klovercloud. OnBoard the application from the create new drop down bar and after that edit the DockerFile as necessary, e.g. exposing the port where application is running.

Create a postgres in the same vpc, and deploy it. Yahoo! Nearly halfway there, with some mouse clicks right. No pesky terminal hassle. All we need to do is just inform the application about the database. The application already expecting one, so lets do it.
