# Requirement Analysis

## Assumptions

- Used internally
- Light load, <2,000 users, DAU < 100
- average 5 appointments ops/sec
- Not using SSOs, therefore need a built in authorization?

## Functional Requirements

Only listing the bare minimum required by the wire frame assuming all extra functionalities are provided. i.e. logins, authorization, ... I may complete the extra functionalities if i'm free enough to.

### Appointment actions

| Action | Object                      | Return          |
| ------ | --------------------------- | --------------- |
| GET    | all appointments, paginated | Appointments[]  |
| GET    | specific appointment        | Appointment     |
| PATCH  | appointment status          | update feedback |
| PATCH  | archive appointment         | update feedback |

### Comments actions

| Action | Object                                       | Return      |
| ------ | -------------------------------------------- | ----------- |
| GET    | all comments for an Appointment, (paginated) | Comments[]  |
| POST   | new comment                                  | Appointment |

## Data Interface

### User Entity

| Name  | Type   | Description                         |
| ----- | ------ | ----------------------------------- |
| uuid  | string | user's identifier                   |
| name  | string | name shown on all pages             |
| image | string | profile image url                   |
| email | string | user's email shown in detailed card |

### Appointment Entity

| Name    | Type   | Description                  |
| ------- | ------ | ---------------------------- |
| uuid    | string | appointment's identifier     |
| details | string | text shown in the boxes      |
| status  | enum   |
| email   | string | email shown in detailed view |

### Status Enum

| Name        | Description                   |
| ----------- | ----------------------------- |
| TODO        |
| IN_PROGRESS |
| DONE        |
| ARCHIVED    | Not shown in all appointments |

transition from todo -> in_progress -> done

_todo: maybe enforce the state transition in BE?_
