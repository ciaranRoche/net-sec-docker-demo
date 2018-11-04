FROM python:2.7-alpine
LABEL maintainer ciaran
COPY . /app
WORKDIR /app
RUN pip install -r requirements.txt
EXPOSE 5000
ENTRYPOINT [ "python" ]
CMD ["application.py"]