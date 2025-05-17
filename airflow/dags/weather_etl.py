from airflow import DAG
from airflow.operators.python import PythonOperator
from datetime import datetime, timedelta
import pandas as pd
from sqlalchemy import create_engine

default_args = {
    'owner': 'airflow',
    'start_date': datetime(2023, 1, 1),
    'retries': 1,
    'retry_delay': timedelta(minutes=5)
}

dag = DAG('weather_daily_etl', default_args=default_args, schedule_interval='@daily')

def export_to_csv():
    engine = create_engine("postgresql://postgres:password@postgres/weather")
    df = pd.read_sql("SELECT * FROM weather_data WHERE recorded_at::date = CURRENT_DATE - INTERVAL '1 day'", engine)
    df.to_csv("/usr/local/airflow/dags/output/weather_summary.csv", index=False)

task = PythonOperator(
    task_id='export_csv',
    python_callable=export_to_csv,
    dag=dag
)
