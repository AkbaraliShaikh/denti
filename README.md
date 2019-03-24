
# denti-go-clean-arch

[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://travis-ci.org/joemccann/dillinger)

This project is created to demonstrate a CleanArchitecture/DDD/HexaArchitecture for dental clinic application built with go gin api including gorm CRUD operations and more.

#Go #Go-Lang #CleanArchitecture #DDD #HexaArchitecture #GoProjectStructure #RealWorldExample #DentalClinic #denti 

# Benefits :
  - flexibility
  - testability

### Dependencies:

- [gin](https://github.com/gin-gonic/gin)                 		 - Http Api Server
- [gorm](https://github.com/jinzhu/gorm)			 	 - Object relational mapping (database)
- [zap](https://github.com/uber-go/zap) 				 - Logger
- [dig](https://github.com/uber-go/dig)					 - Dependency Injection
- [configor](https://github.com/jinzhu/configor) 	 		 - Configuration Helper
- [go.uuid](https://github.com/satori/go.uuid) 				 - UUID v4

###  Project Structure:

  <a target="_blank" href="https://github.com/AkbaraliShaikh/AspNetCore2Docker/blob/master/img/Go_Project_Structure.PNG" class="rich-diff-level-one"><img src="https://github.com/AkbaraliShaikh/AspNetCore2Docker/raw/master/img/Go_Project_Structure.PNG" alt="text" width=75%  height=400px></a>
  
###  Run:
  - $ sudo docker-compose up -d
  - $ curl http://localhost:8282/api/v1/health/
  
###  Test:
```curl --header "Content-Type: application/json" --request POST --data '{"email":"a.a.shaikh55@gmail.com","first_name":"Akbar","last_name":"Shaikh","password":"password","gender":1,"picture":"","phone_number":"9123456789"}'  http://localhost:8283/api/v1/users/```

`curl http://localhost:8283/api/v1/users/`

#####  => Repo work is in-progress, will be keep updating the code, but the basic project structure is ready with the required code to understand, have a look!

## Happy Coding!
