# workers/python/python_worker.py

import asyncio
from temporalio.client import Client
from temporalio.worker import Worker
# Import the activity from the processing folder.
from processing.activities_python import python_add_random_prefix_activity

async def main():
    # Connect to the local Temporal server.
    client = await Client.connect("localhost:7233")
    # Use the same task queue as expected by the workflow.
    task_queue = "python-task-queue"
    # Create the Worker with the imported activity.
    worker = Worker(
        client,
        task_queue=task_queue,
        activities=[python_add_random_prefix_activity],
    )
    print("Python worker started, polling on task queue:", task_queue)
    await worker.run()

if __name__ == "__main__":
    asyncio.run(main())



#run with ```python3 -m workers.python.python_worker```