1. config file (list of jobs / schedule) - environment variable to point to location (JOBSCHED=c:\AUTOJOB\)   jobsched.dat

2. go routine to load config file at startup, and if it changes reload jobs to map, re-calc next-run-time

    read config entries and parse them to (map of) jobs - job name, schedule, next-run-time, etc.

	loop forever through map of jobs and at if time.Now() >= next-run-time, launch go routine to "do" job.   set new next-run-time just prior to launching the go routine

	permanent "heartbeat" job.

	Logging
