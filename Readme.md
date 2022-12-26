go service to generate bash files from json 

## Asumptions: 
* we need a service to sort json tasks and convert them to bash commands
* we may need to reuse it from different applications because it is a http service not a script it self
* we may need to extend to use grpc also so a place for grpc should be considered 
* at the moment there are no usage of a frontend so static files and html generation is not needed 
* as a small service it should have a Dockerfile for easy integration
* Basic approach to sorting is the topological sort problem, no need to implement topo sort, so using 3-td party lib is ok
* No need for an db integration at the moment but maybe there would be usecase for it so it should be considared
    * Db usecases - common jobs saved 




## Actors:
* Go serving instance
* Handler to convert json to go structs
* Struct to bash converter 


## Tasks:
- [x] Setup server with endpoint `/tasks/sort` 
   - [x] Post request json body 
   - [x] validate no circluar dependancies 
   - [x] return json body with sorted commands
- [x] Tests with json file as input and json file as output
  - [x] Good input
  - [x] Invalid input 
  - [x] simple Circular dependancies `task1 requires task2, and task2 requires task1 `
- [x] create package for converting tasks to bash 
- [ ] create test with json input and a .txt file as output
- [x] create endpoint for bash commands generating `/tasks/gen`
- [ ] create test for gen endpoint with .txt file as output
 
   
## Steps 
- [x] Generate init project structure 
- [x] Provide echo service for json 
- [x] Create test for json validation
- [x] Create unit tests for service logic 
- [x] Create bash endpoint 
- [x] Add validation 
- [x] Create unit test for bash endpoint 
- [x] Create integration test for json 
- [x] Create integration test for bash 
- [x] Setup workflow
- [x] Setup dockeer
- [ ] Submit code

Time: 11 hrs
