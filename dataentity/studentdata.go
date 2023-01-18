package dataentity

type Studentdata struct {
	ID          int    `json:"id"`
	StudentName string `json:"studentname"`
	StudentID   string `json:"studentid"`
	EmailID     string `json:"emailid"`
	BranchName  string `json:"branchname"`
	PassingYear int    `json:"passingyear"`
}
