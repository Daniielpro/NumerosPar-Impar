# ODD AND EVEN NUMBERS

## Description
This project is a Go application that allows to know if a number entered in the form is odd or even... 

## Prerequisites
- Go 1.23.3 or higher
- Docker (optional, if you want to run the project in a container)

## Installation
To install the project, follow these steps:

1. Clone the repository:
````bash
git clone https://github.com/Daniielpro/NumerosPar-Impar.git

## Local Execution
To run the project locally, follow these steps:

1. Open terminal from the root folder of the project.

2. Issue the following command
   ````bash
   go run main.go

3. Open the terminal: http://localhost:8084/

## Running Dcoker
To run the project using Docker, follow these steps:

1. Open terminal from the root folder of the project.

2. Issue the following command
   ````bash
     docker build -t username/appname .

3. Issue the following command
   ````bash
    docker run -d -p 8084:8084 username/nameapp
