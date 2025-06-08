import configparser
import psycopg2
import pytest   
from sqlalchemy import URL

def test_config_loading():
    config = configparser.ConfigParser()
    config.read('config.ini')
    
    assert config.has_section('postgres')
    assert config.get('postgres', 'username') is not None
    assert config.get('postgres', 'password') is not None
    assert config.get('postgres', 'db') is not None
    assert config.get('postgres', 'host') is not None
    assert config.get('postgres', 'port') is not None

def test_postgres_connection():
    config = configparser.ConfigParser()
    config.read('config.ini')
    
    postgres_username = config.get('postgres', 'username')
    postgres_password = config.get('postgres', 'password')
    postgres_database = config.get('postgres', 'db')
    postgres_host = config.get('postgres', 'host')
    postgres_port = config.get('postgres', 'port')

    try:
        connection = psycopg2.connect(
            dbname=postgres_database,
            user=postgres_username,
            password=postgres_password,
            host=postgres_host,
            port=postgres_port
        )
        connection.close()
        assert True  # Connection successful
    except Exception as e:
        print(f"Connection failed: {e}")
        assert False

def test_connectionString():
    config = configparser.ConfigParser()
    config.read('config.ini')

    postgres_username = config.get('postgres', 'username')
    postgres_password = config.get('postgres', 'password')
    postgres_database = config.get('postgres', 'db')
    postgres_host = config.get('postgres', 'host')
    postgres_port = config.get('postgres', 'port')
    
    connectionString = URL.create(
        "postgresql+psycopg2",
        username=postgres_username,
        password=postgres_password,
        host=postgres_host,
        port=int(postgres_port),
        database=postgres_database,
    )

    connectionString = str(connectionString)
    want = "postgresql+psycopg2://postgres:***@localhost:5432/Movie-data"

    if connectionString == want :
        assert True  # Connection string matches expected format
    else:
        print(f"Connection string does not match expected: {connectionString}")
        assert False
    