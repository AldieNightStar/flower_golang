((set b (do
    (print name)
    (print age)
    (return 1)
))
(print (call b
    (with "name" "Ihor")
    (with "age" 18)
)))