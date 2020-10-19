package main

import
(
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
type Meeting struct
{
	Id primitive.ObjectID = `json:"Id,omitempty"`
	Title string = `json:"Title,omitempty"`
	Participants object = `json:"Participants,omitempty"`
	Start_time time.Time = `json:"Start_time,omitempty"`
	End_time time.Time = `json:"End_time,omitempty"`   
}
type Participants struct
{
	Name string = `json:"Name,omitempty"` 
	Email string = `json:"Email,omitempty"` 
	RSVP string = `json:"RSVP,omitempty"`
}
//create json structure for meeting and object for participants in meeting
var meet1 Meeting
var meet2 Meeting
var name1 Participants
var name2 Participants
name1 := `{"Name":"Charles","Email":"Charles1968@gmail.com","RSVP":"Yes"}`
name2 := `{"Name":"John","Email":"Johnhappy@yahoo.co.in","RSVP":"No"}`
meet1 := `{"Id":"1001","Title":"GoLang Tutorial","Participants":"name1","Start_time":"09:00:00","End_time":"10:00:00"}`
meet2 := `{"Id":"1002","Title":"Mongodb with GoLang","Participants":"name2","Start_time":"10:00:00","End_time":"11:00:00"}`
//add values to the json structure
func schedule_meeting(w http.ResponseWriter, r *http.Request)
{
	var M Meeting 
	json_Data,_ := json.Marshal(M)
	response,err := http.Post("/meeting","application.json",bytes.NewBuffer(json_Data))
	//request to post from main
	if err!= nil
	{
		http.Error(w, "404 not found.", http.StatusNotFound)
        return
	}
	else
	{
		decoder := json.NewDecoder(response.Body)
		er := decoder.Decode(&M)
	}	
	return decoder
	//returns the details of the schedules meeting
}
func get_meeting(W http.ResponseWriter, r *http.Request)
{
	var m[] Meeting
	jsonData,_ := json.Marshal(m)
	response,err :=http.Get("/meeting/{Id}","application.json",bytes.NewBuffer(jsonData))
	//implement get method
	response Header().Add("content-type","application.json")
	if err!=nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	for response.Next(len(m))
	{
		var meet Meeting
		response.Decode(&meet)
		members := append(meet,m)
	}
	//add the meeting values into the meet array to return it according to id
	if err := response.Err(); err != nil {
		response.WriteHeader(http.StatusInternalSErverError)
		response.Write([]byte(`"message":"` + err.Error() + `"}`))
		return
	}
	return json.NewEncoder(response).Encode(meet) 
}
func meeting_time(W http.ResponseWriter,r *http.Request)
{
	response.Header().Add("content-type", "application/json")
	var meeting []Meeting
	jsonData,_ := json.Marshal(meeting)
	emailid, _ := request.Get(parameter["Email"],bytes.NewBuffer(jsonData))
	//iplements get method
	cursor, err := Mydatabase.Find(response, Meeting{Email: emailid}).Decode(&meeting)
	//gets the vcalue of email id
	if err!=nil {
		log.Fatal(err)
	}
	var filtered []bson.meeting
	// converts it to binary js objects after filtering
	if err = filterCursor.All(cursor, &filtered); err != nil {
		log.Fatal(err)
	}
	fmt.Println(filtered)
}
func all_meeting(W http.ResponseWriter,r *http.Request)
{
	response.Header().Add("content-type", "application/json")
	var meets []Meeting
	jsonData,_ := json.Marshal(meets)
	stiming := request.Get(parameters[Start_time],bytes.NewBuffer(jsonData))
	etiming := request.Get(parameters[End_time],bytes.NewBuffer(jsonData))
	//implements get method
	cursor, err := Mydatabase.Find(response, Meeting{Start_time:stiming,End_time:etiming}).Decode(&meeting)
	if err!=nil {
		log.Fatal(err)
	}
	var filtered []bson.meeting
	//gets stiming and etiming as bson array for the process of filtering 
	if err = filterCursor.All(cursor, &filtered); err != nil {
		log.Fatal(err)
	}
	fmt.Println(filtered)
}
func main()
{
	var client *mongo.Client
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:8080"))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	//create connection with nosql
	Mydatabase = client.Database("MeetingSchedule")
	meetingcollection = Mydatabase.Collection("meetings")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := meetingcollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close()
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
    collection := client.Database("MeetingSchedule").Collection("meetings")
	insertResult, err := collection.InsertOne(context.TODO(), ash)
	//create database in the connection
	if err != nil {
		log.Fatal(err)
	}
	trainers := []interface{}{meet1,meet2}
	insertManyResult, err := collection.InsertMany(context.TODO(), trainers)
	if err != nil {
		log.Fatal(err)
	}
	//conditions to check the errors due to connection and due to creating database in connection
	var result Trainer

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	var results []*Trainer
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()) {
		var elem Trainer
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}
    // add all the json values into mongo connection
    filterCursor, err := meetingcollection.Find(ctx, bson.M{"Start_time": "09:00:00","End_time":"10:00:00","Participants"."RSVP":"Yes"})
	if err != nil {
		log.Fatal(err)
	}
	var meetingsfiltered []bson.M
	if err = filterCursor.All(ctx, &meetingsfiltered); err != nil {
		log.Fatal(err)
	}
	fmt.Println(meetingsfiltered)
    // checks whether ther is overlapping of Start_time and End_time
	filterCursor, err := meetingcollection.Find(ctx, bson.M{"Participants"."Name": "charles", "Id": "1001"})
	if err != nil {
		log.Fatal(err)
	}
	var meetingsfiltered []bson.M
	if err = filterCursor.All(ctx, &meetingsfiltered); err != nil {
		log.Fatal(err)
	}
	fmt.Println(meetingsfiltered)
	//Checks the id for the particular person for the name given
	http.Handlefunc("/meeting",schedule_meeting).Methods("POST")
	http.Handlefunc("/meeting/{Id}",get_meeting).Methods("GET")
	http.HandleFunc("/meeting?Participant=<{email}>",meeting_time).Methods("GET")
	http.HandleFunc("/meeting?start=<{stiming}>&end=<{etiming}>",all_meeting).Methods("GET")
	http.ListenAndServe("mongodb://localhost:8080",router)
	//function call
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.TODO())
	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connection to MongoDB closed.")
	}
    //connection close	
}