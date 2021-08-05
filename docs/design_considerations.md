# Design Considerations
At the time of writing, this project is in it's early stages.

## How big should functions be?
Functions should be small.

Smaller than that.

## Write tests
Tests should not have external dependencies.
We want the test to run fast in order to give quick feedback to developers.

Try to take advantage of TDD and "red, green, refactor" principals.
That is, 
* write a test
* make it pass
* write some code until the test fails
* write more test until it passes
* write more code until the test fails

## Program to an interface
As we build things out, we might not know all the parts that are needed. 

Take advantage of [golang's interfaces](https://gobyexample.com/interfaces) to connect things at different levels of abstractions.  This will allow us to build in a very modular way.

With interfaces, we can build systems without knowing the complete details of their implementations.

i.e. define the signature of an interface and we can implement it later however we like.

At the time of writing this, I've built a few ["in_memory" datastores](https://github.com/HackRVA/hackpoints/tree/main/datastore/in_memory) that will probably be replaced by database methods later.

## No Versions
If it's in the main branch, it should be able to run.
You can make branches, but ideally, they shouldn't live long.

Eventually we will hook this up to automatically deploy (after running tests) what's in the main branch.
