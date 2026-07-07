package main

import "fmt"

func main() {
    dictionary := make(map[int]string)

    // Adding elements (set operation)
    dictionary[1] = "John"
    dictionary[2] = "Mike"
    dictionary[3] = "Emma"

    // Retrieving an element
    fmt.Println("Element with key 1:", dictionary[1])
    // Output: Element with key 1: John

    // Removing an element
    delete(dictionary, 2)

    fmt.Println("Dictionary after removal operation:")
    for key, value := range dictionary {
        fmt.Printf("%d: %s\n", key, value)
    }
    // Output: Dictionary after removal operation: 1: John, 3: Emma
}