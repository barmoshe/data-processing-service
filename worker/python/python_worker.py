import random
import asyncio
from temporalio import activity
from temporalio.client import Client
from temporalio.worker import Worker

# Define the activity with a custom name that matches the one used in the Go Workflow.
@activity.defn(name="PythonAddRandomPrefixActivity")
async def python_add_random_prefix_activity(data: str) -> str:
    prefixes = ["alpha-", "beta-", "gamma-", "delta-", "epsilon-", "zeta-","eta-", "theta-", "iota-", "kappa-", "lambda-", "mu-", "nu-", "xi-", "omicron-", "pi-", "rho-", "sigma-", "tau-", "upsilon-", "phi-", "chi-", "psi-", "omega-"]
    prefix = random.choice(prefixes)
    return f"{prefix}{data}"

async def main():
    # Connect to the local Temporal server.
    client = await Client.connect("localhost:7233")
    # Use the same task queue as in your Go code.
    task_queue = "python-task-queue"
    # Start a Worker for the specified task queue with the Python activity.
    worker = Worker(
        client,
        task_queue=task_queue,
        activities=[python_add_random_prefix_activity],
    )
    print("Python worker started, polling on task queue:", task_queue)
    await worker.run()

if __name__ == "__main__":
    asyncio.run(main())
