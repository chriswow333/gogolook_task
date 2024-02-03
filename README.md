
# Task api application


## How to run 
    local run
        # go run main.go
    docker build
        # docker build -t gogolook .
        # docker run -p 8080:8080 gogolook

## Method
    Get all tasks
      curl: http://localhost:8080/tasks (GET)
    Create a task
      curl http://localhost:8080/tasks (POST)
      payload: 
      
        {
            "id":"testid",
            "name":"testname",
            "status":0,
            "memo":""
        }
      
    Update a task
      curl http://localhost:8080/tasks/testid (PUT)
      payload: 
      
        {
            "name":"testname",
            "status":1,
            "memo":"done"
        }
      

    Delete a task
      curl http://localhost:8080/tasks/testid (DELETE)


