import os
import zipfile
from datetime import datetime, timezone

# convert utc datetime to local datetime without timeozone
def utc_to_local(utc_dt):
    return utc_dt.replace(tzinfo=timezone.utc).astimezone(tz=None).replace(tzinfo=None)


# convert local datetime to utc datetime without timeozone
def local_to_utc(dt):
    return dt.astimezone(timezone.utc).replace(tzinfo=None)


def isoformat(dt):
    return dt.isoformat()


def short(n):
    levels = [
        (None, 1000.0),
        ("K", 1000.0),
        ("M", float("inf")),
    ]

    for _label, scale in levels:
        if n >= scale:
            n /= scale

        else:
            break

    if _label:
        return f"{n:0.2f}{_label}"
    return str(n)

# 
def timeago(utc_dt):
    delta = round((datetime.utcnow() - utc_dt).total_seconds() * 1000)
    levels = [
        ("us", 1),
        ("ms", 1000),
        ("s", 60),
        ("m", 60),
        ("h", 24),
        ("d", 7),
        ("w", 4),
        ("M", 12),
        ("Y", float("inf")),
    ]

    for _label, scale in levels:
        if abs(delta) > scale:
            delta //= scale

        else:
            break

    if delta < 0:
        return f"in {abs(delta)}{_label}"

    else:
        return f"{delta}{_label} ago"



# startdir = '/home/lei/Documents/Notes'  # 要压缩的文件夹路径
# zipname = '/home/lei/Documents/Notes/complete.zip'  # 压缩后文件夹的名字

def compress_zip(zipname, startdir):
    z = zipfile.ZipFile(zipname, 'w', zipfile.ZIP_DEFLATED)  # 参数一：文件夹名
    for dirpath, dirnames, filenames in os.walk(startdir):
        fpath = dirpath.replace(startdir, '')
        # print(fpath)
        fpath = fpath and fpath + os.sep or ''
        # print(fpath)
        for filename in filenames:
            z.write(os.path.join(dirpath, filename), fpath+filename)
    print ('压缩成功')
    z.close()

# compress_zip(zipname, startdir)

