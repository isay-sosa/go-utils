# Go utils
Go-utils is a library written in Go with some functionality that can be helpful for others.

I've programmed in different langueges (Ruby, Java, etc), and always there's something that you want from a specific language to be implemented in other (methods such as `map`, `select` from Ruby, classes like `ArrayList` from Java), etc.

Here you will find these implementations for Go.

## ArrayList
In the Java world, arraylists are very common. Go utils implements an ArrayList struct, just like the ArrayList in Java.

    func main() {
        list := new(ArrayList)

        // Adding an element
        list.Add("Element")
        list.Add(true)
        list.Add(7)

        // Adding several elements
        list.Add("Foo", 20, false, nil, "Bar")

        list.Size() // => 8

        // Getting an element
        el, err := list.Get(1) // => true, nil
        el, err = list.Get(-1) // => nil, index out of range

        // Remove an element
        err := list.RemoveAt(1) // => nil
        list.Size()             // => 7

        err = list.RemoveAt(9)  // => index out of range
        err = list.Remove(true) // => element not found
        err = list.Remove(7)    // => nil
        list.Size()             // => 6

        // To Slice
        slice := list.Slice() // => ["Element", "Foo", 20, false, nil, "Bar"]
    }

## Slices functions
### Combination
This function is based on Ruby's `product` method. It receives several slices and combines all of them in a single slice.

    c, _ := Combination([]int{1, 2, 3}, []int{4, 5})
    // => [[1 4] [1 5] [2 4] [2 5] [3 4] [3 5]]

    c, _ = Combination([]int{1, 2}, []int{1, 2})
    // => [[1 1] [1 2] [2 1] [2 2]]

    c, _ = Combination([]int{1, 2}, []int{3, 4}, []int{5, 6})
    // => [[1 3 5] [1 3 6] [1 4 5] [1 4 6] [2 3 5] [2 3 6] [2 4 5] [2 4 6]]

    c, _ = Combination([]int{1, 2})
    // => [[1] [2]]

    c, _ = Combination([]int{})
    // => []

### Compact
This function is based on Ruby's `compact` method. It receives a slice and returns a new slice with all nil elements removed.

    slice := []interface{}{
        []int{1, 4, 5},
        nil,
        []int{2, 4, 6},
        []int{2, 3, 5},
    }

    c, _ := Compact(slice) // => [[1 4 5] [2 4 6] [2 3 5]]

### IsIncluded
This function is based on Ruby's `include?` method. It receives a slice and an object. It returns true if the object is included in the slice.

    slice := []string{"a", "b", "c", "d"}
    include, err := IsIncluded(slice, "c") // => true, nil

    include, err = IsIncluded(slice, "z")  // => false, element not found

### Map
This function is based on Ruby's `map` method. It receives a slice and a function (mapFunc) and returns a new slice containing the returned values by the mapFunc.

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

        names, _ := Map(users, func(obj interface{}) interface{} {
            return obj.(*User).Name
        })

        fmt.Println(names) // => [User 1, User 2, User 3]

        ids, _ := Map(users, func(obj interface{}) interface{} {
            return obj.(*User).ID
        })

        fmt.Println(ids) // => [1, 2, 3]
    }

### Select
This function is based on Ruby's `select` method. It receives a slice and a function (selectFunc) and returns a new slice containing the elements of which the selectFunc returns true.

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

        adults, _ := Select(users, func(obj interface{}) bool {
            return obj.(*User).Age >= 18
        })

        fmt.Println(adults) // => [&User{1, "User 1", 20}, &User{3, "User 3", 18}]

        teens, _ := Select(users, func(obj interface{}) bool {
            return obj.(*User).Age < 18
        })

        fmt.Println(teens) // => [&User{2, "User 2", 16}]
    }