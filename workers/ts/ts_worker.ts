// workers/ts/ts_worker.ts

import { Worker } from '@temporalio/worker';
// Import the activity from the processing folder.
// Adjust the relative path based on your folder structure.
import { toUpperCaseActivity } from '../../processing/activities_ts';

async function run() {
  // Create a Worker that listens on the "typescript-task-queue"
  const worker = await Worker.create({
    taskQueue: 'typescript-task-queue',
    activities: {
      // Register the activity using the name expected by the Workflow.
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

//run with ```npx tsx workers/ts/ts_worker.ts```
