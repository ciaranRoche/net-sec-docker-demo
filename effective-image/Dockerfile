FROM python:2.7-alpine
LABEL maintainer ciaran
COPY . /app
WORKDIR /app
RUN pip install -r requirements.txt
CMD ["python", "application.py"]