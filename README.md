# Flower
### Easy to learn programming language

# Syntax
* Syntax very similair to `lisp`
* Supported tokens `string`, `number`, `etc (word)`, `comment`
* Supported types for eval `*golisp.Tag`, `*golisp.Value` (It's only one values to parse)
* Data types `number`, `string`, `iterator`, `dict`, `list`, `block`, `function`

* If's
```lisp
; If statement (if boolean (do ...) (do ...))
; Second (do ...) block is not mandatory, unless you want 'else' to process
(if (eq name 'Alan') (do
    (print "Oh, hi Alan")
) (do
    (print "Sad, cause you are not Alan")
))
```
* Functions
```lisp
; a, b, c is an arguments
(set add3
    (def a b c (do
        (return (add a (add b c)))
    )
))

; Call that function
(set summ (add3 10 20 30))
```
* Code blocks
```lisp
; Code block
(set codeblock (do
    ; Asks for name and text
    ; If such variables are not exist - error
    (require name text)
    (print (concat name ":" text))
    (return 123)
))

; Code block call
; Code block allows us to have DSL features.
(set result
    (call codeblock
        (with name "Ihor") ; Adds 'name' variable block
        (with text "Hello there!") ; Adds 'text' variable block
    )
)
```
* Lists
    * List is also iterator. So it could be iterated
```lisp
; Create List
(list.new)

; Create List from iterator
(list.new (of 10 20 30 40 50 60))

; Copy list
; Will make copy with all the values
(list.new otherlist)

; Get list element
(list.get list 10)

; Set list element
(list.set list 10 "Ihor")

; Add element to the list
(list.add list "Adam")

; Get length of the list
(len list)

; iterate over list elements
(iterate list elem (do
    (print elem)
))
```
* Dictionaries
```lisp
; Create dictionary
(set profile
    (dict.new
        (with "name" "Ihor")
        (with "age"  18)
        (with "city" "London")
    )
)

; Get value from dictionary by string
(dict.get profile "name")

; Get value from dictionary as variable
; You can use '.' to enter in deep. Like: (set name profile.portfolio.name)
(print profile.name)

; Set value to the dictionary
; Also you can use (set profile.name "Ihor") as well
(dict.set profile "name" "Ihor")

; Set value to the directory with '.'
; Let's say we want to change profile.cities.first = "Lviv"
(dict.set profile.cities "first" "Lviv")

; Iterate over dictionary keys
(iterate (dict.keys profile) key (do
    (print (concat key ":" (get profile key)))
))

; Get length of dictionary
(len profile)

; Extender. When value not found it will look in extending dict
(dict (extends profile))
```
* Boxes
```lisp
; Language has no any mechanism to assign values upper than it's scope
; But there is an exit. You could use (box) command.
; Box is a 1-key based dict. It's SUPER light and almost has no inpact to perfomance
; Key is "value". So you can change it via (set) command or get the value as from dict

; Create a box
(set b (box))

; Create a box with inital value
(set b (box 0))

; Set new value to the box
(set b.value "Test")

; Get value from the box
(print b.value)
```
* Iterators
```lisp
; Hardcoded Iterator
; Will create object which could be iterated
; Values are not evaluated yet until it iterates
(of 1 2 3 4 5 6)

; Function 'add' will be called ONLY when this will be iterating
(set addings
    (of (add 10 20) (add 40 50) (add 80 100))
)

; Iterator from range 0->100
(iterate (range 0 100) item (do
    (print item)
))

; If you want to poll items one by one
; It will take iterator and get iteration stream
(set iter
    (iteration (of 10 20 30 40))
)

; Poll items from iteration stream one by one
(next iter)

; Poll remain items one by one into a list
(next-all iter)

; Create generator
; It returns values and get counter increased
; If it will return nil value then iteration ends
(generator pos (do ...))

; Generator sample. Will stop if pos is more than 10
; This generator returns multiplied by 2 numbers
(set api.multiplied (generator pos (do
    (if (greater pos 10) (do
        (return nil)
    ))
    (return (mul pos 2))
)))
```
* Stacks
```lisp
; Create empty stack
(stack.new)

; Create stack with iterator
(stack.new (of 10 20 30))

; Push value. Let's assume 's' is a Stack variable
(stack.push s "Some value")

; Pop value. Let's assume 's' is a Stack variable
(stack.pop s)

; Get length of the stack
(len s)
```
* Loops
```lisp
; You can change range value
(iterate (range 0 1000) i (do
    ; Break if number is more than 100
    (if (greater i 100) (do (break)))
    (print i)
))

; Forever loop
(iterate (infinite) i (do
    ; ...
))

; (break) and (continue) are also working
```
* Strings
```lisp
; Iterate over each symbol
; Could be used with (iterate ...) command
(str.iterate "Hello!")

; Join list of strings with "+" symbol. You can set whatever you want
(str.join
    (list.new (of "Hi", "Jack", "Mary"))
    "+"
)

; Concatenation
(str.concat "A" "B" "C")

; Turn everything to string
(str.str obj)

; Get substring from->to
(str.sub "***Ihor*****" 3 7) ; Will return "Ihor"

; Split string by separator
; Will return list of splitted strings by symbol
(str.split "a,b,c" ",") ; will return list["a" "b" "c"]

; Split string by separator with limit
(str.split "a,b,c" "," 2) ; will return list["a" "b,c"]

; Finds first index of "b" in "abc"
; will return -1 if not found
(str.find "abc" "b") ; will return 1

; Get character at some position
; Will return "" if out of bounds
(str.at "abc" 1) ; will return "b"

; Replace substring to another
(str.rep "Ihaoer" "aoe" "o") ; will return "Ihor"

; Multiply strings
(str.mul "Hello" 3) ; Will return "HelloHelloHello"
```
* Time
```lisp
; Sleep one second
(time.sleep 1000)
```
* Assertion
```lisp
; Assert that 2+2=4
(assert (eq (add 2 2) 4))

; Assert with message
(assert (eq (add 2 2) 4) "2 + 2 should be 4")
```
* FileSystem
```lisp
; Import
; Importing file need to return something
; CAUTION: CYCLE dependencies could cause recursion and fail
;    better to have some utils or common file which not imports anything itself
(set utils (fs.import "utils.lsp"))

; Read file as string
(fs.read "test.txt")

; Write string to a file
(fs.write "test.txt" "Some value")

; Delete the file
(fs.delete "test.txt")

; Get list of files/folders
; returns list of dicts with [name, size, isfile]
(fs.list "./")
```
* Nullable operations
```lisp
; If bool is true then return result1, otherwise result2
(maybe bool result1 result2)

; If value is nil then return defaultVal, otherwise value itself
(nil-def value defaultVal)

; True if value is nil
(is-nil value)
```
* OOP
```lisp
; Super type (Class)
(set User (dict.new
    (with "name" "Default")
    (with "age" 0)
    (with "getName" (def self (do
        (return self.name)
    )))
    (with "getAge" (def self (do
        (return self.age)
    )))
))

; Create new user
(set usr1 (dict.new
    (extends User) ; Basic inheritanse
    (with "name" "Ihor")
    (with "age" 18)
))

(print (usr1.getName usr1))
(print (usr1.getAge usr1))
```

# Sample
* Real working lexer in `flower`
```lisp
; Function (parse text)
(set parse (def t (do
    (set arr (list.new))
    (set acc (dict.new (with "value" (list.new))))

    (iterate (str.iterate (str.concat t " ")) c (do
        (if (or (eq c " ") (eq c "\t")) (do
            (list.add arr (str.join acc.value ""))
            (set acc.value (list.new))
        ) (do
            (list.add acc.value c)
        ))
    ))

    (return arr)
)))

; Will return: LIST [[0] = hello, [1] = world, [2] = and, [3] = all]
(print (parse "hello world and all"))
```
