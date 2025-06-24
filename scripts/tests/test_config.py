import psycopg2
import pytest   
from sqlalchemy import URL
from sqlalchemy import Connection, text
from scripts.load_CSV import getConfigDetails, getConnectionString, createEngine, establishConnection, uploadCSVToTable


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






def test_establish_connection():
    postgresUsername, postgresPassword, postgresDatabase, postgresHost, postgresPort = getConfigDetails()
    
    connectionString = getConnectionString(postgresUsername, postgresPassword, postgresDatabase, postgresHost, postgresPort)
    connection = None

    engine = createEngine(connectionString)
    try:
        connection = establishConnection(engine)
        assert connection is not None  
        assert isinstance(connection, Connection)  
        print("Connection established successfully.")
    finally:
            connection.close() 


def test_uploadCSVToTable():
    postgresUsername, postgresPassword, postgresDatabase, postgresHost, postgresPort = getConfigDetails()
    
    connectionString = getConnectionString(postgresUsername, postgresPassword, postgresDatabase, postgresHost, postgresPort)
    engine = createEngine(connectionString)
    connection = establishConnection(engine)

    uploadCSVToTable('test', 'data/test.csv', connection, 'test')

    
    query = text('SELECT * FROM test.test LIMIT 1;')
    result = connection.execute(query)

    want = [('Charlie', 21, "'Hello World'")]
    got = result.all()


    if got == want:
        assert True
        print("CSV uploaded successfully and data matches expected.")
    else:
        print(f"Expected {want} but got {got}")
        print(want, got)
        assert False
    connection.close()  # Ensure the connection is closed after the test
    