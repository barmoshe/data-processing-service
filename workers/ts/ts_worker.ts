import { Worker } from '@temporalio/worker';
import { toUpperCaseActivity } from '../../activities/activities_ts.ts';

async function run() {
  const worker = await Worker.create({
    taskQueue: 'typescript-task-queue',
    activities: {
      // This registers the activity with the name expected by the workflow.
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