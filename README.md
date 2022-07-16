# Go-Mux-CRUD-API-Clean-Architecture

## **Flow of the project**
Check the diagram folder to understand the flow of the project

### **Step 1:** 
  Started with the **_routes/bookstore-routes.go_** to create the routes to connect to the controllers
  
### **Step 2:** 
  Moved to the **_config/app.go_** to create the varible db and to return the variable db which will help the other files to interact with the db
  
### **Step 3:** 
  Next is the **_utils/utils.go_** to unmarshal the data for the request
  
### **Step 4:** 
  Create the **_models/book.go_** to give gives us a structure to help store something in the database
  
### **Step 5:** 
  Create the **_cmd/main/main.go_** to create the server which will define the localhost and show where the routers reside

* routes/bookstore-routes.go required **controllers package**
* config/app.go required **GORM package**
* models/book.go required **config package**

* The control are first in the **ROUTES**
* ROUTES gives controls to the **CONTROLLERS**
* And CONTROLLERS give control to the **MODELS**

* In this sense the users interact with the routes and the routes send
control to the controllers where we have all the logic
Controllers have to perform some operations with the database
The operations of the database reside inside the models which book.go in this case
Hence the functions for the controllers are defined in the models
