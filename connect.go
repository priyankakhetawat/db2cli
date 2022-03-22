package main

import(
    "fmt"
    _ "github.com/ibmdb/go_ibm_db"
    "database/sql"
)

func main(){
    db, err := sql.Open("go_ibm_db", "HOSTNAME=3883e7e4-18f5-4afe-be8c-fa31c41761d2.bs2io90l08kqb1od8lcg.databases.appdomain.cloud;DATABASE=BLUDB;Security=SSL;PORT=31498;UID=nzv68443;PWD=Welcome2ibm@123")
    if err != nil {
        fmt.Println("go ibm driver connection: FAILURE\n" ,err)
    }
    if err = db.Ping(); err != nil {
        fmt.Println("go ibm driver connection: FAILURE\n" ,err)
    } else {
    fmt.Println("go ibm driver connection: SUCCESS\n")
    }
}

