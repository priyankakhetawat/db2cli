var ibmdb = require("ibm_db")
  , connStr = "DATABASE=BLUDB;HOSTNAME=3883e7e4-18f5-4afe-be8c-fa31c41761d2.bs2io90l08kqb1od8lcg.databases.appdomain.cloud;Security=SSL;PORT=31498;PROTOCOL=TCPIP;UID=nzv68443;PWD=Welcome2ibm@123";

ibmdb.open(connStr, function (err, connection) {
    if (err)
    {     
      console.log("\n Nodejs connection in docker container : FAILURE\n\r" , err);
      return;
    }
    else
    {
      console.log("\n Nodejs connection in docker container : SUCCESS\n\r");
    }

});

