
from celery import Celery

app = Celery('proj',
             backend='redis://0.0.0.0',
             broker='pyamqp://guest:guest@127.0.0.1//',
             include=['proj.tasks'])

# Optional configuration, see the application user guide.
app.conf.update(
    task_serializer='json',
    accept_content=['json'],  # Ignore other content
    result_serializer='json',
    timezone='Asia/Shanghai',
    # enable_utc=True,
    result_expires=3600,
)

if __name__ == '__main__':
    app.start()
    
    # celery -A proj worker -l INFO