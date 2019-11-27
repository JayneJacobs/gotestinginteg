# 1. Testing and Integration with go

<!-- TOC -->autoauto- [1. Testing and Integration with go](#1-testing-and-integration-with-go)auto        - [1.1. Course Repo](#11-course-repo)auto        - [1.2. Types of testing](#12-types-of-testing)auto        - [1.3. Artifacts of each Test Type](#13-artifacts-of-each-test-type)auto        - [1.4. Pyramid of Tests](#14-pyramid-of-tests)auto        - [1.5. Steps to follow when writing a test cast](#15-steps-to-follow-when-writing-a-test-cast)auto        - [1.6. Test Coverage](#16-test-coverage)auto        - [1.7. [Benchmark Testing](mygolang-testing/BenchmarkTest.md)](#17-benchmark-testingmygolang-testingbenchmarktestmd)auto- [2. [Assertions](https://drive.google.com/file/d/1r5q5i1sATsP510TCInQ6FhPjWbXVMUyt/view?usp=sharing)](#2-assertionshttpsdrivegooglecomfiled1r5q5i1satsp510tcinq6fhpjwbxvmuytviewuspsharing)auto        - [2.1. [API Library](mygolangTesting/api/domain/locations/providerlocations/scenarios.md)](#21-api-librarymygolangtestingapidomainlocationsproviderlocationsscenariosmd)auto        - [3.1 http using Gin](#31-http-using-gin)auto        - [3.2 Stretchr/testify](#32-stretchrtestify)auto    - [4. Functional Test](#4-functional-test)auto    - [E 4.1 Black Box Test](#e-41-black-box-test)auto        - [4.2 Tests in In test folder](#42-tests-in-in-test-folder)autoauto<!-- /TOC -->

### 1.1. Course Repo

```sh
git clone https://github.com/federicoleon/golang-testing.git
```

### 1.2. Types of testing

![TestType](Resources/TestTypes.png)

**Unit:** Test code to make sure functions and Modules work

- White box testing (best to encapsulate all logic in testing)
- Black Box testing: without regard to code

**Integration:** Integrate the modules of teh application not periferal systems.

More than one module: within the program.

Be sure to validate end to end integrated components.

**Functional:**

User behavior
Rest APIs.
Not as rich or complete

**Systems:**

### 1.3. Artifacts of each Test Type

![TestType](Resources/ArtifactsofTypes.png)

### 1.4. Pyramid of Tests

![Pyramid](Resources/Pyramid.png)

Tests must be delivered with the application

test must be in the package of the function they are testing
File name file_test.go
Test Funtion Name starts with Test. Ex.
func TestFunctionName(t *testing.T) {
    //Init
    //Execute
    //Validate
}

### 1.5. Steps to follow when writing a test cast

1. Initialization
2. Execution
3. Validation

**Run Test**
go test (everything in package)
go test -v .

### 1.6. Test Coverage

go test -cover

```PASML-335382:sort jjacob151$ go test -cover
[9 7 5 3 1 2 4 8 0]
[9 8 7 5 4 3 2 1 0]
[9 7 5 3 1 2 4 6 8 0]
[0 1 2 3 4 5 6 7 8 9]
PASS
coverage: 100.0% of statements
ok      gotrain/GoTestingInteg/mygolang-testing/utils/sort      0.005s
PASML-335382:sort jjacob151$
```

The cover metric does not account for the type of validation used.
It verifies that each line of code is executed.
If there is no validation in the code the test may pass even if it is not correct.

Don't rely blindly on the coverage metric.

You must put good validation in the test case.

t.Error Statement will continue to execute after a fail condition.

[Project Integration](mygolang-testing/IntegrationTest.md)

### 1.7. [Benchmark Testing](mygolang-testing/BenchmarkTest.md)

Bubble Sort is the worst possible type of Sort Algorithm. Compare to Native Sort using Benchmarks.
Test different approaches to solving a problem.

# 2. [Assertions](https://drive.google.com/file/d/1r5q5i1sATsP510TCInQ6FhPjWbXVMUyt/view?usp=sharing)

Go Does not have assertions by default; you can create them with error statements.

github.com/stretcher/testify/assert

### 2.1. [API Library](mygolangTesting/api/domain/locations/providerlocations/scenarios.md)

merdadolibre/golang-restclient

```go get github.com/mercadolibre/golang-restclient/rest```

[Countries API](https://api.mercadolibre.com/countries)

[Testing the API](mygolang-testing/api/domain/locations/providerlocations/provider_locations_test.go)

## 3. [Mockups](https://github.com/mercadolibre/golang-restclient/#mockups)

```go test -mock```

[rest.StartMockupServer()](mygolangTesting/api/domain/locations/providerlocations/provider_locations_test.go)

### 3.1 http using Gin

"github.com/gin-gonic/gin"


[CreateContext](mygolangTesting/api/controllers/controller_locations_test.go)

Controller call --> Service --> Provider --> API
![Controller call --> Service --> Provider --> API](Resources/ArtifactsofTypes.png)

### 3.2 Stretchr/testify

has a mock package

```go
type myMockObject struct {
    mock.Mock
}
```

test.Obj.On("Do Something", 123).Return(true, nil)

github.com/stretchr/testify/mock

## 4. Functional Test

Using MVC Model View Controller Architecture

Commandline testing

```sh
go test -cover
?       gotrain/GoTestingInteg/mygolangTesting/api      [no test files]
PASML-335382:api jjacob151$ cd controllers/
PASML-335382:controllers jjacob151$ go test -cover
Init Service
Init sort Service
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

&{{0xc000032440 -1 200} 0xc0002c6200 0xc00029a270 [{country_id AR}] [] -1  0xc0002c4120 map[]  [] map[] map[]}
Inside Controller
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

&{{0xc000032480 -1 200} 0xc000130d00 0xc000278410 [{country_id BR}] [] -1  0xc0000ecc60 map[]  [] map[] map[]}
Inside Controller
this is the result {BR Brasil 3.00+GMT {{0 0}} []}PASS
coverage: 100.0% of statements
ok      gotrain/GoTestingInteg/mygolangTesting/api/controllers  0.020s
```

##E 4.1 Black Box Test
![BlackBox](Resources/BlackBoxTest.png)

Brief test run 

```sh
 go test .
ok      gotrain/GoTestingInteg/mygolangTesting/api/controllers  (cached)
```
Verbose Test Run

```sh
go test . -v
Init Service
Init sort Service
=== RUN   TestGetCountryNotFound
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

&{{0xc000266480 -1 200} 0xc0002b4200 0xc000278270 [{country_id AR}] [] -1  0xc0002ae120 map[]  [] map[] map[]}
Inside Controller
--- PASS: TestGetCountryNotFound (0.00s)
=== RUN   TestGetCountryNoError
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

&{{0xc0002664c0 -1 200} 0xc0002d2000 0xc0002b00d0 [{country_id BR}] [] -1  0xc0002d0000 map[]  [] map[] map[]}
Inside Controller
this is the result {BR Brasil 3.00+GMT {{0 0}} []}--- PASS: TestGetCountryNoError (0.00s)
PASS
ok      gotrain/GoTestingInteg/mygolangTesting/api/controllers  (cached)
```

### 4.2 [Tests in In test folder]

1. Put all tests in a single folder
2. Use one test main function
3. Run the application and use tests against the running application. 
4. 
