import ibm_db

conn_str = "HOSTNAME=3883e7e4-18f5-4afe-be8c-fa31c41761d2.bs2io90l08kqb1od8lcg.databases.appdomain.cloud;DATABASE=BLUDB;Security=SSL;PORT=31498;UID=nzv68443;PWD=Welcome2ibm@123"

connection = ibm_db.connect(conn_str,"","")

print("\n Connection string used: " + conn_str) 

if (connection):
        print ("\n python ibm_db connection in docker container : SUCCESS\n\r")
else:
        print ("\n python ibm_db connection in docker container : FAILURE\n\r")


