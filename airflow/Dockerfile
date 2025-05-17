FROM python:3.13-alpine

ENV AIRFLOW_VERSION=2.7.1

RUN apk add --no-cache \
        build-base \
        libffi-dev \
        openssl-dev \
        curl \
        bash \
        postgresql-dev \
        && pip install --no-cache-dir "apache-airflow==${AIRFLOW_VERSION}" \
        && apk del build-base

RUN mkdir -p /opt/airflow/dags
COPY dags /opt/airflow/dags

ENV AIRFLOW_HOME=/opt/airflow
WORKDIR $AIRFLOW_HOME

CMD ["airflow", "scheduler"]
