<div align="center"><img src = "assets/glutara.png" width = 20% height= 20%></div>

<div align="center">
An IOT-based system with non-invasive wearable Continuous Glucose Monitor (CGM) for diabetic people
</div>

## Table of Contents

- [App Overview](#app-overview)
- [Prerequisite](#prerequisite)
- [How to Run](#how-to-run)
- [List of Endpoints](#list-of-endpoints)
- [Team](#team)

## App Overview
Millions of individuals worldwide grapple with the relentless challenges of managing diabetes, a chronic condition that demands consistent monitoring and care. Despite the advancements in technology, the process remains burdensome, with traditional finger-pricking glucose monitoring causing discomfort and hindering regular monitoring. This issue is exacerbated for those leading busy lives, leaving little time for necessary health measures. The fear of potential health emergencies, particularly for individuals living alone, further
compounds the need for a more accessible and painless solution. Glutara aims to revolutionize
blood glucose monitoring by offering a seamless and affordable solution, addressing the
fundamental challenges faced by those managing diabetes on a daily basis.

This repository is dedicated to manage the code needed to run Glutara backend server.

## Prerequisite
Make sure you already do these things before running the code
1. Install Go languange on your computer
2. Download the Firebase admin-sdk json file located [here](https://drive.google.com/file/d/18jmUb9Jbsv71MlGfDy0UBbCkTBL0Khf_/view?usp=sharing)

## How to Run
1. Clone this repository from terminal using this following command
    ``` bash
    $ git clone https://github.com/Glutara/glutara-backend.git
    ```
2. Create a .env file inside the repository directory using .env.example file as the template. You can keep the PORT variable blank. The server should automatically use port 8080 as the default port
3. Using Windows PowerShell, navigate to this repository directory
4. Set the GOOGLE_APPLICATION_CREDENTIALS environment variable using this following command
    ``` bash
    $ $env:GOOGLE_APPLICATION_CREDENTIALS="path/to/the/admin-sdk/json/file"
    ```
5. Run the server using this following command
    ``` bash
    $ go run main.go
    ```
6. Glutara backend server should be running. You can also check the server by opening http://localhost:8080/api
    
## List of Endpoints

| Endpoint                             |  Method  |   Usage  |
| ------------------------------------ | :------: | -------- |
| /api/auth/register                   | POST     | Users can register and create account on Glutara App
| /api/auth/login                      | POST     | Users can log in to their previously created account
| /api/{UserID}/reminders              | GET      | Users can see their reminders
| /api/{UserID}/reminders              | POST     | Users can add a new reminder
| /api/{UserID}/reminders/{ReminderID} | DELETE   | Users can delete existing reminder
| /api/{UserID}/reminders/{ReminderID} | PUT      | Users can update existing reminder
| /api/{UserID}/sleeps                 | GET      | Users can see their sleep logs
| /api/{UserID}/sleeps                 | POST     | Users can add a new sleep log
| /api/{UserID}/sleeps/{SleepID}       | DELETE   | Users can delete existing sleep log
| /api/{UserID}/sleeps/{SleepID}       | PUT      | Users can update existing sleep log
| /api/{UserID}/exercises              | GET      | Users can see their exercise logs
| /api/{UserID}/exercises              | POST     | Users can add a new exercise log
| /api/{UserID}/exercises/{ExerciseID} | DELETE   | Users can delete existing exercise log
| /api/{UserID}/exercises/{ExerciseID} | PUT      | Users can update existing exercise log
| /api/{UserID}/meals                  | GET      | Users can see their meal logs
| /api/{UserID}/meals                  | POST     | Users can add a new meal log
| /api/{UserID}/meals/{MealID}         | DELETE   | Users can delete existing meal log
| /api/{UserID}/meals/{MealID}         | PUT      | Users can update existing meal log
| /api/{UserID}/medications                  | GET      | Users can see their medication logs
| /api/{UserID}/medications                  | POST     | Users can add a new medication log
| /api/{UserID}/medications/{MedicationID}   | DELETE   | Users can delete existing medication log
| /api/{UserID}/medications/{MedicationID}   | PUT      | Users can update existing medication log

## Team

Created and developed by AMN:
| Name                           |   Role   |
| ------------------------------ | :------: |
| Michael Leon Putra Widhi       | Hustler  |
| Margaretha Olivia Haryono      | Hipster  |
| Go Dillon Audris               | Hacker   |
| Austin Gabriel Pardosi         | Hacker   |