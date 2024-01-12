<div align="center"><img src = "assets/glutara.png" width = 20% height= 20%></div>

<div align="center">
An IOT-based system with non-invasive wearable Continuous Glucose Monitor (CGM) for diabetic people
</div>

## Table of Contents

- [App Overview](#app-overview)
- [List of Endpoints](#list-of-endpoints)
- [Team](#team)

## App Overview
Millions of individuals worldwide grapple with the relentless challenges of managing diabetes, a chronic condition that demands consistent monitoring and care. Despite the advancements in technology, the process remains burdensome, with traditional finger-pricking glucose monitoring causing discomfort and hindering regular monitoring. This issue is exacerbated for those leading busy lives, leaving little time for necessary health measures. The fear of potential health emergencies, particularly for individuals living alone, further
compounds the need for a more accessible and painless solution. Glutara aims to revolutionize
blood glucose monitoring by offering a seamless and affordable solution, addressing the
fundamental challenges faced by those managing diabetes on a daily basis.

## List of Endpoints

| Endpoint                             |  Method  |   Usage  |
| ------------------------------------ | :------: | -------- |
| /api/auth/register                   | POST     | Users can register and create account on Glutara App
| /api/auth/login                      | POST     | Users can log in to their previously created account
| /api/{UserID}/reminders              | GET      | Users can see their reminders
| /api/{UserID}/reminders              | POST     | Users can add a new reminder
| /api/{UserID}/reminders/{ReminderID} | DELETE   | Users can delete existing reminder

## Team

Created and developed by AMN:
| Name                           |   Role   |
| ------------------------------ | :------: |
| Michael Leon Putra Widhi       | Hustler  |
| Margaretha Olivia Haryono      | Hipster  |
| Go Dillon Audris               | Hacker   |
| Austin Gabriel Pardosi         | Hacker   |