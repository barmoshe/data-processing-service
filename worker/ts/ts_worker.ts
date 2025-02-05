import { Worker } from '@temporalio/worker';

// Define the activity function that uppercases the string.
export async function toUpperCaseActivity(data: string): Promise<string> {
  return data.toUpperCase();
}

async function run() {
  // Create a Worker that listens on the "typescript-task-queue"
  const worker = await Worker.create({
    taskQueue: 'typescript-task-queue',
    activities: {
      // Register the activity using the name the Workflow expects.
      TypeScriptToUppercaseActivity: toUpperCaseActivity,
    },
  });
  console.log("TypeScript worker started, polling on task queue: 'typescript-task-queue'");
  await worker.run();
}

run().catch(err => {
  console.error(err);
  process.exit(1);
});
