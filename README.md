# Booking prototype

This project solely for the purpose of learning docker and handling synchronous problem in ticket booking system

![](https://media.giphy.com/media/26gN0Ix4vuv2Ffkl2/giphy.gif)

### Prerequisite

* Docker CE
* Docker-compose

### Installation

```
git clone https://github.com/vanhtuan0409/prototype-booking-system.git
cd prototype-booking-system
docker-compose up
```

**Note**: After bootstraping docker, you might encounter a problem that API server failed to start. It is because API server started before PostgreSQL finished bootstraping. This can be fixed by stop docker-compose (but don't remove the container) and run docker-compose up again

After running docker successfully, you can go to `localhost/db-manager` and select the following config to view database:

```
System: PostgreSQL
Server: db
Username: booking
Password: password
Database: booking
```

After connect successfully to database dashboard, you can copy contents of file `seed.sql` to insert some resources.

Now your system is ready and running at `localhost`

### Services

There are 3 main service in prototype

#### Frontend service

Just a simple http server that serve React bundle app.

#### Live notification service

A simple Socket.io and Express server, in charge of notify all connected clients when resource booking event occured

#### REST API service

A simple REST API written in Go, which serve the following apis:

* `GET /resources`: return list of all resources and their status (booked or available)
* `POST /resources/:id/book`: booked a resource with provided id
* `POST /resources/restore`: restore all resources to available

When ever a resources was booked successfully or all resources was restored, API service will send to message to Live notification service by `a http request` (I know it is simple and have some drawbacks). Then Live notification service will notify all connected clients to update Web UI.

### How to prevent duplicate booking

Currently, I use database row locking `SELECT ... FOR UPDATE` each resource can only be booked once. This might cause a bottleneck on database access when there is a large number of request.

If you have better solutions, please help me by open an issue for discussion

### Docker and Reversed Proxy

All of the services and database was designed to run in Docker container. On top of that, I use [Traefik](https://traefik.io/) as a Reversed Proxy.

* `:8080`: Point to Traefik dashboard, in which you can review backend services and route mapping
* `/db-manager`: Point to adminer container, serve as a Database dashboard
* `/subcribe/*`: Point to Live notification service
* `/api/*`: Point to REST API service
* `/*` (default): Point to Frontend React app
