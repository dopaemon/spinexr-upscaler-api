import sys
import os
import json
import mimetypes
from datetime import datetime

def get_file_info(file_path):
    if not os.path.exists(file_path):
        return {"error": "File không tồn tại"}

    file_info = os.stat(file_path)
    mime_type, _ = mimetypes.guess_type(file_path)

    return {
        "filename": os.path.basename(file_path),
        "size": file_info.st_size,
        "created": datetime.fromtimestamp(file_info.st_ctime).isoformat(),
        "type": mime_type or "application/octet-stream"
    }

if __name__ == "__main__":
    file_path = sys.argv[1]
    print(json.dumps(get_file_info(file_path)))
