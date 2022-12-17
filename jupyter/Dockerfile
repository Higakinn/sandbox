FROM jupyter/datascience-notebook
ENV TZ=Asia/Tokyo
USER root
RUN apt-get update \
  && apt-get -y install libpq-dev python-dev \
  && pip install psycopg2-binary pymysql \
  && apt-get clean
