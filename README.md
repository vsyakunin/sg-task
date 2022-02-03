# Saiga home assignment

This is a simple tasks API. It is fully dockerized.

## Table of contents

**[1 Response Data format](#1)**   
**[2 Authorization](#2)**
**[3 API](#3)**   
&emsp;**[3.1 Get all tasks](#3.1)**  
&emsp;**[3.2 Get task history](#3.2)**  
&emsp;**[3.3 Download file attached to message](#3.3)**  
**[4 Requirements](#4)**  
**[5 Deploy](#5)**

****
<a name="1">1. Response data format</a>
-  
**Response data format:** JSON

**Sample error response:**
```js
{
    "title": "invalid parameter", 
    "description": "task ID parameter must be positive"
}
```

****
<a name="2">2. Authorization</a>
-  
**Authorization:** Basic Auth  

In this task we implement role based access control with basic auth in sake of simplicity.  
Although there're some libraries providing a better way of handling RBAC.  
For example casbin: https://github.com/casbin/casbin

****
## <a name="3">3. API</a>

### <a name="3.1">3.1 Get all tasks</a>

This endpoint is used to get a list of all tasks.

**URL:**

&emsp;`.../api/v1/tasks`

**METHOD: GET**

**Status codes:**  
&emsp;`200` - OK,  
&emsp;`500` - server error,

**Response Body**
```js
[
    {
        "id": 1,                                             // id of a task, int
        "user_id": 1,                                        // id of user who created a task, int
        "name": "task one",                                  // task name, string
        "category_id": "find a doctor",                      // id of a task category, int
        "status_id": 1,                                      // id of a task status, int
        "started_at": "2022-02-02T19:41:32.213348121+03:00", // date when a task started, datetime
        "updated_at": "2022-02-02T19:41:32.21334819+03:00"   // date when a task was last updated, datetime
    },
    ...
]
```

### <a name="3.2">3.2 Get task history</a>

This endpoint is used to receive chat history for a task.

**URL:**

&emsp;`.../api/v1/tasks/{taskID}/history`

**METHOD: GET**

**URL parameters:**  
&emsp;`taskID` - *int, id of a task*

**Status codes:**  
&emsp;`200` - OK,  
&emsp;`400` - bad request (invalid request body)  
&emsp;`500` - server error,

**Response Body**

```js
[
    {
        "id": 1,                                              // id of a message, int
        "task_id": 1,                                         // if of a task to which the message is related, int
        "from_user_id": 1,                                    // id of a user who wrote the message, int
        "message_text": "message one",                        // contents of message, string
        "file_key": "fileone.txt",                            // s3 key to a file attached to the message, string
        "created_at": "2022-02-02T19:52:56.115236415+03:00"   // date when the message was created, datetime
    }, 
    ...
]
```

### <a name="3.3">3.3 Download file attached to message</a>

This endpoint is used to download file attached to message.

**URL:**

&emsp;`.../api/v1/messages/{messageID}/history`

**METHOD: GET**

**URL parameters:**  
&emsp;`messageID` - *int, id of a message*

**Status codes:**  
&emsp;`200` - OK,  
&emsp;`400` - bad request (invalid request body)  
&emsp;`500` - server error,

**Response Body**  
The file that was attached to the message or nothing if there was no file attached.

## <a name="4">4. Requirements</a>

&emsp; `docker`  
&emsp; `docker-compose`

## <a name="5">5. Deploy</a>

1. Clone this repository.
2. Create container with `docker-compose -f docker-compose.yml up -d`
3. Server will start running. 