from celery import Celery

app = Celery('tasks', backend='redis://0.0.0.0', broker='pyamqp://admin:admin@127.0.0.1//')

app.conf.update(
    task_serializer='json',
    accept_content=['json'],  # Ignore other content
    result_serializer='json',
    timezone='Asia/Shanghai',
    enable_utc=True,
)

@app.task
def add(x, y):
    return x + y

# celery -A tasks worker --loglevel=INFO