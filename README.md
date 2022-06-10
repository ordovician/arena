# Arena Binary Tree

This is a project to demonstrate a bunch of different things:

- How to make and use an Arena allocator
- How to make a Stack abstract data type (ADT)
- Define a binary tree using Go generics which uses an Arena style allocation strategy

So what is the point of all of this code? In part to demonstrate the suitability of Go as a systems programming language. To be a good systems programming language you need to have a fair amount of control over how memory is used and layed out in our programs. That Go is a garbage collected language means you have less memory control than say C/C++.

However since Go has support for proper pointers, value types and taking the address of fields and elements it allows you to define secondary allocators which gives you the ability to have similar kinds of control over memory as older language like C/C++.

Because Go now has generics it is also a good way to show making these allocators is made a lot easier with generics. Now we can write the code once for defining allocators for any type of object.

# Usage

There is no executable here. It is all about running tests and examples. Say you want to run the `TestArenaAllocation` test. You could write the following code (both lines are equivalent):

     ❯ go test -v -timeout 30s -run TestArenaAllocation .
     ❯ go test -v -timeout 30s -run TestArenaAllocation github.com/ordovician/arenatree
     
Or if you want to run all the tests you could write:

    ❯ go test -v -timeout 30s -run .
    
The `-timeout 30s` is obvously not necessary. It is just to prevent tests from hanging. The `-v` for verbose is useful at it shows you what tests actually got run.