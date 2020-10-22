package main

import (
	"context"
	"log"

	"github.com/B5916177/app/controllers"
	_ "github.com/B5916177/app/docs"
	"github.com/B5916177/app/ent"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Patients struct {
	Patient []Patient
}

type Patient struct {
	PatientName    string
	PatientGender  string
	PatientAge     int
	PatientPhone   int
}

type Dentists struct {
	Dentist []Dentist
}

type Dentist struct {
	DentistName   string
	DentistEmail  string
	DentistPhone  int
}

type Employees struct {
	Employee []Employee
}

type Employee struct {
	EmployeeName     string
	EmployeeEmail    string
	EmployeePassword string
	EmployeePhone    int
}

// @title SUT SA Example API
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewQueueController(v1, client)
	controllers.NewPatientController(v1, client)
	controllers.NewDentistController(v1, client)
	controllers.NewEmployeeController(v1, client)

	// Set Patients Data
	patients := Patients{
		Patient: []Patient{
			Patient{"jimjum oroy","female",22,687123488},
			Patient{"padpaw kidow","male",31,656567828},
			Patient{"tomyum kung","male",28,689101058},
			Patient{"pizza havaoian","female",19,612135018},
			
		},
	}

	for _, p := range patients.Patient {
		client.Patient.
			Create().
			SetPatientName(p.PatientName).
		    SetPatientGender(p.PatientGender).
	    	SetPatientAge(p.PatientAge).
		    SetPatientPhone(p.PatientPhone).
			Save(context.Background())
	}

	// Set Dentists Data
	dentists := Dentists{
		Dentist: []Dentist{
			Dentist{"somchai kongman","somchai@gmail.com",887214088},
			Dentist{"anan mesug","anan@gmail.com",856784668},
			Dentist{"panawun jidee","panawun@gmail.com",833314058},
			Dentist{"yindee tonrub","yindee@gmail.com",897299048},
			
		},
	}
	for _, d := range dentists.Dentist {
		client.Dentist.
			Create().
			SetDentistName(d.DentistName).
		    SetDentistEmail(d.DentistEmail).
		    SetDentistPhone(d.DentistPhone).
			Save(context.Background())
	}

	// Set Employees Data
	employees := Employees{
		Employee: []Employee{
			Employee{"tomas andisun","tomas@outlook.com","1234",987214088},
			Employee{"gandarff warlock","gandarff@outlook.com","1234",956783338},
			Employee{"peter pakker","peter@outlook.com","1234",968714088},
			Employee{"tony stark","tony@outlook.com","1234",987244488},
			
		},
	}
	for _, e := range employees.Employee {
		client.Employee.
			Create().
			SetEmployeeName(e.EmployeeName).
		    SetEmployeeEmail(e.EmployeeEmail).
		    SetEmployeePassword(e.EmployeePassword).
		    SetEmployeePhone(e.EmployeePhone).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
