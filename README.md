# Order Viewer App

## Project Structure
The project consists of the following folders:
- db: Contains the initialization script to create the database and tables in PostgreSQL
- import: Contains the csv test data files and a Python script to read and import them into the database.
- api: Contains the GoLang project implementing the REST API for reading orders list.
- nginx: Contains the configuration file to configure NGINX serving front-end static files at /orders while reverse-proxying requests to the REST API via /api/v1/orders.

All project parts are dockerized. For *import* and *api*, I have built custom docker images, while for *db* and *nginx* I just configure and use their official images to do the job. There is also a docker compose file in the project root folder, which starts all the required parts to run and test the project.

## How to Run It

In order to build and run the project, you need either Linux or Windows with docker already installed. Having these requirements ready, you just need to execute the following command:

```
./build-and-run.sh
```
