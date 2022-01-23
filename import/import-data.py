import csv
import psycopg2
from pathlib import Path
import os

conn = psycopg2.connect(dbname=os.environ.get("APP_DB_NAME"), user = os.environ.get("APP_DB_USER"), 
    password = os.environ.get("APP_DB_PASS"), host = os.environ.get("APP_DB_HOST"), port = os.environ.get("APP_DB_PORT"))
cur = conn.cursor()
cwd = Path(__file__)

tables = {'customer_companies':2,'customers':6,'orders':4,'order_items':5,'deliveries':3}

for key, value in tables.items():   
    columns ='' 
    for col in range(0,value):
        columns += '%s'+ ','

    columns = columns.rstrip(',')
                
    table_address = cwd.parent / f'{key}.csv'
    with open(table_address, 'r') as f:
        reader = csv.reader(f)
        next(reader)
        for row in reader:            
            for n,x in enumerate(row):
                if x == '':
                    row[n] = None

            cur.execute(f'INSERT INTO {key} VALUES ({columns})', row)

conn.commit()