## Feature Overview

- API to assign points to ready and nearest agent
- Implement Queue to support Async messaging
- The point will be requeue to ```point channel``` after 1 second if no agent is ready

## Description

The Project is simply assigning points to agents .

agents will start at point(0,0) when application starts 
the number of agents can be modified using ```START_AGENTS_COUNT``` ENV

The speed of agents is defined as ```CONST speedInMilliSecondsPerMeter``` in center service




### Installation
```makefile
make run
```