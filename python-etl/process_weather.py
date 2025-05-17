from kafka import KafkaConsumer
import json
import psycopg2

consumer = KafkaConsumer(
    'weather',
    bootstrap_servers='localhost:9092',
    value_deserializer=lambda x: json.loads(x.decode('utf-8'))
)

conn = psycopg2.connect(
    dbname="weather", user="postgres", password="password", host="localhost"
)
cur = conn.cursor()

for msg in consumer:
    data = msg.value
    cur.execute("""
        INSERT INTO weather_data (location, temperature, recorded_at)
        VALUES (%s, %s, %s)
    """, (data['location'], data['temp_c'], data['time']))
    conn.commit()
    print("Inserted:", data)
