import asyncio
from temporalio.client import Client
from temporalio.worker import Worker
from activities.activities_python import python_add_random_prefix_activity

async def main():
    client = await Client.connect("localhost:7233")
    task_queue = "python-task-queue"
    worker = Worker(
        client,
        task_queue=task_queue,
        activities=[python_add_random_prefix_activity],
    )
    print("Python worker started, polling on task queue:", task_queue)
    await worker.run()

if __name__ == "__main__":
    asyncio.run(main())