package main

//========================================================================
//	 Employee Directory Structure Definitions
//========================================================================

// Reference Pointer
type ChainCode struct {

}

// Employee Directory
type EmpDir struct{
	UID 				string 							`json:"uid"`
    PersonalDetails     Personal                        `json:personalDetails`
	Skills              []Skill        					`json:"skills"`
	EducationDetails    []Education                     `json:"educationDetails"`
	EmploymentDetails   []Employment                    `json:"employmentDetails"`
	Submitted           bool                            `json:"submitted"`
}

// Personal Details
type Personal struct{
	FirstName           string                          `json:"firstName"`
	LastName            string                          `json:"lastName"`
	DOB                 string                          `json:"dob"`
	Father              string                          `json:"father"`
	Mother              string                          `json:"mother"`
	Address             string                          `json:"address"`
	PAN                 string                          `json:"pan"`
	Passport            string                          `json:"passport"`
}

// Skill Details
type Skill struct{
	SkillName           string                          `json:"skillName`
	Experience          int                             `json:"experience"`
	Level               int                             `json:"level"`	    
}

// Education Details
type Education struct{
	Course              string                          `json:"course`
	University          string                          `json:"university"`
	UnivCode            string                          `json:"univcode"`
	College             string                          `json:"college"`
	Enrollment          string                          `json:"enrollment"`
	Year                int                             `json:"level"`	
	Grades              string                         	`json:"grades"`
	Verified            bool                            `json:"verified"`
}

// Employment Details
type Employment struct{
	Employer            string                          `json:"employer"`
	EmployerCode        string                          `json:"employercode"`
	EmployeeCode        string                          `json:"employeecode"`
	FromDate            string                          `json:"fromDate"`
	ToDate              string                          `json:"toDate"`
	Location            string                          `json:"location"`
	Verified            bool                             `json:"verified"`	
}

// Universities
type Univs struct{
	UnivID 				string 							`json:"univid"`
    PersonalDetails     Personal                        `json:"personal"`
	EducationDetails    Education                       `json:"educationDetails"`
	Visible             bool                            `json:"visible"`
}

// Employers
type Employers struct{
	EmploymentID     	string 							`json:"employmentid"`
    PersonalDetails     Personal                        `json:"personal"`
	EmploymentDetails   Employment                      `json:"employmentDetails"`
    Visible             bool                            `json:"visible"`
}