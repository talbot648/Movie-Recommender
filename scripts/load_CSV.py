#import the libraries we are going to use 
import configparser
import pandas as pd
import numpy as np
from sqlalchemy import create_engine, URL, text


config = configparser.ConfigParser()

config.read('config.ini')

postgres_username = config.get('postgres','username')
postgres_password = config.get('postgres','password')
postgres_database = config.get('postgres','db')
postgres_host = config.get('postgres','host')
postgres_port = config.get('postgres','port')

print(f'The username is "{postgres_username}" and the database is "{postgres_database}"')

# we can programmatically create the connection string using the URL class from SQLAlchemy.
conn_string = URL.create(
        "postgresql+psycopg2",
        username=postgres_username,
        password=postgres_password,
        host=postgres_host,
        port=int(postgres_port),
        database=postgres_database,
    )

