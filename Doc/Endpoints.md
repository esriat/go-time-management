# Endpoints

## Users

<details>
    <summary>GET /users</summary>

```Json
[
    {
        "user_id": user_id,
        "contract_id": contract_id,
        "function_id": function_id,
        "role_id": role_id,
        "username": "username",
        "last_name": "last_name",
        "first_name": "first_name",
        "mail": "mail@*uca.fr",
        "theorical_hours_worked": theorical_hours_worked,
        "vacation_hours": vacation_hours
    },
    {
        "user_id": user_id,
        "contract_id": contract_id,
        "function_id": function_id,
        "role_id": role_id,
        "username": "username",
        "last_name": "last_name",
        "first_name": "first_name",
        "mail": "mail@*uca.fr",
        "theorical_hours_worked": theorical_hours_worked,
        "vacation_hours": vacation_hours
    }
]
```
</details>


<details>
    <summary>GET /users/{user_id}</summary>

```Json
{
    "user_id": user_id,
    "contract_id": contract_id,
    "function_id": function_id,
    "role_id": role_id,
    "username": "username",
    "last_name": "last_name",
    "first_name": "first_name",
    "mail": "mail@*uca.fr",
    "theorical_hours_worked": theorical_hours_worked,
    "vacation_hours": vacation_hours
}
```
</details>

<details>
    <summary>GET /companies/{company_id}/users</summary>

```Json
[
    {
        "user_id": user_id,
        "contract_id": contract_id,
        "function_id": function_id,
        "role_id": role_id,
        "username": "username",
        "last_name": "last_name",
        "first_name": "first_name",
        "mail": "mail@*uca.fr",
        "theorical_hours_worked": theorical_hours_worked,
        "vacation_hours": vacation_hours
    },
    {
        "user_id": user_id,
        "contract_id": contract_id,
        "function_id": function_id,
        "role_id": role_id,
        "username": "username",
        "last_name": "last_name",
        "first_name": "first_name",
        "mail": "mail@*uca.fr",
        "theorical_hours_worked": theorical_hours_worked,
        "vacation_hours": vacation_hours
    }
]
```
</details>

<details>
    <summary>GET /projects/{project_id}/users</summary>

```Json
[
    {
        "user_id": user_id,
        "contract_id": contract_id,
        "function_id": function_id,
        "role_id": role_id,
        "username": "username",
        "last_name": "last_name",
        "first_name": "first_name",
        "mail": "mail@*uca.fr",
        "theorical_hours_worked": theorical_hours_worked,
        "vacation_hours": vacation_hours
    },
    {
        "user_id": user_id,
        "contract_id": contract_id,
        "function_id": function_id,
        "role_id": role_id,
        "username": "username",
        "last_name": "last_name",
        "first_name": "first_name",
        "mail": "mail@*uca.fr",
        "theorical_hours_worked": theorical_hours_worked,
        "vacation_hours": vacation_hours
    }
]
```
</details>

<details>
    <summary>GET /schedules/{schedule_id}/users</summary>

```Json
[
    {
        "user_id": user_id,
        "contract_id": contract_id,
        "function_id": function_id,
        "role_id": role_id,
        "username": "username",
        "last_name": "last_name",
        "first_name": "first_name",
        "mail": "mail@*uca.fr",
        "theorical_hours_worked": theorical_hours_worked,
        "vacation_hours": vacation_hours
    },
    {
        "user_id": user_id,
        "contract_id": contract_id,
        "function_id": function_id,
        "role_id": role_id,
        "username": "username",
        "last_name": "last_name",
        "first_name": "first_name",
        "mail": "mail@*uca.fr",
        "theorical_hours_worked": theorical_hours_worked,
        "vacation_hours": vacation_hours
    }
]
```
</details>

<details>
    <summary>POST /users</summary>

##### Request parameters
```Json
{
    "contract_id": contract_id,
    "function_id": function_id,
    "role_id": role_id,
    "username": "username",
    "last_name": "last_name",
    "first_name": "first_name",
    "mail": "mail@*uca.fr",
    "theorical_hours_worked": theorical_hours_worked,
    "vacation_hours": vacation_hours
}
```

##### Return parameters
```
A 200 Code and the ID of the new User.
```

</details>

<details>
    <summary>DELETE /users/{user_id}</summary>

##### Return parameters
```
Just a 200 code.
```
</details>

<details>
    <summary>PATCH /users/{user_id}</summary>

###### Both request and return parameters
```Json
{
    "user_id": user_id,
    "contract_id": contract_id,
    "function_id": function_id,
    "role_id": role_id,
    "username": "username",
    "last_name": "last_name",
    "first_name": "first_name",
    "mail": "mail@*uca.fr",
    "theorical_hours_worked": theorical_hours_worked,
    "vacation_hours": vacation_hours
}
```
</details>

## Companies

<details>
    <summary>GET /companies</summary>

```Json
[
    {
        "company_id": company_id,
        "company_name": "company_name"
    },
    {
        "company_id": company_id,
        "company_name": "company_name"
    }
]
```
</details>

<details>
    <summary>GET /companies/{company_id}</summary>

```Json
{
    "company_id": company_id,
    "company_name": "company_name"
}
```
</details>

<details>
    <summary>POST /companies</summary>

##### Request parameters
```Json
{
    "company_name": company_name,
}
```

##### Return parameters
```
A 200 Code and the Id of the new Company.
```

</details>

<details>
    <summary>DELETE /companies/{company_id}</summary>

##### Returns
```
Just a 200 code.
```
</details>

<details>
    <summary>PATCH /companies/{company_id}</summary>

##### Both request and return parameters
```Json
{
    "company_id": company_id,
    "company_name": company_name
}
```
</details>


## Projects

<details>
    <summary>GET /projects</summary>

```Json
[
    {
        "project_id": project_id,
        "project_name": "project_name"
    },
    {
        "project_id": project_id,
        "project_name": "project_name"
    }
]
```
</details>

<details>
    <summary>GET /projects/{project_id}</summary>

```Json
{
    "project_id": project_id,
    "project_name": "project_name"
}
```    
</details>

<details>
    <summary>GET /companies/{company_id}/projects</summary>

```Json
[
    {
        "project_id": project_id,
        "project_name": "project_name"
    },
    {
        "project_id": project_id,
        "project_name": "project_name"
    }
]
```
</details>

<details>
    <summary>GET /users/{user_id}/projects</summary>

```Json
[
    {
        "project_id": project_id,
        "project_name": "project_name"
    },
    {
        "project_id": project_id,
        "project_name": "project_name"
    }
]
```
</details>

<details>
    <summary>POST /projects</summary>

##### Request parameters
```Json
{
    "project_name": "project_name"
}
```

##### Return parameters
```
A 200 code and the Id of the new Project.
```
</details>

<details>
    <summary>DELETE /projects/{project_id}</summary>

##### Return parameters
```
Just a 200 code.
```
</details>

<details>
    <summary>PATCH /projects/{project_id}</summary>

##### Both request and return parameters
```Json
{
    "project_id": project_id,
    "project_name": project_name
}
```
</details>

## Comments

<details>
    <summary>GET /comments</summary>

```Json
[
    {
        "comment_id": comment_id,
        "comment_text": "comment_text",
        "is_important": true/false,
        "schedule_id": schedule_id
    }
]
```
</details>

<details>
    <summary>GET /comments/{comment_id}</summary>

```Json
{
    "comment_id": comment_id,
    "comment_text": "comment_text",
    "is_important": true/false,
    "schedule_id": schedule_id
}
```
</details>

<details>
    <summary>GET /users/{user_id}/comments</summary>

```Json
[
    {
        "comment_id": comment_id,
        "comment": "comment_text",
        "is_important": true/false,
        "schedule_id": schedule_id
    }
]
```
</details>

<details>
    <summary>GET /schedule/{schedule_id}/comments</summary>

```Json
[
    {
        "comment_id": comment_id,
        "comment": "comment_text",
        "is_important": true/false,
        "schedule_id": schedule_id
    }
]
```
</details>


<details>
    <summary>GET /projects/{project_id}/comments</summary>

```Json
[
    {
        "comment_id": comment_id,
        "comment": "comment_text",
        "is_important": true/false,
        "schedule_id": schedule_id
    }
]
```
</details>

<details>
    <summary>POST /comments</summary>

##### Request parameters
```Json
{
    "schedule_id": schedule_id,
    "comment": comment,
    "is_important": true/false
}
```

##### Return parameters
```
A 200 code and the Id of the new Comment.
```
</details>

<details>
    <summary>DELETE /comments/{comment_id}</summary>

##### Return parameters
```
Just a 200 code.
```
</details>

<details>
    <summary>PATCH /comments/{comment_id}</summary>

##### Both request and return parameters
```Json
{
    "comment_id": comment_id,
    "schedule_id": schedule_id,
    "schedule_id": schedule_id,
    "is_important": true/false
}
```
</details>

## Vacations

<details>
    <summary>GET /user/{user_id]}/vacations</summary>

```Json
[
    {
        "schedule_id": schedule_id,
        "start_date": start_date,
        "end_date": end_date
    }
]
```
</details>

<details>
    <summary>GET /vacations/{vacation_id}</summary>

```Json
{
    "schedule_id": schedule_id,
    "start_date": start_date,
    "end_date": end_date
}
```
</details>

<details>
    <summary>POST /vacations</summary>

##### Request parameters
```Json
{
    "start_date": start_date,
    "end_date": end_date
}
```

##### Return parameters
```
A 200 code and the Id of the new Vacations.
```
</details>

<details>
    <summary>DELETE /vacations/{vacations_id}</summary>

##### Return parameters
```
Just a 200 code.
```
</details>

<details>
    <summary>PATCH /vacations/{vacations_id}</summary>

##### Both request and return parameters
```Json
{
    "schedule_id": schedule_id,
    "start_date": start_date,
    "end_date": end_date
}
```
</details>

## Schedules

<details>
    <summary>GET /schedules/{schedule_id}</summary>

```Json
{
    "schedule_id": schedule_id,
    "project_id": project_id,
    "start_date": start_date,
    "end_date": end_date
}
```
</details>

<details>
    <summary>GET /user/{user_id}/schedules</summary>

```Json
[
    {
        "schedule_id": schedule_id,
        "project_id": project_id,
        "start_date": start_date,
        "end_date": end_date
    }
]
```
</details>

<details>
    <summary>GET /projects/{project_id}/schedules</summary>

```Json
[
    {
        "schedule_id": schedule_id,
        "project_id": project_id,
        "start_date": start_date,
        "end_date": end_date
    }
]
```
</details>

<details>
    <summary>POST /schedules</summary>

##### Request parameters
```Json
{
    "project_id": project_id,
    "start_date": start_date,
    "end_date": end_date
}
```

##### Return parameters
```
A 200 code and the Id of the new Schedule.
```
</details>

<details>
    <summary>DELETE /schedules/{schedule_id}</summary>

##### Return parameters
```
Just a 200 code.
```
</details>

<details>
    <summary>PATCH /schedules/{schedule_id}</summary>

##### Both request and return parameters
```Json
{
    "schedule_id": schedule_id,
    "project_id": project_id,
    "start_date": start_date,
    "end_date": end_date
}
```
</details>

## Roles

<details>
    <summary>GET /roles</summary>

```Json
[
    {
        "role_id": role_id,
        "role_name": "role_name",
        "can_add_and_modify_users": true/false,
        "can_see_other_schedules": true/false,
        "can_add_projects": true/false,
        "can_see_reports": true/false
    },
    {
        "role_id": role_id,
        "role_name": "role_name",
        "can_add_and_modify_users": true/false,
        "can_see_other_schedules": true/false,
        "can_add_projects": true/false,
        "can_see_reports": true/false
    }
]
```
</details>

<details>
    <summary>GET /roles/{role_id}</summary>

```Json
{
    "role_id": role_id,
    "role_name": "role_name",
    "can_add_and_modify_users": true/false,
    "can_see_other_schedules": true/false,
    "can_add_projects": true/false,
    "can_see_reports": true/false
}
```
</details>

<details>
    <summary>GET /users/{user_id}/roles</summary>

```Json
[
    {
        "role_id": role_id,
        "role_name": "role_name",
        "can_add_and_modify_users": true/false,
        "can_see_other_schedules": true/false,
        "can_add_projects": true/false,
        "can_see_reports": true/false
    },
    {
        "role_id": role_id,
        "role_name": "role_name",
        "can_add_and_modify_users": true/false,
        "can_see_other_schedules": true/false,
        "can_add_projects": true/false,
        "can_see_reports": true/false
    }
]
```
</details>

<details>
    <summary>POST /roles</summary>

##### Request parameters
```Json
{
    "role_name": "role_name",
    "can_add_and_modify_users": true/false,
    "can_see_other_schedules": true/false,
    "can_add_projects": true/false,
    "can_see_reports": true/false
}
```

##### Return parameters
```
A 200 code and the Id of the new Role.
```

</details>

<details>
    <summary>DELETE /roles/{role_id}</summary>

##### Returns
```
Just a 200 code.
```
</details>

<details>
    <summary>PATCH /roles/{role_id}</summary>

##### Both request and return parameters
```Json
{
    "role_id": role_id,
    "role_name": "role_name",
    "can_add_and_modify_users": true/false,
    "can_see_other_schedules": true/false,
    "can_add_projects": true/false,
    "can_see_reports": true/false
}
```
</details>

## Contract

<details>
    <summary>GET /contracts</summary>

```Json
[
    {
        "contract_id": contract_id,
        "contract_name": "contract_name"
    },
    {
        "contract_id": contract_id,
        "contract_name": "contract_name"
    }
]
```
</details>

<details>
    <summary>GET /contracts/{contract_id}</summary>

```Json
{
    "contract_id": contract_id,
    "contract_name": "contract_name"
}
```
</details>

<details>
    <summary>GET /users/{user_id}/contracts</summary>

```Json
{
    "contract_id": contract_id,
    "contract_name": "contract_name"
}
```
</details>

<details>
    <summary>POST /contracts</summary>

##### Request parameters
```Json
{
    "contract_name": "contract_name",
}
```

##### Return parameters
```
A 200 code and the Id of the new Contract.
```

</details>

<details>
    <summary>DELETE /contracts/{contract_id}</summary>

##### Returns
```
Just a 200 code.
```
</details>

<details>
    <summary>PATCH /contracts/{contract_id}</summary>

##### Both request and return parameters
```Json
{
    "contract_id": contract_id,
    "contract_name": "contract_name"
}
```
</details>

## Function

<details>
    <summary>GET /functions</summary>

```Json
[
    {
        "function_id": function_id,
        "function_name": "function_name"
    },
    {
        "function_id": function_id,
        "function_name": "function_name"
    }
]
```
</details>

<details>
    <summary>GET /functions/{function_id}</summary>

```Json
{
    "function_id": function_id,
    "function_name": "function_name"
}
```
</details>

<details>
    <summary>GET /users/{user_id}/functions</summary>

```Json
{
    "function_id": function_id,
    "function_name": "function_name"
}
```
</details>

<details>
    <summary>POST /functions</summary>

##### Request parameters
```Json
{
    "function_name": "function_name",
}
```

##### Return parameters
```
A 200 code and the Id of the new Function.
```

</details>

<details>
    <summary>DELETE /functions/{function_id}</summary>

##### Returns
```
Just a 200 code.
```
</details>

<details>
    <summary>PATCH /functions/{function_id}</summary>

##### Both request and return parameters
```Json
{
    "function_id": function_id,
    "function_name": "function_name"
}
```
</details>