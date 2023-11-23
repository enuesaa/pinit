FROM python:3.9

RUN apt update && apt install -y ffmpeg

RUN pip3 install poetry

WORKDIR /app
COPY . .
EXPOSE 3002
