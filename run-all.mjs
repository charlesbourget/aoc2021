#!/usr/bin/env zx
import "fs";
import "path";

$.verbose = false;

const re = /day\d+/g;

for await (const d of await fs.promises.opendir(".")) {
  if (d.isDirectory() && re.exec(d.name)) {
    const dir = d.name;
    let compile = await $`time go build -o ${dir}_bin ./${dir}`;
    let compile_time =
      await $`echo ${compile.toString()} | grep "real" | awk '{print $2}'`;
    console.log(chalk.yellow(`Compile ${dir}: ${compile_time.toString()}`));

    let run = await $`time ./${dir}_bin`;
    let run_time =
      await $`echo ${run.toString()} | grep "real" | awk '{print $2}'`;
    console.log(chalk.green(`Runtime ${dir}: ${run_time.toString()}`));
  }
}

// Clean up
await $`rm -f *_bin`;
