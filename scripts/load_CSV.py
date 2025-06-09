#import the libraries we are going to use 
import configparser
import pandas as pd
import numpy as np
from sqlalchemy import create_engine, URL, text




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


""""
# we can programmatically create the connection string using the URL class from SQLAlchemy.
connectionString = URL.create(
        "postgresql+psycopg2",
        username=postgres_username,
        password=postgres_password,
        host=postgres_host,
        port=int(postgres_port),
        database=postgres_database,
    )

engine = create_engine(connectionString)
connection = engine.connect()

"""

getConfigDetails()