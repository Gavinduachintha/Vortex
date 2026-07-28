from dotenv import load_dotenv
import os
import psycopg2

load_dotenv()

db_password = os.getenv("CROCORACH_DB_PASSWORD")
db_url = f"""postgresql://gavindu:{db_password}@silica-gnoll-30450.j77.aws-ap-south-1.cockroachlabs.cloud:26257/defaultdb?sslmode=verify-full"""


def get_connection():
    conn = psycopg2.connect(db_url)
    return conn

