# Flower
### Easy to learn programming language

# Syntax
* Syntax very similair to `lisp`
* Supported tokens `string`, `number`, `etc (word)`, `comment`
* Supported types for eval `*golisp.Tag`, `*golisp.Value` (It's only one values to parse)
* Data types `number`, `string`, `iterator`, `dict`, `list`, `block`, `function`
```lisp

; Loop (repeat alias max (do ...))
(repeat count 100 (do
    ; Print the count
    (print (str count))
    ; If count > 10 then break
    (if (gt count 10) (do
        ; Will endup all the loop
        (break)
    ))
))

; If statement (if boolean (do ...) (do ...))
; Second (do ...) block is not mandatory, unless you want 'else' to process
(if (eq name 'Alan') (do
    (print "Oh, hi Alan")
) (do
    (print "Sad, cause you are not Alan")
))

; Functions
(set add3
    (def a b c (do
        (return (add a (add b c)))
    )
))

; Function call
(set summ (add3 10 20 30))

; Code blocks
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

; Dictionaries
(set profile
    (dict
        (with name "Ihor")
        (with age 18)
        (with city "London")
    )
)

; Get value from dictionary
(print (get profile "name"))

; Get all keys as iterator from Dictionary
(iterate (keys profile) key (do
    (print (concat
        key ":" (get profile key)
    ))
))

; Get length of array, or dictionary
(print (len profile))

; Functional stuff
; ------------------

; Iterator
; Will create object which could be iterated
; Values are not evaluated yet until it iterates
(of 1 2 3 4 5 6)

; Function 'add' will be called ONLY when this will be iterating
(set addings
    (of (add 10 20) (add 40 50) (add 80 100))
)

; Iteration
; Works ONLY with iterators
; To iterate list, dicts and so on, you need to create iterators of them
(iterate addings item (do
    (print item)
))

; Mapping
; It will create ALTERNATIVE iterator
; Each element will be multiplied by 2
(map addings item (do
    (return (mul item 2))
))

; Filtering
; Will create ALTERNATIVE iterator
; Elements with false will not be computed
(filter addings item (do
    (return (eq (mod item 5) 0))
))

; ---------------------

; Codeblocks could be replaced or reused
; That way no need to write the same code
(set onlyEven (do
    (require item) ; error if such variable is nil
    (return (eq (mod item 2) 0))
))

(set addings (filter addings item onlyEven))

; Iterate over addings(filtered) and put everythin to list
(set numbers (toList addings))

; Convert list to iterator
; And iterate over elements
(iterate (range numbers) item (do
    (print item)
))

; Iterator from range 0->Max
(iterate (range 0 100) item (do
    (print item)
))
```