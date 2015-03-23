# Go utils
This is a utils library for Go with some functions and structs from other languages (like Ruby, Java) that don't come in the Go core and can be helpful.

## ArrayList
In the Java world, arraylists are very common. Go utils implements an ArrayList struct, just like the ArrayList in Java.

## Map
This function is based on Ruby's map. It receives a slice and a function (mapFunc) and returns a new slice containing the returned values by the mapFunc.

    type User struct {
        ID   int
        Name string
    }

    func main() {
        users := []*User{
            &User{1, "User 1"},
            &User{2, "User 2"},
            &User{3, "User 3"},
        }

        names := Map(users, func(obj interface{}) interface{} {
            return obj.(*User).Name
        })

        fmt.Println(names) // => [User 1, User 2, User 3]

        ids := Map(users, func(obj interface{}) interface{} {
            return obj.(*User).ID
        })

        fmt.Println(ids) // => [1, 2, 3]
    }

## Select
This function is based on Ruby's select. It receives a slice and a function (selectFunc) and returns a new slice containing the elements of which the selectFunc returns true.

    type User struct {
        ID   int
        Name string
        Age  int
    }

    func main() {
        users := []*User{
            &User{1, "User 1", 20},
            &User{2, "User 2", 16},
            &User{3, "User 3", 18},
        }

        adults := Select(users, func(obj interface{}) bool {
            return obj.(*User).Age >= 18
        })

        fmt.Println(adults) // => [&User{1, "User 1", 20}, &User{3, "User 3", 18}]

        teens := Select(users, func(obj interface{}) bool {
            return obj.(*User).Age < 18
        })

        fmt.Println(teens) // => [&User{2, "User 2", 16}]
    }