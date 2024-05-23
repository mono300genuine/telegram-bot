import sys
import time
from watchdog.observers import Observer
from watchdog.events import FileSystemEventHandler
import subprocess
import os

class ChangeHandler(FileSystemEventHandler):
    def __init__(self, script_name):
        self.script_name = script_name
        self.process = None
        self.restart_bot()

    def restart_bot(self):
        if self.process:
            self.process.terminate()
        self.process = subprocess.Popen([sys.executable, self.script_name])

    def on_modified(self, event):
        if event.src_path.endswith(self.script_name):
            print(f"{self.script_name} modified; restarting bot...")
            self.restart_bot()

if __name__ == "__main__":
    script_name = "bot.py"  # Replace with the name of your bot script
    event_handler = ChangeHandler(script_name)
    observer = Observer()
    observer.schedule(event_handler, path=os.path.dirname(os.path.abspath(script_name)), recursive=False)
    observer.start()
    try:
        while True:
            time.sleep(1)
    except KeyboardInterrupt:
        observer.stop()
    observer.join()
