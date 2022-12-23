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
   - [ ] Post request json body 
   - [ ] validate no circluar dependancies 
   - [ ] return json body with sorted commands
- [ ] Tests with json file as input and json file as output
  - [ ] Good input
  - [ ] Invalid input 
  - [ ] simple Circular dependancies `task1 requires task2, and task2 requires task1 `
  - [ ] self dependencies 
  - [ ] Complex circular dep `t1->t2->t3->t4->t1`
- [ ] create package for converting tasks to bash 
- [ ] create test with json input and a .txt file as output
- [x] create endpoint for bash commands generating `/tasks/gen`
- [ ] create test for gen endpoint with .txt file as output
 
   
## Steps 
- [x] Generate init project structure 
- [x] Provide echo service for json 
- [ ] Create test for json validation
- [ ] Create unit tests for service logic 
- [ ] Setup dockeer
- [ ] Create bash endpoint 
- [ ] Create unit test for bash endpoint 
- [ ] Create integration test for json 
- [ ] Create integration test for bash 
- [ ] Submit code

Time: 2hrs
