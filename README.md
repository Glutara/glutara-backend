<br>
<div align="center">
    <div >
        <img height="150px" src="https://firebasestorage.googleapis.com/v0/b/upheld-acumen-420202.appspot.com/o/readme-assets%2FGlutara.png?alt=media&token=77d4dd88-6cca-4e4d-94f2-321124c20a61" alt=""/>
    </div>
    <div>
            <h3><b>Glutara</b></h3>
            <p><i>A Key to Your Diabetes Journey</i></p>
    </div>      
</div>
<br>
<h1 align="center">Glutara Backend</h1>
The silent maestro: Glutara's robust backend hums quietly behind the scenes, orchestrating the flow of data, ensuring seamless performance and secure storage. It's the invisible force making everything tick, keeping your health journey smooth and reliable.

## 👨🏻‍💻 &nbsp;Technology Stack

<div align="center">

<a href="https://go.dev/">
<kbd>
<img src="https://firebasestorage.googleapis.com/v0/b/upheld-acumen-420202.appspot.com/o/readme-assets%2Ficons%2FGo.png?alt=media&token=93e65685-5bfa-428d-a9e2-f3cecacbc004" height="60" />
</kbd>
</a>

<a href="https://gin-gonic.com/">
<kbd>
<img src="https://firebasestorage.googleapis.com/v0/b/upheld-acumen-420202.appspot.com/o/readme-assets%2Ficons%2FGin.png?alt=media&token=bdd8b1d9-e6d4-477d-9d4f-0e7dd5a155d0" height="60" />
</kbd>
</a>

<a href="https://mapsplatform.google.com/">
<kbd>
<img src="https://firebasestorage.googleapis.com/v0/b/upheld-acumen-420202.appspot.com/o/readme-assets%2Ficons%2FMaps.png?alt=media&token=5f01c487-3892-4d93-a3e6-dac1909b1e17" height="60" />
</kbd>
</a>

<a href="https://cloud.google.com/">
<kbd>
<img src="https://firebasestorage.googleapis.com/v0/b/upheld-acumen-420202.appspot.com/o/readme-assets%2Ficons%2FGCP.png?alt=media&token=a2af827f-269c-463c-b3d6-567e20822902" height="60" />
</kbd>
</a>

<a href="https://www.docker.com/">
<kbd>
<img src="https://firebasestorage.googleapis.com/v0/b/upheld-acumen-420202.appspot.com/o/readme-assets%2Ficons%2FDocker.png?alt=media&token=3588896c-975f-496f-87d0-e7e1bce0d492" height="60" />
</kbd>
</a>

<a href="https://firebase.google.com/">
<kbd>
<img src="https://firebasestorage.googleapis.com/v0/b/upheld-acumen-420202.appspot.com/o/readme-assets%2Ficons%2FFirebase.png?alt=media&token=da3b3135-dec1-4f6c-b0db-0051541754b6" height="60" />
</kbd>
</a>

<a href="https://developers.google.com/ml-kit">
<kbd>
<img src="https://firebasestorage.googleapis.com/v0/b/upheld-acumen-420202.appspot.com/o/readme-assets%2Ficons%2FMLKit.png?alt=media&token=e1fa8c3a-4cfa-4a5d-ad68-3bc29b7e0d21" height="60" />
</kbd>
</a>

<a href="https://gemini.google.com/">
<kbd>
<img src="https://firebasestorage.googleapis.com/v0/b/upheld-acumen-420202.appspot.com/o/readme-assets%2Ficons%2FGemini%20(1).png?alt=media&token=39c99afa-b82c-45f7-b59c-4df0ccecfe54" height="60" />
</kbd>
</a>

</div>
<div align="center">
<h4>Go | Gin | Google Maps Platform | Google Cloud Platform | Docker | Firebase | MLKit | Gemini</h4>
</div>

## Getting Started
Make sure you already do these things before running the code
1. Install Go language on your computer

## ⚙️ &nbsp;How to Run
1. Clone this repository from terminal using this following command
    ``` bash
    git clone https://github.com/Glutara/glutara-backend.git
    ```
2. Create a .env file inside the repository directory using .env.example file as the template. You can keep the variables blank. The server should automatically use port 8080 as the default port and port 8605 as the model serving port
3. Run the server using this following command
    ``` bash
    go run main.go
    ```
4. Glutara backend server should be running. You can also check the server by opening http://localhost:8080/api
5. You could also check our deployed backend server by opening https://glutara-rest-api-reyoeq7kea-uc.a.run.app/api
    
## 🔑 &nbsp;List of Endpoints

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
| /api/{UserID}/glucoses/info/graphic        | GET      | Users can see graphic of their blood glucose level fluctuation at a certain day
| /api/{UserID}/glucoses/info/average        | GET      | Users can see the average value of their blood glucose level as of today, this week, and this month
| /api/{UserID}/glucoses               | POST     | System can automatically predict and save user's blood glucose level
| /api/{UserID}/logs                   | GET     | Users can view a comprehensive list of all their log activity
| /api/{UserID}/relations              | GET     | Users with 'patient' role can see info and location of another user that is related to them
| /api/{UserID}/relations              | POST    | Users with 'patient' role can add new user as their relative
| /api/{UserID}/relations/related      | GET     | Users with 'relation' role can see info and current blood glucose level of another user with 'patient' role
| /api/{UserID}/scan                   | POST    | Users can post picture of their food to know the name and nutritional value of that food

## 👥 &nbsp;Contributors

| <a href="https://github.com/mikeleo03"><img width="180px" height="180px" src="https://firebasestorage.googleapis.com/v0/b/upheld-acumen-420202.appspot.com/o/readme-assets%2Fpicprof%2FLeon.png?alt=media&token=0ea1884a-32ca-471b-a3af-bf3995bbc605" alt=""/></a> | <a href="https://github.com/GoDillonAudris512"><img width="180px" height="180px" src="https://firebasestorage.googleapis.com/v0/b/upheld-acumen-420202.appspot.com/o/readme-assets%2Fpicprof%2FDillon.png?alt=media&token=bc76cc6b-5606-4351-8472-9c243c8b9da3" alt=""/></a> | <a href="https://github.com/margarethaolivia"><img width="180px" height="180px" src="https://firebasestorage.googleapis.com/v0/b/upheld-acumen-420202.appspot.com/o/readme-assets%2Fpicprof%2FOlivia.png?alt=media&token=d53f9cfd-e1e1-41b6-a28c-440904df29b8" alt=""/></a> | <a href="https://github.com/AustinPardosi"><img width="180px" height="180px" src="https://firebasestorage.googleapis.com/v0/b/upheld-acumen-420202.appspot.com/o/readme-assets%2Fpicprof%2FAustin.png?alt=media&token=f520a334-4aeb-4efe-9437-669451b6dca6" alt=""/></a> |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| <div align="center"><h3><b><a href="https://github.com/mikeleo03">Michael Leon Putra Widhi</a></b></h3><i><p>Bandung Institute of Technology</i></p></div>                                                                               | <div align="center"><h3><b><a href="https://github.com/GoDillonAudris512">Go Dillon Audris</a></b></h3></a><p><i>Bandung Institute of Technology</i></p></div>                                                                          | <div align="center"><h3><b><a href="https://github.com/margarethaolivia">Margaretha Olivia Haryono</a></b></h3></a><p><i>Bandung Institute of Technology</i></p></div>                                                               | <div align="center"><h3><b><a href="https://github.com/AustinPardosi">Austin Gabriel Pardosi</a></b></h3></a><p><i>Bandung Institute of Technology</i></p></div>                                                                            |
