FROM python:3

WORKDIR /usr/src/app

COPY . .

RUN pip install psycopg2-binary
CMD [ "python", "./import-data.py" ]