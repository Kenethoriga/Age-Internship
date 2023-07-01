import json
import psycopg2

conn_string = "dbname='your_database' user='your_username' password='your_password' host='localhost' port='5432'"
conn = psycopg2.connect(conn_string)
cursor = conn.cursor()

cursor.execute("SELECT * FROM public.user_table")
rows = cursor.fetchall()

columns = [desc[0] for desc in cursor.description]
data = []
for row in rows:
    entry = {}
    for i, col in enumerate(columns):
        val = row[i]
        entry[col] = val.decode() if isinstance(val, bytes) else val
    data.append(entry)

result = {
    "driver": {
        "status_code": 200,
        "data": data,
    },
}

json_result = json.dumps(result)
print(json_result)
