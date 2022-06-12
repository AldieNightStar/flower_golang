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
* Dictionaries
```lisp
; Create dictionary
(set profile
    (dict
        (with name "Ihor")
        (with age 18)
        (with city "London")
    )
)

; Get value from dictionary
(print (dict-get profile "name"))

; Iterate over dictionary keys
(iterate (dict-keys profile) key (do
    (print (concat key ":" (get profile key)))
))

; Get length of dictionary
(print (dict-len profile))
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
```
* Loops
```lisp
; You can change range value
(iterate (range 0 1000) i (do
    ; Break if number is more than 100
    (if (greater i 100) (do (break)))
    (print i)
))
```