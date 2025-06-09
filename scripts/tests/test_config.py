import configparser
import psycopg2
import pytest   
from sqlalchemy import URL
from sqlalchemy import create_engine
from scripts.load_CSV import getConfigDetails, getConnectionString, createEngine


def test_config_loading():
    postgresUsername, postgresPassword, postgresDatabase, postgresHost, postgresPort = getConfigDetails()
    assert postgresUsername is not None
    assert postgresPassword is not None
    assert postgresDatabase is not None 
    assert postgresHost is not None
    assert postgresPort is not None


def test_postgres_connection():
    postgresUsername, postgresPassword, postgresDatabase, postgresHost, postgresPort = getConfigDetails()

    try:
        connection = psycopg2.connect(
            dbname=postgresDatabase,
            user=postgresUsername,
            password=postgresPassword,
            host=postgresHost,
            port=postgresPort
        )
        connection.close()
        assert True  # Connection successful
    except Exception as err:
        print(f"Connection failed: {err}")
        assert False



def test_connectionString():
    postgresUsername, postgresPassword, postgresDatabase, postgresHost, postgresPort = getConfigDetails()
    
    connectionString = str(getConnectionString(postgresUsername, postgresPassword, postgresDatabase, postgresHost, postgresPort))

    want = "postgresql+psycopg2://postgres:***@localhost:5432/Movie-data"

    if connectionString == want :
        assert True  # Connection string matches expected format
    else:
        print(f"Connection string does not match expected: {connectionString}")
        assert False # Connection string does not match expected format



def test_establish_connection():
    postgresUsername, postgresPassword, postgresDatabase, postgresHost, postgresPort = getConfigDetails()
    
    connectionString = getConnectionString(postgresUsername, postgresPassword, postgresDatabase, postgresHost, postgresPort)

    engine = createEngine(connectionString)
    try:
        engine.connect()
        print("Connection established successfully.")
        assert True  # Connection established successfully
    except Exception as err:
        print(f"Failed to establish connection: {err}")
        assert False # Connection failed      