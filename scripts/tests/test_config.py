import configparser
import psycopg2
import pytest   

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