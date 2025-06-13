#import the libraries we are going to use 
import configparser
import pandas as pd
import numpy as np
from sqlalchemy import create_engine, URL, text

def main():

    postgresUsername, postgresPassword, postgresDatabase, postgresHost, postgresPort = getConfigDetails()
    connectionString = getConnectionString(postgresUsername, postgresPassword, postgresDatabase, postgresHost, postgresPort)
    engine = createEngine(connectionString)
    connection = establishConnection(engine)
    print(connection)
    uploadCSVToTable('movies_metadata', 'data/movies_metadata.csv', connection,'dbo')



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

def establishConnection(engine):
    return engine.connect()

def uploadCSVToTable(tableName, csvFilePath, connection, schema):
    df = pd.read_csv(csvFilePath)


    df = df.replace({np.nan: None})
    df.to_sql(tableName, con=connection, if_exists='replace', index=False, schema=schema)

    print(f'Data from {csvFilePath} uploaded to {schema}.{tableName} successfully.')

    preview_df = pd.read_sql(f'SELECT * FROM "{schema}"."{tableName}" LIMIT 5;', con=connection)
    print("Sample rows from the table:")
    print(preview_df)
    connection.close()




  
main()

