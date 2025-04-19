# Usage Instructions

### Download and Install Python

- https://www.python.org/downloads/

### Create virtual environment

`python -m venv venv`

### Activate virtual environment

`venv\Scripts\activate`

### Install requirements

`pip install -r requirements.txt`

### Set up .env file with the following

- DB_HOST
- DB_USER
- DB_PASSWORD
- DB_NAME
- PORT

### Setup a mySQL database using the script

- SQL_database_script.sql
- Make sure to use `create database <databaseName>` before running the rest of the script.

### Run API

`python app/app.py`
