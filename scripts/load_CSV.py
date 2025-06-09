#import the libraries we are going to use 
import configparser
import pandas as pd
import numpy as np
from sqlalchemy import create_engine, URL, text

def main():

    postgresUsername, postgresPassword, postgresDatabase, postgresHost, postgresPort = getConfigDetails()
    connectionString = getConnectionString(postgresUsername, postgresPassword, postgresDatabase, postgresHost, postgresPort)
    engine = createEngine(connectionString)


def getConfigDetails():
    config = configparser.ConfigParser()
    config.read('scripts/config.ini')
    postgresUsername = config.get('postgres','username')
    postgresPassword = config.get('postgres','password')
    postgresDatabase = config.get('postgres','db')
    postgresHost = config.get('postgres','host')
    postgresPort = config.get('postgres','port')

    print(f'The username is "{postgresUsername}" and the database is "{postgresDatabase}"')

    return postgresUsername, postgresPassword, postgresDatabase, postgresHost, postgresPort


def getConnectionString(postgresUsername, postgresPassword, postgresDatabase, postgresHost, postgresPort):
# we can programmatically create the connection string using the URL class from SQLAlchemy.
    connectionString = URL.create(
        "postgresql+psycopg2",
        username=postgresUsername,
        password=postgresPassword,
        host=postgresHost,
        port=int(postgresPort),
        database=postgresDatabase,
    )
    print(f'The connection string is "{connectionString}"')
    return connectionString

def createEngine(connectionString):
    # Create an engine instance
    engine = create_engine(connectionString)
    print(f'engine: "{engine}"')
    return engine

 

main()